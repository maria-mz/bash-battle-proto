package bbproto

import (
	"fmt"
)

const HEADER_LEN int = 40

type MessageHeader struct {
	Version     Version
	Action      Action
	PayloadType PayloadType
	StatusCode  StatusCode
	Token       Token
}

func (header MessageHeader) Bytes() ([]byte, error) {
	b := make([]byte, 0, HEADER_LEN)

	b = append(b, byte(header.Version))
	b = append(b, byte(header.Action))
	b = append(b, byte(header.PayloadType))
	b = append(b, byte(header.StatusCode))

	tokenBytes, _ := header.Token.Bytes()

	b = append(b, tokenBytes...)

	return b, nil
}

func (header *MessageHeader) PopulateFromBytes(b []byte) error {
	if len(b) != HEADER_LEN {
		return fmt.Errorf("bad header length %d, expected %d", len(b), HEADER_LEN)
	}

	header.Version.populateFromBytes(b)
	header.Action.populateFromBytes(b)
	header.PayloadType.populateFromBytes(b)
	header.StatusCode.populateFromBytes(b)
	header.Token.PopulateFromBytes(b)

	return nil
}
