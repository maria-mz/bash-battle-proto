package bbproto

import (
	pb "github.com/maria-mz/bash-battle-proto/generated_pb"
	"google.golang.org/protobuf/proto"
)

type GameConfig struct {
	MaxPlayers   int32
	Rounds       int32
	RoundMinutes int32
}

type JoinRegistry struct {
	Username string
}

type CreateGame struct {
	GameConfig GameConfig
}

type GameCreated struct {
	GameId   string
	GameCode string
}

func (msg JoinRegistry) Bytes() ([]byte, error) {
	pb := &pb.JoinRegistry{Username: &msg.Username}

	b, err := proto.Marshal(pb)

	return b, err
}

func (msg *JoinRegistry) PopulateFromBytes(b []byte) error {
	pb := &pb.JoinRegistry{}

	if err := proto.Unmarshal(b, pb); err != nil {
		return err
	}

	msg.Username = *pb.Username

	return nil
}

func (msg CreateGame) Bytes() ([]byte, error) {
	pb := &pb.CreateGame{
		GameConfig: &pb.GameConfig{
			MaxPlayers:   &msg.GameConfig.MaxPlayers,
			Rounds:       &msg.GameConfig.Rounds,
			RoundMinutes: &msg.GameConfig.RoundMinutes,
		},
	}

	b, err := proto.Marshal(pb)

	return b, err
}

func (msg *CreateGame) PopulateFromBytes(b []byte) error {
	pb := &pb.CreateGame{}

	if err := proto.Unmarshal(b, pb); err != nil {
		return err
	}

	config := GameConfig{
		MaxPlayers:   *pb.GameConfig.MaxPlayers,
		Rounds:       *pb.GameConfig.Rounds,
		RoundMinutes: *pb.GameConfig.RoundMinutes,
	}

	msg.GameConfig = config

	return nil
}

func (msg GameCreated) Bytes() ([]byte, error) {
	pb := &pb.GameCreated{
		GameId:   &msg.GameId,
		GameCode: &msg.GameCode,
	}

	b, err := proto.Marshal(pb)

	return b, err
}

func (msg *GameCreated) PopulateFromBytes(b []byte) error {
	pb := &pb.GameCreated{}

	if err := proto.Unmarshal(b, pb); err != nil {
		return err
	}

	msg.GameId = *pb.GameId
	msg.GameCode = *pb.GameCode

	return nil
}
