package torrent

import (
	"crypto/sha1"
	"errors"
	"io/ioutil"
	"touchgrass/cast"
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
	decTorrent, err := cast.ToDict[string, any](decoded)
	if err != nil {
		return nil, err
	}

	decInfo, err := cast.ToDict[string, any](decTorrent["info"])
	if err != nil {
		return nil, err
	}

	/* announce-list is a list of lists with trackers, grouped by the protocol
	for example: [[https://url1.com, https://url2.com], [udp://url3.com]]
	it is preferred over announce key */
	var announceList []string
	if temp, hasKey := decTorrent["announce-list"]; hasKey {
		temp, err := cast.ToList[any](temp)
		if err != nil {
			return nil, err
		}

		for _, tier := range temp { // iterate over tier list
			urlList, err := cast.ToList[any](tier)
			if err != nil {
				return nil, err
			}

			for _, url := range urlList { // iterate over url list
				url, err := cast.To[string](url)
				if err != nil {
					return nil, err
				}

				announceList = append(announceList, url)
			}
		}
	} else if temp, hasKey := decTorrent["announce"]; hasKey {
		if url, err := cast.To[string](temp); err != nil {
			return nil, err
		} else {
			announceList = append(announceList, url)
		}
	} else {
		return nil, errors.New("missing announce key")
	}

	hash, err := createInfoHash(decInfo)
	if err != nil {
		return nil, err
	}

	return &Torrent{
		InfoHash:     hash,
		Announce:     announceList,
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
