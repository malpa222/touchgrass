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

	peers, err := GetPeers(torr)
	if err != nil {
		t.Errorf("%v", err)
		return
	}

	fmt.Printf("%v", peers)
}
