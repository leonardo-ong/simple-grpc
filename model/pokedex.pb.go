// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.0
// 	protoc        v3.17.3
// source: pokedex.proto

package model

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type PokemonId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *PokemonId) Reset() {
	*x = PokemonId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pokedex_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PokemonId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PokemonId) ProtoMessage() {}

func (x *PokemonId) ProtoReflect() protoreflect.Message {
	mi := &file_pokedex_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PokemonId.ProtoReflect.Descriptor instead.
func (*PokemonId) Descriptor() ([]byte, []int) {
	return file_pokedex_proto_rawDescGZIP(), []int{0}
}

func (x *PokemonId) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type PokemonCount struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Count int32 `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *PokemonCount) Reset() {
	*x = PokemonCount{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pokedex_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PokemonCount) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PokemonCount) ProtoMessage() {}

func (x *PokemonCount) ProtoReflect() protoreflect.Message {
	mi := &file_pokedex_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PokemonCount.ProtoReflect.Descriptor instead.
func (*PokemonCount) Descriptor() ([]byte, []int) {
	return file_pokedex_proto_rawDescGZIP(), []int{1}
}

func (x *PokemonCount) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

type Pokemon struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name   string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Height string `protobuf:"bytes,3,opt,name=height,proto3" json:"height,omitempty"`
	Weight string `protobuf:"bytes,4,opt,name=weight,proto3" json:"weight,omitempty"`
	Types  string `protobuf:"bytes,5,opt,name=types,proto3" json:"types,omitempty"`
}

func (x *Pokemon) Reset() {
	*x = Pokemon{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pokedex_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Pokemon) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Pokemon) ProtoMessage() {}

func (x *Pokemon) ProtoReflect() protoreflect.Message {
	mi := &file_pokedex_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Pokemon.ProtoReflect.Descriptor instead.
func (*Pokemon) Descriptor() ([]byte, []int) {
	return file_pokedex_proto_rawDescGZIP(), []int{2}
}

func (x *Pokemon) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Pokemon) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Pokemon) GetHeight() string {
	if x != nil {
		return x.Height
	}
	return ""
}

func (x *Pokemon) GetWeight() string {
	if x != nil {
		return x.Weight
	}
	return ""
}

func (x *Pokemon) GetTypes() string {
	if x != nil {
		return x.Types
	}
	return ""
}

type PokemonList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	List []*Pokemon `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
}

func (x *PokemonList) Reset() {
	*x = PokemonList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pokedex_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PokemonList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PokemonList) ProtoMessage() {}

func (x *PokemonList) ProtoReflect() protoreflect.Message {
	mi := &file_pokedex_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PokemonList.ProtoReflect.Descriptor instead.
func (*PokemonList) Descriptor() ([]byte, []int) {
	return file_pokedex_proto_rawDescGZIP(), []int{3}
}

func (x *PokemonList) GetList() []*Pokemon {
	if x != nil {
		return x.List
	}
	return nil
}

var File_pokedex_proto protoreflect.FileDescriptor

var file_pokedex_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x70, 0x6f, 0x6b, 0x65, 0x64, 0x65, 0x78, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x05, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x1b, 0x0a, 0x09, 0x50, 0x6f, 0x6b, 0x65, 0x6d, 0x6f, 0x6e, 0x49, 0x64,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x22, 0x24, 0x0a, 0x0c, 0x50, 0x6f, 0x6b, 0x65, 0x6d, 0x6f, 0x6e, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x73, 0x0a, 0x07, 0x50, 0x6f, 0x6b, 0x65, 0x6d, 0x6f,
	0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x16, 0x0a,
	0x06, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x77,
	0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x79, 0x70, 0x65, 0x73, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x79, 0x70, 0x65, 0x73, 0x22, 0x31, 0x0a, 0x0b, 0x50,
	0x6f, 0x6b, 0x65, 0x6d, 0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x04, 0x6c, 0x69,
	0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x2e, 0x50, 0x6f, 0x6b, 0x65, 0x6d, 0x6f, 0x6e, 0x52, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x32, 0xba,
	0x01, 0x0a, 0x07, 0x50, 0x6f, 0x6b, 0x65, 0x64, 0x65, 0x78, 0x12, 0x3d, 0x0a, 0x0c, 0x54, 0x6f,
	0x74, 0x61, 0x6c, 0x50, 0x6f, 0x6b, 0x65, 0x6d, 0x6f, 0x6e, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x1a, 0x13, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x50, 0x6f, 0x6b, 0x65, 0x6d,
	0x6f, 0x6e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x00, 0x12, 0x30, 0x0a, 0x0a, 0x47, 0x65, 0x74,
	0x50, 0x6f, 0x6b, 0x65, 0x6d, 0x6f, 0x6e, 0x12, 0x10, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e,
	0x50, 0x6f, 0x6b, 0x65, 0x6d, 0x6f, 0x6e, 0x49, 0x64, 0x1a, 0x0e, 0x2e, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x2e, 0x50, 0x6f, 0x6b, 0x65, 0x6d, 0x6f, 0x6e, 0x22, 0x00, 0x12, 0x3e, 0x0a, 0x0e, 0x47,
	0x65, 0x74, 0x50, 0x6f, 0x6b, 0x65, 0x6d, 0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x16, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x12, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x50, 0x6f,
	0x6b, 0x65, 0x6d, 0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x00, 0x42, 0x14, 0x5a, 0x12, 0x67,
	0x72, 0x70, 0x63, 0x5f, 0x70, 0x6f, 0x6b, 0x65, 0x64, 0x65, 0x78, 0x2f, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pokedex_proto_rawDescOnce sync.Once
	file_pokedex_proto_rawDescData = file_pokedex_proto_rawDesc
)

