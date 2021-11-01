// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.1
// source: db_transport.proto

package dbpb

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

type TransportRecordDTO struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID            string `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	ClusterId     string `protobuf:"bytes,2,opt,name=ClusterId,proto3" json:"ClusterId,omitempty"`
	TransportType string `protobuf:"bytes,3,opt,name=TransportType,proto3" json:"TransportType,omitempty"`
	FilePath      string `protobuf:"bytes,4,opt,name=FilePath,proto3" json:"FilePath,omitempty"`
	TenantId      string `protobuf:"bytes,5,opt,name=TenantId,proto3" json:"TenantId,omitempty"`
	Status        string `protobuf:"bytes,6,opt,name=Status,proto3" json:"Status,omitempty"`
	StartTime     int64  `protobuf:"varint,7,opt,name=StartTime,proto3" json:"StartTime,omitempty"`
	EndTime       int64  `protobuf:"varint,8,opt,name=EndTime,proto3" json:"EndTime,omitempty"`
}

func (x *TransportRecordDTO) Reset() {
	*x = TransportRecordDTO{}
	if protoimpl.UnsafeEnabled {
		mi := &file_db_transport_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransportRecordDTO) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransportRecordDTO) ProtoMessage() {}

func (x *TransportRecordDTO) ProtoReflect() protoreflect.Message {
	mi := &file_db_transport_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransportRecordDTO.ProtoReflect.Descriptor instead.
func (*TransportRecordDTO) Descriptor() ([]byte, []int) {
	return file_db_transport_proto_rawDescGZIP(), []int{0}
}

func (x *TransportRecordDTO) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *TransportRecordDTO) GetClusterId() string {
	if x != nil {
		return x.ClusterId
	}
	return ""
}

func (x *TransportRecordDTO) GetTransportType() string {
	if x != nil {
		return x.TransportType
	}
	return ""
}

func (x *TransportRecordDTO) GetFilePath() string {
	if x != nil {
		return x.FilePath
	}
	return ""
}

func (x *TransportRecordDTO) GetTenantId() string {
	if x != nil {
		return x.TenantId
	}
	return ""
}

func (x *TransportRecordDTO) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *TransportRecordDTO) GetStartTime() int64 {
	if x != nil {
		return x.StartTime
	}
	return 0
}

func (x *TransportRecordDTO) GetEndTime() int64 {
	if x != nil {
		return x.EndTime
	}
	return 0
}

type DBCreateTransportRecordRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Record *TransportRecordDTO `protobuf:"bytes,1,opt,name=record,proto3" json:"record,omitempty"`
}

func (x *DBCreateTransportRecordRequest) Reset() {
	*x = DBCreateTransportRecordRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_db_transport_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DBCreateTransportRecordRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DBCreateTransportRecordRequest) ProtoMessage() {}

func (x *DBCreateTransportRecordRequest) ProtoReflect() protoreflect.Message {
	mi := &file_db_transport_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DBCreateTransportRecordRequest.ProtoReflect.Descriptor instead.
func (*DBCreateTransportRecordRequest) Descriptor() ([]byte, []int) {
	return file_db_transport_proto_rawDescGZIP(), []int{1}
}

func (x *DBCreateTransportRecordRequest) GetRecord() *TransportRecordDTO {
	if x != nil {
		return x.Record
	}
	return nil
}

type DBCreateTransportRecordResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status *DBClusterResponseStatus `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Id     string                   `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DBCreateTransportRecordResponse) Reset() {
	*x = DBCreateTransportRecordResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_db_transport_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DBCreateTransportRecordResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DBCreateTransportRecordResponse) ProtoMessage() {}

func (x *DBCreateTransportRecordResponse) ProtoReflect() protoreflect.Message {
	mi := &file_db_transport_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DBCreateTransportRecordResponse.ProtoReflect.Descriptor instead.
func (*DBCreateTransportRecordResponse) Descriptor() ([]byte, []int) {
	return file_db_transport_proto_rawDescGZIP(), []int{2}
}

func (x *DBCreateTransportRecordResponse) GetStatus() *DBClusterResponseStatus {
	if x != nil {
		return x.Status
	}
	return nil
}

func (x *DBCreateTransportRecordResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DBUpdateTransportRecordRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Record *TransportRecordDTO `protobuf:"bytes,1,opt,name=record,proto3" json:"record,omitempty"`
}

func (x *DBUpdateTransportRecordRequest) Reset() {
	*x = DBUpdateTransportRecordRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_db_transport_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DBUpdateTransportRecordRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DBUpdateTransportRecordRequest) ProtoMessage() {}

func (x *DBUpdateTransportRecordRequest) ProtoReflect() protoreflect.Message {
	mi := &file_db_transport_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DBUpdateTransportRecordRequest.ProtoReflect.Descriptor instead.
func (*DBUpdateTransportRecordRequest) Descriptor() ([]byte, []int) {
	return file_db_transport_proto_rawDescGZIP(), []int{3}
}

func (x *DBUpdateTransportRecordRequest) GetRecord() *TransportRecordDTO {
	if x != nil {
		return x.Record
	}
	return nil
}

type DBUpdateTransportRecordResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status *DBClusterResponseStatus `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *DBUpdateTransportRecordResponse) Reset() {
	*x = DBUpdateTransportRecordResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_db_transport_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DBUpdateTransportRecordResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DBUpdateTransportRecordResponse) ProtoMessage() {}

