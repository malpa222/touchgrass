package client

import (
	"errors"
	"net"
	"time"
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

func (c *Client) connectToPeer() (test int, err error) {
	conn, err := net.DialTimeout("tcp", c.peer.String(), 3*time.Second)
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

	// send the handshake to the peer
	var temp [40]byte
	_, err = conn.Read(temp[:])
	if err != nil {
		return
	}

	hs2, err := handshake.Deserialize(temp[:])
	if hs.InfoHash != hs2.InfoHash {
		return 0, errors.New("the handshake is invalid")
	}

	return
}
