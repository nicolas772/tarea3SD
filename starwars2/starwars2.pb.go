// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: starwars2/starwars2.proto

package go_starwars_grpc

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

type NewCity1 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NombrePlaneta string  `protobuf:"bytes,1,opt,name=nombre_planeta,json=nombrePlaneta,proto3" json:"nombre_planeta,omitempty"`
	NombreCiudad  string  `protobuf:"bytes,2,opt,name=nombre_ciudad,json=nombreCiudad,proto3" json:"nombre_ciudad,omitempty"`
	Action        string  `protobuf:"bytes,3,opt,name=action,proto3" json:"action,omitempty"`
	NuevoValor    *string `protobuf:"bytes,4,opt,name=nuevo_valor,json=nuevoValor,proto3,oneof" json:"nuevo_valor,omitempty"`
}

func (x *NewCity1) Reset() {
	*x = NewCity1{}
	if protoimpl.UnsafeEnabled {
		mi := &file_starwars2_starwars2_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewCity1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewCity1) ProtoMessage() {}

func (x *NewCity1) ProtoReflect() protoreflect.Message {
	mi := &file_starwars2_starwars2_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewCity1.ProtoReflect.Descriptor instead.
func (*NewCity1) Descriptor() ([]byte, []int) {
	return file_starwars2_starwars2_proto_rawDescGZIP(), []int{0}
}

func (x *NewCity1) GetNombrePlaneta() string {
	if x != nil {
		return x.NombrePlaneta
	}
	return ""
}

func (x *NewCity1) GetNombreCiudad() string {
	if x != nil {
		return x.NombreCiudad
	}
	return ""
}

func (x *NewCity1) GetAction() string {
	if x != nil {
		return x.Action
	}
	return ""
}

func (x *NewCity1) GetNuevoValor() string {
	if x != nil && x.NuevoValor != nil {
		return *x.NuevoValor
	}
	return ""
}

type RespFulcrum1 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RelojVector1 int32 `protobuf:"varint,1,opt,name=reloj_vector1,json=relojVector1,proto3" json:"reloj_vector1,omitempty"`
	RelojVector2 int32 `protobuf:"varint,2,opt,name=reloj_vector2,json=relojVector2,proto3" json:"reloj_vector2,omitempty"`
	RelojVector3 int32 `protobuf:"varint,3,opt,name=reloj_vector3,json=relojVector3,proto3" json:"reloj_vector3,omitempty"`
}

func (x *RespFulcrum1) Reset() {
	*x = RespFulcrum1{}
	if protoimpl.UnsafeEnabled {
		mi := &file_starwars2_starwars2_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RespFulcrum1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RespFulcrum1) ProtoMessage() {}

func (x *RespFulcrum1) ProtoReflect() protoreflect.Message {
	mi := &file_starwars2_starwars2_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RespFulcrum1.ProtoReflect.Descriptor instead.
func (*RespFulcrum1) Descriptor() ([]byte, []int) {
	return file_starwars2_starwars2_proto_rawDescGZIP(), []int{1}
}

func (x *RespFulcrum1) GetRelojVector1() int32 {
	if x != nil {
		return x.RelojVector1
	}
	return 0
}

func (x *RespFulcrum1) GetRelojVector2() int32 {
	if x != nil {
		return x.RelojVector2
	}
	return 0
}

func (x *RespFulcrum1) GetRelojVector3() int32 {
	if x != nil {
		return x.RelojVector3
	}
	return 0
}

var File_starwars2_starwars2_proto protoreflect.FileDescriptor

var file_starwars2_starwars2_proto_rawDesc = []byte{
	0x0a, 0x19, 0x73, 0x74, 0x61, 0x72, 0x77, 0x61, 0x72, 0x73, 0x32, 0x2f, 0x73, 0x74, 0x61, 0x72,
	0x77, 0x61, 0x72, 0x73, 0x32, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x73, 0x74, 0x61,
	0x72, 0x77, 0x61, 0x72, 0x73, 0x32, 0x22, 0xa4, 0x01, 0x0a, 0x08, 0x4e, 0x65, 0x77, 0x43, 0x69,
	0x74, 0x79, 0x31, 0x12, 0x25, 0x0a, 0x0e, 0x6e, 0x6f, 0x6d, 0x62, 0x72, 0x65, 0x5f, 0x70, 0x6c,
	0x61, 0x6e, 0x65, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6e, 0x6f, 0x6d,
	0x62, 0x72, 0x65, 0x50, 0x6c, 0x61, 0x6e, 0x65, 0x74, 0x61, 0x12, 0x23, 0x0a, 0x0d, 0x6e, 0x6f,
	0x6d, 0x62, 0x72, 0x65, 0x5f, 0x63, 0x69, 0x75, 0x64, 0x61, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0c, 0x6e, 0x6f, 0x6d, 0x62, 0x72, 0x65, 0x43, 0x69, 0x75, 0x64, 0x61, 0x64, 0x12,
	0x16, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x24, 0x0a, 0x0b, 0x6e, 0x75, 0x65, 0x76, 0x6f,
	0x5f, 0x76, 0x61, 0x6c, 0x6f, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0a,
	0x6e, 0x75, 0x65, 0x76, 0x6f, 0x56, 0x61, 0x6c, 0x6f, 0x72, 0x88, 0x01, 0x01, 0x42, 0x0e, 0x0a,
	0x0c, 0x5f, 0x6e, 0x75, 0x65, 0x76, 0x6f, 0x5f, 0x76, 0x61, 0x6c, 0x6f, 0x72, 0x22, 0x7d, 0x0a,
	0x0c, 0x52, 0x65, 0x73, 0x70, 0x46, 0x75, 0x6c, 0x63, 0x72, 0x75, 0x6d, 0x31, 0x12, 0x23, 0x0a,
	0x0d, 0x72, 0x65, 0x6c, 0x6f, 0x6a, 0x5f, 0x76, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x31, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x72, 0x65, 0x6c, 0x6f, 0x6a, 0x56, 0x65, 0x63, 0x74, 0x6f,
	0x72, 0x31, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x6c, 0x6f, 0x6a, 0x5f, 0x76, 0x65, 0x63, 0x74,
	0x6f, 0x72, 0x32, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x72, 0x65, 0x6c, 0x6f, 0x6a,
	0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x32, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x6c, 0x6f, 0x6a,
	0x5f, 0x76, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x33, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c,
	0x72, 0x65, 0x6c, 0x6f, 0x6a, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x33, 0x32, 0x4e, 0x0a, 0x09,
	0x53, 0x74, 0x61, 0x72, 0x57, 0x61, 0x72, 0x73, 0x31, 0x12, 0x41, 0x0a, 0x0f, 0x43, 0x69, 0x74,
	0x79, 0x4d, 0x67, 0x6d, 0x74, 0x46, 0x75, 0x6c, 0x63, 0x72, 0x75, 0x6d, 0x12, 0x13, 0x2e, 0x73,
	0x74, 0x61, 0x72, 0x77, 0x61, 0x72, 0x73, 0x32, 0x2e, 0x4e, 0x65, 0x77, 0x43, 0x69, 0x74, 0x79,
	0x31, 0x1a, 0x17, 0x2e, 0x73, 0x74, 0x61, 0x72, 0x77, 0x61, 0x72, 0x73, 0x32, 0x2e, 0x52, 0x65,
	0x73, 0x70, 0x46, 0x75, 0x6c, 0x63, 0x72, 0x75, 0x6d, 0x31, 0x22, 0x00, 0x42, 0x2f, 0x5a, 0x2d,
	0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2d, 0x73,
	0x74, 0x61, 0x72, 0x77, 0x61, 0x72, 0x73, 0x2d, 0x67, 0x72, 0x70, 0x63, 0x3b, 0x67, 0x6f, 0x5f,
	0x73, 0x74, 0x61, 0x72, 0x77, 0x61, 0x72, 0x73, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_starwars2_starwars2_proto_rawDescOnce sync.Once
	file_starwars2_starwars2_proto_rawDescData = file_starwars2_starwars2_proto_rawDesc
)

func file_starwars2_starwars2_proto_rawDescGZIP() []byte {
	file_starwars2_starwars2_proto_rawDescOnce.Do(func() {
		file_starwars2_starwars2_proto_rawDescData = protoimpl.X.CompressGZIP(file_starwars2_starwars2_proto_rawDescData)
	})
	return file_starwars2_starwars2_proto_rawDescData
}

var file_starwars2_starwars2_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_starwars2_starwars2_proto_goTypes = []interface{}{
	(*NewCity1)(nil),     // 0: starwars2.NewCity1
	(*RespFulcrum1)(nil), // 1: starwars2.RespFulcrum1
}
var file_starwars2_starwars2_proto_depIdxs = []int32{
	0, // 0: starwars2.StarWars1.CityMgmtFulcrum:input_type -> starwars2.NewCity1
	1, // 1: starwars2.StarWars1.CityMgmtFulcrum:output_type -> starwars2.RespFulcrum1
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_starwars2_starwars2_proto_init() }
func file_starwars2_starwars2_proto_init() {
	if File_starwars2_starwars2_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_starwars2_starwars2_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewCity1); i {
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
		file_starwars2_starwars2_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RespFulcrum1); i {
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
	file_starwars2_starwars2_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_starwars2_starwars2_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_starwars2_starwars2_proto_goTypes,
		DependencyIndexes: file_starwars2_starwars2_proto_depIdxs,
		MessageInfos:      file_starwars2_starwars2_proto_msgTypes,
	}.Build()
	File_starwars2_starwars2_proto = out.File
	file_starwars2_starwars2_proto_rawDesc = nil
	file_starwars2_starwars2_proto_goTypes = nil
	file_starwars2_starwars2_proto_depIdxs = nil
}