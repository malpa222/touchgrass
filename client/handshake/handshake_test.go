package handshake

import (
	"bytes"
	"testing"
	"touchgrass/torrent"
)

func TestSerialize(t *testing.T) {
	if torr, err := torrent.ParseTorrent("../../torrent/testfile.torrent"); err != nil {
		t.Errorf("got an error:\n%v", err)
	} else {
		hs := &Handshake{
			InfoHash: torr.InfoHash,
			PeerID:   torr.InfoHash,
		}

		ser := Serialize(hs)
		if len(*ser) != 68 {
			t.Errorf("serialization went wrong, expected 68 byte long array, got:\n%v", ser)
		}
	}
}

func TestDeserialize(t *testing.T) {
	if torr, err := torrent.ParseTorrent("../../torrent/testfile.torrent"); err != nil {
		t.Errorf("got an error:\n%v", err)
	} else {
		hs := &Handshake{
			InfoHash: torr.InfoHash,
			PeerID:   torr.InfoHash,
		}

		out := Serialize(hs)
		if len(*out) != 68 {
			t.Errorf("serialization went wrong, expected 68 byte long array, got:\n%v", err)
		}

		temp := bytes.NewReader(*out)
		if deser, err := Deserialize(temp); err != nil {
			t.Errorf("got an error:\n%v", err)
		} else if *deser != *hs {
			t.Errorf("data doesn't match\nexpected:%v\ngot:%v", hs, deser)
		}
	}
}
