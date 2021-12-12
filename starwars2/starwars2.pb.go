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
	Sender        string  `protobuf:"bytes,5,opt,name=sender,proto3" json:"sender,omitempty"`
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

func (x *NewCity1) GetSender() string {
	if x != nil {
		return x.Sender
	}
	return ""
}

type Planeta struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NombrePlaneta string `protobuf:"bytes,1,opt,name=nombre_planeta,json=nombrePlaneta,proto3" json:"nombre_planeta,omitempty"`
}

func (x *Planeta) Reset() {
	*x = Planeta{}
	if protoimpl.UnsafeEnabled {
		mi := &file_starwars2_starwars2_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Planeta) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Planeta) ProtoMessage() {}

func (x *Planeta) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use Planeta.ProtoReflect.Descriptor instead.
func (*Planeta) Descriptor() ([]byte, []int) {
	return file_starwars2_starwars2_proto_rawDescGZIP(), []int{1}
}

func (x *Planeta) GetNombrePlaneta() string {
	if x != nil {
		return x.NombrePlaneta
	}
	return ""
}

type RespFulcrum1 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RelojVector  []int32 `protobuf:"varint,1,rep,packed,name=reloj_vector,json=relojVector,proto3" json:"reloj_vector,omitempty"`
	Planeta      string  `protobuf:"bytes,2,opt,name=planeta,proto3" json:"planeta,omitempty"`
	SeRealizoMod string  `protobuf:"bytes,3,opt,name=se_realizo_mod,json=seRealizoMod,proto3" json:"se_realizo_mod,omitempty"`
}

func (x *RespFulcrum1) Reset() {
	*x = RespFulcrum1{}
	if protoimpl.UnsafeEnabled {
		mi := &file_starwars2_starwars2_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RespFulcrum1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RespFulcrum1) ProtoMessage() {}

func (x *RespFulcrum1) ProtoReflect() protoreflect.Message {
	mi := &file_starwars2_starwars2_proto_msgTypes[2]
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
	return file_starwars2_starwars2_proto_rawDescGZIP(), []int{2}
}

func (x *RespFulcrum1) GetRelojVector() []int32 {
	if x != nil {
		return x.RelojVector
	}
	return nil
}

func (x *RespFulcrum1) GetPlaneta() string {
	if x != nil {
		return x.Planeta
	}
	return ""
}

func (x *RespFulcrum1) GetSeRealizoMod() string {
	if x != nil {
		return x.SeRealizoMod
	}
	return ""
}

type Vector struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RelojVector []int32 `protobuf:"varint,1,rep,packed,name=reloj_vector,json=relojVector,proto3" json:"reloj_vector,omitempty"`
	Planeta     string  `protobuf:"bytes,2,opt,name=planeta,proto3" json:"planeta,omitempty"`
}

func (x *Vector) Reset() {
	*x = Vector{}
	if protoimpl.UnsafeEnabled {
		mi := &file_starwars2_starwars2_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Vector) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Vector) ProtoMessage() {}

func (x *Vector) ProtoReflect() protoreflect.Message {
	mi := &file_starwars2_starwars2_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Vector.ProtoReflect.Descriptor instead.
func (*Vector) Descriptor() ([]byte, []int) {
	return file_starwars2_starwars2_proto_rawDescGZIP(), []int{3}
}

func (x *Vector) GetRelojVector() []int32 {
	if x != nil {
		return x.RelojVector
	}
	return nil
}

func (x *Vector) GetPlaneta() string {
	if x != nil {
		return x.Planeta
	}
	return ""
}

type VectoresList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Vectores []*Vector `protobuf:"bytes,1,rep,name=vectores,proto3" json:"vectores,omitempty"`
}

func (x *VectoresList) Reset() {
	*x = VectoresList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_starwars2_starwars2_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VectoresList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VectoresList) ProtoMessage() {}

func (x *VectoresList) ProtoReflect() protoreflect.Message {
	mi := &file_starwars2_starwars2_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VectoresList.ProtoReflect.Descriptor instead.
func (*VectoresList) Descriptor() ([]byte, []int) {
	return file_starwars2_starwars2_proto_rawDescGZIP(), []int{4}
}

func (x *VectoresList) GetVectores() []*Vector {
	if x != nil {
		return x.Vectores
	}
	return nil
}

type RegistroModificado struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Planeta               string  `protobuf:"bytes,1,opt,name=planeta,proto3" json:"planeta,omitempty"`
	RelojVector           []int32 `protobuf:"varint,2,rep,packed,name=reloj_vector,json=relojVector,proto3" json:"reloj_vector,omitempty"`
	UltimoServidorFulcrum string  `protobuf:"bytes,3,opt,name=ultimo_servidor_fulcrum,json=ultimoServidorFulcrum,proto3" json:"ultimo_servidor_fulcrum,omitempty"`
}

