package client

import (
	"touchgrass/client/p2p"
	t "touchgrass/torrent"
)

type workPiece struct {
	Index int
	Hash  [20]byte
}

type Piece struct {
	Index int
	Data  []byte
	Hash  [20]byte
}

func Download(peerId [20]byte, torrent *t.Torrent) (string, error) { // TODO decide on parameters and return values
	// initialize work and results channels
	workChan := make(chan *workPiece, torrent.PieceLength)
	resultChan := make(chan *Piece)

	peers, err := p2p.GetPeers(peerId, torrent)
	if err != nil {
		return "", err
	}

	// populate the work channel with pieces
	for i, hash := range torrent.PieceHashes {
		workChan <- &workPiece{
			Index: i,
			Hash:  hash,
		}
	}

	for _, peer := range *peers {
		go startWorker(&peer, workChan, resultChan)
	}

	return "", nil
}

func Upload(torrent *t.Torrent) (path string, err error) { // TODO implement uploading
	return
}

func startWorker(peer *p2p.Peer, queue chan *workPiece, results chan *Piece) error {
	return nil
}
