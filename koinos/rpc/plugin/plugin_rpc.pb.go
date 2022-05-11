// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.4
// source: koinos/rpc/plugin/plugin_rpc.proto

package plugin

import (
	rpc "github.com/koinos/koinos-proto-golang/koinos/rpc"
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

type SubmitDataRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *SubmitDataRequest) Reset() {
	*x = SubmitDataRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_koinos_rpc_plugin_plugin_rpc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubmitDataRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubmitDataRequest) ProtoMessage() {}

func (x *SubmitDataRequest) ProtoReflect() protoreflect.Message {
	mi := &file_koinos_rpc_plugin_plugin_rpc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubmitDataRequest.ProtoReflect.Descriptor instead.
func (*SubmitDataRequest) Descriptor() ([]byte, []int) {
	return file_koinos_rpc_plugin_plugin_rpc_proto_rawDescGZIP(), []int{0}
}

func (x *SubmitDataRequest) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type SubmitDataResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *SubmitDataResponse) Reset() {
	*x = SubmitDataResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_koinos_rpc_plugin_plugin_rpc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubmitDataResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubmitDataResponse) ProtoMessage() {}

func (x *SubmitDataResponse) ProtoReflect() protoreflect.Message {
	mi := &file_koinos_rpc_plugin_plugin_rpc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubmitDataResponse.ProtoReflect.Descriptor instead.
func (*SubmitDataResponse) Descriptor() ([]byte, []int) {
	return file_koinos_rpc_plugin_plugin_rpc_proto_rawDescGZIP(), []int{1}
}

func (x *SubmitDataResponse) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type PluginRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Request:
	//	*PluginRequest_Reserved
	//	*PluginRequest_SubmitData
	Request isPluginRequest_Request `protobuf_oneof:"request"`
}

func (x *PluginRequest) Reset() {
	*x = PluginRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_koinos_rpc_plugin_plugin_rpc_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PluginRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PluginRequest) ProtoMessage() {}

func (x *PluginRequest) ProtoReflect() protoreflect.Message {
	mi := &file_koinos_rpc_plugin_plugin_rpc_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PluginRequest.ProtoReflect.Descriptor instead.
func (*PluginRequest) Descriptor() ([]byte, []int) {
	return file_koinos_rpc_plugin_plugin_rpc_proto_rawDescGZIP(), []int{2}
}

func (m *PluginRequest) GetRequest() isPluginRequest_Request {
	if m != nil {
		return m.Request
	}
	return nil
}

func (x *PluginRequest) GetReserved() *rpc.ReservedRpc {
	if x, ok := x.GetRequest().(*PluginRequest_Reserved); ok {
		return x.Reserved
	}
	return nil
}

func (x *PluginRequest) GetSubmitData() *SubmitDataRequest {
	if x, ok := x.GetRequest().(*PluginRequest_SubmitData); ok {
		return x.SubmitData
	}
	return nil
}

type isPluginRequest_Request interface {
	isPluginRequest_Request()
}

type PluginRequest_Reserved struct {
	Reserved *rpc.ReservedRpc `protobuf:"bytes,1,opt,name=reserved,proto3,oneof"`
}

type PluginRequest_SubmitData struct {
	SubmitData *SubmitDataRequest `protobuf:"bytes,2,opt,name=submit_data,json=submitData,proto3,oneof"`
}

func (*PluginRequest_Reserved) isPluginRequest_Request() {}

func (*PluginRequest_SubmitData) isPluginRequest_Request() {}

type PluginResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Response:
	//	*PluginResponse_Reserved
	//	*PluginResponse_Error
	//	*PluginResponse_SubmitData
	Response isPluginResponse_Response `protobuf_oneof:"response"`
}

func (x *PluginResponse) Reset() {
	*x = PluginResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_koinos_rpc_plugin_plugin_rpc_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PluginResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PluginResponse) ProtoMessage() {}

