package bbproto

import "github.com/google/uuid"

type Token string

func NewToken() Token {
	return Token(uuid.New().String())
}

func (token Token) Bytes() ([]byte, error) {
	return []byte(token), nil
}

func (token *Token) PopulateFromBytes(b []byte) error {
	*token = Token(b[4:40])
	return nil // TODO: Can check size
}
