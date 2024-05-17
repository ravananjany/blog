// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: protos/blog/blog.proto

package blog

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

type Post struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PostId          string   `protobuf:"bytes,1,opt,name=post_id,json=postId,proto3" json:"post_id,omitempty"`
	Title           string   `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Content         string   `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
	Author          string   `protobuf:"bytes,4,opt,name=author,proto3" json:"author,omitempty"`
	PublicationData string   `protobuf:"bytes,5,opt,name=publication_data,json=publicationData,proto3" json:"publication_data,omitempty"`
	Tags            []string `protobuf:"bytes,6,rep,name=tags,proto3" json:"tags,omitempty"`
}

func (x *Post) Reset() {
	*x = Post{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_blog_blog_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Post) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Post) ProtoMessage() {}

func (x *Post) ProtoReflect() protoreflect.Message {
	mi := &file_protos_blog_blog_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Post.ProtoReflect.Descriptor instead.
func (*Post) Descriptor() ([]byte, []int) {
	return file_protos_blog_blog_proto_rawDescGZIP(), []int{0}
}

func (x *Post) GetPostId() string {
	if x != nil {
		return x.PostId
	}
	return ""
}

func (x *Post) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Post) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *Post) GetAuthor() string {
	if x != nil {
		return x.Author
	}
	return ""
}

func (x *Post) GetPublicationData() string {
	if x != nil {
		return x.PublicationData
	}
	return ""
}

func (x *Post) GetTags() []string {
	if x != nil {
		return x.Tags
	}
	return nil
}

type Datas struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Posts []*Post `protobuf:"bytes,1,rep,name=Posts,proto3" json:"Posts,omitempty"`
}

func (x *Datas) Reset() {
	*x = Datas{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_blog_blog_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Datas) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Datas) ProtoMessage() {}

func (x *Datas) ProtoReflect() protoreflect.Message {
	mi := &file_protos_blog_blog_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Datas.ProtoReflect.Descriptor instead.
func (*Datas) Descriptor() ([]byte, []int) {
	return file_protos_blog_blog_proto_rawDescGZIP(), []int{1}
}

func (x *Datas) GetPosts() []*Post {
	if x != nil {
		return x.Posts
	}
	return nil
}

type Readall struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Readall) Reset() {
	*x = Readall{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_blog_blog_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Readall) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Readall) ProtoMessage() {}

func (x *Readall) ProtoReflect() protoreflect.Message {
	mi := &file_protos_blog_blog_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Readall.ProtoReflect.Descriptor instead.
func (*Readall) Descriptor() ([]byte, []int) {
	return file_protos_blog_blog_proto_rawDescGZIP(), []int{2}
}

type Id struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PostId string `protobuf:"bytes,1,opt,name=post_id,json=postId,proto3" json:"post_id,omitempty"`
}

func (x *Id) Reset() {
	*x = Id{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_blog_blog_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Id) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Id) ProtoMessage() {}

func (x *Id) ProtoReflect() protoreflect.Message {
	mi := &file_protos_blog_blog_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Id.ProtoReflect.Descriptor instead.
func (*Id) Descriptor() ([]byte, []int) {
	return file_protos_blog_blog_proto_rawDescGZIP(), []int{3}
}

func (x *Id) GetPostId() string {
	if x != nil {
		return x.PostId
	}
	return ""
}

type DeleteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *DeleteResponse) Reset() {
	*x = DeleteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_blog_blog_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteResponse) ProtoMessage() {}

