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
// source: google/ads/googleads/v2/resources/click_view.proto

package resources

import (
	reflect "reflect"
	sync "sync"

	proto "github.com/golang/protobuf/proto"
	common "google.golang.org/genproto/googleapis/ads/googleads/v2/common"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
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

// A click view with metrics aggregated at each click level, including both
// valid and invalid clicks. For non-Search campaigns, metrics.clicks
// represents the number of valid and invalid interactions.
// Queries including ClickView must have a filter limiting the results to one
// day and can be requested for dates back to 90 days before the time of the
// request.
type ClickView struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Output only. The resource name of the click view.
	// Click view resource names have the form:
	//
	// `customers/{customer_id}/clickViews/{date (yyyy-MM-dd)}~{gclid}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// Output only. The Google Click ID.
	Gclid *wrapperspb.StringValue `protobuf:"bytes,2,opt,name=gclid,proto3" json:"gclid,omitempty"`
	// Output only. The location criteria matching the area of interest associated with the
	// impression.
	AreaOfInterest *common.ClickLocation `protobuf:"bytes,3,opt,name=area_of_interest,json=areaOfInterest,proto3" json:"area_of_interest,omitempty"`
	// Output only. The location criteria matching the location of presence associated with the
	// impression.
	LocationOfPresence *common.ClickLocation `protobuf:"bytes,4,opt,name=location_of_presence,json=locationOfPresence,proto3" json:"location_of_presence,omitempty"`
	// Output only. Page number in search results where the ad was shown.
	PageNumber *wrapperspb.Int64Value `protobuf:"bytes,5,opt,name=page_number,json=pageNumber,proto3" json:"page_number,omitempty"`
	// Output only. The associated ad.
	AdGroupAd *wrapperspb.StringValue `protobuf:"bytes,7,opt,name=ad_group_ad,json=adGroupAd,proto3" json:"ad_group_ad,omitempty"`
}

func (x *ClickView) Reset() {
	*x = ClickView{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v2_resources_click_view_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClickView) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClickView) ProtoMessage() {}

func (x *ClickView) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v2_resources_click_view_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClickView.ProtoReflect.Descriptor instead.
func (*ClickView) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v2_resources_click_view_proto_rawDescGZIP(), []int{0}
}

func (x *ClickView) GetResourceName() string {
	if x != nil {
		return x.ResourceName
	}
	return ""
}

func (x *ClickView) GetGclid() *wrapperspb.StringValue {
	if x != nil {
		return x.Gclid
	}
	return nil
}

func (x *ClickView) GetAreaOfInterest() *common.ClickLocation {
	if x != nil {
		return x.AreaOfInterest
	}
	return nil
}

func (x *ClickView) GetLocationOfPresence() *common.ClickLocation {
	if x != nil {
		return x.LocationOfPresence
	}
	return nil
}

func (x *ClickView) GetPageNumber() *wrapperspb.Int64Value {
	if x != nil {
		return x.PageNumber
	}
	return nil
}

func (x *ClickView) GetAdGroupAd() *wrapperspb.StringValue {
	if x != nil {
		return x.AdGroupAd
	}
	return nil
}

var File_google_ads_googleads_v2_resources_click_view_proto protoreflect.FileDescriptor

