package client

import (
	"errors"
	"log"
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

func (c *Client) Connect(peer *tracker.Peer) (test int, err error) {
	conn, err := net.Dial("tcp", peer.String())
	if err != nil {
		return
	}

	defer conn.Close()

	hs := &handshake.Handshake{
		PeerID:   c.peerId,
		InfoHash: c.infoHash,
	}

	// send the handshake to the peer
	_, err = conn.Write(hs.Serialize())
	if err != nil {
		return
	}

	// read the incoming handshake
	var temp [68]byte
	_, err = conn.Read(temp[:])
	if err != nil {
		return
	}

	hs2, err := handshake.Deserialize(temp[:])
	if hs2 == nil {
		return 0, errors.New("received no handshake")
	} else if hs.InfoHash != hs2.InfoHash {
		return 0, errors.New("the handshake is invalid")
	}

	log.Printf("their handshake:\n %#v", hs2)

	return
}
