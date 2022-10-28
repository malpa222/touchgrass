package bitfield

type Bitfield []byte

func (b Bitfield) IsFlipped(index uint) bool {
	temp := b[index/8]  // get the byte of interest from the bitfield array
	offset := index % 8 // get the bit index

	return temp>>offset&1 == 1 // check if the last byte is flipped to 1
}

func (b Bitfield) FlipBit(index uint) {

}
