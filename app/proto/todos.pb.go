// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.15.8
// source: protoc/todos.proto

package protoc

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

type AddTodoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Content string `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`
	Source  string `protobuf:"bytes,2,opt,name=source,proto3" json:"source,omitempty"`
	Size    string `protobuf:"bytes,3,opt,name=size,proto3" json:"size,omitempty"`
	Format  string `protobuf:"bytes,4,opt,name=format,proto3" json:"format,omitempty"`
}

func (x *AddTodoRequest) Reset() {
	*x = AddTodoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protoc_todos_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddTodoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddTodoRequest) ProtoMessage() {}

func (x *AddTodoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protoc_todos_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddTodoRequest.ProtoReflect.Descriptor instead.
func (*AddTodoRequest) Descriptor() ([]byte, []int) {
	return file_protoc_todos_proto_rawDescGZIP(), []int{0}
}

func (x *AddTodoRequest) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *AddTodoRequest) GetSource() string {
	if x != nil {
		return x.Source
	}
	return ""
}

func (x *AddTodoRequest) GetSize() string {
	if x != nil {
		return x.Size
	}
	return ""
}

func (x *AddTodoRequest) GetFormat() string {
	if x != nil {
		return x.Format
	}
	return ""
}

type Todo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string    `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Content    string    `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
	LastUpdate string    `protobuf:"bytes,3,opt,name=last_update,json=lastUpdate,proto3" json:"last_update,omitempty"`
	Metadata   *TodoMeta `protobuf:"bytes,4,opt,name=metadata,proto3" json:"metadata,omitempty"`
}

func (x *Todo) Reset() {
	*x = Todo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protoc_todos_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Todo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Todo) ProtoMessage() {}

func (x *Todo) ProtoReflect() protoreflect.Message {
	mi := &file_protoc_todos_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Todo.ProtoReflect.Descriptor instead.
func (*Todo) Descriptor() ([]byte, []int) {
	return file_protoc_todos_proto_rawDescGZIP(), []int{1}
}

func (x *Todo) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Todo) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *Todo) GetLastUpdate() string {
	if x != nil {
		return x.LastUpdate
	}
	return ""
}

func (x *Todo) GetMetadata() *TodoMeta {
	if x != nil {
		return x.Metadata
	}
	return nil
}

type TodoMeta struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Source string `protobuf:"bytes,1,opt,name=source,proto3" json:"source,omitempty"`
	Size   string `protobuf:"bytes,2,opt,name=size,proto3" json:"size,omitempty"`
	Format string `protobuf:"bytes,3,opt,name=format,proto3" json:"format,omitempty"`
}

func (x *TodoMeta) Reset() {
	*x = TodoMeta{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protoc_todos_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TodoMeta) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TodoMeta) ProtoMessage() {}

func (x *TodoMeta) ProtoReflect() protoreflect.Message {
	mi := &file_protoc_todos_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TodoMeta.ProtoReflect.Descriptor instead.
func (*TodoMeta) Descriptor() ([]byte, []int) {
	return file_protoc_todos_proto_rawDescGZIP(), []int{2}
}

func (x *TodoMeta) GetSource() string {
	if x != nil {
		return x.Source
	}
	return ""
}

func (x *TodoMeta) GetSize() string {
	if x != nil {
		return x.Size
	}
	return ""
}

func (x *TodoMeta) GetFormat() string {
	if x != nil {
		return x.Format
	}
	return ""
}

type AddTodoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result *Todo `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *AddTodoResponse) Reset() {
	*x = AddTodoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protoc_todos_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddTodoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddTodoResponse) ProtoMessage() {}

func (x *AddTodoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protoc_todos_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddTodoResponse.ProtoReflect.Descriptor instead.
func (*AddTodoResponse) Descriptor() ([]byte, []int) {
	return file_protoc_todos_proto_rawDescGZIP(), []int{3}
}

func (x *AddTodoResponse) GetResult() *Todo {
	if x != nil {
		return x.Result
	}
	return nil
}

type GetTodosRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Size   int32 `protobuf:"varint,1,opt,name=size,proto3" json:"size,omitempty"`
	Offset int32 `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
}

func (x *GetTodosRequest) Reset() {
	*x = GetTodosRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protoc_todos_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTodosRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTodosRequest) ProtoMessage() {}

func (x *GetTodosRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protoc_todos_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTodosRequest.ProtoReflect.Descriptor instead.
func (*GetTodosRequest) Descriptor() ([]byte, []int) {
	return file_protoc_todos_proto_rawDescGZIP(), []int{4}
}

func (x *GetTodosRequest) GetSize() int32 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *GetTodosRequest) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

type GetTodosResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Results []*Todo `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"`
}

func (x *GetTodosResponse) Reset() {
	*x = GetTodosResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protoc_todos_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTodosResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTodosResponse) ProtoMessage() {}

