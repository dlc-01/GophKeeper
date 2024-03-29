// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v4.25.3
// source: text.proto

package gen

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

type NoteMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Note     string `protobuf:"bytes,2,opt,name=note,proto3" json:"note,omitempty"`
	Metadata string `protobuf:"bytes,3,opt,name=metadata,proto3" json:"metadata,omitempty"`
}

func (x *NoteMsg) Reset() {
	*x = NoteMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_text_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NoteMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NoteMsg) ProtoMessage() {}

func (x *NoteMsg) ProtoReflect() protoreflect.Message {
	mi := &file_text_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NoteMsg.ProtoReflect.Descriptor instead.
func (*NoteMsg) Descriptor() ([]byte, []int) {
	return file_text_proto_rawDescGZIP(), []int{0}
}

func (x *NoteMsg) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *NoteMsg) GetNote() string {
	if x != nil {
		return x.Note
	}
	return ""
}

func (x *NoteMsg) GetMetadata() string {
	if x != nil {
		return x.Metadata
	}
	return ""
}

type CreateTextRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Note  *NoteMsg `protobuf:"bytes,2,opt,name=note,proto3" json:"note,omitempty"`
}

func (x *CreateTextRequest) Reset() {
	*x = CreateTextRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_text_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTextRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTextRequest) ProtoMessage() {}

func (x *CreateTextRequest) ProtoReflect() protoreflect.Message {
	mi := &file_text_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTextRequest.ProtoReflect.Descriptor instead.
func (*CreateTextRequest) Descriptor() ([]byte, []int) {
	return file_text_proto_rawDescGZIP(), []int{1}
}

func (x *CreateTextRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *CreateTextRequest) GetNote() *NoteMsg {
	if x != nil {
		return x.Note
	}
	return nil
}

type CreateTextResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Note  *NoteMsg `protobuf:"bytes,1,opt,name=note,proto3" json:"note,omitempty"`
	Error string   `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *CreateTextResponse) Reset() {
	*x = CreateTextResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_text_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTextResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTextResponse) ProtoMessage() {}

func (x *CreateTextResponse) ProtoReflect() protoreflect.Message {
	mi := &file_text_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTextResponse.ProtoReflect.Descriptor instead.
func (*CreateTextResponse) Descriptor() ([]byte, []int) {
	return file_text_proto_rawDescGZIP(), []int{2}
}

func (x *CreateTextResponse) GetNote() *NoteMsg {
	if x != nil {
		return x.Note
	}
	return nil
}

func (x *CreateTextResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type GetTextRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *GetTextRequest) Reset() {
	*x = GetTextRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_text_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTextRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTextRequest) ProtoMessage() {}

func (x *GetTextRequest) ProtoReflect() protoreflect.Message {
	mi := &file_text_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTextRequest.ProtoReflect.Descriptor instead.
func (*GetTextRequest) Descriptor() ([]byte, []int) {
	return file_text_proto_rawDescGZIP(), []int{3}
}

func (x *GetTextRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type GetTextResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Notes []*NoteMsg `protobuf:"bytes,1,rep,name=notes,proto3" json:"notes,omitempty"`
	Error string     `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *GetTextResponse) Reset() {
	*x = GetTextResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_text_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTextResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTextResponse) ProtoMessage() {}

