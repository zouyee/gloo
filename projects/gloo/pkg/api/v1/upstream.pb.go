// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: upstream.proto

package v1 // import "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"
import core "github.com/solo-io/solo-kit/pkg/api/v1/resources/core"

import bytes "bytes"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

//
// @solo-kit:resource.short_name=us
// @solo-kit:resource.plural_name=upstreams
// @solo-kit:resource.resource_groups=api.gloo.solo.io,discovery.gloo.solo.io,translator.supergloo.solo.io
//
// Upstreams represent destination for routing HTTP requests. Upstreams can be compared to
// [clusters](https://www.envoyproxy.io/docs/envoy/latest/api-v1/cluster_manager/cluster.html?highlight=cluster) in Envoy terminology.
// Each upstream in Gloo has a type. Supported types include `static`, `kubernetes`, `aws`, `consul`, and more.
// Each upstream type is handled by a corresponding Gloo plugin.
type Upstream struct {
	// Type-specific configuration. Examples include static, kubernetes, and aws.
	// The type-specific config for the upstream is called a spec.
	UpstreamSpec *UpstreamSpec `protobuf:"bytes,2,opt,name=upstream_spec,json=upstreamSpec" json:"upstream_spec,omitempty"`
	// Status indicates the validation status of the resource. Status is read-only by clients, and set by gloo during validation
	Status core.Status `protobuf:"bytes,6,opt,name=status" json:"status" testdiff:"ignore"`
	// Metadata contains the object metadata for this resource
	Metadata core.Metadata `protobuf:"bytes,7,opt,name=metadata" json:"metadata"`
	// Upstreams and their configuration can be automatically by Gloo Discovery
	// if this upstream is created or modified by Discovery, metadata about the operation will be placed here.
	DiscoveryMetadata    *DiscoveryMetadata `protobuf:"bytes,8,opt,name=discovery_metadata,json=discoveryMetadata" json:"discovery_metadata,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *Upstream) Reset()         { *m = Upstream{} }
func (m *Upstream) String() string { return proto.CompactTextString(m) }
func (*Upstream) ProtoMessage()    {}
func (*Upstream) Descriptor() ([]byte, []int) {
	return fileDescriptor_upstream_8296731ebe0b6f5e, []int{0}
}
func (m *Upstream) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Upstream.Unmarshal(m, b)
}
func (m *Upstream) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Upstream.Marshal(b, m, deterministic)
}
func (dst *Upstream) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Upstream.Merge(dst, src)
}
func (m *Upstream) XXX_Size() int {
	return xxx_messageInfo_Upstream.Size(m)
}
func (m *Upstream) XXX_DiscardUnknown() {
	xxx_messageInfo_Upstream.DiscardUnknown(m)
}

var xxx_messageInfo_Upstream proto.InternalMessageInfo

func (m *Upstream) GetUpstreamSpec() *UpstreamSpec {
	if m != nil {
		return m.UpstreamSpec
	}
	return nil
}

func (m *Upstream) GetStatus() core.Status {
	if m != nil {
		return m.Status
	}
	return core.Status{}
}

func (m *Upstream) GetMetadata() core.Metadata {
	if m != nil {
		return m.Metadata
	}
	return core.Metadata{}
}

func (m *Upstream) GetDiscoveryMetadata() *DiscoveryMetadata {
	if m != nil {
		return m.DiscoveryMetadata
	}
	return nil
}

// created by discovery services
type DiscoveryMetadata struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DiscoveryMetadata) Reset()         { *m = DiscoveryMetadata{} }
func (m *DiscoveryMetadata) String() string { return proto.CompactTextString(m) }
func (*DiscoveryMetadata) ProtoMessage()    {}
func (*DiscoveryMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_upstream_8296731ebe0b6f5e, []int{1}
}
func (m *DiscoveryMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DiscoveryMetadata.Unmarshal(m, b)
}
func (m *DiscoveryMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DiscoveryMetadata.Marshal(b, m, deterministic)
}
func (dst *DiscoveryMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DiscoveryMetadata.Merge(dst, src)
}
func (m *DiscoveryMetadata) XXX_Size() int {
	return xxx_messageInfo_DiscoveryMetadata.Size(m)
}
func (m *DiscoveryMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_DiscoveryMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_DiscoveryMetadata proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Upstream)(nil), "gloo.solo.io.Upstream")
	proto.RegisterType((*DiscoveryMetadata)(nil), "gloo.solo.io.DiscoveryMetadata")
}
func (this *Upstream) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Upstream)
	if !ok {
		that2, ok := that.(Upstream)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.UpstreamSpec.Equal(that1.UpstreamSpec) {
		return false
	}
	if !this.Status.Equal(&that1.Status) {
		return false
	}
	if !this.Metadata.Equal(&that1.Metadata) {
		return false
	}
	if !this.DiscoveryMetadata.Equal(that1.DiscoveryMetadata) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *DiscoveryMetadata) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*DiscoveryMetadata)
	if !ok {
		that2, ok := that.(DiscoveryMetadata)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}

func init() { proto.RegisterFile("upstream.proto", fileDescriptor_upstream_8296731ebe0b6f5e) }

var fileDescriptor_upstream_8296731ebe0b6f5e = []byte{
	// 323 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x90, 0xcf, 0x4a, 0x33, 0x31,
	0x14, 0xc5, 0xbf, 0x96, 0x8f, 0x5a, 0x62, 0x15, 0x1a, 0x8b, 0xd4, 0x2e, 0xac, 0xcc, 0xca, 0x8d,
	0x89, 0x55, 0x10, 0xe9, 0x46, 0x28, 0x82, 0x2b, 0x5d, 0x4c, 0x71, 0xe3, 0xa6, 0xa4, 0x99, 0x34,
	0xc6, 0xfe, 0xb9, 0x21, 0xc9, 0x14, 0x7c, 0x19, 0xd7, 0x3e, 0x8a, 0x4f, 0xd1, 0x85, 0x8f, 0xe0,
	0x13, 0xc8, 0x64, 0x32, 0x43, 0xab, 0x2e, 0xba, 0x9a, 0xb9, 0xb9, 0xe7, 0x77, 0xb8, 0xe7, 0xa0,
	0xfd, 0x54, 0x5b, 0x67, 0x04, 0x9b, 0x13, 0x6d, 0xc0, 0x01, 0x6e, 0xc8, 0x19, 0x00, 0xb1, 0x30,
	0x03, 0xa2, 0xa0, 0xd3, 0x92, 0x20, 0xc1, 0x2f, 0x68, 0xf6, 0x97, 0x6b, 0x3a, 0x3d, 0xa9, 0xdc,
	0x73, 0x3a, 0x26, 0x1c, 0xe6, 0x34, 0x53, 0x9e, 0x29, 0xc8, 0xbf, 0x53, 0xe5, 0x28, 0xd3, 0x8a,
	0x2e, 0x7b, 0x74, 0x2e, 0x1c, 0x4b, 0x98, 0x63, 0x01, 0xa1, 0x5b, 0x20, 0xd6, 0x31, 0x97, 0xda,
	0x00, 0xf4, 0xff, 0x00, 0xb2, 0xd3, 0xa8, 0x36, 0xf0, 0x22, 0xb8, 0xb3, 0xf9, 0x14, 0x50, 0x3d,
	0x4b, 0xa5, 0x5a, 0x04, 0x36, 0x7a, 0xab, 0xa2, 0xfa, 0x63, 0x88, 0x85, 0x6f, 0xd0, 0x5e, 0x11,
	0x71, 0x64, 0xb5, 0xe0, 0xed, 0xea, 0x49, 0xe5, 0x74, 0xf7, 0xa2, 0x43, 0xd6, 0x83, 0x92, 0x42,
	0x3e, 0xd4, 0x82, 0xc7, 0x8d, 0x74, 0x6d, 0xc2, 0x77, 0xa8, 0x96, 0x5f, 0xd6, 0xae, 0x79, 0xb2,
	0x45, 0x38, 0x18, 0x51, 0x92, 0x43, 0xbf, 0x1b, 0x1c, 0x7d, 0xac, 0xba, 0xff, 0xbe, 0x56, 0xdd,
	0xa6, 0x13, 0xd6, 0x25, 0x6a, 0x32, 0xe9, 0x47, 0x4a, 0x2e, 0xc0, 0x88, 0x28, 0x0e, 0x38, 0xbe,
	0x46, 0xf5, 0xa2, 0x95, 0xf6, 0x8e, 0xb7, 0x3a, 0xdc, 0xb4, 0xba, 0x0f, 0xdb, 0xc1, 0xff, 0xcc,
	0x2c, 0x2e, 0xd5, 0xf8, 0x01, 0xe1, 0x44, 0x59, 0x0e, 0x4b, 0x61, 0x5e, 0x47, 0xa5, 0x47, 0xdd,
	0x7b, 0x74, 0x37, 0x83, 0xdc, 0x16, 0xba, 0xc2, 0x2c, 0x6e, 0x26, 0x3f, 0x9f, 0xa2, 0x03, 0xd4,
	0xfc, 0xa5, 0x1b, 0x5c, 0xbd, 0x7f, 0x1e, 0x57, 0x9e, 0xce, 0xb7, 0xeb, 0x5d, 0x4f, 0x65, 0xe8,
	0x7e, 0x5c, 0xf3, 0xa5, 0x5f, 0x7e, 0x07, 0x00, 0x00, 0xff, 0xff, 0xbc, 0x86, 0xf2, 0x46, 0x4a,
	0x02, 0x00, 0x00,
}