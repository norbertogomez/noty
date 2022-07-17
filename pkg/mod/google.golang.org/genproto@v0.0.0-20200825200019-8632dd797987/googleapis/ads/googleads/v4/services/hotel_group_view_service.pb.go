// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.3
// source: google/ads/googleads/v4/services/hotel_group_view_service.proto

package services

import (
	context "context"
	reflect "reflect"
	sync "sync"

	proto "github.com/golang/protobuf/proto"
	resources "google.golang.org/genproto/googleapis/ads/googleads/v4/resources"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

// Request message for [HotelGroupViewService.GetHotelGroupView][google.ads.googleads.v4.services.HotelGroupViewService.GetHotelGroupView].
type GetHotelGroupViewRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. Resource name of the Hotel Group View to fetch.
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
}

func (x *GetHotelGroupViewRequest) Reset() {
	*x = GetHotelGroupViewRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v4_services_hotel_group_view_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetHotelGroupViewRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetHotelGroupViewRequest) ProtoMessage() {}

func (x *GetHotelGroupViewRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v4_services_hotel_group_view_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetHotelGroupViewRequest.ProtoReflect.Descriptor instead.
func (*GetHotelGroupViewRequest) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v4_services_hotel_group_view_service_proto_rawDescGZIP(), []int{0}
}

func (x *GetHotelGroupViewRequest) GetResourceName() string {
	if x != nil {
		return x.ResourceName
	}
	return ""
}

var File_google_ads_googleads_v4_services_hotel_group_view_service_proto protoreflect.FileDescriptor

var file_google_ads_googleads_v4_services_hotel_group_view_service_proto_rawDesc = []byte{
	0x0a, 0x3f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x34, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2f, 0x68, 0x6f, 0x74, 0x65, 0x6c, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x76,
	0x69, 0x65, 0x77, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x34, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x1a, 0x38, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x34, 0x2f, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x68, 0x6f, 0x74, 0x65, 0x6c, 0x5f, 0x67, 0x72, 0x6f,
	0x75, 0x70, 0x5f, 0x76, 0x69, 0x65, 0x77, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x70, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x48, 0x6f, 0x74, 0x65, 0x6c, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x56, 0x69, 0x65, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x54, 0x0a, 0x0d,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x2f, 0xe0, 0x41, 0x02, 0xfa, 0x41, 0x29, 0x0a, 0x27, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69,
	0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x48, 0x6f, 0x74, 0x65, 0x6c, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x56, 0x69, 0x65, 0x77, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4e, 0x61,
	0x6d, 0x65, 0x32, 0x84, 0x02, 0x0a, 0x15, 0x48, 0x6f, 0x74, 0x65, 0x6c, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x56, 0x69, 0x65, 0x77, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0xcd, 0x01, 0x0a,
	0x11, 0x47, 0x65, 0x74, 0x48, 0x6f, 0x74, 0x65, 0x6c, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x56, 0x69,
	0x65, 0x77, 0x12, 0x3a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x34, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x48, 0x6f, 0x74, 0x65, 0x6c, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x56, 0x69, 0x65, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x31,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x34, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x73, 0x2e, 0x48, 0x6f, 0x74, 0x65, 0x6c, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x56, 0x69, 0x65,
	0x77, 0x22, 0x49, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x33, 0x12, 0x31, 0x2f, 0x76, 0x34, 0x2f, 0x7b,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x3d, 0x63, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73, 0x2f, 0x2a, 0x2f, 0x68, 0x6f, 0x74, 0x65, 0x6c, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x56, 0x69, 0x65, 0x77, 0x73, 0x2f, 0x2a, 0x7d, 0xda, 0x41, 0x0d, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x1a, 0x1b, 0xca, 0x41,
	0x18, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x42, 0x81, 0x02, 0x0a, 0x24, 0x63, 0x6f,
	0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x34, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x42, 0x1a, 0x48, 0x6f, 0x74, 0x65, 0x6c, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x56, 0x69,
	0x65, 0x77, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x48, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e,
	0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x34, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x3b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0xa2, 0x02, 0x03, 0x47, 0x41, 0x41,
	0xaa, 0x02, 0x20, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x41, 0x64, 0x73, 0x2e, 0x47, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x2e, 0x56, 0x34, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0xca, 0x02, 0x20, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x41, 0x64, 0x73,
	0x5c, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x5c, 0x56, 0x34, 0x5c, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0xea, 0x02, 0x24, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a,
	0x3a, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x3a,
	0x3a, 0x56, 0x34, 0x3a, 0x3a, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_ads_googleads_v4_services_hotel_group_view_service_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v4_services_hotel_group_view_service_proto_rawDescData = file_google_ads_googleads_v4_services_hotel_group_view_service_proto_rawDesc
)

