package tracker

import (
	"math/rand"
	"net/http"
	"net/url"
	"strconv"
	"time"
	t "touchgrass/torrent"
)

type Event byte

const (
	Started Event = iota
	Stopped
	Empty
	Completed
)

type TrackerReq struct {
	Port       int
	Uploaded   int
	Downloaded int
	Left       int
	Event      Event
	Compact    bool
}

func GetPeers(torrent *t.Torrent, req *TrackerReq) (*http.Response, error) {
	trackerUrl, err := buildUrl(torrent, req)
	if err != nil {
		return nil, err
	}

	resp, err := http.Get(trackerUrl)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func buildUrl(torrent *t.Torrent, req *TrackerReq) (string, error) {
	base, err := url.Parse(torrent.Announce)
	if err != nil {
		return "", err
	}

	compact := "0"
	if req.Compact {
		compact = "1"
	}

	params := url.Values{
		"info_hash":  []string{string(torrent.InfoHash[:])},
		"peer_id":    []string{createPeerId()},
		"port":       []string{strconv.Itoa(req.Port)},
		"uploaded":   []string{strconv.Itoa(req.Uploaded)},
		"downloaded": []string{strconv.Itoa(req.Downloaded)},
		"left":       []string{strconv.Itoa(req.Left)},
		"compact":    []string{compact},
	}

	base.RawQuery = params.Encode()
	return base.String(), nil
}

func createPeerId() string {
	rand.Seed(time.Now().UnixNano())
	charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPRSTUWQXYZ1234567890"

	id := make([]byte, 20)
	for i := 0; i < 20; i++ {
		id[i] = charset[rand.Intn(len(charset))]
	}

	return string(id)
}
