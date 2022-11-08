package client

import (
	"touchgrass/client/tracker"
	t "touchgrass/torrent"
)

type Piece struct {
	Index int
	Data  []byte
	Hash  [20]byte
}

func Download(peerId [20]byte, torrent *t.Torrent) (string, error) { // TODO decide on parameters and return values
	// initialize work and results channels
	workChan := make(chan *Piece, torrent.PieceLength)
	resultChan := make(chan *Piece)

	peers, err := tracker.GetPeers(peerId, torrent)
	if err != nil {
		return "", err
	}

	// populate the work channel with pieces
	for i, hash := range torrent.PieceHashes {
		workChan <- &Piece{
			Index: i,
			Hash:  hash,
		}
	}

	for _, peer := range *peers {
		workChan := make(chan Piece)
	}

	return "", nil
}

func Upload(torrent *t.Torrent) (path string, err error) { // TODO implement uploading
	return
}

// generates a slice of empty pieces that need to be downloaded
func genWorkQueue(torrent t.Torrent) (pieces []Piece) {
	for i, p := range torrent.PieceHashes {
		pieces[i] = Piece{Hash: p}
	}

	return
}
