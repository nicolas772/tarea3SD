// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: starwars/starwars.proto

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

type NewCity struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NombrePlaneta string  `protobuf:"bytes,1,opt,name=nombre_planeta,json=nombrePlaneta,proto3" json:"nombre_planeta,omitempty"`
	NombreCiudad  string  `protobuf:"bytes,2,opt,name=nombre_ciudad,json=nombreCiudad,proto3" json:"nombre_ciudad,omitempty"`
	Action        string  `protobuf:"bytes,3,opt,name=action,proto3" json:"action,omitempty"`
	NuevoValor    *string `protobuf:"bytes,4,opt,name=nuevo_valor,json=nuevoValor,proto3,oneof" json:"nuevo_valor,omitempty"`
}

func (x *NewCity) Reset() {
	*x = NewCity{}
	if protoimpl.UnsafeEnabled {
		mi := &file_starwars_starwars_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewCity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewCity) ProtoMessage() {}

func (x *NewCity) ProtoReflect() protoreflect.Message {
	mi := &file_starwars_starwars_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewCity.ProtoReflect.Descriptor instead.
func (*NewCity) Descriptor() ([]byte, []int) {
	return file_starwars_starwars_proto_rawDescGZIP(), []int{0}
}

func (x *NewCity) GetNombrePlaneta() string {
	if x != nil {
		return x.NombrePlaneta
	}
	return ""
}

func (x *NewCity) GetNombreCiudad() string {
	if x != nil {
		return x.NombreCiudad
	}
	return ""
}

func (x *NewCity) GetAction() string {
	if x != nil {
		return x.Action
	}
	return ""
}

func (x *NewCity) GetNuevoValor() string {
	if x != nil && x.NuevoValor != nil {
		return *x.NuevoValor
	}
	return ""
}

type RespBroker1 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DireccionServidor string `protobuf:"bytes,1,opt,name=direccion_servidor,json=direccionServidor,proto3" json:"direccion_servidor,omitempty"`
}

func (x *RespBroker1) Reset() {
	*x = RespBroker1{}
	if protoimpl.UnsafeEnabled {
		mi := &file_starwars_starwars_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RespBroker1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RespBroker1) ProtoMessage() {}

func (x *RespBroker1) ProtoReflect() protoreflect.Message {
	mi := &file_starwars_starwars_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RespBroker1.ProtoReflect.Descriptor instead.
func (*RespBroker1) Descriptor() ([]byte, []int) {
	return file_starwars_starwars_proto_rawDescGZIP(), []int{1}
}

func (x *RespBroker1) GetDireccionServidor() string {
	if x != nil {
		return x.DireccionServidor
	}
	return ""
}

type RespBroker2 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CantRebeldes       int32   `protobuf:"varint,1,opt,name=cant_rebeldes,json=cantRebeldes,proto3" json:"cant_rebeldes,omitempty"`
	RelojVector        []int32 `protobuf:"varint,2,rep,packed,name=reloj_vector,json=relojVector,proto3" json:"reloj_vector,omitempty"`
	ServidorContactado string  `protobuf:"bytes,3,opt,name=servidor_contactado,json=servidorContactado,proto3" json:"servidor_contactado,omitempty"`
}

func (x *RespBroker2) Reset() {
	*x = RespBroker2{}
	if protoimpl.UnsafeEnabled {
		mi := &file_starwars_starwars_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RespBroker2) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RespBroker2) ProtoMessage() {}

func (x *RespBroker2) ProtoReflect() protoreflect.Message {
	mi := &file_starwars_starwars_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RespBroker2.ProtoReflect.Descriptor instead.
func (*RespBroker2) Descriptor() ([]byte, []int) {
	return file_starwars_starwars_proto_rawDescGZIP(), []int{2}
}

func (x *RespBroker2) GetCantRebeldes() int32 {
	if x != nil {
		return x.CantRebeldes
	}
	return 0
}

func (x *RespBroker2) GetRelojVector() []int32 {
	if x != nil {
		return x.RelojVector
	}
	return nil
}

func (x *RespBroker2) GetServidorContactado() string {
	if x != nil {
		return x.ServidorContactado
	}
	return ""
}

var File_starwars_starwars_proto protoreflect.FileDescriptor

var file_starwars_starwars_proto_rawDesc = []byte{
	0x0a, 0x17, 0x73, 0x74, 0x61, 0x72, 0x77, 0x61, 0x72, 0x73, 0x2f, 0x73, 0x74, 0x61, 0x72, 0x77,
	0x61, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73, 0x74, 0x61, 0x72, 0x77,
	0x61, 0x72, 0x73, 0x22, 0xa3, 0x01, 0x0a, 0x07, 0x4e, 0x65, 0x77, 0x43, 0x69, 0x74, 0x79, 0x12,
	0x25, 0x0a, 0x0e, 0x6e, 0x6f, 0x6d, 0x62, 0x72, 0x65, 0x5f, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x74,
	0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6e, 0x6f, 0x6d, 0x62, 0x72, 0x65, 0x50,
	0x6c, 0x61, 0x6e, 0x65, 0x74, 0x61, 0x12, 0x23, 0x0a, 0x0d, 0x6e, 0x6f, 0x6d, 0x62, 0x72, 0x65,
	0x5f, 0x63, 0x69, 0x75, 0x64, 0x61, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6e,
	0x6f, 0x6d, 0x62, 0x72, 0x65, 0x43, 0x69, 0x75, 0x64, 0x61, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x24, 0x0a, 0x0b, 0x6e, 0x75, 0x65, 0x76, 0x6f, 0x5f, 0x76, 0x61, 0x6c,
	0x6f, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0a, 0x6e, 0x75, 0x65, 0x76,
	0x6f, 0x56, 0x61, 0x6c, 0x6f, 0x72, 0x88, 0x01, 0x01, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x6e, 0x75,
	0x65, 0x76, 0x6f, 0x5f, 0x76, 0x61, 0x6c, 0x6f, 0x72, 0x22, 0x3c, 0x0a, 0x0b, 0x52, 0x65, 0x73,
	0x70, 0x42, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x31, 0x12, 0x2d, 0x0a, 0x12, 0x64, 0x69, 0x72, 0x65,
	0x63, 0x63, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x64, 0x6f, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x64, 0x69, 0x72, 0x65, 0x63, 0x63, 0x69, 0x6f, 0x6e, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x64, 0x6f, 0x72, 0x22, 0x86, 0x01, 0x0a, 0x0b, 0x52, 0x65, 0x73, 0x70,
	0x42, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x32, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x61, 0x6e, 0x74, 0x5f,
	0x72, 0x65, 0x62, 0x65, 0x6c, 0x64, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c,
	0x63, 0x61, 0x6e, 0x74, 0x52, 0x65, 0x62, 0x65, 0x6c, 0x64, 0x65, 0x73, 0x12, 0x21, 0x0a, 0x0c,
	0x72, 0x65, 0x6c, 0x6f, 0x6a, 0x5f, 0x76, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x05, 0x52, 0x0b, 0x72, 0x65, 0x6c, 0x6f, 0x6a, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x12,
	0x2f, 0x0a, 0x13, 0x73, 0x65, 0x72, 0x76, 0x69, 0x64, 0x6f, 0x72, 0x5f, 0x63, 0x6f, 0x6e, 0x74,
	0x61, 0x63, 0x74, 0x61, 0x64, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x64, 0x6f, 0x72, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x61, 0x64, 0x6f,
	0x32, 0x86, 0x01, 0x0a, 0x08, 0x53, 0x74, 0x61, 0x72, 0x57, 0x61, 0x72, 0x73, 0x12, 0x3c, 0x0a,
	0x0e, 0x43, 0x69, 0x74, 0x79, 0x4d, 0x67, 0x6d, 0x74, 0x42, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x12,
	0x11, 0x2e, 0x73, 0x74, 0x61, 0x72, 0x77, 0x61, 0x72, 0x73, 0x2e, 0x4e, 0x65, 0x77, 0x43, 0x69,
	0x74, 0x79, 0x1a, 0x15, 0x2e, 0x73, 0x74, 0x61, 0x72, 0x77, 0x61, 0x72, 0x73, 0x2e, 0x52, 0x65,
	0x73, 0x70, 0x42, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x31, 0x22, 0x00, 0x12, 0x3c, 0x0a, 0x0e, 0x43,
	0x69, 0x74, 0x79, 0x4c, 0x65, 0x69, 0x61, 0x42, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x12, 0x11, 0x2e,
	0x73, 0x74, 0x61, 0x72, 0x77, 0x61, 0x72, 0x73, 0x2e, 0x4e, 0x65, 0x77, 0x43, 0x69, 0x74, 0x79,
	0x1a, 0x15, 0x2e, 0x73, 0x74, 0x61, 0x72, 0x77, 0x61, 0x72, 0x73, 0x2e, 0x52, 0x65, 0x73, 0x70,
	0x42, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x32, 0x22, 0x00, 0x42, 0x2f, 0x5a, 0x2d, 0x65, 0x78, 0x61,
	0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2d, 0x73, 0x74, 0x61, 0x72,
	0x77, 0x61, 0x72, 0x73, 0x2d, 0x67, 0x72, 0x70, 0x63, 0x3b, 0x67, 0x6f, 0x5f, 0x73, 0x74, 0x61,
	0x72, 0x77, 0x61, 0x72, 0x73, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_starwars_starwars_proto_rawDescOnce sync.Once
	file_starwars_starwars_proto_rawDescData = file_starwars_starwars_proto_rawDesc
)

func file_starwars_starwars_proto_rawDescGZIP() []byte {
	file_starwars_starwars_proto_rawDescOnce.Do(func() {
		file_starwars_starwars_proto_rawDescData = protoimpl.X.CompressGZIP(file_starwars_starwars_proto_rawDescData)
	})
	return file_starwars_starwars_proto_rawDescData
}

var file_starwars_starwars_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_starwars_starwars_proto_goTypes = []interface{}{
	(*NewCity)(nil),     // 0: starwars.NewCity
	(*RespBroker1)(nil), // 1: starwars.RespBroker1
	(*RespBroker2)(nil), // 2: starwars.RespBroker2
}
var file_starwars_starwars_proto_depIdxs = []int32{
	0, // 0: starwars.StarWars.CityMgmtBroker:input_type -> starwars.NewCity
	0, // 1: starwars.StarWars.CityLeiaBroker:input_type -> starwars.NewCity
	1, // 2: starwars.StarWars.CityMgmtBroker:output_type -> starwars.RespBroker1
	2, // 3: starwars.StarWars.CityLeiaBroker:output_type -> starwars.RespBroker2
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_starwars_starwars_proto_init() }
func file_starwars_starwars_proto_init() {
	if File_starwars_starwars_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_starwars_starwars_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewCity); i {
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
		file_starwars_starwars_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RespBroker1); i {
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
		file_starwars_starwars_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RespBroker2); i {
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
	file_starwars_starwars_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_starwars_starwars_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_starwars_starwars_proto_goTypes,
		DependencyIndexes: file_starwars_starwars_proto_depIdxs,
		MessageInfos:      file_starwars_starwars_proto_msgTypes,
	}.Build()
	File_starwars_starwars_proto = out.File
	file_starwars_starwars_proto_rawDesc = nil
	file_starwars_starwars_proto_goTypes = nil
	file_starwars_starwars_proto_depIdxs = nil
}
