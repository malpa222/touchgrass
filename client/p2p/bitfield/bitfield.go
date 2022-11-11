package bitfield

// using binary notation here because it is easier to visualize

type Bitfield struct {
	Data []byte
}

func (b Bitfield) HasPiece(index uint) bool {
	temp := b.Data[index/8] // get the byte of interest from the bitfield array
	offset := index % 8     // get the bit index

	return temp>>offset&0b0001 == 0b0001 // check if the last byte is flipped to 1
}

func (b Bitfield) FlipBit(index uint) {
	byteIdx := index / 8 // get the byte of interest from the bitfield array
	offset := index % 8  // get the bit index

	/* shift one by `offset` bits like: 0b0001 << 3 = 0b1000 -> 8
	then XOR the shifted byte with the b[byteIdx] and save the result

	for example, index == 3, so byteIdx == 0, offset = 3 and b[byteIdx] == 0b0101
	0b0001 << 3 == 0b1000 == 8

	  0b0101
	^ 0b1000
	  ------
	  0b1101 == 13 */
	b.Data[byteIdx] ^= 0b0001 << offset
}
