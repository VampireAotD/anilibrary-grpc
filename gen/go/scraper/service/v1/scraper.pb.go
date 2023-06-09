// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        (unknown)
// source: scraper/service/v1/scraper.proto

package scraper

import (
	v1 "github.com/VampireAotD/anilibrary-grpc-gateway/gen/go/scraper/model/v1"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type ScrapeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *ScrapeRequest) Reset() {
	*x = ScrapeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_scraper_service_v1_scraper_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ScrapeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ScrapeRequest) ProtoMessage() {}

func (x *ScrapeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_scraper_service_v1_scraper_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ScrapeRequest.ProtoReflect.Descriptor instead.
func (*ScrapeRequest) Descriptor() ([]byte, []int) {
	return file_scraper_service_v1_scraper_proto_rawDescGZIP(), []int{0}
}

func (x *ScrapeRequest) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

type ScrapeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Anime *v1.Anime `protobuf:"bytes,1,opt,name=anime,proto3" json:"anime,omitempty"`
}

func (x *ScrapeResponse) Reset() {
	*x = ScrapeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_scraper_service_v1_scraper_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ScrapeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ScrapeResponse) ProtoMessage() {}

func (x *ScrapeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_scraper_service_v1_scraper_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ScrapeResponse.ProtoReflect.Descriptor instead.
func (*ScrapeResponse) Descriptor() ([]byte, []int) {
	return file_scraper_service_v1_scraper_proto_rawDescGZIP(), []int{1}
}

func (x *ScrapeResponse) GetAnime() *v1.Anime {
	if x != nil {
		return x.Anime
	}
	return nil
}

var File_scraper_service_v1_scraper_proto protoreflect.FileDescriptor

var file_scraper_service_v1_scraper_proto_rawDesc = []byte{
	0x0a, 0x20, 0x73, 0x63, 0x72, 0x61, 0x70, 0x65, 0x72, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x63, 0x72, 0x61, 0x70, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x12, 0x73, 0x63, 0x72, 0x61, 0x70, 0x65, 0x72, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65,
	0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x73,
	0x63, 0x72, 0x61, 0x70, 0x65, 0x72, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x76, 0x31, 0x2f,
	0x61, 0x6e, 0x69, 0x6d, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x59, 0x0a, 0x0d, 0x53,
	0x63, 0x72, 0x61, 0x70, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x48, 0x0a, 0x03,
	0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x36, 0xfa, 0x42, 0x33, 0x72, 0x31,
	0x32, 0x2f, 0x28, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x61, 0x6e, 0x69, 0x6d, 0x65,
	0x67, 0x6f, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x7c, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f,
	0x61, 0x6e, 0x69, 0x6d, 0x65, 0x76, 0x6f, 0x73, 0x74, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x29, 0x2e,
	0x2b, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x22, 0x3f, 0x0a, 0x0e, 0x53, 0x63, 0x72, 0x61, 0x70, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2d, 0x0a, 0x05, 0x61, 0x6e, 0x69, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x73, 0x63, 0x72, 0x61, 0x70, 0x65,
	0x72, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x6e, 0x69, 0x6d, 0x65,
	0x52, 0x05, 0x61, 0x6e, 0x69, 0x6d, 0x65, 0x32, 0x7e, 0x0a, 0x0e, 0x53, 0x63, 0x72, 0x61, 0x70,
	0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x6c, 0x0a, 0x06, 0x53, 0x63, 0x72,
	0x61, 0x70, 0x65, 0x12, 0x21, 0x2e, 0x73, 0x63, 0x72, 0x61, 0x70, 0x65, 0x72, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x63, 0x72, 0x61, 0x70, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x73, 0x63, 0x72, 0x61, 0x70, 0x65, 0x72,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x63, 0x72, 0x61,
	0x70, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1b, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x15, 0x3a, 0x01, 0x2a, 0x22, 0x10, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x6e, 0x69, 0x6d, 0x65,
	0x2f, 0x73, 0x63, 0x72, 0x61, 0x70, 0x65, 0x42, 0x52, 0x5a, 0x50, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x56, 0x61, 0x6d, 0x70, 0x69, 0x72, 0x65, 0x41, 0x6f, 0x74,
	0x44, 0x2f, 0x61, 0x6e, 0x69, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x2d, 0x67, 0x72, 0x70,
	0x63, 0x2d, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f,
	0x2f, 0x73, 0x63, 0x72, 0x61, 0x70, 0x65, 0x72, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2f, 0x76, 0x31, 0x3b, 0x73, 0x63, 0x72, 0x61, 0x70, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_scraper_service_v1_scraper_proto_rawDescOnce sync.Once
	file_scraper_service_v1_scraper_proto_rawDescData = file_scraper_service_v1_scraper_proto_rawDesc
)

func file_scraper_service_v1_scraper_proto_rawDescGZIP() []byte {
	file_scraper_service_v1_scraper_proto_rawDescOnce.Do(func() {
		file_scraper_service_v1_scraper_proto_rawDescData = protoimpl.X.CompressGZIP(file_scraper_service_v1_scraper_proto_rawDescData)
	})
	return file_scraper_service_v1_scraper_proto_rawDescData
}

var file_scraper_service_v1_scraper_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_scraper_service_v1_scraper_proto_goTypes = []interface{}{
	(*ScrapeRequest)(nil),  // 0: scraper.service.v1.ScrapeRequest
	(*ScrapeResponse)(nil), // 1: scraper.service.v1.ScrapeResponse
	(*v1.Anime)(nil),       // 2: scraper.model.v1.Anime
}
var file_scraper_service_v1_scraper_proto_depIdxs = []int32{
	2, // 0: scraper.service.v1.ScrapeResponse.anime:type_name -> scraper.model.v1.Anime
	0, // 1: scraper.service.v1.ScraperService.Scrape:input_type -> scraper.service.v1.ScrapeRequest
	1, // 2: scraper.service.v1.ScraperService.Scrape:output_type -> scraper.service.v1.ScrapeResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_scraper_service_v1_scraper_proto_init() }
func file_scraper_service_v1_scraper_proto_init() {
	if File_scraper_service_v1_scraper_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_scraper_service_v1_scraper_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ScrapeRequest); i {
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
		file_scraper_service_v1_scraper_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ScrapeResponse); i {
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
			RawDescriptor: file_scraper_service_v1_scraper_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_scraper_service_v1_scraper_proto_goTypes,
		DependencyIndexes: file_scraper_service_v1_scraper_proto_depIdxs,
		MessageInfos:      file_scraper_service_v1_scraper_proto_msgTypes,
	}.Build()
	File_scraper_service_v1_scraper_proto = out.File
	file_scraper_service_v1_scraper_proto_rawDesc = nil
	file_scraper_service_v1_scraper_proto_goTypes = nil
	file_scraper_service_v1_scraper_proto_depIdxs = nil
}
