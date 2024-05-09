package bbproto

import (
	"crypto/rand"

	"github.com/google/uuid"
)

func genRandByte() byte {
	b := make([]byte, 1)
	rand.Read(b)
	return b[0]
}

func genRandInt32() int32 {
	return int32(genRandByte())
}

func genRandString() string {
	return uuid.New().String()
}

func genRandMessageHeader() *MessageHeader {
	return &MessageHeader{
		Version: Version(genRandByte()),
		Action:  Action(genRandByte()),
		Error:   Error(genRandByte()),
		Token:   NewToken(),
	}
}

func genRandJoinRegistry() *JoinRegistry {
	return &JoinRegistry{
		Username: genRandString(),
	}
}

func genRandCreateGame() *CreateGame {
	gameConfig := GameConfig{
		MaxPlayers:   genRandInt32(),
		Rounds:       genRandInt32(),
		RoundMinutes: genRandInt32(),
	}
	return &CreateGame{
		GameConfig: gameConfig,
	}
}

func genRandGameCreated() *GameCreated {
	return &GameCreated{
		GameId:   genRandString(),
		GameCode: genRandString(),
	}
}
