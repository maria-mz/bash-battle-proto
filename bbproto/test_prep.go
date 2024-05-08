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
		Version:     Version(genRandByte()),
		Action:      Action(genRandByte()),
		PayloadType: PayloadType(genRandByte()),
		StatusCode:  StatusCode(genRandByte()),
		Token:       NewToken(),
	}
}

func genRandCreateGameRequest() *CreateGameRequest {
	gameConfig := GameConfig{
		MaxPlayers:   genRandInt32(),
		Rounds:       genRandInt32(),
		RoundMinutes: genRandInt32(),
	}
	return &CreateGameRequest{
		GameConfig: gameConfig,
		Username:   genRandString(),
	}
}

func genRandCreateGameResponse() *CreateGameResponse {
	return &CreateGameResponse{
		GameId:   genRandString(),
		GameCode: genRandString(),
	}
}
