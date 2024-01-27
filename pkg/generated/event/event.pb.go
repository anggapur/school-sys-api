// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.1
// source: api/proto/event/event.proto

package event

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

type EventStatus int32

const (
	EventStatus_EVENT_CREATED         EventStatus = 0
	EventStatus_EVENT_UPDATED         EventStatus = 1
	EventStatus_EVENT_FAILED          EventStatus = 2
	EventStatus_EVENT_DUPLICATE_ENTRY EventStatus = 3
)

// Enum value maps for EventStatus.
var (
	EventStatus_name = map[int32]string{
		0: "EVENT_CREATED",
		1: "EVENT_UPDATED",
		2: "EVENT_FAILED",
		3: "EVENT_DUPLICATE_ENTRY",
	}
	EventStatus_value = map[string]int32{
		"EVENT_CREATED":         0,
		"EVENT_UPDATED":         1,
		"EVENT_FAILED":          2,
		"EVENT_DUPLICATE_ENTRY": 3,
	}
)

func (x EventStatus) Enum() *EventStatus {
	p := new(EventStatus)
	*p = x
	return p
}

func (x EventStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EventStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_api_proto_event_event_proto_enumTypes[0].Descriptor()
}

func (EventStatus) Type() protoreflect.EnumType {
	return &file_api_proto_event_event_proto_enumTypes[0]
}

func (x EventStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EventStatus.Descriptor instead.
func (EventStatus) EnumDescriptor() ([]byte, []int) {
	return file_api_proto_event_event_proto_rawDescGZIP(), []int{0}
}

type EventDetails struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	EventName   string `protobuf:"bytes,2,opt,name=event_name,json=eventName,proto3" json:"event_name,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *EventDetails) Reset() {
	*x = EventDetails{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_event_event_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventDetails) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventDetails) ProtoMessage() {}

func (x *EventDetails) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_event_event_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventDetails.ProtoReflect.Descriptor instead.
func (*EventDetails) Descriptor() ([]byte, []int) {
	return file_api_proto_event_event_proto_rawDescGZIP(), []int{0}
}

func (x *EventDetails) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *EventDetails) GetEventName() string {
	if x != nil {
		return x.EventName
	}
	return ""
}

func (x *EventDetails) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

var File_api_proto_event_event_proto protoreflect.FileDescriptor

var file_api_proto_event_event_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5f, 0x0a,
	0x0c, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1d, 0x0a,
	0x0a, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x2a, 0x60,
	0x0a, 0x0b, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x11, 0x0a,
	0x0d, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x43, 0x52, 0x45, 0x41, 0x54, 0x45, 0x44, 0x10, 0x00,
	0x12, 0x11, 0x0a, 0x0d, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x55, 0x50, 0x44, 0x41, 0x54, 0x45,
	0x44, 0x10, 0x01, 0x12, 0x10, 0x0a, 0x0c, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x46, 0x41, 0x49,
	0x4c, 0x45, 0x44, 0x10, 0x02, 0x12, 0x19, 0x0a, 0x15, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x44,
	0x55, 0x50, 0x4c, 0x49, 0x43, 0x41, 0x54, 0x45, 0x5f, 0x45, 0x4e, 0x54, 0x52, 0x59, 0x10, 0x03,
	0x42, 0x08, 0x5a, 0x06, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_api_proto_event_event_proto_rawDescOnce sync.Once
	file_api_proto_event_event_proto_rawDescData = file_api_proto_event_event_proto_rawDesc
)

func file_api_proto_event_event_proto_rawDescGZIP() []byte {
	file_api_proto_event_event_proto_rawDescOnce.Do(func() {
		file_api_proto_event_event_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_proto_event_event_proto_rawDescData)
	})
	return file_api_proto_event_event_proto_rawDescData
}

var file_api_proto_event_event_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_api_proto_event_event_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_api_proto_event_event_proto_goTypes = []interface{}{
	(EventStatus)(0),     // 0: EventStatus
	(*EventDetails)(nil), // 1: EventDetails
}
var file_api_proto_event_event_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_proto_event_event_proto_init() }
func file_api_proto_event_event_proto_init() {
	if File_api_proto_event_event_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_proto_event_event_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EventDetails); i {
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
			RawDescriptor: file_api_proto_event_event_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_proto_event_event_proto_goTypes,
		DependencyIndexes: file_api_proto_event_event_proto_depIdxs,
		EnumInfos:         file_api_proto_event_event_proto_enumTypes,
		MessageInfos:      file_api_proto_event_event_proto_msgTypes,
	}.Build()
	File_api_proto_event_event_proto = out.File
	file_api_proto_event_event_proto_rawDesc = nil
	file_api_proto_event_event_proto_goTypes = nil
	file_api_proto_event_event_proto_depIdxs = nil
}