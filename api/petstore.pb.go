// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/alanchchen/go-example-twirp/api/petstore.proto

package api

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import empty "github.com/golang/protobuf/ptypes/empty"
import _ "google.golang.org/genproto/googleapis/api/annotations"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type AddPetRequest struct {
	// Pet to add to the store
	Pet                  *AddPetRequest_PetMessage `protobuf:"bytes,1,opt,name=pet,proto3" json:"pet,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *AddPetRequest) Reset()         { *m = AddPetRequest{} }
func (m *AddPetRequest) String() string { return proto.CompactTextString(m) }
func (*AddPetRequest) ProtoMessage()    {}
func (*AddPetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_petstore_648dbf3c11f4d833, []int{0}
}
func (m *AddPetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddPetRequest.Unmarshal(m, b)
}
func (m *AddPetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddPetRequest.Marshal(b, m, deterministic)
}
func (dst *AddPetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddPetRequest.Merge(dst, src)
}
func (m *AddPetRequest) XXX_Size() int {
	return xxx_messageInfo_AddPetRequest.Size(m)
}
func (m *AddPetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AddPetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AddPetRequest proto.InternalMessageInfo

func (m *AddPetRequest) GetPet() *AddPetRequest_PetMessage {
	if m != nil {
		return m.Pet
	}
	return nil
}

type AddPetRequest_PetMessage struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Tag                  string   `protobuf:"bytes,3,opt,name=tag,proto3" json:"tag,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddPetRequest_PetMessage) Reset()         { *m = AddPetRequest_PetMessage{} }
func (m *AddPetRequest_PetMessage) String() string { return proto.CompactTextString(m) }
func (*AddPetRequest_PetMessage) ProtoMessage()    {}
func (*AddPetRequest_PetMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_petstore_648dbf3c11f4d833, []int{0, 0}
}
func (m *AddPetRequest_PetMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddPetRequest_PetMessage.Unmarshal(m, b)
}
func (m *AddPetRequest_PetMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddPetRequest_PetMessage.Marshal(b, m, deterministic)
}
func (dst *AddPetRequest_PetMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddPetRequest_PetMessage.Merge(dst, src)
}
func (m *AddPetRequest_PetMessage) XXX_Size() int {
	return xxx_messageInfo_AddPetRequest_PetMessage.Size(m)
}
func (m *AddPetRequest_PetMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_AddPetRequest_PetMessage.DiscardUnknown(m)
}

var xxx_messageInfo_AddPetRequest_PetMessage proto.InternalMessageInfo

func (m *AddPetRequest_PetMessage) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *AddPetRequest_PetMessage) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *AddPetRequest_PetMessage) GetTag() string {
	if m != nil {
		return m.Tag
	}
	return ""
}

