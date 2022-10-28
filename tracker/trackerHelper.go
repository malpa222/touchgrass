package tracker

import (
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"math/rand"
	"net"
	"net/http"
	"net/url"
	"strconv"
	"time"
	t "touchgrass/torrent"
	"touchgrass/torrent/bencode"
)

type Req struct {
	Port       int
	Uploaded   int
	Downloaded int
	Left       int
	Compact    bool
}

type Resp struct {
	Failure  string
	Interval int
	Peers    []*Peer
}

type Peer struct {
	IP   net.IP
	Port uint16
}

func GetPeersHTTP(torrent *t.Torrent, req *Req) (*Resp, error) {
	trackerUrl, err := buildUrl(torrent, req)
	if err != nil {
		return nil, err
	}

	resp, err := http.Get(trackerUrl)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if decodedResp, err := unmarshalResponse(body); err != nil {
		return nil, err
	} else {
		return decodedResp, nil
	}
}

func buildUrl(torrent *t.Torrent, req *Req) (string, error) {
	base, err := url.Parse(torrent.Announce)
	if err != nil {
		return "", err
	}

	// generate a peer id
	rand.Seed(time.Now().UnixNano())
	chars := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPRSTUWQXYZ1234567890"
	id := make([]byte, 20)
	for i := 0; i < 20; i++ {
		id[i] = chars[rand.Intn(len(chars))]
	}

	params := url.Values{
		"info_hash":  []string{string(torrent.InfoHash[:])},
		"peer_id":    []string{string(id)},
		"port":       []string{strconv.Itoa(req.Port)},
		"uploaded":   []string{strconv.Itoa(req.Uploaded)},
		"downloaded": []string{strconv.Itoa(req.Downloaded)},
		"left":       []string{strconv.Itoa(req.Left)},
		"compact":    []string{"1"},
	}

	base.RawQuery = params.Encode()
	return base.String(), nil
}

func unmarshalResponse(body []byte) (*Resp, error) {
	_, decoded := bencode.Decode(body)
	var temp map[string]any

	// expecting a dictionary
	switch v := decoded.(type) {
	case map[string]any:
		temp = v
	default:
		return nil, errors.New(fmt.Sprintf("got an incorrect response from server:\n%v", v))
	}

	// first check if the server returned failure
	if failure, hasFailure := temp["failure reason"].(string); hasFailure {
		return &Resp{Failure: failure}, nil
	}

	// then check if it has returned peer list at all
	var peersBlob []byte
	if blob, hasPeers := temp["peers"]; !hasPeers {
		return nil, errors.New("missing peer list")
	} else {
		peersBlob = []byte(blob.(string))
	}

	// unmarshall the peer list from bytes to Peer struct
	// according to bep_0023, a peer in compact form consists of:
	// 4 bytes -> IP address
	// 2 bytes -> port
	peerSize := 6
	if len(peersBlob)%peerSize != 0 {
		return nil, errors.New(fmt.Sprintf("received incorrect peer list:\n%v", peersBlob))
	}

	numPeers := len(peersBlob) / peerSize
	peers := make([]*Peer, numPeers)
	for i := 0; i < numPeers; i++ {
		offset := i * peerSize

		peers[i] = &Peer{
			IP:   peersBlob[offset : offset+4],
			Port: binary.BigEndian.Uint16(peersBlob[offset+4 : offset+6]),
		}
	}

	// once peers are processed, check if the tracker returned the interval
	interval, hasInterval := temp["interval"].(int)
	if !hasInterval {
		return nil, errors.New("received incorrect response: missing interval")
	}

	return &Resp{
		Failure:  "",
		Interval: interval,
		Peers:    nil,
	}, nil
}
