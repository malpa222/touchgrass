package torrent

import (
	"testing"
)

const PATH = "./testfile.torrent"

func TestParseTorrent(t *testing.T) {
	torr, err := ParseTorrent(PATH)

	if err != nil || torr == nil {
		t.Errorf("Error:\n%v", err)
	}

	println(torr)
}