func (x *DeleteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_blog_blog_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteResponse.ProtoReflect.Descriptor instead.
func (*DeleteResponse) Descriptor() ([]byte, []int) {
	return file_protos_blog_blog_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_protos_blog_blog_proto protoreflect.FileDescriptor

var file_protos_blog_blog_proto_rawDesc = []byte{
	0x0a, 0x16, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x62, 0x6c, 0x6f, 0x67, 0x2f, 0x62, 0x6c,
	0x6f, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x62, 0x6c, 0x6f, 0x67, 0x22, 0xa6,
	0x01, 0x0a, 0x04, 0x50, 0x6f, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x70, 0x6f, 0x73, 0x74, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x6f, 0x73, 0x74, 0x49, 0x64,
	0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x12, 0x16, 0x0a, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x12, 0x29, 0x0a, 0x10, 0x70, 0x75, 0x62, 0x6c,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44,
	0x61, 0x74, 0x61, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x22, 0x29, 0x0a, 0x05, 0x44, 0x61, 0x74, 0x61, 0x73,
	0x12, 0x20, 0x0a, 0x05, 0x50, 0x6f, 0x73, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x0a, 0x2e, 0x62, 0x6c, 0x6f, 0x67, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x05, 0x50, 0x6f, 0x73,
	0x74, 0x73, 0x22, 0x09, 0x0a, 0x07, 0x52, 0x65, 0x61, 0x64, 0x61, 0x6c, 0x6c, 0x22, 0x1d, 0x0a,
	0x02, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x70, 0x6f, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x6f, 0x73, 0x74, 0x49, 0x64, 0x22, 0x2a, 0x0a, 0x0e,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0xdd, 0x01, 0x0a, 0x0b, 0x42, 0x6c, 0x6f,
	0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x28, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x42, 0x6c, 0x6f, 0x67, 0x50, 0x6f, 0x73, 0x74, 0x12, 0x0a, 0x2e, 0x62, 0x6c, 0x6f,
	0x67, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x1a, 0x0a, 0x2e, 0x62, 0x6c, 0x6f, 0x67, 0x2e, 0x50, 0x6f,
	0x73, 0x74, 0x12, 0x24, 0x0a, 0x0c, 0x52, 0x65, 0x61, 0x64, 0x42, 0x6c, 0x6f, 0x67, 0x50, 0x6f,
	0x73, 0x74, 0x12, 0x08, 0x2e, 0x62, 0x6c, 0x6f, 0x67, 0x2e, 0x49, 0x64, 0x1a, 0x0a, 0x2e, 0x62,
	0x6c, 0x6f, 0x67, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x12, 0x28, 0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x42, 0x6c, 0x6f, 0x67, 0x50, 0x6f, 0x73, 0x74, 0x12, 0x0a, 0x2e, 0x62, 0x6c, 0x6f,
	0x67, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x1a, 0x0a, 0x2e, 0x62, 0x6c, 0x6f, 0x67, 0x2e, 0x50, 0x6f,
	0x73, 0x74, 0x12, 0x30, 0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x6c, 0x6f, 0x67,
	0x50, 0x6f, 0x73, 0x74, 0x12, 0x08, 0x2e, 0x62, 0x6c, 0x6f, 0x67, 0x2e, 0x49, 0x64, 0x1a, 0x14,
	0x2e, 0x62, 0x6c, 0x6f, 0x67, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x22, 0x0a, 0x04, 0x52, 0x65, 0x61, 0x64, 0x12, 0x0d, 0x2e, 0x62,
	0x6c, 0x6f, 0x67, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x61, 0x6c, 0x6c, 0x1a, 0x0b, 0x2e, 0x62, 0x6c,
	0x6f, 0x67, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x73, 0x42, 0x22, 0x5a, 0x20, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6a, 0x61, 0x6e, 0x79, 0x2f, 0x62, 0x6c, 0x6f, 0x67,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x62, 0x6c, 0x6f, 0x67, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_blog_blog_proto_rawDescOnce sync.Once
	file_protos_blog_blog_proto_rawDescData = file_protos_blog_blog_proto_rawDesc
)

func file_protos_blog_blog_proto_rawDescGZIP() []byte {
	file_protos_blog_blog_proto_rawDescOnce.Do(func() {
		file_protos_blog_blog_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_blog_blog_proto_rawDescData)
	})
	return file_protos_blog_blog_proto_rawDescData
}

var file_protos_blog_blog_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_protos_blog_blog_proto_goTypes = []interface{}{
	(*Post)(nil),           // 0: blog.Post
	(*Datas)(nil),          // 1: blog.Datas
	(*Readall)(nil),        // 2: blog.Readall
	(*Id)(nil),             // 3: blog.Id
	(*DeleteResponse)(nil), // 4: blog.DeleteResponse
}
var file_protos_blog_blog_proto_depIdxs = []int32{
	0, // 0: blog.Datas.Posts:type_name -> blog.Post
	0, // 1: blog.BlogService.CreateBlogPost:input_type -> blog.Post
	3, // 2: blog.BlogService.ReadBlogPost:input_type -> blog.Id
	0, // 3: blog.BlogService.UpdateBlogPost:input_type -> blog.Post
	3, // 4: blog.BlogService.DeleteBlogPost:input_type -> blog.Id
	2, // 5: blog.BlogService.Read:input_type -> blog.Readall
	0, // 6: blog.BlogService.CreateBlogPost:output_type -> blog.Post
	0, // 7: blog.BlogService.ReadBlogPost:output_type -> blog.Post
	0, // 8: blog.BlogService.UpdateBlogPost:output_type -> blog.Post
	4, // 9: blog.BlogService.DeleteBlogPost:output_type -> blog.DeleteResponse
	1, // 10: blog.BlogService.Read:output_type -> blog.Datas
	6, // [6:11] is the sub-list for method output_type
	1, // [1:6] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_protos_blog_blog_proto_init() }
func file_protos_blog_blog_proto_init() {
	if File_protos_blog_blog_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_blog_blog_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Post); i {
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
		file_protos_blog_blog_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Datas); i {
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
		file_protos_blog_blog_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Readall); i {
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
		file_protos_blog_blog_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Id); i {
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
		file_protos_blog_blog_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteResponse); i {
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
			RawDescriptor: file_protos_blog_blog_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_blog_blog_proto_goTypes,
		DependencyIndexes: file_protos_blog_blog_proto_depIdxs,
		MessageInfos:      file_protos_blog_blog_proto_msgTypes,
	}.Build()
	File_protos_blog_blog_proto = out.File
	file_protos_blog_blog_proto_rawDesc = nil
	file_protos_blog_blog_proto_goTypes = nil
	file_protos_blog_blog_proto_depIdxs = nil
}
