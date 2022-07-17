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
// source: google/maps/routes/v1/route_matrix_element.proto

package routes

import (
	reflect "reflect"
	sync "sync"

	proto "github.com/golang/protobuf/proto"
	status "google.golang.org/genproto/googleapis/rpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
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

// The condition of the route being returned.
type RouteMatrixElementCondition int32

const (
	// Not used.
	RouteMatrixElementCondition_ROUTE_MATRIX_ELEMENT_CONDITION_UNSPECIFIED RouteMatrixElementCondition = 0
	// A route was found, and the corresponding information was filled out for the
	// element.
	RouteMatrixElementCondition_ROUTE_EXISTS RouteMatrixElementCondition = 1
	// No route could be found. Fields containing route information, such as
	// distance_meters or duration, will not be filled out in the element.
	RouteMatrixElementCondition_ROUTE_NOT_FOUND RouteMatrixElementCondition = 2
)

// Enum value maps for RouteMatrixElementCondition.
var (
	RouteMatrixElementCondition_name = map[int32]string{
		0: "ROUTE_MATRIX_ELEMENT_CONDITION_UNSPECIFIED",
		1: "ROUTE_EXISTS",
		2: "ROUTE_NOT_FOUND",
	}
	RouteMatrixElementCondition_value = map[string]int32{
		"ROUTE_MATRIX_ELEMENT_CONDITION_UNSPECIFIED": 0,
		"ROUTE_EXISTS":    1,
		"ROUTE_NOT_FOUND": 2,
	}
)

func (x RouteMatrixElementCondition) Enum() *RouteMatrixElementCondition {
	p := new(RouteMatrixElementCondition)
	*p = x
	return p
}

func (x RouteMatrixElementCondition) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RouteMatrixElementCondition) Descriptor() protoreflect.EnumDescriptor {
	return file_google_maps_routes_v1_route_matrix_element_proto_enumTypes[0].Descriptor()
}

func (RouteMatrixElementCondition) Type() protoreflect.EnumType {
	return &file_google_maps_routes_v1_route_matrix_element_proto_enumTypes[0]
}

func (x RouteMatrixElementCondition) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RouteMatrixElementCondition.Descriptor instead.
func (RouteMatrixElementCondition) EnumDescriptor() ([]byte, []int) {
	return file_google_maps_routes_v1_route_matrix_element_proto_rawDescGZIP(), []int{0}
}

// Encapsulates route information computed for an origin/destination pair in the
// ComputeRouteMatrix API. This proto can be streamed to the client.
type RouteMatrixElement struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Zero-based index of the origin in the request.
	OriginIndex int32 `protobuf:"varint,1,opt,name=origin_index,json=originIndex,proto3" json:"origin_index,omitempty"`
	// Zero-based index of the destination in the request.
	DestinationIndex int32 `protobuf:"varint,2,opt,name=destination_index,json=destinationIndex,proto3" json:"destination_index,omitempty"`
	// Error status code for this element.
	Status *status.Status `protobuf:"bytes,3,opt,name=status,proto3" json:"status,omitempty"`
	// Indicates whether the route was found or not. Independent of status.
	Condition RouteMatrixElementCondition `protobuf:"varint,9,opt,name=condition,proto3,enum=google.maps.routes.v1.RouteMatrixElementCondition" json:"condition,omitempty"`
	// The travel distance of the route, in meters.
	DistanceMeters int32 `protobuf:"varint,4,opt,name=distance_meters,json=distanceMeters,proto3" json:"distance_meters,omitempty"`
	// The length of time needed to navigate the route. If you set the
	// `route_preference` to `TRAFFIC_UNAWARE`, then this value is the same as
	// `static_duration`. If you set the `route_preference` to either
	// `TRAFFIC_AWARE` or `TRAFFIC_AWARE_OPTIMAL`, then this value is calculated
	// taking traffic conditions into account.
	Duration *durationpb.Duration `protobuf:"bytes,5,opt,name=duration,proto3" json:"duration,omitempty"`
	// The duration of traveling through the route without taking traffic
	// conditions into consideration.
	StaticDuration *durationpb.Duration `protobuf:"bytes,6,opt,name=static_duration,json=staticDuration,proto3" json:"static_duration,omitempty"`
	// Additional information about the route. For example: restriction
	// information and toll information
	TravelAdvisory *RouteTravelAdvisory `protobuf:"bytes,7,opt,name=travel_advisory,json=travelAdvisory,proto3" json:"travel_advisory,omitempty"`
	// In some cases when the server is not able to compute the route with the
	// given preferences for this particular origin/destination pair, it may
	// fall back to using a different mode of computation. When fallback mode is
	// used, this field contains detailed information about the fallback response.
	// Otherwise this field is unset.
	FallbackInfo *FallbackInfo `protobuf:"bytes,8,opt,name=fallback_info,json=fallbackInfo,proto3" json:"fallback_info,omitempty"`
}

