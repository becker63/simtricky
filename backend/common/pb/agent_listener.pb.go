// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.4
// 	protoc        v5.29.3
// source: agent_listener.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Status struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Status        bool                   `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Status) Reset() {
	*x = Status{}
	mi := &file_agent_listener_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Status) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Status) ProtoMessage() {}

func (x *Status) ProtoReflect() protoreflect.Message {
	mi := &file_agent_listener_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Status.ProtoReflect.Descriptor instead.
func (*Status) Descriptor() ([]byte, []int) {
	return file_agent_listener_proto_rawDescGZIP(), []int{0}
}

func (x *Status) GetStatus() bool {
	if x != nil {
		return x.Status
	}
	return false
}

type StartStopReq struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	MachineId     uint32                 `protobuf:"varint,1,opt,name=Machine_id,json=MachineId,proto3" json:"Machine_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *StartStopReq) Reset() {
	*x = StartStopReq{}
	mi := &file_agent_listener_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *StartStopReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartStopReq) ProtoMessage() {}

func (x *StartStopReq) ProtoReflect() protoreflect.Message {
	mi := &file_agent_listener_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartStopReq.ProtoReflect.Descriptor instead.
func (*StartStopReq) Descriptor() ([]byte, []int) {
	return file_agent_listener_proto_rawDescGZIP(), []int{1}
}

func (x *StartStopReq) GetMachineId() uint32 {
	if x != nil {
		return x.MachineId
	}
	return 0
}

type Config struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	MachineId     uint32                 `protobuf:"varint,1,opt,name=Machine_id,json=MachineId,proto3" json:"Machine_id,omitempty"`
	Dockerfile    string                 `protobuf:"bytes,2,opt,name=Dockerfile,proto3" json:"Dockerfile,omitempty"`
	Files         *Dir                   `protobuf:"bytes,3,opt,name=Files,proto3,oneof" json:"Files,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Config) Reset() {
	*x = Config{}
	mi := &file_agent_listener_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Config) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Config) ProtoMessage() {}

func (x *Config) ProtoReflect() protoreflect.Message {
	mi := &file_agent_listener_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Config.ProtoReflect.Descriptor instead.
func (*Config) Descriptor() ([]byte, []int) {
	return file_agent_listener_proto_rawDescGZIP(), []int{2}
}

func (x *Config) GetMachineId() uint32 {
	if x != nil {
		return x.MachineId
	}
	return 0
}

func (x *Config) GetDockerfile() string {
	if x != nil {
		return x.Dockerfile
	}
	return ""
}

func (x *Config) GetFiles() *Dir {
	if x != nil {
		return x.Files
	}
	return nil
}

type ChangedConfig struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	MachineId     uint32                 `protobuf:"varint,1,opt,name=machine_id,json=machineId,proto3" json:"machine_id,omitempty"`
	Config        *Config                `protobuf:"bytes,2,opt,name=config,proto3" json:"config,omitempty"`                           // The updated values
	UpdateMask    *fieldmaskpb.FieldMask `protobuf:"bytes,3,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"` // Specifies which fields to update
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ChangedConfig) Reset() {
	*x = ChangedConfig{}
	mi := &file_agent_listener_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ChangedConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChangedConfig) ProtoMessage() {}

