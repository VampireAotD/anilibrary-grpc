// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        (unknown)
// source: scraper/model/v1/anime.proto

package scraper_model

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

type Anime struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Image       string   `protobuf:"bytes,1,opt,name=image,proto3" json:"image,omitempty"`
	Title       string   `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Status      string   `protobuf:"bytes,3,opt,name=status,proto3" json:"status,omitempty"`
	Episodes    string   `protobuf:"bytes,4,opt,name=episodes,proto3" json:"episodes,omitempty"`
	Genres      []string `protobuf:"bytes,5,rep,name=genres,proto3" json:"genres,omitempty"`
	VoiceActing []string `protobuf:"bytes,6,rep,name=voice_acting,json=voiceActing,proto3" json:"voice_acting,omitempty"`
	Synonyms    []string `protobuf:"bytes,7,rep,name=synonyms,proto3" json:"synonyms,omitempty"`
	Rating      float32  `protobuf:"fixed32,8,opt,name=rating,proto3" json:"rating,omitempty"`
}

func (x *Anime) Reset() {
	*x = Anime{}
	if protoimpl.UnsafeEnabled {
		mi := &file_scraper_model_v1_anime_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Anime) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Anime) ProtoMessage() {}

func (x *Anime) ProtoReflect() protoreflect.Message {
	mi := &file_scraper_model_v1_anime_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Anime.ProtoReflect.Descriptor instead.
func (*Anime) Descriptor() ([]byte, []int) {
	return file_scraper_model_v1_anime_proto_rawDescGZIP(), []int{0}
}

func (x *Anime) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

func (x *Anime) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Anime) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *Anime) GetEpisodes() string {
	if x != nil {
		return x.Episodes
	}
	return ""
}

func (x *Anime) GetGenres() []string {
	if x != nil {
		return x.Genres
	}
	return nil
}

func (x *Anime) GetVoiceActing() []string {
	if x != nil {
		return x.VoiceActing
	}
	return nil
}

func (x *Anime) GetSynonyms() []string {
	if x != nil {
		return x.Synonyms
	}
	return nil
}

func (x *Anime) GetRating() float32 {
	if x != nil {
		return x.Rating
	}
	return 0
}

var File_scraper_model_v1_anime_proto protoreflect.FileDescriptor

var file_scraper_model_v1_anime_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x73, 0x63, 0x72, 0x61, 0x70, 0x65, 0x72, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f,
	0x76, 0x31, 0x2f, 0x61, 0x6e, 0x69, 0x6d, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10,
	0x73, 0x63, 0x72, 0x61, 0x70, 0x65, 0x72, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x76, 0x31,
	0x22, 0xd6, 0x01, 0x0a, 0x05, 0x41, 0x6e, 0x69, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6d,
	0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1a,
	0x0a, 0x08, 0x65, 0x70, 0x69, 0x73, 0x6f, 0x64, 0x65, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x65, 0x70, 0x69, 0x73, 0x6f, 0x64, 0x65, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x67, 0x65,
	0x6e, 0x72, 0x65, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x67, 0x65, 0x6e, 0x72,
	0x65, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x5f, 0x61, 0x63, 0x74, 0x69,
	0x6e, 0x67, 0x18, 0x06, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x41,
	0x63, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x79, 0x6e, 0x6f, 0x6e, 0x79, 0x6d,
	0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x73, 0x79, 0x6e, 0x6f, 0x6e, 0x79, 0x6d,
	0x73, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x02, 0x52, 0x06, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x42, 0x56, 0x5a, 0x54, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x56, 0x61, 0x6d, 0x70, 0x69, 0x72, 0x65, 0x41,
	0x6f, 0x74, 0x44, 0x2f, 0x61, 0x6e, 0x69, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x2d, 0x67,
	0x72, 0x70, 0x63, 0x2d, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2f, 0x67, 0x65, 0x6e, 0x2f,
	0x67, 0x6f, 0x2f, 0x73, 0x63, 0x72, 0x61, 0x70, 0x65, 0x72, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x2f, 0x76, 0x31, 0x3b, 0x73, 0x63, 0x72, 0x61, 0x70, 0x65, 0x72, 0x5f, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_scraper_model_v1_anime_proto_rawDescOnce sync.Once
	file_scraper_model_v1_anime_proto_rawDescData = file_scraper_model_v1_anime_proto_rawDesc
)

func file_scraper_model_v1_anime_proto_rawDescGZIP() []byte {
	file_scraper_model_v1_anime_proto_rawDescOnce.Do(func() {
		file_scraper_model_v1_anime_proto_rawDescData = protoimpl.X.CompressGZIP(file_scraper_model_v1_anime_proto_rawDescData)
	})
	return file_scraper_model_v1_anime_proto_rawDescData
}

var file_scraper_model_v1_anime_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_scraper_model_v1_anime_proto_goTypes = []interface{}{
	(*Anime)(nil), // 0: scraper.model.v1.Anime
}
var file_scraper_model_v1_anime_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_scraper_model_v1_anime_proto_init() }
func file_scraper_model_v1_anime_proto_init() {
	if File_scraper_model_v1_anime_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_scraper_model_v1_anime_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Anime); i {
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
			RawDescriptor: file_scraper_model_v1_anime_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_scraper_model_v1_anime_proto_goTypes,
		DependencyIndexes: file_scraper_model_v1_anime_proto_depIdxs,
		MessageInfos:      file_scraper_model_v1_anime_proto_msgTypes,
	}.Build()
	File_scraper_model_v1_anime_proto = out.File
	file_scraper_model_v1_anime_proto_rawDesc = nil
	file_scraper_model_v1_anime_proto_goTypes = nil
	file_scraper_model_v1_anime_proto_depIdxs = nil
}
