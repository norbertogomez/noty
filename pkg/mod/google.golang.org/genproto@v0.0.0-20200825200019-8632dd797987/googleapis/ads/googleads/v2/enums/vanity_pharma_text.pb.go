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
// source: google/ads/googleads/v2/enums/vanity_pharma_text.proto

package enums

import (
	reflect "reflect"
	sync "sync"

	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// Enum describing possible text.
type VanityPharmaTextEnum_VanityPharmaText int32

const (
	// Not specified.
	VanityPharmaTextEnum_UNSPECIFIED VanityPharmaTextEnum_VanityPharmaText = 0
	// Used for return value only. Represents value unknown in this version.
	VanityPharmaTextEnum_UNKNOWN VanityPharmaTextEnum_VanityPharmaText = 1
	// Prescription treatment website with website content in English.
	VanityPharmaTextEnum_PRESCRIPTION_TREATMENT_WEBSITE_EN VanityPharmaTextEnum_VanityPharmaText = 2
	// Prescription treatment website with website content in Spanish
	// (Sitio de tratamientos con receta).
	VanityPharmaTextEnum_PRESCRIPTION_TREATMENT_WEBSITE_ES VanityPharmaTextEnum_VanityPharmaText = 3
	// Prescription device website with website content in English.
	VanityPharmaTextEnum_PRESCRIPTION_DEVICE_WEBSITE_EN VanityPharmaTextEnum_VanityPharmaText = 4
	// Prescription device website with website content in Spanish (Sitio de
	// dispositivos con receta).
	VanityPharmaTextEnum_PRESCRIPTION_DEVICE_WEBSITE_ES VanityPharmaTextEnum_VanityPharmaText = 5
	// Medical device website with website content in English.
	VanityPharmaTextEnum_MEDICAL_DEVICE_WEBSITE_EN VanityPharmaTextEnum_VanityPharmaText = 6
	// Medical device website with website content in Spanish (Sitio de
	// dispositivos médicos).
	VanityPharmaTextEnum_MEDICAL_DEVICE_WEBSITE_ES VanityPharmaTextEnum_VanityPharmaText = 7
	// Preventative treatment website with website content in English.
	VanityPharmaTextEnum_PREVENTATIVE_TREATMENT_WEBSITE_EN VanityPharmaTextEnum_VanityPharmaText = 8
	// Preventative treatment website with website content in Spanish (Sitio de
	// tratamientos preventivos).
	VanityPharmaTextEnum_PREVENTATIVE_TREATMENT_WEBSITE_ES VanityPharmaTextEnum_VanityPharmaText = 9
	// Prescription contraception website with website content in English.
	VanityPharmaTextEnum_PRESCRIPTION_CONTRACEPTION_WEBSITE_EN VanityPharmaTextEnum_VanityPharmaText = 10
	// Prescription contraception website with website content in Spanish (Sitio
	// de anticonceptivos con receta).
	VanityPharmaTextEnum_PRESCRIPTION_CONTRACEPTION_WEBSITE_ES VanityPharmaTextEnum_VanityPharmaText = 11
	// Prescription vaccine website with website content in English.
	VanityPharmaTextEnum_PRESCRIPTION_VACCINE_WEBSITE_EN VanityPharmaTextEnum_VanityPharmaText = 12
	// Prescription vaccine website with website content in Spanish (Sitio de
	// vacunas con receta).
	VanityPharmaTextEnum_PRESCRIPTION_VACCINE_WEBSITE_ES VanityPharmaTextEnum_VanityPharmaText = 13
)

// Enum value maps for VanityPharmaTextEnum_VanityPharmaText.
var (
	VanityPharmaTextEnum_VanityPharmaText_name = map[int32]string{
		0:  "UNSPECIFIED",
		1:  "UNKNOWN",
		2:  "PRESCRIPTION_TREATMENT_WEBSITE_EN",
		3:  "PRESCRIPTION_TREATMENT_WEBSITE_ES",
		4:  "PRESCRIPTION_DEVICE_WEBSITE_EN",
		5:  "PRESCRIPTION_DEVICE_WEBSITE_ES",
		6:  "MEDICAL_DEVICE_WEBSITE_EN",
		7:  "MEDICAL_DEVICE_WEBSITE_ES",
		8:  "PREVENTATIVE_TREATMENT_WEBSITE_EN",
		9:  "PREVENTATIVE_TREATMENT_WEBSITE_ES",
		10: "PRESCRIPTION_CONTRACEPTION_WEBSITE_EN",
		11: "PRESCRIPTION_CONTRACEPTION_WEBSITE_ES",
		12: "PRESCRIPTION_VACCINE_WEBSITE_EN",
		13: "PRESCRIPTION_VACCINE_WEBSITE_ES",
	}
	VanityPharmaTextEnum_VanityPharmaText_value = map[string]int32{
		"UNSPECIFIED":                           0,
		"UNKNOWN":                               1,
		"PRESCRIPTION_TREATMENT_WEBSITE_EN":     2,
		"PRESCRIPTION_TREATMENT_WEBSITE_ES":     3,
		"PRESCRIPTION_DEVICE_WEBSITE_EN":        4,
		"PRESCRIPTION_DEVICE_WEBSITE_ES":        5,
		"MEDICAL_DEVICE_WEBSITE_EN":             6,
		"MEDICAL_DEVICE_WEBSITE_ES":             7,
		"PREVENTATIVE_TREATMENT_WEBSITE_EN":     8,
		"PREVENTATIVE_TREATMENT_WEBSITE_ES":     9,
		"PRESCRIPTION_CONTRACEPTION_WEBSITE_EN": 10,
		"PRESCRIPTION_CONTRACEPTION_WEBSITE_ES": 11,
		"PRESCRIPTION_VACCINE_WEBSITE_EN":       12,
		"PRESCRIPTION_VACCINE_WEBSITE_ES":       13,
	}
)

func (x VanityPharmaTextEnum_VanityPharmaText) Enum() *VanityPharmaTextEnum_VanityPharmaText {
	p := new(VanityPharmaTextEnum_VanityPharmaText)
	*p = x
	return p
}

func (x VanityPharmaTextEnum_VanityPharmaText) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (VanityPharmaTextEnum_VanityPharmaText) Descriptor() protoreflect.EnumDescriptor {
	return file_google_ads_googleads_v2_enums_vanity_pharma_text_proto_enumTypes[0].Descriptor()
}

func (VanityPharmaTextEnum_VanityPharmaText) Type() protoreflect.EnumType {
	return &file_google_ads_googleads_v2_enums_vanity_pharma_text_proto_enumTypes[0]
}

func (x VanityPharmaTextEnum_VanityPharmaText) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use VanityPharmaTextEnum_VanityPharmaText.Descriptor instead.
func (VanityPharmaTextEnum_VanityPharmaText) EnumDescriptor() ([]byte, []int) {
	return file_google_ads_googleads_v2_enums_vanity_pharma_text_proto_rawDescGZIP(), []int{0, 0}
}

// The text that will be displayed in display URL of the text ad when website
// description is the selected display mode for vanity pharma URLs.
type VanityPharmaTextEnum struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *VanityPharmaTextEnum) Reset() {
	*x = VanityPharmaTextEnum{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v2_enums_vanity_pharma_text_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VanityPharmaTextEnum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VanityPharmaTextEnum) ProtoMessage() {}

func (x *VanityPharmaTextEnum) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v2_enums_vanity_pharma_text_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VanityPharmaTextEnum.ProtoReflect.Descriptor instead.
func (*VanityPharmaTextEnum) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v2_enums_vanity_pharma_text_proto_rawDescGZIP(), []int{0}
}