func (x *RegistroModificado) Reset() {
	*x = RegistroModificado{}
	if protoimpl.UnsafeEnabled {
		mi := &file_starwars2_starwars2_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegistroModificado) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegistroModificado) ProtoMessage() {}

func (x *RegistroModificado) ProtoReflect() protoreflect.Message {
	mi := &file_starwars2_starwars2_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegistroModificado.ProtoReflect.Descriptor instead.
func (*RegistroModificado) Descriptor() ([]byte, []int) {
	return file_starwars2_starwars2_proto_rawDescGZIP(), []int{5}
}

func (x *RegistroModificado) GetPlaneta() string {
	if x != nil {
		return x.Planeta
	}
	return ""
}

func (x *RegistroModificado) GetRelojVector() []int32 {
	if x != nil {
		return x.RelojVector
	}
	return nil
}

func (x *RegistroModificado) GetUltimoServidorFulcrum() string {
	if x != nil {
		return x.UltimoServidorFulcrum
	}
	return ""
}

type RegistrosModificadosList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Registros []*RegistroModificado `protobuf:"bytes,1,rep,name=registros,proto3" json:"registros,omitempty"`
}

func (x *RegistrosModificadosList) Reset() {
	*x = RegistrosModificadosList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_starwars2_starwars2_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegistrosModificadosList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegistrosModificadosList) ProtoMessage() {}

func (x *RegistrosModificadosList) ProtoReflect() protoreflect.Message {
	mi := &file_starwars2_starwars2_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegistrosModificadosList.ProtoReflect.Descriptor instead.
func (*RegistrosModificadosList) Descriptor() ([]byte, []int) {
	return file_starwars2_starwars2_proto_rawDescGZIP(), []int{6}
}

func (x *RegistrosModificadosList) GetRegistros() []*RegistroModificado {
	if x != nil {
		return x.Registros
	}
	return nil
}

type CiudadSolicitada struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ciudad                string  `protobuf:"bytes,1,opt,name=ciudad,proto3" json:"ciudad,omitempty"`
	Planeta               string  `protobuf:"bytes,2,opt,name=planeta,proto3" json:"planeta,omitempty"`
	RelojVector           []int32 `protobuf:"varint,3,rep,packed,name=reloj_vector,json=relojVector,proto3" json:"reloj_vector,omitempty"`
	UltimoServidorFulcrum string  `protobuf:"bytes,4,opt,name=ultimo_servidor_fulcrum,json=ultimoServidorFulcrum,proto3" json:"ultimo_servidor_fulcrum,omitempty"`
}

func (x *CiudadSolicitada) Reset() {
	*x = CiudadSolicitada{}
	if protoimpl.UnsafeEnabled {
		mi := &file_starwars2_starwars2_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CiudadSolicitada) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CiudadSolicitada) ProtoMessage() {}

func (x *CiudadSolicitada) ProtoReflect() protoreflect.Message {
	mi := &file_starwars2_starwars2_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CiudadSolicitada.ProtoReflect.Descriptor instead.
func (*CiudadSolicitada) Descriptor() ([]byte, []int) {
	return file_starwars2_starwars2_proto_rawDescGZIP(), []int{7}
}

func (x *CiudadSolicitada) GetCiudad() string {
	if x != nil {
		return x.Ciudad
	}
	return ""
}

func (x *CiudadSolicitada) GetPlaneta() string {
	if x != nil {
		return x.Planeta
	}
	return ""
}

func (x *CiudadSolicitada) GetRelojVector() []int32 {
	if x != nil {
		return x.RelojVector
	}
	return nil
}

func (x *CiudadSolicitada) GetUltimoServidorFulcrum() string {
	if x != nil {
		return x.UltimoServidorFulcrum
	}
	return ""
}

type CiudadesSolicitadasList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ciudades []*CiudadSolicitada `protobuf:"bytes,1,rep,name=ciudades,proto3" json:"ciudades,omitempty"`
}

func (x *CiudadesSolicitadasList) Reset() {
	*x = CiudadesSolicitadasList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_starwars2_starwars2_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CiudadesSolicitadasList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CiudadesSolicitadasList) ProtoMessage() {}

