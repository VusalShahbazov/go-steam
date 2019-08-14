// Code generated by protoc-gen-go. DO NOT EDIT.
// source: steammessages_player.steamclient.proto

package unified

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package unified is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package unified to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type CPlayer_GetGameBadgeLevels_Request struct {
	Appid                *uint32  `protobuf:"varint,1,opt,name=appid" json:"appid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CPlayer_GetGameBadgeLevels_Request) Reset()         { *m = CPlayer_GetGameBadgeLevels_Request{} }
func (m *CPlayer_GetGameBadgeLevels_Request) String() string { return proto.CompactTextString(m) }
func (*CPlayer_GetGameBadgeLevels_Request) ProtoMessage()    {}
func (*CPlayer_GetGameBadgeLevels_Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_2b474a6cd0239a66, []int{0}
}

func (m *CPlayer_GetGameBadgeLevels_Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CPlayer_GetGameBadgeLevels_Request.Unmarshal(m, b)
}
func (m *CPlayer_GetGameBadgeLevels_Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CPlayer_GetGameBadgeLevels_Request.Marshal(b, m, deterministic)
}
func (m *CPlayer_GetGameBadgeLevels_Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CPlayer_GetGameBadgeLevels_Request.Merge(m, src)
}
func (m *CPlayer_GetGameBadgeLevels_Request) XXX_Size() int {
	return xxx_messageInfo_CPlayer_GetGameBadgeLevels_Request.Size(m)
}
func (m *CPlayer_GetGameBadgeLevels_Request) XXX_DiscardUnknown() {
	xxx_messageInfo_CPlayer_GetGameBadgeLevels_Request.DiscardUnknown(m)
}

var xxx_messageInfo_CPlayer_GetGameBadgeLevels_Request proto.InternalMessageInfo

func (m *CPlayer_GetGameBadgeLevels_Request) GetAppid() uint32 {
	if m != nil && m.Appid != nil {
		return *m.Appid
	}
	return 0
}

type CPlayer_GetGameBadgeLevels_Response struct {
	PlayerLevel          *uint32                                      `protobuf:"varint,1,opt,name=player_level" json:"player_level,omitempty"`
	Badges               []*CPlayer_GetGameBadgeLevels_Response_Badge `protobuf:"bytes,2,rep,name=badges" json:"badges,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                     `json:"-"`
	XXX_unrecognized     []byte                                       `json:"-"`
	XXX_sizecache        int32                                        `json:"-"`
}

func (m *CPlayer_GetGameBadgeLevels_Response) Reset()         { *m = CPlayer_GetGameBadgeLevels_Response{} }
func (m *CPlayer_GetGameBadgeLevels_Response) String() string { return proto.CompactTextString(m) }
func (*CPlayer_GetGameBadgeLevels_Response) ProtoMessage()    {}
func (*CPlayer_GetGameBadgeLevels_Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_2b474a6cd0239a66, []int{1}
}

func (m *CPlayer_GetGameBadgeLevels_Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CPlayer_GetGameBadgeLevels_Response.Unmarshal(m, b)
}
func (m *CPlayer_GetGameBadgeLevels_Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CPlayer_GetGameBadgeLevels_Response.Marshal(b, m, deterministic)
}
func (m *CPlayer_GetGameBadgeLevels_Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CPlayer_GetGameBadgeLevels_Response.Merge(m, src)
}
func (m *CPlayer_GetGameBadgeLevels_Response) XXX_Size() int {
	return xxx_messageInfo_CPlayer_GetGameBadgeLevels_Response.Size(m)
}
func (m *CPlayer_GetGameBadgeLevels_Response) XXX_DiscardUnknown() {
	xxx_messageInfo_CPlayer_GetGameBadgeLevels_Response.DiscardUnknown(m)
}

var xxx_messageInfo_CPlayer_GetGameBadgeLevels_Response proto.InternalMessageInfo

func (m *CPlayer_GetGameBadgeLevels_Response) GetPlayerLevel() uint32 {
	if m != nil && m.PlayerLevel != nil {
		return *m.PlayerLevel
	}
	return 0
}

func (m *CPlayer_GetGameBadgeLevels_Response) GetBadges() []*CPlayer_GetGameBadgeLevels_Response_Badge {
	if m != nil {
		return m.Badges
	}
	return nil
}

