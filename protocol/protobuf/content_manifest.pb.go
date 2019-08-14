// Code generated by protoc-gen-go. DO NOT EDIT.
// source: content_manifest.proto

package protobuf

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
// is compatible with the proto package protobuf is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package protobuf to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type ContentManifestPayload struct {
	Mappings             []*ContentManifestPayload_FileMapping `protobuf:"bytes,1,rep,name=mappings" json:"mappings,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                              `json:"-"`
	XXX_unrecognized     []byte                                `json:"-"`
	XXX_sizecache        int32                                 `json:"-"`
}

func (m *ContentManifestPayload) Reset()         { *m = ContentManifestPayload{} }
func (m *ContentManifestPayload) String() string { return proto.CompactTextString(m) }
func (*ContentManifestPayload) ProtoMessage()    {}
func (*ContentManifestPayload) Descriptor() ([]byte, []int) {
	return fileDescriptor_e3cda137a29253ba, []int{0}
}

func (m *ContentManifestPayload) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ContentManifestPayload.Unmarshal(m, b)
}
func (m *ContentManifestPayload) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ContentManifestPayload.Marshal(b, m, deterministic)
}
func (m *ContentManifestPayload) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContentManifestPayload.Merge(m, src)
}
func (m *ContentManifestPayload) XXX_Size() int {
	return xxx_messageInfo_ContentManifestPayload.Size(m)
}
func (m *ContentManifestPayload) XXX_DiscardUnknown() {
	xxx_messageInfo_ContentManifestPayload.DiscardUnknown(m)
}

var xxx_messageInfo_ContentManifestPayload proto.InternalMessageInfo

func (m *ContentManifestPayload) GetMappings() []*ContentManifestPayload_FileMapping {
	if m != nil {
		return m.Mappings
	}
	return nil
}

type ContentManifestPayload_FileMapping struct {
	Filename             *string                                         `protobuf:"bytes,1,opt,name=filename" json:"filename,omitempty"`
	Size                 *uint64                                         `protobuf:"varint,2,opt,name=size" json:"size,omitempty"`
	Flags                *uint32                                         `protobuf:"varint,3,opt,name=flags" json:"flags,omitempty"`
	ShaFilename          []byte                                          `protobuf:"bytes,4,opt,name=sha_filename" json:"sha_filename,omitempty"`
	ShaContent           []byte                                          `protobuf:"bytes,5,opt,name=sha_content" json:"sha_content,omitempty"`
	Chunks               []*ContentManifestPayload_FileMapping_ChunkData `protobuf:"bytes,6,rep,name=chunks" json:"chunks,omitempty"`
	Linktarget           *string                                         `protobuf:"bytes,7,opt,name=linktarget" json:"linktarget,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                        `json:"-"`
	XXX_unrecognized     []byte                                          `json:"-"`
	XXX_sizecache        int32                                           `json:"-"`
}

func (m *ContentManifestPayload_FileMapping) Reset()         { *m = ContentManifestPayload_FileMapping{} }
func (m *ContentManifestPayload_FileMapping) String() string { return proto.CompactTextString(m) }
func (*ContentManifestPayload_FileMapping) ProtoMessage()    {}
func (*ContentManifestPayload_FileMapping) Descriptor() ([]byte, []int) {
	return fileDescriptor_e3cda137a29253ba, []int{0, 0}
}

func (m *ContentManifestPayload_FileMapping) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ContentManifestPayload_FileMapping.Unmarshal(m, b)
}
func (m *ContentManifestPayload_FileMapping) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ContentManifestPayload_FileMapping.Marshal(b, m, deterministic)
}
func (m *ContentManifestPayload_FileMapping) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContentManifestPayload_FileMapping.Merge(m, src)
}
func (m *ContentManifestPayload_FileMapping) XXX_Size() int {
	return xxx_messageInfo_ContentManifestPayload_FileMapping.Size(m)
}
func (m *ContentManifestPayload_FileMapping) XXX_DiscardUnknown() {
	xxx_messageInfo_ContentManifestPayload_FileMapping.DiscardUnknown(m)
}

