// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.6.1
// source: admin_service.proto

package pb

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

type CreateReportRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductId   int64  `protobuf:"varint,1,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *CreateReportRequest) Reset() {
	*x = CreateReportRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admin_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateReportRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateReportRequest) ProtoMessage() {}

func (x *CreateReportRequest) ProtoReflect() protoreflect.Message {
	mi := &file_admin_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateReportRequest.ProtoReflect.Descriptor instead.
func (*CreateReportRequest) Descriptor() ([]byte, []int) {
	return file_admin_service_proto_rawDescGZIP(), []int{0}
}

func (x *CreateReportRequest) GetProductId() int64 {
	if x != nil {
		return x.ProductId
	}
	return 0
}

func (x *CreateReportRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type CreateReportResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *CreateReportResponse) Reset() {
	*x = CreateReportResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admin_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateReportResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateReportResponse) ProtoMessage() {}

func (x *CreateReportResponse) ProtoReflect() protoreflect.Message {
	mi := &file_admin_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateReportResponse.ProtoReflect.Descriptor instead.
func (*CreateReportResponse) Descriptor() ([]byte, []int) {
	return file_admin_service_proto_rawDescGZIP(), []int{1}
}

func (x *CreateReportResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type HandleReportRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ReportId int64 `protobuf:"varint,1,opt,name=report_id,json=reportId,proto3" json:"report_id,omitempty"`
}

func (x *HandleReportRequest) Reset() {
	*x = HandleReportRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admin_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HandleReportRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HandleReportRequest) ProtoMessage() {}

func (x *HandleReportRequest) ProtoReflect() protoreflect.Message {
	mi := &file_admin_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HandleReportRequest.ProtoReflect.Descriptor instead.
func (*HandleReportRequest) Descriptor() ([]byte, []int) {
	return file_admin_service_proto_rawDescGZIP(), []int{2}
}

func (x *HandleReportRequest) GetReportId() int64 {
	if x != nil {
		return x.ReportId
	}
	return 0
}

type HandleReportResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *HandleReportResponse) Reset() {
	*x = HandleReportResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admin_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HandleReportResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HandleReportResponse) ProtoMessage() {}

func (x *HandleReportResponse) ProtoReflect() protoreflect.Message {
	mi := &file_admin_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HandleReportResponse.ProtoReflect.Descriptor instead.
func (*HandleReportResponse) Descriptor() ([]byte, []int) {
	return file_admin_service_proto_rawDescGZIP(), []int{3}
}

func (x *HandleReportResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type GetAllReportResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ListReport []*GetAllReportResponse_ProductReport `protobuf:"bytes,1,rep,name=listReport,proto3" json:"listReport,omitempty"`
}

func (x *GetAllReportResponse) Reset() {
	*x = GetAllReportResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admin_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllReportResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllReportResponse) ProtoMessage() {}

func (x *GetAllReportResponse) ProtoReflect() protoreflect.Message {
	mi := &file_admin_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllReportResponse.ProtoReflect.Descriptor instead.
func (*GetAllReportResponse) Descriptor() ([]byte, []int) {
	return file_admin_service_proto_rawDescGZIP(), []int{4}
}

func (x *GetAllReportResponse) GetListReport() []*GetAllReportResponse_ProductReport {
	if x != nil {
		return x.ListReport
	}
	return nil
}

type GetAllReportResponse_ProductReport struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductImage string `protobuf:"bytes,1,opt,name=productImage,proto3" json:"productImage,omitempty"`
	ProductName  string `protobuf:"bytes,2,opt,name=productName,proto3" json:"productName,omitempty"`
	Description  string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	ReportId     int64  `protobuf:"varint,4,opt,name=reportId,proto3" json:"reportId,omitempty"`
}

func (x *GetAllReportResponse_ProductReport) Reset() {
	*x = GetAllReportResponse_ProductReport{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admin_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllReportResponse_ProductReport) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllReportResponse_ProductReport) ProtoMessage() {}

func (x *GetAllReportResponse_ProductReport) ProtoReflect() protoreflect.Message {
	mi := &file_admin_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllReportResponse_ProductReport.ProtoReflect.Descriptor instead.
func (*GetAllReportResponse_ProductReport) Descriptor() ([]byte, []int) {
	return file_admin_service_proto_rawDescGZIP(), []int{4, 0}
}

func (x *GetAllReportResponse_ProductReport) GetProductImage() string {
	if x != nil {
		return x.ProductImage
	}
	return ""
}

func (x *GetAllReportResponse_ProductReport) GetProductName() string {
	if x != nil {
		return x.ProductName
	}
	return ""
}

func (x *GetAllReportResponse_ProductReport) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *GetAllReportResponse_ProductReport) GetReportId() int64 {
	if x != nil {
		return x.ReportId
	}
	return 0
}

var File_admin_service_proto protoreflect.FileDescriptor

var file_admin_service_proto_rawDesc = []byte{
	0x0a, 0x13, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x63, 0x65,
	0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0d, 0x67,
	0x65, 0x6e, 0x65, 0x72, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x56, 0x0a, 0x13,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x49, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x22, 0x30, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x32, 0x0a, 0x13, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65,
	0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a,
	0x09, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x08, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x49, 0x64, 0x22, 0x30, 0x0a, 0x14, 0x48, 0x61,
	0x6e, 0x64, 0x6c, 0x65, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0xfb, 0x01, 0x0a,
	0x14, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4d, 0x0a, 0x0a, 0x6c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x70,
	0x6f, 0x72, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x65, 0x63, 0x6f, 0x6d,
	0x6d, 0x65, 0x72, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x70, 0x6f,
	0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x0a, 0x6c, 0x69, 0x73, 0x74, 0x52, 0x65,
	0x70, 0x6f, 0x72, 0x74, 0x1a, 0x93, 0x01, 0x0a, 0x0d, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x22, 0x0a, 0x0c, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x70, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x70, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a,
	0x0a, 0x08, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x49, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x08, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x49, 0x64, 0x32, 0x85, 0x03, 0x0a, 0x0c, 0x41,
	0x64, 0x6d, 0x69, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x31, 0x0a, 0x04, 0x50,
	0x69, 0x6e, 0x67, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x0f, 0x2e, 0x65, 0x63,
	0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x63, 0x65, 0x2e, 0x50, 0x6f, 0x6e, 0x67, 0x22, 0x00, 0x12, 0x51,
	0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x1e,
	0x2e, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x63, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f,
	0x2e, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x63, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x51, 0x0a, 0x0c, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x52, 0x65, 0x70, 0x6f, 0x72,
	0x74, 0x12, 0x1e, 0x2e, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x63, 0x65, 0x2e, 0x48, 0x61,
	0x6e, 0x64, 0x6c, 0x65, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1f, 0x2e, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x63, 0x65, 0x2e, 0x48, 0x61,
	0x6e, 0x64, 0x6c, 0x65, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x51, 0x0a, 0x0c, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65,
	0x70, 0x6f, 0x72, 0x74, 0x12, 0x1e, 0x2e, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x63, 0x65,
	0x2e, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x63, 0x65,
	0x2e, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x49, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x41, 0x6c,
	0x6c, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a,
	0x1f, 0x2e, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x41,
	0x6c, 0x6c, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_admin_service_proto_rawDescOnce sync.Once
	file_admin_service_proto_rawDescData = file_admin_service_proto_rawDesc
)

func file_admin_service_proto_rawDescGZIP() []byte {
	file_admin_service_proto_rawDescOnce.Do(func() {
		file_admin_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_admin_service_proto_rawDescData)
	})
	return file_admin_service_proto_rawDescData
}

var file_admin_service_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_admin_service_proto_goTypes = []interface{}{
	(*CreateReportRequest)(nil),                // 0: ecommerce.CreateReportRequest
	(*CreateReportResponse)(nil),               // 1: ecommerce.CreateReportResponse
	(*HandleReportRequest)(nil),                // 2: ecommerce.HandleReportRequest
	(*HandleReportResponse)(nil),               // 3: ecommerce.HandleReportResponse
	(*GetAllReportResponse)(nil),               // 4: ecommerce.GetAllReportResponse
	(*GetAllReportResponse_ProductReport)(nil), // 5: ecommerce.GetAllReportResponse.ProductReport
	(*empty.Empty)(nil),                        // 6: google.protobuf.Empty
	(*Pong)(nil),                               // 7: ecommerce.Pong
}
var file_admin_service_proto_depIdxs = []int32{
	5, // 0: ecommerce.GetAllReportResponse.listReport:type_name -> ecommerce.GetAllReportResponse.ProductReport
	6, // 1: ecommerce.AdminService.Ping:input_type -> google.protobuf.Empty
	0, // 2: ecommerce.AdminService.CreateReport:input_type -> ecommerce.CreateReportRequest
	2, // 3: ecommerce.AdminService.HandleReport:input_type -> ecommerce.HandleReportRequest
	2, // 4: ecommerce.AdminService.DeleteReport:input_type -> ecommerce.HandleReportRequest
	6, // 5: ecommerce.AdminService.GetAllReport:input_type -> google.protobuf.Empty
	7, // 6: ecommerce.AdminService.Ping:output_type -> ecommerce.Pong
	1, // 7: ecommerce.AdminService.CreateReport:output_type -> ecommerce.CreateReportResponse
	3, // 8: ecommerce.AdminService.HandleReport:output_type -> ecommerce.HandleReportResponse
	3, // 9: ecommerce.AdminService.DeleteReport:output_type -> ecommerce.HandleReportResponse
	4, // 10: ecommerce.AdminService.GetAllReport:output_type -> ecommerce.GetAllReportResponse
	6, // [6:11] is the sub-list for method output_type
	1, // [1:6] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_admin_service_proto_init() }
func file_admin_service_proto_init() {
	if File_admin_service_proto != nil {
		return
	}
	file_general_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_admin_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateReportRequest); i {
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
		file_admin_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateReportResponse); i {
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
		file_admin_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HandleReportRequest); i {
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
		file_admin_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HandleReportResponse); i {
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
		file_admin_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllReportResponse); i {
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
		file_admin_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllReportResponse_ProductReport); i {
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
			RawDescriptor: file_admin_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_admin_service_proto_goTypes,
		DependencyIndexes: file_admin_service_proto_depIdxs,
		MessageInfos:      file_admin_service_proto_msgTypes,
	}.Build()
	File_admin_service_proto = out.File
	file_admin_service_proto_rawDesc = nil
	file_admin_service_proto_goTypes = nil
	file_admin_service_proto_depIdxs = nil
}
