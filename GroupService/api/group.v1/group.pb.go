// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v3.12.4
// source: api/proto/group.proto

package group_v1

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

type RegisterReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User string `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *RegisterReq) Reset() {
	*x = RegisterReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_group_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterReq) ProtoMessage() {}

func (x *RegisterReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_group_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterReq.ProtoReflect.Descriptor instead.
func (*RegisterReq) Descriptor() ([]byte, []int) {
	return file_api_proto_group_proto_rawDescGZIP(), []int{0}
}

func (x *RegisterReq) GetUser() string {
	if x != nil {
		return x.User
	}
	return ""
}

func (x *RegisterReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type BoolResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *BoolResp) Reset() {
	*x = BoolResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_group_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BoolResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BoolResp) ProtoMessage() {}

func (x *BoolResp) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_group_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BoolResp.ProtoReflect.Descriptor instead.
func (*BoolResp) Descriptor() ([]byte, []int) {
	return file_api_proto_group_proto_rawDescGZIP(), []int{1}
}

func (x *BoolResp) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type GetReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User string `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *GetReq) Reset() {
	*x = GetReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_group_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetReq) ProtoMessage() {}

func (x *GetReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_group_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetReq.ProtoReflect.Descriptor instead.
func (*GetReq) Descriptor() ([]byte, []int) {
	return file_api_proto_group_proto_rawDescGZIP(), []int{2}
}

func (x *GetReq) GetUser() string {
	if x != nil {
		return x.User
	}
	return ""
}

func (x *GetReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type GetCodeResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code string `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
}

func (x *GetCodeResp) Reset() {
	*x = GetCodeResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_group_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCodeResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCodeResp) ProtoMessage() {}

func (x *GetCodeResp) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_group_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCodeResp.ProtoReflect.Descriptor instead.
func (*GetCodeResp) Descriptor() ([]byte, []int) {
	return file_api_proto_group_proto_rawDescGZIP(), []int{3}
}

func (x *GetCodeResp) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

type AddReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code string `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	User string `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *AddReq) Reset() {
	*x = AddReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_group_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddReq) ProtoMessage() {}

func (x *AddReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_group_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddReq.ProtoReflect.Descriptor instead.
func (*AddReq) Descriptor() ([]byte, []int) {
	return file_api_proto_group_proto_rawDescGZIP(), []int{4}
}

func (x *AddReq) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *AddReq) GetUser() string {
	if x != nil {
		return x.User
	}
	return ""
}

type RemoveReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User   string `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	Group  string `protobuf:"bytes,2,opt,name=group,proto3" json:"group,omitempty"`
	Remove string `protobuf:"bytes,3,opt,name=remove,proto3" json:"remove,omitempty"`
}

func (x *RemoveReq) Reset() {
	*x = RemoveReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_group_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveReq) ProtoMessage() {}

func (x *RemoveReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_group_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveReq.ProtoReflect.Descriptor instead.
func (*RemoveReq) Descriptor() ([]byte, []int) {
	return file_api_proto_group_proto_rawDescGZIP(), []int{5}
}

func (x *RemoveReq) GetUser() string {
	if x != nil {
		return x.User
	}
	return ""
}

func (x *RemoveReq) GetGroup() string {
	if x != nil {
		return x.Group
	}
	return ""
}

func (x *RemoveReq) GetRemove() string {
	if x != nil {
		return x.Remove
	}
	return ""
}

type GetMembersResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Members []string `protobuf:"bytes,1,rep,name=members,proto3" json:"members,omitempty"`
}

func (x *GetMembersResp) Reset() {
	*x = GetMembersResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_group_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMembersResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMembersResp) ProtoMessage() {}

func (x *GetMembersResp) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_group_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMembersResp.ProtoReflect.Descriptor instead.
func (*GetMembersResp) Descriptor() ([]byte, []int) {
	return file_api_proto_group_proto_rawDescGZIP(), []int{6}
}

func (x *GetMembersResp) GetMembers() []string {
	if x != nil {
		return x.Members
	}
	return nil
}

var File_api_proto_group_proto protoreflect.FileDescriptor

var file_api_proto_group_proto_rawDesc = []byte{
	0x0a, 0x15, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x72, 0x6f, 0x75,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x22, 0x35,
	0x0a, 0x0b, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x12, 0x12, 0x0a,
	0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x73, 0x65,
	0x72, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x24, 0x0a, 0x08, 0x42, 0x6f, 0x6f, 0x6c, 0x52, 0x65, 0x73,
	0x70, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0x30, 0x0a, 0x06, 0x47,
	0x65, 0x74, 0x52, 0x65, 0x71, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x21, 0x0a,
	0x0b, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x12, 0x0a, 0x04,
	0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65,
	0x22, 0x30, 0x0a, 0x06, 0x41, 0x64, 0x64, 0x52, 0x65, 0x71, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x73,
	0x65, 0x72, 0x22, 0x4d, 0x0a, 0x09, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x52, 0x65, 0x71, 0x12,
	0x12, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75,
	0x73, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x6d,
	0x6f, 0x76, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x6d, 0x6f, 0x76,
	0x65, 0x22, 0x2a, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x32, 0x88, 0x02,
	0x0a, 0x05, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x31, 0x0a, 0x08, 0x52, 0x65, 0x67, 0x69, 0x73,
	0x74, 0x65, 0x72, 0x12, 0x12, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x52, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x1a, 0x0f, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e,
	0x42, 0x6f, 0x6f, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x12, 0x2e, 0x0a, 0x07, 0x47, 0x65,
	0x74, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x0d, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x47, 0x65,
	0x74, 0x52, 0x65, 0x71, 0x1a, 0x12, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x47, 0x65, 0x74,
	0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x12, 0x2e, 0x0a, 0x0a, 0x41, 0x64,
	0x64, 0x54, 0x6f, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x0d, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70,
	0x2e, 0x41, 0x64, 0x64, 0x52, 0x65, 0x71, 0x1a, 0x0f, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e,
	0x42, 0x6f, 0x6f, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x12, 0x36, 0x0a, 0x0f, 0x52, 0x65,
	0x6d, 0x6f, 0x76, 0x65, 0x46, 0x72, 0x6f, 0x6d, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x10, 0x2e,
	0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x52, 0x65, 0x71, 0x1a,
	0x0f, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x52, 0x65, 0x73, 0x70,
	0x22, 0x00, 0x12, 0x34, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73,
	0x12, 0x0d, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x1a,
	0x15, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x6d, 0x62, 0x65,
	0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x42, 0x0a, 0x5a, 0x08, 0x67, 0x72, 0x6f, 0x75,
	0x70, 0x2e, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_proto_group_proto_rawDescOnce sync.Once
	file_api_proto_group_proto_rawDescData = file_api_proto_group_proto_rawDesc
)

func file_api_proto_group_proto_rawDescGZIP() []byte {
	file_api_proto_group_proto_rawDescOnce.Do(func() {
		file_api_proto_group_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_proto_group_proto_rawDescData)
	})
	return file_api_proto_group_proto_rawDescData
}

