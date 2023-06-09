// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.6.1
// source: release_namer.proto

package releasenamer

import (
	empty "github.com/golang/protobuf/ptypes/empty"
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

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_release_namer_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_release_namer_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_release_namer_proto_rawDescGZIP(), []int{0}
}

type CandyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Candies []*Candy `protobuf:"bytes,1,rep,name=Candies,proto3" json:"Candies,omitempty"`
}

func (x *CandyResponse) Reset() {
	*x = CandyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_release_namer_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CandyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CandyResponse) ProtoMessage() {}

func (x *CandyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_release_namer_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CandyResponse.ProtoReflect.Descriptor instead.
func (*CandyResponse) Descriptor() ([]byte, []int) {
	return file_release_namer_proto_rawDescGZIP(), []int{1}
}

func (x *CandyResponse) GetCandies() []*Candy {
	if x != nil {
		return x.Candies
	}
	return nil
}

type Candy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Candy) Reset() {
	*x = Candy{}
	if protoimpl.UnsafeEnabled {
		mi := &file_release_namer_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Candy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Candy) ProtoMessage() {}

func (x *Candy) ProtoReflect() protoreflect.Message {
	mi := &file_release_namer_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Candy.ProtoReflect.Descriptor instead.
func (*Candy) Descriptor() ([]byte, []int) {
	return file_release_namer_proto_rawDescGZIP(), []int{2}
}

func (x *Candy) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_release_namer_proto protoreflect.FileDescriptor

var file_release_namer_proto_rawDesc = []byte{
	0x0a, 0x13, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x6e, 0x61,
	0x6d, 0x65, 0x72, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x3e, 0x0a, 0x0d, 0x43, 0x61, 0x6e,
	0x64, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2d, 0x0a, 0x07, 0x43, 0x61,
	0x6e, 0x64, 0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x72, 0x65,
	0x6c, 0x65, 0x61, 0x73, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x72, 0x2e, 0x43, 0x61, 0x6e, 0x64, 0x79,
	0x52, 0x07, 0x43, 0x61, 0x6e, 0x64, 0x69, 0x65, 0x73, 0x22, 0x1b, 0x0a, 0x05, 0x43, 0x61, 0x6e,
	0x64, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x32, 0x9c, 0x01, 0x0a, 0x0c, 0x52, 0x65, 0x6c, 0x65, 0x61,
	0x73, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x72, 0x12, 0x43, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x43, 0x61,
	0x6e, 0x64, 0x69, 0x65, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1b, 0x2e,
	0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x72, 0x2e, 0x43, 0x61, 0x6e,
	0x64, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x47, 0x0a, 0x0e,
	0x47, 0x65, 0x74, 0x52, 0x61, 0x6e, 0x64, 0x6f, 0x6d, 0x43, 0x61, 0x6e, 0x64, 0x79, 0x12, 0x16,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1b, 0x2e, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65,
	0x6e, 0x61, 0x6d, 0x65, 0x72, 0x2e, 0x43, 0x61, 0x6e, 0x64, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x6a, 0x0a, 0x1d, 0x69, 0x6f, 0x2e, 0x67, 0x72, 0x70, 0x63,
	0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2e, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73,
	0x65, 0x6e, 0x61, 0x6d, 0x65, 0x72, 0x42, 0x11, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x4e,
	0x61, 0x6d, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x34, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x77, 0x69, 0x6c, 0x6c, 0x69, 0x61, 0x6d, 0x7a,
	0x65, 0x6c, 0x65, 0x73, 0x6e, 0x79, 0x2f, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x2d, 0x6e,
	0x61, 0x6d, 0x65, 0x72, 0x2f, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x6e, 0x61, 0x6d, 0x65,
	0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_release_namer_proto_rawDescOnce sync.Once
	file_release_namer_proto_rawDescData = file_release_namer_proto_rawDesc
)

func file_release_namer_proto_rawDescGZIP() []byte {
	file_release_namer_proto_rawDescOnce.Do(func() {
		file_release_namer_proto_rawDescData = protoimpl.X.CompressGZIP(file_release_namer_proto_rawDescData)
	})
	return file_release_namer_proto_rawDescData
}

var file_release_namer_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_release_namer_proto_goTypes = []interface{}{
	(*Empty)(nil),         // 0: releasenamer.Empty
	(*CandyResponse)(nil), // 1: releasenamer.CandyResponse
	(*Candy)(nil),         // 2: releasenamer.Candy
	(*empty.Empty)(nil),   // 3: google.protobuf.Empty
}
var file_release_namer_proto_depIdxs = []int32{
	2, // 0: releasenamer.CandyResponse.Candies:type_name -> releasenamer.Candy
	3, // 1: releasenamer.ReleaseNamer.GetCandies:input_type -> google.protobuf.Empty
	3, // 2: releasenamer.ReleaseNamer.GetRandomCandy:input_type -> google.protobuf.Empty
	1, // 3: releasenamer.ReleaseNamer.GetCandies:output_type -> releasenamer.CandyResponse
	1, // 4: releasenamer.ReleaseNamer.GetRandomCandy:output_type -> releasenamer.CandyResponse
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_release_namer_proto_init() }
func file_release_namer_proto_init() {
	if File_release_namer_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_release_namer_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
		file_release_namer_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CandyResponse); i {
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
		file_release_namer_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Candy); i {
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
			RawDescriptor: file_release_namer_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_release_namer_proto_goTypes,
		DependencyIndexes: file_release_namer_proto_depIdxs,
		MessageInfos:      file_release_namer_proto_msgTypes,
	}.Build()
	File_release_namer_proto = out.File
	file_release_namer_proto_rawDesc = nil
	file_release_namer_proto_goTypes = nil
	file_release_namer_proto_depIdxs = nil
}