func file_google_ads_googleads_v4_services_hotel_group_view_service_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v4_services_hotel_group_view_service_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v4_services_hotel_group_view_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_ads_googleads_v4_services_hotel_group_view_service_proto_rawDescData)
	})
	return file_google_ads_googleads_v4_services_hotel_group_view_service_proto_rawDescData
}

var file_google_ads_googleads_v4_services_hotel_group_view_service_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_ads_googleads_v4_services_hotel_group_view_service_proto_goTypes = []interface{}{
	(*GetHotelGroupViewRequest)(nil), // 0: google.ads.googleads.v4.services.GetHotelGroupViewRequest
	(*resources.HotelGroupView)(nil), // 1: google.ads.googleads.v4.resources.HotelGroupView
}
var file_google_ads_googleads_v4_services_hotel_group_view_service_proto_depIdxs = []int32{
	0, // 0: google.ads.googleads.v4.services.HotelGroupViewService.GetHotelGroupView:input_type -> google.ads.googleads.v4.services.GetHotelGroupViewRequest
	1, // 1: google.ads.googleads.v4.services.HotelGroupViewService.GetHotelGroupView:output_type -> google.ads.googleads.v4.resources.HotelGroupView
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v4_services_hotel_group_view_service_proto_init() }
func file_google_ads_googleads_v4_services_hotel_group_view_service_proto_init() {
	if File_google_ads_googleads_v4_services_hotel_group_view_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_ads_googleads_v4_services_hotel_group_view_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetHotelGroupViewRequest); i {
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
			RawDescriptor: file_google_ads_googleads_v4_services_hotel_group_view_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_google_ads_googleads_v4_services_hotel_group_view_service_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v4_services_hotel_group_view_service_proto_depIdxs,
		MessageInfos:      file_google_ads_googleads_v4_services_hotel_group_view_service_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v4_services_hotel_group_view_service_proto = out.File
	file_google_ads_googleads_v4_services_hotel_group_view_service_proto_rawDesc = nil
	file_google_ads_googleads_v4_services_hotel_group_view_service_proto_goTypes = nil
	file_google_ads_googleads_v4_services_hotel_group_view_service_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// HotelGroupViewServiceClient is the client API for HotelGroupViewService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type HotelGroupViewServiceClient interface {
	// Returns the requested Hotel Group View in full detail.
	GetHotelGroupView(ctx context.Context, in *GetHotelGroupViewRequest, opts ...grpc.CallOption) (*resources.HotelGroupView, error)
}

type hotelGroupViewServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewHotelGroupViewServiceClient(cc grpc.ClientConnInterface) HotelGroupViewServiceClient {
	return &hotelGroupViewServiceClient{cc}
}

func (c *hotelGroupViewServiceClient) GetHotelGroupView(ctx context.Context, in *GetHotelGroupViewRequest, opts ...grpc.CallOption) (*resources.HotelGroupView, error) {
	out := new(resources.HotelGroupView)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v4.services.HotelGroupViewService/GetHotelGroupView", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HotelGroupViewServiceServer is the server API for HotelGroupViewService service.
type HotelGroupViewServiceServer interface {
	// Returns the requested Hotel Group View in full detail.
	GetHotelGroupView(context.Context, *GetHotelGroupViewRequest) (*resources.HotelGroupView, error)
}

// UnimplementedHotelGroupViewServiceServer can be embedded to have forward compatible implementations.
type UnimplementedHotelGroupViewServiceServer struct {
}

func (*UnimplementedHotelGroupViewServiceServer) GetHotelGroupView(context.Context, *GetHotelGroupViewRequest) (*resources.HotelGroupView, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetHotelGroupView not implemented")
}

func RegisterHotelGroupViewServiceServer(s *grpc.Server, srv HotelGroupViewServiceServer) {
	s.RegisterService(&_HotelGroupViewService_serviceDesc, srv)
}

func _HotelGroupViewService_GetHotelGroupView_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetHotelGroupViewRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HotelGroupViewServiceServer).GetHotelGroupView(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v4.services.HotelGroupViewService/GetHotelGroupView",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HotelGroupViewServiceServer).GetHotelGroupView(ctx, req.(*GetHotelGroupViewRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _HotelGroupViewService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v4.services.HotelGroupViewService",
	HandlerType: (*HotelGroupViewServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetHotelGroupView",
			Handler:    _HotelGroupViewService_GetHotelGroupView_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v4/services/hotel_group_view_service.proto",
}