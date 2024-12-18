// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.1
// source: secret_feeder.proto

package server

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

type ReqHandshake struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id  string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Pub []byte `protobuf:"bytes,2,opt,name=pub,proto3" json:"pub,omitempty"`
}

func (x *ReqHandshake) Reset() {
	*x = ReqHandshake{}
	if protoimpl.UnsafeEnabled {
		mi := &file_secret_feeder_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReqHandshake) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReqHandshake) ProtoMessage() {}

func (x *ReqHandshake) ProtoReflect() protoreflect.Message {
	mi := &file_secret_feeder_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReqHandshake.ProtoReflect.Descriptor instead.
func (*ReqHandshake) Descriptor() ([]byte, []int) {
	return file_secret_feeder_proto_rawDescGZIP(), []int{0}
}

func (x *ReqHandshake) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ReqHandshake) GetPub() []byte {
	if x != nil {
		return x.Pub
	}
	return nil
}

type RespHandshake struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pub []byte `protobuf:"bytes,1,opt,name=pub,proto3" json:"pub,omitempty"`
}

func (x *RespHandshake) Reset() {
	*x = RespHandshake{}
	if protoimpl.UnsafeEnabled {
		mi := &file_secret_feeder_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RespHandshake) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RespHandshake) ProtoMessage() {}

func (x *RespHandshake) ProtoReflect() protoreflect.Message {
	mi := &file_secret_feeder_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RespHandshake.ProtoReflect.Descriptor instead.
func (*RespHandshake) Descriptor() ([]byte, []int) {
	return file_secret_feeder_proto_rawDescGZIP(), []int{1}
}

func (x *RespHandshake) GetPub() []byte {
	if x != nil {
		return x.Pub
	}
	return nil
}

type ReqNewSecret struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Address []byte `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	Secret  []byte `protobuf:"bytes,3,opt,name=secret,proto3" json:"secret,omitempty"`
}

func (x *ReqNewSecret) Reset() {
	*x = ReqNewSecret{}
	if protoimpl.UnsafeEnabled {
		mi := &file_secret_feeder_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReqNewSecret) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReqNewSecret) ProtoMessage() {}

func (x *ReqNewSecret) ProtoReflect() protoreflect.Message {
	mi := &file_secret_feeder_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReqNewSecret.ProtoReflect.Descriptor instead.
func (*ReqNewSecret) Descriptor() ([]byte, []int) {
	return file_secret_feeder_proto_rawDescGZIP(), []int{2}
}

func (x *ReqNewSecret) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ReqNewSecret) GetAddress() []byte {
	if x != nil {
		return x.Address
	}
	return nil
}

func (x *ReqNewSecret) GetSecret() []byte {
	if x != nil {
		return x.Secret
	}
	return nil
}

type RespNewSecret struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result bool `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *RespNewSecret) Reset() {
	*x = RespNewSecret{}
	if protoimpl.UnsafeEnabled {
		mi := &file_secret_feeder_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RespNewSecret) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RespNewSecret) ProtoMessage() {}

func (x *RespNewSecret) ProtoReflect() protoreflect.Message {
	mi := &file_secret_feeder_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RespNewSecret.ProtoReflect.Descriptor instead.
func (*RespNewSecret) Descriptor() ([]byte, []int) {
	return file_secret_feeder_proto_rawDescGZIP(), []int{3}
}

func (x *RespNewSecret) GetResult() bool {
	if x != nil {
		return x.Result
	}
	return false
}

type ReqGetSecret struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Address []byte `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
}

func (x *ReqGetSecret) Reset() {
	*x = ReqGetSecret{}
	if protoimpl.UnsafeEnabled {
		mi := &file_secret_feeder_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReqGetSecret) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReqGetSecret) ProtoMessage() {}

func (x *ReqGetSecret) ProtoReflect() protoreflect.Message {
	mi := &file_secret_feeder_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReqGetSecret.ProtoReflect.Descriptor instead.
func (*ReqGetSecret) Descriptor() ([]byte, []int) {
	return file_secret_feeder_proto_rawDescGZIP(), []int{4}
}

func (x *ReqGetSecret) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ReqGetSecret) GetAddress() []byte {
	if x != nil {
		return x.Address
	}
	return nil
}

type RespGetSecret struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address []byte `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Secret  []byte `protobuf:"bytes,2,opt,name=secret,proto3" json:"secret,omitempty"`
}

func (x *RespGetSecret) Reset() {
	*x = RespGetSecret{}
	if protoimpl.UnsafeEnabled {
		mi := &file_secret_feeder_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RespGetSecret) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RespGetSecret) ProtoMessage() {}

