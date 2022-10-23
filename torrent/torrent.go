package torrent

import (
	"io/ioutil"
	"touchgrass/torrent/bencode"
)

type Torrent struct {
	Announce     string
	Comment      string
	CreatedBy    string
	CreationDate int
	Info         *Info
}

type Info struct {
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
		Comment:      rawTorrent["comment"].(string),
		CreatedBy:    rawTorrent["created by"].(string),
		CreationDate: rawTorrent["creation date"].(int),
	}

	if rawInfo, ok := rawTorrent["info"]; ok {
		rawInfo := rawInfo.(bencode.Dictionary)

		torrent.Info = &Info{
			Name:        rawInfo["name"].(string),
			PieceLength: rawInfo["piece length"].(int),
			Pieces:      rawInfo["pieces"].(string),
			Length:      rawInfo["length"].(int),
		}
	}

	return torrent, nil
}
