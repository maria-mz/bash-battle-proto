package bbproto

import (
	"reflect"
	"testing"
)

type byteableTest struct {
	name      string
	byteable1 Byteable
	byteable2 Byteable
}

func (bt *byteableTest) run(t *testing.T) {
	t.Logf("byteable1 = %+v", bt.byteable1)

	b, err := bt.byteable1.Bytes()

	if err != nil {
		t.Fatalf("error getting bytes: %s", err)
	}

	t.Log("bytes:", b)

	_ = bt.byteable2.PopulateFromBytes(b)

	if err != nil {
		t.Fatalf("error populating from bytes: %s", err)
	}

	t.Logf("byteable2 = %+v", bt.byteable2)

	if reflect.DeepEqual(bt.byteable1, bt.byteable2) != true {
		t.Fatalf("byteables don't match!")
	}
}

var byteableTests = []byteableTest{
	{
		name:      "random message header",
		byteable1: genRandMessageHeader(),
		byteable2: &MessageHeader{},
	},
	{
		name:      "random game request",
		byteable1: genRandCreateGameRequest(),
		byteable2: &CreateGameRequest{},
	},
	{
		name:      "random game response",
		byteable1: genRandCreateGameResponse(),
		byteable2: &CreateGameResponse{},
	},
}

func TestByteables(t *testing.T) {
	for _, bt := range byteableTests {
		t.Run(bt.name, bt.run)
	}
}
