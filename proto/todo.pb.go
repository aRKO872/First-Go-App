// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.25.1
// source: proto/todo.proto

package todo

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

type NewTodo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *NewTodo) Reset() {
	*x = NewTodo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_todo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewTodo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewTodo) ProtoMessage() {}

func (x *NewTodo) ProtoReflect() protoreflect.Message {
	mi := &file_proto_todo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewTodo.ProtoReflect.Descriptor instead.
func (*NewTodo) Descriptor() ([]byte, []int) {
	return file_proto_todo_proto_rawDescGZIP(), []int{0}
}

func (x *NewTodo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *NewTodo) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type Todos struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Done        bool   `protobuf:"varint,4,opt,name=done,proto3" json:"done,omitempty"`
}

func (x *Todos) Reset() {
	*x = Todos{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_todo_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Todos) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Todos) ProtoMessage() {}

func (x *Todos) ProtoReflect() protoreflect.Message {
	mi := &file_proto_todo_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Todos.ProtoReflect.Descriptor instead.
func (*Todos) Descriptor() ([]byte, []int) {
	return file_proto_todo_proto_rawDescGZIP(), []int{1}
}

func (x *Todos) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Todos) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Todos) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Todos) GetDone() bool {
	if x != nil {
		return x.Done
	}
	return false
}

type EditOrDeleteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *EditOrDeleteRequest) Reset() {
	*x = EditOrDeleteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_todo_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EditOrDeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EditOrDeleteRequest) ProtoMessage() {}

