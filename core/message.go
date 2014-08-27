package core

import (
	"bytes"
	"errors"
	"github.com/izqui/helpers"
)

type Message struct {
	Identifier byte
	Options    []byte
	Data       []byte

	Reply chan Message
}

func NewMessage(id byte) *Message {

	return &Message{Identifier: id}
}

func (m *Message) MarshalBinary() ([]byte, error) {

	buf := new(bytes.Buffer)

	buf.WriteByte(m.Identifier)
	buf.Write(helpers.FitBytesInto(m.Options, MESSAGE_OPTIONS_SIZE))
	buf.Write(m.Data)

	return buf.Bytes(), nil

}

func (m *Message) UnmarshalBinary(d []byte) error {

	buf := bytes.NewBuffer(d)

	if len(d) < MESSAGE_OPTIONS_SIZE+MESSAGE_TYPE_SIZE {
		return errors.New("Insuficient message size")
	}
	m.Identifier = buf.Next(1)[0]
	m.Options = helpers.StripByte(buf.Next(MESSAGE_OPTIONS_SIZE), 0)
	m.Data = buf.Next(helpers.MaxInt)

	return nil
}
