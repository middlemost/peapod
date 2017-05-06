// Code generated by protoc-gen-gogo.
// source: bolt.proto
// DO NOT EDIT!

/*
Package bolt is a generated protocol buffer package.

It is generated from these files:
	bolt.proto

It has these top-level messages:
	Job
	Playlist
	Track
	User
*/
package bolt

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type Job struct {
	ID         int64  `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	OwnerID    int64  `protobuf:"varint,2,opt,name=OwnerID,proto3" json:"OwnerID,omitempty"`
	Type       string `protobuf:"bytes,3,opt,name=Type,proto3" json:"Type,omitempty"`
	Status     string `protobuf:"bytes,4,opt,name=Status,proto3" json:"Status,omitempty"`
	PlaylistID int64  `protobuf:"varint,5,opt,name=PlaylistID,proto3" json:"PlaylistID,omitempty"`
	Title      string `protobuf:"bytes,10,opt,name=Title,proto3" json:"Title,omitempty"`
	URL        string `protobuf:"bytes,6,opt,name=URL,proto3" json:"URL,omitempty"`
	Text       string `protobuf:"bytes,11,opt,name=Text,proto3" json:"Text,omitempty"`
	Error      string `protobuf:"bytes,7,opt,name=Error,proto3" json:"Error,omitempty"`
	CreatedAt  int64  `protobuf:"varint,8,opt,name=CreatedAt,proto3" json:"CreatedAt,omitempty"`
	UpdatedAt  int64  `protobuf:"varint,9,opt,name=UpdatedAt,proto3" json:"UpdatedAt,omitempty"`
}

func (m *Job) Reset()                    { *m = Job{} }
func (m *Job) String() string            { return proto.CompactTextString(m) }
func (*Job) ProtoMessage()               {}
func (*Job) Descriptor() ([]byte, []int) { return fileDescriptorBolt, []int{0} }

type Playlist struct {
	ID        int64  `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	OwnerID   int64  `protobuf:"varint,2,opt,name=OwnerID,proto3" json:"OwnerID,omitempty"`
	Token     string `protobuf:"bytes,3,opt,name=Token,proto3" json:"Token,omitempty"`
	Name      string `protobuf:"bytes,4,opt,name=Name,proto3" json:"Name,omitempty"`
	CreatedAt int64  `protobuf:"varint,5,opt,name=CreatedAt,proto3" json:"CreatedAt,omitempty"`
	UpdatedAt int64  `protobuf:"varint,6,opt,name=UpdatedAt,proto3" json:"UpdatedAt,omitempty"`
}

func (m *Playlist) Reset()                    { *m = Playlist{} }
func (m *Playlist) String() string            { return proto.CompactTextString(m) }
func (*Playlist) ProtoMessage()               {}
func (*Playlist) Descriptor() ([]byte, []int) { return fileDescriptorBolt, []int{1} }

type Track struct {
	ID          int64  `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	PlaylistID  int64  `protobuf:"varint,2,opt,name=PlaylistID,proto3" json:"PlaylistID,omitempty"`
	Filename    string `protobuf:"bytes,3,opt,name=Filename,proto3" json:"Filename,omitempty"`
	ContentType string `protobuf:"bytes,4,opt,name=ContentType,proto3" json:"ContentType,omitempty"`
	Title       string `protobuf:"bytes,5,opt,name=Title,proto3" json:"Title,omitempty"`
	Description string `protobuf:"bytes,10,opt,name=Description,proto3" json:"Description,omitempty"`
	Duration    int64  `protobuf:"varint,6,opt,name=Duration,proto3" json:"Duration,omitempty"`
	FileSize    int64  `protobuf:"varint,7,opt,name=FileSize,proto3" json:"FileSize,omitempty"`
	CreatedAt   int64  `protobuf:"varint,8,opt,name=CreatedAt,proto3" json:"CreatedAt,omitempty"`
	UpdatedAt   int64  `protobuf:"varint,9,opt,name=UpdatedAt,proto3" json:"UpdatedAt,omitempty"`
}

func (m *Track) Reset()                    { *m = Track{} }
func (m *Track) String() string            { return proto.CompactTextString(m) }
func (*Track) ProtoMessage()               {}
func (*Track) Descriptor() ([]byte, []int) { return fileDescriptorBolt, []int{2} }

type User struct {
	ID           int64  `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	MobileNumber string `protobuf:"bytes,2,opt,name=MobileNumber,proto3" json:"MobileNumber,omitempty"`
	CreatedAt    int64  `protobuf:"varint,3,opt,name=CreatedAt,proto3" json:"CreatedAt,omitempty"`
	UpdatedAt    int64  `protobuf:"varint,4,opt,name=UpdatedAt,proto3" json:"UpdatedAt,omitempty"`
}

func (m *User) Reset()                    { *m = User{} }
func (m *User) String() string            { return proto.CompactTextString(m) }
func (*User) ProtoMessage()               {}
func (*User) Descriptor() ([]byte, []int) { return fileDescriptorBolt, []int{3} }

func init() {
	proto.RegisterType((*Job)(nil), "bolt.Job")
	proto.RegisterType((*Playlist)(nil), "bolt.Playlist")
	proto.RegisterType((*Track)(nil), "bolt.Track")
	proto.RegisterType((*User)(nil), "bolt.User")
}