type CPlayer_GetGameBadgeLevels_Response_Badge struct {
	Level                *int32   `protobuf:"varint,1,opt,name=level" json:"level,omitempty"`
	Series               *int32   `protobuf:"varint,2,opt,name=series" json:"series,omitempty"`
	BorderColor          *uint32  `protobuf:"varint,3,opt,name=border_color" json:"border_color,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CPlayer_GetGameBadgeLevels_Response_Badge) Reset() {
	*m = CPlayer_GetGameBadgeLevels_Response_Badge{}
}
func (m *CPlayer_GetGameBadgeLevels_Response_Badge) String() string { return proto.CompactTextString(m) }
func (*CPlayer_GetGameBadgeLevels_Response_Badge) ProtoMessage()    {}
func (*CPlayer_GetGameBadgeLevels_Response_Badge) Descriptor() ([]byte, []int) {
	return fileDescriptor_2b474a6cd0239a66, []int{1, 0}
}

func (m *CPlayer_GetGameBadgeLevels_Response_Badge) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CPlayer_GetGameBadgeLevels_Response_Badge.Unmarshal(m, b)
}
func (m *CPlayer_GetGameBadgeLevels_Response_Badge) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CPlayer_GetGameBadgeLevels_Response_Badge.Marshal(b, m, deterministic)
}
func (m *CPlayer_GetGameBadgeLevels_Response_Badge) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CPlayer_GetGameBadgeLevels_Response_Badge.Merge(m, src)
}
func (m *CPlayer_GetGameBadgeLevels_Response_Badge) XXX_Size() int {
	return xxx_messageInfo_CPlayer_GetGameBadgeLevels_Response_Badge.Size(m)
}
func (m *CPlayer_GetGameBadgeLevels_Response_Badge) XXX_DiscardUnknown() {
	xxx_messageInfo_CPlayer_GetGameBadgeLevels_Response_Badge.DiscardUnknown(m)
}

var xxx_messageInfo_CPlayer_GetGameBadgeLevels_Response_Badge proto.InternalMessageInfo

func (m *CPlayer_GetGameBadgeLevels_Response_Badge) GetLevel() int32 {
	if m != nil && m.Level != nil {
		return *m.Level
	}
	return 0
}

func (m *CPlayer_GetGameBadgeLevels_Response_Badge) GetSeries() int32 {
	if m != nil && m.Series != nil {
		return *m.Series
	}
	return 0
}

func (m *CPlayer_GetGameBadgeLevels_Response_Badge) GetBorderColor() uint32 {
	if m != nil && m.BorderColor != nil {
		return *m.BorderColor
	}
	return 0
}

type CPlayer_GetLastPlayedTimes_Request struct {
	MinLastPlayed        *uint32  `protobuf:"varint,1,opt,name=min_last_played" json:"min_last_played,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CPlayer_GetLastPlayedTimes_Request) Reset()         { *m = CPlayer_GetLastPlayedTimes_Request{} }
func (m *CPlayer_GetLastPlayedTimes_Request) String() string { return proto.CompactTextString(m) }
func (*CPlayer_GetLastPlayedTimes_Request) ProtoMessage()    {}
func (*CPlayer_GetLastPlayedTimes_Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_2b474a6cd0239a66, []int{2}
}

func (m *CPlayer_GetLastPlayedTimes_Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CPlayer_GetLastPlayedTimes_Request.Unmarshal(m, b)
}
func (m *CPlayer_GetLastPlayedTimes_Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CPlayer_GetLastPlayedTimes_Request.Marshal(b, m, deterministic)
}
func (m *CPlayer_GetLastPlayedTimes_Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CPlayer_GetLastPlayedTimes_Request.Merge(m, src)
}
func (m *CPlayer_GetLastPlayedTimes_Request) XXX_Size() int {
	return xxx_messageInfo_CPlayer_GetLastPlayedTimes_Request.Size(m)
}
func (m *CPlayer_GetLastPlayedTimes_Request) XXX_DiscardUnknown() {
	xxx_messageInfo_CPlayer_GetLastPlayedTimes_Request.DiscardUnknown(m)
}

var xxx_messageInfo_CPlayer_GetLastPlayedTimes_Request proto.InternalMessageInfo

func (m *CPlayer_GetLastPlayedTimes_Request) GetMinLastPlayed() uint32 {
	if m != nil && m.MinLastPlayed != nil {
		return *m.MinLastPlayed
	}
	return 0
}

type CPlayer_GetLastPlayedTimes_Response struct {
	Games                []*CPlayer_GetLastPlayedTimes_Response_Game `protobuf:"bytes,1,rep,name=games" json:"games,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                    `json:"-"`
	XXX_unrecognized     []byte                                      `json:"-"`
	XXX_sizecache        int32                                       `json:"-"`
}

func (m *CPlayer_GetLastPlayedTimes_Response) Reset()         { *m = CPlayer_GetLastPlayedTimes_Response{} }
func (m *CPlayer_GetLastPlayedTimes_Response) String() string { return proto.CompactTextString(m) }
func (*CPlayer_GetLastPlayedTimes_Response) ProtoMessage()    {}
func (*CPlayer_GetLastPlayedTimes_Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_2b474a6cd0239a66, []int{3}
}

func (m *CPlayer_GetLastPlayedTimes_Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CPlayer_GetLastPlayedTimes_Response.Unmarshal(m, b)
}
func (m *CPlayer_GetLastPlayedTimes_Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CPlayer_GetLastPlayedTimes_Response.Marshal(b, m, deterministic)
}
func (m *CPlayer_GetLastPlayedTimes_Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CPlayer_GetLastPlayedTimes_Response.Merge(m, src)
}
func (m *CPlayer_GetLastPlayedTimes_Response) XXX_Size() int {
	return xxx_messageInfo_CPlayer_GetLastPlayedTimes_Response.Size(m)
}
func (m *CPlayer_GetLastPlayedTimes_Response) XXX_DiscardUnknown() {
	xxx_messageInfo_CPlayer_GetLastPlayedTimes_Response.DiscardUnknown(m)
}

var xxx_messageInfo_CPlayer_GetLastPlayedTimes_Response proto.InternalMessageInfo

func (m *CPlayer_GetLastPlayedTimes_Response) GetGames() []*CPlayer_GetLastPlayedTimes_Response_Game {
	if m != nil {
		return m.Games
	}
	return nil
}

type CPlayer_GetLastPlayedTimes_Response_Game struct {
	Appid                *int32   `protobuf:"varint,1,opt,name=appid" json:"appid,omitempty"`
	LastPlaytime         *uint32  `protobuf:"varint,2,opt,name=last_playtime" json:"last_playtime,omitempty"`
	Playtime_2Weeks      *int32   `protobuf:"varint,3,opt,name=playtime_2weeks" json:"playtime_2weeks,omitempty"`
	PlaytimeForever      *int32   `protobuf:"varint,4,opt,name=playtime_forever" json:"playtime_forever,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CPlayer_GetLastPlayedTimes_Response_Game) Reset() {
	*m = CPlayer_GetLastPlayedTimes_Response_Game{}
}
func (m *CPlayer_GetLastPlayedTimes_Response_Game) String() string { return proto.CompactTextString(m) }
func (*CPlayer_GetLastPlayedTimes_Response_Game) ProtoMessage()    {}
func (*CPlayer_GetLastPlayedTimes_Response_Game) Descriptor() ([]byte, []int) {
	return fileDescriptor_2b474a6cd0239a66, []int{3, 0}
}

func (m *CPlayer_GetLastPlayedTimes_Response_Game) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CPlayer_GetLastPlayedTimes_Response_Game.Unmarshal(m, b)
}
func (m *CPlayer_GetLastPlayedTimes_Response_Game) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CPlayer_GetLastPlayedTimes_Response_Game.Marshal(b, m, deterministic)
}
func (m *CPlayer_GetLastPlayedTimes_Response_Game) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CPlayer_GetLastPlayedTimes_Response_Game.Merge(m, src)
}
func (m *CPlayer_GetLastPlayedTimes_Response_Game) XXX_Size() int {
	return xxx_messageInfo_CPlayer_GetLastPlayedTimes_Response_Game.Size(m)
}
func (m *CPlayer_GetLastPlayedTimes_Response_Game) XXX_DiscardUnknown() {
	xxx_messageInfo_CPlayer_GetLastPlayedTimes_Response_Game.DiscardUnknown(m)
}

