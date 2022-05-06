// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.20.0
// source: team.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type TeamResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatusCode int32    `protobuf:"varint,1,opt,name=StatusCode,proto3" json:"StatusCode,omitempty"`
	Errors     []string `protobuf:"bytes,2,rep,name=errors,proto3" json:"errors,omitempty"`
	Data       *Team    `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *TeamResponse) Reset() {
	*x = TeamResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_team_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TeamResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TeamResponse) ProtoMessage() {}

func (x *TeamResponse) ProtoReflect() protoreflect.Message {
	mi := &file_team_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TeamResponse.ProtoReflect.Descriptor instead.
func (*TeamResponse) Descriptor() ([]byte, []int) {
	return file_team_proto_rawDescGZIP(), []int{0}
}

func (x *TeamResponse) GetStatusCode() int32 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *TeamResponse) GetErrors() []string {
	if x != nil {
		return x.Errors
	}
	return nil
}

func (x *TeamResponse) GetData() *Team {
	if x != nil {
		return x.Data
	}
	return nil
}

type TeamPagination struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*Team             `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
	Meta  *PaginationMetadata `protobuf:"bytes,2,opt,name=meta,proto3" json:"meta,omitempty"`
}

func (x *TeamPagination) Reset() {
	*x = TeamPagination{}
	if protoimpl.UnsafeEnabled {
		mi := &file_team_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TeamPagination) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TeamPagination) ProtoMessage() {}

func (x *TeamPagination) ProtoReflect() protoreflect.Message {
	mi := &file_team_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TeamPagination.ProtoReflect.Descriptor instead.
func (*TeamPagination) Descriptor() ([]byte, []int) {
	return file_team_proto_rawDescGZIP(), []int{1}
}

func (x *TeamPagination) GetItems() []*Team {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *TeamPagination) GetMeta() *PaginationMetadata {
	if x != nil {
		return x.Meta
	}
	return nil
}

type TeamListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatusCode int32    `protobuf:"varint,1,opt,name=StatusCode,proto3" json:"StatusCode,omitempty"`
	Errors     []string `protobuf:"bytes,2,rep,name=errors,proto3" json:"errors,omitempty"`
	Data       []*Team  `protobuf:"bytes,3,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *TeamListResponse) Reset() {
	*x = TeamListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_team_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TeamListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TeamListResponse) ProtoMessage() {}

func (x *TeamListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_team_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TeamListResponse.ProtoReflect.Descriptor instead.
func (*TeamListResponse) Descriptor() ([]byte, []int) {
	return file_team_proto_rawDescGZIP(), []int{2}
}

func (x *TeamListResponse) GetStatusCode() int32 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *TeamListResponse) GetErrors() []string {
	if x != nil {
		return x.Errors
	}
	return nil
}

func (x *TeamListResponse) GetData() []*Team {
	if x != nil {
		return x.Data
	}
	return nil
}

type TeamPaginationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatusCode int32           `protobuf:"varint,1,opt,name=StatusCode,proto3" json:"StatusCode,omitempty"`
	Errors     []string        `protobuf:"bytes,2,rep,name=errors,proto3" json:"errors,omitempty"`
	Data       *TeamPagination `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *TeamPaginationResponse) Reset() {
	*x = TeamPaginationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_team_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TeamPaginationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TeamPaginationResponse) ProtoMessage() {}

func (x *TeamPaginationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_team_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TeamPaginationResponse.ProtoReflect.Descriptor instead.
func (*TeamPaginationResponse) Descriptor() ([]byte, []int) {
	return file_team_proto_rawDescGZIP(), []int{3}
}

func (x *TeamPaginationResponse) GetStatusCode() int32 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *TeamPaginationResponse) GetErrors() []string {
	if x != nil {
		return x.Errors
	}
	return nil
}

func (x *TeamPaginationResponse) GetData() *TeamPagination {
	if x != nil {
		return x.Data
	}
	return nil
}

type FindAllTeamRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Limit int64 `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	Page  int64 `protobuf:"varint,2,opt,name=page,proto3" json:"page,omitempty"`
}

func (x *FindAllTeamRequest) Reset() {
	*x = FindAllTeamRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_team_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindAllTeamRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindAllTeamRequest) ProtoMessage() {}

func (x *FindAllTeamRequest) ProtoReflect() protoreflect.Message {
	mi := &file_team_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindAllTeamRequest.ProtoReflect.Descriptor instead.
func (*FindAllTeamRequest) Descriptor() ([]byte, []int) {
	return file_team_proto_rawDescGZIP(), []int{4}
}

func (x *FindAllTeamRequest) GetLimit() int64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *FindAllTeamRequest) GetPage() int64 {
	if x != nil {
		return x.Page
	}
	return 0
}

type FindOneTeamRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *FindOneTeamRequest) Reset() {
	*x = FindOneTeamRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_team_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindOneTeamRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindOneTeamRequest) ProtoMessage() {}

func (x *FindOneTeamRequest) ProtoReflect() protoreflect.Message {
	mi := &file_team_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindOneTeamRequest.ProtoReflect.Descriptor instead.
func (*FindOneTeamRequest) Descriptor() ([]byte, []int) {
	return file_team_proto_rawDescGZIP(), []int{5}
}

func (x *FindOneTeamRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type FindMultiTeamRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ids []uint32 `protobuf:"varint,1,rep,packed,name=ids,proto3" json:"ids,omitempty"`
}

func (x *FindMultiTeamRequest) Reset() {
	*x = FindMultiTeamRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_team_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindMultiTeamRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindMultiTeamRequest) ProtoMessage() {}

func (x *FindMultiTeamRequest) ProtoReflect() protoreflect.Message {
	mi := &file_team_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindMultiTeamRequest.ProtoReflect.Descriptor instead.
func (*FindMultiTeamRequest) Descriptor() ([]byte, []int) {
	return file_team_proto_rawDescGZIP(), []int{6}
}

func (x *FindMultiTeamRequest) GetIds() []uint32 {
	if x != nil {
		return x.Ids
	}
	return nil
}

type CreateTeamRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Team *Team `protobuf:"bytes,1,opt,name=team,proto3" json:"team,omitempty"`
}

func (x *CreateTeamRequest) Reset() {
	*x = CreateTeamRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_team_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTeamRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTeamRequest) ProtoMessage() {}

func (x *CreateTeamRequest) ProtoReflect() protoreflect.Message {
	mi := &file_team_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTeamRequest.ProtoReflect.Descriptor instead.
func (*CreateTeamRequest) Descriptor() ([]byte, []int) {
	return file_team_proto_rawDescGZIP(), []int{7}
}

func (x *CreateTeamRequest) GetTeam() *Team {
	if x != nil {
		return x.Team
	}
	return nil
}

type UpdateTeamRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Team *Team `protobuf:"bytes,1,opt,name=team,proto3" json:"team,omitempty"`
}

func (x *UpdateTeamRequest) Reset() {
	*x = UpdateTeamRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_team_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateTeamRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateTeamRequest) ProtoMessage() {}

func (x *UpdateTeamRequest) ProtoReflect() protoreflect.Message {
	mi := &file_team_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateTeamRequest.ProtoReflect.Descriptor instead.
func (*UpdateTeamRequest) Descriptor() ([]byte, []int) {
	return file_team_proto_rawDescGZIP(), []int{8}
}

func (x *UpdateTeamRequest) GetTeam() *Team {
	if x != nil {
		return x.Team
	}
	return nil
}

type DeleteTeamRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteTeamRequest) Reset() {
	*x = DeleteTeamRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_team_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteTeamRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteTeamRequest) ProtoMessage() {}

func (x *DeleteTeamRequest) ProtoReflect() protoreflect.Message {
	mi := &file_team_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteTeamRequest.ProtoReflect.Descriptor instead.
func (*DeleteTeamRequest) Descriptor() ([]byte, []int) {
	return file_team_proto_rawDescGZIP(), []int{9}
}

func (x *DeleteTeamRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

var File_team_proto protoreflect.FileDescriptor

var file_team_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x74, 0x65, 0x61, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x74, 0x65,
	0x61, 0x6d, 0x1a, 0x09, 0x64, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x65, 0x0a, 0x0c, 0x54,
	0x65, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0a, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x73, 0x12, 0x1d, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x09, 0x2e, 0x64, 0x74, 0x6f, 0x2e, 0x54, 0x65, 0x61, 0x6d, 0x52, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x22, 0x61, 0x0a, 0x0e, 0x54, 0x65, 0x61, 0x6d, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1f, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x64, 0x74, 0x6f, 0x2e, 0x54, 0x65, 0x61, 0x6d, 0x52, 0x05,
	0x69, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x2e, 0x0a, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x50, 0x61, 0x67,
	0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52,
	0x04, 0x6d, 0x65, 0x74, 0x61, 0x22, 0x69, 0x0a, 0x10, 0x54, 0x65, 0x61, 0x6d, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x73, 0x12, 0x1d, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x09, 0x2e, 0x64, 0x74, 0x6f, 0x2e, 0x54, 0x65, 0x61, 0x6d, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x22, 0x7a, 0x0a, 0x16, 0x54, 0x65, 0x61, 0x6d, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x73, 0x12, 0x28, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x14, 0x2e, 0x74, 0x65, 0x61, 0x6d, 0x2e, 0x54, 0x65, 0x61, 0x6d, 0x50, 0x61, 0x67, 0x69,
	0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x3e, 0x0a, 0x12,
	0x46, 0x69, 0x6e, 0x64, 0x41, 0x6c, 0x6c, 0x54, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x22, 0x24, 0x0a, 0x12,
	0x46, 0x69, 0x6e, 0x64, 0x4f, 0x6e, 0x65, 0x54, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02,
	0x69, 0x64, 0x22, 0x28, 0x0a, 0x14, 0x46, 0x69, 0x6e, 0x64, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x54,
	0x65, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x64,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x03, 0x69, 0x64, 0x73, 0x22, 0x32, 0x0a, 0x11,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1d, 0x0a, 0x04, 0x74, 0x65, 0x61, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x09, 0x2e, 0x64, 0x74, 0x6f, 0x2e, 0x54, 0x65, 0x61, 0x6d, 0x52, 0x04, 0x74, 0x65, 0x61, 0x6d,
	0x22, 0x32, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x65, 0x61, 0x6d, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x04, 0x74, 0x65, 0x61, 0x6d, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x64, 0x74, 0x6f, 0x2e, 0x54, 0x65, 0x61, 0x6d, 0x52, 0x04,
	0x74, 0x65, 0x61, 0x6d, 0x22, 0x23, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x65,
	0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x32, 0xfb, 0x02, 0x0a, 0x0b, 0x54, 0x65,
	0x61, 0x6d, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x43, 0x0a, 0x07, 0x46, 0x69, 0x6e,
	0x64, 0x41, 0x6c, 0x6c, 0x12, 0x18, 0x2e, 0x74, 0x65, 0x61, 0x6d, 0x2e, 0x46, 0x69, 0x6e, 0x64,
	0x41, 0x6c, 0x6c, 0x54, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c,
	0x2e, 0x74, 0x65, 0x61, 0x6d, 0x2e, 0x54, 0x65, 0x61, 0x6d, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x39,
	0x0a, 0x07, 0x46, 0x69, 0x6e, 0x64, 0x4f, 0x6e, 0x65, 0x12, 0x18, 0x2e, 0x74, 0x65, 0x61, 0x6d,
	0x2e, 0x46, 0x69, 0x6e, 0x64, 0x4f, 0x6e, 0x65, 0x54, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x74, 0x65, 0x61, 0x6d, 0x2e, 0x54, 0x65, 0x61, 0x6d, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x41, 0x0a, 0x09, 0x46, 0x69, 0x6e,
	0x64, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x12, 0x1a, 0x2e, 0x74, 0x65, 0x61, 0x6d, 0x2e, 0x46, 0x69,
	0x6e, 0x64, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x54, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x16, 0x2e, 0x74, 0x65, 0x61, 0x6d, 0x2e, 0x54, 0x65, 0x61, 0x6d, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x37, 0x0a, 0x06,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x17, 0x2e, 0x74, 0x65, 0x61, 0x6d, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x54, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x12, 0x2e, 0x74, 0x65, 0x61, 0x6d, 0x2e, 0x54, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x37, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12,
	0x17, 0x2e, 0x74, 0x65, 0x61, 0x6d, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x65, 0x61,
	0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x74, 0x65, 0x61, 0x6d, 0x2e,
	0x54, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x37,
	0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x17, 0x2e, 0x74, 0x65, 0x61, 0x6d, 0x2e,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x12, 0x2e, 0x74, 0x65, 0x61, 0x6d, 0x2e, 0x54, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x39, 0x5a, 0x37, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x61, 0x6d, 0x69, 0x74, 0x68, 0x69, 0x77, 0x61, 0x74,
	0x2f, 0x73, 0x61, 0x6d, 0x69, 0x74, 0x68, 0x69, 0x77, 0x61, 0x74, 0x2d, 0x62, 0x61, 0x63, 0x6b,
	0x65, 0x6e, 0x64, 0x2d, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x73, 0x72, 0x63, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_team_proto_rawDescOnce sync.Once
	file_team_proto_rawDescData = file_team_proto_rawDesc
)

func file_team_proto_rawDescGZIP() []byte {
	file_team_proto_rawDescOnce.Do(func() {
		file_team_proto_rawDescData = protoimpl.X.CompressGZIP(file_team_proto_rawDescData)
	})
	return file_team_proto_rawDescData
}

var file_team_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_team_proto_goTypes = []interface{}{
	(*TeamResponse)(nil),           // 0: team.TeamResponse
	(*TeamPagination)(nil),         // 1: team.TeamPagination
	(*TeamListResponse)(nil),       // 2: team.TeamListResponse
	(*TeamPaginationResponse)(nil), // 3: team.TeamPaginationResponse
	(*FindAllTeamRequest)(nil),     // 4: team.FindAllTeamRequest
	(*FindOneTeamRequest)(nil),     // 5: team.FindOneTeamRequest
	(*FindMultiTeamRequest)(nil),   // 6: team.FindMultiTeamRequest
	(*CreateTeamRequest)(nil),      // 7: team.CreateTeamRequest
	(*UpdateTeamRequest)(nil),      // 8: team.UpdateTeamRequest
	(*DeleteTeamRequest)(nil),      // 9: team.DeleteTeamRequest
	(*Team)(nil),                   // 10: dto.Team
	(*PaginationMetadata)(nil),     // 11: common.PaginationMetadata
}
var file_team_proto_depIdxs = []int32{
	10, // 0: team.TeamResponse.data:type_name -> dto.Team
	10, // 1: team.TeamPagination.items:type_name -> dto.Team
	11, // 2: team.TeamPagination.meta:type_name -> common.PaginationMetadata
	10, // 3: team.TeamListResponse.data:type_name -> dto.Team
	1,  // 4: team.TeamPaginationResponse.data:type_name -> team.TeamPagination
	10, // 5: team.CreateTeamRequest.team:type_name -> dto.Team
	10, // 6: team.UpdateTeamRequest.team:type_name -> dto.Team
	4,  // 7: team.TeamService.FindAll:input_type -> team.FindAllTeamRequest
	5,  // 8: team.TeamService.FindOne:input_type -> team.FindOneTeamRequest
	6,  // 9: team.TeamService.FindMulti:input_type -> team.FindMultiTeamRequest
	7,  // 10: team.TeamService.Create:input_type -> team.CreateTeamRequest
	8,  // 11: team.TeamService.Update:input_type -> team.UpdateTeamRequest
	9,  // 12: team.TeamService.Delete:input_type -> team.DeleteTeamRequest
	3,  // 13: team.TeamService.FindAll:output_type -> team.TeamPaginationResponse
	0,  // 14: team.TeamService.FindOne:output_type -> team.TeamResponse
	2,  // 15: team.TeamService.FindMulti:output_type -> team.TeamListResponse
	0,  // 16: team.TeamService.Create:output_type -> team.TeamResponse
	0,  // 17: team.TeamService.Update:output_type -> team.TeamResponse
	0,  // 18: team.TeamService.Delete:output_type -> team.TeamResponse
	13, // [13:19] is the sub-list for method output_type
	7,  // [7:13] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_team_proto_init() }
func file_team_proto_init() {
	if File_team_proto != nil {
		return
	}
	file_dto_proto_init()
	file_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_team_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TeamResponse); i {
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
		file_team_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TeamPagination); i {
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
		file_team_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TeamListResponse); i {
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
		file_team_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TeamPaginationResponse); i {
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
		file_team_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindAllTeamRequest); i {
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
		file_team_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindOneTeamRequest); i {
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
		file_team_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindMultiTeamRequest); i {
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
		file_team_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateTeamRequest); i {
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
		file_team_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateTeamRequest); i {
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
		file_team_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteTeamRequest); i {
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
			RawDescriptor: file_team_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_team_proto_goTypes,
		DependencyIndexes: file_team_proto_depIdxs,
		MessageInfos:      file_team_proto_msgTypes,
	}.Build()
	File_team_proto = out.File
	file_team_proto_rawDesc = nil
	file_team_proto_goTypes = nil
	file_team_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// TeamServiceClient is the client API for TeamService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TeamServiceClient interface {
	FindAll(ctx context.Context, in *FindAllTeamRequest, opts ...grpc.CallOption) (*TeamPaginationResponse, error)
	FindOne(ctx context.Context, in *FindOneTeamRequest, opts ...grpc.CallOption) (*TeamResponse, error)
	FindMulti(ctx context.Context, in *FindMultiTeamRequest, opts ...grpc.CallOption) (*TeamListResponse, error)
	Create(ctx context.Context, in *CreateTeamRequest, opts ...grpc.CallOption) (*TeamResponse, error)
	Update(ctx context.Context, in *UpdateTeamRequest, opts ...grpc.CallOption) (*TeamResponse, error)
	Delete(ctx context.Context, in *DeleteTeamRequest, opts ...grpc.CallOption) (*TeamResponse, error)
}

type teamServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTeamServiceClient(cc grpc.ClientConnInterface) TeamServiceClient {
	return &teamServiceClient{cc}
}

func (c *teamServiceClient) FindAll(ctx context.Context, in *FindAllTeamRequest, opts ...grpc.CallOption) (*TeamPaginationResponse, error) {
	out := new(TeamPaginationResponse)
	err := c.cc.Invoke(ctx, "/team.TeamService/FindAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *teamServiceClient) FindOne(ctx context.Context, in *FindOneTeamRequest, opts ...grpc.CallOption) (*TeamResponse, error) {
	out := new(TeamResponse)
	err := c.cc.Invoke(ctx, "/team.TeamService/FindOne", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *teamServiceClient) FindMulti(ctx context.Context, in *FindMultiTeamRequest, opts ...grpc.CallOption) (*TeamListResponse, error) {
	out := new(TeamListResponse)
	err := c.cc.Invoke(ctx, "/team.TeamService/FindMulti", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *teamServiceClient) Create(ctx context.Context, in *CreateTeamRequest, opts ...grpc.CallOption) (*TeamResponse, error) {
	out := new(TeamResponse)
	err := c.cc.Invoke(ctx, "/team.TeamService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *teamServiceClient) Update(ctx context.Context, in *UpdateTeamRequest, opts ...grpc.CallOption) (*TeamResponse, error) {
	out := new(TeamResponse)
	err := c.cc.Invoke(ctx, "/team.TeamService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *teamServiceClient) Delete(ctx context.Context, in *DeleteTeamRequest, opts ...grpc.CallOption) (*TeamResponse, error) {
	out := new(TeamResponse)
	err := c.cc.Invoke(ctx, "/team.TeamService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TeamServiceServer is the server API for TeamService service.
type TeamServiceServer interface {
	FindAll(context.Context, *FindAllTeamRequest) (*TeamPaginationResponse, error)
	FindOne(context.Context, *FindOneTeamRequest) (*TeamResponse, error)
	FindMulti(context.Context, *FindMultiTeamRequest) (*TeamListResponse, error)
	Create(context.Context, *CreateTeamRequest) (*TeamResponse, error)
	Update(context.Context, *UpdateTeamRequest) (*TeamResponse, error)
	Delete(context.Context, *DeleteTeamRequest) (*TeamResponse, error)
}

// UnimplementedTeamServiceServer can be embedded to have forward compatible implementations.
type UnimplementedTeamServiceServer struct {
}

func (*UnimplementedTeamServiceServer) FindAll(context.Context, *FindAllTeamRequest) (*TeamPaginationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindAll not implemented")
}
func (*UnimplementedTeamServiceServer) FindOne(context.Context, *FindOneTeamRequest) (*TeamResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindOne not implemented")
}
func (*UnimplementedTeamServiceServer) FindMulti(context.Context, *FindMultiTeamRequest) (*TeamListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindMulti not implemented")
}
func (*UnimplementedTeamServiceServer) Create(context.Context, *CreateTeamRequest) (*TeamResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (*UnimplementedTeamServiceServer) Update(context.Context, *UpdateTeamRequest) (*TeamResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (*UnimplementedTeamServiceServer) Delete(context.Context, *DeleteTeamRequest) (*TeamResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}

func RegisterTeamServiceServer(s *grpc.Server, srv TeamServiceServer) {
	s.RegisterService(&_TeamService_serviceDesc, srv)
}

func _TeamService_FindAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindAllTeamRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TeamServiceServer).FindAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/team.TeamService/FindAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TeamServiceServer).FindAll(ctx, req.(*FindAllTeamRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TeamService_FindOne_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindOneTeamRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TeamServiceServer).FindOne(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/team.TeamService/FindOne",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TeamServiceServer).FindOne(ctx, req.(*FindOneTeamRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TeamService_FindMulti_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindMultiTeamRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TeamServiceServer).FindMulti(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/team.TeamService/FindMulti",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TeamServiceServer).FindMulti(ctx, req.(*FindMultiTeamRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TeamService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTeamRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TeamServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/team.TeamService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TeamServiceServer).Create(ctx, req.(*CreateTeamRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TeamService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateTeamRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TeamServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/team.TeamService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TeamServiceServer).Update(ctx, req.(*UpdateTeamRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TeamService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteTeamRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TeamServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/team.TeamService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TeamServiceServer).Delete(ctx, req.(*DeleteTeamRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _TeamService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "team.TeamService",
	HandlerType: (*TeamServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FindAll",
			Handler:    _TeamService_FindAll_Handler,
		},
		{
			MethodName: "FindOne",
			Handler:    _TeamService_FindOne_Handler,
		},
		{
			MethodName: "FindMulti",
			Handler:    _TeamService_FindMulti_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _TeamService_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _TeamService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _TeamService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "team.proto",
}