package p2p

import (
	"errors"
	"fmt"
	"net"
	"touchgrass/client/p2p/bitfield"
	"touchgrass/client/p2p/handshake"
	"touchgrass/client/p2p/message"
)

type P2P struct {
	handshake handshake.Handshake
	peer      Peer

	Conn     net.Conn
	Chocked  bool
	Bitfield bitfield.Bitfield
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

		Conn:     conn,
		Chocked:  true,
		Bitfield: bf,
	}, nil
}

// ReadIncoming reads incoming messages
func (p *P2P) ReadIncoming() (message message.Message, err error) {
	var buf []byte
	if _, err := p.Conn.Read(buf); err != nil {
		return message, err
	}

	return
}

func (p *P2P) SendUnchoke() error {
	if err := p.SendMsg(message.Message{MessageId: message.MsgUnchoke}); err != nil {
		return err
	}

	return nil
}

func (p *P2P) SendChoke() error {
	if err := p.SendMsg(message.Message{MessageId: message.MsgChoke}); err != nil {
		return err
	}

	return nil
}

func (p *P2P) SendInterested() error {
	if err := p.SendMsg(message.Message{MessageId: message.MsgInterested}); err != nil {
		return err
	}

	return nil
}

func (p *P2P) SendNotInterested() error {
	if err := p.SendMsg(message.Message{MessageId: message.MsgNotInterested}); err != nil {
		return err
	}

	return nil
}

func (p *P2P) SendMsg(msg message.Message) error {
	serialized := msg.Serialize()
	if num, err := p.Conn.Write(serialized); err != nil {
		return err
	} else if num != len(serialized) {
		return errors.New("unable to send whole message. aborting... ")
	}

	return nil
}

func shakeHands(hs handshake.Handshake, conn net.Conn) error {
	// send the handshake to the peer
	_, err := conn.Write(hs.Serialize())
	if err != nil {
		conn.Close()
		return err
	}

	// read the incoming handshake
	var temp [68]byte
	_, err = conn.Read(temp[:])
	if err != nil {
		conn.Close()
		return err
	}

	// check if handshakes match
	hs2, err := handshake.Deserialize(temp[:])
	if hs2 == nil {
		conn.Close()
		return errors.New("received no handshake")
	} else if hs.InfoHash != hs2.InfoHash {
		conn.Close()
		return errors.New("the handshake is invalid")
	}

	return nil
}

func getBitfield(conn net.Conn) (bitfield.Bitfield, error) {
	var temp []byte
	if _, err := conn.Read(temp[:]); err != nil {
		return bitfield.Bitfield{}, err
	}

	msg, err := message.Deserialize(&temp)
	if err != nil {
		return bitfield.Bitfield{}, err
	}

	if msg.MessageId != message.MsgBitfield {
		return bitfield.Bitfield{}, errors.New(fmt.Sprintf("Expected bitfield, received: %v", msg.MessageId))
	}

	return bitfield.Bitfield{Data: msg.Payload}, nil
}