var xxx_messageInfo_CPlayer_GetLastPlayedTimes_Response_Game proto.InternalMessageInfo

func (m *CPlayer_GetLastPlayedTimes_Response_Game) GetAppid() int32 {
	if m != nil && m.Appid != nil {
		return *m.Appid
	}
	return 0
}

func (m *CPlayer_GetLastPlayedTimes_Response_Game) GetLastPlaytime() uint32 {
	if m != nil && m.LastPlaytime != nil {
		return *m.LastPlaytime
	}
	return 0
}

func (m *CPlayer_GetLastPlayedTimes_Response_Game) GetPlaytime_2Weeks() int32 {
	if m != nil && m.Playtime_2Weeks != nil {
		return *m.Playtime_2Weeks
	}
	return 0
}

func (m *CPlayer_GetLastPlayedTimes_Response_Game) GetPlaytimeForever() int32 {
	if m != nil && m.PlaytimeForever != nil {
		return *m.PlaytimeForever
	}
	return 0
}

type CPlayer_AcceptSSA_Request struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CPlayer_AcceptSSA_Request) Reset()         { *m = CPlayer_AcceptSSA_Request{} }
func (m *CPlayer_AcceptSSA_Request) String() string { return proto.CompactTextString(m) }
func (*CPlayer_AcceptSSA_Request) ProtoMessage()    {}
func (*CPlayer_AcceptSSA_Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_2b474a6cd0239a66, []int{4}
}

func (m *CPlayer_AcceptSSA_Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CPlayer_AcceptSSA_Request.Unmarshal(m, b)
}
func (m *CPlayer_AcceptSSA_Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CPlayer_AcceptSSA_Request.Marshal(b, m, deterministic)
}
func (m *CPlayer_AcceptSSA_Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CPlayer_AcceptSSA_Request.Merge(m, src)
}
func (m *CPlayer_AcceptSSA_Request) XXX_Size() int {
	return xxx_messageInfo_CPlayer_AcceptSSA_Request.Size(m)
}
func (m *CPlayer_AcceptSSA_Request) XXX_DiscardUnknown() {
	xxx_messageInfo_CPlayer_AcceptSSA_Request.DiscardUnknown(m)
}

var xxx_messageInfo_CPlayer_AcceptSSA_Request proto.InternalMessageInfo

type CPlayer_AcceptSSA_Response struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CPlayer_AcceptSSA_Response) Reset()         { *m = CPlayer_AcceptSSA_Response{} }
func (m *CPlayer_AcceptSSA_Response) String() string { return proto.CompactTextString(m) }
func (*CPlayer_AcceptSSA_Response) ProtoMessage()    {}
func (*CPlayer_AcceptSSA_Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_2b474a6cd0239a66, []int{5}
}

func (m *CPlayer_AcceptSSA_Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CPlayer_AcceptSSA_Response.Unmarshal(m, b)
}
func (m *CPlayer_AcceptSSA_Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CPlayer_AcceptSSA_Response.Marshal(b, m, deterministic)
}
func (m *CPlayer_AcceptSSA_Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CPlayer_AcceptSSA_Response.Merge(m, src)
}
func (m *CPlayer_AcceptSSA_Response) XXX_Size() int {
	return xxx_messageInfo_CPlayer_AcceptSSA_Response.Size(m)
}
func (m *CPlayer_AcceptSSA_Response) XXX_DiscardUnknown() {
	xxx_messageInfo_CPlayer_AcceptSSA_Response.DiscardUnknown(m)
}

var xxx_messageInfo_CPlayer_AcceptSSA_Response proto.InternalMessageInfo

