package tracker

import (
	"testing"
	"touchgrass/torrent"
)

func TestTrackerResponse(t *testing.T) {
	torr, err := torrent.ParseTorrent("./bencode/debian.torrent")
	if err != nil {
		return
	}

	req := TrackerReq{}
}
