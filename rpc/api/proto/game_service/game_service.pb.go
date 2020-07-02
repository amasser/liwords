// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.24.0
// 	protoc        v3.10.1
// source: api/proto/game_service/game_service.proto

package game_service

import (
	realtime "github.com/domino14/liwords/rpc/api/proto/realtime"
	macondo "github.com/domino14/macondo/gen/api/proto/macondo"
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// Meta information about a game, including its players.
type GameInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GameId string `protobuf:"bytes,1,opt,name=game_id,json=gameId,proto3" json:"game_id,omitempty"`
}

func (x *GameInfoRequest) Reset() {
	*x = GameInfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_game_service_game_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GameInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GameInfoRequest) ProtoMessage() {}

func (x *GameInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_game_service_game_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GameInfoRequest.ProtoReflect.Descriptor instead.
func (*GameInfoRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_game_service_game_service_proto_rawDescGZIP(), []int{0}
}

func (x *GameInfoRequest) GetGameId() string {
	if x != nil {
		return x.GameId
	}
	return ""
}

// Meta information about the player of a particular game.
type PlayerInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId      string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Nickname    string `protobuf:"bytes,2,opt,name=nickname,proto3" json:"nickname,omitempty"`
	FullName    string `protobuf:"bytes,3,opt,name=full_name,json=fullName,proto3" json:"full_name,omitempty"`
	CountryCode string `protobuf:"bytes,4,opt,name=country_code,json=countryCode,proto3" json:"country_code,omitempty"`
	// Rating for the particular mode of the game. String because it could be
	// provisional or some other strings.
	Rating string `protobuf:"bytes,5,opt,name=rating,proto3" json:"rating,omitempty"`
	Title  string `protobuf:"bytes,6,opt,name=title,proto3" json:"title,omitempty"`
	// The avatar URL should be set if the player has an avatar, and blank
	// otherwise. It'll probably be a fixed pattern for the url.
	AvatarUrl string `protobuf:"bytes,7,opt,name=avatar_url,json=avatarUrl,proto3" json:"avatar_url,omitempty"`
}

func (x *PlayerInfo) Reset() {
	*x = PlayerInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_game_service_game_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlayerInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlayerInfo) ProtoMessage() {}

func (x *PlayerInfo) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_game_service_game_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlayerInfo.ProtoReflect.Descriptor instead.
func (*PlayerInfo) Descriptor() ([]byte, []int) {
	return file_api_proto_game_service_game_service_proto_rawDescGZIP(), []int{1}
}

func (x *PlayerInfo) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *PlayerInfo) GetNickname() string {
	if x != nil {
		return x.Nickname
	}
	return ""
}

func (x *PlayerInfo) GetFullName() string {
	if x != nil {
		return x.FullName
	}
	return ""
}

func (x *PlayerInfo) GetCountryCode() string {
	if x != nil {
		return x.CountryCode
	}
	return ""
}

func (x *PlayerInfo) GetRating() string {
	if x != nil {
		return x.Rating
	}
	return ""
}

func (x *PlayerInfo) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *PlayerInfo) GetAvatarUrl() string {
	if x != nil {
		return x.AvatarUrl
	}
	return ""
}

type GameInfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Players         []*PlayerInfo         `protobuf:"bytes,1,rep,name=players,proto3" json:"players,omitempty"`
	Lexicon         string                `protobuf:"bytes,2,opt,name=lexicon,proto3" json:"lexicon,omitempty"`
	Variant         string                `protobuf:"bytes,3,opt,name=variant,proto3" json:"variant,omitempty"`
	TimeControlName string                `protobuf:"bytes,4,opt,name=time_control_name,json=timeControlName,proto3" json:"time_control_name,omitempty"`
	TimeControl     string                `protobuf:"bytes,5,opt,name=time_control,json=timeControl,proto3" json:"time_control,omitempty"`
	TournamentName  string                `protobuf:"bytes,6,opt,name=tournament_name,json=tournamentName,proto3" json:"tournament_name,omitempty"`
	ChallengeRule   macondo.ChallengeRule `protobuf:"varint,7,opt,name=challenge_rule,json=challengeRule,proto3,enum=macondo.ChallengeRule" json:"challenge_rule,omitempty"`
	RatingMode      realtime.RatingMode   `protobuf:"varint,8,opt,name=rating_mode,json=ratingMode,proto3,enum=liwords.RatingMode" json:"rating_mode,omitempty"`
	// done - is game done?
	Done bool `protobuf:"varint,9,opt,name=done,proto3" json:"done,omitempty"`
}