type AddPetResponse struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Tag                  string   `protobuf:"bytes,3,opt,name=tag,proto3" json:"tag,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddPetResponse) Reset()         { *m = AddPetResponse{} }
func (m *AddPetResponse) String() string { return proto.CompactTextString(m) }
func (*AddPetResponse) ProtoMessage()    {}
func (*AddPetResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_petstore_648dbf3c11f4d833, []int{1}
}
func (m *AddPetResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddPetResponse.Unmarshal(m, b)
}
func (m *AddPetResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddPetResponse.Marshal(b, m, deterministic)
}
func (dst *AddPetResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddPetResponse.Merge(dst, src)
}
func (m *AddPetResponse) XXX_Size() int {
	return xxx_messageInfo_AddPetResponse.Size(m)
}
func (m *AddPetResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AddPetResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AddPetResponse proto.InternalMessageInfo

func (m *AddPetResponse) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *AddPetResponse) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *AddPetResponse) GetTag() string {
	if m != nil {
		return m.Tag
	}
	return ""
}

type DeletePetRequest struct {
	// ID of pet to delete
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeletePetRequest) Reset()         { *m = DeletePetRequest{} }
func (m *DeletePetRequest) String() string { return proto.CompactTextString(m) }
func (*DeletePetRequest) ProtoMessage()    {}
func (*DeletePetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_petstore_648dbf3c11f4d833, []int{2}
}
func (m *DeletePetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeletePetRequest.Unmarshal(m, b)
}
func (m *DeletePetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeletePetRequest.Marshal(b, m, deterministic)
}
func (dst *DeletePetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeletePetRequest.Merge(dst, src)
}
func (m *DeletePetRequest) XXX_Size() int {
	return xxx_messageInfo_DeletePetRequest.Size(m)
}
func (m *DeletePetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeletePetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeletePetRequest proto.InternalMessageInfo

func (m *DeletePetRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type FindPetByIdRequest struct {
	// ID of pet to fetch
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FindPetByIdRequest) Reset()         { *m = FindPetByIdRequest{} }
func (m *FindPetByIdRequest) String() string { return proto.CompactTextString(m) }
func (*FindPetByIdRequest) ProtoMessage()    {}
func (*FindPetByIdRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_petstore_648dbf3c11f4d833, []int{3}
}
func (m *FindPetByIdRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindPetByIdRequest.Unmarshal(m, b)
}
func (m *FindPetByIdRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindPetByIdRequest.Marshal(b, m, deterministic)
}
func (dst *FindPetByIdRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindPetByIdRequest.Merge(dst, src)
}
func (m *FindPetByIdRequest) XXX_Size() int {
	return xxx_messageInfo_FindPetByIdRequest.Size(m)
}
func (m *FindPetByIdRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FindPetByIdRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FindPetByIdRequest proto.InternalMessageInfo

func (m *FindPetByIdRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type FindPetByIdResponse struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Tag                  string   `protobuf:"bytes,3,opt,name=tag,proto3" json:"tag,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FindPetByIdResponse) Reset()         { *m = FindPetByIdResponse{} }
func (m *FindPetByIdResponse) String() string { return proto.CompactTextString(m) }
func (*FindPetByIdResponse) ProtoMessage()    {}
func (*FindPetByIdResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_petstore_648dbf3c11f4d833, []int{4}
}
func (m *FindPetByIdResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindPetByIdResponse.Unmarshal(m, b)
}
func (m *FindPetByIdResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindPetByIdResponse.Marshal(b, m, deterministic)
}
func (dst *FindPetByIdResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindPetByIdResponse.Merge(dst, src)
}
func (m *FindPetByIdResponse) XXX_Size() int {
	return xxx_messageInfo_FindPetByIdResponse.Size(m)
}
func (m *FindPetByIdResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_FindPetByIdResponse.DiscardUnknown(m)
}

var xxx_messageInfo_FindPetByIdResponse proto.InternalMessageInfo

func (m *FindPetByIdResponse) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *FindPetByIdResponse) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *FindPetByIdResponse) GetTag() string {
	if m != nil {
		return m.Tag
	}
	return ""
}

type FindPetsByIdsRequest struct {
	Ids []string `protobuf:"bytes,1,rep,name=ids,proto3" json:"ids,omitempty"`
	// maximum number of results to return
	Limit                int32    `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FindPetsByIdsRequest) Reset()         { *m = FindPetsByIdsRequest{} }
func (m *FindPetsByIdsRequest) String() string { return proto.CompactTextString(m) }
func (*FindPetsByIdsRequest) ProtoMessage()    {}
func (*FindPetsByIdsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_petstore_648dbf3c11f4d833, []int{5}
}
func (m *FindPetsByIdsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindPetsByIdsRequest.Unmarshal(m, b)
}
func (m *FindPetsByIdsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindPetsByIdsRequest.Marshal(b, m, deterministic)
}
func (dst *FindPetsByIdsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindPetsByIdsRequest.Merge(dst, src)
}
func (m *FindPetsByIdsRequest) XXX_Size() int {
	return xxx_messageInfo_FindPetsByIdsRequest.Size(m)
}
func (m *FindPetsByIdsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FindPetsByIdsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FindPetsByIdsRequest proto.InternalMessageInfo

func (m *FindPetsByIdsRequest) GetIds() []string {
	if m != nil {
		return m.Ids
	}
	return nil
}

func (m *FindPetsByIdsRequest) GetLimit() int32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

type FindPetsByIdsResponse struct {
	Pets                 []*FindPetsByIdsResponse_PetsMessage `protobuf:"bytes,1,rep,name=pets,proto3" json:"pets,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                             `json:"-"`
	XXX_unrecognized     []byte                               `json:"-"`
	XXX_sizecache        int32                                `json:"-"`
}