func file_pokedex_proto_rawDescGZIP() []byte {
	file_pokedex_proto_rawDescOnce.Do(func() {
		file_pokedex_proto_rawDescData = protoimpl.X.CompressGZIP(file_pokedex_proto_rawDescData)
	})
	return file_pokedex_proto_rawDescData
}

var file_pokedex_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_pokedex_proto_goTypes = []interface{}{
	(*PokemonId)(nil),     // 0: model.PokemonId
	(*PokemonCount)(nil),  // 1: model.PokemonCount
	(*Pokemon)(nil),       // 2: model.Pokemon
	(*PokemonList)(nil),   // 3: model.PokemonList
	(*emptypb.Empty)(nil), // 4: google.protobuf.Empty
}
var file_pokedex_proto_depIdxs = []int32{
	2, // 0: model.PokemonList.list:type_name -> model.Pokemon
	4, // 1: model.Pokedex.TotalPokemon:input_type -> google.protobuf.Empty
	0, // 2: model.Pokedex.GetPokemon:input_type -> model.PokemonId
	4, // 3: model.Pokedex.GetPokemonList:input_type -> google.protobuf.Empty
	1, // 4: model.Pokedex.TotalPokemon:output_type -> model.PokemonCount
	2, // 5: model.Pokedex.GetPokemon:output_type -> model.Pokemon
	3, // 6: model.Pokedex.GetPokemonList:output_type -> model.PokemonList
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_pokedex_proto_init() }
func file_pokedex_proto_init() {
	if File_pokedex_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pokedex_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PokemonId); i {
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
		file_pokedex_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PokemonCount); i {
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
		file_pokedex_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Pokemon); i {
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
		file_pokedex_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PokemonList); i {
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
			RawDescriptor: file_pokedex_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pokedex_proto_goTypes,
		DependencyIndexes: file_pokedex_proto_depIdxs,
		MessageInfos:      file_pokedex_proto_msgTypes,
	}.Build()
	File_pokedex_proto = out.File
	file_pokedex_proto_rawDesc = nil
	file_pokedex_proto_goTypes = nil
	file_pokedex_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// PokedexClient is the client API for Pokedex service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PokedexClient interface {
	TotalPokemon(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*PokemonCount, error)
	GetPokemon(ctx context.Context, in *PokemonId, opts ...grpc.CallOption) (*Pokemon, error)
	GetPokemonList(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*PokemonList, error)
}

type pokedexClient struct {
	cc grpc.ClientConnInterface
}

func NewPokedexClient(cc grpc.ClientConnInterface) PokedexClient {
	return &pokedexClient{cc}
}

func (c *pokedexClient) TotalPokemon(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*PokemonCount, error) {
	out := new(PokemonCount)
	err := c.cc.Invoke(ctx, "/model.Pokedex/TotalPokemon", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pokedexClient) GetPokemon(ctx context.Context, in *PokemonId, opts ...grpc.CallOption) (*Pokemon, error) {
	out := new(Pokemon)
	err := c.cc.Invoke(ctx, "/model.Pokedex/GetPokemon", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pokedexClient) GetPokemonList(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*PokemonList, error) {
	out := new(PokemonList)
	err := c.cc.Invoke(ctx, "/model.Pokedex/GetPokemonList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PokedexServer is the server API for Pokedex service.
type PokedexServer interface {
	TotalPokemon(context.Context, *emptypb.Empty) (*PokemonCount, error)
	GetPokemon(context.Context, *PokemonId) (*Pokemon, error)
	GetPokemonList(context.Context, *emptypb.Empty) (*PokemonList, error)
}

// UnimplementedPokedexServer can be embedded to have forward compatible implementations.
type UnimplementedPokedexServer struct {
}

func (*UnimplementedPokedexServer) TotalPokemon(context.Context, *emptypb.Empty) (*PokemonCount, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TotalPokemon not implemented")
}
func (*UnimplementedPokedexServer) GetPokemon(context.Context, *PokemonId) (*Pokemon, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPokemon not implemented")
}
func (*UnimplementedPokedexServer) GetPokemonList(context.Context, *emptypb.Empty) (*PokemonList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPokemonList not implemented")
}

func RegisterPokedexServer(s *grpc.Server, srv PokedexServer) {
	s.RegisterService(&_Pokedex_serviceDesc, srv)
}

func _Pokedex_TotalPokemon_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PokedexServer).TotalPokemon(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/model.Pokedex/TotalPokemon",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PokedexServer).TotalPokemon(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Pokedex_GetPokemon_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PokemonId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PokedexServer).GetPokemon(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/model.Pokedex/GetPokemon",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PokedexServer).GetPokemon(ctx, req.(*PokemonId))
	}
	return interceptor(ctx, in, info, handler)
}

func _Pokedex_GetPokemonList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PokedexServer).GetPokemonList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/model.Pokedex/GetPokemonList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PokedexServer).GetPokemonList(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _Pokedex_serviceDesc = grpc.ServiceDesc{
	ServiceName: "model.Pokedex",
	HandlerType: (*PokedexServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "TotalPokemon",
			Handler:    _Pokedex_TotalPokemon_Handler,
		},
		{
			MethodName: "GetPokemon",
			Handler:    _Pokedex_GetPokemon_Handler,
		},
		{
			MethodName: "GetPokemonList",
			Handler:    _Pokedex_GetPokemonList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pokedex.proto",
}