func (x *GetTodosResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protoc_todos_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTodosResponse.ProtoReflect.Descriptor instead.
func (*GetTodosResponse) Descriptor() ([]byte, []int) {
	return file_protoc_todos_proto_rawDescGZIP(), []int{5}
}

func (x *GetTodosResponse) GetResults() []*Todo {
	if x != nil {
		return x.Results
	}
	return nil
}

var File_protoc_todos_proto protoreflect.FileDescriptor

var file_protoc_todos_proto_rawDesc = []byte{
	0x0a, 0x12, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2f, 0x74, 0x6f, 0x64, 0x6f, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x65, 0x73, 0x6d, 0x6f, 0x6e, 0x67, 0x6f, 0x22, 0x6e, 0x0a,
	0x0e, 0x41, 0x64, 0x64, 0x54, 0x6f, 0x64, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x22, 0x80, 0x01,
	0x0a, 0x04, 0x54, 0x6f, 0x64, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x12, 0x1f, 0x0a, 0x0b, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6c, 0x61, 0x73, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x12, 0x2d, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x65, 0x73, 0x6d, 0x6f, 0x6e, 0x67, 0x6f, 0x2e, 0x54, 0x6f,
	0x64, 0x6f, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x22, 0x4e, 0x0a, 0x08, 0x54, 0x6f, 0x64, 0x6f, 0x4d, 0x65, 0x74, 0x61, 0x12, 0x16, 0x0a, 0x06,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x6f, 0x72, 0x6d,
	0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74,
	0x22, 0x38, 0x0a, 0x0f, 0x41, 0x64, 0x64, 0x54, 0x6f, 0x64, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x25, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x65, 0x73, 0x6d, 0x6f, 0x6e, 0x67, 0x6f, 0x2e, 0x54, 0x6f,
	0x64, 0x6f, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x3d, 0x0a, 0x0f, 0x47, 0x65,
	0x74, 0x54, 0x6f, 0x64, 0x6f, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x73, 0x69, 0x7a,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x22, 0x3b, 0x0a, 0x10, 0x47, 0x65, 0x74,
	0x54, 0x6f, 0x64, 0x6f, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x27, 0x0a,
	0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d,
	0x2e, 0x65, 0x73, 0x6d, 0x6f, 0x6e, 0x67, 0x6f, 0x2e, 0x54, 0x6f, 0x64, 0x6f, 0x52, 0x07, 0x72,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x32, 0x8d, 0x01, 0x0a, 0x0c, 0x54, 0x6f, 0x64, 0x6f, 0x73,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3c, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x54, 0x6f,
	0x64, 0x6f, 0x12, 0x17, 0x2e, 0x65, 0x73, 0x6d, 0x6f, 0x6e, 0x67, 0x6f, 0x2e, 0x41, 0x64, 0x64,
	0x54, 0x6f, 0x64, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x65, 0x73,
	0x6d, 0x6f, 0x6e, 0x67, 0x6f, 0x2e, 0x41, 0x64, 0x64, 0x54, 0x6f, 0x64, 0x6f, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3f, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x64, 0x6f,
	0x73, 0x12, 0x18, 0x2e, 0x65, 0x73, 0x6d, 0x6f, 0x6e, 0x67, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x54,
	0x6f, 0x64, 0x6f, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x65, 0x73,
	0x6d, 0x6f, 0x6e, 0x67, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x64, 0x6f, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protoc_todos_proto_rawDescOnce sync.Once
	file_protoc_todos_proto_rawDescData = file_protoc_todos_proto_rawDesc
)

func file_protoc_todos_proto_rawDescGZIP() []byte {
	file_protoc_todos_proto_rawDescOnce.Do(func() {
		file_protoc_todos_proto_rawDescData = protoimpl.X.CompressGZIP(file_protoc_todos_proto_rawDescData)
	})
	return file_protoc_todos_proto_rawDescData
}

var file_protoc_todos_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_protoc_todos_proto_goTypes = []interface{}{
	(*AddTodoRequest)(nil),   // 0: esmongo.AddTodoRequest
	(*Todo)(nil),             // 1: esmongo.Todo
	(*TodoMeta)(nil),         // 2: esmongo.TodoMeta
	(*AddTodoResponse)(nil),  // 3: esmongo.AddTodoResponse
	(*GetTodosRequest)(nil),  // 4: esmongo.GetTodosRequest
	(*GetTodosResponse)(nil), // 5: esmongo.GetTodosResponse
}
var file_protoc_todos_proto_depIdxs = []int32{
	2, // 0: esmongo.Todo.metadata:type_name -> esmongo.TodoMeta
	1, // 1: esmongo.AddTodoResponse.result:type_name -> esmongo.Todo
	1, // 2: esmongo.GetTodosResponse.results:type_name -> esmongo.Todo
	0, // 3: esmongo.TodosService.AddTodo:input_type -> esmongo.AddTodoRequest
	4, // 4: esmongo.TodosService.GetTodos:input_type -> esmongo.GetTodosRequest
	3, // 5: esmongo.TodosService.AddTodo:output_type -> esmongo.AddTodoResponse
	5, // 6: esmongo.TodosService.GetTodos:output_type -> esmongo.GetTodosResponse
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_protoc_todos_proto_init() }
func file_protoc_todos_proto_init() {
	if File_protoc_todos_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protoc_todos_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddTodoRequest); i {
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
		file_protoc_todos_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Todo); i {
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
		file_protoc_todos_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TodoMeta); i {
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
		file_protoc_todos_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddTodoResponse); i {
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
		file_protoc_todos_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTodosRequest); i {
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
		file_protoc_todos_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTodosResponse); i {
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
			RawDescriptor: file_protoc_todos_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protoc_todos_proto_goTypes,
		DependencyIndexes: file_protoc_todos_proto_depIdxs,
		MessageInfos:      file_protoc_todos_proto_msgTypes,
	}.Build()
	File_protoc_todos_proto = out.File
	file_protoc_todos_proto_rawDesc = nil
	file_protoc_todos_proto_goTypes = nil
	file_protoc_todos_proto_depIdxs = nil
}
