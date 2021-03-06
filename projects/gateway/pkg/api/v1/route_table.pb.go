// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/gateway/api/v1/route_table.proto

package v1

import (
	bytes "bytes"
	fmt "fmt"
	math "math"

	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/solo-io/protoc-gen-ext/extproto"
	core "github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

//
//
// The **RouteTable** is a child Routing object for the Gloo Gateway.
//
// A **RouteTable** gets built into the complete routing configuration
// when it is referenced by a `delegateAction`, either
// in a parent VirtualService or another RouteTable.
//
// Routes specified in a RouteTable must have their paths start with the prefix provided in the
// parent's matcher.
//
// For example, the following configuration:
//
//
// ```
// virtualService: mydomain.com
// match: /a
// delegate: a-routes
// ---
// routeTable: a-routes
// match: /1
//
// ```
//
// would *not be valid*, while
//
// ```
// virtualService: mydomain.com
// match: /a
// delegate: a-routes
// ---
// routeTable: a-routes
// match: /a/1
//
// ```
//
// *would* be valid.
//
//
// A complete configuration might look as follows:
//
// ```yaml
// apiVersion: gateway.solo.io/v1
// kind: VirtualService
// metadata:
//   name: 'any'
//   namespace: 'any'
// spec:
//   virtualHost:
//     domains:
//     - 'any.com'
//     routes:
//     - matcher:
//         prefix: '/a' # delegate ownership of routes for `any.com/a`
//       delegateAction:
//         name: 'a-routes'
//         namespace: 'a'
//     - matcher:
//         prefix: '/b' # delegate ownership of routes for `any.com/b`
//       delegateAction:
//         name: 'b-routes'
//         namespace: 'b'
// ```
//
// * A root-level **VirtualService** which delegates routing to to the `a-routes` and `b-routes` **RouteTables**.
// * Routes with `delegateActions` can only use a `prefix` matcher.
//
// ```yaml
// apiVersion: gateway.solo.io/v1
// kind: RouteTable
// metadata:
//   name: 'a-routes'
//   namespace: 'a'
// spec:
//   routes:
//     - matcher:
//         # the path matchers in this RouteTable must begin with the prefix `/a/`
//         prefix: '/a/1'
//       routeAction:
//         single:
//           upstream:
//             name: 'foo-upstream'
//
//     - matcher:
//         prefix: '/a/2'
//       routeAction:
//         single:
//           upstream:
//             name: 'bar-upstream'
// ```
//
// * A **RouteTable** which defines two routes.
//
// ```yaml
// apiVersion: gateway.solo.io/v1
// kind: RouteTable
// metadata:
//   name: 'b-routes'
//   namespace: 'b'
// spec:
//   routes:
//     - matcher:
//         # the path matchers in this RouteTable must begin with the prefix `/b/`
//         regex: '/b/3'
//       routeAction:
//         single:
//           upstream:
//             name: 'bar-upstream'
//     - matcher:
//         prefix: '/b/c/'
//       # routes in the RouteTable can perform any action, including a delegateAction
//       delegateAction:
//         name: 'c-routes'
//         namespace: 'c'
//
// ```
//
// * A **RouteTable** which both *defines a route* and *delegates to* another **RouteTable**.
//
//
// ```yaml
// apiVersion: gateway.solo.io/v1
// kind: RouteTable
// metadata:
//   name: 'c-routes'
//   namespace: 'c'
// spec:
//   routes:
//     - matcher:
//         exact: '/b/c/4'
//       routeAction:
//         single:
//           upstream:
//             name: 'qux-upstream'
// ```
//
// * A RouteTable which is a child of another route table.
//
//
// Would produce the following route config for `mydomain.com`:
//
// ```
// /a/1 -> foo-upstream
// /a/2 -> bar-upstream
// /b/3 -> baz-upstream
// /b/c/4 -> qux-upstream
// ```
//
type RouteTable struct {
	// the list of routes for the route table
	Routes []*Route `protobuf:"bytes,1,rep,name=routes,proto3" json:"routes,omitempty"`
	// Status indicates the validation status of this resource.
	// Status is read-only by clients, and set by gloo during validation
	Status core.Status `protobuf:"bytes,6,opt,name=status,proto3" json:"status" testdiff:"ignore"`
	// Metadata contains the object metadata for this resource
	Metadata             core.Metadata `protobuf:"bytes,7,opt,name=metadata,proto3" json:"metadata"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *RouteTable) Reset()         { *m = RouteTable{} }
func (m *RouteTable) String() string { return proto.CompactTextString(m) }
func (*RouteTable) ProtoMessage()    {}
func (*RouteTable) Descriptor() ([]byte, []int) {
	return fileDescriptor_4d1ea5a66e7f9a13, []int{0}
}
func (m *RouteTable) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RouteTable.Unmarshal(m, b)
}
func (m *RouteTable) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RouteTable.Marshal(b, m, deterministic)
}
func (m *RouteTable) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RouteTable.Merge(m, src)
}
func (m *RouteTable) XXX_Size() int {
	return xxx_messageInfo_RouteTable.Size(m)
}
func (m *RouteTable) XXX_DiscardUnknown() {
	xxx_messageInfo_RouteTable.DiscardUnknown(m)
}

var xxx_messageInfo_RouteTable proto.InternalMessageInfo

func (m *RouteTable) GetRoutes() []*Route {
	if m != nil {
		return m.Routes
	}
	return nil
}

func (m *RouteTable) GetStatus() core.Status {
	if m != nil {
		return m.Status
	}
	return core.Status{}
}

func (m *RouteTable) GetMetadata() core.Metadata {
	if m != nil {
		return m.Metadata
	}
	return core.Metadata{}
}

func init() {
	proto.RegisterType((*RouteTable)(nil), "gateway.solo.io.RouteTable")
}

func init() {
	proto.RegisterFile("github.com/solo-io/gloo/projects/gateway/api/v1/route_table.proto", fileDescriptor_4d1ea5a66e7f9a13)
}

var fileDescriptor_4d1ea5a66e7f9a13 = []byte{
	// 337 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0xc1, 0x4e, 0xfa, 0x40,
	0x10, 0xc6, 0xff, 0xfd, 0xdb, 0x54, 0xb3, 0x18, 0x8d, 0x0d, 0x21, 0x84, 0x28, 0x10, 0x4e, 0x5c,
	0xdc, 0x8d, 0x70, 0x31, 0x24, 0x1e, 0xe4, 0x6a, 0xbc, 0x54, 0x4f, 0x5e, 0xc8, 0x52, 0x96, 0x75,
	0xa5, 0x38, 0xcd, 0xee, 0x14, 0xf1, 0xca, 0xd3, 0xf8, 0x08, 0x3e, 0x82, 0x4f, 0x81, 0x89, 0x6f,
	0x80, 0x89, 0x77, 0xd3, 0x65, 0x21, 0x4a, 0xa2, 0xf1, 0xd6, 0xce, 0xf7, 0xcd, 0x37, 0xfb, 0x9b,
	0x21, 0xe7, 0x52, 0xe1, 0x6d, 0xd6, 0xa7, 0x31, 0x8c, 0x99, 0x81, 0x04, 0x8e, 0x15, 0x30, 0x99,
	0x00, 0xb0, 0x54, 0xc3, 0x9d, 0x88, 0xd1, 0x30, 0xc9, 0x51, 0x3c, 0xf0, 0x47, 0xc6, 0x53, 0xc5,
	0x26, 0x27, 0x4c, 0x43, 0x86, 0xa2, 0x87, 0xbc, 0x9f, 0x08, 0x9a, 0x6a, 0x40, 0x08, 0xf7, 0x9d,
	0x83, 0xe6, 0xfd, 0x54, 0x41, 0xa5, 0x28, 0x41, 0x82, 0xd5, 0x58, 0xfe, 0xb5, 0xb4, 0x55, 0x42,
	0x31, 0xc5, 0x65, 0x51, 0x4c, 0xd1, 0xd5, 0xaa, 0x76, 0xe4, 0x48, 0xe1, 0x2a, 0x7d, 0x2c, 0x90,
	0x0f, 0x38, 0x72, 0xa7, 0x1f, 0x6e, 0xea, 0x06, 0x39, 0x66, 0xe6, 0xa7, 0xee, 0xd5, 0xbf, 0xd3,
	0x5b, 0xbf, 0x82, 0x4c, 0x94, 0xc6, 0x8c, 0x27, 0x3d, 0x23, 0xf4, 0x44, 0xc5, 0x0e, 0xa6, 0xf1,
	0xea, 0x11, 0x12, 0xe5, 0x88, 0xd7, 0x39, 0x61, 0x48, 0x49, 0x60, 0x81, 0x4d, 0xd9, 0xab, 0x6f,
	0x35, 0x0b, 0xad, 0x12, 0xdd, 0x80, 0xa5, 0xd6, 0x1c, 0x39, 0x57, 0x78, 0x41, 0x82, 0xe5, 0x13,
	0xcb, 0x41, 0xdd, 0x6b, 0x16, 0x5a, 0x45, 0x1a, 0x83, 0x16, 0x6b, 0xf3, 0x95, 0xd5, 0xba, 0x47,
	0xcf, 0x1f, 0xbe, 0xf7, 0x32, 0xaf, 0xfd, 0x7b, 0x9f, 0xd7, 0x0e, 0x50, 0x18, 0x1c, 0xa8, 0xe1,
	0xb0, 0xd3, 0x50, 0xf2, 0x1e, 0xb4, 0x68, 0x44, 0x2e, 0x22, 0x3c, 0x25, 0x3b, 0xab, 0x7d, 0x94,
	0xb7, 0x6d, 0x5c, 0xe9, 0x7b, 0xdc, 0xa5, 0x53, 0xbb, 0x7e, 0x1e, 0x16, 0xad, 0xdd, 0x9d, 0xca,
	0x6c, 0xe1, 0xfb, 0xe4, 0xbf, 0xc6, 0xd9, 0xc2, 0xdf, 0x0b, 0x77, 0xbf, 0xdc, 0xcc, 0x74, 0xcf,
	0xf2, 0xe1, 0x4f, 0x6f, 0x55, 0xef, 0xa6, 0xfd, 0xe7, 0xdb, 0xa7, 0x23, 0xe9, 0xd6, 0xd6, 0x0f,
	0xec, 0x9e, 0xda, 0x9f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x64, 0x20, 0xb5, 0x0d, 0x39, 0x02, 0x00,
	0x00,
}

func (this *RouteTable) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*RouteTable)
	if !ok {
		that2, ok := that.(RouteTable)
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
	if len(this.Routes) != len(that1.Routes) {
		return false
	}
	for i := range this.Routes {
		if !this.Routes[i].Equal(that1.Routes[i]) {
			return false
		}
	}
	if !this.Status.Equal(&that1.Status) {
		return false
	}
	if !this.Metadata.Equal(&that1.Metadata) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