func init() { proto.RegisterFile("bolt.proto", fileDescriptorBolt) }

var fileDescriptorBolt = []byte{
	// 376 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xa4, 0x53, 0xcb, 0x6a, 0xe3, 0x40,
	0x10, 0x44, 0x4f, 0x5b, 0xed, 0x65, 0x59, 0x86, 0x65, 0x19, 0x96, 0x10, 0x84, 0x4f, 0x3e, 0xe5,
	0x92, 0x2f, 0x08, 0x56, 0x02, 0x0a, 0x89, 0x13, 0x64, 0xfb, 0x03, 0x24, 0xbb, 0x0f, 0xc2, 0xb2,
	0x46, 0x8c, 0xc6, 0x89, 0x9d, 0x3f, 0xc8, 0x2f, 0xe4, 0x96, 0x3f, 0x0d, 0xd3, 0x92, 0x65, 0xc9,
	0x06, 0x43, 0xc8, 0xad, 0xab, 0x9a, 0x56, 0xd5, 0x54, 0x21, 0x80, 0x44, 0x64, 0xea, 0xaa, 0x90,
	0x42, 0x09, 0x66, 0xeb, 0x79, 0xf8, 0x6e, 0x82, 0x75, 0x2f, 0x12, 0xf6, 0x1b, 0xcc, 0x30, 0xe0,
	0x86, 0x6f, 0x8c, 0xac, 0xc8, 0x0c, 0x03, 0xc6, 0xa1, 0xf7, 0xf4, 0x9a, 0xa3, 0x0c, 0x03, 0x6e,
	0x12, 0xb9, 0x87, 0x8c, 0x81, 0x3d, 0xdb, 0x15, 0xc8, 0x2d, 0xdf, 0x18, 0x79, 0x11, 0xcd, 0xec,
	0x1f, 0xb8, 0x53, 0x15, 0xab, 0x4d, 0xc9, 0x6d, 0x62, 0x6b, 0xc4, 0x2e, 0x01, 0x9e, 0xb3, 0x78,
	0x97, 0xa5, 0xa5, 0x0a, 0x03, 0xee, 0xd0, 0x87, 0x5a, 0x0c, 0xfb, 0x0b, 0xce, 0x2c, 0x55, 0x19,
	0x72, 0xa0, 0xb3, 0x0a, 0xb0, 0x3f, 0x60, 0xcd, 0xa3, 0x07, 0xee, 0x12, 0xa7, 0x47, 0xd2, 0xc4,
	0xad, 0xe2, 0x83, 0x5a, 0x13, 0xb7, 0x4a, 0xdf, 0xde, 0x4a, 0x29, 0x24, 0xef, 0x55, 0xb7, 0x04,
	0xd8, 0x05, 0x78, 0x63, 0x89, 0xb1, 0xc2, 0xe5, 0x8d, 0xe2, 0x7d, 0x12, 0x3c, 0x10, 0x7a, 0x3b,
	0x2f, 0x96, 0xf5, 0xd6, 0xab, 0xb6, 0x0d, 0x31, 0xfc, 0x30, 0xa0, 0xbf, 0x37, 0xf7, 0x8d, 0x40,
	0xf4, 0x23, 0xc4, 0x0a, 0xf3, 0x3a, 0x91, 0x0a, 0x68, 0xcb, 0x93, 0x78, 0x8d, 0x75, 0x20, 0x34,
	0x77, 0xcd, 0x39, 0x67, 0xcd, 0xb9, 0xc7, 0xe6, 0x3e, 0x4d, 0x70, 0x66, 0x32, 0x5e, 0xac, 0x4e,
	0x9c, 0x75, 0x43, 0x36, 0x4f, 0x42, 0xfe, 0x0f, 0xfd, 0xbb, 0x34, 0xc3, 0x5c, 0xbb, 0xa9, 0x2c,
	0x36, 0x98, 0xf9, 0x30, 0x18, 0x8b, 0x5c, 0x61, 0xae, 0xa8, 0xd3, 0xca, 0x6c, 0x9b, 0x3a, 0x54,
	0xe4, 0xb4, 0x2b, 0xf2, 0x61, 0x10, 0x60, 0xb9, 0x90, 0x69, 0xa1, 0x52, 0x91, 0xd7, 0xf5, 0xb5,
	0x29, 0xad, 0x1a, 0x6c, 0x64, 0x4c, 0xeb, 0xea, 0x31, 0x0d, 0xde, 0x3b, 0x9a, 0xa6, 0x6f, 0x48,
	0xed, 0x59, 0x51, 0x83, 0x7f, 0x54, 0xe0, 0x0b, 0xd8, 0xf3, 0x12, 0xe5, 0x49, 0x42, 0x43, 0xf8,
	0xf5, 0x28, 0x92, 0x34, 0xc3, 0xc9, 0x66, 0x9d, 0xa0, 0xa4, 0x8c, 0xbc, 0xa8, 0xc3, 0x75, 0x75,
	0xad, 0xb3, 0xba, 0xf6, 0x91, 0x6e, 0xe2, 0xd2, 0x1f, 0x75, 0xfd, 0x15, 0x00, 0x00, 0xff, 0xff,
	0x14, 0x1e, 0x30, 0x83, 0x5f, 0x03, 0x00, 0x00,
}