func (m *FindPetsByIdsResponse) Reset()         { *m = FindPetsByIdsResponse{} }
func (m *FindPetsByIdsResponse) String() string { return proto.CompactTextString(m) }
func (*FindPetsByIdsResponse) ProtoMessage()    {}
func (*FindPetsByIdsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_petstore_648dbf3c11f4d833, []int{6}
}
func (m *FindPetsByIdsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindPetsByIdsResponse.Unmarshal(m, b)
}
func (m *FindPetsByIdsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindPetsByIdsResponse.Marshal(b, m, deterministic)
}
func (dst *FindPetsByIdsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindPetsByIdsResponse.Merge(dst, src)
}
func (m *FindPetsByIdsResponse) XXX_Size() int {
	return xxx_messageInfo_FindPetsByIdsResponse.Size(m)
}
func (m *FindPetsByIdsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_FindPetsByIdsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_FindPetsByIdsResponse proto.InternalMessageInfo

func (m *FindPetsByIdsResponse) GetPets() []*FindPetsByIdsResponse_PetsMessage {
	if m != nil {
		return m.Pets
	}
	return nil
}

type FindPetsByIdsResponse_PetsMessage struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Tag                  string   `protobuf:"bytes,3,opt,name=tag,proto3" json:"tag,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FindPetsByIdsResponse_PetsMessage) Reset()         { *m = FindPetsByIdsResponse_PetsMessage{} }
func (m *FindPetsByIdsResponse_PetsMessage) String() string { return proto.CompactTextString(m) }
func (*FindPetsByIdsResponse_PetsMessage) ProtoMessage()    {}
func (*FindPetsByIdsResponse_PetsMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_petstore_648dbf3c11f4d833, []int{6, 0}
}
func (m *FindPetsByIdsResponse_PetsMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindPetsByIdsResponse_PetsMessage.Unmarshal(m, b)
}
func (m *FindPetsByIdsResponse_PetsMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindPetsByIdsResponse_PetsMessage.Marshal(b, m, deterministic)
}
func (dst *FindPetsByIdsResponse_PetsMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindPetsByIdsResponse_PetsMessage.Merge(dst, src)
}
func (m *FindPetsByIdsResponse_PetsMessage) XXX_Size() int {
	return xxx_messageInfo_FindPetsByIdsResponse_PetsMessage.Size(m)
}
func (m *FindPetsByIdsResponse_PetsMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_FindPetsByIdsResponse_PetsMessage.DiscardUnknown(m)
}

var xxx_messageInfo_FindPetsByIdsResponse_PetsMessage proto.InternalMessageInfo

func (m *FindPetsByIdsResponse_PetsMessage) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *FindPetsByIdsResponse_PetsMessage) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *FindPetsByIdsResponse_PetsMessage) GetTag() string {
	if m != nil {
		return m.Tag
	}
	return ""
}

type FindPetsRequest struct {
	// maximum number of results to return
	Limit int32 `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	// tags to filter by
	Tags                 []string `protobuf:"bytes,2,rep,name=tags,proto3" json:"tags,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FindPetsRequest) Reset()         { *m = FindPetsRequest{} }
func (m *FindPetsRequest) String() string { return proto.CompactTextString(m) }
func (*FindPetsRequest) ProtoMessage()    {}
func (*FindPetsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_petstore_648dbf3c11f4d833, []int{7}
}
func (m *FindPetsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindPetsRequest.Unmarshal(m, b)
}
func (m *FindPetsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindPetsRequest.Marshal(b, m, deterministic)
}
func (dst *FindPetsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindPetsRequest.Merge(dst, src)
}
func (m *FindPetsRequest) XXX_Size() int {
	return xxx_messageInfo_FindPetsRequest.Size(m)
}
func (m *FindPetsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FindPetsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FindPetsRequest proto.InternalMessageInfo

func (m *FindPetsRequest) GetLimit() int32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *FindPetsRequest) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

type FindPetsResponse struct {
	Pets                 []*FindPetsResponse_PetsMessage `protobuf:"bytes,1,rep,name=pets,proto3" json:"pets,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                        `json:"-"`
	XXX_unrecognized     []byte                          `json:"-"`
	XXX_sizecache        int32                           `json:"-"`
}

