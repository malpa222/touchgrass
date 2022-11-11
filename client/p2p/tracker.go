package p2p

import (
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"net/url"
	"strconv"
	t "touchgrass/torrent"
	"touchgrass/torrent/bencode"
)

type request struct {
	announce   string
	peerId     [20]byte
	infoHash   [20]byte
	port       int
	uploaded   int
	downloaded int
	left       int
}

type response struct {
	failure  string
	interval int
	peers    *[]Peer
}

type Peer struct {
	IP   net.IP
	Port uint16
}

func (p *Peer) String() string {
	return net.JoinHostPort(p.IP.String(), strconv.Itoa(int(p.Port)))
}

func GetPeers(peerId [20]byte, torrent *t.Torrent) (peers *[]Peer, err error) {
	req := &request{
		announce:   torrent.Announce[0], // TODO
		peerId:     peerId,
		infoHash:   torrent.InfoHash,
		port:       6881,
		uploaded:   0,
		downloaded: 0,
		left:       0,
	}

	trackerUrl, err := buildUrl(req)
	if err != nil {
		return
	}
	res, err := http.Get(trackerUrl)
	if err != nil {
		return
	}
	log.Printf("tracker returned: %d", res.StatusCode)
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return
	}

	if decoded, err := unmarshalResponse(body); err != nil {
		return nil, errors.New(decoded.failure)
	} else if decoded.peers == nil {
		return nil, errors.New("received empty peer list")
	} else {
		return decoded.peers, nil
	}
}

func buildUrl(req *request) (string, error) {
	base, err := url.Parse(req.announce)
	if err != nil {
		return "", err
	}

	params := url.Values{
		"info_hash":  []string{string(req.infoHash[:])},
		"peer_id":    []string{string(req.peerId[:])},
		"port":       []string{strconv.Itoa(req.port)},
		"uploaded":   []string{strconv.Itoa(req.uploaded)},
		"downloaded": []string{strconv.Itoa(req.downloaded)},
		"left":       []string{strconv.Itoa(req.left)},
		"compact":    []string{"1"},
	}

	base.RawQuery = params.Encode()
	return base.String(), nil
}

func unmarshalResponse(body []byte) (res *response, err error) {
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
		return &response{failure: failure}, nil
	}

	// then check if it has returned peer list at all and if it's not empty
	var peersBlob []byte
	if blob, hasPeers := temp["peers"]; !hasPeers {
		return nil, errors.New("missing peer list")
	} else {
		peersBlob = []byte(blob.(string))
	}

	/* unmarshall the peer list from bytes to Peer struct
	according to bep_0023, a peer in compact form consists of:
	4 bytes -> IP address
	2 bytes -> port */
	peerSize := 6
	if len(peersBlob)%peerSize != 0 {
		return nil, errors.New(fmt.Sprintf("received incorrect peer list:\n%v", peersBlob))
	}

	numPeers := len(peersBlob) / peerSize
	peers := make([]Peer, numPeers)
	for i := 0; i < numPeers; i++ {
		offset := i * peerSize

		peers[i] = Peer{
			IP:   peersBlob[offset : offset+4],
			Port: binary.BigEndian.Uint16(peersBlob[offset+4 : offset+6]),
		}
	}

	// once peers are processed, check if the tracker returned the interval
	interval, hasInterval := temp["interval"].(int)
	if !hasInterval {
		interval = 15
	}

	return &response{
		interval: interval,
		peers:    &peers,
	}, nil
}