func (x *EditOrDeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_todo_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EditOrDeleteRequest.ProtoReflect.Descriptor instead.
func (*EditOrDeleteRequest) Descriptor() ([]byte, []int) {
	return file_proto_todo_proto_rawDescGZIP(), []int{2}
}

func (x *EditOrDeleteRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type EmptyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EmptyRequest) Reset() {
	*x = EmptyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_todo_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmptyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmptyRequest) ProtoMessage() {}

func (x *EmptyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_todo_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmptyRequest.ProtoReflect.Descriptor instead.
func (*EmptyRequest) Descriptor() ([]byte, []int) {
	return file_proto_todo_proto_rawDescGZIP(), []int{3}
}

type ArrayOfTodos struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TodoArr []*Todos `protobuf:"bytes,1,rep,name=todoArr,proto3" json:"todoArr,omitempty"`
}

func (x *ArrayOfTodos) Reset() {
	*x = ArrayOfTodos{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_todo_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ArrayOfTodos) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ArrayOfTodos) ProtoMessage() {}

func (x *ArrayOfTodos) ProtoReflect() protoreflect.Message {
	mi := &file_proto_todo_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ArrayOfTodos.ProtoReflect.Descriptor instead.
func (*ArrayOfTodos) Descriptor() ([]byte, []int) {
	return file_proto_todo_proto_rawDescGZIP(), []int{4}
}

func (x *ArrayOfTodos) GetTodoArr() []*Todos {
	if x != nil {
		return x.TodoArr
	}
	return nil
}

var File_proto_todo_proto protoreflect.FileDescriptor

var file_proto_todo_proto_rawDesc = []byte{
	0x0a, 0x10, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x6f, 0x64, 0x6f, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x04, 0x74, 0x6f, 0x64, 0x6f, 0x22, 0x3f, 0x0a, 0x07, 0x4e, 0x65, 0x77, 0x54,
	0x6f, 0x64, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x61, 0x0a, 0x05, 0x54, 0x6f, 0x64,
	0x6f, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x6f, 0x6e, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x64, 0x6f, 0x6e, 0x65, 0x22, 0x25, 0x0a, 0x13,
	0x45, 0x64, 0x69, 0x74, 0x4f, 0x72, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x22, 0x0e, 0x0a, 0x0c, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x22, 0x35, 0x0a, 0x0c, 0x41, 0x72, 0x72, 0x61, 0x79, 0x4f, 0x66, 0x54, 0x6f,
	0x64, 0x6f, 0x73, 0x12, 0x25, 0x0a, 0x07, 0x74, 0x6f, 0x64, 0x6f, 0x41, 0x72, 0x72, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x2e, 0x54, 0x6f, 0x64, 0x6f,
	0x73, 0x52, 0x07, 0x74, 0x6f, 0x64, 0x6f, 0x41, 0x72, 0x72, 0x32, 0xe6, 0x01, 0x0a, 0x0b, 0x54,
	0x6f, 0x64, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x28, 0x0a, 0x0a, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x64, 0x6f, 0x12, 0x0d, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x2e,
	0x4e, 0x65, 0x77, 0x54, 0x6f, 0x64, 0x6f, 0x1a, 0x0b, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x2e, 0x54,
	0x6f, 0x64, 0x6f, 0x73, 0x12, 0x37, 0x0a, 0x0d, 0x46, 0x65, 0x74, 0x63, 0x68, 0x41, 0x6c, 0x6c,
	0x54, 0x6f, 0x64, 0x6f, 0x73, 0x12, 0x12, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x74, 0x6f, 0x64, 0x6f,
	0x2e, 0x41, 0x72, 0x72, 0x61, 0x79, 0x4f, 0x66, 0x54, 0x6f, 0x64, 0x6f, 0x73, 0x12, 0x38, 0x0a,
	0x0e, 0x45, 0x64, 0x69, 0x74, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x54, 0x6f, 0x64, 0x6f, 0x12,
	0x19, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x2e, 0x45, 0x64, 0x69, 0x74, 0x4f, 0x72, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0b, 0x2e, 0x74, 0x6f, 0x64,
	0x6f, 0x2e, 0x54, 0x6f, 0x64, 0x6f, 0x73, 0x12, 0x3a, 0x0a, 0x10, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x54, 0x6f, 0x64, 0x6f, 0x12, 0x19, 0x2e, 0x74, 0x6f,
	0x64, 0x6f, 0x2e, 0x45, 0x64, 0x69, 0x74, 0x4f, 0x72, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0b, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x2e, 0x54, 0x6f,
	0x64, 0x6f, 0x73, 0x42, 0x2c, 0x5a, 0x2a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x61, 0x52, 0x4b, 0x4f, 0x38, 0x37, 0x32, 0x2f, 0x66, 0x69, 0x72, 0x73, 0x74, 0x2d,
	0x67, 0x6f, 0x2d, 0x61, 0x70, 0x70, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x6f, 0x64,
	0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_todo_proto_rawDescOnce sync.Once
	file_proto_todo_proto_rawDescData = file_proto_todo_proto_rawDesc
)

func file_proto_todo_proto_rawDescGZIP() []byte {
	file_proto_todo_proto_rawDescOnce.Do(func() {
		file_proto_todo_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_todo_proto_rawDescData)
	})
	return file_proto_todo_proto_rawDescData
}

var file_proto_todo_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_proto_todo_proto_goTypes = []interface{}{
	(*NewTodo)(nil),             // 0: todo.NewTodo
	(*Todos)(nil),               // 1: todo.Todos
	(*EditOrDeleteRequest)(nil), // 2: todo.EditOrDeleteRequest
	(*EmptyRequest)(nil),        // 3: todo.EmptyRequest
	(*ArrayOfTodos)(nil),        // 4: todo.ArrayOfTodos
}
var file_proto_todo_proto_depIdxs = []int32{
	1, // 0: todo.ArrayOfTodos.todoArr:type_name -> todo.Todos
	0, // 1: todo.TodoService.CreateTodo:input_type -> todo.NewTodo
	3, // 2: todo.TodoService.FetchAllTodos:input_type -> todo.EmptyRequest
	2, // 3: todo.TodoService.EditSingleTodo:input_type -> todo.EditOrDeleteRequest
	2, // 4: todo.TodoService.DeleteSingleTodo:input_type -> todo.EditOrDeleteRequest
	1, // 5: todo.TodoService.CreateTodo:output_type -> todo.Todos
	4, // 6: todo.TodoService.FetchAllTodos:output_type -> todo.ArrayOfTodos
	1, // 7: todo.TodoService.EditSingleTodo:output_type -> todo.Todos
	1, // 8: todo.TodoService.DeleteSingleTodo:output_type -> todo.Todos
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_todo_proto_init() }
func file_proto_todo_proto_init() {
	if File_proto_todo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_todo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewTodo); i {
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
		file_proto_todo_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Todos); i {
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
		file_proto_todo_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EditOrDeleteRequest); i {
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
		file_proto_todo_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmptyRequest); i {
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
		file_proto_todo_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ArrayOfTodos); i {
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
			RawDescriptor: file_proto_todo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_todo_proto_goTypes,
		DependencyIndexes: file_proto_todo_proto_depIdxs,
		MessageInfos:      file_proto_todo_proto_msgTypes,
	}.Build()
	File_proto_todo_proto = out.File
	file_proto_todo_proto_rawDesc = nil
	file_proto_todo_proto_goTypes = nil
	file_proto_todo_proto_depIdxs = nil
}
