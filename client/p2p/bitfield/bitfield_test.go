package bitfield

import "testing"

// using binary notation here because it is easier to visualize

func TestBitfield_HasPiece(t *testing.T) {
	b := Bitfield{0b0001, 0b0010}

	if !b.HasPiece(0) {
		t.Errorf("Expected to recognize the flipped bit. %04b", b[0])
	}

	if !b.HasPiece(9) {
		t.Errorf("Expected to recognize the flipped bit. %04b", b[1])
	}

}
func TestBitfield_FlipBit(t *testing.T) {
	b := Bitfield{0b0101, 0b0010}

	b.FlipBit(3) // 0b0001 --> 0b0101 = 5
	if b[0] != 0b1101 {
		t.Errorf("Expected to recognize the flipped bit. %04b", b[0])
	}

	b.FlipBit(9) // 0b0010 --> 0b0000
	if b[1] != 0b0000 {
		t.Errorf("Expected to recognize the flipped bit. %04b", b[1])
	}
}
