// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.1
// source: api/proto/movie/movie.proto

package movie

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

type MovieStatus int32

const (
	MovieStatus_CREATED         MovieStatus = 0
	MovieStatus_UPDATED         MovieStatus = 1
	MovieStatus_FAILED          MovieStatus = 2
	MovieStatus_DUPLICATE_ENTRY MovieStatus = 3
)

// Enum value maps for MovieStatus.
var (
	MovieStatus_name = map[int32]string{
		0: "CREATED",
		1: "UPDATED",
		2: "FAILED",
		3: "DUPLICATE_ENTRY",
	}
	MovieStatus_value = map[string]int32{
		"CREATED":         0,
		"UPDATED":         1,
		"FAILED":          2,
		"DUPLICATE_ENTRY": 3,
	}
)

func (x MovieStatus) Enum() *MovieStatus {
	p := new(MovieStatus)
	*p = x
	return p
}

func (x MovieStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MovieStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_api_proto_movie_movie_proto_enumTypes[0].Descriptor()
}

func (MovieStatus) Type() protoreflect.EnumType {
	return &file_api_proto_movie_movie_proto_enumTypes[0]
}

func (x MovieStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MovieStatus.Descriptor instead.
func (MovieStatus) EnumDescriptor() ([]byte, []int) {
	return file_api_proto_movie_movie_proto_rawDescGZIP(), []int{0}
}

type MovieDetails struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	MovieName   string `protobuf:"bytes,2,opt,name=movie_name,json=movieName,proto3" json:"movie_name,omitempty"`
	Genre       string `protobuf:"bytes,3,opt,name=genre,proto3" json:"genre,omitempty"`
	Description string `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	Ratings     string `protobuf:"bytes,5,opt,name=ratings,proto3" json:"ratings,omitempty"`
}

func (x *MovieDetails) Reset() {
	*x = MovieDetails{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_movie_movie_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MovieDetails) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MovieDetails) ProtoMessage() {}

func (x *MovieDetails) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_movie_movie_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MovieDetails.ProtoReflect.Descriptor instead.
func (*MovieDetails) Descriptor() ([]byte, []int) {
	return file_api_proto_movie_movie_proto_rawDescGZIP(), []int{0}
}

func (x *MovieDetails) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *MovieDetails) GetMovieName() string {
	if x != nil {
		return x.MovieName
	}
	return ""
}

func (x *MovieDetails) GetGenre() string {
	if x != nil {
		return x.Genre
	}
	return ""
}

func (x *MovieDetails) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *MovieDetails) GetRatings() string {
	if x != nil {
		return x.Ratings
	}
	return ""
}

var File_api_proto_movie_movie_proto protoreflect.FileDescriptor

var file_api_proto_movie_movie_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x6f, 0x76, 0x69,
	0x65, 0x2f, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8f, 0x01,
	0x0a, 0x0c, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1d,
	0x0a, 0x0a, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x67, 0x65, 0x6e, 0x72, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x67, 0x65,
	0x6e, 0x72, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x73,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x2a,
	0x48, 0x0a, 0x0b, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0b,
	0x0a, 0x07, 0x43, 0x52, 0x45, 0x41, 0x54, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x55,
	0x50, 0x44, 0x41, 0x54, 0x45, 0x44, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x46, 0x41, 0x49, 0x4c,
	0x45, 0x44, 0x10, 0x02, 0x12, 0x13, 0x0a, 0x0f, 0x44, 0x55, 0x50, 0x4c, 0x49, 0x43, 0x41, 0x54,
	0x45, 0x5f, 0x45, 0x4e, 0x54, 0x52, 0x59, 0x10, 0x03, 0x42, 0x08, 0x5a, 0x06, 0x6d, 0x6f, 0x76,
	0x69, 0x65, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_proto_movie_movie_proto_rawDescOnce sync.Once
	file_api_proto_movie_movie_proto_rawDescData = file_api_proto_movie_movie_proto_rawDesc
)

func file_api_proto_movie_movie_proto_rawDescGZIP() []byte {
	file_api_proto_movie_movie_proto_rawDescOnce.Do(func() {
		file_api_proto_movie_movie_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_proto_movie_movie_proto_rawDescData)
	})
	return file_api_proto_movie_movie_proto_rawDescData
}

var file_api_proto_movie_movie_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_api_proto_movie_movie_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_api_proto_movie_movie_proto_goTypes = []interface{}{
	(MovieStatus)(0),     // 0: MovieStatus
	(*MovieDetails)(nil), // 1: MovieDetails
}
var file_api_proto_movie_movie_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_proto_movie_movie_proto_init() }
func file_api_proto_movie_movie_proto_init() {
	if File_api_proto_movie_movie_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_proto_movie_movie_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MovieDetails); i {
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
			RawDescriptor: file_api_proto_movie_movie_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_proto_movie_movie_proto_goTypes,
		DependencyIndexes: file_api_proto_movie_movie_proto_depIdxs,
		EnumInfos:         file_api_proto_movie_movie_proto_enumTypes,
		MessageInfos:      file_api_proto_movie_movie_proto_msgTypes,
	}.Build()
	File_api_proto_movie_movie_proto = out.File
	file_api_proto_movie_movie_proto_rawDesc = nil
	file_api_proto_movie_movie_proto_goTypes = nil
	file_api_proto_movie_movie_proto_depIdxs = nil
}