var xxx_messageInfo_ContentManifestPayload_FileMapping proto.InternalMessageInfo

func (m *ContentManifestPayload_FileMapping) GetFilename() string {
	if m != nil && m.Filename != nil {
		return *m.Filename
	}
	return ""
}

func (m *ContentManifestPayload_FileMapping) GetSize() uint64 {
	if m != nil && m.Size != nil {
		return *m.Size
	}
	return 0
}

func (m *ContentManifestPayload_FileMapping) GetFlags() uint32 {
	if m != nil && m.Flags != nil {
		return *m.Flags
	}
	return 0
}

func (m *ContentManifestPayload_FileMapping) GetShaFilename() []byte {
	if m != nil {
		return m.ShaFilename
	}
	return nil
}

func (m *ContentManifestPayload_FileMapping) GetShaContent() []byte {
	if m != nil {
		return m.ShaContent
	}
	return nil
}

func (m *ContentManifestPayload_FileMapping) GetChunks() []*ContentManifestPayload_FileMapping_ChunkData {
	if m != nil {
		return m.Chunks
	}
	return nil
}

func (m *ContentManifestPayload_FileMapping) GetLinktarget() string {
	if m != nil && m.Linktarget != nil {
		return *m.Linktarget
	}
	return ""
}

type ContentManifestPayload_FileMapping_ChunkData struct {
	Sha                  []byte   `protobuf:"bytes,1,opt,name=sha" json:"sha,omitempty"`
	Crc                  *uint32  `protobuf:"fixed32,2,opt,name=crc" json:"crc,omitempty"`
	Offset               *uint64  `protobuf:"varint,3,opt,name=offset" json:"offset,omitempty"`
	CbOriginal           *uint32  `protobuf:"varint,4,opt,name=cb_original" json:"cb_original,omitempty"`
	CbCompressed         *uint32  `protobuf:"varint,5,opt,name=cb_compressed" json:"cb_compressed,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ContentManifestPayload_FileMapping_ChunkData) Reset() {
	*m = ContentManifestPayload_FileMapping_ChunkData{}
}
func (m *ContentManifestPayload_FileMapping_ChunkData) String() string {
	return proto.CompactTextString(m)
}
func (*ContentManifestPayload_FileMapping_ChunkData) ProtoMessage() {}
func (*ContentManifestPayload_FileMapping_ChunkData) Descriptor() ([]byte, []int) {
	return fileDescriptor_e3cda137a29253ba, []int{0, 0, 0}
}

func (m *ContentManifestPayload_FileMapping_ChunkData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ContentManifestPayload_FileMapping_ChunkData.Unmarshal(m, b)
}
func (m *ContentManifestPayload_FileMapping_ChunkData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ContentManifestPayload_FileMapping_ChunkData.Marshal(b, m, deterministic)
}
func (m *ContentManifestPayload_FileMapping_ChunkData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContentManifestPayload_FileMapping_ChunkData.Merge(m, src)
}
func (m *ContentManifestPayload_FileMapping_ChunkData) XXX_Size() int {
	return xxx_messageInfo_ContentManifestPayload_FileMapping_ChunkData.Size(m)
}
func (m *ContentManifestPayload_FileMapping_ChunkData) XXX_DiscardUnknown() {
	xxx_messageInfo_ContentManifestPayload_FileMapping_ChunkData.DiscardUnknown(m)
}

var xxx_messageInfo_ContentManifestPayload_FileMapping_ChunkData proto.InternalMessageInfo

func (m *ContentManifestPayload_FileMapping_ChunkData) GetSha() []byte {
	if m != nil {
		return m.Sha
	}
	return nil
}

func (m *ContentManifestPayload_FileMapping_ChunkData) GetCrc() uint32 {
	if m != nil && m.Crc != nil {
		return *m.Crc
	}
	return 0
}

func (m *ContentManifestPayload_FileMapping_ChunkData) GetOffset() uint64 {
	if m != nil && m.Offset != nil {
		return *m.Offset
	}
	return 0
}

func (m *ContentManifestPayload_FileMapping_ChunkData) GetCbOriginal() uint32 {
	if m != nil && m.CbOriginal != nil {
		return *m.CbOriginal
	}
	return 0
}

func (m *ContentManifestPayload_FileMapping_ChunkData) GetCbCompressed() uint32 {
	if m != nil && m.CbCompressed != nil {
		return *m.CbCompressed
	}
	return 0
}

type ContentManifestMetadata struct {
	DepotId              *uint32  `protobuf:"varint,1,opt,name=depot_id" json:"depot_id,omitempty"`
	GidManifest          *uint64  `protobuf:"varint,2,opt,name=gid_manifest" json:"gid_manifest,omitempty"`
	CreationTime         *uint32  `protobuf:"varint,3,opt,name=creation_time" json:"creation_time,omitempty"`
	FilenamesEncrypted   *bool    `protobuf:"varint,4,opt,name=filenames_encrypted" json:"filenames_encrypted,omitempty"`
	CbDiskOriginal       *uint64  `protobuf:"varint,5,opt,name=cb_disk_original" json:"cb_disk_original,omitempty"`
	CbDiskCompressed     *uint64  `protobuf:"varint,6,opt,name=cb_disk_compressed" json:"cb_disk_compressed,omitempty"`
	UniqueChunks         *uint32  `protobuf:"varint,7,opt,name=unique_chunks" json:"unique_chunks,omitempty"`
	CrcEncrypted         *uint32  `protobuf:"varint,8,opt,name=crc_encrypted" json:"crc_encrypted,omitempty"`
	CrcClear             *uint32  `protobuf:"varint,9,opt,name=crc_clear" json:"crc_clear,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ContentManifestMetadata) Reset()         { *m = ContentManifestMetadata{} }
func (m *ContentManifestMetadata) String() string { return proto.CompactTextString(m) }
func (*ContentManifestMetadata) ProtoMessage()    {}
func (*ContentManifestMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_e3cda137a29253ba, []int{1}
}

func (m *ContentManifestMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ContentManifestMetadata.Unmarshal(m, b)
}
func (m *ContentManifestMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ContentManifestMetadata.Marshal(b, m, deterministic)
}
func (m *ContentManifestMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContentManifestMetadata.Merge(m, src)
}
func (m *ContentManifestMetadata) XXX_Size() int {
	return xxx_messageInfo_ContentManifestMetadata.Size(m)
}
func (m *ContentManifestMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_ContentManifestMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_ContentManifestMetadata proto.InternalMessageInfo

func (m *ContentManifestMetadata) GetDepotId() uint32 {
	if m != nil && m.DepotId != nil {
		return *m.DepotId
	}
	return 0
}

func (m *ContentManifestMetadata) GetGidManifest() uint64 {
	if m != nil && m.GidManifest != nil {
		return *m.GidManifest
	}
	return 0
}

func (m *ContentManifestMetadata) GetCreationTime() uint32 {
	if m != nil && m.CreationTime != nil {
		return *m.CreationTime
	}
	return 0
}

func (m *ContentManifestMetadata) GetFilenamesEncrypted() bool {
	if m != nil && m.FilenamesEncrypted != nil {
		return *m.FilenamesEncrypted
	}
	return false
}

func (m *ContentManifestMetadata) GetCbDiskOriginal() uint64 {
	if m != nil && m.CbDiskOriginal != nil {
		return *m.CbDiskOriginal
	}
	return 0
}

func (m *ContentManifestMetadata) GetCbDiskCompressed() uint64 {
	if m != nil && m.CbDiskCompressed != nil {
		return *m.CbDiskCompressed
	}
	return 0
}

func (m *ContentManifestMetadata) GetUniqueChunks() uint32 {
	if m != nil && m.UniqueChunks != nil {
		return *m.UniqueChunks
	}
	return 0
}

func (m *ContentManifestMetadata) GetCrcEncrypted() uint32 {
	if m != nil && m.CrcEncrypted != nil {
		return *m.CrcEncrypted
	}
	return 0
}

func (m *ContentManifestMetadata) GetCrcClear() uint32 {
	if m != nil && m.CrcClear != nil {
		return *m.CrcClear
	}
	return 0
}

type ContentManifestSignature struct {
	Signature            []byte   `protobuf:"bytes,1,opt,name=signature" json:"signature,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ContentManifestSignature) Reset()         { *m = ContentManifestSignature{} }
func (m *ContentManifestSignature) String() string { return proto.CompactTextString(m) }
func (*ContentManifestSignature) ProtoMessage()    {}
func (*ContentManifestSignature) Descriptor() ([]byte, []int) {
	return fileDescriptor_e3cda137a29253ba, []int{2}
}

func (m *ContentManifestSignature) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ContentManifestSignature.Unmarshal(m, b)
}
func (m *ContentManifestSignature) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ContentManifestSignature.Marshal(b, m, deterministic)
}
func (m *ContentManifestSignature) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContentManifestSignature.Merge(m, src)
}
func (m *ContentManifestSignature) XXX_Size() int {
	return xxx_messageInfo_ContentManifestSignature.Size(m)
}
func (m *ContentManifestSignature) XXX_DiscardUnknown() {
	xxx_messageInfo_ContentManifestSignature.DiscardUnknown(m)
}

var xxx_messageInfo_ContentManifestSignature proto.InternalMessageInfo

func (m *ContentManifestSignature) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

func init() {
	proto.RegisterType((*ContentManifestPayload)(nil), "ContentManifestPayload")
	proto.RegisterType((*ContentManifestPayload_FileMapping)(nil), "ContentManifestPayload.FileMapping")
	proto.RegisterType((*ContentManifestPayload_FileMapping_ChunkData)(nil), "ContentManifestPayload.FileMapping.ChunkData")
	proto.RegisterType((*ContentManifestMetadata)(nil), "ContentManifestMetadata")
	proto.RegisterType((*ContentManifestSignature)(nil), "ContentManifestSignature")
}

func init() { proto.RegisterFile("content_manifest.proto", fileDescriptor_e3cda137a29253ba) }

var fileDescriptor_e3cda137a29253ba = []byte{
	// 412 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0xcd, 0x8e, 0xd3, 0x30,
	0x14, 0x85, 0x49, 0x7f, 0xd3, 0xdb, 0x06, 0x15, 0x0f, 0x0c, 0x56, 0xd9, 0x44, 0xc3, 0x26, 0x9b,
	0xe9, 0x02, 0x89, 0x25, 0x1b, 0x06, 0x21, 0x36, 0x95, 0x90, 0x78, 0x00, 0xeb, 0xd6, 0xb9, 0x49,
	0xad, 0x26, 0x76, 0xb0, 0xdd, 0xc5, 0xb0, 0xe2, 0x45, 0x78, 0x43, 0x24, 0x5e, 0x01, 0xd9, 0x4d,
	0x3b, 0x55, 0xc5, 0x82, 0xe5, 0xfd, 0xf1, 0x3d, 0xdf, 0x39, 0x86, 0x5b, 0x69, 0xb4, 0x27, 0xed,
	0x45, 0x8b, 0x5a, 0x55, 0xe4, 0xfc, 0xba, 0xb3, 0xc6, 0x9b, 0xbb, 0x3f, 0x03, 0xb8, 0x7d, 0x38,
	0x8e, 0x36, 0xfd, 0xe4, 0x2b, 0x3e, 0x36, 0x06, 0x4b, 0xf6, 0x1e, 0xd2, 0x16, 0xbb, 0x4e, 0xe9,
	0xda, 0xf1, 0x24, 0x1f, 0x16, 0xf3, 0x77, 0x6f, 0xd7, 0xff, 0x5e, 0x5d, 0x7f, 0x56, 0x0d, 0x6d,
	0x8e, 0xbb, 0xab, 0x5f, 0x03, 0x98, 0x5f, 0xd4, 0x6c, 0x09, 0x69, 0xa5, 0x1a, 0xd2, 0xd8, 0x12,
	0x4f, 0xf2, 0xa4, 0x98, 0xb1, 0x05, 0x8c, 0x9c, 0xfa, 0x41, 0x7c, 0x90, 0x27, 0xc5, 0x88, 0x65,
	0x30, 0xae, 0x1a, 0xac, 0x1d, 0x1f, 0xe6, 0x49, 0x91, 0xb1, 0x97, 0xb0, 0x70, 0x3b, 0x14, 0xe7,
	0x27, 0xa3, 0x3c, 0x29, 0x16, 0xec, 0x06, 0xe6, 0xa1, 0xdb, 0x9b, 0xe0, 0xe3, 0xd8, 0xfc, 0x00,
	0x13, 0xb9, 0x3b, 0xe8, 0xbd, 0xe3, 0x93, 0x88, 0x77, 0xff, 0x1f, 0x78, 0xeb, 0x87, 0xf0, 0xe2,
	0x13, 0x7a, 0x64, 0x0c, 0xa0, 0x51, 0x7a, 0xef, 0xd1, 0xd6, 0xe4, 0xf9, 0x34, 0xa0, 0xad, 0x10,
	0x66, 0x4f, 0x0b, 0x73, 0x18, 0xba, 0x1d, 0x46, 0xe8, 0x45, 0x28, 0xa4, 0x95, 0x91, 0x79, 0xca,
	0x9e, 0xc3, 0xc4, 0x54, 0x95, 0x23, 0x1f, 0xa1, 0x47, 0x01, 0x4f, 0x6e, 0x85, 0xb1, 0xaa, 0x56,
	0x1a, 0x9b, 0xc8, 0x9c, 0xb1, 0x57, 0x90, 0xc9, 0xad, 0x90, 0xa6, 0xed, 0x2c, 0x39, 0x47, 0x65,
	0xa4, 0xce, 0xee, 0x7e, 0x27, 0xf0, 0xfa, 0x8a, 0x73, 0x43, 0x1e, 0xcb, 0xa0, 0xb8, 0x84, 0xb4,
	0xa4, 0xce, 0x78, 0xa1, 0xca, 0x28, 0x1b, 0xe3, 0xa8, 0x55, 0x79, 0xfe, 0xb5, 0x3e, 0xb3, 0x70,
	0xda, 0x12, 0x7a, 0x65, 0xb4, 0xf0, 0xaa, 0xa5, 0x3e, 0xbb, 0x37, 0x70, 0x73, 0xca, 0xcd, 0x09,
	0xd2, 0xd2, 0x3e, 0x76, 0x9e, 0xca, 0x88, 0x93, 0x32, 0x0e, 0x4b, 0xb9, 0x15, 0xa5, 0x72, 0xfb,
	0x27, 0xd0, 0x71, 0xbc, 0xb6, 0x02, 0x76, 0x9a, 0x5c, 0xd0, 0x4e, 0x4e, 0x4a, 0x07, 0xad, 0xbe,
	0x1f, 0x48, 0xf4, 0x51, 0x4f, 0xcf, 0xde, 0xac, 0xbc, 0xd0, 0x48, 0x63, 0xfb, 0x05, 0xcc, 0x42,
	0x5b, 0x36, 0x84, 0x96, 0xcf, 0xa2, 0xdd, 0x7b, 0xe0, 0x57, 0x6e, 0xbf, 0xa9, 0x5a, 0xa3, 0x3f,
	0x58, 0x0a, 0xeb, 0xee, 0x54, 0x1c, 0x63, 0xfe, 0x38, 0xfe, 0x92, 0xfc, 0x4c, 0x9e, 0xfd, 0x0d,
	0x00, 0x00, 0xff, 0xff, 0xc6, 0x87, 0xdb, 0xe6, 0xaf, 0x02, 0x00, 0x00,
}