func (x *GameInfoResponse) Reset() {
	*x = GameInfoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_game_service_game_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GameInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GameInfoResponse) ProtoMessage() {}

func (x *GameInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_game_service_game_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GameInfoResponse.ProtoReflect.Descriptor instead.
func (*GameInfoResponse) Descriptor() ([]byte, []int) {
	return file_api_proto_game_service_game_service_proto_rawDescGZIP(), []int{2}
}

func (x *GameInfoResponse) GetPlayers() []*PlayerInfo {
	if x != nil {
		return x.Players
	}
	return nil
}

func (x *GameInfoResponse) GetLexicon() string {
	if x != nil {
		return x.Lexicon
	}
	return ""
}

func (x *GameInfoResponse) GetVariant() string {
	if x != nil {
		return x.Variant
	}
	return ""
}

func (x *GameInfoResponse) GetTimeControlName() string {
	if x != nil {
		return x.TimeControlName
	}
	return ""
}

func (x *GameInfoResponse) GetTimeControl() string {
	if x != nil {
		return x.TimeControl
	}
	return ""
}

func (x *GameInfoResponse) GetTournamentName() string {
	if x != nil {
		return x.TournamentName
	}
	return ""
}

func (x *GameInfoResponse) GetChallengeRule() macondo.ChallengeRule {
	if x != nil {
		return x.ChallengeRule
	}
	return macondo.ChallengeRule_VOID
}

func (x *GameInfoResponse) GetRatingMode() realtime.RatingMode {
	if x != nil {
		return x.RatingMode
	}
	return realtime.RatingMode_RATED
}

func (x *GameInfoResponse) GetDone() bool {
	if x != nil {
		return x.Done
	}
	return false
}

var File_api_proto_game_service_game_service_proto protoreflect.FileDescriptor

var file_api_proto_game_service_game_service_proto_rawDesc = []byte{
	0x0a, 0x29, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x61, 0x6d, 0x65,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x67, 0x61, 0x6d, 0x65, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x67, 0x61, 0x6d,
	0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x1a, 0x21, 0x61, 0x70, 0x69, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x72, 0x65, 0x61, 0x6c, 0x74, 0x69, 0x6d, 0x65, 0x2f, 0x72, 0x65,
	0x61, 0x6c, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x27, 0x6d, 0x61,
	0x63, 0x6f, 0x6e, 0x64, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x6d, 0x61, 0x63, 0x6f, 0x6e, 0x64, 0x6f, 0x2f, 0x6d, 0x61, 0x63, 0x6f, 0x6e, 0x64, 0x6f, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2a, 0x0a, 0x0f, 0x47, 0x61, 0x6d, 0x65, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x67, 0x61, 0x6d, 0x65,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x67, 0x61, 0x6d, 0x65, 0x49,
	0x64, 0x22, 0xce, 0x01, 0x0a, 0x0a, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x69, 0x63,
	0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x69, 0x63,
	0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x75, 0x6c, 0x6c, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x75, 0x6c, 0x6c, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x5f, 0x63, 0x6f,
	0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72,
	0x79, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x14, 0x0a,
	0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69,
	0x74, 0x6c, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x5f, 0x75, 0x72,
	0x6c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x55,
	0x72, 0x6c, 0x22, 0xfb, 0x02, 0x0a, 0x10, 0x47, 0x61, 0x6d, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x32, 0x0a, 0x07, 0x70, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x07, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6c,
	0x65, 0x78, 0x69, 0x63, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6c, 0x65,
	0x78, 0x69, 0x63, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x74,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x74, 0x12,
	0x2a, 0x0a, 0x11, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x74, 0x69, 0x6d, 0x65,
	0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x74,
	0x69, 0x6d, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x74, 0x69, 0x6d, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x12, 0x27,
	0x0a, 0x0f, 0x74, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x74, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6d,
	0x65, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x3d, 0x0a, 0x0e, 0x63, 0x68, 0x61, 0x6c, 0x6c,
	0x65, 0x6e, 0x67, 0x65, 0x5f, 0x72, 0x75, 0x6c, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x16, 0x2e, 0x6d, 0x61, 0x63, 0x6f, 0x6e, 0x64, 0x6f, 0x2e, 0x43, 0x68, 0x61, 0x6c, 0x6c, 0x65,
	0x6e, 0x67, 0x65, 0x52, 0x75, 0x6c, 0x65, 0x52, 0x0d, 0x63, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e,
	0x67, 0x65, 0x52, 0x75, 0x6c, 0x65, 0x12, 0x34, 0x0a, 0x0b, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67,
	0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x13, 0x2e, 0x6c, 0x69,
	0x77, 0x6f, 0x72, 0x64, 0x73, 0x2e, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x4d, 0x6f, 0x64, 0x65,
	0x52, 0x0a, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x64, 0x6f, 0x6e, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x64, 0x6f, 0x6e, 0x65,
	0x32, 0x63, 0x0a, 0x13, 0x47, 0x61, 0x6d, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4c, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x1d, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x61, 0x6d, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x61, 0x6d, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x38, 0x5a, 0x36, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x6f, 0x6d, 0x69, 0x6e, 0x6f, 0x31, 0x34, 0x2f, 0x6c, 0x69, 0x77,
	0x6f, 0x72, 0x64, 0x73, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x67, 0x61, 0x6d, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_proto_game_service_game_service_proto_rawDescOnce sync.Once
	file_api_proto_game_service_game_service_proto_rawDescData = file_api_proto_game_service_game_service_proto_rawDesc
)

func file_api_proto_game_service_game_service_proto_rawDescGZIP() []byte {
	file_api_proto_game_service_game_service_proto_rawDescOnce.Do(func() {
		file_api_proto_game_service_game_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_proto_game_service_game_service_proto_rawDescData)
	})
	return file_api_proto_game_service_game_service_proto_rawDescData
}