func (x *RouteMatrixElement) Reset() {
	*x = RouteMatrixElement{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_maps_routes_v1_route_matrix_element_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RouteMatrixElement) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RouteMatrixElement) ProtoMessage() {}

func (x *RouteMatrixElement) ProtoReflect() protoreflect.Message {
	mi := &file_google_maps_routes_v1_route_matrix_element_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RouteMatrixElement.ProtoReflect.Descriptor instead.
func (*RouteMatrixElement) Descriptor() ([]byte, []int) {
	return file_google_maps_routes_v1_route_matrix_element_proto_rawDescGZIP(), []int{0}
}

func (x *RouteMatrixElement) GetOriginIndex() int32 {
	if x != nil {
		return x.OriginIndex
	}
	return 0
}

func (x *RouteMatrixElement) GetDestinationIndex() int32 {
	if x != nil {
		return x.DestinationIndex
	}
	return 0
}

func (x *RouteMatrixElement) GetStatus() *status.Status {
	if x != nil {
		return x.Status
	}
	return nil
}

func (x *RouteMatrixElement) GetCondition() RouteMatrixElementCondition {
	if x != nil {
		return x.Condition
	}
	return RouteMatrixElementCondition_ROUTE_MATRIX_ELEMENT_CONDITION_UNSPECIFIED
}

func (x *RouteMatrixElement) GetDistanceMeters() int32 {
	if x != nil {
		return x.DistanceMeters
	}
	return 0
}

func (x *RouteMatrixElement) GetDuration() *durationpb.Duration {
	if x != nil {
		return x.Duration
	}
	return nil
}

func (x *RouteMatrixElement) GetStaticDuration() *durationpb.Duration {
	if x != nil {
		return x.StaticDuration
	}
	return nil
}

func (x *RouteMatrixElement) GetTravelAdvisory() *RouteTravelAdvisory {
	if x != nil {
		return x.TravelAdvisory
	}
	return nil
}

func (x *RouteMatrixElement) GetFallbackInfo() *FallbackInfo {
	if x != nil {
		return x.FallbackInfo
	}
	return nil
}

var File_google_maps_routes_v1_route_matrix_element_proto protoreflect.FileDescriptor

var file_google_maps_routes_v1_route_matrix_element_proto_rawDesc = []byte{
	0x0a, 0x30, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x6d, 0x61, 0x70, 0x73, 0x2f, 0x72, 0x6f,
	0x75, 0x74, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x5f, 0x6d, 0x61,
	0x74, 0x72, 0x69, 0x78, 0x5f, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x15, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x6d, 0x61, 0x70, 0x73, 0x2e,
	0x72, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x29, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x6d, 0x61, 0x70, 0x73, 0x2f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x2f, 0x76, 0x31,
	0x2f, 0x66, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x6d, 0x61, 0x70,
	0x73, 0x2f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x6f, 0x75, 0x74,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x72, 0x70, 0x63, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xa5, 0x04, 0x0a, 0x12, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x4d, 0x61, 0x74, 0x72, 0x69, 0x78,
	0x45, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x6f, 0x72, 0x69, 0x67, 0x69,
	0x6e, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x6f,
	0x72, 0x69, 0x67, 0x69, 0x6e, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x2b, 0x0a, 0x11, 0x64, 0x65,
	0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x10, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x2a, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x72, 0x70, 0x63, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x50, 0x0a, 0x09, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x32, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x6d, 0x61, 0x70, 0x73, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x52,
	0x6f, 0x75, 0x74, 0x65, 0x4d, 0x61, 0x74, 0x72, 0x69, 0x78, 0x45, 0x6c, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x09, 0x63, 0x6f, 0x6e, 0x64,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x27, 0x0a, 0x0f, 0x64, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x63,
	0x65, 0x5f, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0e,
	0x64, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x4d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x12, 0x35,
	0x0a, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x64, 0x75, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x42, 0x0a, 0x0f, 0x73, 0x74, 0x61, 0x74, 0x69, 0x63, 0x5f,
	0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0e, 0x73, 0x74, 0x61, 0x74, 0x69,
	0x63, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x53, 0x0a, 0x0f, 0x74, 0x72, 0x61,
	0x76, 0x65, 0x6c, 0x5f, 0x61, 0x64, 0x76, 0x69, 0x73, 0x6f, 0x72, 0x79, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x6d, 0x61, 0x70, 0x73,
	0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x6f, 0x75, 0x74, 0x65,
	0x54, 0x72, 0x61, 0x76, 0x65, 0x6c, 0x41, 0x64, 0x76, 0x69, 0x73, 0x6f, 0x72, 0x79, 0x52, 0x0e,
	0x74, 0x72, 0x61, 0x76, 0x65, 0x6c, 0x41, 0x64, 0x76, 0x69, 0x73, 0x6f, 0x72, 0x79, 0x12, 0x48,
	0x0a, 0x0d, 0x66, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x6d,
	0x61, 0x70, 0x73, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x61,
	0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0c, 0x66, 0x61, 0x6c, 0x6c,
	0x62, 0x61, 0x63, 0x6b, 0x49, 0x6e, 0x66, 0x6f, 0x2a, 0x74, 0x0a, 0x1b, 0x52, 0x6f, 0x75, 0x74,
	0x65, 0x4d, 0x61, 0x74, 0x72, 0x69, 0x78, 0x45, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x43, 0x6f,
	0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2e, 0x0a, 0x2a, 0x52, 0x4f, 0x55, 0x54, 0x45,
	0x5f, 0x4d, 0x41, 0x54, 0x52, 0x49, 0x58, 0x5f, 0x45, 0x4c, 0x45, 0x4d, 0x45, 0x4e, 0x54, 0x5f,
	0x43, 0x4f, 0x4e, 0x44, 0x49, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43,
	0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x10, 0x0a, 0x0c, 0x52, 0x4f, 0x55, 0x54, 0x45,
	0x5f, 0x45, 0x58, 0x49, 0x53, 0x54, 0x53, 0x10, 0x01, 0x12, 0x13, 0x0a, 0x0f, 0x52, 0x4f, 0x55,
	0x54, 0x45, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x46, 0x4f, 0x55, 0x4e, 0x44, 0x10, 0x02, 0x42, 0xb4,
	0x01, 0x0a, 0x19, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x6d, 0x61,
	0x70, 0x73, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x42, 0x1e, 0x43, 0x6f,
	0x6d, 0x70, 0x75, 0x74, 0x65, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x4d, 0x61, 0x74, 0x72, 0x69, 0x78,
	0x45, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3b,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72,
	0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x6d, 0x61, 0x70, 0x73, 0x2f, 0x72, 0x6f, 0x75, 0x74, 0x65,
	0x73, 0x2f, 0x76, 0x31, 0x3b, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x73, 0xf8, 0x01, 0x01, 0xa2, 0x02,
	0x04, 0x47, 0x4d, 0x52, 0x53, 0xaa, 0x02, 0x15, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x4d,
	0x61, 0x70, 0x73, 0x2e, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x15,
	0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x4d, 0x61, 0x70, 0x73, 0x5c, 0x52, 0x6f, 0x75, 0x74,
	0x65, 0x73, 0x5c, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_maps_routes_v1_route_matrix_element_proto_rawDescOnce sync.Once
	file_google_maps_routes_v1_route_matrix_element_proto_rawDescData = file_google_maps_routes_v1_route_matrix_element_proto_rawDesc
)

func file_google_maps_routes_v1_route_matrix_element_proto_rawDescGZIP() []byte {
	file_google_maps_routes_v1_route_matrix_element_proto_rawDescOnce.Do(func() {
		file_google_maps_routes_v1_route_matrix_element_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_maps_routes_v1_route_matrix_element_proto_rawDescData)
	})
	return file_google_maps_routes_v1_route_matrix_element_proto_rawDescData
}

