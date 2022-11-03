package tracker

import (
	"fmt"
	"testing"
	"touchgrass/torrent"
)

func TestTrackerResponse(t *testing.T) {
	torr, err := torrent.ParseTorrent("../torrent/testfile.torrent")
	if err != nil {
		t.Errorf("%v", err)
		return
	}

	req := &Req{
		Port: 6881,
	}

	peers, err := GetPeers(torr, req)
	if err != nil {
		t.Errorf("%v", err)
		return
	}

	fmt.Printf("%v", peers)
}
