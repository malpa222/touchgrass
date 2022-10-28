package handshake

import (
	"errors"
	"io"
)

const pstr = "BitTorrent protocol"

const lenPstr = 0x13    // length of the pstr
const lenReserved = 0x8 // length of the reserved bytes

type Handshake struct {
	InfoHash [20]byte
	PeerID   [20]byte
}

func Serialize(hs *Handshake) *[]byte {
	buf := make([]byte, 68)
	buf[0] = lenPstr

	offset := 1
	offset += copy(buf[offset:], pstr)                      // protocol id
	offset += copy(buf[offset:], make([]byte, lenReserved)) // empty 8 bytes
	offset += copy(buf[offset:], hs.InfoHash[:])            // the infohash of the torrent
	offset += copy(buf[offset:], hs.PeerID[:])              // the peerid used for connecting to tracker

	return &buf
}

func Deserialize(reader io.Reader) (hs *Handshake, err error) {
	// check if the protocol length is correct
	buf := make([]byte, 1)
	if num, err := reader.Read(buf); err != nil {
		return nil, err
	} else if buf[0] != lenPstr || num == 0 {
		return nil, errors.New("invalid length")
	}

	// check if the pstr matches
	buf = make([]byte, lenPstr)
	if num, err := reader.Read(buf); err != nil {
		return nil, err
	} else if num != lenPstr || string(buf[:]) != pstr {
		return nil, errors.New("invalid protocol specification")
	}

	// check the flags
	buf = make([]byte, lenReserved)
	if num, err := reader.Read(buf); err != nil {
		return nil, err
	} else if num != lenReserved {
		return nil, errors.New("invalid flags")
	}

	// try the rest of the Handshake
	buf = make([]byte, 40)
	if num, err := reader.Read(buf); err != nil {
		return nil, err
	} else if num != len(buf) {
		return nil, errors.New("invalid peer data")
	}

	// copy the peerID and infohash to new buffers
	hs = &Handshake{}
	copy(hs.PeerID[:], buf[:20])
	copy(hs.InfoHash[:], buf[20:])

	return hs, nil
}
