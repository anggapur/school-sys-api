// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.1
// source: api/proto/file/file-service.proto

package file

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type FileUploadRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *FileUploadRequest) Reset() {
	*x = FileUploadRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_file_file_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileUploadRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileUploadRequest) ProtoMessage() {}

func (x *FileUploadRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_file_file_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileUploadRequest.ProtoReflect.Descriptor instead.
func (*FileUploadRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_file_file_service_proto_rawDescGZIP(), []int{0}
}

func (x *FileUploadRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type FileUploadDetailResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileDetail *FileUploadDetail `protobuf:"bytes,1,opt,name=fileDetail,proto3" json:"fileDetail,omitempty"`
	Status     FileUploadStatus  `protobuf:"varint,2,opt,name=status,proto3,enum=FileUploadStatus" json:"status,omitempty"`
}

func (x *FileUploadDetailResponse) Reset() {
	*x = FileUploadDetailResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_file_file_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileUploadDetailResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileUploadDetailResponse) ProtoMessage() {}

func (x *FileUploadDetailResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_file_file_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileUploadDetailResponse.ProtoReflect.Descriptor instead.
func (*FileUploadDetailResponse) Descriptor() ([]byte, []int) {
	return file_api_proto_file_file_service_proto_rawDescGZIP(), []int{1}
}

func (x *FileUploadDetailResponse) GetFileDetail() *FileUploadDetail {
	if x != nil {
		return x.FileDetail
	}
	return nil
}

func (x *FileUploadDetailResponse) GetStatus() FileUploadStatus {
	if x != nil {
		return x.Status
	}
	return FileUploadStatus_FILE_DRAFT
}

var File_api_proto_file_file_service_proto protoreflect.FileDescriptor

var file_api_proto_file_file_service_proto_rawDesc = []byte{
	0x0a, 0x21, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x66, 0x69, 0x6c, 0x65,
	0x2f, 0x66, 0x69, 0x6c, 0x65, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x66,
	0x69, 0x6c, 0x65, 0x2f, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x26,
	0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x27, 0x0a, 0x11, 0x46, 0x69, 0x6c, 0x65, 0x55, 0x70,
	0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22,
	0x78, 0x0a, 0x18, 0x46, 0x69, 0x6c, 0x65, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x44, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x31, 0x0a, 0x0a, 0x66,
	0x69, 0x6c, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x11, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x44, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x52, 0x0a, 0x66, 0x69, 0x6c, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x29,
	0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x11,
	0x2e, 0x46, 0x69, 0x6c, 0x65, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x32, 0x78, 0x0a, 0x12, 0x46, 0x69, 0x6c,
	0x65, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x12,
	0x62, 0x0a, 0x11, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x55, 0x70,
	0x6c, 0x6f, 0x61, 0x64, 0x12, 0x12, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x55, 0x70, 0x6c, 0x6f, 0x61,
	0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x55,
	0x70, 0x6c, 0x6f, 0x61, 0x64, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x1e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x18, 0x3a, 0x01, 0x2a, 0x22, 0x13,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x66, 0x69, 0x6c, 0x65, 0x2f, 0x75, 0x70, 0x6c,
	0x6f, 0x61, 0x64, 0x42, 0x07, 0x5a, 0x05, 0x66, 0x69, 0x6c, 0x65, 0x2f, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_proto_file_file_service_proto_rawDescOnce sync.Once
	file_api_proto_file_file_service_proto_rawDescData = file_api_proto_file_file_service_proto_rawDesc
)

func file_api_proto_file_file_service_proto_rawDescGZIP() []byte {
	file_api_proto_file_file_service_proto_rawDescOnce.Do(func() {
		file_api_proto_file_file_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_proto_file_file_service_proto_rawDescData)
	})
	return file_api_proto_file_file_service_proto_rawDescData
}

var file_api_proto_file_file_service_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_api_proto_file_file_service_proto_goTypes = []interface{}{
	(*FileUploadRequest)(nil),        // 0: FileUploadRequest
	(*FileUploadDetailResponse)(nil), // 1: FileUploadDetailResponse
	(*FileUploadDetail)(nil),         // 2: FileUploadDetail
	(FileUploadStatus)(0),            // 3: FileUploadStatus
}
var file_api_proto_file_file_service_proto_depIdxs = []int32{
	2, // 0: FileUploadDetailResponse.fileDetail:type_name -> FileUploadDetail
	3, // 1: FileUploadDetailResponse.status:type_name -> FileUploadStatus
	0, // 2: FileUploadPlatform.RequestFileUpload:input_type -> FileUploadRequest
	1, // 3: FileUploadPlatform.RequestFileUpload:output_type -> FileUploadDetailResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_api_proto_file_file_service_proto_init() }
func file_api_proto_file_file_service_proto_init() {
	if File_api_proto_file_file_service_proto != nil {
		return
	}
	file_api_proto_file_file_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_api_proto_file_file_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileUploadRequest); i {
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
		file_api_proto_file_file_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileUploadDetailResponse); i {
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
			RawDescriptor: file_api_proto_file_file_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_proto_file_file_service_proto_goTypes,
		DependencyIndexes: file_api_proto_file_file_service_proto_depIdxs,
		MessageInfos:      file_api_proto_file_file_service_proto_msgTypes,
	}.Build()
	File_api_proto_file_file_service_proto = out.File
	file_api_proto_file_file_service_proto_rawDesc = nil
	file_api_proto_file_file_service_proto_goTypes = nil
	file_api_proto_file_file_service_proto_depIdxs = nil
}