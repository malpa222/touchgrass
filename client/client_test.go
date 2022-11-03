package client

import (
	"testing"
)

func ConnectTest(t *testing.T) {
	c := &Client{
		Conn:     nil,
		Choked:   false,
		Bitfield: nil,
	}
}