func (x *DBUpdateTransportRecordResponse) ProtoReflect() protoreflect.Message {
	mi := &file_db_transport_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DBUpdateTransportRecordResponse.ProtoReflect.Descriptor instead.
func (*DBUpdateTransportRecordResponse) Descriptor() ([]byte, []int) {
	return file_db_transport_proto_rawDescGZIP(), []int{4}
}

func (x *DBUpdateTransportRecordResponse) GetStatus() *DBClusterResponseStatus {
	if x != nil {
		return x.Status
	}
	return nil
}

type DBFindTransportRecordByIDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RecordId string `protobuf:"bytes,1,opt,name=recordId,proto3" json:"recordId,omitempty"`
}

func (x *DBFindTransportRecordByIDRequest) Reset() {
	*x = DBFindTransportRecordByIDRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_db_transport_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DBFindTransportRecordByIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DBFindTransportRecordByIDRequest) ProtoMessage() {}

func (x *DBFindTransportRecordByIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_db_transport_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DBFindTransportRecordByIDRequest.ProtoReflect.Descriptor instead.
func (*DBFindTransportRecordByIDRequest) Descriptor() ([]byte, []int) {
	return file_db_transport_proto_rawDescGZIP(), []int{5}
}

func (x *DBFindTransportRecordByIDRequest) GetRecordId() string {
	if x != nil {
		return x.RecordId
	}
	return ""
}

type DBFindTransportRecordByIDResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status *DBClusterResponseStatus `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Record *TransportRecordDTO      `protobuf:"bytes,2,opt,name=record,proto3" json:"record,omitempty"`
}

func (x *DBFindTransportRecordByIDResponse) Reset() {
	*x = DBFindTransportRecordByIDResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_db_transport_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DBFindTransportRecordByIDResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DBFindTransportRecordByIDResponse) ProtoMessage() {}

func (x *DBFindTransportRecordByIDResponse) ProtoReflect() protoreflect.Message {
	mi := &file_db_transport_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DBFindTransportRecordByIDResponse.ProtoReflect.Descriptor instead.
func (*DBFindTransportRecordByIDResponse) Descriptor() ([]byte, []int) {
	return file_db_transport_proto_rawDescGZIP(), []int{6}
}

func (x *DBFindTransportRecordByIDResponse) GetStatus() *DBClusterResponseStatus {
	if x != nil {
		return x.Status
	}
	return nil
}

func (x *DBFindTransportRecordByIDResponse) GetRecord() *TransportRecordDTO {
	if x != nil {
		return x.Record
	}
	return nil
}

type DBListTransportRecordRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page      *DBPageDTO `protobuf:"bytes,1,opt,name=page,proto3" json:"page,omitempty"`
	ClusterId string     `protobuf:"bytes,2,opt,name=clusterId,proto3" json:"clusterId,omitempty"`
	RecordId  string     `protobuf:"bytes,3,opt,name=recordId,proto3" json:"recordId,omitempty"`
}

func (x *DBListTransportRecordRequest) Reset() {
	*x = DBListTransportRecordRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_db_transport_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DBListTransportRecordRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DBListTransportRecordRequest) ProtoMessage() {}

func (x *DBListTransportRecordRequest) ProtoReflect() protoreflect.Message {
	mi := &file_db_transport_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DBListTransportRecordRequest.ProtoReflect.Descriptor instead.
func (*DBListTransportRecordRequest) Descriptor() ([]byte, []int) {
	return file_db_transport_proto_rawDescGZIP(), []int{7}
}

func (x *DBListTransportRecordRequest) GetPage() *DBPageDTO {
	if x != nil {
		return x.Page
	}
	return nil
}

func (x *DBListTransportRecordRequest) GetClusterId() string {
	if x != nil {
		return x.ClusterId
	}
	return ""
}

func (x *DBListTransportRecordRequest) GetRecordId() string {
	if x != nil {
		return x.RecordId
	}
	return ""
}

type DBListTransportRecordResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status  *DBClusterResponseStatus `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Page    *DBPageDTO               `protobuf:"bytes,2,opt,name=page,proto3" json:"page,omitempty"`
	Records []*TransportRecordDTO    `protobuf:"bytes,3,rep,name=records,proto3" json:"records,omitempty"`
}