func (m *FindPetsResponse) Reset()         { *m = FindPetsResponse{} }
func (m *FindPetsResponse) String() string { return proto.CompactTextString(m) }
func (*FindPetsResponse) ProtoMessage()    {}
func (*FindPetsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_petstore_648dbf3c11f4d833, []int{8}
}
func (m *FindPetsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindPetsResponse.Unmarshal(m, b)
}
func (m *FindPetsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindPetsResponse.Marshal(b, m, deterministic)
}
func (dst *FindPetsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindPetsResponse.Merge(dst, src)
}
func (m *FindPetsResponse) XXX_Size() int {
	return xxx_messageInfo_FindPetsResponse.Size(m)
}
func (m *FindPetsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_FindPetsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_FindPetsResponse proto.InternalMessageInfo

func (m *FindPetsResponse) GetPets() []*FindPetsResponse_PetsMessage {
	if m != nil {
		return m.Pets
	}
	return nil
}

type FindPetsResponse_PetsMessage struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Tag                  string   `protobuf:"bytes,3,opt,name=tag,proto3" json:"tag,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FindPetsResponse_PetsMessage) Reset()         { *m = FindPetsResponse_PetsMessage{} }
func (m *FindPetsResponse_PetsMessage) String() string { return proto.CompactTextString(m) }
func (*FindPetsResponse_PetsMessage) ProtoMessage()    {}
func (*FindPetsResponse_PetsMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_petstore_648dbf3c11f4d833, []int{8, 0}
}
func (m *FindPetsResponse_PetsMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindPetsResponse_PetsMessage.Unmarshal(m, b)
}
func (m *FindPetsResponse_PetsMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindPetsResponse_PetsMessage.Marshal(b, m, deterministic)
}
func (dst *FindPetsResponse_PetsMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindPetsResponse_PetsMessage.Merge(dst, src)
}
func (m *FindPetsResponse_PetsMessage) XXX_Size() int {
	return xxx_messageInfo_FindPetsResponse_PetsMessage.Size(m)
}
func (m *FindPetsResponse_PetsMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_FindPetsResponse_PetsMessage.DiscardUnknown(m)
}

var xxx_messageInfo_FindPetsResponse_PetsMessage proto.InternalMessageInfo

func (m *FindPetsResponse_PetsMessage) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *FindPetsResponse_PetsMessage) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *FindPetsResponse_PetsMessage) GetTag() string {
	if m != nil {
		return m.Tag
	}
	return ""
}

func init() {
	proto.RegisterType((*AddPetRequest)(nil), "am.is.AddPetRequest")
	proto.RegisterType((*AddPetRequest_PetMessage)(nil), "am.is.AddPetRequest.PetMessage")
	proto.RegisterType((*AddPetResponse)(nil), "am.is.AddPetResponse")
	proto.RegisterType((*DeletePetRequest)(nil), "am.is.DeletePetRequest")
	proto.RegisterType((*FindPetByIdRequest)(nil), "am.is.FindPetByIdRequest")
	proto.RegisterType((*FindPetByIdResponse)(nil), "am.is.FindPetByIdResponse")
	proto.RegisterType((*FindPetsByIdsRequest)(nil), "am.is.FindPetsByIdsRequest")
	proto.RegisterType((*FindPetsByIdsResponse)(nil), "am.is.FindPetsByIdsResponse")
	proto.RegisterType((*FindPetsByIdsResponse_PetsMessage)(nil), "am.is.FindPetsByIdsResponse.PetsMessage")
	proto.RegisterType((*FindPetsRequest)(nil), "am.is.FindPetsRequest")
	proto.RegisterType((*FindPetsResponse)(nil), "am.is.FindPetsResponse")
	proto.RegisterType((*FindPetsResponse_PetsMessage)(nil), "am.is.FindPetsResponse.PetsMessage")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// PetstoreServiceClient is the client API for PetstoreService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PetstoreServiceClient interface {
	// Creates a new pet in the store.  Duplicates are allowed
	AddPet(ctx context.Context, in *AddPetRequest, opts ...grpc.CallOption) (*AddPetResponse, error)
	// deletes a single pet based on the ID supplied
	DeletePet(ctx context.Context, in *DeletePetRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	// Returns a user based on a single ID, if the user does not have access to the pet
	FindPetById(ctx context.Context, in *FindPetByIdRequest, opts ...grpc.CallOption) (*FindPetByIdResponse, error)
	// Returns all pets from the system that the user has access to
	FindPets(ctx context.Context, in *FindPetsRequest, opts ...grpc.CallOption) (*FindPetsResponse, error)
	// Returns all pets from the system that the user has access to
	FindPetsByIds(ctx context.Context, in *FindPetsByIdsRequest, opts ...grpc.CallOption) (*FindPetsByIdsResponse, error)
}

type petstoreServiceClient struct {
	cc *grpc.ClientConn
}

func NewPetstoreServiceClient(cc *grpc.ClientConn) PetstoreServiceClient {
	return &petstoreServiceClient{cc}
}

func (c *petstoreServiceClient) AddPet(ctx context.Context, in *AddPetRequest, opts ...grpc.CallOption) (*AddPetResponse, error) {
	out := new(AddPetResponse)
	err := c.cc.Invoke(ctx, "/am.is.PetstoreService/AddPet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *petstoreServiceClient) DeletePet(ctx context.Context, in *DeletePetRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/am.is.PetstoreService/DeletePet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *petstoreServiceClient) FindPetById(ctx context.Context, in *FindPetByIdRequest, opts ...grpc.CallOption) (*FindPetByIdResponse, error) {
	out := new(FindPetByIdResponse)
	err := c.cc.Invoke(ctx, "/am.is.PetstoreService/FindPetById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *petstoreServiceClient) FindPets(ctx context.Context, in *FindPetsRequest, opts ...grpc.CallOption) (*FindPetsResponse, error) {
	out := new(FindPetsResponse)
	err := c.cc.Invoke(ctx, "/am.is.PetstoreService/FindPets", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *petstoreServiceClient) FindPetsByIds(ctx context.Context, in *FindPetsByIdsRequest, opts ...grpc.CallOption) (*FindPetsByIdsResponse, error) {
	out := new(FindPetsByIdsResponse)
	err := c.cc.Invoke(ctx, "/am.is.PetstoreService/FindPetsByIds", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PetstoreServiceServer is the server API for PetstoreService service.
type PetstoreServiceServer interface {
	// Creates a new pet in the store.  Duplicates are allowed
	AddPet(context.Context, *AddPetRequest) (*AddPetResponse, error)
	// deletes a single pet based on the ID supplied
	DeletePet(context.Context, *DeletePetRequest) (*empty.Empty, error)
	// Returns a user based on a single ID, if the user does not have access to the pet
	FindPetById(context.Context, *FindPetByIdRequest) (*FindPetByIdResponse, error)
	// Returns all pets from the system that the user has access to
	FindPets(context.Context, *FindPetsRequest) (*FindPetsResponse, error)
	// Returns all pets from the system that the user has access to
	FindPetsByIds(context.Context, *FindPetsByIdsRequest) (*FindPetsByIdsResponse, error)
}

func RegisterPetstoreServiceServer(s *grpc.Server, srv PetstoreServiceServer) {
	s.RegisterService(&_PetstoreService_serviceDesc, srv)
}

func _PetstoreService_AddPet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddPetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PetstoreServiceServer).AddPet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/am.is.PetstoreService/AddPet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PetstoreServiceServer).AddPet(ctx, req.(*AddPetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PetstoreService_DeletePet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeletePetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PetstoreServiceServer).DeletePet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/am.is.PetstoreService/DeletePet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PetstoreServiceServer).DeletePet(ctx, req.(*DeletePetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PetstoreService_FindPetById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindPetByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PetstoreServiceServer).FindPetById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/am.is.PetstoreService/FindPetById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PetstoreServiceServer).FindPetById(ctx, req.(*FindPetByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PetstoreService_FindPets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindPetsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PetstoreServiceServer).FindPets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/am.is.PetstoreService/FindPets",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PetstoreServiceServer).FindPets(ctx, req.(*FindPetsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PetstoreService_FindPetsByIds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindPetsByIdsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PetstoreServiceServer).FindPetsByIds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/am.is.PetstoreService/FindPetsByIds",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PetstoreServiceServer).FindPetsByIds(ctx, req.(*FindPetsByIdsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _PetstoreService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "am.is.PetstoreService",
	HandlerType: (*PetstoreServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddPet",
			Handler:    _PetstoreService_AddPet_Handler,
		},
		{
			MethodName: "DeletePet",
			Handler:    _PetstoreService_DeletePet_Handler,
		},
		{
			MethodName: "FindPetById",
			Handler:    _PetstoreService_FindPetById_Handler,
		},
		{
			MethodName: "FindPets",
			Handler:    _PetstoreService_FindPets_Handler,
		},
		{
			MethodName: "FindPetsByIds",
			Handler:    _PetstoreService_FindPetsByIds_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/alanchchen/go-example-twirp/api/petstore.proto",
}

func init() {
	proto.RegisterFile("github.com/alanchchen/go-example-twirp/api/petstore.proto", fileDescriptor_petstore_648dbf3c11f4d833)
}

var fileDescriptor_petstore_648dbf3c11f4d833 = []byte{
	// 554 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x53, 0x4f, 0x8b, 0x13, 0x4f,
	0x10, 0x65, 0x32, 0x9b, 0xe5, 0x97, 0x0a, 0x9b, 0x64, 0xeb, 0x97, 0x4d, 0xe2, 0xec, 0x82, 0x61,
	0xf4, 0x10, 0x84, 0x9d, 0xc1, 0x78, 0x10, 0xff, 0x20, 0x18, 0x75, 0x41, 0x64, 0x25, 0x44, 0x4f,
	0xe2, 0xa5, 0x93, 0x29, 0x27, 0x0d, 0x99, 0x3f, 0xa6, 0x3b, 0xea, 0x22, 0x7b, 0xf1, 0xe0, 0x59,
	0xf0, 0xe6, 0x37, 0xf2, 0xec, 0x57, 0xf0, 0x83, 0x48, 0xf7, 0xcc, 0x24, 0x33, 0x63, 0xd6, 0xc3,
	0x8a, 0xb7, 0xee, 0xea, 0x57, 0xaf, 0xea, 0x75, 0xbd, 0x82, 0x3b, 0x3e, 0x97, 0xf3, 0xd5, 0xd4,
	0x99, 0x45, 0x81, 0xcb, 0x16, 0x2c, 0x9c, 0xcd, 0x67, 0x73, 0x0a, 0x5d, 0x3f, 0x3a, 0xa6, 0x0f,
	0x2c, 0x88, 0x17, 0x74, 0x2c, 0xdf, 0xf3, 0x65, 0xec, 0xb2, 0x98, 0xbb, 0x31, 0x49, 0x21, 0xa3,
	0x25, 0x39, 0xf1, 0x32, 0x92, 0x11, 0x56, 0x59, 0xe0, 0x70, 0x61, 0x1d, 0xf9, 0x51, 0xe4, 0x2f,
	0x48, 0x23, 0x58, 0x18, 0x46, 0x92, 0x49, 0x1e, 0x85, 0x22, 0x01, 0x59, 0x87, 0xe9, 0xab, 0xbe,
	0x4d, 0x57, 0x6f, 0x5c, 0x0a, 0x62, 0x79, 0x96, 0x3c, 0xda, 0x9f, 0x0d, 0xd8, 0x7b, 0xe8, 0x79,
	0x63, 0x92, 0x13, 0x7a, 0xbb, 0x22, 0x21, 0xf1, 0x26, 0x98, 0x31, 0xc9, 0x9e, 0xd1, 0x37, 0x06,
	0xf5, 0xe1, 0x55, 0x47, 0x57, 0x70, 0x0a, 0x10, 0x67, 0x4c, 0xf2, 0x94, 0x84, 0x60, 0x3e, 0x4d,
	0x14, 0xd6, 0x1a, 0x01, 0x6c, 0x42, 0xd8, 0x80, 0x0a, 0xf7, 0x74, 0xbe, 0x39, 0xa9, 0x70, 0x0f,
	0x11, 0x76, 0x42, 0x16, 0x50, 0xaf, 0xd2, 0x37, 0x06, 0xb5, 0x89, 0x3e, 0x63, 0x0b, 0x4c, 0xc9,
	0xfc, 0x9e, 0xa9, 0x43, 0xea, 0x68, 0x9f, 0x40, 0x23, 0x2b, 0x22, 0xe2, 0x28, 0x14, 0x97, 0xe5,
	0xb1, 0xa1, 0xf5, 0x98, 0x16, 0x24, 0x29, 0x27, 0xa9, 0xc4, 0x64, 0x5f, 0x07, 0x3c, 0xe1, 0xa1,
	0x2a, 0x36, 0x3a, 0x7b, 0xea, 0x5d, 0x84, 0x7a, 0x06, 0xff, 0x17, 0x50, 0x7f, 0xd5, 0xd6, 0x03,
	0x68, 0xa7, 0x64, 0x42, 0xb1, 0x89, 0xac, 0x68, 0x0b, 0x4c, 0xee, 0x89, 0x9e, 0xd1, 0x37, 0x15,
	0x92, 0x7b, 0x02, 0xdb, 0x50, 0x5d, 0xf0, 0x80, 0x4b, 0x4d, 0x58, 0x9d, 0x24, 0x17, 0xfb, 0x9b,
	0x01, 0x07, 0x25, 0x82, 0xb4, 0x9f, 0xfb, 0xb0, 0xa3, 0x5c, 0xa1, 0x29, 0xea, 0xc3, 0x41, 0x3a,
	0xb0, 0xad, 0x58, 0x35, 0x38, 0x91, 0x4d, 0x4e, 0x67, 0x59, 0x8f, 0xa0, 0x9e, 0x0b, 0x5e, 0x52,
	0xdc, 0x3d, 0x68, 0x66, 0xf5, 0x32, 0x5d, 0x6b, 0x15, 0x46, 0x4e, 0x85, 0xa2, 0x93, 0xcc, 0x17,
	0xbd, 0x8a, 0x96, 0xab, 0xcf, 0xf6, 0x17, 0x03, 0x5a, 0x9b, 0xec, 0x54, 0xd4, 0xed, 0x82, 0xa8,
	0x6b, 0x25, 0x51, 0xff, 0x58, 0xcf, 0xf0, 0xbb, 0x09, 0xcd, 0x71, 0xba, 0x69, 0x2f, 0x68, 0xf9,
	0x8e, 0xcf, 0x08, 0x4f, 0x61, 0x37, 0xf1, 0x27, 0xb6, 0xb7, 0xed, 0x84, 0x75, 0x50, 0x8a, 0x26,
	0x1d, 0xda, 0x9d, 0x4f, 0x3f, 0x7e, 0x7e, 0xad, 0xb4, 0xec, 0xda, 0x7a, 0x7d, 0xef, 0xaa, 0x95,
	0xc1, 0x97, 0x50, 0x5b, 0xdb, 0x14, 0xbb, 0x69, 0x6e, 0xd9, 0xb8, 0x56, 0xc7, 0x49, 0x76, 0xd7,
	0xc9, 0x76, 0xd7, 0x79, 0xa2, 0x76, 0x37, 0x63, 0xbd, 0xd1, 0x58, 0xb3, 0xba, 0x1f, 0xb9, 0x77,
	0x8e, 0xaf, 0xa1, 0x9e, 0xb3, 0x2c, 0x5e, 0x29, 0xfe, 0x5b, 0xce, 0xec, 0x96, 0xb5, 0xed, 0xa9,
	0xd8, 0x33, 0x96, 0xd9, 0x9f, 0xc3, 0x7f, 0xd9, 0x04, 0xb0, 0xf3, 0xdb, 0x48, 0x12, 0xde, 0xee,
	0x05, 0xa3, 0xb2, 0xf7, 0x35, 0x69, 0x1d, 0x37, 0x1f, 0x81, 0x33, 0xd8, 0x2b, 0xd8, 0x14, 0x0f,
	0xb7, 0x9b, 0x37, 0x61, 0x3e, 0xfa, 0x93, 0xb3, 0xed, 0xae, 0xa6, 0xdf, 0xc7, 0x66, 0xa1, 0x67,
	0x71, 0x3e, 0xaa, 0xbe, 0x32, 0x59, 0xcc, 0xa7, 0xbb, 0xfa, 0x07, 0x6f, 0xfd, 0x0a, 0x00, 0x00,
	0xff, 0xff, 0xe3, 0x9f, 0x33, 0x4e, 0x6d, 0x05, 0x00, 0x00,
}
