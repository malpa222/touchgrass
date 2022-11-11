package message

import (
	"bytes"
	"testing"
)

func TestSerialize(t *testing.T) {
	s1 := "dziooooooooooooooooo"
	msg := &Message{
		MessageId: 0,
		Payload:   []byte(s1),
	}

	ser := msg.Serialize()
	if len(*ser) != len(s1)+4+1 {
		t.Errorf("serialization went wrong, expected %v byte long array, got:\n%v",
			len(s1)+4+1, len(*ser))
	}
}

func TestDeserialize(t *testing.T) {
	s1 := "dziooooooooooooooooo"
	msg := &Message{
		MessageId: 0,
		Payload:   []byte(s1),
	}

	ser := msg.Serialize()
	if len(*ser) != len(s1)+4+1 {
		t.Errorf("serialization went wrong, expected %v byte long array, got:\n%v",
			len(s1)+4+1, len(*ser))
	}

	r := bytes.NewReader(*ser)
	if _, err := Read(r); err != nil {
		t.Errorf("got an error:\n%v", err)
	}
}
