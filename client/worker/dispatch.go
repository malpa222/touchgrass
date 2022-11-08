package dispatch

import (
	"touchgrass/client"
	"touchgrass/torrent"
)

type piece struct {
}

type dispatch struct {
	torrent    torrent.Torrent
	queue      map[string][]byte
	pieces     []*piece
	downloaded []*piece
	workers    []*worker
}

func start(t torrent.Torrent) *dispatch {
	return &dispatch{
		torrent: t,
		queue:   make(map[string][]byte, 0), // TODO
	}
}

func (d *dispatch) newWorker(c *client.Client) {
}

// gets a piece from the queue and passes it to a worker
func getPiece() *piece {
	return nil
}
