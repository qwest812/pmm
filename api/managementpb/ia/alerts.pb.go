// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: managementpb/ia/alerts.proto

package iav1beta1

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	_ "github.com/mwitkow/go-proto-validators"
	managementpb "github.com/percona/pmm/api/managementpb"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// Alert represents Alert.
type Alert struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID.
	AlertId string `protobuf:"bytes,1,opt,name=alert_id,json=alertId,proto3" json:"alert_id,omitempty"`
	// Human-readable summary.
	Summary string `protobuf:"bytes,2,opt,name=summary,proto3" json:"summary,omitempty"`
	// Severity.
	Severity managementpb.Severity `protobuf:"varint,3,opt,name=severity,proto3,enum=management.Severity" json:"severity,omitempty"`
	// Status.
	Status Status `protobuf:"varint,4,opt,name=status,proto3,enum=ia.v1beta1.Status" json:"status,omitempty"`
	// Combined labels.
	Labels map[string]string `protobuf:"bytes,5,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// The current Alert Rule.
	Rule *Rule `protobuf:"bytes,6,opt,name=rule,proto3" json:"rule,omitempty"`
	// Alert creation time.
	CreatedAt *timestamp.Timestamp `protobuf:"bytes,7,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	// Alert last update time.
	UpdatedAt *timestamp.Timestamp `protobuf:"bytes,8,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *Alert) Reset() {
	*x = Alert{}
	if protoimpl.UnsafeEnabled {
		mi := &file_managementpb_ia_alerts_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Alert) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Alert) ProtoMessage() {}

func (x *Alert) ProtoReflect() protoreflect.Message {
	mi := &file_managementpb_ia_alerts_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Alert.ProtoReflect.Descriptor instead.
func (*Alert) Descriptor() ([]byte, []int) {
	return file_managementpb_ia_alerts_proto_rawDescGZIP(), []int{0}
}

func (x *Alert) GetAlertId() string {
	if x != nil {
		return x.AlertId
	}
	return ""
}

func (x *Alert) GetSummary() string {
	if x != nil {
		return x.Summary
	}
	return ""
}

func (x *Alert) GetSeverity() managementpb.Severity {
	if x != nil {
		return x.Severity
	}
	return managementpb.Severity_SEVERITY_INVALID
}

func (x *Alert) GetStatus() Status {
	if x != nil {
		return x.Status
	}
	return Status_STATUS_INVALID
}

func (x *Alert) GetLabels() map[string]string {
	if x != nil {
		return x.Labels
	}
	return nil
}

func (x *Alert) GetRule() *Rule {
	if x != nil {
		return x.Rule
	}
	return nil
}

func (x *Alert) GetCreatedAt() *timestamp.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Alert) GetUpdatedAt() *timestamp.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

type ListAlertsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Page request.
	PageParams *PageParams `protobuf:"bytes,1,opt,name=page_params,json=pageParams,proto3" json:"page_params,omitempty"`
}

func (x *ListAlertsRequest) Reset() {
	*x = ListAlertsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_managementpb_ia_alerts_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListAlertsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAlertsRequest) ProtoMessage() {}

func (x *ListAlertsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_managementpb_ia_alerts_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListAlertsRequest.ProtoReflect.Descriptor instead.
func (*ListAlertsRequest) Descriptor() ([]byte, []int) {
	return file_managementpb_ia_alerts_proto_rawDescGZIP(), []int{1}
}

func (x *ListAlertsRequest) GetPageParams() *PageParams {
	if x != nil {
		return x.PageParams
	}
	return nil
}

type ListAlertsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Alerts []*Alert `protobuf:"bytes,1,rep,name=alerts,proto3" json:"alerts,omitempty"`
	// Total items and pages.
	Totals *PageTotals `protobuf:"bytes,2,opt,name=totals,proto3" json:"totals,omitempty"`
}

func (x *ListAlertsResponse) Reset() {
	*x = ListAlertsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_managementpb_ia_alerts_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListAlertsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAlertsResponse) ProtoMessage() {}

func (x *ListAlertsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_managementpb_ia_alerts_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListAlertsResponse.ProtoReflect.Descriptor instead.
func (*ListAlertsResponse) Descriptor() ([]byte, []int) {
	return file_managementpb_ia_alerts_proto_rawDescGZIP(), []int{2}
}

func (x *ListAlertsResponse) GetAlerts() []*Alert {
	if x != nil {
		return x.Alerts
	}
	return nil
}

func (x *ListAlertsResponse) GetTotals() *PageTotals {
	if x != nil {
		return x.Totals
	}
	return nil
}

type ToggleAlertRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AlertId string `protobuf:"bytes,1,opt,name=alert_id,json=alertId,proto3" json:"alert_id,omitempty"`
	// Silences or unsilences alert if set.
	Silenced BooleanFlag `protobuf:"varint,2,opt,name=silenced,proto3,enum=ia.v1beta1.BooleanFlag" json:"silenced,omitempty"`
}

func (x *ToggleAlertRequest) Reset() {
	*x = ToggleAlertRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_managementpb_ia_alerts_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ToggleAlertRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ToggleAlertRequest) ProtoMessage() {}

func (x *ToggleAlertRequest) ProtoReflect() protoreflect.Message {
	mi := &file_managementpb_ia_alerts_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ToggleAlertRequest.ProtoReflect.Descriptor instead.
func (*ToggleAlertRequest) Descriptor() ([]byte, []int) {
	return file_managementpb_ia_alerts_proto_rawDescGZIP(), []int{3}
}

func (x *ToggleAlertRequest) GetAlertId() string {
	if x != nil {
		return x.AlertId
	}
	return ""
}

func (x *ToggleAlertRequest) GetSilenced() BooleanFlag {
	if x != nil {
		return x.Silenced
	}
	return BooleanFlag_DO_NOT_CHANGE
}

type ToggleAlertResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ToggleAlertResponse) Reset() {
	*x = ToggleAlertResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_managementpb_ia_alerts_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ToggleAlertResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ToggleAlertResponse) ProtoMessage() {}

func (x *ToggleAlertResponse) ProtoReflect() protoreflect.Message {
	mi := &file_managementpb_ia_alerts_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ToggleAlertResponse.ProtoReflect.Descriptor instead.
func (*ToggleAlertResponse) Descriptor() ([]byte, []int) {
	return file_managementpb_ia_alerts_proto_rawDescGZIP(), []int{4}
}

var File_managementpb_ia_alerts_proto protoreflect.FileDescriptor

var file_managementpb_ia_alerts_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x70, 0x62, 0x2f, 0x69,
	0x61, 0x2f, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a,
	0x69, 0x61, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x1a, 0x36, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x77, 0x69, 0x74, 0x6b, 0x6f, 0x77, 0x2f, 0x67,
	0x6f, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f,
	0x72, 0x73, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x22, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x70, 0x62, 0x2f,
	0x69, 0x61, 0x2f, 0x62, 0x6f, 0x6f, 0x6c, 0x65, 0x61, 0x6e, 0x5f, 0x66, 0x6c, 0x61, 0x67, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x70, 0x62, 0x2f, 0x69, 0x61, 0x2f, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d,
	0x65, 0x6e, 0x74, 0x70, 0x62, 0x2f, 0x69, 0x61, 0x2f, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74,
	0x70, 0x62, 0x2f, 0x69, 0x61, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1b, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x70, 0x62,
	0x2f, 0x73, 0x65, 0x76, 0x65, 0x72, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xa8, 0x03, 0x0a, 0x05, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x61, 0x6c, 0x65,
	0x72, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x6c, 0x65,
	0x72, 0x74, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x12, 0x30,
	0x0a, 0x08, 0x73, 0x65, 0x76, 0x65, 0x72, 0x69, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x14, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x53, 0x65,
	0x76, 0x65, 0x72, 0x69, 0x74, 0x79, 0x52, 0x08, 0x73, 0x65, 0x76, 0x65, 0x72, 0x69, 0x74, 0x79,
	0x12, 0x2a, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x12, 0x2e, 0x69, 0x61, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x35, 0x0a, 0x06,
	0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x69,
	0x61, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x2e,
	0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x6c, 0x61, 0x62,
	0x65, 0x6c, 0x73, 0x12, 0x24, 0x0a, 0x04, 0x72, 0x75, 0x6c, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x10, 0x2e, 0x69, 0x61, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x52,
	0x75, 0x6c, 0x65, 0x52, 0x04, 0x72, 0x75, 0x6c, 0x65, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f,
	0x61, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x1a,
	0x39, 0x0a, 0x0b, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x4c, 0x0a, 0x11, 0x4c, 0x69,
	0x73, 0x74, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x37, 0x0a, 0x0b, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x69, 0x61, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x31, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x52, 0x0a, 0x70, 0x61,
	0x67, 0x65, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x22, 0x6f, 0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74,
	0x41, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x29,
	0x0a, 0x06, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11,
	0x2e, 0x69, 0x61, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x41, 0x6c, 0x65, 0x72,
	0x74, 0x52, 0x06, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x12, 0x2e, 0x0a, 0x06, 0x74, 0x6f, 0x74,
	0x61, 0x6c, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x69, 0x61, 0x2e, 0x76,
	0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x74, 0x61, 0x6c,
	0x73, 0x52, 0x06, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x73, 0x22, 0x6c, 0x0a, 0x12, 0x54, 0x6f, 0x67,
	0x67, 0x6c, 0x65, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x21, 0x0a, 0x08, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x06, 0xe2, 0xdf, 0x1f, 0x02, 0x58, 0x01, 0x52, 0x07, 0x61, 0x6c, 0x65, 0x72, 0x74,
	0x49, 0x64, 0x12, 0x33, 0x0a, 0x08, 0x73, 0x69, 0x6c, 0x65, 0x6e, 0x63, 0x65, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x17, 0x2e, 0x69, 0x61, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x31, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x65, 0x61, 0x6e, 0x46, 0x6c, 0x61, 0x67, 0x52, 0x08, 0x73,
	0x69, 0x6c, 0x65, 0x6e, 0x63, 0x65, 0x64, 0x22, 0x15, 0x0a, 0x13, 0x54, 0x6f, 0x67, 0x67, 0x6c,
	0x65, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0xfb,
	0x01, 0x0a, 0x06, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x12, 0x75, 0x0a, 0x0a, 0x4c, 0x69, 0x73,
	0x74, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x12, 0x1d, 0x2e, 0x69, 0x61, 0x2e, 0x76, 0x31, 0x62,
	0x65, 0x74, 0x61, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x69, 0x61, 0x2e, 0x76, 0x31, 0x62, 0x65,
	0x74, 0x61, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x28, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x22, 0x22, 0x1d,
	0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x69,
	0x61, 0x2f, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x2f, 0x4c, 0x69, 0x73, 0x74, 0x3a, 0x01, 0x2a,
	0x12, 0x7a, 0x0a, 0x0b, 0x54, 0x6f, 0x67, 0x67, 0x6c, 0x65, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x12,
	0x1e, 0x2e, 0x69, 0x61, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x54, 0x6f, 0x67,
	0x67, 0x6c, 0x65, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1f, 0x2e, 0x69, 0x61, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x54, 0x6f, 0x67,
	0x67, 0x6c, 0x65, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x2a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x24, 0x22, 0x1f, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x61,
	0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x69, 0x61, 0x2f, 0x41, 0x6c, 0x65, 0x72,
	0x74, 0x73, 0x2f, 0x54, 0x6f, 0x67, 0x67, 0x6c, 0x65, 0x3a, 0x01, 0x2a, 0x42, 0x1f, 0x5a, 0x1d,
	0x61, 0x70, 0x69, 0x2f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x70, 0x62,
	0x2f, 0x69, 0x61, 0x3b, 0x69, 0x61, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_managementpb_ia_alerts_proto_rawDescOnce sync.Once
	file_managementpb_ia_alerts_proto_rawDescData = file_managementpb_ia_alerts_proto_rawDesc
)

func file_managementpb_ia_alerts_proto_rawDescGZIP() []byte {
	file_managementpb_ia_alerts_proto_rawDescOnce.Do(func() {
		file_managementpb_ia_alerts_proto_rawDescData = protoimpl.X.CompressGZIP(file_managementpb_ia_alerts_proto_rawDescData)
	})
	return file_managementpb_ia_alerts_proto_rawDescData
}

var file_managementpb_ia_alerts_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_managementpb_ia_alerts_proto_goTypes = []interface{}{
	(*Alert)(nil),               // 0: ia.v1beta1.Alert
	(*ListAlertsRequest)(nil),   // 1: ia.v1beta1.ListAlertsRequest
	(*ListAlertsResponse)(nil),  // 2: ia.v1beta1.ListAlertsResponse
	(*ToggleAlertRequest)(nil),  // 3: ia.v1beta1.ToggleAlertRequest
	(*ToggleAlertResponse)(nil), // 4: ia.v1beta1.ToggleAlertResponse
	nil,                         // 5: ia.v1beta1.Alert.LabelsEntry
	(managementpb.Severity)(0),  // 6: management.Severity
	(Status)(0),                 // 7: ia.v1beta1.Status
	(*Rule)(nil),                // 8: ia.v1beta1.Rule
	(*timestamp.Timestamp)(nil), // 9: google.protobuf.Timestamp
	(*PageParams)(nil),          // 10: ia.v1beta1.PageParams
	(*PageTotals)(nil),          // 11: ia.v1beta1.PageTotals
	(BooleanFlag)(0),            // 12: ia.v1beta1.BooleanFlag
}
var file_managementpb_ia_alerts_proto_depIdxs = []int32{
	6,  // 0: ia.v1beta1.Alert.severity:type_name -> management.Severity
	7,  // 1: ia.v1beta1.Alert.status:type_name -> ia.v1beta1.Status
	5,  // 2: ia.v1beta1.Alert.labels:type_name -> ia.v1beta1.Alert.LabelsEntry
	8,  // 3: ia.v1beta1.Alert.rule:type_name -> ia.v1beta1.Rule
	9,  // 4: ia.v1beta1.Alert.created_at:type_name -> google.protobuf.Timestamp
	9,  // 5: ia.v1beta1.Alert.updated_at:type_name -> google.protobuf.Timestamp
	10, // 6: ia.v1beta1.ListAlertsRequest.page_params:type_name -> ia.v1beta1.PageParams
	0,  // 7: ia.v1beta1.ListAlertsResponse.alerts:type_name -> ia.v1beta1.Alert
	11, // 8: ia.v1beta1.ListAlertsResponse.totals:type_name -> ia.v1beta1.PageTotals
	12, // 9: ia.v1beta1.ToggleAlertRequest.silenced:type_name -> ia.v1beta1.BooleanFlag
	1,  // 10: ia.v1beta1.Alerts.ListAlerts:input_type -> ia.v1beta1.ListAlertsRequest
	3,  // 11: ia.v1beta1.Alerts.ToggleAlert:input_type -> ia.v1beta1.ToggleAlertRequest
	2,  // 12: ia.v1beta1.Alerts.ListAlerts:output_type -> ia.v1beta1.ListAlertsResponse
	4,  // 13: ia.v1beta1.Alerts.ToggleAlert:output_type -> ia.v1beta1.ToggleAlertResponse
	12, // [12:14] is the sub-list for method output_type
	10, // [10:12] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_managementpb_ia_alerts_proto_init() }
func file_managementpb_ia_alerts_proto_init() {
	if File_managementpb_ia_alerts_proto != nil {
		return
	}
	file_managementpb_ia_boolean_flag_proto_init()
	file_managementpb_ia_pagination_proto_init()
	file_managementpb_ia_rules_proto_init()
	file_managementpb_ia_status_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_managementpb_ia_alerts_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Alert); i {
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
		file_managementpb_ia_alerts_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListAlertsRequest); i {
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
		file_managementpb_ia_alerts_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListAlertsResponse); i {
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
		file_managementpb_ia_alerts_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ToggleAlertRequest); i {
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
		file_managementpb_ia_alerts_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ToggleAlertResponse); i {
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
			RawDescriptor: file_managementpb_ia_alerts_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_managementpb_ia_alerts_proto_goTypes,
		DependencyIndexes: file_managementpb_ia_alerts_proto_depIdxs,
		MessageInfos:      file_managementpb_ia_alerts_proto_msgTypes,
	}.Build()
	File_managementpb_ia_alerts_proto = out.File
	file_managementpb_ia_alerts_proto_rawDesc = nil
	file_managementpb_ia_alerts_proto_goTypes = nil
	file_managementpb_ia_alerts_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// AlertsClient is the client API for Alerts service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AlertsClient interface {
	// ListAlerts returns a list of all Alerts.
	ListAlerts(ctx context.Context, in *ListAlertsRequest, opts ...grpc.CallOption) (*ListAlertsResponse, error)
	// ToggleAlert allows to switch between silenced and unsilenced states of an Alert.
	ToggleAlert(ctx context.Context, in *ToggleAlertRequest, opts ...grpc.CallOption) (*ToggleAlertResponse, error)
}

type alertsClient struct {
	cc grpc.ClientConnInterface
}

func NewAlertsClient(cc grpc.ClientConnInterface) AlertsClient {
	return &alertsClient{cc}
}

func (c *alertsClient) ListAlerts(ctx context.Context, in *ListAlertsRequest, opts ...grpc.CallOption) (*ListAlertsResponse, error) {
	out := new(ListAlertsResponse)
	err := c.cc.Invoke(ctx, "/ia.v1beta1.Alerts/ListAlerts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *alertsClient) ToggleAlert(ctx context.Context, in *ToggleAlertRequest, opts ...grpc.CallOption) (*ToggleAlertResponse, error) {
	out := new(ToggleAlertResponse)
	err := c.cc.Invoke(ctx, "/ia.v1beta1.Alerts/ToggleAlert", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AlertsServer is the server API for Alerts service.
type AlertsServer interface {
	// ListAlerts returns a list of all Alerts.
	ListAlerts(context.Context, *ListAlertsRequest) (*ListAlertsResponse, error)
	// ToggleAlert allows to switch between silenced and unsilenced states of an Alert.
	ToggleAlert(context.Context, *ToggleAlertRequest) (*ToggleAlertResponse, error)
}

// UnimplementedAlertsServer can be embedded to have forward compatible implementations.
type UnimplementedAlertsServer struct {
}

func (*UnimplementedAlertsServer) ListAlerts(context.Context, *ListAlertsRequest) (*ListAlertsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAlerts not implemented")
}
func (*UnimplementedAlertsServer) ToggleAlert(context.Context, *ToggleAlertRequest) (*ToggleAlertResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ToggleAlert not implemented")
}

func RegisterAlertsServer(s *grpc.Server, srv AlertsServer) {
	s.RegisterService(&_Alerts_serviceDesc, srv)
}

func _Alerts_ListAlerts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListAlertsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlertsServer).ListAlerts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ia.v1beta1.Alerts/ListAlerts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlertsServer).ListAlerts(ctx, req.(*ListAlertsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Alerts_ToggleAlert_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ToggleAlertRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlertsServer).ToggleAlert(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ia.v1beta1.Alerts/ToggleAlert",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlertsServer).ToggleAlert(ctx, req.(*ToggleAlertRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Alerts_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ia.v1beta1.Alerts",
	HandlerType: (*AlertsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListAlerts",
			Handler:    _Alerts_ListAlerts_Handler,
		},
		{
			MethodName: "ToggleAlert",
			Handler:    _Alerts_ToggleAlert_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "managementpb/ia/alerts.proto",
}
