package torrent

import (
	"bytes"
	"encoding/gob"
	"io/ioutil"
	"touchgrass/torrent/bencode"
)

type Torrent struct {
	Announce     string
	InfoHash     [20]byte
	PieceHashes  [][20]byte
	PieceLength  int
	CreatedBy    string
	CreationDate int
}

type torrentInfo struct {
	Name        string
	PieceLength int
	Pieces      string
	Length      int
	Path        []string
}

func ParseTorrent(path string) (*Torrent, error) {
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	_, rawTorrent := bencode.GetDict(buf)

	torrent := &Torrent{
		Announce:     rawTorrent["announce"].(string),
		CreatedBy:    rawTorrent["created by"].(string),
		CreationDate: rawTorrent["creation date"].(int),
	}

	if rawInfo, ok := rawTorrent["info"]; ok {
		rawInfo := rawInfo.(bencode.Dictionary)

		// copy the rawInfo map
		var temp bencode.Dictionary
		for k, v := range rawInfo {
			temp[k] = v
		}

		var hashBuf bytes.Buffer
		enc := gob.NewEncoder(&hashBuf)
		if err := enc.Encode(temp); err != nil {

		}

		//torrent.Info = &Info{
		//	Name:        rawInfo["name"].(string),
		//	PieceLength: rawInfo["piece length"].(int),
		//	Pieces:      rawInfo["pieces"].(string),
		//	Length:      rawInfo["length"].(int),
		//}
	}

	return torrent, nil
}
