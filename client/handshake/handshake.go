package handshake

import (
	"bytes"
	"errors"
)

const pstr = "BitTorrent protocol"

const lenPstr = 0x13    // length of the pstr
const lenReserved = 0x8 // length of the reserved bytes

type Handshake struct {
	InfoHash [20]byte
	PeerID   [20]byte
}

func (hs *Handshake) Serialize() []byte {
	buf := make([]byte, 68)
	buf[0] = lenPstr

	offset := 1
	offset += copy(buf[offset:], pstr)                      // protocol id
	offset += copy(buf[offset:], make([]byte, lenReserved)) // empty 8 bytes
	offset += copy(buf[offset:], hs.InfoHash[:])            // the infohash of the torrent
	offset += copy(buf[offset:], hs.PeerID[:])              // the peerid used for connecting to tracker

	return buf
}

func Deserialize(buf []byte) (*Handshake, error) {
	reader := bytes.NewReader(buf)

	// check if the protocol length is correct
	temp := make([]byte, 1)
	if num, err := reader.Read(temp); err != nil {
		return nil, err
	} else if temp[0] != lenPstr || num == 0 {
		return nil, errors.New("invalid length")
	}

	// check if the pstr matches
	temp = make([]byte, lenPstr)
	if num, err := reader.Read(temp); err != nil {
		return nil, err
	} else if num != lenPstr || string(temp[:]) != pstr {
		return nil, errors.New("invalid protocol specification")
	}

	// check the flags
	temp = make([]byte, lenReserved)
	if num, err := reader.Read(temp); err != nil {
		return nil, err
	} else if num != lenReserved {
		return nil, errors.New("invalid flags")
	}

	// try the rest of the Handshake
	temp = make([]byte, 40)
	if num, err := reader.Read(temp); err != nil {
		return nil, err
	} else if num != len(temp) {
		return nil, errors.New("invalid peer data")
	}

	hs := &Handshake{}
	copy(hs.InfoHash[:], temp[:20])
	copy(hs.PeerID[:], temp[20:])

	return hs, nil
}
