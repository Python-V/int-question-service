// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.1
// source: questionset.proto

package question

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

type QuestionSet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	QuestionId  uint32 `protobuf:"varint,2,opt,name=question_id,json=questionId,proto3" json:"question_id,omitempty"`
	SetId       uint32 `protobuf:"varint,3,opt,name=set_id,json=setId,proto3" json:"set_id,omitempty"`
	CreatedById uint32 `protobuf:"varint,4,opt,name=created_by_id,json=createdById,proto3" json:"created_by_id,omitempty"`
	CreatedDate string `protobuf:"bytes,5,opt,name=created_date,json=createdDate,proto3" json:"created_date,omitempty"`
	UpdatedDate string `protobuf:"bytes,6,opt,name=updated_date,json=updatedDate,proto3" json:"updated_date,omitempty"`
}

func (x *QuestionSet) Reset() {
	*x = QuestionSet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_questionset_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QuestionSet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QuestionSet) ProtoMessage() {}

func (x *QuestionSet) ProtoReflect() protoreflect.Message {
	mi := &file_questionset_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QuestionSet.ProtoReflect.Descriptor instead.
func (*QuestionSet) Descriptor() ([]byte, []int) {
	return file_questionset_proto_rawDescGZIP(), []int{0}
}

func (x *QuestionSet) GetQuestionId() uint32 {
	if x != nil {
		return x.QuestionId
	}
	return 0
}

func (x *QuestionSet) GetSetId() uint32 {
	if x != nil {
		return x.SetId
	}
	return 0
}

func (x *QuestionSet) GetCreatedById() uint32 {
	if x != nil {
		return x.CreatedById
	}
	return 0
}

func (x *QuestionSet) GetCreatedDate() string {
	if x != nil {
		return x.CreatedDate
	}
	return ""
}

func (x *QuestionSet) GetUpdatedDate() string {
	if x != nil {
		return x.UpdatedDate
	}
	return ""
}

type CreateQuestionSetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	QuestionSet *QuestionSet `protobuf:"bytes,1,opt,name=question_set,json=questionSet,proto3" json:"question_set,omitempty"`
}

func (x *CreateQuestionSetRequest) Reset() {
	*x = CreateQuestionSetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_questionset_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateQuestionSetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateQuestionSetRequest) ProtoMessage() {}

func (x *CreateQuestionSetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_questionset_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateQuestionSetRequest.ProtoReflect.Descriptor instead.
func (*CreateQuestionSetRequest) Descriptor() ([]byte, []int) {
	return file_questionset_proto_rawDescGZIP(), []int{1}
}

func (x *CreateQuestionSetRequest) GetQuestionSet() *QuestionSet {
	if x != nil {
		return x.QuestionSet
	}
	return nil
}

type CreateQuestionSetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	QuestionSet *QuestionSet `protobuf:"bytes,1,opt,name=question_set,json=questionSet,proto3" json:"question_set,omitempty"`
}

func (x *CreateQuestionSetResponse) Reset() {
	*x = CreateQuestionSetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_questionset_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateQuestionSetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateQuestionSetResponse) ProtoMessage() {}

