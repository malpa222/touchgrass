package message

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
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

func (msg *Message) Serialize() []byte {
	length := uint32(len(msg.Payload) + 1) // length = payload + 1 byte for the MessageId

	buf := make([]byte, length+4) // 4 bytes for the length identifier
	binary.BigEndian.PutUint32(buf, length)

	buf[4] = byte(msg.MessageId) // put the message id into the buffer
	copy(buf[5:], msg.Payload)   // copy the payload into the buffer

	return buf
}

func parseRequest() {

}

func (msg *Message) ParseHave() (Index int, err error) {
	if msg.MessageId != MsgHave {
		return 0, fmt.Errorf("expected HAVE (ID %d), got ID %d", MsgHave, msg.MessageId)
	} else if len(msg.Payload) != 4 {
		return 0, fmt.Errorf("expected payload length 4, got length %d", len(msg.Payload))
	}

	return int(binary.BigEndian.Uint32(msg.Payload)), nil
}

func (msg *Message) ParsePiece(index int, buf []byte) (Index int, err error) {
	if msg.MessageId != MsgPiece {
		return 0, fmt.Errorf("expected PIECE (ID %d), got ID %d", MsgPiece, msg.MessageId)
	}

	if len(msg.Payload) < 8 {
		return 0, fmt.Errorf("payload too short. %d < 8", len(msg.Payload))
	}

	parsedIndex := int(binary.BigEndian.Uint32(msg.Payload[0:4]))
	if parsedIndex != index {
		return 0, fmt.Errorf("expected index %d, got %d", index, parsedIndex)
	}

	begin := int(binary.BigEndian.Uint32(msg.Payload[4:8]))
	if begin >= len(buf) {
		return 0, fmt.Errorf("begin offset too high. %d >= %d", begin, len(buf))
	}

	data := msg.Payload[8:]
	if begin+len(data) > len(buf) {
		return 0, fmt.Errorf("data too long [%d] for offset %d with length %d", len(data), begin, len(buf))
	}

	copy(buf[begin:], data)
	return len(data), nil
}

func FormatRequest(index int, begin int, length int) []byte {
	payload := make([]byte, 12)
	binary.BigEndian.PutUint32(payload[0:4], uint32(index))
	binary.BigEndian.PutUint32(payload[4:8], uint32(begin))
	binary.BigEndian.PutUint32(payload[8:12], uint32(length))

	return payload
}

func FormatHave(index int) []byte {
	payload := make([]byte, 4)
	binary.BigEndian.PutUint32(payload, uint32(index))

	return payload
}

func Deserialize(buf *[]byte) (*Message, error) {
	reader := bytes.NewReader(*buf)

	// get the length
	temp := make([]byte, 4)
	if num, err := reader.Read(temp); err != nil {
		return nil, err
	} else if num == 0 {
		return nil, nil // keepalive
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