type CPlayer_LastPlayedTimes_Notification struct {
	Games                []*CPlayer_GetLastPlayedTimes_Response_Game `protobuf:"bytes,1,rep,name=games" json:"games,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                    `json:"-"`
	XXX_unrecognized     []byte                                      `json:"-"`
	XXX_sizecache        int32                                       `json:"-"`
}

func (m *CPlayer_LastPlayedTimes_Notification) Reset()         { *m = CPlayer_LastPlayedTimes_Notification{} }
func (m *CPlayer_LastPlayedTimes_Notification) String() string { return proto.CompactTextString(m) }
func (*CPlayer_LastPlayedTimes_Notification) ProtoMessage()    {}
func (*CPlayer_LastPlayedTimes_Notification) Descriptor() ([]byte, []int) {
	return fileDescriptor_2b474a6cd0239a66, []int{6}
}

func (m *CPlayer_LastPlayedTimes_Notification) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CPlayer_LastPlayedTimes_Notification.Unmarshal(m, b)
}
func (m *CPlayer_LastPlayedTimes_Notification) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CPlayer_LastPlayedTimes_Notification.Marshal(b, m, deterministic)
}
func (m *CPlayer_LastPlayedTimes_Notification) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CPlayer_LastPlayedTimes_Notification.Merge(m, src)
}
func (m *CPlayer_LastPlayedTimes_Notification) XXX_Size() int {
	return xxx_messageInfo_CPlayer_LastPlayedTimes_Notification.Size(m)
}
func (m *CPlayer_LastPlayedTimes_Notification) XXX_DiscardUnknown() {
	xxx_messageInfo_CPlayer_LastPlayedTimes_Notification.DiscardUnknown(m)
}

var xxx_messageInfo_CPlayer_LastPlayedTimes_Notification proto.InternalMessageInfo

func (m *CPlayer_LastPlayedTimes_Notification) GetGames() []*CPlayer_GetLastPlayedTimes_Response_Game {
	if m != nil {
		return m.Games
	}
	return nil
}

type CPlayerClient_GetSystemInformation_Request struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CPlayerClient_GetSystemInformation_Request) Reset() {
	*m = CPlayerClient_GetSystemInformation_Request{}
}
func (m *CPlayerClient_GetSystemInformation_Request) String() string {
	return proto.CompactTextString(m)
}
func (*CPlayerClient_GetSystemInformation_Request) ProtoMessage() {}
func (*CPlayerClient_GetSystemInformation_Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_2b474a6cd0239a66, []int{7}
}

func (m *CPlayerClient_GetSystemInformation_Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CPlayerClient_GetSystemInformation_Request.Unmarshal(m, b)
}
func (m *CPlayerClient_GetSystemInformation_Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CPlayerClient_GetSystemInformation_Request.Marshal(b, m, deterministic)
}
func (m *CPlayerClient_GetSystemInformation_Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CPlayerClient_GetSystemInformation_Request.Merge(m, src)
}
func (m *CPlayerClient_GetSystemInformation_Request) XXX_Size() int {
	return xxx_messageInfo_CPlayerClient_GetSystemInformation_Request.Size(m)
}
func (m *CPlayerClient_GetSystemInformation_Request) XXX_DiscardUnknown() {
	xxx_messageInfo_CPlayerClient_GetSystemInformation_Request.DiscardUnknown(m)
}

var xxx_messageInfo_CPlayerClient_GetSystemInformation_Request proto.InternalMessageInfo

type CClientSystemInfo struct {
	Cpu                  *CClientSystemInfo_CPU       `protobuf:"bytes,1,opt,name=cpu" json:"cpu,omitempty"`
	VideoCard            *CClientSystemInfo_VideoCard `protobuf:"bytes,2,opt,name=video_card" json:"video_card,omitempty"`
	OperatingSystem      *string                      `protobuf:"bytes,3,opt,name=operating_system" json:"operating_system,omitempty"`
	Os_64Bit             *bool                        `protobuf:"varint,4,opt,name=os_64bit" json:"os_64bit,omitempty"`
	SystemRamMb          *int32                       `protobuf:"varint,5,opt,name=system_ram_mb" json:"system_ram_mb,omitempty"`
	AudioDevice          *string                      `protobuf:"bytes,6,opt,name=audio_device" json:"audio_device,omitempty"`
	AudioDriverVersion   *string                      `protobuf:"bytes,7,opt,name=audio_driver_version" json:"audio_driver_version,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *CClientSystemInfo) Reset()         { *m = CClientSystemInfo{} }
func (m *CClientSystemInfo) String() string { return proto.CompactTextString(m) }
func (*CClientSystemInfo) ProtoMessage()    {}
func (*CClientSystemInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_2b474a6cd0239a66, []int{8}
}

func (m *CClientSystemInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CClientSystemInfo.Unmarshal(m, b)
}
func (m *CClientSystemInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CClientSystemInfo.Marshal(b, m, deterministic)
}
func (m *CClientSystemInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CClientSystemInfo.Merge(m, src)
}
func (m *CClientSystemInfo) XXX_Size() int {
	return xxx_messageInfo_CClientSystemInfo.Size(m)
}
func (m *CClientSystemInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_CClientSystemInfo.DiscardUnknown(m)
}

var xxx_messageInfo_CClientSystemInfo proto.InternalMessageInfo

func (m *CClientSystemInfo) GetCpu() *CClientSystemInfo_CPU {
	if m != nil {
		return m.Cpu
	}
	return nil
}

func (m *CClientSystemInfo) GetVideoCard() *CClientSystemInfo_VideoCard {
	if m != nil {
		return m.VideoCard
	}
	return nil
}

func (m *CClientSystemInfo) GetOperatingSystem() string {
	if m != nil && m.OperatingSystem != nil {
		return *m.OperatingSystem
	}
	return ""
}

func (m *CClientSystemInfo) GetOs_64Bit() bool {
	if m != nil && m.Os_64Bit != nil {
		return *m.Os_64Bit
	}
	return false
}

func (m *CClientSystemInfo) GetSystemRamMb() int32 {
	if m != nil && m.SystemRamMb != nil {
		return *m.SystemRamMb
	}
	return 0
}

func (m *CClientSystemInfo) GetAudioDevice() string {
	if m != nil && m.AudioDevice != nil {
		return *m.AudioDevice
	}
	return ""
}

func (m *CClientSystemInfo) GetAudioDriverVersion() string {
	if m != nil && m.AudioDriverVersion != nil {
		return *m.AudioDriverVersion
	}
	return ""
}

type CClientSystemInfo_CPU struct {
	SpeedMhz             *int32   `protobuf:"varint,1,opt,name=speed_mhz" json:"speed_mhz,omitempty"`
	Vendor               *string  `protobuf:"bytes,2,opt,name=vendor" json:"vendor,omitempty"`
	LogicalProcessors    *int32   `protobuf:"varint,3,opt,name=logical_processors" json:"logical_processors,omitempty"`
	PhysicalProcessors   *int32   `protobuf:"varint,4,opt,name=physical_processors" json:"physical_processors,omitempty"`
	Hyperthreading       *bool    `protobuf:"varint,5,opt,name=hyperthreading" json:"hyperthreading,omitempty"`
	Fcmov                *bool    `protobuf:"varint,6,opt,name=fcmov" json:"fcmov,omitempty"`
	Sse2                 *bool    `protobuf:"varint,7,opt,name=sse2" json:"sse2,omitempty"`
	Sse3                 *bool    `protobuf:"varint,8,opt,name=sse3" json:"sse3,omitempty"`
	Ssse3                *bool    `protobuf:"varint,9,opt,name=ssse3" json:"ssse3,omitempty"`
	Sse4A                *bool    `protobuf:"varint,10,opt,name=sse4a" json:"sse4a,omitempty"`
	Sse41                *bool    `protobuf:"varint,11,opt,name=sse41" json:"sse41,omitempty"`
	Sse42                *bool    `protobuf:"varint,12,opt,name=sse42" json:"sse42,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CClientSystemInfo_CPU) Reset()         { *m = CClientSystemInfo_CPU{} }
func (m *CClientSystemInfo_CPU) String() string { return proto.CompactTextString(m) }
func (*CClientSystemInfo_CPU) ProtoMessage()    {}
func (*CClientSystemInfo_CPU) Descriptor() ([]byte, []int) {
	return fileDescriptor_2b474a6cd0239a66, []int{8, 0}
}

func (m *CClientSystemInfo_CPU) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CClientSystemInfo_CPU.Unmarshal(m, b)
}
func (m *CClientSystemInfo_CPU) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CClientSystemInfo_CPU.Marshal(b, m, deterministic)
}
func (m *CClientSystemInfo_CPU) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CClientSystemInfo_CPU.Merge(m, src)
}
func (m *CClientSystemInfo_CPU) XXX_Size() int {
	return xxx_messageInfo_CClientSystemInfo_CPU.Size(m)
}
func (m *CClientSystemInfo_CPU) XXX_DiscardUnknown() {
	xxx_messageInfo_CClientSystemInfo_CPU.DiscardUnknown(m)
}

var xxx_messageInfo_CClientSystemInfo_CPU proto.InternalMessageInfo

func (m *CClientSystemInfo_CPU) GetSpeedMhz() int32 {
	if m != nil && m.SpeedMhz != nil {
		return *m.SpeedMhz
	}
	return 0
}

func (m *CClientSystemInfo_CPU) GetVendor() string {
	if m != nil && m.Vendor != nil {
		return *m.Vendor
	}
	return ""
}

func (m *CClientSystemInfo_CPU) GetLogicalProcessors() int32 {
	if m != nil && m.LogicalProcessors != nil {
		return *m.LogicalProcessors
	}
	return 0
}

func (m *CClientSystemInfo_CPU) GetPhysicalProcessors() int32 {
	if m != nil && m.PhysicalProcessors != nil {
		return *m.PhysicalProcessors
	}
	return 0
}

func (m *CClientSystemInfo_CPU) GetHyperthreading() bool {
	if m != nil && m.Hyperthreading != nil {
		return *m.Hyperthreading
	}
	return false
}

func (m *CClientSystemInfo_CPU) GetFcmov() bool {
	if m != nil && m.Fcmov != nil {
		return *m.Fcmov
	}
	return false
}

func (m *CClientSystemInfo_CPU) GetSse2() bool {
	if m != nil && m.Sse2 != nil {
		return *m.Sse2
	}
	return false
}

func (m *CClientSystemInfo_CPU) GetSse3() bool {
	if m != nil && m.Sse3 != nil {
		return *m.Sse3
	}
	return false
}

func (m *CClientSystemInfo_CPU) GetSsse3() bool {
	if m != nil && m.Ssse3 != nil {
		return *m.Ssse3
	}
	return false
}

func (m *CClientSystemInfo_CPU) GetSse4A() bool {
	if m != nil && m.Sse4A != nil {
		return *m.Sse4A
	}
	return false
}

func (m *CClientSystemInfo_CPU) GetSse41() bool {
	if m != nil && m.Sse41 != nil {
		return *m.Sse41
	}
	return false
}

func (m *CClientSystemInfo_CPU) GetSse42() bool {
	if m != nil && m.Sse42 != nil {
		return *m.Sse42
	}
	return false
}

type CClientSystemInfo_VideoCard struct {
	Driver               *string  `protobuf:"bytes,1,opt,name=driver" json:"driver,omitempty"`
	DriverVersion        *string  `protobuf:"bytes,2,opt,name=driver_version" json:"driver_version,omitempty"`
	DriverDate           *uint32  `protobuf:"varint,3,opt,name=driver_date" json:"driver_date,omitempty"`
	DirectxVersion       *string  `protobuf:"bytes,4,opt,name=directx_version" json:"directx_version,omitempty"`
	OpenglVersion        *string  `protobuf:"bytes,5,opt,name=opengl_version" json:"opengl_version,omitempty"`
	Vendorid             *int32   `protobuf:"varint,6,opt,name=vendorid" json:"vendorid,omitempty"`
	Deviceid             *int32   `protobuf:"varint,7,opt,name=deviceid" json:"deviceid,omitempty"`
	VramMb               *int32   `protobuf:"varint,8,opt,name=vram_mb" json:"vram_mb,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CClientSystemInfo_VideoCard) Reset()         { *m = CClientSystemInfo_VideoCard{} }
func (m *CClientSystemInfo_VideoCard) String() string { return proto.CompactTextString(m) }
func (*CClientSystemInfo_VideoCard) ProtoMessage()    {}
func (*CClientSystemInfo_VideoCard) Descriptor() ([]byte, []int) {
	return fileDescriptor_2b474a6cd0239a66, []int{8, 1}
}

func (m *CClientSystemInfo_VideoCard) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CClientSystemInfo_VideoCard.Unmarshal(m, b)
}
func (m *CClientSystemInfo_VideoCard) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CClientSystemInfo_VideoCard.Marshal(b, m, deterministic)
}
func (m *CClientSystemInfo_VideoCard) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CClientSystemInfo_VideoCard.Merge(m, src)
}
func (m *CClientSystemInfo_VideoCard) XXX_Size() int {
	return xxx_messageInfo_CClientSystemInfo_VideoCard.Size(m)
}
func (m *CClientSystemInfo_VideoCard) XXX_DiscardUnknown() {
	xxx_messageInfo_CClientSystemInfo_VideoCard.DiscardUnknown(m)
}