func (x *CreateQuestionSetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_questionset_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateQuestionSetResponse.ProtoReflect.Descriptor instead.
func (*CreateQuestionSetResponse) Descriptor() ([]byte, []int) {
	return file_questionset_proto_rawDescGZIP(), []int{2}
}

func (x *CreateQuestionSetResponse) GetQuestionSet() *QuestionSet {
	if x != nil {
		return x.QuestionSet
	}
	return nil
}

type UpdateQuestionSetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	QuestionSet *QuestionSet `protobuf:"bytes,1,opt,name=question_set,json=questionSet,proto3" json:"question_set,omitempty"`
	Id          uint32       `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *UpdateQuestionSetRequest) Reset() {
	*x = UpdateQuestionSetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_questionset_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateQuestionSetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateQuestionSetRequest) ProtoMessage() {}

func (x *UpdateQuestionSetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_questionset_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateQuestionSetRequest.ProtoReflect.Descriptor instead.
func (*UpdateQuestionSetRequest) Descriptor() ([]byte, []int) {
	return file_questionset_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateQuestionSetRequest) GetQuestionSet() *QuestionSet {
	if x != nil {
		return x.QuestionSet
	}
	return nil
}

func (x *UpdateQuestionSetRequest) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type UpdateQuestionSetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	QuestionSet *QuestionSet `protobuf:"bytes,1,opt,name=question_set,json=questionSet,proto3" json:"question_set,omitempty"`
}

func (x *UpdateQuestionSetResponse) Reset() {
	*x = UpdateQuestionSetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_questionset_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateQuestionSetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateQuestionSetResponse) ProtoMessage() {}

func (x *UpdateQuestionSetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_questionset_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateQuestionSetResponse.ProtoReflect.Descriptor instead.
func (*UpdateQuestionSetResponse) Descriptor() ([]byte, []int) {
	return file_questionset_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateQuestionSetResponse) GetQuestionSet() *QuestionSet {
	if x != nil {
		return x.QuestionSet
	}
	return nil
}

type ListQuestionSetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListQuestionSetRequest) Reset() {
	*x = ListQuestionSetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_questionset_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListQuestionSetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListQuestionSetRequest) ProtoMessage() {}

func (x *ListQuestionSetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_questionset_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListQuestionSetRequest.ProtoReflect.Descriptor instead.
func (*ListQuestionSetRequest) Descriptor() ([]byte, []int) {
	return file_questionset_proto_rawDescGZIP(), []int{5}
}

type ListQuestionSetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	QuestionSet []*QuestionSet `protobuf:"bytes,1,rep,name=question_set,json=questionSet,proto3" json:"question_set,omitempty"`
}

func (x *ListQuestionSetResponse) Reset() {
	*x = ListQuestionSetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_questionset_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListQuestionSetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListQuestionSetResponse) ProtoMessage() {}

func (x *ListQuestionSetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_questionset_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListQuestionSetResponse.ProtoReflect.Descriptor instead.
func (*ListQuestionSetResponse) Descriptor() ([]byte, []int) {
	return file_questionset_proto_rawDescGZIP(), []int{6}
}

func (x *ListQuestionSetResponse) GetQuestionSet() []*QuestionSet {
	if x != nil {
		return x.QuestionSet
	}
	return nil
}

type RetrieveQuestionSetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *RetrieveQuestionSetRequest) Reset() {
	*x = RetrieveQuestionSetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_questionset_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RetrieveQuestionSetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RetrieveQuestionSetRequest) ProtoMessage() {}

func (x *RetrieveQuestionSetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_questionset_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RetrieveQuestionSetRequest.ProtoReflect.Descriptor instead.
func (*RetrieveQuestionSetRequest) Descriptor() ([]byte, []int) {
	return file_questionset_proto_rawDescGZIP(), []int{7}
}

func (x *RetrieveQuestionSetRequest) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type RetrieveQuestionSetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	QuestionSet *QuestionSet `protobuf:"bytes,1,opt,name=question_set,json=questionSet,proto3" json:"question_set,omitempty"`
}

func (x *RetrieveQuestionSetResponse) Reset() {
	*x = RetrieveQuestionSetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_questionset_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RetrieveQuestionSetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RetrieveQuestionSetResponse) ProtoMessage() {}

func (x *RetrieveQuestionSetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_questionset_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RetrieveQuestionSetResponse.ProtoReflect.Descriptor instead.
func (*RetrieveQuestionSetResponse) Descriptor() ([]byte, []int) {
	return file_questionset_proto_rawDescGZIP(), []int{8}
}

func (x *RetrieveQuestionSetResponse) GetQuestionSet() *QuestionSet {
	if x != nil {
		return x.QuestionSet
	}
	return nil
}

type DeleteQuestionSetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteQuestionSetRequest) Reset() {
	*x = DeleteQuestionSetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_questionset_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteQuestionSetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteQuestionSetRequest) ProtoMessage() {}

func (x *DeleteQuestionSetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_questionset_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteQuestionSetRequest.ProtoReflect.Descriptor instead.
func (*DeleteQuestionSetRequest) Descriptor() ([]byte, []int) {
	return file_questionset_proto_rawDescGZIP(), []int{9}
}

func (x *DeleteQuestionSetRequest) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type DeleteQuestionSetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success  bool   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Response string `protobuf:"bytes,2,opt,name=response,proto3" json:"response,omitempty"`
}

func (x *DeleteQuestionSetResponse) Reset() {
	*x = DeleteQuestionSetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_questionset_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteQuestionSetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteQuestionSetResponse) ProtoMessage() {}

func (x *DeleteQuestionSetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_questionset_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteQuestionSetResponse.ProtoReflect.Descriptor instead.
func (*DeleteQuestionSetResponse) Descriptor() ([]byte, []int) {
	return file_questionset_proto_rawDescGZIP(), []int{10}
}

func (x *DeleteQuestionSetResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *DeleteQuestionSetResponse) GetResponse() string {
	if x != nil {
		return x.Response
	}
	return ""
}

var File_questionset_proto protoreflect.FileDescriptor

var file_questionset_proto_rawDesc = []byte{
	0x0a, 0x11, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x65, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xb9, 0x01, 0x0a, 0x0b, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e,
	0x53, 0x65, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69,
	0x6f, 0x6e, 0x49, 0x64, 0x12, 0x15, 0x0a, 0x06, 0x73, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x73, 0x65, 0x74, 0x49, 0x64, 0x12, 0x22, 0x0a, 0x0d, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x62, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x49, 0x64, 0x12,
	0x21, 0x0a, 0x0c, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x44, 0x61,
	0x74, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x64, 0x61,
	0x74, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x44, 0x61, 0x74, 0x65, 0x4a, 0x04, 0x08, 0x01, 0x10, 0x02, 0x52, 0x02, 0x69, 0x64, 0x22,
	0x4b, 0x0a, 0x18, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f,
	0x6e, 0x53, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2f, 0x0a, 0x0c, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0c, 0x2e, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x74, 0x52,
	0x0b, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x74, 0x22, 0x4c, 0x0a, 0x19,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2f, 0x0a, 0x0c, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0c, 0x2e, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x74, 0x52, 0x0b, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x74, 0x22, 0x5b, 0x0a, 0x18, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2f, 0x0a, 0x0c, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x73, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x51,
	0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x74, 0x52, 0x0b, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x22, 0x4c, 0x0a, 0x19, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2f, 0x0a, 0x0c, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x73, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x51, 0x75, 0x65,
	0x73, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x74, 0x52, 0x0b, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69,
	0x6f, 0x6e, 0x53, 0x65, 0x74, 0x22, 0x18, 0x0a, 0x16, 0x4c, 0x69, 0x73, 0x74, 0x51, 0x75, 0x65,
	0x73, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22,
	0x4a, 0x0a, 0x17, 0x4c, 0x69, 0x73, 0x74, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x53,
	0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2f, 0x0a, 0x0c, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x65, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x0c, 0x2e, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x74, 0x52, 0x0b,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x74, 0x22, 0x2c, 0x0a, 0x1a, 0x52,
	0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x65, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x53,
	0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x22, 0x4e, 0x0a, 0x1b, 0x52, 0x65, 0x74,
	0x72, 0x69, 0x65, 0x76, 0x65, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2f, 0x0a, 0x0c, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c,
	0x2e, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x74, 0x52, 0x0b, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x74, 0x22, 0x2a, 0x0a, 0x18, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x02, 0x69, 0x64, 0x22, 0x51, 0x0a, 0x19, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x51,
	0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x1a, 0x0a, 0x08,
	0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x11, 0x5a, 0x0f, 0x2e, 0x2f, 0x6c, 0x69,
	0x62, 0x73, 0x3b, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_questionset_proto_rawDescOnce sync.Once
	file_questionset_proto_rawDescData = file_questionset_proto_rawDesc
)

func file_questionset_proto_rawDescGZIP() []byte {
	file_questionset_proto_rawDescOnce.Do(func() {
		file_questionset_proto_rawDescData = protoimpl.X.CompressGZIP(file_questionset_proto_rawDescData)
	})
	return file_questionset_proto_rawDescData
}

var file_questionset_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_questionset_proto_goTypes = []interface{}{
	(*QuestionSet)(nil),                 // 0: QuestionSet
	(*CreateQuestionSetRequest)(nil),    // 1: CreateQuestionSetRequest
	(*CreateQuestionSetResponse)(nil),   // 2: CreateQuestionSetResponse
	(*UpdateQuestionSetRequest)(nil),    // 3: UpdateQuestionSetRequest
	(*UpdateQuestionSetResponse)(nil),   // 4: UpdateQuestionSetResponse
	(*ListQuestionSetRequest)(nil),      // 5: ListQuestionSetRequest
	(*ListQuestionSetResponse)(nil),     // 6: ListQuestionSetResponse
	(*RetrieveQuestionSetRequest)(nil),  // 7: RetrieveQuestionSetRequest
	(*RetrieveQuestionSetResponse)(nil), // 8: RetrieveQuestionSetResponse
	(*DeleteQuestionSetRequest)(nil),    // 9: DeleteQuestionSetRequest
	(*DeleteQuestionSetResponse)(nil),   // 10: DeleteQuestionSetResponse
}
var file_questionset_proto_depIdxs = []int32{
	0, // 0: CreateQuestionSetRequest.question_set:type_name -> QuestionSet
	0, // 1: CreateQuestionSetResponse.question_set:type_name -> QuestionSet
	0, // 2: UpdateQuestionSetRequest.question_set:type_name -> QuestionSet
	0, // 3: UpdateQuestionSetResponse.question_set:type_name -> QuestionSet
	0, // 4: ListQuestionSetResponse.question_set:type_name -> QuestionSet
	0, // 5: RetrieveQuestionSetResponse.question_set:type_name -> QuestionSet
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_questionset_proto_init() }
func file_questionset_proto_init() {
	if File_questionset_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_questionset_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QuestionSet); i {
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
		file_questionset_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateQuestionSetRequest); i {
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
		file_questionset_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateQuestionSetResponse); i {
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
		file_questionset_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateQuestionSetRequest); i {
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
		file_questionset_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateQuestionSetResponse); i {
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
		file_questionset_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListQuestionSetRequest); i {
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
		file_questionset_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListQuestionSetResponse); i {
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
		file_questionset_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RetrieveQuestionSetRequest); i {
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
		file_questionset_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RetrieveQuestionSetResponse); i {
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
		file_questionset_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteQuestionSetRequest); i {
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
		file_questionset_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteQuestionSetResponse); i {
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
			RawDescriptor: file_questionset_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_questionset_proto_goTypes,
		DependencyIndexes: file_questionset_proto_depIdxs,
		MessageInfos:      file_questionset_proto_msgTypes,
	}.Build()
	File_questionset_proto = out.File
	file_questionset_proto_rawDesc = nil
	file_questionset_proto_goTypes = nil
	file_questionset_proto_depIdxs = nil
}
