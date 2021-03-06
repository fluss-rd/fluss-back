// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: app/river-management/handlers/grpc/service.proto

package grpchandler

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

type UpdateModuleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ModuleID string `protobuf:"bytes,1,opt,name=moduleID,proto3" json:"moduleID,omitempty"`
	Status   string `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *UpdateModuleRequest) Reset() {
	*x = UpdateModuleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_river_management_handlers_grpc_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateModuleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateModuleRequest) ProtoMessage() {}

func (x *UpdateModuleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_app_river_management_handlers_grpc_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateModuleRequest.ProtoReflect.Descriptor instead.
func (*UpdateModuleRequest) Descriptor() ([]byte, []int) {
	return file_app_river_management_handlers_grpc_service_proto_rawDescGZIP(), []int{0}
}

func (x *UpdateModuleRequest) GetModuleID() string {
	if x != nil {
		return x.ModuleID
	}
	return ""
}

func (x *UpdateModuleRequest) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type GetModuleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PhoneNumber string `protobuf:"bytes,1,opt,name=phoneNumber,proto3" json:"phoneNumber,omitempty"`
}

func (x *GetModuleRequest) Reset() {
	*x = GetModuleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_river_management_handlers_grpc_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetModuleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetModuleRequest) ProtoMessage() {}

func (x *GetModuleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_app_river_management_handlers_grpc_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetModuleRequest.ProtoReflect.Descriptor instead.
func (*GetModuleRequest) Descriptor() ([]byte, []int) {
	return file_app_river_management_handlers_grpc_service_proto_rawDescGZIP(), []int{1}
}

func (x *GetModuleRequest) GetPhoneNumber() string {
	if x != nil {
		return x.PhoneNumber
	}
	return ""
}

type Module struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ModuleID     string `protobuf:"bytes,1,opt,name=moduleID,proto3" json:"moduleID,omitempty"`
	PhoneNumber  string `protobuf:"bytes,2,opt,name=phoneNumber,proto3" json:"phoneNumber,omitempty"`
	Alias        string `protobuf:"bytes,3,opt,name=alias,proto3" json:"alias,omitempty"`
	RiverID      string `protobuf:"bytes,4,opt,name=riverID,proto3" json:"riverID,omitempty"`
	RiverName    string `protobuf:"bytes,5,opt,name=riverName,proto3" json:"riverName,omitempty"`
	UserID       string `protobuf:"bytes,6,opt,name=userID,proto3" json:"userID,omitempty"`
	CreationDate int64  `protobuf:"varint,7,opt,name=creationDate,proto3" json:"creationDate,omitempty"`
	UpdateDate   int64  `protobuf:"varint,8,opt,name=updateDate,proto3" json:"updateDate,omitempty"`
	Currentstate string `protobuf:"bytes,9,opt,name=currentstate,proto3" json:"currentstate,omitempty"`
	Serial       string `protobuf:"bytes,10,opt,name=serial,proto3" json:"serial,omitempty"`
	Location     *Point `protobuf:"bytes,11,opt,name=location,proto3" json:"location,omitempty"`
}

func (x *Module) Reset() {
	*x = Module{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_river_management_handlers_grpc_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Module) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Module) ProtoMessage() {}

func (x *Module) ProtoReflect() protoreflect.Message {
	mi := &file_app_river_management_handlers_grpc_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Module.ProtoReflect.Descriptor instead.
func (*Module) Descriptor() ([]byte, []int) {
	return file_app_river_management_handlers_grpc_service_proto_rawDescGZIP(), []int{2}
}

func (x *Module) GetModuleID() string {
	if x != nil {
		return x.ModuleID
	}
	return ""
}

func (x *Module) GetPhoneNumber() string {
	if x != nil {
		return x.PhoneNumber
	}
	return ""
}

func (x *Module) GetAlias() string {
	if x != nil {
		return x.Alias
	}
	return ""
}

func (x *Module) GetRiverID() string {
	if x != nil {
		return x.RiverID
	}
	return ""
}

func (x *Module) GetRiverName() string {
	if x != nil {
		return x.RiverName
	}
	return ""
}

func (x *Module) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

func (x *Module) GetCreationDate() int64 {
	if x != nil {
		return x.CreationDate
	}
	return 0
}

func (x *Module) GetUpdateDate() int64 {
	if x != nil {
		return x.UpdateDate
	}
	return 0
}

func (x *Module) GetCurrentstate() string {
	if x != nil {
		return x.Currentstate
	}
	return ""
}

func (x *Module) GetSerial() string {
	if x != nil {
		return x.Serial
	}
	return ""
}

func (x *Module) GetLocation() *Point {
	if x != nil {
		return x.Location
	}
	return nil
}

type Point struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Latitude  float64 `protobuf:"fixed64,1,opt,name=latitude,proto3" json:"latitude,omitempty"`
	Longitude float64 `protobuf:"fixed64,64,opt,name=longitude,proto3" json:"longitude,omitempty"`
}

func (x *Point) Reset() {
	*x = Point{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_river_management_handlers_grpc_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Point) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Point) ProtoMessage() {}

func (x *Point) ProtoReflect() protoreflect.Message {
	mi := &file_app_river_management_handlers_grpc_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Point.ProtoReflect.Descriptor instead.
func (*Point) Descriptor() ([]byte, []int) {
	return file_app_river_management_handlers_grpc_service_proto_rawDescGZIP(), []int{3}
}

func (x *Point) GetLatitude() float64 {
	if x != nil {
		return x.Latitude
	}
	return 0
}

func (x *Point) GetLongitude() float64 {
	if x != nil {
		return x.Longitude
	}
	return 0
}

var File_app_river_management_handlers_grpc_service_proto protoreflect.FileDescriptor

var file_app_river_management_handlers_grpc_service_proto_rawDesc = []byte{
	0x0a, 0x30, 0x61, 0x70, 0x70, 0x2f, 0x72, 0x69, 0x76, 0x65, 0x72, 0x2d, 0x6d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x73, 0x2f,
	0x67, 0x72, 0x70, 0x63, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x49, 0x0a, 0x13, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a,
	0x08, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x22, 0x34, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x68, 0x6f, 0x6e,
	0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0xd6, 0x02, 0x0a, 0x06, 0x4d, 0x6f, 0x64, 0x75,
	0x6c, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x49, 0x44, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x49, 0x44, 0x12, 0x20,
	0x0a, 0x0b, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x12, 0x14, 0x0a, 0x05, 0x61, 0x6c, 0x69, 0x61, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x61, 0x6c, 0x69, 0x61, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x69, 0x76, 0x65, 0x72, 0x49,
	0x44, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x72, 0x69, 0x76, 0x65, 0x72, 0x49, 0x44,
	0x12, 0x1c, 0x0a, 0x09, 0x72, 0x69, 0x76, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x69, 0x76, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x22, 0x0a, 0x0c, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x44, 0x61, 0x74, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x44, 0x61, 0x74, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x44, 0x61, 0x74, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x63, 0x75,
	0x72, 0x72, 0x65, 0x6e, 0x74, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x12, 0x28, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x22, 0x41, 0x0a, 0x05, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x74,
	0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x6c, 0x61, 0x74,
	0x69, 0x74, 0x75, 0x64, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75,
	0x64, 0x65, 0x18, 0x40, 0x20, 0x01, 0x28, 0x01, 0x52, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74,
	0x75, 0x64, 0x65, 0x32, 0x97, 0x01, 0x0a, 0x07, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x40, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x42, 0x79, 0x50, 0x68,
	0x6f, 0x6e, 0x65, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x0d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4d, 0x6f, 0x64, 0x75, 0x6c,
	0x65, 0x12, 0x4a, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x6f, 0x64, 0x75, 0x6c,
	0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x42, 0x30, 0x5a,
	0x2e, 0x61, 0x70, 0x70, 0x2f, 0x72, 0x69, 0x76, 0x65, 0x72, 0x2d, 0x6d, 0x61, 0x6e, 0x61, 0x67,
	0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x73, 0x2f, 0x67,
	0x72, 0x70, 0x63, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_app_river_management_handlers_grpc_service_proto_rawDescOnce sync.Once
	file_app_river_management_handlers_grpc_service_proto_rawDescData = file_app_river_management_handlers_grpc_service_proto_rawDesc
)

func file_app_river_management_handlers_grpc_service_proto_rawDescGZIP() []byte {
	file_app_river_management_handlers_grpc_service_proto_rawDescOnce.Do(func() {
		file_app_river_management_handlers_grpc_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_app_river_management_handlers_grpc_service_proto_rawDescData)
	})
	return file_app_river_management_handlers_grpc_service_proto_rawDescData
}

var file_app_river_management_handlers_grpc_service_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_app_river_management_handlers_grpc_service_proto_goTypes = []interface{}{
	(*UpdateModuleRequest)(nil), // 0: proto.UpdateModuleRequest
	(*GetModuleRequest)(nil),    // 1: proto.GetModuleRequest
	(*Module)(nil),              // 2: proto.Module
	(*Point)(nil),               // 3: proto.Point
	(*emptypb.Empty)(nil),       // 4: google.protobuf.Empty
}
var file_app_river_management_handlers_grpc_service_proto_depIdxs = []int32{
	3, // 0: proto.Module.location:type_name -> proto.Point
	1, // 1: proto.Service.GetModuleByPhonenumber:input_type -> proto.GetModuleRequest
	0, // 2: proto.Service.UpdateModuleStatus:input_type -> proto.UpdateModuleRequest
	2, // 3: proto.Service.GetModuleByPhonenumber:output_type -> proto.Module
	4, // 4: proto.Service.UpdateModuleStatus:output_type -> google.protobuf.Empty
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_app_river_management_handlers_grpc_service_proto_init() }
func file_app_river_management_handlers_grpc_service_proto_init() {
	if File_app_river_management_handlers_grpc_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_app_river_management_handlers_grpc_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateModuleRequest); i {
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
		file_app_river_management_handlers_grpc_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetModuleRequest); i {
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
		file_app_river_management_handlers_grpc_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Module); i {
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
		file_app_river_management_handlers_grpc_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Point); i {
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
			RawDescriptor: file_app_river_management_handlers_grpc_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_app_river_management_handlers_grpc_service_proto_goTypes,
		DependencyIndexes: file_app_river_management_handlers_grpc_service_proto_depIdxs,
		MessageInfos:      file_app_river_management_handlers_grpc_service_proto_msgTypes,
	}.Build()
	File_app_river_management_handlers_grpc_service_proto = out.File
	file_app_river_management_handlers_grpc_service_proto_rawDesc = nil
	file_app_river_management_handlers_grpc_service_proto_goTypes = nil
	file_app_river_management_handlers_grpc_service_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ServiceClient is the client API for Service service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ServiceClient interface {
	GetModuleByPhonenumber(ctx context.Context, in *GetModuleRequest, opts ...grpc.CallOption) (*Module, error)
	UpdateModuleStatus(ctx context.Context, in *UpdateModuleRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type serviceClient struct {
	cc grpc.ClientConnInterface
}

func NewServiceClient(cc grpc.ClientConnInterface) ServiceClient {
	return &serviceClient{cc}
}

func (c *serviceClient) GetModuleByPhonenumber(ctx context.Context, in *GetModuleRequest, opts ...grpc.CallOption) (*Module, error) {
	out := new(Module)
	err := c.cc.Invoke(ctx, "/proto.Service/GetModuleByPhonenumber", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) UpdateModuleStatus(ctx context.Context, in *UpdateModuleRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/proto.Service/UpdateModuleStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServiceServer is the server API for Service service.
type ServiceServer interface {
	GetModuleByPhonenumber(context.Context, *GetModuleRequest) (*Module, error)
	UpdateModuleStatus(context.Context, *UpdateModuleRequest) (*emptypb.Empty, error)
}

// UnimplementedServiceServer can be embedded to have forward compatible implementations.
type UnimplementedServiceServer struct {
}

func (*UnimplementedServiceServer) GetModuleByPhonenumber(context.Context, *GetModuleRequest) (*Module, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetModuleByPhonenumber not implemented")
}
func (*UnimplementedServiceServer) UpdateModuleStatus(context.Context, *UpdateModuleRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateModuleStatus not implemented")
}

func RegisterServiceServer(s *grpc.Server, srv ServiceServer) {
	s.RegisterService(&_Service_serviceDesc, srv)
}

func _Service_GetModuleByPhonenumber_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetModuleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).GetModuleByPhonenumber(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Service/GetModuleByPhonenumber",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).GetModuleByPhonenumber(ctx, req.(*GetModuleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_UpdateModuleStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateModuleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).UpdateModuleStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Service/UpdateModuleStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).UpdateModuleStatus(ctx, req.(*UpdateModuleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Service_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Service",
	HandlerType: (*ServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetModuleByPhonenumber",
			Handler:    _Service_GetModuleByPhonenumber_Handler,
		},
		{
			MethodName: "UpdateModuleStatus",
			Handler:    _Service_UpdateModuleStatus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "app/river-management/handlers/grpc/service.proto",
}