var File_google_ads_googleads_v2_enums_vanity_pharma_text_proto protoreflect.FileDescriptor

var file_google_ads_googleads_v2_enums_vanity_pharma_text_proto_rawDesc = []byte{
	0x0a, 0x36, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x32, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f,
	0x76, 0x61, 0x6e, 0x69, 0x74, 0x79, 0x5f, 0x70, 0x68, 0x61, 0x72, 0x6d, 0x61, 0x5f, 0x74, 0x65,
	0x78, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1d, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76,
	0x32, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8b, 0x04, 0x0a, 0x14, 0x56, 0x61, 0x6e, 0x69, 0x74, 0x79,
	0x50, 0x68, 0x61, 0x72, 0x6d, 0x61, 0x54, 0x65, 0x78, 0x74, 0x45, 0x6e, 0x75, 0x6d, 0x22, 0xf2,
	0x03, 0x0a, 0x10, 0x56, 0x61, 0x6e, 0x69, 0x74, 0x79, 0x50, 0x68, 0x61, 0x72, 0x6d, 0x61, 0x54,
	0x65, 0x78, 0x74, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49,
	0x45, 0x44, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10,
	0x01, 0x12, 0x25, 0x0a, 0x21, 0x50, 0x52, 0x45, 0x53, 0x43, 0x52, 0x49, 0x50, 0x54, 0x49, 0x4f,
	0x4e, 0x5f, 0x54, 0x52, 0x45, 0x41, 0x54, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x57, 0x45, 0x42, 0x53,
	0x49, 0x54, 0x45, 0x5f, 0x45, 0x4e, 0x10, 0x02, 0x12, 0x25, 0x0a, 0x21, 0x50, 0x52, 0x45, 0x53,
	0x43, 0x52, 0x49, 0x50, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x52, 0x45, 0x41, 0x54, 0x4d, 0x45,
	0x4e, 0x54, 0x5f, 0x57, 0x45, 0x42, 0x53, 0x49, 0x54, 0x45, 0x5f, 0x45, 0x53, 0x10, 0x03, 0x12,
	0x22, 0x0a, 0x1e, 0x50, 0x52, 0x45, 0x53, 0x43, 0x52, 0x49, 0x50, 0x54, 0x49, 0x4f, 0x4e, 0x5f,
	0x44, 0x45, 0x56, 0x49, 0x43, 0x45, 0x5f, 0x57, 0x45, 0x42, 0x53, 0x49, 0x54, 0x45, 0x5f, 0x45,
	0x4e, 0x10, 0x04, 0x12, 0x22, 0x0a, 0x1e, 0x50, 0x52, 0x45, 0x53, 0x43, 0x52, 0x49, 0x50, 0x54,
	0x49, 0x4f, 0x4e, 0x5f, 0x44, 0x45, 0x56, 0x49, 0x43, 0x45, 0x5f, 0x57, 0x45, 0x42, 0x53, 0x49,
	0x54, 0x45, 0x5f, 0x45, 0x53, 0x10, 0x05, 0x12, 0x1d, 0x0a, 0x19, 0x4d, 0x45, 0x44, 0x49, 0x43,
	0x41, 0x4c, 0x5f, 0x44, 0x45, 0x56, 0x49, 0x43, 0x45, 0x5f, 0x57, 0x45, 0x42, 0x53, 0x49, 0x54,
	0x45, 0x5f, 0x45, 0x4e, 0x10, 0x06, 0x12, 0x1d, 0x0a, 0x19, 0x4d, 0x45, 0x44, 0x49, 0x43, 0x41,
	0x4c, 0x5f, 0x44, 0x45, 0x56, 0x49, 0x43, 0x45, 0x5f, 0x57, 0x45, 0x42, 0x53, 0x49, 0x54, 0x45,
	0x5f, 0x45, 0x53, 0x10, 0x07, 0x12, 0x25, 0x0a, 0x21, 0x50, 0x52, 0x45, 0x56, 0x45, 0x4e, 0x54,
	0x41, 0x54, 0x49, 0x56, 0x45, 0x5f, 0x54, 0x52, 0x45, 0x41, 0x54, 0x4d, 0x45, 0x4e, 0x54, 0x5f,
	0x57, 0x45, 0x42, 0x53, 0x49, 0x54, 0x45, 0x5f, 0x45, 0x4e, 0x10, 0x08, 0x12, 0x25, 0x0a, 0x21,
	0x50, 0x52, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x41, 0x54, 0x49, 0x56, 0x45, 0x5f, 0x54, 0x52, 0x45,
	0x41, 0x54, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x57, 0x45, 0x42, 0x53, 0x49, 0x54, 0x45, 0x5f, 0x45,
	0x53, 0x10, 0x09, 0x12, 0x29, 0x0a, 0x25, 0x50, 0x52, 0x45, 0x53, 0x43, 0x52, 0x49, 0x50, 0x54,
	0x49, 0x4f, 0x4e, 0x5f, 0x43, 0x4f, 0x4e, 0x54, 0x52, 0x41, 0x43, 0x45, 0x50, 0x54, 0x49, 0x4f,
	0x4e, 0x5f, 0x57, 0x45, 0x42, 0x53, 0x49, 0x54, 0x45, 0x5f, 0x45, 0x4e, 0x10, 0x0a, 0x12, 0x29,
	0x0a, 0x25, 0x50, 0x52, 0x45, 0x53, 0x43, 0x52, 0x49, 0x50, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x43,
	0x4f, 0x4e, 0x54, 0x52, 0x41, 0x43, 0x45, 0x50, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x57, 0x45, 0x42,
	0x53, 0x49, 0x54, 0x45, 0x5f, 0x45, 0x53, 0x10, 0x0b, 0x12, 0x23, 0x0a, 0x1f, 0x50, 0x52, 0x45,
	0x53, 0x43, 0x52, 0x49, 0x50, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x56, 0x41, 0x43, 0x43, 0x49, 0x4e,
	0x45, 0x5f, 0x57, 0x45, 0x42, 0x53, 0x49, 0x54, 0x45, 0x5f, 0x45, 0x4e, 0x10, 0x0c, 0x12, 0x23,
	0x0a, 0x1f, 0x50, 0x52, 0x45, 0x53, 0x43, 0x52, 0x49, 0x50, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x56,
	0x41, 0x43, 0x43, 0x49, 0x4e, 0x45, 0x5f, 0x57, 0x45, 0x42, 0x53, 0x49, 0x54, 0x45, 0x5f, 0x45,
	0x53, 0x10, 0x0d, 0x42, 0xea, 0x01, 0x0a, 0x21, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73,
	0x2e, 0x76, 0x32, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x42, 0x15, 0x56, 0x61, 0x6e, 0x69, 0x74,
	0x79, 0x50, 0x68, 0x61, 0x72, 0x6d, 0x61, 0x54, 0x65, 0x78, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x50, 0x01, 0x5a, 0x42, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e,
	0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x32, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73,
	0x3b, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0xa2, 0x02, 0x03, 0x47, 0x41, 0x41, 0xaa, 0x02, 0x1d, 0x47,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x41, 0x64, 0x73, 0x2e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x41, 0x64, 0x73, 0x2e, 0x56, 0x32, 0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0xca, 0x02, 0x1d, 0x47,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x41, 0x64, 0x73, 0x5c, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x41, 0x64, 0x73, 0x5c, 0x56, 0x32, 0x5c, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0xea, 0x02, 0x21, 0x47,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x47, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x56, 0x32, 0x3a, 0x3a, 0x45, 0x6e, 0x75, 0x6d, 0x73,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_ads_googleads_v2_enums_vanity_pharma_text_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v2_enums_vanity_pharma_text_proto_rawDescData = file_google_ads_googleads_v2_enums_vanity_pharma_text_proto_rawDesc
)