func (x *PluginResponse) ProtoReflect() protoreflect.Message {
	mi := &file_koinos_rpc_plugin_plugin_rpc_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PluginResponse.ProtoReflect.Descriptor instead.
func (*PluginResponse) Descriptor() ([]byte, []int) {
	return file_koinos_rpc_plugin_plugin_rpc_proto_rawDescGZIP(), []int{3}
}

func (m *PluginResponse) GetResponse() isPluginResponse_Response {
	if m != nil {
		return m.Response
	}
	return nil
}

func (x *PluginResponse) GetReserved() *rpc.ReservedRpc {
	if x, ok := x.GetResponse().(*PluginResponse_Reserved); ok {
		return x.Reserved
	}
	return nil
}

func (x *PluginResponse) GetError() *rpc.ErrorResponse {
	if x, ok := x.GetResponse().(*PluginResponse_Error); ok {
		return x.Error
	}
	return nil
}

func (x *PluginResponse) GetSubmitData() *SubmitDataResponse {
	if x, ok := x.GetResponse().(*PluginResponse_SubmitData); ok {
		return x.SubmitData
	}
	return nil
}

type isPluginResponse_Response interface {
	isPluginResponse_Response()
}

type PluginResponse_Reserved struct {
	Reserved *rpc.ReservedRpc `protobuf:"bytes,1,opt,name=reserved,proto3,oneof"`
}

type PluginResponse_Error struct {
	Error *rpc.ErrorResponse `protobuf:"bytes,2,opt,name=error,proto3,oneof"`
}

type PluginResponse_SubmitData struct {
	SubmitData *SubmitDataResponse `protobuf:"bytes,3,opt,name=submit_data,json=submitData,proto3,oneof"`
}

func (*PluginResponse_Reserved) isPluginResponse_Response() {}

func (*PluginResponse_Error) isPluginResponse_Response() {}

func (*PluginResponse_SubmitData) isPluginResponse_Response() {}

var File_koinos_rpc_plugin_plugin_rpc_proto protoreflect.FileDescriptor

var file_koinos_rpc_plugin_plugin_rpc_proto_rawDesc = []byte{
	0x0a, 0x22, 0x6b, 0x6f, 0x69, 0x6e, 0x6f, 0x73, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x6c, 0x75,
	0x67, 0x69, 0x6e, 0x2f, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x5f, 0x72, 0x70, 0x63, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x6b, 0x6f, 0x69, 0x6e, 0x6f, 0x73, 0x2e, 0x72, 0x70, 0x63,
	0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x1a, 0x14, 0x6b, 0x6f, 0x69, 0x6e, 0x6f, 0x73, 0x2f,
	0x72, 0x70, 0x63, 0x2f, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x29, 0x0a,
	0x13, 0x73, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x2a, 0x0a, 0x14, 0x73, 0x75, 0x62, 0x6d,
	0x69, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x22, 0x9e, 0x01, 0x0a, 0x0e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x5f,
	0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x36, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x6b, 0x6f, 0x69, 0x6e,
	0x6f, 0x73, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x64, 0x5f,
	0x72, 0x70, 0x63, 0x48, 0x00, 0x52, 0x08, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x64, 0x12,
	0x49, 0x0a, 0x0b, 0x73, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x6b, 0x6f, 0x69, 0x6e, 0x6f, 0x73, 0x2e, 0x72, 0x70,
	0x63, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x73, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x5f,
	0x64, 0x61, 0x74, 0x61, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x00, 0x52, 0x0a,
	0x73, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x44, 0x61, 0x74, 0x61, 0x42, 0x09, 0x0a, 0x07, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0xd5, 0x01, 0x0a, 0x0f, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e,
	0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x36, 0x0a, 0x08, 0x72, 0x65, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x6b, 0x6f,
	0x69, 0x6e, 0x6f, 0x73, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x64, 0x5f, 0x72, 0x70, 0x63, 0x48, 0x00, 0x52, 0x08, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x64, 0x12, 0x32, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x6b, 0x6f, 0x69, 0x6e, 0x6f, 0x73, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x48, 0x00, 0x52, 0x05,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x4a, 0x0a, 0x0b, 0x73, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x5f,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x6b, 0x6f, 0x69,
	0x6e, 0x6f, 0x73, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x73,
	0x75, 0x62, 0x6d, 0x69, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x48, 0x00, 0x52, 0x0a, 0x73, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x44, 0x61, 0x74,
	0x61, 0x42, 0x0a, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x39, 0x5a,
	0x37, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6b, 0x6f, 0x69, 0x6e,
	0x6f, 0x73, 0x2f, 0x6b, 0x6f, 0x69, 0x6e, 0x6f, 0x73, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d,
	0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2f, 0x6b, 0x6f, 0x69, 0x6e, 0x6f, 0x73, 0x2f, 0x72, 0x70,
	0x63, 0x2f, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_koinos_rpc_plugin_plugin_rpc_proto_rawDescOnce sync.Once
	file_koinos_rpc_plugin_plugin_rpc_proto_rawDescData = file_koinos_rpc_plugin_plugin_rpc_proto_rawDesc
)

func file_koinos_rpc_plugin_plugin_rpc_proto_rawDescGZIP() []byte {
	file_koinos_rpc_plugin_plugin_rpc_proto_rawDescOnce.Do(func() {
		file_koinos_rpc_plugin_plugin_rpc_proto_rawDescData = protoimpl.X.CompressGZIP(file_koinos_rpc_plugin_plugin_rpc_proto_rawDescData)
	})
	return file_koinos_rpc_plugin_plugin_rpc_proto_rawDescData
}

var file_koinos_rpc_plugin_plugin_rpc_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_koinos_rpc_plugin_plugin_rpc_proto_goTypes = []interface{}{
	(*SubmitDataRequest)(nil),  // 0: koinos.rpc.plugin.submit_data_request
	(*SubmitDataResponse)(nil), // 1: koinos.rpc.plugin.submit_data_response
	(*PluginRequest)(nil),      // 2: koinos.rpc.plugin.plugin_request
	(*PluginResponse)(nil),     // 3: koinos.rpc.plugin.plugin_response
	(*rpc.ReservedRpc)(nil),    // 4: koinos.rpc.reserved_rpc
	(*rpc.ErrorResponse)(nil),  // 5: koinos.rpc.error_response
}
var file_koinos_rpc_plugin_plugin_rpc_proto_depIdxs = []int32{
	4, // 0: koinos.rpc.plugin.plugin_request.reserved:type_name -> koinos.rpc.reserved_rpc
	0, // 1: koinos.rpc.plugin.plugin_request.submit_data:type_name -> koinos.rpc.plugin.submit_data_request
	4, // 2: koinos.rpc.plugin.plugin_response.reserved:type_name -> koinos.rpc.reserved_rpc
	5, // 3: koinos.rpc.plugin.plugin_response.error:type_name -> koinos.rpc.error_response
	1, // 4: koinos.rpc.plugin.plugin_response.submit_data:type_name -> koinos.rpc.plugin.submit_data_response
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_koinos_rpc_plugin_plugin_rpc_proto_init() }
func file_koinos_rpc_plugin_plugin_rpc_proto_init() {
	if File_koinos_rpc_plugin_plugin_rpc_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_koinos_rpc_plugin_plugin_rpc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubmitDataRequest); i {
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
		file_koinos_rpc_plugin_plugin_rpc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubmitDataResponse); i {
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
		file_koinos_rpc_plugin_plugin_rpc_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PluginRequest); i {
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
		file_koinos_rpc_plugin_plugin_rpc_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PluginResponse); i {
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
	file_koinos_rpc_plugin_plugin_rpc_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*PluginRequest_Reserved)(nil),
		(*PluginRequest_SubmitData)(nil),
	}
	file_koinos_rpc_plugin_plugin_rpc_proto_msgTypes[3].OneofWrappers = []interface{}{
		(*PluginResponse_Reserved)(nil),
		(*PluginResponse_Error)(nil),
		(*PluginResponse_SubmitData)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_koinos_rpc_plugin_plugin_rpc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_koinos_rpc_plugin_plugin_rpc_proto_goTypes,
		DependencyIndexes: file_koinos_rpc_plugin_plugin_rpc_proto_depIdxs,
		MessageInfos:      file_koinos_rpc_plugin_plugin_rpc_proto_msgTypes,
	}.Build()
	File_koinos_rpc_plugin_plugin_rpc_proto = out.File
	file_koinos_rpc_plugin_plugin_rpc_proto_rawDesc = nil
	file_koinos_rpc_plugin_plugin_rpc_proto_goTypes = nil
	file_koinos_rpc_plugin_plugin_rpc_proto_depIdxs = nil
}
