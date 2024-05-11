// Composable messages

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.20.3
// source: composable.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GameStatus int32

const (
	GameStatus_NotStarted GameStatus = 0
	GameStatus_InProgress GameStatus = 1
	GameStatus_Cancelled  GameStatus = 2
	GameStatus_Done       GameStatus = 3
)

// Enum value maps for GameStatus.
var (
	GameStatus_name = map[int32]string{
		0: "NotStarted",
		1: "InProgress",
		2: "Cancelled",
		3: "Done",
	}
	GameStatus_value = map[string]int32{
		"NotStarted": 0,
		"InProgress": 1,
		"Cancelled":  2,
		"Done":       3,
	}
)

func (x GameStatus) Enum() *GameStatus {
	p := new(GameStatus)
	*p = x
	return p
}

func (x GameStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (GameStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_composable_proto_enumTypes[0].Descriptor()
}

func (GameStatus) Type() protoreflect.EnumType {
	return &file_composable_proto_enumTypes[0]
}

func (x GameStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use GameStatus.Descriptor instead.
func (GameStatus) EnumDescriptor() ([]byte, []int) {
	return file_composable_proto_rawDescGZIP(), []int{0}
}

type GameLookupStatus int32

const (
	GameLookupStatus_GameFound    GameLookupStatus = 0
	GameLookupStatus_GameNotFound GameLookupStatus = 1
)

// Enum value maps for GameLookupStatus.
var (
	GameLookupStatus_name = map[int32]string{
		0: "GameFound",
		1: "GameNotFound",
	}
	GameLookupStatus_value = map[string]int32{
		"GameFound":    0,
		"GameNotFound": 1,
	}
)

func (x GameLookupStatus) Enum() *GameLookupStatus {
	p := new(GameLookupStatus)
	*p = x
	return p
}

func (x GameLookupStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (GameLookupStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_composable_proto_enumTypes[1].Descriptor()
}

func (GameLookupStatus) Type() protoreflect.EnumType {
	return &file_composable_proto_enumTypes[1]
}

func (x GameLookupStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use GameLookupStatus.Descriptor instead.
func (GameLookupStatus) EnumDescriptor() ([]byte, []int) {
	return file_composable_proto_rawDescGZIP(), []int{1}
}

type EmptyMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EmptyMessage) Reset() {
	*x = EmptyMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_composable_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmptyMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmptyMessage) ProtoMessage() {}

func (x *EmptyMessage) ProtoReflect() protoreflect.Message {
	mi := &file_composable_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmptyMessage.ProtoReflect.Descriptor instead.
func (*EmptyMessage) Descriptor() ([]byte, []int) {
	return file_composable_proto_rawDescGZIP(), []int{0}
}

type GameConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MaxPlayers   int32 `protobuf:"varint,1,opt,name=maxPlayers,proto3" json:"maxPlayers,omitempty"`
	Rounds       int32 `protobuf:"varint,2,opt,name=rounds,proto3" json:"rounds,omitempty"`
	RoundMinutes int32 `protobuf:"varint,3,opt,name=roundMinutes,proto3" json:"roundMinutes,omitempty"`
}

func (x *GameConfig) Reset() {
	*x = GameConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_composable_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GameConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GameConfig) ProtoMessage() {}

func (x *GameConfig) ProtoReflect() protoreflect.Message {
	mi := &file_composable_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GameConfig.ProtoReflect.Descriptor instead.
func (*GameConfig) Descriptor() ([]byte, []int) {
	return file_composable_proto_rawDescGZIP(), []int{1}
}

func (x *GameConfig) GetMaxPlayers() int32 {
	if x != nil {
		return x.MaxPlayers
	}
	return 0
}

func (x *GameConfig) GetRounds() int32 {
	if x != nil {
		return x.Rounds
	}
	return 0
}

func (x *GameConfig) GetRoundMinutes() int32 {
	if x != nil {
		return x.RoundMinutes
	}
	return 0
}

type Player struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name int32 `protobuf:"varint,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Player) Reset() {
	*x = Player{}
	if protoimpl.UnsafeEnabled {
		mi := &file_composable_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Player) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Player) ProtoMessage() {}

func (x *Player) ProtoReflect() protoreflect.Message {
	mi := &file_composable_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Player.ProtoReflect.Descriptor instead.
func (*Player) Descriptor() ([]byte, []int) {
	return file_composable_proto_rawDescGZIP(), []int{2}
}

func (x *Player) GetName() int32 {
	if x != nil {
		return x.Name
	}
	return 0
}

type PlayerRoundStats struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BeatRound bool    `protobuf:"varint,1,opt,name=beatRound,proto3" json:"beatRound,omitempty"`
	Command   *string `protobuf:"bytes,2,opt,name=command,proto3,oneof" json:"command,omitempty"`
}

func (x *PlayerRoundStats) Reset() {
	*x = PlayerRoundStats{}
	if protoimpl.UnsafeEnabled {
		mi := &file_composable_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlayerRoundStats) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlayerRoundStats) ProtoMessage() {}

func (x *PlayerRoundStats) ProtoReflect() protoreflect.Message {
	mi := &file_composable_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlayerRoundStats.ProtoReflect.Descriptor instead.
func (*PlayerRoundStats) Descriptor() ([]byte, []int) {
	return file_composable_proto_rawDescGZIP(), []int{3}
}

func (x *PlayerRoundStats) GetBeatRound() bool {
	if x != nil {
		return x.BeatRound
	}
	return false
}

func (x *PlayerRoundStats) GetCommand() string {
	if x != nil && x.Command != nil {
		return *x.Command
	}
	return ""
}

var File_composable_proto protoreflect.FileDescriptor

var file_composable_proto_rawDesc = []byte{
	0x0a, 0x10, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x0e, 0x0a, 0x0c, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x68, 0x0a, 0x0a, 0x47, 0x61, 0x6d,
	0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x1e, 0x0a, 0x0a, 0x6d, 0x61, 0x78, 0x50, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x6d, 0x61, 0x78,
	0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x6f, 0x75, 0x6e, 0x64,
	0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x73, 0x12,
	0x22, 0x0a, 0x0c, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x4d, 0x69, 0x6e, 0x75, 0x74, 0x65, 0x73, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x4d, 0x69, 0x6e, 0x75,
	0x74, 0x65, 0x73, 0x22, 0x1c, 0x0a, 0x06, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x22, 0x5b, 0x0a, 0x10, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x52, 0x6f, 0x75, 0x6e, 0x64,
	0x53, 0x74, 0x61, 0x74, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x62, 0x65, 0x61, 0x74, 0x52, 0x6f, 0x75,
	0x6e, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x62, 0x65, 0x61, 0x74, 0x52, 0x6f,
	0x75, 0x6e, 0x64, 0x12, 0x1d, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x88,
	0x01, 0x01, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x2a, 0x45,
	0x0a, 0x0a, 0x47, 0x61, 0x6d, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0e, 0x0a, 0x0a,
	0x4e, 0x6f, 0x74, 0x53, 0x74, 0x61, 0x72, 0x74, 0x65, 0x64, 0x10, 0x00, 0x12, 0x0e, 0x0a, 0x0a,
	0x49, 0x6e, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x10, 0x01, 0x12, 0x0d, 0x0a, 0x09,
	0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x6c, 0x65, 0x64, 0x10, 0x02, 0x12, 0x08, 0x0a, 0x04, 0x44,
	0x6f, 0x6e, 0x65, 0x10, 0x03, 0x2a, 0x33, 0x0a, 0x10, 0x47, 0x61, 0x6d, 0x65, 0x4c, 0x6f, 0x6f,
	0x6b, 0x75, 0x70, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0d, 0x0a, 0x09, 0x47, 0x61, 0x6d,
	0x65, 0x46, 0x6f, 0x75, 0x6e, 0x64, 0x10, 0x00, 0x12, 0x10, 0x0a, 0x0c, 0x47, 0x61, 0x6d, 0x65,
	0x4e, 0x6f, 0x74, 0x46, 0x6f, 0x75, 0x6e, 0x64, 0x10, 0x01, 0x42, 0x27, 0x5a, 0x25, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x61, 0x72, 0x69, 0x61, 0x2d, 0x6d,
	0x7a, 0x2f, 0x62, 0x61, 0x73, 0x68, 0x2d, 0x62, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_composable_proto_rawDescOnce sync.Once
	file_composable_proto_rawDescData = file_composable_proto_rawDesc
)

func file_composable_proto_rawDescGZIP() []byte {
	file_composable_proto_rawDescOnce.Do(func() {
		file_composable_proto_rawDescData = protoimpl.X.CompressGZIP(file_composable_proto_rawDescData)
	})
	return file_composable_proto_rawDescData
}

var file_composable_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_composable_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_composable_proto_goTypes = []interface{}{
	(GameStatus)(0),          // 0: proto.GameStatus
	(GameLookupStatus)(0),    // 1: proto.GameLookupStatus
	(*EmptyMessage)(nil),     // 2: proto.EmptyMessage
	(*GameConfig)(nil),       // 3: proto.GameConfig
	(*Player)(nil),           // 4: proto.Player
	(*PlayerRoundStats)(nil), // 5: proto.PlayerRoundStats
}
var file_composable_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_composable_proto_init() }
func file_composable_proto_init() {
	if File_composable_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_composable_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmptyMessage); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_composable_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GameConfig); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_composable_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Player); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_composable_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlayerRoundStats); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_composable_proto_msgTypes[3].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_composable_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_composable_proto_goTypes,
		DependencyIndexes: file_composable_proto_depIdxs,
		EnumInfos:         file_composable_proto_enumTypes,
		MessageInfos:      file_composable_proto_msgTypes,
	}.Build()
	File_composable_proto = out.File
	file_composable_proto_rawDesc = nil
	file_composable_proto_goTypes = nil
	file_composable_proto_depIdxs = nil
}