func file_google_ads_googleads_v2_enums_vanity_pharma_text_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v2_enums_vanity_pharma_text_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v2_enums_vanity_pharma_text_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_ads_googleads_v2_enums_vanity_pharma_text_proto_rawDescData)
	})
	return file_google_ads_googleads_v2_enums_vanity_pharma_text_proto_rawDescData
}

var file_google_ads_googleads_v2_enums_vanity_pharma_text_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_ads_googleads_v2_enums_vanity_pharma_text_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_ads_googleads_v2_enums_vanity_pharma_text_proto_goTypes = []interface{}{
	(VanityPharmaTextEnum_VanityPharmaText)(0), // 0: google.ads.googleads.v2.enums.VanityPharmaTextEnum.VanityPharmaText
	(*VanityPharmaTextEnum)(nil),               // 1: google.ads.googleads.v2.enums.VanityPharmaTextEnum
}
var file_google_ads_googleads_v2_enums_vanity_pharma_text_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v2_enums_vanity_pharma_text_proto_init() }
func file_google_ads_googleads_v2_enums_vanity_pharma_text_proto_init() {
	if File_google_ads_googleads_v2_enums_vanity_pharma_text_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_ads_googleads_v2_enums_vanity_pharma_text_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VanityPharmaTextEnum); i {
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
			RawDescriptor: file_google_ads_googleads_v2_enums_vanity_pharma_text_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_ads_googleads_v2_enums_vanity_pharma_text_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v2_enums_vanity_pharma_text_proto_depIdxs,
		EnumInfos:         file_google_ads_googleads_v2_enums_vanity_pharma_text_proto_enumTypes,
		MessageInfos:      file_google_ads_googleads_v2_enums_vanity_pharma_text_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v2_enums_vanity_pharma_text_proto = out.File
	file_google_ads_googleads_v2_enums_vanity_pharma_text_proto_rawDesc = nil
	file_google_ads_googleads_v2_enums_vanity_pharma_text_proto_goTypes = nil
	file_google_ads_googleads_v2_enums_vanity_pharma_text_proto_depIdxs = nil
}
