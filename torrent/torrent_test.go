package torrent

import (
	"testing"
)

const TORRENT = `d8:announce41:http://bttracker.debian.org:6969/announce7:comment35:"Debian CD from cdimage.debian.org"10:created by13:mktorrent 1.113:creation datei1662813552e4:infod6:lengthi400556032e4:name31:debian-11.5.0-amd64-netinst.iso12:piecelength i262144e6:pieces26:(huge binary blob of data)ee`
const PATH = "./debian.torrent"

func TestParseTorrent(t *testing.T) {
	if _, err := ParseTorrent(PATH); err != nil {
		t.Errorf("Error:\n%v", err)
	}
}
