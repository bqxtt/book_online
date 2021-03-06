// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: userpb/user.proto

package userpb

import (
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type USER_ROLE int32

const (
	USER_ROLE_UNKNOWN USER_ROLE = 0
	USER_ROLE_NORMAL  USER_ROLE = 1
	USER_ROLE_ADMIN   USER_ROLE = 2
)

// Enum value maps for USER_ROLE.
var (
	USER_ROLE_name = map[int32]string{
		0: "UNKNOWN",
		1: "NORMAL",
		2: "ADMIN",
	}
	USER_ROLE_value = map[string]int32{
		"UNKNOWN": 0,
		"NORMAL":  1,
		"ADMIN":   2,
	}
)

func (x USER_ROLE) Enum() *USER_ROLE {
	p := new(USER_ROLE)
	*p = x
	return p
}

func (x USER_ROLE) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (USER_ROLE) Descriptor() protoreflect.EnumDescriptor {
	return file_userpb_user_proto_enumTypes[0].Descriptor()
}

func (USER_ROLE) Type() protoreflect.EnumType {
	return &file_userpb_user_proto_enumTypes[0]
}

func (x USER_ROLE) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use USER_ROLE.Descriptor instead.
func (USER_ROLE) EnumDescriptor() ([]byte, []int) {
	return file_userpb_user_proto_rawDescGZIP(), []int{0}
}

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId     int64                  `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Name       string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	AvatarUrl  string                 `protobuf:"bytes,3,opt,name=avatar_url,json=avatarUrl,proto3" json:"avatar_url,omitempty"`
	Role       USER_ROLE              `protobuf:"varint,4,opt,name=role,proto3,enum=sdk.USER_ROLE" json:"role,omitempty"`
	Department string                 `protobuf:"bytes,5,opt,name=department,proto3" json:"department,omitempty"`
	Class      string                 `protobuf:"bytes,6,opt,name=class,proto3" json:"class,omitempty"`
	Motto      string                 `protobuf:"bytes,7,opt,name=motto,proto3" json:"motto,omitempty"`
	CreatedAt  *timestamppb.Timestamp `protobuf:"bytes,8,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_userpb_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_userpb_user_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_userpb_user_proto_rawDescGZIP(), []int{0}
}

func (x *User) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *User) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *User) GetAvatarUrl() string {
	if x != nil {
		return x.AvatarUrl
	}
	return ""
}

func (x *User) GetRole() USER_ROLE {
	if x != nil {
		return x.Role
	}
	return USER_ROLE_UNKNOWN
}

func (x *User) GetDepartment() string {
	if x != nil {
		return x.Department
	}
	return ""
}

func (x *User) GetClass() string {
	if x != nil {
		return x.Class
	}
	return ""
}

func (x *User) GetMotto() string {
	if x != nil {
		return x.Motto
	}
	return ""
}

func (x *User) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

type UserAuth struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId    int64  `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	PwdDigest string `protobuf:"bytes,2,opt,name=pwd_digest,json=pwdDigest,proto3" json:"pwd_digest,omitempty"`
}

func (x *UserAuth) Reset() {
	*x = UserAuth{}
	if protoimpl.UnsafeEnabled {
		mi := &file_userpb_user_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserAuth) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserAuth) ProtoMessage() {}

func (x *UserAuth) ProtoReflect() protoreflect.Message {
	mi := &file_userpb_user_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserAuth.ProtoReflect.Descriptor instead.
func (*UserAuth) Descriptor() ([]byte, []int) {
	return file_userpb_user_proto_rawDescGZIP(), []int{1}
}

func (x *UserAuth) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *UserAuth) GetPwdDigest() string {
	if x != nil {
		return x.PwdDigest
	}
	return ""
}

var File_userpb_user_proto protoreflect.FileDescriptor

var file_userpb_user_proto_rawDesc = []byte{
	0x0a, 0x11, 0x75, 0x73, 0x65, 0x72, 0x70, 0x62, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x03, 0x73, 0x64, 0x6b, 0x1a, 0x27, 0x69, 0x6e, 0x63, 0x6c, 0x75, 0x64,
	0x65, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xfd, 0x01, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x76, 0x61, 0x74, 0x61,
	0x72, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x76, 0x61,
	0x74, 0x61, 0x72, 0x55, 0x72, 0x6c, 0x12, 0x22, 0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x0e, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x55, 0x53, 0x45, 0x52, 0x5f,
	0x52, 0x4f, 0x4c, 0x45, 0x52, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x64, 0x65,
	0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x64, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6c,
	0x61, 0x73, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x63, 0x6c, 0x61, 0x73, 0x73,
	0x12, 0x14, 0x0a, 0x05, 0x6d, 0x6f, 0x74, 0x74, 0x6f, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x6d, 0x6f, 0x74, 0x74, 0x6f, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x5f, 0x61, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x22, 0x42, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x41, 0x75, 0x74, 0x68, 0x12, 0x17, 0x0a,
	0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x77, 0x64, 0x5f, 0x64, 0x69,
	0x67, 0x65, 0x73, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x77, 0x64, 0x44,
	0x69, 0x67, 0x65, 0x73, 0x74, 0x2a, 0x2f, 0x0a, 0x09, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x52, 0x4f,
	0x4c, 0x45, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12,
	0x0a, 0x0a, 0x06, 0x4e, 0x4f, 0x52, 0x4d, 0x41, 0x4c, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x41,
	0x44, 0x4d, 0x49, 0x4e, 0x10, 0x02, 0x42, 0x32, 0x5a, 0x30, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x71, 0x78, 0x74, 0x74, 0x2f, 0x62, 0x6f, 0x6f, 0x6b, 0x5f,
	0x6f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x70, 0x6b, 0x67, 0x2f,
	0x73, 0x64, 0x6b, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_userpb_user_proto_rawDescOnce sync.Once
	file_userpb_user_proto_rawDescData = file_userpb_user_proto_rawDesc
)

func file_userpb_user_proto_rawDescGZIP() []byte {
	file_userpb_user_proto_rawDescOnce.Do(func() {
		file_userpb_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_userpb_user_proto_rawDescData)
	})
	return file_userpb_user_proto_rawDescData
}

var file_userpb_user_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_userpb_user_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_userpb_user_proto_goTypes = []interface{}{
	(USER_ROLE)(0),                // 0: sdk.USER_ROLE
	(*User)(nil),                  // 1: sdk.User
	(*UserAuth)(nil),              // 2: sdk.UserAuth
	(*timestamppb.Timestamp)(nil), // 3: google.protobuf.Timestamp
}
var file_userpb_user_proto_depIdxs = []int32{
	0, // 0: sdk.User.role:type_name -> sdk.USER_ROLE
	3, // 1: sdk.User.created_at:type_name -> google.protobuf.Timestamp
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_userpb_user_proto_init() }
func file_userpb_user_proto_init() {
	if File_userpb_user_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_userpb_user_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
		file_userpb_user_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserAuth); i {
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
			RawDescriptor: file_userpb_user_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_userpb_user_proto_goTypes,
		DependencyIndexes: file_userpb_user_proto_depIdxs,
		EnumInfos:         file_userpb_user_proto_enumTypes,
		MessageInfos:      file_userpb_user_proto_msgTypes,
	}.Build()
	File_userpb_user_proto = out.File
	file_userpb_user_proto_rawDesc = nil
	file_userpb_user_proto_goTypes = nil
	file_userpb_user_proto_depIdxs = nil
}
