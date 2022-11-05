package client

import (
	"errors"
	"net"
	"touchgrass/client/bitfield"
	"touchgrass/client/handshake"
	"touchgrass/client/tracker"
)

// Client is a tcp connection with peer
type Client struct {
	peer     tracker.Peer
	infoHash [20]byte
	peerId   [20]byte

	Conn     net.Conn
	Choked   bool
	Bitfield bitfield.Bitfield
}

func New(peer tracker.Peer, infoHash [20]byte, peerId [20]byte) *Client {
	return &Client{
		peer:     peer,
		infoHash: infoHash,
		peerId:   peerId,
	}
}

func (c *Client) Connect(peer *tracker.Peer) error {
	conn, err := net.Dial("tcp", peer.String())
	if err != nil {
		return err
	}
	defer conn.Close()

	if err := c.shakeHands(conn); err != nil {
		return err
	}

	return nil
}

func (c *Client) shakeHands(conn net.Conn) error {
	hs := &handshake.Handshake{
		PeerID:   c.peerId,
		InfoHash: c.infoHash,
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
