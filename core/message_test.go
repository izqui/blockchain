package core

import (
	"github.com/izqui/helpers"
	"reflect"
	"testing"
)

func TestMessageMarshalling(t *testing.T) {

	mes := &Message{Identifier: MESSAGE_GET_NODES, Options: []byte{1, 2, 3, 4}, Data: []byte(helpers.RandomString(helpers.RandomInt(1024, 1024*4)))}
	bs, err := mes.MarshalBinary()

	if err != nil {
		t.Error(err)
	}

	newMes := new(Message)
	err = newMes.UnmarshalBinary(bs)

	if err != nil {
		t.Error(err)
	}

	if !reflect.DeepEqual(newMes, mes) {

		t.Error("Marshall unmarshall message error")
	}
}
