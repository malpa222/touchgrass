package torrent

import (
	"crypto/sha1"
	"io/ioutil"
	"touchgrass/torrent/bencode"
)

type Torrent struct {
	Announce     []string
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

	// get the underlying value and assert if the file is using correct format
	_, decoded := bencode.Decode(buf)
	decTorrent, err := getDict(decoded)
	if err != nil {
		return nil, err
	}

	decInfo, err := getDict(decTorrent["info"])
	if err != nil {
		return nil, err
	}

	announceList, exists := getList(decTorrent["announce-list"])
	if exists {
		for _, url := range announceList {

		}
	} else {

	}

	hash, err := createInfoHash(decInfo)
	if err != nil {
		return nil, err
	}

	return &Torrent{
		InfoHash:     hash,
		Announce:     decTorrent["announce"].(string),
		CreatedBy:    decTorrent["created by"].(string),
		CreationDate: decTorrent["creation date"].(int),
		PieceHashes:  splitPieces(decInfo["pieces"].(string)),
		PieceLength:  decInfo["piece length"].(int),
	}, nil
}

func createInfoHash(info map[string]any) ([20]byte, error) {
	if encoded, err := bencode.Encode(info); err != nil {
		return [20]byte{}, err
	} else {
		return sha1.Sum([]byte(encoded)), nil
	}
}

func splitPieces(data string) [][20]byte {
	var chunks [][20]byte

	for i := 0; i < len(data); i += 20 {
		temp := []byte(data[i : i+20])
		chunks = append(chunks, *(*[20]byte)(temp))
	}

	return chunks
}