var file_api_proto_game_service_game_service_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_api_proto_game_service_game_service_proto_goTypes = []interface{}{
	(*GameInfoRequest)(nil),    // 0: game_service.GameInfoRequest
	(*PlayerInfo)(nil),         // 1: game_service.PlayerInfo
	(*GameInfoResponse)(nil),   // 2: game_service.GameInfoResponse
	(macondo.ChallengeRule)(0), // 3: macondo.ChallengeRule
	(realtime.RatingMode)(0),   // 4: liwords.RatingMode
}
var file_api_proto_game_service_game_service_proto_depIdxs = []int32{
	1, // 0: game_service.GameInfoResponse.players:type_name -> game_service.PlayerInfo
	3, // 1: game_service.GameInfoResponse.challenge_rule:type_name -> macondo.ChallengeRule
	4, // 2: game_service.GameInfoResponse.rating_mode:type_name -> liwords.RatingMode
	0, // 3: game_service.GameMetadataService.GetMetadata:input_type -> game_service.GameInfoRequest
	2, // 4: game_service.GameMetadataService.GetMetadata:output_type -> game_service.GameInfoResponse
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_api_proto_game_service_game_service_proto_init() }
func file_api_proto_game_service_game_service_proto_init() {
	if File_api_proto_game_service_game_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_proto_game_service_game_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GameInfoRequest); i {
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
		file_api_proto_game_service_game_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlayerInfo); i {
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
		file_api_proto_game_service_game_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GameInfoResponse); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_proto_game_service_game_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_proto_game_service_game_service_proto_goTypes,
		DependencyIndexes: file_api_proto_game_service_game_service_proto_depIdxs,
		MessageInfos:      file_api_proto_game_service_game_service_proto_msgTypes,
	}.Build()
	File_api_proto_game_service_game_service_proto = out.File
	file_api_proto_game_service_game_service_proto_rawDesc = nil
	file_api_proto_game_service_game_service_proto_goTypes = nil
	file_api_proto_game_service_game_service_proto_depIdxs = nil
}