var xxx_messageInfo_CClientSystemInfo_VideoCard proto.InternalMessageInfo

func (m *CClientSystemInfo_VideoCard) GetDriver() string {
	if m != nil && m.Driver != nil {
		return *m.Driver
	}
	return ""
}

func (m *CClientSystemInfo_VideoCard) GetDriverVersion() string {
	if m != nil && m.DriverVersion != nil {
		return *m.DriverVersion
	}
	return ""
}

func (m *CClientSystemInfo_VideoCard) GetDriverDate() uint32 {
	if m != nil && m.DriverDate != nil {
		return *m.DriverDate
	}
	return 0
}

func (m *CClientSystemInfo_VideoCard) GetDirectxVersion() string {
	if m != nil && m.DirectxVersion != nil {
		return *m.DirectxVersion
	}
	return ""
}

func (m *CClientSystemInfo_VideoCard) GetOpenglVersion() string {
	if m != nil && m.OpenglVersion != nil {
		return *m.OpenglVersion
	}
	return ""
}

func (m *CClientSystemInfo_VideoCard) GetVendorid() int32 {
	if m != nil && m.Vendorid != nil {
		return *m.Vendorid
	}
	return 0
}

func (m *CClientSystemInfo_VideoCard) GetDeviceid() int32 {
	if m != nil && m.Deviceid != nil {
		return *m.Deviceid
	}
	return 0
}

func (m *CClientSystemInfo_VideoCard) GetVramMb() int32 {
	if m != nil && m.VramMb != nil {
		return *m.VramMb
	}
	return 0
}

