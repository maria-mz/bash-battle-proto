package bbproto

import (
	"testing"
)

func TestToken(t *testing.T) {
	token1 := NewToken()
	t.Logf("token1 = %s", token1)

	if length := len(token1); length != 36 {
		t.Errorf("token length is %d; expected %d", length, 36)
	}

	tokenBytes, _ := token1.Bytes()

	token2 := Token(tokenBytes)

	t.Logf("token2 = %s", token2)

	if token2 != token1 {
		t.Errorf("tokens don't match")
	}
}
