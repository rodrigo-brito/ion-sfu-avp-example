// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.12.4
// source: cmd/signal/grpc/proto/avp.proto

package proto

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type SignalRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Payload:
	//	*SignalRequest_Process
	Payload isSignalRequest_Payload `protobuf_oneof:"payload"`
}

func (x *SignalRequest) Reset() {
	*x = SignalRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmd_server_grpc_proto_avp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignalRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignalRequest) ProtoMessage() {}

func (x *SignalRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_server_grpc_proto_avp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignalRequest.ProtoReflect.Descriptor instead.
func (*SignalRequest) Descriptor() ([]byte, []int) {
	return file_cmd_server_grpc_proto_avp_proto_rawDescGZIP(), []int{0}
}

func (m *SignalRequest) GetPayload() isSignalRequest_Payload {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (x *SignalRequest) GetProcess() *Process {
	if x, ok := x.GetPayload().(*SignalRequest_Process); ok {
		return x.Process
	}
	return nil
}

type isSignalRequest_Payload interface {
	isSignalRequest_Payload()
}

type SignalRequest_Process struct {
	Process *Process `protobuf:"bytes,1,opt,name=process,proto3,oneof"`
}

func (*SignalRequest_Process) isSignalRequest_Payload() {}

type SignalReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SignalReply) Reset() {
	*x = SignalReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmd_server_grpc_proto_avp_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignalReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignalReply) ProtoMessage() {}

func (x *SignalReply) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_server_grpc_proto_avp_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignalReply.ProtoReflect.Descriptor instead.
func (*SignalReply) Descriptor() ([]byte, []int) {
	return file_cmd_server_grpc_proto_avp_proto_rawDescGZIP(), []int{1}
}

// Process describes an a/v process
type Process struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sfu    string `protobuf:"bytes,1,opt,name=sfu,proto3" json:"sfu,omitempty"` // media sfu
	Pid    string `protobuf:"bytes,2,opt,name=pid,proto3" json:"pid,omitempty"` // pipeline id
	Sid    string `protobuf:"bytes,3,opt,name=sid,proto3" json:"sid,omitempty"` // session id
	Tid    string `protobuf:"bytes,4,opt,name=tid,proto3" json:"tid,omitempty"` // track id
	Eid    string `protobuf:"bytes,5,opt,name=eid,proto3" json:"eid,omitempty"` // element id
	Config []byte `protobuf:"bytes,6,opt,name=config,proto3" json:"config,omitempty"`
}

func (x *Process) Reset() {
	*x = Process{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmd_server_grpc_proto_avp_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Process) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Process) ProtoMessage() {}

func (x *Process) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_server_grpc_proto_avp_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Process.ProtoReflect.Descriptor instead.
func (*Process) Descriptor() ([]byte, []int) {
	return file_cmd_server_grpc_proto_avp_proto_rawDescGZIP(), []int{2}
}

func (x *Process) GetSfu() string {
	if x != nil {
		return x.Sfu
	}
	return ""
}

func (x *Process) GetPid() string {
	if x != nil {
		return x.Pid
	}
	return ""
}

func (x *Process) GetSid() string {
	if x != nil {
		return x.Sid
	}
	return ""
}

func (x *Process) GetTid() string {
	if x != nil {
		return x.Tid
	}
	return ""
}

func (x *Process) GetEid() string {
	if x != nil {
		return x.Eid
	}
	return ""
}

func (x *Process) GetConfig() []byte {
	if x != nil {
		return x.Config
	}
	return nil
}

var File_cmd_server_grpc_proto_avp_proto protoreflect.FileDescriptor

var file_cmd_server_grpc_proto_avp_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x63, 0x6d, 0x64, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x67, 0x72, 0x70,
	0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x76, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x03, 0x61, 0x76, 0x70, 0x22, 0x44, 0x0a, 0x0d, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x6c,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x28, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x63, 0x65,
	0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x61, 0x76, 0x70, 0x2e, 0x50,
	0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x48, 0x00, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73,
	0x73, 0x42, 0x09, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x22, 0x0d, 0x0a, 0x0b,
	0x53, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x7b, 0x0a, 0x07, 0x50,
	0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x66, 0x75, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x73, 0x66, 0x75, 0x12, 0x10, 0x0a, 0x03, 0x70, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x70, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x69,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x73, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03,
	0x74, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x74, 0x69, 0x64, 0x12, 0x10,
	0x0a, 0x03, 0x65, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x65, 0x69, 0x64,
	0x12, 0x16, 0x0a, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x32, 0x3b, 0x0a, 0x03, 0x41, 0x56, 0x50, 0x12,
	0x34, 0x0a, 0x06, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x12, 0x12, 0x2e, 0x61, 0x76, 0x70, 0x2e,
	0x53, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e,
	0x61, 0x76, 0x70, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22,
	0x00, 0x28, 0x01, 0x30, 0x01, 0x42, 0x2f, 0x5a, 0x2d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x69, 0x6f, 0x6e, 0x2f, 0x69, 0x6f, 0x6e, 0x2d, 0x61, 0x76, 0x70,
	0x2f, 0x63, 0x6d, 0x64, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x67, 0x72, 0x70, 0x63,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cmd_server_grpc_proto_avp_proto_rawDescOnce sync.Once
	file_cmd_server_grpc_proto_avp_proto_rawDescData = file_cmd_server_grpc_proto_avp_proto_rawDesc
)

func file_cmd_server_grpc_proto_avp_proto_rawDescGZIP() []byte {
	file_cmd_server_grpc_proto_avp_proto_rawDescOnce.Do(func() {
		file_cmd_server_grpc_proto_avp_proto_rawDescData = protoimpl.X.CompressGZIP(file_cmd_server_grpc_proto_avp_proto_rawDescData)
	})
	return file_cmd_server_grpc_proto_avp_proto_rawDescData
}

var file_cmd_server_grpc_proto_avp_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_cmd_server_grpc_proto_avp_proto_goTypes = []interface{}{
	(*SignalRequest)(nil), // 0: avp.SignalRequest
	(*SignalReply)(nil),   // 1: avp.SignalReply
	(*Process)(nil),       // 2: avp.Process
}
var file_cmd_server_grpc_proto_avp_proto_depIdxs = []int32{
	2, // 0: avp.SignalRequest.process:type_name -> avp.Process
	0, // 1: avp.AVP.Signal:input_type -> avp.SignalRequest
	1, // 2: avp.AVP.Signal:output_type -> avp.SignalReply
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_cmd_server_grpc_proto_avp_proto_init() }
func file_cmd_server_grpc_proto_avp_proto_init() {
	if File_cmd_server_grpc_proto_avp_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cmd_server_grpc_proto_avp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignalRequest); i {
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
		file_cmd_server_grpc_proto_avp_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignalReply); i {
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
		file_cmd_server_grpc_proto_avp_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Process); i {
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
	file_cmd_server_grpc_proto_avp_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*SignalRequest_Process)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_cmd_server_grpc_proto_avp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_cmd_server_grpc_proto_avp_proto_goTypes,
		DependencyIndexes: file_cmd_server_grpc_proto_avp_proto_depIdxs,
		MessageInfos:      file_cmd_server_grpc_proto_avp_proto_msgTypes,
	}.Build()
	File_cmd_server_grpc_proto_avp_proto = out.File
	file_cmd_server_grpc_proto_avp_proto_rawDesc = nil
	file_cmd_server_grpc_proto_avp_proto_goTypes = nil
	file_cmd_server_grpc_proto_avp_proto_depIdxs = nil
}
