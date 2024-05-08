package bbproto

import (
	pb "github.com/maria-mz/bash-battle-proto/generated_protobuf"
	"google.golang.org/protobuf/proto"
)

type GameConfig struct {
	MaxPlayers   int32
	Rounds       int32
	RoundMinutes int32
}

type CreateGameRequest struct {
	GameConfig GameConfig
	Username   string
}

type CreateGameResponse struct {
	GameId   string
	GameCode string
}

func (req CreateGameRequest) Bytes() ([]byte, error) {
	gameConfigPb := &pb.GameConfig{
		MaxPlayers:   &req.GameConfig.MaxPlayers,
		Rounds:       &req.GameConfig.Rounds,
		RoundMinutes: &req.GameConfig.RoundMinutes,
	}

	reqPb := &pb.CreateGameRequest{
		GameConfig: gameConfigPb,
		Username:   &req.Username,
	}

	b, err := proto.Marshal(reqPb)

	return b, err
}

func (req *CreateGameRequest) PopulateFromBytes(b []byte) error {
	reqPb := &pb.CreateGameRequest{}

	if err := proto.Unmarshal(b, reqPb); err != nil {
		return err
	}

	gameConfig := GameConfig{
		MaxPlayers:   *reqPb.GameConfig.MaxPlayers,
		Rounds:       *reqPb.GameConfig.Rounds,
		RoundMinutes: *reqPb.GameConfig.RoundMinutes,
	}

	req.GameConfig = gameConfig
	req.Username = *reqPb.Username

	return nil
}

func (resp CreateGameResponse) Bytes() ([]byte, error) {
	reqPb := &pb.CreateGameResponse{
		GameId:   &resp.GameId,
		GameCode: &resp.GameCode,
	}

	b, err := proto.Marshal(reqPb)

	return b, err
}

func (resp *CreateGameResponse) PopulateFromBytes(b []byte) error {
	reqPb := &pb.CreateGameResponse{}

	if err := proto.Unmarshal(b, reqPb); err != nil {
		return err
	}

	resp.GameId = *reqPb.GameId
	resp.GameCode = *reqPb.GameCode

	return nil
}