func (x *ChangedConfig) ProtoReflect() protoreflect.Message {
	mi := &file_agent_listener_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChangedConfig.ProtoReflect.Descriptor instead.
func (*ChangedConfig) Descriptor() ([]byte, []int) {
	return file_agent_listener_proto_rawDescGZIP(), []int{3}
}

func (x *ChangedConfig) GetMachineId() uint32 {
	if x != nil {
		return x.MachineId
	}
	return 0
}

func (x *ChangedConfig) GetConfig() *Config {
	if x != nil {
		return x.Config
	}
	return nil
}

func (x *ChangedConfig) GetUpdateMask() *fieldmaskpb.FieldMask {
	if x != nil {
		return x.UpdateMask
	}
	return nil
}

var File_agent_listener_proto protoreflect.FileDescriptor

var file_agent_listener_proto_rawDesc = string([]byte{
	0x0a, 0x14, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x65, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x10, 0x66, 0x69, 0x6c, 0x65,
	0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x66, 0x69,
	0x65, 0x6c, 0x64, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x20,
	0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x22, 0x2e, 0x0a, 0x0d, 0x53, 0x74, 0x61, 0x72, 0x74, 0x53, 0x74, 0x6f, 0x70, 0x5f, 0x72, 0x65,
	0x71, 0x12, 0x1d, 0x0a, 0x0a, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x49, 0x64,
	0x22, 0x75, 0x0a, 0x06, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x1d, 0x0a, 0x0a, 0x4d, 0x61,
	0x63, 0x68, 0x69, 0x6e, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09,
	0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x44, 0x6f, 0x63,
	0x6b, 0x65, 0x72, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x44,
	0x6f, 0x63, 0x6b, 0x65, 0x72, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x22, 0x0a, 0x05, 0x46, 0x69, 0x6c,
	0x65, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x70, 0x62, 0x2e, 0x44, 0x69,
	0x72, 0x48, 0x00, 0x52, 0x05, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x88, 0x01, 0x01, 0x42, 0x08, 0x0a,
	0x06, 0x5f, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x22, 0x8f, 0x01, 0x0a, 0x0d, 0x43, 0x68, 0x61, 0x6e,
	0x67, 0x65, 0x64, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x1d, 0x0a, 0x0a, 0x6d, 0x61, 0x63,
	0x68, 0x69, 0x6e, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x6d,
	0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x49, 0x64, 0x12, 0x22, 0x0a, 0x06, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x52, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x3b, 0x0a, 0x0b,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4d, 0x61, 0x73, 0x6b, 0x52, 0x0a, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x73, 0x6b, 0x32, 0xb7, 0x01, 0x0a, 0x0e, 0x41, 0x67,
	0x65, 0x6e, 0x74, 0x5f, 0x4c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x12, 0x26, 0x0a, 0x05,
	0x53, 0x74, 0x61, 0x72, 0x74, 0x12, 0x11, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x74, 0x61, 0x72, 0x74,
	0x53, 0x74, 0x6f, 0x70, 0x5f, 0x72, 0x65, 0x71, 0x1a, 0x0a, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x25, 0x0a, 0x04, 0x53, 0x74, 0x6f, 0x70, 0x12, 0x11, 0x2e, 0x70,
	0x62, 0x2e, 0x53, 0x74, 0x61, 0x72, 0x74, 0x53, 0x74, 0x6f, 0x70, 0x5f, 0x72, 0x65, 0x71, 0x1a,
	0x0a, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x2d, 0x0a, 0x0c, 0x43,
	0x68, 0x61, 0x6e, 0x67, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x11, 0x2e, 0x70, 0x62,
	0x2e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x64, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x1a, 0x0a,
	0x2e, 0x70, 0x62, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x27, 0x0a, 0x0d, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x5f, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x0a, 0x2e, 0x70, 0x62,
	0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x1a, 0x0a, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x42, 0x29, 0x5a, 0x27, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x62, 0x65, 0x63, 0x6b, 0x65, 0x72, 0x36, 0x33, 0x2f, 0x73, 0x69, 0x6d, 0x74, 0x72,
	0x69, 0x63, 0x6b, 0x79, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x70, 0x62, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_agent_listener_proto_rawDescOnce sync.Once
	file_agent_listener_proto_rawDescData []byte
)

func file_agent_listener_proto_rawDescGZIP() []byte {
	file_agent_listener_proto_rawDescOnce.Do(func() {
		file_agent_listener_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_agent_listener_proto_rawDesc), len(file_agent_listener_proto_rawDesc)))
	})
	return file_agent_listener_proto_rawDescData
}

var file_agent_listener_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_agent_listener_proto_goTypes = []any{
	(*Status)(nil),                // 0: pb.Status
	(*StartStopReq)(nil),          // 1: pb.StartStop_req
	(*Config)(nil),                // 2: pb.Config
	(*ChangedConfig)(nil),         // 3: pb.ChangedConfig
	(*Dir)(nil),                   // 4: pb.Dir
	(*fieldmaskpb.FieldMask)(nil), // 5: google.protobuf.FieldMask
}
var file_agent_listener_proto_depIdxs = []int32{
	4, // 0: pb.Config.Files:type_name -> pb.Dir
	2, // 1: pb.ChangedConfig.config:type_name -> pb.Config
	5, // 2: pb.ChangedConfig.update_mask:type_name -> google.protobuf.FieldMask
	1, // 3: pb.Agent_Listener.Start:input_type -> pb.StartStop_req
	1, // 4: pb.Agent_Listener.Stop:input_type -> pb.StartStop_req
	3, // 5: pb.Agent_Listener.ChangeConfig:input_type -> pb.ChangedConfig
	2, // 6: pb.Agent_Listener.Create_Server:input_type -> pb.Config
	0, // 7: pb.Agent_Listener.Start:output_type -> pb.Status
	0, // 8: pb.Agent_Listener.Stop:output_type -> pb.Status
	0, // 9: pb.Agent_Listener.ChangeConfig:output_type -> pb.Status
	0, // 10: pb.Agent_Listener.Create_Server:output_type -> pb.Status
	7, // [7:11] is the sub-list for method output_type
	3, // [3:7] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_agent_listener_proto_init() }
func file_agent_listener_proto_init() {
	if File_agent_listener_proto != nil {
		return
	}
	file_filesystem_proto_init()
	file_agent_listener_proto_msgTypes[2].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_agent_listener_proto_rawDesc), len(file_agent_listener_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_agent_listener_proto_goTypes,
		DependencyIndexes: file_agent_listener_proto_depIdxs,
		MessageInfos:      file_agent_listener_proto_msgTypes,
	}.Build()
	File_agent_listener_proto = out.File
	file_agent_listener_proto_goTypes = nil
	file_agent_listener_proto_depIdxs = nil
}
