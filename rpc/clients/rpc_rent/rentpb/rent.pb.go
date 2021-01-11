// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: rentpb/rent.proto

package rentpb

import (
	bookpb "github.com/bqxtt/book_online/book/pkg/sdk/bookpb"
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

type RentStatus int32

const (
	RentStatus_FINISHED   RentStatus = 0
	RentStatus_IN_PROCESS RentStatus = 1
)

// Enum value maps for RentStatus.
var (
	RentStatus_name = map[int32]string{
		0: "FINISHED",
		1: "IN_PROCESS",
	}
	RentStatus_value = map[string]int32{
		"FINISHED":   0,
		"IN_PROCESS": 1,
	}
)

func (x RentStatus) Enum() *RentStatus {
	p := new(RentStatus)
	*p = x
	return p
}

func (x RentStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RentStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_rentpb_rent_proto_enumTypes[0].Descriptor()
}

func (RentStatus) Type() protoreflect.EnumType {
	return &file_rentpb_rent_proto_enumTypes[0]
}

func (x RentStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RentStatus.Descriptor instead.
func (RentStatus) EnumDescriptor() ([]byte, []int) {
	return file_rentpb_rent_proto_rawDescGZIP(), []int{0}
}

type RentRecord struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId     int64                  `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Book       *bookpb.Book           `protobuf:"bytes,3,opt,name=book,proto3" json:"book,omitempty"`
	Status     RentStatus             `protobuf:"varint,4,opt,name=status,proto3,enum=sdk.RentStatus" json:"status,omitempty"`
	BorrowedAt *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=borrowed_at,json=borrowedAt,proto3" json:"borrowed_at,omitempty"`
	ReturnedAt *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=returned_at,json=returnedAt,proto3" json:"returned_at,omitempty"`
	Deadline   *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=deadline,proto3" json:"deadline,omitempty"`
}

func (x *RentRecord) Reset() {
	*x = RentRecord{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rentpb_rent_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RentRecord) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RentRecord) ProtoMessage() {}

func (x *RentRecord) ProtoReflect() protoreflect.Message {
	mi := &file_rentpb_rent_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RentRecord.ProtoReflect.Descriptor instead.
func (*RentRecord) Descriptor() ([]byte, []int) {
	return file_rentpb_rent_proto_rawDescGZIP(), []int{0}
}

func (x *RentRecord) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *RentRecord) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *RentRecord) GetBook() *bookpb.Book {
	if x != nil {
		return x.Book
	}
	return nil
}

func (x *RentRecord) GetStatus() RentStatus {
	if x != nil {
		return x.Status
	}
	return RentStatus_FINISHED
}

func (x *RentRecord) GetBorrowedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.BorrowedAt
	}
	return nil
}

func (x *RentRecord) GetReturnedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.ReturnedAt
	}
	return nil
}

func (x *RentRecord) GetDeadline() *timestamppb.Timestamp {
	if x != nil {
		return x.Deadline
	}
	return nil
}

var File_rentpb_rent_proto protoreflect.FileDescriptor

var file_rentpb_rent_proto_rawDesc = []byte{
	0x0a, 0x11, 0x72, 0x65, 0x6e, 0x74, 0x70, 0x62, 0x2f, 0x72, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x03, 0x73, 0x64, 0x6b, 0x1a, 0x27, 0x69, 0x6e, 0x63, 0x6c, 0x75, 0x64,
	0x65, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x19, 0x69, 0x6e, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x2f, 0x62, 0x6f, 0x6f, 0x6b, 0x70,
	0x62, 0x2f, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xaf, 0x02, 0x0a,
	0x0a, 0x52, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x04, 0x62, 0x6f, 0x6f, 0x6b, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x09, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x04, 0x62,
	0x6f, 0x6f, 0x6b, 0x12, 0x27, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x0f, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x52, 0x65, 0x6e, 0x74, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x3b, 0x0a, 0x0b,
	0x62, 0x6f, 0x72, 0x72, 0x6f, 0x77, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x62,
	0x6f, 0x72, 0x72, 0x6f, 0x77, 0x65, 0x64, 0x41, 0x74, 0x12, 0x3b, 0x0a, 0x0b, 0x72, 0x65, 0x74,
	0x75, 0x72, 0x6e, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x72, 0x65, 0x74, 0x75,
	0x72, 0x6e, 0x65, 0x64, 0x41, 0x74, 0x12, 0x36, 0x0a, 0x08, 0x64, 0x65, 0x61, 0x64, 0x6c, 0x69,
	0x6e, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x08, 0x64, 0x65, 0x61, 0x64, 0x6c, 0x69, 0x6e, 0x65, 0x2a, 0x2a,
	0x0a, 0x0a, 0x52, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0c, 0x0a, 0x08,
	0x46, 0x49, 0x4e, 0x49, 0x53, 0x48, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0e, 0x0a, 0x0a, 0x49, 0x4e,
	0x5f, 0x50, 0x52, 0x4f, 0x43, 0x45, 0x53, 0x53, 0x10, 0x01, 0x42, 0x32, 0x5a, 0x30, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x71, 0x78, 0x74, 0x74, 0x2f, 0x62,
	0x6f, 0x6f, 0x6b, 0x5f, 0x6f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x2f, 0x72, 0x65, 0x6e, 0x74, 0x2f,
	0x70, 0x6b, 0x67, 0x2f, 0x73, 0x64, 0x6b, 0x2f, 0x72, 0x65, 0x6e, 0x74, 0x70, 0x62, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_rentpb_rent_proto_rawDescOnce sync.Once
	file_rentpb_rent_proto_rawDescData = file_rentpb_rent_proto_rawDesc
)

func file_rentpb_rent_proto_rawDescGZIP() []byte {
	file_rentpb_rent_proto_rawDescOnce.Do(func() {
		file_rentpb_rent_proto_rawDescData = protoimpl.X.CompressGZIP(file_rentpb_rent_proto_rawDescData)
	})
	return file_rentpb_rent_proto_rawDescData
}

var file_rentpb_rent_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_rentpb_rent_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_rentpb_rent_proto_goTypes = []interface{}{
	(RentStatus)(0),               // 0: sdk.RentStatus
	(*RentRecord)(nil),            // 1: sdk.RentRecord
	(*bookpb.Book)(nil),           // 2: sdk.Book
	(*timestamppb.Timestamp)(nil), // 3: google.protobuf.Timestamp
}
var file_rentpb_rent_proto_depIdxs = []int32{
	2, // 0: sdk.RentRecord.book:type_name -> sdk.Book
	0, // 1: sdk.RentRecord.status:type_name -> sdk.RentStatus
	3, // 2: sdk.RentRecord.borrowed_at:type_name -> google.protobuf.Timestamp
	3, // 3: sdk.RentRecord.returned_at:type_name -> google.protobuf.Timestamp
	3, // 4: sdk.RentRecord.deadline:type_name -> google.protobuf.Timestamp
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_rentpb_rent_proto_init() }
func file_rentpb_rent_proto_init() {
	if File_rentpb_rent_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_rentpb_rent_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RentRecord); i {
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
			RawDescriptor: file_rentpb_rent_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_rentpb_rent_proto_goTypes,
		DependencyIndexes: file_rentpb_rent_proto_depIdxs,
		EnumInfos:         file_rentpb_rent_proto_enumTypes,
		MessageInfos:      file_rentpb_rent_proto_msgTypes,
	}.Build()
	File_rentpb_rent_proto = out.File
	file_rentpb_rent_proto_rawDesc = nil
	file_rentpb_rent_proto_goTypes = nil
	file_rentpb_rent_proto_depIdxs = nil
}
