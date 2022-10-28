package bitfield

import "testing"

func TestBitfield_IsFlipped(t *testing.T) {
	b := Bitfield{0b0001, 0b0010}

	if !b.IsFlipped(0) {
		t.Errorf("Expected to recognize the flipped bit. %b", b[0])
	}

	if !b.IsFlipped(9) {
		t.Errorf("Expected to recognize the flipped bit. %b", b[1])
	}
}
