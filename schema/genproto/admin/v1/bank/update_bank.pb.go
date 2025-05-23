// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v5.27.0
// source: update_bank.proto

package admin

import (
	_ "github.com/srikrsna/protoc-gen-gotag/tagger"
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

type UpdateBankRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty" binding:"required"`
	Content    string   `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty" binding:"required"`
	OpenGrade  []string `protobuf:"bytes,3,rep,name=open_grade,json=openGrade,proto3" json:"open_grade,omitempty"`
	QuestionId []int32  `protobuf:"varint,4,rep,packed,name=question_id,json=questionId,proto3" json:"question_id,omitempty"`
}

func (x *UpdateBankRequest) Reset() {
	*x = UpdateBankRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_update_bank_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateBankRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateBankRequest) ProtoMessage() {}

func (x *UpdateBankRequest) ProtoReflect() protoreflect.Message {
	mi := &file_update_bank_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateBankRequest.ProtoReflect.Descriptor instead.
func (*UpdateBankRequest) Descriptor() ([]byte, []int) {
	return file_update_bank_proto_rawDescGZIP(), []int{0}
}

func (x *UpdateBankRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdateBankRequest) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *UpdateBankRequest) GetOpenGrade() []string {
	if x != nil {
		return x.OpenGrade
	}
	return nil
}

func (x *UpdateBankRequest) GetQuestionId() []int32 {
	if x != nil {
		return x.QuestionId
	}
	return nil
}

type UpdateBankResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateBankResponse) Reset() {
	*x = UpdateBankResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_update_bank_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateBankResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateBankResponse) ProtoMessage() {}

func (x *UpdateBankResponse) ProtoReflect() protoreflect.Message {
	mi := &file_update_bank_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateBankResponse.ProtoReflect.Descriptor instead.
func (*UpdateBankResponse) Descriptor() ([]byte, []int) {
	return file_update_bank_proto_rawDescGZIP(), []int{1}
}

var File_update_bank_proto protoreflect.FileDescriptor

var file_update_bank_proto_rawDesc = []byte{
	0x0a, 0x11, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x62, 0x61, 0x6e, 0x6b, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x05, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x1a, 0x1f, 0x74, 0x68, 0x69, 0x72,
	0x64, 0x5f, 0x70, 0x61, 0x72, 0x74, 0x79, 0x2f, 0x74, 0x61, 0x67, 0x67, 0x65, 0x72, 0x2f, 0x74,
	0x61, 0x67, 0x67, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xaf, 0x01, 0x0a, 0x11,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x61, 0x6e, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x27, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x42, 0x17, 0x9a,
	0x84, 0x9e, 0x03, 0x12, 0x62, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x3a, 0x22, 0x72, 0x65, 0x71,
	0x75, 0x69, 0x72, 0x65, 0x64, 0x22, 0x52, 0x02, 0x69, 0x64, 0x12, 0x31, 0x0a, 0x07, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x17, 0x9a, 0x84, 0x9e,
	0x03, 0x12, 0x62, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x3a, 0x22, 0x72, 0x65, 0x71, 0x75, 0x69,
	0x72, 0x65, 0x64, 0x22, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x1d, 0x0a,
	0x0a, 0x6f, 0x70, 0x65, 0x6e, 0x5f, 0x67, 0x72, 0x61, 0x64, 0x65, 0x18, 0x03, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x09, 0x6f, 0x70, 0x65, 0x6e, 0x47, 0x72, 0x61, 0x64, 0x65, 0x12, 0x1f, 0x0a, 0x0b,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x03, 0x28,
	0x05, 0x52, 0x0a, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x22, 0x14, 0x0a,
	0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x61, 0x6e, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x42, 0x15, 0x5a, 0x13, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x76, 0x31, 0x2f,
	0x62, 0x61, 0x6e, 0x6b, 0x3b, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_update_bank_proto_rawDescOnce sync.Once
	file_update_bank_proto_rawDescData = file_update_bank_proto_rawDesc
)

func file_update_bank_proto_rawDescGZIP() []byte {
	file_update_bank_proto_rawDescOnce.Do(func() {
		file_update_bank_proto_rawDescData = protoimpl.X.CompressGZIP(file_update_bank_proto_rawDescData)
	})
	return file_update_bank_proto_rawDescData
}

var file_update_bank_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_update_bank_proto_goTypes = []interface{}{
	(*UpdateBankRequest)(nil),  // 0: admin.UpdateBankRequest
	(*UpdateBankResponse)(nil), // 1: admin.UpdateBankResponse
}
var file_update_bank_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_update_bank_proto_init() }
func file_update_bank_proto_init() {
	if File_update_bank_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_update_bank_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateBankRequest); i {
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
		file_update_bank_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateBankResponse); i {
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
			RawDescriptor: file_update_bank_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_update_bank_proto_goTypes,
		DependencyIndexes: file_update_bank_proto_depIdxs,
		MessageInfos:      file_update_bank_proto_msgTypes,
	}.Build()
	File_update_bank_proto = out.File
	file_update_bank_proto_rawDesc = nil
	file_update_bank_proto_goTypes = nil
	file_update_bank_proto_depIdxs = nil
}
