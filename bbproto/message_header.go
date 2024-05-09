package bbproto

import (
	"fmt"

	"github.com/google/uuid"
)

type Version byte
type Action byte
type Error byte
type Token string

type MessageHeader struct {
	Version Version
	Action  Action
	Error   Error
	Token   Token
}

const HEADER_LEN int = 39

const VERSION = 1

const (
	Action_UNSPECIFIED   Action = 0
	Action_JOIN_REGISTRY Action = 1
	Action_CREATE_GAME   Action = 2
)

const (
	Error_NO_ERR     Error = 0
	Error_CLIENT_ERR Error = 1
	Error_SERVER_ERR Error = 2
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

func (error *Error) populateFromBytes(b []byte) {
	*error = Error(b[2])
}

func (token *Token) populateFromBytes(b []byte) error {
	*token = Token(b[3:39])
	return nil // TODO: Can check size
}

func (header MessageHeader) Bytes() ([]byte, error) {
	b := make([]byte, 0, HEADER_LEN)

	b = append(b, byte(header.Version))
	b = append(b, byte(header.Action))
	b = append(b, byte(header.Error))

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
	header.Error.populateFromBytes(b)
	header.Token.populateFromBytes(b)

	return nil
}

func NewMessageHeader(b []byte) (*MessageHeader, error) {
	header := &MessageHeader{}
	err := header.PopulateFromBytes(b)
	return header, err
}

func NewToken() Token {
	return Token(uuid.New().String())
}