func (x *DBListTransportRecordResponse) Reset() {
	*x = DBListTransportRecordResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_db_transport_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DBListTransportRecordResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DBListTransportRecordResponse) ProtoMessage() {}

func (x *DBListTransportRecordResponse) ProtoReflect() protoreflect.Message {
	mi := &file_db_transport_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DBListTransportRecordResponse.ProtoReflect.Descriptor instead.
func (*DBListTransportRecordResponse) Descriptor() ([]byte, []int) {
	return file_db_transport_proto_rawDescGZIP(), []int{8}
}

func (x *DBListTransportRecordResponse) GetStatus() *DBClusterResponseStatus {
	if x != nil {
		return x.Status
	}
	return nil
}

func (x *DBListTransportRecordResponse) GetPage() *DBPageDTO {
	if x != nil {
		return x.Page
	}
	return nil
}

func (x *DBListTransportRecordResponse) GetRecords() []*TransportRecordDTO {
	if x != nil {
		return x.Records
	}
	return nil
}

var File_db_transport_proto protoreflect.FileDescriptor

var file_db_transport_proto_rawDesc = []byte{
	0x0a, 0x12, 0x64, 0x62, 0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x10, 0x64, 0x62, 0x5f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf0, 0x01, 0x0a, 0x12, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x44, 0x54, 0x4f, 0x12, 0x0e, 0x0a,
	0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x12, 0x1c, 0x0a,
	0x09, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x49, 0x64, 0x12, 0x24, 0x0a, 0x0d, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x54, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0d, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x46, 0x69, 0x6c, 0x65, 0x50, 0x61, 0x74, 0x68, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x46, 0x69, 0x6c, 0x65, 0x50, 0x61, 0x74, 0x68, 0x12, 0x1a, 0x0a,
	0x08, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x1c, 0x0a, 0x09, 0x53, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x53, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x45, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x07, 0x45, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x4d, 0x0a, 0x1e, 0x44, 0x42, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65,
	0x63, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2b, 0x0a, 0x06, 0x72,
	0x65, 0x63, 0x6f, 0x72, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x44, 0x54, 0x4f,
	0x52, 0x06, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x22, 0x63, 0x0a, 0x1f, 0x44, 0x42, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x63,
	0x6f, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x30, 0x0a, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x44, 0x42,
	0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x4d, 0x0a,
	0x1e, 0x44, 0x42, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f,
	0x72, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x2b, 0x0a, 0x06, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x13, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x72,
	0x64, 0x44, 0x54, 0x4f, 0x52, 0x06, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x22, 0x53, 0x0a, 0x1f,
	0x44, 0x42, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72,
	0x74, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x30, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x18, 0x2e, 0x44, 0x42, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x22, 0x3e, 0x0a, 0x20, 0x44, 0x42, 0x46, 0x69, 0x6e, 0x64, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x49,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x49,
	0x64, 0x22, 0x82, 0x01, 0x0a, 0x21, 0x44, 0x42, 0x46, 0x69, 0x6e, 0x64, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x42, 0x79, 0x49, 0x44, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x30, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x44, 0x42, 0x43, 0x6c, 0x75, 0x73,
	0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x2b, 0x0a, 0x06, 0x72, 0x65, 0x63,
	0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x44, 0x54, 0x4f, 0x52, 0x06,
	0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x22, 0x78, 0x0a, 0x1c, 0x44, 0x42, 0x4c, 0x69, 0x73, 0x74,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x44, 0x42, 0x50, 0x61, 0x67, 0x65, 0x44, 0x54, 0x4f,
	0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65,
	0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6c, 0x75, 0x73, 0x74,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x49, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x49, 0x64,
	0x22, 0xa0, 0x01, 0x0a, 0x1d, 0x44, 0x42, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x30, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x18, 0x2e, 0x44, 0x42, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x1e, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x44, 0x42, 0x50, 0x61, 0x67, 0x65, 0x44, 0x54, 0x4f, 0x52, 0x04,
	0x70, 0x61, 0x67, 0x65, 0x12, 0x2d, 0x0a, 0x07, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x18,
	0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72,
	0x74, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x44, 0x54, 0x4f, 0x52, 0x07, 0x72, 0x65, 0x63, 0x6f,
	0x72, 0x64, 0x73, 0x42, 0x11, 0x5a, 0x0f, 0x2e, 0x2f, 0x2e, 0x2e, 0x2f, 0x64, 0x62, 0x70, 0x62,
	0x2f, 0x3b, 0x64, 0x62, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_db_transport_proto_rawDescOnce sync.Once
	file_db_transport_proto_rawDescData = file_db_transport_proto_rawDesc
)

func file_db_transport_proto_rawDescGZIP() []byte {
	file_db_transport_proto_rawDescOnce.Do(func() {
		file_db_transport_proto_rawDescData = protoimpl.X.CompressGZIP(file_db_transport_proto_rawDescData)
	})
	return file_db_transport_proto_rawDescData
}

var file_db_transport_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_db_transport_proto_goTypes = []interface{}{
	(*TransportRecordDTO)(nil),                // 0: TransportRecordDTO
	(*DBCreateTransportRecordRequest)(nil),    // 1: DBCreateTransportRecordRequest
	(*DBCreateTransportRecordResponse)(nil),   // 2: DBCreateTransportRecordResponse
	(*DBUpdateTransportRecordRequest)(nil),    // 3: DBUpdateTransportRecordRequest
	(*DBUpdateTransportRecordResponse)(nil),   // 4: DBUpdateTransportRecordResponse
	(*DBFindTransportRecordByIDRequest)(nil),  // 5: DBFindTransportRecordByIDRequest
	(*DBFindTransportRecordByIDResponse)(nil), // 6: DBFindTransportRecordByIDResponse
	(*DBListTransportRecordRequest)(nil),      // 7: DBListTransportRecordRequest
	(*DBListTransportRecordResponse)(nil),     // 8: DBListTransportRecordResponse
	(*DBClusterResponseStatus)(nil),           // 9: DBClusterResponseStatus
	(*DBPageDTO)(nil),                         // 10: DBPageDTO
}
var file_db_transport_proto_depIdxs = []int32{
	0,  // 0: DBCreateTransportRecordRequest.record:type_name -> TransportRecordDTO
	9,  // 1: DBCreateTransportRecordResponse.status:type_name -> DBClusterResponseStatus
	0,  // 2: DBUpdateTransportRecordRequest.record:type_name -> TransportRecordDTO
	9,  // 3: DBUpdateTransportRecordResponse.status:type_name -> DBClusterResponseStatus
	9,  // 4: DBFindTransportRecordByIDResponse.status:type_name -> DBClusterResponseStatus
	0,  // 5: DBFindTransportRecordByIDResponse.record:type_name -> TransportRecordDTO
	10, // 6: DBListTransportRecordRequest.page:type_name -> DBPageDTO
	9,  // 7: DBListTransportRecordResponse.status:type_name -> DBClusterResponseStatus
	10, // 8: DBListTransportRecordResponse.page:type_name -> DBPageDTO
	0,  // 9: DBListTransportRecordResponse.records:type_name -> TransportRecordDTO
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_db_transport_proto_init() }
func file_db_transport_proto_init() {
	if File_db_transport_proto != nil {
		return
	}
	file_db_cluster_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_db_transport_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransportRecordDTO); i {
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
		file_db_transport_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DBCreateTransportRecordRequest); i {
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
		file_db_transport_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DBCreateTransportRecordResponse); i {
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
		file_db_transport_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DBUpdateTransportRecordRequest); i {
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
		file_db_transport_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DBUpdateTransportRecordResponse); i {
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
		file_db_transport_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DBFindTransportRecordByIDRequest); i {
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
		file_db_transport_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DBFindTransportRecordByIDResponse); i {
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
		file_db_transport_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DBListTransportRecordRequest); i {
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
		file_db_transport_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DBListTransportRecordResponse); i {
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
			RawDescriptor: file_db_transport_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_db_transport_proto_goTypes,
		DependencyIndexes: file_db_transport_proto_depIdxs,
		MessageInfos:      file_db_transport_proto_msgTypes,
	}.Build()
	File_db_transport_proto = out.File
	file_db_transport_proto_rawDesc = nil
	file_db_transport_proto_goTypes = nil
	file_db_transport_proto_depIdxs = nil
}
