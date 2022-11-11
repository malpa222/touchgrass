package p2p

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"net"
	"touchgrass/client/p2p/bitfield"
	"touchgrass/client/p2p/handshake"
	"touchgrass/client/p2p/message"
)

type P2P struct {
	handshake handshake.Handshake
	peer      Peer

	Connection net.Conn
	Chocked    bool
	Bitfield   bitfield.Bitfield
}

func New(peerId [20]byte, infoHash [20]byte, peer Peer) (p2p *P2P, err error) {
	conn, err := net.Dial("tcp", peer.String())
	if err != nil {
		return nil, err
	}

	hs := handshake.Handshake{
		InfoHash: infoHash,
		PeerID:   peerId,
	}
	if err := shakeHands(hs, conn); err != nil {
		return nil, err
	}

	bf, err := getBitfield(conn)
	if err != nil {
		return nil, err
	}

	return &P2P{
		handshake: hs,
		peer:      peer,

		Connection: conn,
		Chocked:    true,
		Bitfield:   bf,
	}, nil
}

// Read reads incoming messages
func (p *P2P) Read(r io.Reader) (message message.Message, err error) {
	return
}

// SendRequest sends a piece request to the peer
func (p *P2P) SendRequest(index int) {

}

func shakeHands(hs handshake.Handshake, conn net.Conn) error {
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

	// check if handshakes match
	hs2, err := handshake.Read(bytes.NewReader(temp[:]))
	if hs2 == nil {
		return errors.New("received no handshake")
	} else if hs.InfoHash != hs2.InfoHash {
		return errors.New("the handshake is invalid")
	}

	return nil
}

func getBitfield(conn net.Conn) (bitfield.Bitfield, error) {
	msg, err := message.Read(conn)
	if err != nil {
		return bitfield.Bitfield{}, err
	}

	if msg.MessageId != message.MsgBitfield {
		return bitfield.Bitfield{}, errors.New(fmt.Sprintf("Expected bitfield, received: %v", msg.MessageId))
	}

	return bitfield.Bitfield{Data: msg.Payload}, nil
}