func (x *RespGetSecret) ProtoReflect() protoreflect.Message {
	mi := &file_secret_feeder_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RespGetSecret.ProtoReflect.Descriptor instead.
func (*RespGetSecret) Descriptor() ([]byte, []int) {
	return file_secret_feeder_proto_rawDescGZIP(), []int{5}
}

func (x *RespGetSecret) GetAddress() []byte {
	if x != nil {
		return x.Address
	}
	return nil
}

func (x *RespGetSecret) GetSecret() []byte {
	if x != nil {
		return x.Secret
	}
	return nil
}

type ReqUpdateSecret struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Address   []byte `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	OldSecret []byte `protobuf:"bytes,3,opt,name=old_secret,json=oldSecret,proto3" json:"old_secret,omitempty"`
	NewSecret []byte `protobuf:"bytes,4,opt,name=new_secret,json=newSecret,proto3" json:"new_secret,omitempty"`
}

func (x *ReqUpdateSecret) Reset() {
	*x = ReqUpdateSecret{}
	if protoimpl.UnsafeEnabled {
		mi := &file_secret_feeder_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReqUpdateSecret) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReqUpdateSecret) ProtoMessage() {}

func (x *ReqUpdateSecret) ProtoReflect() protoreflect.Message {
	mi := &file_secret_feeder_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReqUpdateSecret.ProtoReflect.Descriptor instead.
func (*ReqUpdateSecret) Descriptor() ([]byte, []int) {
	return file_secret_feeder_proto_rawDescGZIP(), []int{6}
}

func (x *ReqUpdateSecret) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ReqUpdateSecret) GetAddress() []byte {
	if x != nil {
		return x.Address
	}
	return nil
}

func (x *ReqUpdateSecret) GetOldSecret() []byte {
	if x != nil {
		return x.OldSecret
	}
	return nil
}

func (x *ReqUpdateSecret) GetNewSecret() []byte {
	if x != nil {
		return x.NewSecret
	}
	return nil
}

type RespUpdateSecret struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result bool `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *RespUpdateSecret) Reset() {
	*x = RespUpdateSecret{}
	if protoimpl.UnsafeEnabled {
		mi := &file_secret_feeder_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RespUpdateSecret) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RespUpdateSecret) ProtoMessage() {}