type CPlayerClient_GetSystemInformation_Response struct {
	SystemInfo           *CClientSystemInfo `protobuf:"bytes,1,opt,name=system_info" json:"system_info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *CPlayerClient_GetSystemInformation_Response) Reset() {
	*m = CPlayerClient_GetSystemInformation_Response{}
}
func (m *CPlayerClient_GetSystemInformation_Response) String() string {
	return proto.CompactTextString(m)
}
func (*CPlayerClient_GetSystemInformation_Response) ProtoMessage() {}
func (*CPlayerClient_GetSystemInformation_Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_2b474a6cd0239a66, []int{9}
}

func (m *CPlayerClient_GetSystemInformation_Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CPlayerClient_GetSystemInformation_Response.Unmarshal(m, b)
}
func (m *CPlayerClient_GetSystemInformation_Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CPlayerClient_GetSystemInformation_Response.Marshal(b, m, deterministic)
}
func (m *CPlayerClient_GetSystemInformation_Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CPlayerClient_GetSystemInformation_Response.Merge(m, src)
}
func (m *CPlayerClient_GetSystemInformation_Response) XXX_Size() int {
	return xxx_messageInfo_CPlayerClient_GetSystemInformation_Response.Size(m)
}
func (m *CPlayerClient_GetSystemInformation_Response) XXX_DiscardUnknown() {
	xxx_messageInfo_CPlayerClient_GetSystemInformation_Response.DiscardUnknown(m)
}

var xxx_messageInfo_CPlayerClient_GetSystemInformation_Response proto.InternalMessageInfo

func (m *CPlayerClient_GetSystemInformation_Response) GetSystemInfo() *CClientSystemInfo {
	if m != nil {
		return m.SystemInfo
	}
	return nil
}

func init() {
	proto.RegisterType((*CPlayer_GetGameBadgeLevels_Request)(nil), "CPlayer_GetGameBadgeLevels_Request")
	proto.RegisterType((*CPlayer_GetGameBadgeLevels_Response)(nil), "CPlayer_GetGameBadgeLevels_Response")
	proto.RegisterType((*CPlayer_GetGameBadgeLevels_Response_Badge)(nil), "CPlayer_GetGameBadgeLevels_Response.Badge")
	proto.RegisterType((*CPlayer_GetLastPlayedTimes_Request)(nil), "CPlayer_GetLastPlayedTimes_Request")
	proto.RegisterType((*CPlayer_GetLastPlayedTimes_Response)(nil), "CPlayer_GetLastPlayedTimes_Response")
	proto.RegisterType((*CPlayer_GetLastPlayedTimes_Response_Game)(nil), "CPlayer_GetLastPlayedTimes_Response.Game")
	proto.RegisterType((*CPlayer_AcceptSSA_Request)(nil), "CPlayer_AcceptSSA_Request")
	proto.RegisterType((*CPlayer_AcceptSSA_Response)(nil), "CPlayer_AcceptSSA_Response")
	proto.RegisterType((*CPlayer_LastPlayedTimes_Notification)(nil), "CPlayer_LastPlayedTimes_Notification")
	proto.RegisterType((*CPlayerClient_GetSystemInformation_Request)(nil), "CPlayerClient_GetSystemInformation_Request")
	proto.RegisterType((*CClientSystemInfo)(nil), "CClientSystemInfo")
	proto.RegisterType((*CClientSystemInfo_CPU)(nil), "CClientSystemInfo.CPU")
	proto.RegisterType((*CClientSystemInfo_VideoCard)(nil), "CClientSystemInfo.VideoCard")
	proto.RegisterType((*CPlayerClient_GetSystemInformation_Response)(nil), "CPlayerClient_GetSystemInformation_Response")
}

func init() {
	proto.RegisterFile("steammessages_player.steamclient.proto", fileDescriptor_2b474a6cd0239a66)
}

var fileDescriptor_2b474a6cd0239a66 = []byte{
	// 1073 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x55, 0x4f, 0x6f, 0xdb, 0xc6,
	0x13, 0x05, 0x23, 0xcb, 0x91, 0x46, 0x76, 0x9c, 0x6c, 0x9c, 0xfc, 0x18, 0xda, 0x3f, 0x60, 0x21,
	0xa7, 0xf9, 0xe3, 0x38, 0x44, 0xab, 0x04, 0x45, 0xd1, 0x16, 0x08, 0x6c, 0x1d, 0x8c, 0x02, 0x46,
	0xe0, 0xc6, 0x71, 0x4e, 0x05, 0xd8, 0x15, 0x39, 0x94, 0x16, 0x21, 0xb9, 0xec, 0xee, 0x4a, 0xa9,
	0x7a, 0x2a, 0x74, 0xee, 0xad, 0xe8, 0xd7, 0xe8, 0x4d, 0x3d, 0x07, 0xe8, 0x17, 0xe9, 0x17, 0xe8,
	0xb1, 0xf7, 0x62, 0x97, 0xa2, 0xfe, 0x58, 0x76, 0xaa, 0xf6, 0x24, 0xcd, 0xcc, 0xe3, 0xe0, 0xcd,
	0xdb, 0xd9, 0xb7, 0xf0, 0x40, 0x69, 0x64, 0x69, 0x8a, 0x4a, 0xb1, 0x2e, 0xaa, 0x20, 0x4f, 0xd8,
	0x10, 0xa5, 0x6f, 0x93, 0x61, 0xc2, 0x31, 0xd3, 0x7e, 0x2e, 0x85, 0x16, 0xde, 0xc1, 0x22, 0xae,
	0x9f, 0xf1, 0x98, 0x63, 0x14, 0x74, 0x98, 0xc2, 0x65, 0x74, 0xf3, 0x19, 0x34, 0xdb, 0xa7, 0xb6,
	0x55, 0x70, 0x8c, 0xfa, 0x98, 0xa5, 0x78, 0xc4, 0xa2, 0x2e, 0x9e, 0xe0, 0x00, 0x13, 0x15, 0xbc,
	0xc2, 0xef, 0xfa, 0xa8, 0x34, 0xd9, 0x84, 0x2a, 0xcb, 0x73, 0x1e, 0xb9, 0x0e, 0x75, 0x1e, 0x6d,
	0x36, 0xc7, 0x0e, 0xec, 0x7d, 0xf0, 0x2b, 0x95, 0x8b, 0x4c, 0x21, 0xd9, 0x86, 0x8d, 0x82, 0x66,
	0x90, 0x98, 0x4a, 0xf1, 0x35, 0xf9, 0x1c, 0xd6, 0x3b, 0x06, 0xad, 0xdc, 0x6b, 0xb4, 0xf2, 0xa8,
	0xd1, 0xda, 0xf7, 0x57, 0xe8, 0xe5, 0xdb, 0xa4, 0xf7, 0x25, 0x54, 0xed, 0x1f, 0xc3, 0x68, 0xd6,
	0xb3, 0x4a, 0x6e, 0xc0, 0xba, 0x42, 0xc9, 0x6d, 0x4f, 0x13, 0x6f, 0xc3, 0x46, 0x47, 0xc8, 0x08,
	0x65, 0x10, 0x8a, 0x44, 0x48, 0xb7, 0x62, 0x79, 0x8f, 0x9c, 0x85, 0x69, 0x4f, 0x98, 0xd2, 0x36,
	0x8a, 0x5e, 0xf3, 0x14, 0x67, 0xd3, 0x7e, 0x03, 0x5b, 0x29, 0xcf, 0x82, 0x84, 0x29, 0x5d, 0xc8,
	0x3c, 0x99, 0xfb, 0xa8, 0x3d, 0x1a, 0xbb, 0x2f, 0x5e, 0xf7, 0x90, 0xa6, 0x42, 0x69, 0x2a, 0x31,
	0xc4, 0x4c, 0x53, 0x03, 0x7b, 0x5a, 0xc0, 0xa8, 0xe6, 0x29, 0x52, 0xdd, 0x43, 0x5a, 0x68, 0x4c,
	0x59, 0x22, 0x91, 0x45, 0x43, 0xfa, 0x36, 0x13, 0xef, 0x14, 0x65, 0x1d, 0xd1, 0xd7, 0xcd, 0xf7,
	0x8b, 0xe2, 0x2d, 0x93, 0x98, 0x88, 0xf7, 0x19, 0x54, 0xbb, 0x2c, 0x45, 0xe5, 0x3a, 0x56, 0xa5,
	0xc7, 0xfe, 0x0a, 0x1f, 0xf9, 0x46, 0x3e, 0x2f, 0x80, 0x35, 0xf3, 0xbb, 0x78, 0x6a, 0x55, 0x72,
	0x07, 0x36, 0xa7, 0x23, 0x19, 0xa2, 0x56, 0xaa, 0x4d, 0xf2, 0x3f, 0xd8, 0x2a, 0x33, 0x41, 0xeb,
	0x1d, 0xe2, 0x5b, 0x65, 0xd5, 0xaa, 0x12, 0x17, 0x6e, 0x4e, 0x0b, 0xb1, 0x90, 0x38, 0x40, 0xe9,
	0xae, 0x99, 0x4a, 0x73, 0x07, 0xee, 0x95, 0x64, 0x0e, 0xc3, 0x10, 0x73, 0x7d, 0x76, 0x76, 0x58,
	0xaa, 0xd7, 0xdc, 0x05, 0xef, 0xb2, 0x62, 0x41, 0xb0, 0xf9, 0x2d, 0xdc, 0x2f, 0xab, 0x17, 0x87,
	0x78, 0x29, 0x34, 0x8f, 0x79, 0xc8, 0x34, 0x17, 0xd9, 0x7f, 0x9f, 0xbe, 0x79, 0x00, 0xfb, 0x13,
	0x6c, 0xdb, 0x1e, 0x82, 0xf9, 0xe2, 0x6c, 0xa8, 0x34, 0xa6, 0x5f, 0x65, 0xb1, 0x90, 0xa9, 0xed,
	0x3f, 0x65, 0xfb, 0xd7, 0x1a, 0xdc, 0x6a, 0x17, 0xc0, 0x19, 0x88, 0xec, 0x41, 0x25, 0xcc, 0xfb,
	0x56, 0xb7, 0x46, 0xeb, 0xae, 0xbf, 0x04, 0xf0, 0xdb, 0xa7, 0xe7, 0xe4, 0x63, 0x80, 0x01, 0x8f,
	0x50, 0x04, 0x21, 0x93, 0x91, 0x15, 0xb3, 0xd1, 0xda, 0xbd, 0x04, 0xfb, 0xc6, 0x80, 0xda, 0x4c,
	0x46, 0x46, 0x51, 0x91, 0xa3, 0x64, 0x9a, 0x67, 0xdd, 0x40, 0x59, 0x84, 0xd5, 0xba, 0x4e, 0x6e,
	0x42, 0x4d, 0xa8, 0xe0, 0xd3, 0xe7, 0x1d, 0xae, 0xad, 0xc6, 0x35, 0x73, 0x5a, 0x05, 0x22, 0x90,
	0x2c, 0x0d, 0xd2, 0x8e, 0x5b, 0x2d, 0x17, 0x9b, 0xf5, 0x23, 0x2e, 0x82, 0x08, 0x07, 0x3c, 0x44,
	0x77, 0xdd, 0x7e, 0xbe, 0x0b, 0xdb, 0x93, 0xac, 0xe4, 0x03, 0x94, 0xc1, 0x00, 0xa5, 0xe2, 0x22,
	0x73, 0xaf, 0x9b, 0xaa, 0xf7, 0x87, 0x03, 0x15, 0x43, 0xf8, 0x16, 0xd4, 0x55, 0x8e, 0x18, 0x05,
	0x69, 0xef, 0x87, 0xd9, 0xbd, 0x19, 0x60, 0x16, 0x09, 0x69, 0xf9, 0xd7, 0x89, 0x07, 0x24, 0x11,
	0x5d, 0x1e, 0xb2, 0x24, 0xc8, 0xa5, 0x08, 0x51, 0x29, 0x21, 0xcb, 0x7d, 0xd8, 0x81, 0xdb, 0x79,
	0x6f, 0xa8, 0x2e, 0x16, 0xed, 0x4a, 0x90, 0xbb, 0x70, 0xa3, 0x37, 0xcc, 0x51, 0xea, 0x9e, 0xd9,
	0x78, 0x9e, 0x75, 0x2d, 0xdf, 0x9a, 0xd9, 0xc1, 0x38, 0x4c, 0xc5, 0xc0, 0x12, 0xad, 0x91, 0x0d,
	0x58, 0x53, 0x0a, 0x5b, 0x96, 0x58, 0x19, 0x3d, 0x73, 0x6b, 0x25, 0x54, 0xd9, 0xb0, 0x3e, 0x0b,
	0xf1, 0x39, 0x73, 0x61, 0x3e, 0xfc, 0xc4, 0x6d, 0xcc, 0x87, 0x2d, 0x77, 0xc3, 0x84, 0xde, 0xaf,
	0x0e, 0xd4, 0x67, 0x3a, 0xdf, 0x80, 0xf5, 0x42, 0x08, 0x3b, 0x65, 0xdd, 0x90, 0xbb, 0x20, 0x4c,
	0x31, 0xed, 0x6d, 0x68, 0x4c, 0xf2, 0x11, 0xd3, 0x58, 0x98, 0x84, 0xb9, 0x0f, 0x11, 0x97, 0x18,
	0xea, 0xef, 0xa7, 0xe8, 0xb5, 0xb2, 0x8b, 0xc8, 0x31, 0xeb, 0x26, 0xd3, 0x7c, 0xb5, 0x3c, 0xbb,
	0x42, 0x43, 0x1e, 0xd9, 0x29, 0xab, 0x26, 0x53, 0x1c, 0x0f, 0x8f, 0xec, 0xa4, 0x55, 0xb2, 0x05,
	0xd7, 0x07, 0x93, 0x73, 0xac, 0xd9, 0x2b, 0xf4, 0x06, 0x9e, 0xac, 0xb4, 0xa5, 0x13, 0x33, 0x78,
	0x08, 0x8d, 0xc9, 0x36, 0xf0, 0x2c, 0x16, 0x93, 0xc5, 0x24, 0xcb, 0xcb, 0xd6, 0xfa, 0xb3, 0x02,
	0xeb, 0x45, 0x5f, 0xf2, 0x9b, 0x03, 0x64, 0xd9, 0x51, 0xc9, 0x9e, 0xff, 0xcf, 0x86, 0xef, 0xdd,
	0x5f, 0xc5, 0x93, 0x9b, 0xe7, 0xa3, 0xb1, 0xfb, 0xf5, 0x2b, 0xd4, 0x7d, 0x99, 0x29, 0x6b, 0x7b,
	0x67, 0xe6, 0x7d, 0xa1, 0x16, 0x46, 0x45, 0x4c, 0x19, 0xed, 0x2b, 0x94, 0x07, 0xb6, 0x64, 0x1b,
	0x50, 0xeb, 0xd9, 0x34, 0x16, 0xd2, 0xe6, 0xcc, 0x0d, 0x3f, 0xa0, 0x2c, 0x8b, 0x28, 0x8f, 0x29,
	0xd7, 0x0f, 0x15, 0x8d, 0x05, 0x4f, 0xc8, 0x2f, 0x0e, 0xb8, 0xc5, 0x60, 0xcb, 0x97, 0x7d, 0x91,
	0xfe, 0x15, 0x0e, 0xbe, 0x48, 0xff, 0x2a, 0xbb, 0x68, 0xfa, 0xa3, 0xb1, 0xbb, 0x7f, 0x8c, 0xba,
	0xe0, 0x7e, 0xd1, 0xc7, 0xd5, 0x94, 0x26, 0x0b, 0x43, 0xd1, 0xcf, 0x34, 0x09, 0xa1, 0x3e, 0x75,
	0x34, 0xe2, 0xf9, 0x57, 0x5a, 0xa0, 0xb7, 0xe3, 0x7f, 0xc0, 0x01, 0xff, 0x3f, 0x1a, 0xbb, 0xf7,
	0xce, 0x15, 0x4a, 0xca, 0x95, 0x69, 0x8d, 0xb9, 0xb1, 0x83, 0x42, 0xbe, 0xb3, 0x43, 0xef, 0xe9,
	0x68, 0xec, 0x3e, 0x3e, 0xa4, 0x0a, 0xa5, 0x59, 0x20, 0xcb, 0xc1, 0x80, 0x94, 0x32, 0xa0, 0x42,
	0xdf, 0xe2, 0x4d, 0xa5, 0x11, 0xd3, 0xac, 0xf5, 0x73, 0x05, 0x36, 0xe6, 0xf7, 0x88, 0xfc, 0xe4,
	0xc0, 0x1d, 0xeb, 0xa4, 0xc3, 0x8b, 0xca, 0x7d, 0xe4, 0xaf, 0xe2, 0xbc, 0x5e, 0xc3, 0x7f, 0x29,
	0xa6, 0x64, 0x5f, 0x8c, 0xc6, 0xee, 0x17, 0xf3, 0x65, 0x1a, 0x4b, 0x91, 0x5a, 0x76, 0x28, 0xa9,
	0x16, 0xe5, 0x43, 0x27, 0x62, 0x9a, 0x0a, 0x89, 0xe5, 0x9b, 0x68, 0x18, 0x5a, 0x11, 0xc9, 0xef,
	0x0e, 0x6c, 0x5f, 0xb6, 0xda, 0xe4, 0x89, 0xbf, 0xba, 0x4b, 0x7b, 0x07, 0xfe, 0xbf, 0xb8, 0x2c,
	0xcd, 0xd3, 0xd1, 0xd8, 0x3d, 0x99, 0x7c, 0x5a, 0xf0, 0x35, 0xe2, 0xce, 0x38, 0xcf, 0x3d, 0xd0,
	0x46, 0x5e, 0x3e, 0xeb, 0x51, 0x3c, 0xd1, 0x16, 0x60, 0x36, 0x57, 0xd1, 0xe2, 0xd2, 0x79, 0x66,
	0x53, 0x1e, 0x2c, 0x89, 0x5f, 0x36, 0xc9, 0xe6, 0xf4, 0x51, 0xef, 0xc7, 0xee, 0xb5, 0xa3, 0xca,
	0x8f, 0x8e, 0xf3, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0x9e, 0xe2, 0xe2, 0x67, 0xb1, 0x09, 0x00,
	0x00,
}
