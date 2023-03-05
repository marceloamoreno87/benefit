package msbenefit

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type NewBenefit struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Doc string `protobuf:"bytes,1,opt,name=doc,proto3" json:"doc,omitempty"`
}

func (x *NewBenefit) Reset() {
	*x = NewBenefit{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_benefit_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewBenefit) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewBenefit) ProtoMessage() {}

func (x *NewBenefit) ProtoReflect() protoreflect.Message {
	mi := &file_proto_benefit_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (*NewBenefit) Descriptor() ([]byte, []int) {
	return file_proto_benefit_proto_rawDescGZIP(), []int{0}
}

func (x *NewBenefit) GetDoc() string {
	if x != nil {
		return x.Doc
	}
	return ""
}

type Benefit struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Nb string `protobuf:"bytes,1,opt,name=nb,proto3" json:"nb,omitempty"`
}

func (x *Benefit) Reset() {
	*x = Benefit{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_benefit_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Benefit) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Benefit) ProtoMessage() {}

func (x *Benefit) ProtoReflect() protoreflect.Message {
	mi := &file_proto_benefit_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (*Benefit) Descriptor() ([]byte, []int) {
	return file_proto_benefit_proto_rawDescGZIP(), []int{1}
}

func (x *Benefit) GetNb() string {
	if x != nil {
		return x.Nb
	}
	return ""
}

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cpf        string     `protobuf:"bytes,1,opt,name=cpf,proto3" json:"cpf,omitempty"`
	Beneficios []*Benefit `protobuf:"bytes,2,rep,name=beneficios,proto3" json:"beneficios,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_benefit_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_proto_benefit_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (*User) Descriptor() ([]byte, []int) {
	return file_proto_benefit_proto_rawDescGZIP(), []int{2}
}

func (x *User) GetCpf() string {
	if x != nil {
		return x.Cpf
	}
	return ""
}

func (x *User) GetBeneficios() []*Benefit {
	if x != nil {
		return x.Beneficios
	}
	return nil
}

var File_proto_benefit_proto protoreflect.FileDescriptor

var file_proto_benefit_proto_rawDesc = []byte{
	0x0a, 0x13, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x62, 0x65, 0x6e, 0x65, 0x66, 0x69, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x1e, 0x0a, 0x0a,
	0x4e, 0x65, 0x77, 0x42, 0x65, 0x6e, 0x65, 0x66, 0x69, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x64, 0x6f,
	0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x64, 0x6f, 0x63, 0x22, 0x19, 0x0a, 0x07,
	0x42, 0x65, 0x6e, 0x65, 0x66, 0x69, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x6e, 0x62, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x6e, 0x62, 0x22, 0x48, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12,
	0x10, 0x0a, 0x03, 0x63, 0x70, 0x66, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x63, 0x70,
	0x66, 0x12, 0x2e, 0x0a, 0x0a, 0x62, 0x65, 0x6e, 0x65, 0x66, 0x69, 0x63, 0x69, 0x6f, 0x73, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x42, 0x65,
	0x6e, 0x65, 0x66, 0x69, 0x74, 0x52, 0x0a, 0x62, 0x65, 0x6e, 0x65, 0x66, 0x69, 0x63, 0x69, 0x6f,
	0x73, 0x32, 0x40, 0x0a, 0x0e, 0x42, 0x65, 0x6e, 0x65, 0x66, 0x69, 0x74, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x2e, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x42, 0x65, 0x6e, 0x65, 0x66, 0x69,
	0x74, 0x12, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4e, 0x65, 0x77, 0x42, 0x65, 0x6e,
	0x65, 0x66, 0x69, 0x74, 0x1a, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x73, 0x65,
	0x72, 0x22, 0x00, 0x42, 0x25, 0x5a, 0x23, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x6d, 0x61, 0x72, 0x63, 0x65, 0x6c, 0x6f, 0x61, 0x6d, 0x6f, 0x72, 0x65, 0x6e, 0x6f,
	0x38, 0x37, 0x2f, 0x63, 0x72, 0x61, 0x77, 0x6c, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_proto_benefit_proto_rawDescOnce sync.Once
	file_proto_benefit_proto_rawDescData = file_proto_benefit_proto_rawDesc
)

func file_proto_benefit_proto_rawDescGZIP() []byte {
	file_proto_benefit_proto_rawDescOnce.Do(func() {
		file_proto_benefit_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_benefit_proto_rawDescData)
	})
	return file_proto_benefit_proto_rawDescData
}

var file_proto_benefit_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_proto_benefit_proto_goTypes = []interface{}{
	(*NewBenefit)(nil),
	(*Benefit)(nil),
	(*User)(nil),
}
var file_proto_benefit_proto_depIdxs = []int32{
	1,
	0,
	2,
	2,
	1,
	1,
	1,
	0,
}

func init() { file_proto_benefit_proto_init() }
func file_proto_benefit_proto_init() {
	if File_proto_benefit_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_benefit_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewBenefit); i {
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
		file_proto_benefit_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Benefit); i {
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
		file_proto_benefit_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
			RawDescriptor: file_proto_benefit_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_benefit_proto_goTypes,
		DependencyIndexes: file_proto_benefit_proto_depIdxs,
		MessageInfos:      file_proto_benefit_proto_msgTypes,
	}.Build()
	File_proto_benefit_proto = out.File
	file_proto_benefit_proto_rawDesc = nil
	file_proto_benefit_proto_goTypes = nil
	file_proto_benefit_proto_depIdxs = nil
}