func (x *CiudadesSolicitadasList) ProtoReflect() protoreflect.Message {
	mi := &file_starwars2_starwars2_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CiudadesSolicitadasList.ProtoReflect.Descriptor instead.
func (*CiudadesSolicitadasList) Descriptor() ([]byte, []int) {
	return file_starwars2_starwars2_proto_rawDescGZIP(), []int{8}
}

func (x *CiudadesSolicitadasList) GetCiudades() []*CiudadSolicitada {
	if x != nil {
		return x.Ciudades
	}
	return nil
}

type RespFulcrum2 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CantRebeldes       int32   `protobuf:"varint,1,opt,name=cant_rebeldes,json=cantRebeldes,proto3" json:"cant_rebeldes,omitempty"`
	RelojVector        []int32 `protobuf:"varint,2,rep,packed,name=reloj_vector,json=relojVector,proto3" json:"reloj_vector,omitempty"`
	ServidorContactado string  `protobuf:"bytes,3,opt,name=servidor_contactado,json=servidorContactado,proto3" json:"servidor_contactado,omitempty"`
}

func (x *RespFulcrum2) Reset() {
	*x = RespFulcrum2{}
	if protoimpl.UnsafeEnabled {
		mi := &file_starwars2_starwars2_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RespFulcrum2) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RespFulcrum2) ProtoMessage() {}

