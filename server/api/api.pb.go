// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.25.1
// source: protos/api.proto

package api

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

type GetDevicesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetDevicesRequest) Reset() {
	*x = GetDevicesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDevicesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDevicesRequest) ProtoMessage() {}

func (x *GetDevicesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDevicesRequest.ProtoReflect.Descriptor instead.
func (*GetDevicesRequest) Descriptor() ([]byte, []int) {
	return file_protos_api_proto_rawDescGZIP(), []int{0}
}

type GetDevicesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Devices []*Device `protobuf:"bytes,1,rep,name=devices,proto3" json:"devices,omitempty"`
}

func (x *GetDevicesResponse) Reset() {
	*x = GetDevicesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDevicesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDevicesResponse) ProtoMessage() {}

func (x *GetDevicesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDevicesResponse.ProtoReflect.Descriptor instead.
func (*GetDevicesResponse) Descriptor() ([]byte, []int) {
	return file_protos_api_proto_rawDescGZIP(), []int{1}
}

func (x *GetDevicesResponse) GetDevices() []*Device {
	if x != nil {
		return x.Devices
	}
	return nil
}

type Address struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ip        string `protobuf:"bytes,1,opt,name=ip,proto3" json:"ip,omitempty"`
	Netmask   string `protobuf:"bytes,2,opt,name=netmask,proto3" json:"netmask,omitempty"`
	Broadcast string `protobuf:"bytes,3,opt,name=broadcast,proto3" json:"broadcast,omitempty"`
	P2P       string `protobuf:"bytes,4,opt,name=p2p,proto3" json:"p2p,omitempty"`
}

func (x *Address) Reset() {
	*x = Address{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Address) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Address) ProtoMessage() {}

func (x *Address) ProtoReflect() protoreflect.Message {
	mi := &file_protos_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Address.ProtoReflect.Descriptor instead.
func (*Address) Descriptor() ([]byte, []int) {
	return file_protos_api_proto_rawDescGZIP(), []int{2}
}

func (x *Address) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

func (x *Address) GetNetmask() string {
	if x != nil {
		return x.Netmask
	}
	return ""
}

func (x *Address) GetBroadcast() string {
	if x != nil {
		return x.Broadcast
	}
	return ""
}

func (x *Address) GetP2P() string {
	if x != nil {
		return x.P2P
	}
	return ""
}

type Device struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string     `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description string     `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Flags       int32      `protobuf:"varint,3,opt,name=flags,proto3" json:"flags,omitempty"`
	Addresses   []*Address `protobuf:"bytes,4,rep,name=addresses,proto3" json:"addresses,omitempty"`
}

func (x *Device) Reset() {
	*x = Device{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Device) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Device) ProtoMessage() {}

func (x *Device) ProtoReflect() protoreflect.Message {
	mi := &file_protos_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Device.ProtoReflect.Descriptor instead.
func (*Device) Descriptor() ([]byte, []int) {
	return file_protos_api_proto_rawDescGZIP(), []int{3}
}

func (x *Device) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Device) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Device) GetFlags() int32 {
	if x != nil {
		return x.Flags
	}
	return 0
}

func (x *Device) GetAddresses() []*Address {
	if x != nil {
		return x.Addresses
	}
	return nil
}

var File_protos_api_proto protoreflect.FileDescriptor

var file_protos_api_proto_rawDesc = []byte{
	0x0a, 0x10, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x03, 0x61, 0x70, 0x69, 0x22, 0x13, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x44, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x3b, 0x0a, 0x12,
	0x47, 0x65, 0x74, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x25, 0x0a, 0x07, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x52, 0x07, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x22, 0x63, 0x0a, 0x07, 0x41, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x6e, 0x65, 0x74, 0x6d, 0x61, 0x73, 0x6b, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6e, 0x65, 0x74, 0x6d, 0x61, 0x73, 0x6b, 0x12, 0x1c,
	0x0a, 0x09, 0x62, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x62, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03,
	0x70, 0x32, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x70, 0x32, 0x70, 0x22, 0x80,
	0x01, 0x0a, 0x06, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x14, 0x0a, 0x05, 0x66, 0x6c, 0x61, 0x67, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05,
	0x66, 0x6c, 0x61, 0x67, 0x73, 0x12, 0x2a, 0x0a, 0x09, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x41,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x09, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x65,
	0x73, 0x32, 0x44, 0x0a, 0x03, 0x41, 0x70, 0x69, 0x12, 0x3d, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x44,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x12, 0x16, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x74,
	0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x07, 0x5a, 0x05, 0x2e, 0x2f, 0x61, 0x70, 0x69,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_api_proto_rawDescOnce sync.Once
	file_protos_api_proto_rawDescData = file_protos_api_proto_rawDesc
)

func file_protos_api_proto_rawDescGZIP() []byte {
	file_protos_api_proto_rawDescOnce.Do(func() {
		file_protos_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_api_proto_rawDescData)
	})
	return file_protos_api_proto_rawDescData
}

var file_protos_api_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_protos_api_proto_goTypes = []interface{}{
	(*GetDevicesRequest)(nil),  // 0: api.GetDevicesRequest
	(*GetDevicesResponse)(nil), // 1: api.GetDevicesResponse
	(*Address)(nil),            // 2: api.Address
	(*Device)(nil),             // 3: api.Device
}
var file_protos_api_proto_depIdxs = []int32{
	3, // 0: api.GetDevicesResponse.devices:type_name -> api.Device
	2, // 1: api.Device.addresses:type_name -> api.Address
	0, // 2: api.Api.GetDevices:input_type -> api.GetDevicesRequest
	1, // 3: api.Api.GetDevices:output_type -> api.GetDevicesResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_protos_api_proto_init() }
func file_protos_api_proto_init() {
	if File_protos_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDevicesRequest); i {
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
		file_protos_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDevicesResponse); i {
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
		file_protos_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Address); i {
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
		file_protos_api_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Device); i {
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
			RawDescriptor: file_protos_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_api_proto_goTypes,
		DependencyIndexes: file_protos_api_proto_depIdxs,
		MessageInfos:      file_protos_api_proto_msgTypes,
	}.Build()
	File_protos_api_proto = out.File
	file_protos_api_proto_rawDesc = nil
	file_protos_api_proto_goTypes = nil
	file_protos_api_proto_depIdxs = nil
}