var file_google_ads_googleads_v2_resources_click_view_proto_rawDesc = []byte{
	0x0a, 0x32, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x32, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x73, 0x2f, 0x63, 0x6c, 0x69, 0x63, 0x6b, 0x5f, 0x76, 0x69, 0x65, 0x77, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x21, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x32, 0x2e, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x1a, 0x33, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x32,
	0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63, 0x6c, 0x69, 0x63, 0x6b, 0x5f, 0x6c, 0x6f,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62,
	0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65,
	0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xdd, 0x04, 0x0a, 0x09, 0x43, 0x6c, 0x69, 0x63, 0x6b,
	0x56, 0x69, 0x65, 0x77, 0x12, 0x4f, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x2a, 0xe0, 0x41, 0x03,
	0xfa, 0x41, 0x24, 0x0a, 0x22, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x43, 0x6c,
	0x69, 0x63, 0x6b, 0x56, 0x69, 0x65, 0x77, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x37, 0x0a, 0x05, 0x67, 0x63, 0x6c, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x05, 0x67, 0x63, 0x6c, 0x69, 0x64, 0x12, 0x5c,
	0x0a, 0x10, 0x61, 0x72, 0x65, 0x61, 0x5f, 0x6f, 0x66, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x65,
	0x73, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e,
	0x76, 0x32, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x43, 0x6c, 0x69, 0x63, 0x6b, 0x4c,
	0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x0e, 0x61, 0x72,
	0x65, 0x61, 0x4f, 0x66, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x65, 0x73, 0x74, 0x12, 0x64, 0x0a, 0x14,
	0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6f, 0x66, 0x5f, 0x70, 0x72, 0x65, 0x73,
	0x65, 0x6e, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64,
	0x73, 0x2e, 0x76, 0x32, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x43, 0x6c, 0x69, 0x63,
	0x6b, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x12,
	0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x66, 0x50, 0x72, 0x65, 0x73, 0x65, 0x6e,
	0x63, 0x65, 0x12, 0x41, 0x0a, 0x0b, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x4e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x68, 0x0a, 0x0b, 0x61, 0x64, 0x5f, 0x67, 0x72, 0x6f, 0x75,
	0x70, 0x5f, 0x61, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72,
	0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x2a, 0xe0, 0x41, 0x03, 0xfa, 0x41, 0x24,
	0x0a, 0x22, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x41, 0x64, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x41, 0x64, 0x52, 0x09, 0x61, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x41, 0x64, 0x3a,
	0x55, 0xea, 0x41, 0x52, 0x0a, 0x22, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x43,
	0x6c, 0x69, 0x63, 0x6b, 0x56, 0x69, 0x65, 0x77, 0x12, 0x2c, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x65, 0x72, 0x73, 0x2f, 0x7b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x7d, 0x2f, 0x63,
	0x6c, 0x69, 0x63, 0x6b, 0x56, 0x69, 0x65, 0x77, 0x73, 0x2f, 0x7b, 0x63, 0x6c, 0x69, 0x63, 0x6b,
	0x5f, 0x76, 0x69, 0x65, 0x77, 0x7d, 0x42, 0xfb, 0x01, 0x0a, 0x25, 0x63, 0x6f, 0x6d, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x61, 0x64, 0x73, 0x2e, 0x76, 0x32, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73,
	0x42, 0x0e, 0x43, 0x6c, 0x69, 0x63, 0x6b, 0x56, 0x69, 0x65, 0x77, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x50, 0x01, 0x5a, 0x4a, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e,
	0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x32, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x73, 0x3b, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0xa2, 0x02,
	0x03, 0x47, 0x41, 0x41, 0xaa, 0x02, 0x21, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x41, 0x64,
	0x73, 0x2e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x2e, 0x56, 0x32, 0x2e, 0x52,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0xca, 0x02, 0x21, 0x47, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x5c, 0x41, 0x64, 0x73, 0x5c, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x5c,
	0x56, 0x32, 0x5c, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0xea, 0x02, 0x25, 0x47,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x47, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x56, 0x32, 0x3a, 0x3a, 0x52, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_ads_googleads_v2_resources_click_view_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v2_resources_click_view_proto_rawDescData = file_google_ads_googleads_v2_resources_click_view_proto_rawDesc
)

func file_google_ads_googleads_v2_resources_click_view_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v2_resources_click_view_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v2_resources_click_view_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_ads_googleads_v2_resources_click_view_proto_rawDescData)
	})
	return file_google_ads_googleads_v2_resources_click_view_proto_rawDescData
}

var file_google_ads_googleads_v2_resources_click_view_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_ads_googleads_v2_resources_click_view_proto_goTypes = []interface{}{
	(*ClickView)(nil),              // 0: google.ads.googleads.v2.resources.ClickView
	(*wrapperspb.StringValue)(nil), // 1: google.protobuf.StringValue
	(*common.ClickLocation)(nil),   // 2: google.ads.googleads.v2.common.ClickLocation
	(*wrapperspb.Int64Value)(nil),  // 3: google.protobuf.Int64Value
}
var file_google_ads_googleads_v2_resources_click_view_proto_depIdxs = []int32{
	1, // 0: google.ads.googleads.v2.resources.ClickView.gclid:type_name -> google.protobuf.StringValue
	2, // 1: google.ads.googleads.v2.resources.ClickView.area_of_interest:type_name -> google.ads.googleads.v2.common.ClickLocation
	2, // 2: google.ads.googleads.v2.resources.ClickView.location_of_presence:type_name -> google.ads.googleads.v2.common.ClickLocation
	3, // 3: google.ads.googleads.v2.resources.ClickView.page_number:type_name -> google.protobuf.Int64Value
	1, // 4: google.ads.googleads.v2.resources.ClickView.ad_group_ad:type_name -> google.protobuf.StringValue
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v2_resources_click_view_proto_init() }
func file_google_ads_googleads_v2_resources_click_view_proto_init() {
	if File_google_ads_googleads_v2_resources_click_view_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_ads_googleads_v2_resources_click_view_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClickView); i {
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
			RawDescriptor: file_google_ads_googleads_v2_resources_click_view_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_ads_googleads_v2_resources_click_view_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v2_resources_click_view_proto_depIdxs,
		MessageInfos:      file_google_ads_googleads_v2_resources_click_view_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v2_resources_click_view_proto = out.File
	file_google_ads_googleads_v2_resources_click_view_proto_rawDesc = nil
	file_google_ads_googleads_v2_resources_click_view_proto_goTypes = nil
	file_google_ads_googleads_v2_resources_click_view_proto_depIdxs = nil
}