func (x *RespFulcrum2) ProtoReflect() protoreflect.Message {
	mi := &file_starwars2_starwars2_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RespFulcrum2.ProtoReflect.Descriptor instead.
func (*RespFulcrum2) Descriptor() ([]byte, []int) {
	return file_starwars2_starwars2_proto_rawDescGZIP(), []int{9}
}

func (x *RespFulcrum2) GetCantRebeldes() int32 {
	if x != nil {
		return x.CantRebeldes
	}
	return 0
}

func (x *RespFulcrum2) GetRelojVector() []int32 {
	if x != nil {
		return x.RelojVector
	}
	return nil
}

func (x *RespFulcrum2) GetServidorContactado() string {
	if x != nil {
		return x.ServidorContactado
	}
	return ""
}

var File_starwars2_starwars2_proto protoreflect.FileDescriptor

var file_starwars2_starwars2_proto_rawDesc = []byte{
	0x0a, 0x19, 0x73, 0x74, 0x61, 0x72, 0x77, 0x61, 0x72, 0x73, 0x32, 0x2f, 0x73, 0x74, 0x61, 0x72,
	0x77, 0x61, 0x72, 0x73, 0x32, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x73, 0x74, 0x61,
	0x72, 0x77, 0x61, 0x72, 0x73, 0x32, 0x22, 0xbc, 0x01, 0x0a, 0x08, 0x4e, 0x65, 0x77, 0x43, 0x69,
	0x74, 0x79, 0x31, 0x12, 0x25, 0x0a, 0x0e, 0x6e, 0x6f, 0x6d, 0x62, 0x72, 0x65, 0x5f, 0x70, 0x6c,
	0x61, 0x6e, 0x65, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6e, 0x6f, 0x6d,
	0x62, 0x72, 0x65, 0x50, 0x6c, 0x61, 0x6e, 0x65, 0x74, 0x61, 0x12, 0x23, 0x0a, 0x0d, 0x6e, 0x6f,
	0x6d, 0x62, 0x72, 0x65, 0x5f, 0x63, 0x69, 0x75, 0x64, 0x61, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0c, 0x6e, 0x6f, 0x6d, 0x62, 0x72, 0x65, 0x43, 0x69, 0x75, 0x64, 0x61, 0x64, 0x12,
	0x16, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x24, 0x0a, 0x0b, 0x6e, 0x75, 0x65, 0x76, 0x6f,
	0x5f, 0x76, 0x61, 0x6c, 0x6f, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0a,
	0x6e, 0x75, 0x65, 0x76, 0x6f, 0x56, 0x61, 0x6c, 0x6f, 0x72, 0x88, 0x01, 0x01, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73,
	0x65, 0x6e, 0x64, 0x65, 0x72, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x6e, 0x75, 0x65, 0x76, 0x6f, 0x5f,
	0x76, 0x61, 0x6c, 0x6f, 0x72, 0x22, 0x30, 0x0a, 0x07, 0x50, 0x6c, 0x61, 0x6e, 0x65, 0x74, 0x61,
	0x12, 0x25, 0x0a, 0x0e, 0x6e, 0x6f, 0x6d, 0x62, 0x72, 0x65, 0x5f, 0x70, 0x6c, 0x61, 0x6e, 0x65,
	0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6e, 0x6f, 0x6d, 0x62, 0x72, 0x65,
	0x50, 0x6c, 0x61, 0x6e, 0x65, 0x74, 0x61, 0x22, 0x71, 0x0a, 0x0c, 0x52, 0x65, 0x73, 0x70, 0x46,
	0x75, 0x6c, 0x63, 0x72, 0x75, 0x6d, 0x31, 0x12, 0x21, 0x0a, 0x0c, 0x72, 0x65, 0x6c, 0x6f, 0x6a,
	0x5f, 0x76, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x03, 0x28, 0x05, 0x52, 0x0b, 0x72,
	0x65, 0x6c, 0x6f, 0x6a, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x6c,
	0x61, 0x6e, 0x65, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x6c, 0x61,
	0x6e, 0x65, 0x74, 0x61, 0x12, 0x24, 0x0a, 0x0e, 0x73, 0x65, 0x5f, 0x72, 0x65, 0x61, 0x6c, 0x69,
	0x7a, 0x6f, 0x5f, 0x6d, 0x6f, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x73, 0x65,
	0x52, 0x65, 0x61, 0x6c, 0x69, 0x7a, 0x6f, 0x4d, 0x6f, 0x64, 0x22, 0x45, 0x0a, 0x06, 0x56, 0x65,
	0x63, 0x74, 0x6f, 0x72, 0x12, 0x21, 0x0a, 0x0c, 0x72, 0x65, 0x6c, 0x6f, 0x6a, 0x5f, 0x76, 0x65,
	0x63, 0x74, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x03, 0x28, 0x05, 0x52, 0x0b, 0x72, 0x65, 0x6c, 0x6f,
	0x6a, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x6c, 0x61, 0x6e, 0x65,
	0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x74,
	0x61, 0x22, 0x3d, 0x0a, 0x0c, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x65, 0x73, 0x4c, 0x69, 0x73,
	0x74, 0x12, 0x2d, 0x0a, 0x08, 0x76, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x65, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x73, 0x74, 0x61, 0x72, 0x77, 0x61, 0x72, 0x73, 0x32, 0x2e,
	0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x08, 0x76, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x65, 0x73,
	0x22, 0x89, 0x01, 0x0a, 0x12, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x6f, 0x4d, 0x6f, 0x64,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x64, 0x6f, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x6c, 0x61, 0x6e, 0x65,
	0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x74,
	0x61, 0x12, 0x21, 0x0a, 0x0c, 0x72, 0x65, 0x6c, 0x6f, 0x6a, 0x5f, 0x76, 0x65, 0x63, 0x74, 0x6f,
	0x72, 0x18, 0x02, 0x20, 0x03, 0x28, 0x05, 0x52, 0x0b, 0x72, 0x65, 0x6c, 0x6f, 0x6a, 0x56, 0x65,
	0x63, 0x74, 0x6f, 0x72, 0x12, 0x36, 0x0a, 0x17, 0x75, 0x6c, 0x74, 0x69, 0x6d, 0x6f, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x64, 0x6f, 0x72, 0x5f, 0x66, 0x75, 0x6c, 0x63, 0x72, 0x75, 0x6d, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x15, 0x75, 0x6c, 0x74, 0x69, 0x6d, 0x6f, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x64, 0x6f, 0x72, 0x46, 0x75, 0x6c, 0x63, 0x72, 0x75, 0x6d, 0x22, 0x57, 0x0a, 0x18,
	0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x6f, 0x73, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x64, 0x6f, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x3b, 0x0a, 0x09, 0x72, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x72, 0x6f, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x73, 0x74,
	0x61, 0x72, 0x77, 0x61, 0x72, 0x73, 0x32, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x6f,
	0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x63, 0x61, 0x64, 0x6f, 0x52, 0x09, 0x72, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x72, 0x6f, 0x73, 0x22, 0x9f, 0x01, 0x0a, 0x10, 0x43, 0x69, 0x75, 0x64, 0x61, 0x64,
	0x53, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x74, 0x61, 0x64, 0x61, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x69,
	0x75, 0x64, 0x61, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x69, 0x75, 0x64,
	0x61, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x74, 0x61, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x74, 0x61, 0x12, 0x21, 0x0a, 0x0c,
	0x72, 0x65, 0x6c, 0x6f, 0x6a, 0x5f, 0x76, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x05, 0x52, 0x0b, 0x72, 0x65, 0x6c, 0x6f, 0x6a, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x12,
	0x36, 0x0a, 0x17, 0x75, 0x6c, 0x74, 0x69, 0x6d, 0x6f, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x64,
	0x6f, 0x72, 0x5f, 0x66, 0x75, 0x6c, 0x63, 0x72, 0x75, 0x6d, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x15, 0x75, 0x6c, 0x74, 0x69, 0x6d, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x64, 0x6f, 0x72,
	0x46, 0x75, 0x6c, 0x63, 0x72, 0x75, 0x6d, 0x22, 0x52, 0x0a, 0x17, 0x43, 0x69, 0x75, 0x64, 0x61,
	0x64, 0x65, 0x73, 0x53, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x74, 0x61, 0x64, 0x61, 0x73, 0x4c, 0x69,
	0x73, 0x74, 0x12, 0x37, 0x0a, 0x08, 0x63, 0x69, 0x75, 0x64, 0x61, 0x64, 0x65, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x73, 0x74, 0x61, 0x72, 0x77, 0x61, 0x72, 0x73, 0x32,
	0x2e, 0x43, 0x69, 0x75, 0x64, 0x61, 0x64, 0x53, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x74, 0x61, 0x64,
	0x61, 0x52, 0x08, 0x63, 0x69, 0x75, 0x64, 0x61, 0x64, 0x65, 0x73, 0x22, 0x87, 0x01, 0x0a, 0x0c,
	0x52, 0x65, 0x73, 0x70, 0x46, 0x75, 0x6c, 0x63, 0x72, 0x75, 0x6d, 0x32, 0x12, 0x23, 0x0a, 0x0d,
	0x63, 0x61, 0x6e, 0x74, 0x5f, 0x72, 0x65, 0x62, 0x65, 0x6c, 0x64, 0x65, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0c, 0x63, 0x61, 0x6e, 0x74, 0x52, 0x65, 0x62, 0x65, 0x6c, 0x64, 0x65,
	0x73, 0x12, 0x21, 0x0a, 0x0c, 0x72, 0x65, 0x6c, 0x6f, 0x6a, 0x5f, 0x76, 0x65, 0x63, 0x74, 0x6f,
	0x72, 0x18, 0x02, 0x20, 0x03, 0x28, 0x05, 0x52, 0x0b, 0x72, 0x65, 0x6c, 0x6f, 0x6a, 0x56, 0x65,
	0x63, 0x74, 0x6f, 0x72, 0x12, 0x2f, 0x0a, 0x13, 0x73, 0x65, 0x72, 0x76, 0x69, 0x64, 0x6f, 0x72,
	0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x61, 0x64, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x12, 0x73, 0x65, 0x72, 0x76, 0x69, 0x64, 0x6f, 0x72, 0x43, 0x6f, 0x6e, 0x74, 0x61,
	0x63, 0x74, 0x61, 0x64, 0x6f, 0x32, 0xda, 0x01, 0x0a, 0x09, 0x53, 0x74, 0x61, 0x72, 0x57, 0x61,
	0x72, 0x73, 0x31, 0x12, 0x41, 0x0a, 0x0f, 0x43, 0x69, 0x74, 0x79, 0x4d, 0x67, 0x6d, 0x74, 0x46,
	0x75, 0x6c, 0x63, 0x72, 0x75, 0x6d, 0x12, 0x13, 0x2e, 0x73, 0x74, 0x61, 0x72, 0x77, 0x61, 0x72,
	0x73, 0x32, 0x2e, 0x4e, 0x65, 0x77, 0x43, 0x69, 0x74, 0x79, 0x31, 0x1a, 0x17, 0x2e, 0x73, 0x74,
	0x61, 0x72, 0x77, 0x61, 0x72, 0x73, 0x32, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x46, 0x75, 0x6c, 0x63,
	0x72, 0x75, 0x6d, 0x31, 0x22, 0x00, 0x12, 0x43, 0x0a, 0x11, 0x43, 0x69, 0x74, 0x79, 0x42, 0x72,
	0x6f, 0x6b, 0x65, 0x72, 0x46, 0x75, 0x6c, 0x63, 0x72, 0x75, 0x6d, 0x12, 0x13, 0x2e, 0x73, 0x74,
	0x61, 0x72, 0x77, 0x61, 0x72, 0x73, 0x32, 0x2e, 0x4e, 0x65, 0x77, 0x43, 0x69, 0x74, 0x79, 0x31,
	0x1a, 0x17, 0x2e, 0x73, 0x74, 0x61, 0x72, 0x77, 0x61, 0x72, 0x73, 0x32, 0x2e, 0x52, 0x65, 0x73,
	0x70, 0x46, 0x75, 0x6c, 0x63, 0x72, 0x75, 0x6d, 0x32, 0x22, 0x00, 0x12, 0x45, 0x0a, 0x14, 0x52,
	0x65, 0x6c, 0x6f, 0x6a, 0x65, 0x73, 0x42, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x46, 0x75, 0x6c, 0x63,
	0x72, 0x75, 0x6d, 0x12, 0x12, 0x2e, 0x73, 0x74, 0x61, 0x72, 0x77, 0x61, 0x72, 0x73, 0x32, 0x2e,
	0x50, 0x6c, 0x61, 0x6e, 0x65, 0x74, 0x61, 0x1a, 0x17, 0x2e, 0x73, 0x74, 0x61, 0x72, 0x77, 0x61,
	0x72, 0x73, 0x32, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x46, 0x75, 0x6c, 0x63, 0x72, 0x75, 0x6d, 0x31,
	0x22, 0x00, 0x42, 0x2f, 0x5a, 0x2d, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x67, 0x6f, 0x2d, 0x73, 0x74, 0x61, 0x72, 0x77, 0x61, 0x72, 0x73, 0x2d, 0x67, 0x72,
	0x70, 0x63, 0x3b, 0x67, 0x6f, 0x5f, 0x73, 0x74, 0x61, 0x72, 0x77, 0x61, 0x72, 0x73, 0x5f, 0x67,
	0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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

var file_starwars2_starwars2_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_starwars2_starwars2_proto_goTypes = []interface{}{
	(*NewCity1)(nil),                 // 0: starwars2.NewCity1
	(*Planeta)(nil),                  // 1: starwars2.Planeta
	(*RespFulcrum1)(nil),             // 2: starwars2.RespFulcrum1
	(*Vector)(nil),                   // 3: starwars2.Vector
	(*VectoresList)(nil),             // 4: starwars2.VectoresList
	(*RegistroModificado)(nil),       // 5: starwars2.RegistroModificado
	(*RegistrosModificadosList)(nil), // 6: starwars2.RegistrosModificadosList
	(*CiudadSolicitada)(nil),         // 7: starwars2.CiudadSolicitada
	(*CiudadesSolicitadasList)(nil),  // 8: starwars2.CiudadesSolicitadasList
	(*RespFulcrum2)(nil),             // 9: starwars2.RespFulcrum2
}
var file_starwars2_starwars2_proto_depIdxs = []int32{
	3, // 0: starwars2.VectoresList.vectores:type_name -> starwars2.Vector
	5, // 1: starwars2.RegistrosModificadosList.registros:type_name -> starwars2.RegistroModificado
	7, // 2: starwars2.CiudadesSolicitadasList.ciudades:type_name -> starwars2.CiudadSolicitada
	0, // 3: starwars2.StarWars1.CityMgmtFulcrum:input_type -> starwars2.NewCity1
	0, // 4: starwars2.StarWars1.CityBrokerFulcrum:input_type -> starwars2.NewCity1
	1, // 5: starwars2.StarWars1.RelojesBrokerFulcrum:input_type -> starwars2.Planeta
	2, // 6: starwars2.StarWars1.CityMgmtFulcrum:output_type -> starwars2.RespFulcrum1
	9, // 7: starwars2.StarWars1.CityBrokerFulcrum:output_type -> starwars2.RespFulcrum2
	2, // 8: starwars2.StarWars1.RelojesBrokerFulcrum:output_type -> starwars2.RespFulcrum1
	6, // [6:9] is the sub-list for method output_type
	3, // [3:6] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
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
			switch v := v.(*Planeta); i {
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
		file_starwars2_starwars2_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
		file_starwars2_starwars2_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Vector); i {
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
		file_starwars2_starwars2_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VectoresList); i {
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
		file_starwars2_starwars2_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegistroModificado); i {
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
		file_starwars2_starwars2_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegistrosModificadosList); i {
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
		file_starwars2_starwars2_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CiudadSolicitada); i {
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
		file_starwars2_starwars2_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CiudadesSolicitadasList); i {
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
		file_starwars2_starwars2_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RespFulcrum2); i {
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
			NumMessages:   10,
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
