package bbproto

import (
	"fmt"

	"github.com/google/uuid"
)

type Version byte
type Action byte
type PayloadType byte
type StatusCode byte
type Token string

type MessageHeader struct {
	Version     Version
	Action      Action
	PayloadType PayloadType
	StatusCode  StatusCode
	Token       Token
}

const HEADER_LEN int = 40

const VERSION = 1

const (
	Action_UNSPECIFIED Action = 0
	Action_BCST        Action = 1 // Broadcast
	Action_REQ         Action = 2 // Request
	Action_RESP        Action = 8 // Response
	Action_AWK         Action = 9 // Acknowledgement
)

const (
	PayloadType_UNSPECIFIED      PayloadType = 0
	PayloadType_CREATE_GAME_REQ  PayloadType = 1
	PayloadType_CREATE_GAME_RESP PayloadType = 1
	// TODO: more as they come
)

const (
	Status_Code_UNSPECIFIED  StatusCode = 0
	Status_Code_OK           StatusCode = 0
	Status_Code_SERVER_ERROR StatusCode = 1
	// TODO: more as they come
)

func (token Token) bytes() ([]byte, error) {
	return []byte(token), nil
}

func (v *Version) populateFromBytes(b []byte) {
	*v = Version(b[0])
}

func (action *Action) populateFromBytes(b []byte) {
	*action = Action(b[1])
}

func (payloadType *PayloadType) populateFromBytes(b []byte) {
	*payloadType = PayloadType(b[2])
}

func (statusCode *StatusCode) populateFromBytes(b []byte) {
	*statusCode = StatusCode(b[3])
}

func (token *Token) populateFromBytes(b []byte) error {
	*token = Token(b[4:40])
	return nil // TODO: Can check size
}

func (header MessageHeader) Bytes() ([]byte, error) {
	b := make([]byte, 0, HEADER_LEN)

	b = append(b, byte(header.Version))
	b = append(b, byte(header.Action))
	b = append(b, byte(header.PayloadType))
	b = append(b, byte(header.StatusCode))

	tokenBytes, _ := header.Token.bytes()

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
	header.Token.populateFromBytes(b)

	return nil
}

func NewToken() Token {
	return Token(uuid.New().String())
}