func (x *GetTextResponse) ProtoReflect() protoreflect.Message {
	mi := &file_text_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTextResponse.ProtoReflect.Descriptor instead.
func (*GetTextResponse) Descriptor() ([]byte, []int) {
	return file_text_proto_rawDescGZIP(), []int{4}
}

func (x *GetTextResponse) GetNotes() []*NoteMsg {
	if x != nil {
		return x.Notes
	}
	return nil
}

func (x *GetTextResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type UpdateTextRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Note *NoteMsg `protobuf:"bytes,1,opt,name=note,proto3" json:"note,omitempty"`
}

func (x *UpdateTextRequest) Reset() {
	*x = UpdateTextRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_text_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateTextRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateTextRequest) ProtoMessage() {}

func (x *UpdateTextRequest) ProtoReflect() protoreflect.Message {
	mi := &file_text_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateTextRequest.ProtoReflect.Descriptor instead.
func (*UpdateTextRequest) Descriptor() ([]byte, []int) {
	return file_text_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateTextRequest) GetNote() *NoteMsg {
	if x != nil {
		return x.Note
	}
	return nil
}

type UpdateTextResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Note  *NoteMsg `protobuf:"bytes,1,opt,name=note,proto3" json:"note,omitempty"`
	Error string   `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *UpdateTextResponse) Reset() {
	*x = UpdateTextResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_text_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateTextResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateTextResponse) ProtoMessage() {}

func (x *UpdateTextResponse) ProtoReflect() protoreflect.Message {
	mi := &file_text_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateTextResponse.ProtoReflect.Descriptor instead.
func (*UpdateTextResponse) Descriptor() ([]byte, []int) {
	return file_text_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateTextResponse) GetNote() *NoteMsg {
	if x != nil {
		return x.Note
	}
	return nil
}

func (x *UpdateTextResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type DeleteTextResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error string `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *DeleteTextResponse) Reset() {
	*x = DeleteTextResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_text_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteTextResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteTextResponse) ProtoMessage() {}

func (x *DeleteTextResponse) ProtoReflect() protoreflect.Message {
	mi := &file_text_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteTextResponse.ProtoReflect.Descriptor instead.
func (*DeleteTextResponse) Descriptor() ([]byte, []int) {
	return file_text_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteTextResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type DeleteTextRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *DeleteTextRequest) Reset() {
	*x = DeleteTextRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_text_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteTextRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteTextRequest) ProtoMessage() {}

func (x *DeleteTextRequest) ProtoReflect() protoreflect.Message {
	mi := &file_text_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteTextRequest.ProtoReflect.Descriptor instead.
func (*DeleteTextRequest) Descriptor() ([]byte, []int) {
	return file_text_proto_rawDescGZIP(), []int{8}
}

func (x *DeleteTextRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

var File_text_proto protoreflect.FileDescriptor

var file_text_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x74, 0x65, 0x78, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x49, 0x0a, 0x07, 0x4e, 0x6f, 0x74, 0x65, 0x4d, 0x73, 0x67, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x6f, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x6f,
	0x74, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x22, 0x4d,
	0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x65, 0x78, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x22, 0x0a, 0x04, 0x6e, 0x6f, 0x74,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x4e, 0x6f, 0x74, 0x65, 0x4d, 0x73, 0x67, 0x52, 0x04, 0x6e, 0x6f, 0x74, 0x65, 0x22, 0x4e, 0x0a,
	0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x65, 0x78, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x22, 0x0a, 0x04, 0x6e, 0x6f, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4e, 0x6f, 0x74, 0x65, 0x4d, 0x73,
	0x67, 0x52, 0x04, 0x6e, 0x6f, 0x74, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x26, 0x0a,
	0x0e, 0x47, 0x65, 0x74, 0x54, 0x65, 0x78, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x4d, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x54, 0x65, 0x78, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x24, 0x0a, 0x05, 0x6e, 0x6f, 0x74, 0x65,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x4e, 0x6f, 0x74, 0x65, 0x4d, 0x73, 0x67, 0x52, 0x05, 0x6e, 0x6f, 0x74, 0x65, 0x73, 0x12, 0x14,
	0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x22, 0x37, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x65,
	0x78, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x04, 0x6e, 0x6f, 0x74,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x4e, 0x6f, 0x74, 0x65, 0x4d, 0x73, 0x67, 0x52, 0x04, 0x6e, 0x6f, 0x74, 0x65, 0x22, 0x4e, 0x0a,
	0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x65, 0x78, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x22, 0x0a, 0x04, 0x6e, 0x6f, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4e, 0x6f, 0x74, 0x65, 0x4d, 0x73,
	0x67, 0x52, 0x04, 0x6e, 0x6f, 0x74, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x2a, 0x0a,
	0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x65, 0x78, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x29, 0x0a, 0x11, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x54, 0x65, 0x78, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14,
	0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x32, 0x89, 0x02, 0x0a, 0x04, 0x54, 0x65, 0x78, 0x74, 0x12, 0x41, 0x0a,
	0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x65, 0x78, 0x74, 0x12, 0x18, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x65, 0x78, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x54, 0x65, 0x78, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x38, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x54, 0x65, 0x78, 0x74, 0x12, 0x15, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x65, 0x78, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x65,
	0x78, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x41, 0x0a, 0x0a, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x54, 0x65, 0x78, 0x74, 0x12, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x65, 0x78, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x54, 0x65, 0x78, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x41, 0x0a,
	0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x65, 0x78, 0x74, 0x12, 0x18, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x65, 0x78, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x54, 0x65, 0x78, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x42, 0x2f, 0x5a, 0x2d, 0x47, 0x6f, 0x70, 0x68, 0x4b, 0x65, 0x65, 0x70, 0x65, 0x72, 0x2f, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x61,
	0x64, 0x61, 0x70, 0x74, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x65, 0x6e,
	0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_text_proto_rawDescOnce sync.Once
	file_text_proto_rawDescData = file_text_proto_rawDesc
)

func file_text_proto_rawDescGZIP() []byte {
	file_text_proto_rawDescOnce.Do(func() {
		file_text_proto_rawDescData = protoimpl.X.CompressGZIP(file_text_proto_rawDescData)
	})
	return file_text_proto_rawDescData
}

var file_text_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_text_proto_goTypes = []interface{}{
	(*NoteMsg)(nil),            // 0: proto.NoteMsg
	(*CreateTextRequest)(nil),  // 1: proto.CreateTextRequest
	(*CreateTextResponse)(nil), // 2: proto.CreateTextResponse
	(*GetTextRequest)(nil),     // 3: proto.GetTextRequest
	(*GetTextResponse)(nil),    // 4: proto.GetTextResponse
	(*UpdateTextRequest)(nil),  // 5: proto.UpdateTextRequest
	(*UpdateTextResponse)(nil), // 6: proto.UpdateTextResponse
	(*DeleteTextResponse)(nil), // 7: proto.DeleteTextResponse
	(*DeleteTextRequest)(nil),  // 8: proto.DeleteTextRequest
}
var file_text_proto_depIdxs = []int32{
	0, // 0: proto.CreateTextRequest.note:type_name -> proto.NoteMsg
	0, // 1: proto.CreateTextResponse.note:type_name -> proto.NoteMsg
	0, // 2: proto.GetTextResponse.notes:type_name -> proto.NoteMsg
	0, // 3: proto.UpdateTextRequest.note:type_name -> proto.NoteMsg
	0, // 4: proto.UpdateTextResponse.note:type_name -> proto.NoteMsg
	1, // 5: proto.Text.CreateText:input_type -> proto.CreateTextRequest
	3, // 6: proto.Text.GetText:input_type -> proto.GetTextRequest
	5, // 7: proto.Text.UpdateText:input_type -> proto.UpdateTextRequest
	8, // 8: proto.Text.DeleteText:input_type -> proto.DeleteTextRequest
	2, // 9: proto.Text.CreateText:output_type -> proto.CreateTextResponse
	4, // 10: proto.Text.GetText:output_type -> proto.GetTextResponse
	6, // 11: proto.Text.UpdateText:output_type -> proto.UpdateTextResponse
	7, // 12: proto.Text.DeleteText:output_type -> proto.DeleteTextResponse
	9, // [9:13] is the sub-list for method output_type
	5, // [5:9] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_text_proto_init() }
func file_text_proto_init() {
	if File_text_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_text_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NoteMsg); i {
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
		file_text_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateTextRequest); i {
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
		file_text_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateTextResponse); i {
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
		file_text_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTextRequest); i {
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
		file_text_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTextResponse); i {
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
		file_text_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateTextRequest); i {
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
		file_text_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateTextResponse); i {
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
		file_text_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteTextResponse); i {
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
		file_text_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteTextRequest); i {
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
			RawDescriptor: file_text_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_text_proto_goTypes,
		DependencyIndexes: file_text_proto_depIdxs,
		MessageInfos:      file_text_proto_msgTypes,
	}.Build()
	File_text_proto = out.File
	file_text_proto_rawDesc = nil
	file_text_proto_goTypes = nil
	file_text_proto_depIdxs = nil
}
