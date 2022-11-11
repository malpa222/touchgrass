package client

import (
	"touchgrass/client/p2p"
	t "touchgrass/torrent"
)

type client struct {
	peerId  [20]byte
	torrent *t.Torrent

	queue   chan workPiece
	results chan *Piece
}

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
	c := &client{
		peerId:  peerId,
		torrent: torrent,
		queue:   make(chan workPiece, torrent.PieceLength),
		results: make(chan *Piece),
	}

	peers, err := p2p.GetPeers(peerId, torrent)
	if err != nil {
		return "", err
	}

	// populate the work channel with pieces
	for i, hash := range torrent.PieceHashes {
		c.queue <- workPiece{
			Index: i,
			Hash:  hash,
		}
	}

	for _, peer := range *peers {
		go startWorker(c, peer)
	}

	return "", nil
}

func startWorker(client *client, peer p2p.Peer) error {
	// connect to a peer
	p, err := p2p.New(client.peerId, client.torrent.InfoHash, peer)
	if err != nil {
		return err
	}
	defer p.Connection.Close()

	// send unchoke
	// send interested
	// check if bitfield has the piece

	return nil
}