var file_api_proto_group_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_api_proto_group_proto_goTypes = []interface{}{
	(*RegisterReq)(nil),    // 0: group.RegisterReq
	(*BoolResp)(nil),       // 1: group.BoolResp
	(*GetReq)(nil),         // 2: group.GetReq
	(*GetCodeResp)(nil),    // 3: group.GetCodeResp
	(*AddReq)(nil),         // 4: group.AddReq
	(*RemoveReq)(nil),      // 5: group.RemoveReq
	(*GetMembersResp)(nil), // 6: group.GetMembersResp
}
var file_api_proto_group_proto_depIdxs = []int32{
	0, // 0: group.Group.Register:input_type -> group.RegisterReq
	2, // 1: group.Group.GetCode:input_type -> group.GetReq
	4, // 2: group.Group.AddToGroup:input_type -> group.AddReq
	5, // 3: group.Group.RemoveFromGroup:input_type -> group.RemoveReq
	2, // 4: group.Group.GetMembers:input_type -> group.GetReq
	1, // 5: group.Group.Register:output_type -> group.BoolResp
	3, // 6: group.Group.GetCode:output_type -> group.GetCodeResp
	1, // 7: group.Group.AddToGroup:output_type -> group.BoolResp
	1, // 8: group.Group.RemoveFromGroup:output_type -> group.BoolResp
	6, // 9: group.Group.GetMembers:output_type -> group.GetMembersResp
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_proto_group_proto_init() }
func file_api_proto_group_proto_init() {
	if File_api_proto_group_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_proto_group_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterReq); i {
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
		file_api_proto_group_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BoolResp); i {
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
		file_api_proto_group_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetReq); i {
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
		file_api_proto_group_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCodeResp); i {
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
		file_api_proto_group_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddReq); i {
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
		file_api_proto_group_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveReq); i {
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
		file_api_proto_group_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMembersResp); i {
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
			RawDescriptor: file_api_proto_group_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_proto_group_proto_goTypes,
		DependencyIndexes: file_api_proto_group_proto_depIdxs,
		MessageInfos:      file_api_proto_group_proto_msgTypes,
	}.Build()
	File_api_proto_group_proto = out.File
	file_api_proto_group_proto_rawDesc = nil
	file_api_proto_group_proto_goTypes = nil
	file_api_proto_group_proto_depIdxs = nil
}
