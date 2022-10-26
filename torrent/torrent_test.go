package torrent

import (
	"testing"
)

const PATH = "./debian.torrent"

func TestParseTorrent(t *testing.T) {
	if torr, err := ParseTorrent(PATH); err != nil || torr == nil {
		t.Errorf("Error:\n%v", err)
	}
}