var file_google_maps_routes_v1_route_matrix_element_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_maps_routes_v1_route_matrix_element_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_maps_routes_v1_route_matrix_element_proto_goTypes = []interface{}{
	(RouteMatrixElementCondition)(0), // 0: google.maps.routes.v1.RouteMatrixElementCondition
	(*RouteMatrixElement)(nil),       // 1: google.maps.routes.v1.RouteMatrixElement
	(*status.Status)(nil),            // 2: google.rpc.Status
	(*durationpb.Duration)(nil),      // 3: google.protobuf.Duration
	(*RouteTravelAdvisory)(nil),      // 4: google.maps.routes.v1.RouteTravelAdvisory
	(*FallbackInfo)(nil),             // 5: google.maps.routes.v1.FallbackInfo
}
var file_google_maps_routes_v1_route_matrix_element_proto_depIdxs = []int32{
	2, // 0: google.maps.routes.v1.RouteMatrixElement.status:type_name -> google.rpc.Status
	0, // 1: google.maps.routes.v1.RouteMatrixElement.condition:type_name -> google.maps.routes.v1.RouteMatrixElementCondition
	3, // 2: google.maps.routes.v1.RouteMatrixElement.duration:type_name -> google.protobuf.Duration
	3, // 3: google.maps.routes.v1.RouteMatrixElement.static_duration:type_name -> google.protobuf.Duration
	4, // 4: google.maps.routes.v1.RouteMatrixElement.travel_advisory:type_name -> google.maps.routes.v1.RouteTravelAdvisory
	5, // 5: google.maps.routes.v1.RouteMatrixElement.fallback_info:type_name -> google.maps.routes.v1.FallbackInfo
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_google_maps_routes_v1_route_matrix_element_proto_init() }
func file_google_maps_routes_v1_route_matrix_element_proto_init() {
	if File_google_maps_routes_v1_route_matrix_element_proto != nil {
		return
	}
	file_google_maps_routes_v1_fallback_info_proto_init()
	file_google_maps_routes_v1_route_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_google_maps_routes_v1_route_matrix_element_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RouteMatrixElement); i {
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
			RawDescriptor: file_google_maps_routes_v1_route_matrix_element_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_maps_routes_v1_route_matrix_element_proto_goTypes,
		DependencyIndexes: file_google_maps_routes_v1_route_matrix_element_proto_depIdxs,
		EnumInfos:         file_google_maps_routes_v1_route_matrix_element_proto_enumTypes,
		MessageInfos:      file_google_maps_routes_v1_route_matrix_element_proto_msgTypes,
	}.Build()
	File_google_maps_routes_v1_route_matrix_element_proto = out.File
	file_google_maps_routes_v1_route_matrix_element_proto_rawDesc = nil
	file_google_maps_routes_v1_route_matrix_element_proto_goTypes = nil
	file_google_maps_routes_v1_route_matrix_element_proto_depIdxs = nil
}