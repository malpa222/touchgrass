package torrent

import (
	"crypto/sha1"
	"errors"
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

func ParseTorrent(path string) (*Torrent, error) {
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	_, rawTorrent := bencode.Decode(buf)
	rawInfo, ok := rawTorrent["info"].(bencode.Dictionary)
	if !ok {
		return nil, errors.New("missing the info dictionary")
	}

	infoBytes, err := bencode.ToBytes(rawInfo)
	if err != nil {
		return nil, err
	}

	return &Torrent{
		InfoHash:     sha1.Sum(*infoBytes),
		Announce:     rawTorrent["announce"].(string),
		CreatedBy:    rawTorrent["created by"].(string),
		CreationDate: rawTorrent["creation date"].(int),
		PieceHashes:  splitPieces(rawInfo["pieces"].(string)),
		PieceLength:  rawInfo["piece length"].(int),
	}, nil
}

func splitPieces(data string) [][20]byte {
	var chunks [][20]byte

	for i := 0; i < len(data); i += 20 {
		temp := []byte(data[i : i+20])
		chunks = append(chunks, *(*[20]byte)(temp))
	}

	return chunks
}
