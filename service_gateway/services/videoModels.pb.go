// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.9.1
// source: videoInfo.proto

package services

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

type VideoModel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: json:"file_id"
	FileId uint32 `protobuf:"varint,1,opt,name=FileId,proto3" json:"FileId,omitempty"`
	// @inject_tag: json:"file_name"
	FileName string `protobuf:"bytes,2,opt,name=FileName,proto3" json:"FileName,omitempty"`
	// @inject_tag: json:"file_original_name"
	FileOriginalName string `protobuf:"bytes,3,opt,name=FileOriginalName,proto3" json:"FileOriginalName,omitempty"`
	// @inject_tag: json:"file_path"
	FilePath string `protobuf:"bytes,4,opt,name=FilePath,proto3" json:"FilePath,omitempty"`
	// @inject_tag: json:"file_url"
	FileUrl string `protobuf:"bytes,5,opt,name=FileUrl,proto3" json:"FileUrl,omitempty"`
	// @inject_tag: json:"file_type"
	FileType string `protobuf:"bytes,6,opt,name=FileType,proto3" json:"FileType,omitempty"`
	// @inject_tag: json:"mime_type"
	MimeType string `protobuf:"bytes,7,opt,name=MimeType,proto3" json:"MimeType,omitempty"`
	// @inject_tag: json:"file_size"
	FileSize string `protobuf:"bytes,8,opt,name=FileSize,proto3" json:"FileSize,omitempty"`
	// @inject_tag: json:"file_status"
	FileStatus string `protobuf:"bytes,9,opt,name=FileStatus,proto3" json:"FileStatus,omitempty"`
	// @inject_tag: json:"update_time"
	UpdateTime int64 `protobuf:"varint,10,opt,name=UpdateTime,proto3" json:"UpdateTime,omitempty"`
	// @inject_tag: json:"create_time"
	CreateTime int64 `protobuf:"varint,11,opt,name=CreateTime,proto3" json:"CreateTime,omitempty"`
}

func (x *VideoModel) Reset() {
	*x = VideoModel{}
	if protoimpl.UnsafeEnabled {
		mi := &file_videoModels_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VideoModel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VideoModel) ProtoMessage() {}

func (x *VideoModel) ProtoReflect() protoreflect.Message {
	mi := &file_videoModels_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VideoModel.ProtoReflect.Descriptor instead.
func (*VideoModel) Descriptor() ([]byte, []int) {
	return file_videoModels_proto_rawDescGZIP(), []int{0}
}

func (x *VideoModel) GetFileId() uint32 {
	if x != nil {
		return x.FileId
	}
	return 0
}

func (x *VideoModel) GetFileName() string {
	if x != nil {
		return x.FileName
	}
	return ""
}

func (x *VideoModel) GetFileOriginalName() string {
	if x != nil {
		return x.FileOriginalName
	}
	return ""
}

func (x *VideoModel) GetFilePath() string {
	if x != nil {
		return x.FilePath
	}
	return ""
}

func (x *VideoModel) GetFileUrl() string {
	if x != nil {
		return x.FileUrl
	}
	return ""
}

func (x *VideoModel) GetFileType() string {
	if x != nil {
		return x.FileType
	}
	return ""
}

func (x *VideoModel) GetMimeType() string {
	if x != nil {
		return x.MimeType
	}
	return ""
}

func (x *VideoModel) GetFileSize() string {
	if x != nil {
		return x.FileSize
	}
	return ""
}

func (x *VideoModel) GetFileStatus() string {
	if x != nil {
		return x.FileStatus
	}
	return ""
}

func (x *VideoModel) GetUpdateTime() int64 {
	if x != nil {
		return x.UpdateTime
	}
	return 0
}

func (x *VideoModel) GetCreateTime() int64 {
	if x != nil {
		return x.CreateTime
	}
	return 0
}

var File_videoModels_proto protoreflect.FileDescriptor

var file_videoModels_proto_rawDesc = []byte{
	0x0a, 0x11, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x22, 0xd6, 0x02,
	0x0a, 0x0a, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x12, 0x16, 0x0a, 0x06,
	0x46, 0x69, 0x6c, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x46, 0x69,
	0x6c, 0x65, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x46, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x46, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x2a, 0x0a, 0x10, 0x46, 0x69, 0x6c, 0x65, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c,
	0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x46, 0x69, 0x6c, 0x65,
	0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x46, 0x69, 0x6c, 0x65, 0x50, 0x61, 0x74, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x46, 0x69, 0x6c, 0x65, 0x50, 0x61, 0x74, 0x68, 0x12, 0x18, 0x0a, 0x07, 0x46, 0x69, 0x6c, 0x65,
	0x55, 0x72, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x46, 0x69, 0x6c, 0x65, 0x55,
	0x72, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x46, 0x69, 0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x46, 0x69, 0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x4d, 0x69, 0x6d, 0x65, 0x54, 0x79, 0x70, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x4d, 0x69, 0x6d, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x46, 0x69,
	0x6c, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x46, 0x69,
	0x6c, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x46, 0x69, 0x6c, 0x65, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x46, 0x69, 0x6c, 0x65,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x54, 0x69, 0x6d, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x54, 0x69, 0x6d, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x42, 0x0b, 0x5a, 0x09, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_videoModels_proto_rawDescOnce sync.Once
	file_videoModels_proto_rawDescData = file_videoModels_proto_rawDesc
)

func file_videoModels_proto_rawDescGZIP() []byte {
	file_videoModels_proto_rawDescOnce.Do(func() {
		file_videoModels_proto_rawDescData = protoimpl.X.CompressGZIP(file_videoModels_proto_rawDescData)
	})
	return file_videoModels_proto_rawDescData
}

var file_videoModels_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_videoModels_proto_goTypes = []interface{}{
	(*VideoModel)(nil), // 0: services.VideoModel
}
var file_videoModels_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_videoModels_proto_init() }
func file_videoModels_proto_init() {
	if File_videoModels_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_videoModels_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VideoModel); i {
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
			RawDescriptor: file_videoModels_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_videoModels_proto_goTypes,
		DependencyIndexes: file_videoModels_proto_depIdxs,
		MessageInfos:      file_videoModels_proto_msgTypes,
	}.Build()
	File_videoModels_proto = out.File
	file_videoModels_proto_rawDesc = nil
	file_videoModels_proto_goTypes = nil
	file_videoModels_proto_depIdxs = nil
}
