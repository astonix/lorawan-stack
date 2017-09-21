// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/TheThingsNetwork/ttn/api/rights.proto

package ttnpb

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Right is the enum that defines all the different rights to do something in
// the network.
type Right int32

const (
	RightApplicationDelete        Right = 0
	RightApplicationSettings      Right = 1
	RightApplicationCollaborators Right = 2
	RightGatewayOwner             Right = 10
)

var Right_name = map[int32]string{
	0:  "APPLICATION_DELETE",
	1:  "APPLICATION_SETTINGS",
	2:  "APPLICATION_COLLABORATORS",
	10: "GATEWAY_OWNER",
}
var Right_value = map[string]int32{
	"APPLICATION_DELETE":        0,
	"APPLICATION_SETTINGS":      1,
	"APPLICATION_COLLABORATORS": 2,
	"GATEWAY_OWNER":             10,
}

func (Right) EnumDescriptor() ([]byte, []int) { return fileDescriptorRights, []int{0} }

func init() {
	proto.RegisterEnum("ttn.v3.Right", Right_name, Right_value)
}

func init() {
	proto.RegisterFile("github.com/TheThingsNetwork/ttn/api/rights.proto", fileDescriptorRights)
}

var fileDescriptorRights = []byte{
	// 331 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x32, 0x48, 0xcf, 0x2c, 0xc9,
	0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x0f, 0xc9, 0x48, 0x0d, 0xc9, 0xc8, 0xcc, 0x4b, 0x2f,
	0xf6, 0x4b, 0x2d, 0x29, 0xcf, 0x2f, 0xca, 0xd6, 0x2f, 0x29, 0xc9, 0xd3, 0x4f, 0x2c, 0xc8, 0xd4,
	0x2f, 0xca, 0x4c, 0xcf, 0x28, 0x29, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2b, 0x29,
	0xc9, 0xd3, 0x2b, 0x33, 0x96, 0xd2, 0x45, 0xd2, 0x99, 0x9e, 0x9f, 0x9e, 0xaf, 0x0f, 0x96, 0x4e,
	0x2a, 0x4d, 0x03, 0xf3, 0xc0, 0x1c, 0x30, 0x0b, 0xa2, 0x4d, 0xeb, 0x25, 0x23, 0x17, 0x6b, 0x10,
	0xc8, 0x1c, 0x21, 0x23, 0x2e, 0x21, 0xc7, 0x80, 0x00, 0x1f, 0x4f, 0x67, 0xc7, 0x10, 0x4f, 0x7f,
	0xbf, 0x78, 0x17, 0x57, 0x1f, 0xd7, 0x10, 0x57, 0x01, 0x06, 0x29, 0xa9, 0xae, 0xb9, 0x0a, 0x62,
	0x60, 0x25, 0x8e, 0x05, 0x05, 0x39, 0x99, 0xc9, 0x89, 0x25, 0x99, 0xf9, 0x79, 0x2e, 0xa9, 0x39,
	0xa9, 0x25, 0xa9, 0x42, 0x66, 0x5c, 0x22, 0xc8, 0x7a, 0x82, 0x5d, 0x43, 0x42, 0x3c, 0xfd, 0xdc,
	0x83, 0x05, 0x18, 0xa5, 0x64, 0xba, 0xe6, 0x2a, 0x48, 0xa0, 0xeb, 0x0a, 0x4e, 0x2d, 0x29, 0x01,
	0x79, 0x44, 0xc8, 0x81, 0x4b, 0x12, 0x59, 0x9f, 0xb3, 0xbf, 0x8f, 0x8f, 0xa3, 0x93, 0x7f, 0x90,
	0x63, 0x88, 0x7f, 0x50, 0xb0, 0x00, 0x93, 0x94, 0x62, 0xd7, 0x5c, 0x05, 0x59, 0x74, 0xcd, 0xce,
	0xf9, 0x39, 0x39, 0x89, 0x49, 0xf9, 0x45, 0x89, 0x25, 0xf9, 0x45, 0xc5, 0x42, 0x1a, 0x5c, 0xbc,
	0xee, 0x8e, 0x21, 0xae, 0xe1, 0x8e, 0x91, 0xf1, 0xfe, 0xe1, 0x7e, 0xae, 0x41, 0x02, 0x5c, 0x52,
	0xa2, 0x5d, 0x73, 0x15, 0x04, 0xc1, 0xba, 0xdc, 0x13, 0x4b, 0x52, 0xcb, 0x13, 0x2b, 0xfd, 0xcb,
	0xf3, 0x52, 0x8b, 0xa4, 0x38, 0x3a, 0x16, 0xcb, 0x31, 0x6c, 0x58, 0x22, 0xc7, 0xe0, 0xe4, 0x7e,
	0xe3, 0xa1, 0x1c, 0x43, 0xc3, 0x23, 0x39, 0xc6, 0x13, 0x8f, 0xe4, 0x18, 0x2f, 0x3c, 0x92, 0x63,
	0x7c, 0xf0, 0x48, 0x8e, 0xf1, 0xc5, 0x23, 0x39, 0x86, 0x0f, 0x8f, 0xe4, 0x18, 0xa3, 0x34, 0x09,
	0x05, 0x79, 0x41, 0x76, 0x3a, 0x88, 0x2e, 0x48, 0x4a, 0x62, 0x03, 0x87, 0x9d, 0x31, 0x20, 0x00,
	0x00, 0xff, 0xff, 0x55, 0x09, 0x7c, 0x22, 0xa6, 0x01, 0x00, 0x00,
}