func (x *RespUpdateSecret) ProtoReflect() protoreflect.Message {
	mi := &file_secret_feeder_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RespUpdateSecret.ProtoReflect.Descriptor instead.
func (*RespUpdateSecret) Descriptor() ([]byte, []int) {
	return file_secret_feeder_proto_rawDescGZIP(), []int{7}
}

func (x *RespUpdateSecret) GetResult() bool {
	if x != nil {
		return x.Result
	}
	return false
}

var File_secret_feeder_proto protoreflect.FileDescriptor

var file_secret_feeder_proto_rawDesc = []byte{
	0x0a, 0x13, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x5f, 0x66, 0x65, 0x65, 0x64, 0x65, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x22, 0x30, 0x0a,
	0x0c, 0x52, 0x65, 0x71, 0x48, 0x61, 0x6e, 0x64, 0x73, 0x68, 0x61, 0x6b, 0x65, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x10, 0x0a,
	0x03, 0x70, 0x75, 0x62, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x03, 0x70, 0x75, 0x62, 0x22,
	0x21, 0x0a, 0x0d, 0x52, 0x65, 0x73, 0x70, 0x48, 0x61, 0x6e, 0x64, 0x73, 0x68, 0x61, 0x6b, 0x65,
	0x12, 0x10, 0x0a, 0x03, 0x70, 0x75, 0x62, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x03, 0x70,
	0x75, 0x62, 0x22, 0x50, 0x0a, 0x0c, 0x52, 0x65, 0x71, 0x4e, 0x65, 0x77, 0x53, 0x65, 0x63, 0x72,
	0x65, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x16, 0x0a, 0x06,
	0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x73, 0x65,
	0x63, 0x72, 0x65, 0x74, 0x22, 0x27, 0x0a, 0x0d, 0x52, 0x65, 0x73, 0x70, 0x4e, 0x65, 0x77, 0x53,
	0x65, 0x63, 0x72, 0x65, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x38, 0x0a,
	0x0c, 0x52, 0x65, 0x71, 0x47, 0x65, 0x74, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x18, 0x0a,
	0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07,
	0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0x41, 0x0a, 0x0d, 0x52, 0x65, 0x73, 0x70, 0x47,
	0x65, 0x74, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x06, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x22, 0x79, 0x0a, 0x0f, 0x52, 0x65,
	0x71, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x18, 0x0a,
	0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07,
	0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x6f, 0x6c, 0x64, 0x5f, 0x73,
	0x65, 0x63, 0x72, 0x65, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x6f, 0x6c, 0x64,
	0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x6e, 0x65, 0x77, 0x5f, 0x73, 0x65,
	0x63, 0x72, 0x65, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x6e, 0x65, 0x77, 0x53,
	0x65, 0x63, 0x72, 0x65, 0x74, 0x22, 0x2a, 0x0a, 0x10, 0x52, 0x65, 0x73, 0x70, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x32, 0x8a, 0x02, 0x0a, 0x0f, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x46, 0x65, 0x65, 0x64,
	0x65, 0x72, 0x53, 0x76, 0x63, 0x12, 0x3a, 0x0a, 0x09, 0x48, 0x61, 0x6e, 0x64, 0x73, 0x68, 0x61,
	0x6b, 0x65, 0x12, 0x14, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x71, 0x48,
	0x61, 0x6e, 0x64, 0x73, 0x68, 0x61, 0x6b, 0x65, 0x1a, 0x15, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x48, 0x61, 0x6e, 0x64, 0x73, 0x68, 0x61, 0x6b, 0x65, 0x22,
	0x00, 0x12, 0x3a, 0x0a, 0x09, 0x4e, 0x65, 0x77, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x12, 0x14,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x71, 0x4e, 0x65, 0x77, 0x53, 0x65,
	0x63, 0x72, 0x65, 0x74, 0x1a, 0x15, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x52, 0x65,
	0x73, 0x70, 0x4e, 0x65, 0x77, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x22, 0x00, 0x12, 0x3a, 0x0a,
	0x09, 0x47, 0x65, 0x74, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x12, 0x14, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x71, 0x47, 0x65, 0x74, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74,
	0x1a, 0x15, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x47, 0x65,
	0x74, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x22, 0x00, 0x12, 0x43, 0x0a, 0x0c, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x12, 0x17, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x2e, 0x52, 0x65, 0x71, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x65, 0x63, 0x72,
	0x65, 0x74, 0x1a, 0x18, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x73, 0x70,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x22, 0x00, 0x42, 0x22,
	0x5a, 0x20, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x65, 0x61,
	0x74, 0x6f, 0x7a, 0x2f, 0x73, 0x66, 0x65, 0x65, 0x64, 0x65, 0x72, 0x2f, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_secret_feeder_proto_rawDescOnce sync.Once
	file_secret_feeder_proto_rawDescData = file_secret_feeder_proto_rawDesc
)

func file_secret_feeder_proto_rawDescGZIP() []byte {
	file_secret_feeder_proto_rawDescOnce.Do(func() {
		file_secret_feeder_proto_rawDescData = protoimpl.X.CompressGZIP(file_secret_feeder_proto_rawDescData)
	})
	return file_secret_feeder_proto_rawDescData
}

var file_secret_feeder_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_secret_feeder_proto_goTypes = []any{
	(*ReqHandshake)(nil),     // 0: server.ReqHandshake
	(*RespHandshake)(nil),    // 1: server.RespHandshake
	(*ReqNewSecret)(nil),     // 2: server.ReqNewSecret
	(*RespNewSecret)(nil),    // 3: server.RespNewSecret
	(*ReqGetSecret)(nil),     // 4: server.ReqGetSecret
	(*RespGetSecret)(nil),    // 5: server.RespGetSecret
	(*ReqUpdateSecret)(nil),  // 6: server.ReqUpdateSecret
	(*RespUpdateSecret)(nil), // 7: server.RespUpdateSecret
}
var file_secret_feeder_proto_depIdxs = []int32{
	0, // 0: server.SecretFeederSvc.Handshake:input_type -> server.ReqHandshake
	2, // 1: server.SecretFeederSvc.NewSecret:input_type -> server.ReqNewSecret
	4, // 2: server.SecretFeederSvc.GetSecret:input_type -> server.ReqGetSecret
	6, // 3: server.SecretFeederSvc.UpdateSecret:input_type -> server.ReqUpdateSecret
	1, // 4: server.SecretFeederSvc.Handshake:output_type -> server.RespHandshake
	3, // 5: server.SecretFeederSvc.NewSecret:output_type -> server.RespNewSecret
	5, // 6: server.SecretFeederSvc.GetSecret:output_type -> server.RespGetSecret
	7, // 7: server.SecretFeederSvc.UpdateSecret:output_type -> server.RespUpdateSecret
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_secret_feeder_proto_init() }
func file_secret_feeder_proto_init() {
	if File_secret_feeder_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_secret_feeder_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*ReqHandshake); i {
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
		file_secret_feeder_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*RespHandshake); i {
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
		file_secret_feeder_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*ReqNewSecret); i {
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
		file_secret_feeder_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*RespNewSecret); i {
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
		file_secret_feeder_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*ReqGetSecret); i {
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
		file_secret_feeder_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*RespGetSecret); i {
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
		file_secret_feeder_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*ReqUpdateSecret); i {
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
		file_secret_feeder_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*RespUpdateSecret); i {
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
			RawDescriptor: file_secret_feeder_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_secret_feeder_proto_goTypes,
		DependencyIndexes: file_secret_feeder_proto_depIdxs,
		MessageInfos:      file_secret_feeder_proto_msgTypes,
	}.Build()
	File_secret_feeder_proto = out.File
	file_secret_feeder_proto_rawDesc = nil
	file_secret_feeder_proto_goTypes = nil
	file_secret_feeder_proto_depIdxs = nil
}
