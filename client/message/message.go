package message

import (
	"bytes"
	"encoding/binary"
	"errors"
)

type ID byte

const (
	MsgChoke ID = iota
	MsgUnchoke
	MsgInterested
	MsgNotInterested
	MsgHave
	MsgBitfield
	MsgRequest
	MsgPiece
	MsgCancel
)

type Message struct {
	MessageId ID
	Payload   []byte
}

func (msg *Message) Serialize() *[]byte {
	length := uint32(len(msg.Payload) + 1) // length = payload + 1 byte for the MessageId

	buf := make([]byte, length+4) // 4 bytes for the length identifier
	binary.BigEndian.PutUint32(buf, length)

	buf[4] = byte(msg.MessageId) // put the message id into the buffer
	copy(buf[5:], msg.Payload)   // copy the payload into the buffer

	return &buf
}

func Deserialize(buf *[]byte) (msg *Message, err error) {
	reader := bytes.NewReader(*buf)

	// get the length
	temp := make([]byte, 4)
	if num, err := reader.Read(temp); err != nil {
		return nil, err
	} else if num == 0 {
		return msg, nil // keepalive
	} else if num != 4 {
		return nil, errors.New("invalid length")
	}

	length := binary.BigEndian.Uint32(temp)

	temp = make([]byte, length)
	if num, err := reader.Read(temp); err != nil {
		return nil, err
	} else if num != int(length) {
		return nil, errors.New("malformed payload")
	}

	return &Message{
		MessageId: ID(temp[0]),
		Payload:   temp[1:],
	}, nil
}
