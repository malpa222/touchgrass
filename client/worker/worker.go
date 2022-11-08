package dispatch

import (
	"errors"
	"net"
	"touchgrass/client/handshake"
	"touchgrass/client/tracker"
)

// a worker is a wrapper over peer which handles the
type worker struct {
	peerID   [20]byte
	infoHash [20]byte
}

func New(peer tracker.Peer) *worker {

}

func (c *Client) Connect(peer *tracker.Peer) error {
	conn, err := net.Dial("tcp", peer.String())
	if err != nil {
		return err
	}

	if err := c.shakeHands(conn); err != nil {
		return err
	}

	// Assign the connection to peer
	c.conn = conn
	defer conn.Close()

	return nil
}

func (w *worker) shakeHands(conn net.Conn) error {
	hs := &handshake.Handshake{
		PeerID:   w.peerID,
		InfoHash: w.infoHash,
	}

	// send the handshake to the peer
	_, err := conn.Write(hs.Serialize())
	if err != nil {
		return err
	}

	// read the incoming handshake
	var temp [68]byte
	_, err = conn.Read(temp[:])
	if err != nil {
		return err
	}

	hs2, err := handshake.Deserialize(temp[:])
	if hs2 == nil {
		return errors.New("received no handshake")
	} else if hs.InfoHash != hs2.InfoHash {
		return errors.New("the handshake is invalid")
	}

	return nil
}

func (c *Client) readMessage() error {
	if c.conn == nil {
		return errors.New("the the client is not connected to a peer")
	}

	return nil
}
