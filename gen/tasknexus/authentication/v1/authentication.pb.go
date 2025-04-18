// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: tasknexus/authentication/v1/authentication.proto

package authenticationv1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type AuthenticateUserRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Username      string                 `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Passwordd     string                 `protobuf:"bytes,2,opt,name=passwordd,proto3" json:"passwordd,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AuthenticateUserRequest) Reset() {
	*x = AuthenticateUserRequest{}
	mi := &file_tasknexus_authentication_v1_authentication_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AuthenticateUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthenticateUserRequest) ProtoMessage() {}

func (x *AuthenticateUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tasknexus_authentication_v1_authentication_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthenticateUserRequest.ProtoReflect.Descriptor instead.
func (*AuthenticateUserRequest) Descriptor() ([]byte, []int) {
	return file_tasknexus_authentication_v1_authentication_proto_rawDescGZIP(), []int{0}
}

func (x *AuthenticateUserRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *AuthenticateUserRequest) GetPasswordd() string {
	if x != nil {
		return x.Passwordd
	}
	return ""
}

type AuthenticateUserResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Success       bool                   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Message       string                 `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Token         string                 `protobuf:"bytes,3,opt,name=token,proto3" json:"token,omitempty"` // JWT token for authenticated user
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AuthenticateUserResponse) Reset() {
	*x = AuthenticateUserResponse{}
	mi := &file_tasknexus_authentication_v1_authentication_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AuthenticateUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthenticateUserResponse) ProtoMessage() {}

func (x *AuthenticateUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tasknexus_authentication_v1_authentication_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthenticateUserResponse.ProtoReflect.Descriptor instead.
func (*AuthenticateUserResponse) Descriptor() ([]byte, []int) {
	return file_tasknexus_authentication_v1_authentication_proto_rawDescGZIP(), []int{1}
}

func (x *AuthenticateUserResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *AuthenticateUserResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *AuthenticateUserResponse) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

var File_tasknexus_authentication_v1_authentication_proto protoreflect.FileDescriptor

const file_tasknexus_authentication_v1_authentication_proto_rawDesc = "" +
	"\n" +
	"0tasknexus/authentication/v1/authentication.proto\x12\x1btasknexus.authentication.v1\"S\n" +
	"\x17AuthenticateUserRequest\x12\x1a\n" +
	"\busername\x18\x01 \x01(\tR\busername\x12\x1c\n" +
	"\tpasswordd\x18\x02 \x01(\tR\tpasswordd\"d\n" +
	"\x18AuthenticateUserResponse\x12\x18\n" +
	"\asuccess\x18\x01 \x01(\bR\asuccess\x12\x18\n" +
	"\amessage\x18\x02 \x01(\tR\amessage\x12\x14\n" +
	"\x05token\x18\x03 \x01(\tR\x05token2\x98\x01\n" +
	"\x15AuthenticationService\x12\x7f\n" +
	"\x10AuthenticateUser\x124.tasknexus.authentication.v1.AuthenticateUserRequest\x1a5.tasknexus.authentication.v1.AuthenticateUserResponseB\x9a\x02\n" +
	"\x1fcom.tasknexus.authentication.v1B\x13AuthenticationProtoP\x01ZTgithub.com/cnc-csku/task-nexus-proto-go/tasknexus/authentication/v1;authenticationv1\xa2\x02\x03TAX\xaa\x02\x1bTasknexus.Authentication.V1\xca\x02\x1bTasknexus\\Authentication\\V1\xe2\x02'Tasknexus\\Authentication\\V1\\GPBMetadata\xea\x02\x1dTasknexus::Authentication::V1b\x06proto3"

var (
	file_tasknexus_authentication_v1_authentication_proto_rawDescOnce sync.Once
	file_tasknexus_authentication_v1_authentication_proto_rawDescData []byte
)

func file_tasknexus_authentication_v1_authentication_proto_rawDescGZIP() []byte {
	file_tasknexus_authentication_v1_authentication_proto_rawDescOnce.Do(func() {
		file_tasknexus_authentication_v1_authentication_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_tasknexus_authentication_v1_authentication_proto_rawDesc), len(file_tasknexus_authentication_v1_authentication_proto_rawDesc)))
	})
	return file_tasknexus_authentication_v1_authentication_proto_rawDescData
}

var file_tasknexus_authentication_v1_authentication_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_tasknexus_authentication_v1_authentication_proto_goTypes = []any{
	(*AuthenticateUserRequest)(nil),  // 0: tasknexus.authentication.v1.AuthenticateUserRequest
	(*AuthenticateUserResponse)(nil), // 1: tasknexus.authentication.v1.AuthenticateUserResponse
}
var file_tasknexus_authentication_v1_authentication_proto_depIdxs = []int32{
	0, // 0: tasknexus.authentication.v1.AuthenticationService.AuthenticateUser:input_type -> tasknexus.authentication.v1.AuthenticateUserRequest
	1, // 1: tasknexus.authentication.v1.AuthenticationService.AuthenticateUser:output_type -> tasknexus.authentication.v1.AuthenticateUserResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_tasknexus_authentication_v1_authentication_proto_init() }
func file_tasknexus_authentication_v1_authentication_proto_init() {
	if File_tasknexus_authentication_v1_authentication_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_tasknexus_authentication_v1_authentication_proto_rawDesc), len(file_tasknexus_authentication_v1_authentication_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_tasknexus_authentication_v1_authentication_proto_goTypes,
		DependencyIndexes: file_tasknexus_authentication_v1_authentication_proto_depIdxs,
		MessageInfos:      file_tasknexus_authentication_v1_authentication_proto_msgTypes,
	}.Build()
	File_tasknexus_authentication_v1_authentication_proto = out.File
	file_tasknexus_authentication_v1_authentication_proto_goTypes = nil
	file_tasknexus_authentication_v1_authentication_proto_depIdxs = nil
}
