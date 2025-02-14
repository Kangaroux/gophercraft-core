// Code generated by protoc-gen-gcraft. DO NOT EDIT.
// source: club_ban.proto

package v1

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/Gophercraft/core/bnet/bgs/protocol"
	v2 "github.com/Gophercraft/core/bnet/bgs/protocol/v2"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type AddBanOptions struct {
	TargetId             *MemberId       `protobuf:"bytes,1,opt,name=target_id,json=targetId" json:"target_id,omitempty"`
	Attribute            []*v2.Attribute `protobuf:"bytes,2,rep,name=attribute" json:"attribute,omitempty"`
	Reason               *string         `protobuf:"bytes,3,opt,name=reason" json:"reason,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *AddBanOptions) Reset()         { *m = AddBanOptions{} }
func (m *AddBanOptions) String() string { return proto.CompactTextString(m) }
func (*AddBanOptions) ProtoMessage()    {}
func (*AddBanOptions) Descriptor() ([]byte, []int) {
	return fileDescriptor_ee1cc033b0fb491a, []int{0}
}

func (m *AddBanOptions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddBanOptions.Unmarshal(m, b)
}
func (m *AddBanOptions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddBanOptions.Marshal(b, m, deterministic)
}
func (m *AddBanOptions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddBanOptions.Merge(m, src)
}
func (m *AddBanOptions) XXX_Size() int {
	return xxx_messageInfo_AddBanOptions.Size(m)
}
func (m *AddBanOptions) XXX_DiscardUnknown() {
	xxx_messageInfo_AddBanOptions.DiscardUnknown(m)
}

var xxx_messageInfo_AddBanOptions proto.InternalMessageInfo

func (m *AddBanOptions) GetTargetId() *MemberId {
	if m != nil {
		return m.TargetId
	}
	return nil
}

func (m *AddBanOptions) GetAttribute() []*v2.Attribute {
	if m != nil {
		return m.Attribute
	}
	return nil
}

func (m *AddBanOptions) GetReason() string {
	if m != nil && m.Reason != nil {
		return *m.Reason
	}
	return ""
}

type ClubBan struct {
	Id                   *MemberId          `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	BattleTag            *string            `protobuf:"bytes,2,opt,name=battle_tag,json=battleTag" json:"battle_tag,omitempty"`
	Creator              *MemberDescription `protobuf:"bytes,3,opt,name=creator" json:"creator,omitempty"`
	Attribute            []*v2.Attribute    `protobuf:"bytes,4,rep,name=attribute" json:"attribute,omitempty"`
	Reason               *string            `protobuf:"bytes,5,opt,name=reason" json:"reason,omitempty"`
	CreationTime         *uint64            `protobuf:"varint,6,opt,name=creation_time,json=creationTime" json:"creation_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *ClubBan) Reset()         { *m = ClubBan{} }
func (m *ClubBan) String() string { return proto.CompactTextString(m) }
func (*ClubBan) ProtoMessage()    {}
func (*ClubBan) Descriptor() ([]byte, []int) {
	return fileDescriptor_ee1cc033b0fb491a, []int{1}
}

func (m *ClubBan) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClubBan.Unmarshal(m, b)
}
func (m *ClubBan) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClubBan.Marshal(b, m, deterministic)
}
func (m *ClubBan) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClubBan.Merge(m, src)
}
func (m *ClubBan) XXX_Size() int {
	return xxx_messageInfo_ClubBan.Size(m)
}
func (m *ClubBan) XXX_DiscardUnknown() {
	xxx_messageInfo_ClubBan.DiscardUnknown(m)
}

var xxx_messageInfo_ClubBan proto.InternalMessageInfo

func (m *ClubBan) GetId() *MemberId {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *ClubBan) GetBattleTag() string {
	if m != nil && m.BattleTag != nil {
		return *m.BattleTag
	}
	return ""
}

func (m *ClubBan) GetCreator() *MemberDescription {
	if m != nil {
		return m.Creator
	}
	return nil
}

func (m *ClubBan) GetAttribute() []*v2.Attribute {
	if m != nil {
		return m.Attribute
	}
	return nil
}

func (m *ClubBan) GetReason() string {
	if m != nil && m.Reason != nil {
		return *m.Reason
	}
	return ""
}

func (m *ClubBan) GetCreationTime() uint64 {
	if m != nil && m.CreationTime != nil {
		return *m.CreationTime
	}
	return 0
}

func init() {
	proto.RegisterType((*AddBanOptions)(nil), "bgs.protocol.club.v1.AddBanOptions")
	proto.RegisterType((*ClubBan)(nil), "bgs.protocol.club.v1.ClubBan")
}

func init() { proto.RegisterFile("club_ban.proto", fileDescriptor_ee1cc033b0fb491a) }

var fileDescriptor_ee1cc033b0fb491a = []byte{
	// 363 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x50, 0x5f, 0xab, 0xd3, 0x30,
	0x14, 0xa7, 0xbd, 0xd7, 0x7b, 0x6d, 0xe6, 0x04, 0x83, 0x48, 0x19, 0x28, 0x65, 0x43, 0xec, 0x53,
	0xb2, 0xf5, 0x49, 0xd4, 0x97, 0x4d, 0x5f, 0xf6, 0x20, 0x42, 0xd9, 0x93, 0x2f, 0x25, 0x69, 0xcf,
	0xb2, 0x40, 0x9a, 0x94, 0xf4, 0xb4, 0xe8, 0xd7, 0xf0, 0x03, 0xf8, 0x59, 0x65, 0xed, 0xaa, 0x0e,
	0x44, 0xc4, 0xb7, 0x9c, 0x5f, 0xce, 0xef, 0xdf, 0x21, 0x8f, 0x4b, 0xd3, 0xc9, 0x42, 0x0a, 0xcb,
	0x1a, 0xef, 0xd0, 0xd1, 0xa7, 0x52, 0xb5, 0xe3, 0xb3, 0x74, 0x86, 0x9d, 0x3f, 0x59, 0xbf, 0x59,
	0xbc, 0x54, 0xc6, 0x49, 0x61, 0x0a, 0xf8, 0x82, 0x60, 0x5b, 0xed, 0x6c, 0xcb, 0x8f, 0x1a, 0x4c,
	0x55, 0xb8, 0x06, 0xcf, 0xd3, 0xc8, 0x58, 0x3c, 0x19, 0xc4, 0x6a, 0xa8, 0x25, 0xf8, 0x0b, 0xb4,
	0x12, 0x8d, 0xe6, 0xa5, 0xd1, 0x60, 0x91, 0xf7, 0x19, 0x17, 0x88, 0x5e, 0xcb, 0x0e, 0xa1, 0xc0,
	0xaf, 0x0d, 0x5c, 0x78, 0xcb, 0xef, 0x01, 0x99, 0x6f, 0xab, 0x6a, 0x27, 0xec, 0xa7, 0x51, 0x8f,
	0xbe, 0x25, 0x11, 0x0a, 0xaf, 0x00, 0x0b, 0x5d, 0xc5, 0x41, 0x12, 0xa4, 0xb3, 0xec, 0x05, 0xfb,
	0x53, 0x34, 0xf6, 0x71, 0x70, 0xdb, 0x57, 0xf9, 0xc3, 0x91, 0xb0, 0xaf, 0xe8, 0x6b, 0x12, 0xfd,
	0xf4, 0x89, 0xc3, 0xe4, 0x26, 0x9d, 0x65, 0x8b, 0x6b, 0x72, 0x9f, 0xb1, 0xed, 0xb4, 0x91, 0xff,
	0x5a, 0xa6, 0xcf, 0xc8, 0x9d, 0x07, 0xd1, 0x3a, 0x1b, 0xdf, 0x24, 0x41, 0x1a, 0xe5, 0x97, 0x69,
	0xf9, 0x2d, 0x24, 0xf7, 0xef, 0x4d, 0x27, 0x77, 0xc2, 0x52, 0x46, 0xc2, 0x7f, 0xce, 0x14, 0xea,
	0x8a, 0x3e, 0x27, 0x44, 0x0a, 0x44, 0x03, 0x05, 0x0a, 0x15, 0x87, 0x83, 0x6e, 0x34, 0x22, 0x07,
	0xa1, 0xe8, 0x96, 0xdc, 0x97, 0x1e, 0x04, 0x3a, 0x3f, 0x78, 0xce, 0xb2, 0x57, 0x7f, 0xd3, 0xfc,
	0x00, 0x6d, 0xe9, 0xf5, 0x70, 0xa4, 0x7c, 0xe2, 0x5d, 0xf7, 0xbd, 0xfd, 0xbf, 0xbe, 0x0f, 0x7e,
	0xef, 0x4b, 0x57, 0x64, 0x3e, 0x88, 0x6b, 0x67, 0x0b, 0xd4, 0x35, 0xc4, 0x77, 0x49, 0x90, 0xde,
	0xe6, 0x8f, 0x26, 0xf0, 0xa0, 0x6b, 0xd8, 0xbd, 0xfb, 0xfc, 0x46, 0x69, 0x3c, 0x75, 0x92, 0x95,
	0xae, 0xe6, 0x6d, 0xd7, 0x80, 0x6f, 0xd6, 0x6b, 0xe4, 0xca, 0x35, 0x27, 0xf0, 0xa5, 0x17, 0x47,
	0xe4, 0xd2, 0x02, 0x72, 0xa9, 0x5a, 0x3e, 0x45, 0xe1, 0xe7, 0x3e, 0xbc, 0xdf, 0xfc, 0x08, 0x00,
	0x00, 0xff, 0xff, 0x17, 0x19, 0x27, 0x95, 0x79, 0x02, 0x00, 0x00,
}
