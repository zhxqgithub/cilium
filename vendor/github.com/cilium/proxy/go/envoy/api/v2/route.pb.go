// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/api/v2/route.proto

package envoy_api_v2

import (
	fmt "fmt"
	core "github.com/cilium/proxy/go/envoy/api/v2/core"
	route "github.com/cilium/proxy/go/envoy/api/v2/route"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// [#next-free-field: 11]
type RouteConfiguration struct {
	// The name of the route configuration. For example, it might match
	// :ref:`route_config_name
	// <envoy_api_field_config.filter.network.http_connection_manager.v2.Rds.route_config_name>` in
	// :ref:`envoy_api_msg_config.filter.network.http_connection_manager.v2.Rds`.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// An array of virtual hosts that make up the route table.
	VirtualHosts []*route.VirtualHost `protobuf:"bytes,2,rep,name=virtual_hosts,json=virtualHosts,proto3" json:"virtual_hosts,omitempty"`
	// An array of virtual hosts will be dynamically loaded via the VHDS API.
	// Both *virtual_hosts* and *vhds* fields will be used when present. *virtual_hosts* can be used
	// for a base routing table or for infrequently changing virtual hosts. *vhds* is used for
	// on-demand discovery of virtual hosts. The contents of these two fields will be merged to
	// generate a routing table for a given RouteConfiguration, with *vhds* derived configuration
	// taking precedence.
	Vhds *Vhds `protobuf:"bytes,9,opt,name=vhds,proto3" json:"vhds,omitempty"`
	// Optionally specifies a list of HTTP headers that the connection manager
	// will consider to be internal only. If they are found on external requests they will be cleaned
	// prior to filter invocation. See :ref:`config_http_conn_man_headers_x-envoy-internal` for more
	// information.
	InternalOnlyHeaders []string `protobuf:"bytes,3,rep,name=internal_only_headers,json=internalOnlyHeaders,proto3" json:"internal_only_headers,omitempty"`
	// Specifies a list of HTTP headers that should be added to each response that
	// the connection manager encodes. Headers specified at this level are applied
	// after headers from any enclosed :ref:`envoy_api_msg_route.VirtualHost` or
	// :ref:`envoy_api_msg_route.RouteAction`. For more information, including details on
	// header value syntax, see the documentation on :ref:`custom request headers
	// <config_http_conn_man_headers_custom_request_headers>`.
	ResponseHeadersToAdd []*core.HeaderValueOption `protobuf:"bytes,4,rep,name=response_headers_to_add,json=responseHeadersToAdd,proto3" json:"response_headers_to_add,omitempty"`
	// Specifies a list of HTTP headers that should be removed from each response
	// that the connection manager encodes.
	ResponseHeadersToRemove []string `protobuf:"bytes,5,rep,name=response_headers_to_remove,json=responseHeadersToRemove,proto3" json:"response_headers_to_remove,omitempty"`
	// Specifies a list of HTTP headers that should be added to each request
	// routed by the HTTP connection manager. Headers specified at this level are
	// applied after headers from any enclosed :ref:`envoy_api_msg_route.VirtualHost` or
	// :ref:`envoy_api_msg_route.RouteAction`. For more information, including details on
	// header value syntax, see the documentation on :ref:`custom request headers
	// <config_http_conn_man_headers_custom_request_headers>`.
	RequestHeadersToAdd []*core.HeaderValueOption `protobuf:"bytes,6,rep,name=request_headers_to_add,json=requestHeadersToAdd,proto3" json:"request_headers_to_add,omitempty"`
	// Specifies a list of HTTP headers that should be removed from each request
	// routed by the HTTP connection manager.
	RequestHeadersToRemove []string `protobuf:"bytes,8,rep,name=request_headers_to_remove,json=requestHeadersToRemove,proto3" json:"request_headers_to_remove,omitempty"`
	// By default, headers that should be added/removed are evaluated from most to least specific:
	//
	// * route level
	// * virtual host level
	// * connection manager level
	//
	// To allow setting overrides at the route or virtual host level, this order can be reversed
	// by setting this option to true. Defaults to false.
	//
	// [#next-major-version: In the v3 API, this will default to true.]
	MostSpecificHeaderMutationsWins bool `protobuf:"varint,10,opt,name=most_specific_header_mutations_wins,json=mostSpecificHeaderMutationsWins,proto3" json:"most_specific_header_mutations_wins,omitempty"`
	// An optional boolean that specifies whether the clusters that the route
	// table refers to will be validated by the cluster manager. If set to true
	// and a route refers to a non-existent cluster, the route table will not
	// load. If set to false and a route refers to a non-existent cluster, the
	// route table will load and the router filter will return a 404 if the route
	// is selected at runtime. This setting defaults to true if the route table
	// is statically defined via the :ref:`route_config
	// <envoy_api_field_config.filter.network.http_connection_manager.v2.HttpConnectionManager.route_config>`
	// option. This setting default to false if the route table is loaded dynamically via the
	// :ref:`rds
	// <envoy_api_field_config.filter.network.http_connection_manager.v2.HttpConnectionManager.rds>`
	// option. Users may wish to override the default behavior in certain cases (for example when
	// using CDS with a static route table).
	ValidateClusters     *wrappers.BoolValue `protobuf:"bytes,7,opt,name=validate_clusters,json=validateClusters,proto3" json:"validate_clusters,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *RouteConfiguration) Reset()         { *m = RouteConfiguration{} }
func (m *RouteConfiguration) String() string { return proto.CompactTextString(m) }
func (*RouteConfiguration) ProtoMessage()    {}
func (*RouteConfiguration) Descriptor() ([]byte, []int) {
	return fileDescriptor_1f33b4742f398551, []int{0}
}

func (m *RouteConfiguration) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RouteConfiguration.Unmarshal(m, b)
}
func (m *RouteConfiguration) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RouteConfiguration.Marshal(b, m, deterministic)
}
func (m *RouteConfiguration) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RouteConfiguration.Merge(m, src)
}
func (m *RouteConfiguration) XXX_Size() int {
	return xxx_messageInfo_RouteConfiguration.Size(m)
}
func (m *RouteConfiguration) XXX_DiscardUnknown() {
	xxx_messageInfo_RouteConfiguration.DiscardUnknown(m)
}

var xxx_messageInfo_RouteConfiguration proto.InternalMessageInfo

func (m *RouteConfiguration) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *RouteConfiguration) GetVirtualHosts() []*route.VirtualHost {
	if m != nil {
		return m.VirtualHosts
	}
	return nil
}

func (m *RouteConfiguration) GetVhds() *Vhds {
	if m != nil {
		return m.Vhds
	}
	return nil
}

func (m *RouteConfiguration) GetInternalOnlyHeaders() []string {
	if m != nil {
		return m.InternalOnlyHeaders
	}
	return nil
}

func (m *RouteConfiguration) GetResponseHeadersToAdd() []*core.HeaderValueOption {
	if m != nil {
		return m.ResponseHeadersToAdd
	}
	return nil
}

func (m *RouteConfiguration) GetResponseHeadersToRemove() []string {
	if m != nil {
		return m.ResponseHeadersToRemove
	}
	return nil
}

func (m *RouteConfiguration) GetRequestHeadersToAdd() []*core.HeaderValueOption {
	if m != nil {
		return m.RequestHeadersToAdd
	}
	return nil
}

func (m *RouteConfiguration) GetRequestHeadersToRemove() []string {
	if m != nil {
		return m.RequestHeadersToRemove
	}
	return nil
}

func (m *RouteConfiguration) GetMostSpecificHeaderMutationsWins() bool {
	if m != nil {
		return m.MostSpecificHeaderMutationsWins
	}
	return false
}

func (m *RouteConfiguration) GetValidateClusters() *wrappers.BoolValue {
	if m != nil {
		return m.ValidateClusters
	}
	return nil
}

type Vhds struct {
	// Configuration source specifier for VHDS.
	ConfigSource         *core.ConfigSource `protobuf:"bytes,1,opt,name=config_source,json=configSource,proto3" json:"config_source,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *Vhds) Reset()         { *m = Vhds{} }
func (m *Vhds) String() string { return proto.CompactTextString(m) }
func (*Vhds) ProtoMessage()    {}
func (*Vhds) Descriptor() ([]byte, []int) {
	return fileDescriptor_1f33b4742f398551, []int{1}
}

func (m *Vhds) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Vhds.Unmarshal(m, b)
}
func (m *Vhds) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Vhds.Marshal(b, m, deterministic)
}
func (m *Vhds) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Vhds.Merge(m, src)
}
func (m *Vhds) XXX_Size() int {
	return xxx_messageInfo_Vhds.Size(m)
}
func (m *Vhds) XXX_DiscardUnknown() {
	xxx_messageInfo_Vhds.DiscardUnknown(m)
}

var xxx_messageInfo_Vhds proto.InternalMessageInfo

func (m *Vhds) GetConfigSource() *core.ConfigSource {
	if m != nil {
		return m.ConfigSource
	}
	return nil
}

func init() {
	proto.RegisterType((*RouteConfiguration)(nil), "envoy.api.v2.RouteConfiguration")
	proto.RegisterType((*Vhds)(nil), "envoy.api.v2.Vhds")
}

func init() { proto.RegisterFile("envoy/api/v2/route.proto", fileDescriptor_1f33b4742f398551) }

var fileDescriptor_1f33b4742f398551 = []byte{
	// 568 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x53, 0x4f, 0x6f, 0xd3, 0x30,
	0x1c, 0x55, 0xd6, 0x3f, 0x6b, 0xbd, 0x4e, 0x2a, 0x1e, 0x5b, 0x43, 0x84, 0x58, 0x35, 0xfe, 0xa8,
	0x5c, 0x12, 0xa9, 0x3b, 0x21, 0x4e, 0x64, 0x48, 0xdb, 0x01, 0xd8, 0x94, 0xa1, 0x72, 0x8c, 0xdc,
	0xc4, 0x6d, 0x2d, 0x25, 0xfe, 0x05, 0xdb, 0xc9, 0xe8, 0x57, 0xe0, 0x82, 0xc4, 0x89, 0x6f, 0xc0,
	0x57, 0xe3, 0xcc, 0x71, 0x07, 0x84, 0x62, 0x27, 0xb0, 0xac, 0x3b, 0xed, 0x52, 0xc5, 0x7d, 0xef,
	0xfd, 0xde, 0xb3, 0xfd, 0x8c, 0x6c, 0xca, 0x0b, 0x58, 0x7b, 0x24, 0x63, 0x5e, 0x31, 0xf5, 0x04,
	0xe4, 0x8a, 0xba, 0x99, 0x00, 0x05, 0x78, 0xa0, 0x11, 0x97, 0x64, 0xcc, 0x2d, 0xa6, 0xce, 0xe3,
	0x06, 0x2f, 0x02, 0x41, 0xbd, 0x39, 0x91, 0x15, 0xd7, 0x79, 0xbe, 0x89, 0x46, 0xc0, 0x17, 0x6c,
	0x19, 0x4a, 0xc8, 0x45, 0x54, 0xd3, 0x5e, 0x6e, 0x9a, 0x99, 0xdf, 0x30, 0x82, 0x34, 0x03, 0x4e,
	0xb9, 0x92, 0x15, 0xf5, 0xc9, 0x12, 0x60, 0x99, 0x50, 0x4f, 0xaf, 0xe6, 0xf9, 0xc2, 0xbb, 0x12,
	0x24, 0xcb, 0xa8, 0xf8, 0x87, 0xe7, 0x71, 0x46, 0x3c, 0xc2, 0x39, 0x28, 0xa2, 0x18, 0x70, 0xe9,
	0xa5, 0x6c, 0x29, 0x48, 0x9d, 0xde, 0x19, 0x15, 0x24, 0x61, 0x31, 0x51, 0xd4, 0xab, 0x3f, 0x0c,
	0x70, 0xf4, 0xb3, 0x83, 0x70, 0x50, 0x7a, 0x9e, 0xe8, 0x80, 0xb9, 0xd0, 0x72, 0x8c, 0x51, 0x9b,
	0x93, 0x94, 0xda, 0xd6, 0xd8, 0x9a, 0xf4, 0x03, 0xfd, 0x8d, 0xdf, 0xa2, 0xdd, 0x82, 0x09, 0x95,
	0x93, 0x24, 0x5c, 0x81, 0x54, 0xd2, 0xde, 0x1a, 0xb7, 0x26, 0x3b, 0xd3, 0x43, 0xf7, 0xe6, 0xc9,
	0xb8, 0xe6, 0xcc, 0x66, 0x86, 0x78, 0x06, 0x52, 0x05, 0x83, 0xe2, 0xff, 0x42, 0xe2, 0x17, 0xa8,
	0x5d, 0xac, 0x62, 0x69, 0xf7, 0xc7, 0xd6, 0x64, 0x67, 0x8a, 0x9b, 0xe2, 0xd9, 0x2a, 0x96, 0x81,
	0xc6, 0xf1, 0x14, 0xed, 0x33, 0xae, 0xa8, 0xe0, 0x24, 0x09, 0x81, 0x27, 0xeb, 0x70, 0x45, 0x49,
	0x4c, 0x85, 0xb4, 0x5b, 0xe3, 0xd6, 0xa4, 0x1f, 0xec, 0xd5, 0xe0, 0x39, 0x4f, 0xd6, 0x67, 0x06,
	0xc2, 0x0b, 0x34, 0x12, 0x54, 0x66, 0xc0, 0x25, 0xad, 0xe9, 0xa1, 0x82, 0x90, 0xc4, 0xb1, 0xdd,
	0xd6, 0x59, 0x9f, 0x35, 0xed, 0xca, 0x9b, 0x71, 0x8d, 0x78, 0x46, 0x92, 0x9c, 0x9e, 0x67, 0xe5,
	0xe6, 0xfd, 0xfe, 0xb5, 0xdf, 0xfd, 0x6e, 0xb5, 0x86, 0xbf, 0xb6, 0x83, 0x87, 0xf5, 0xbc, 0xca,
	0xe2, 0x23, 0xbc, 0x89, 0x63, 0xfc, 0x1a, 0x39, 0x77, 0xf9, 0x08, 0x9a, 0x42, 0x41, 0xed, 0x8e,
	0x0e, 0x38, 0xda, 0x50, 0x06, 0x1a, 0xc6, 0x31, 0x3a, 0x10, 0xf4, 0x73, 0x4e, 0xa5, 0xba, 0x9d,
	0xb1, 0x7b, 0xbf, 0x8c, 0x7b, 0xd5, 0xb8, 0x46, 0xc4, 0x57, 0xe8, 0xd1, 0x1d, 0x2e, 0x55, 0xc2,
	0x9e, 0x4e, 0x78, 0x70, 0x5b, 0x57, 0x05, 0x7c, 0x87, 0x9e, 0xa6, 0x20, 0x55, 0x28, 0x33, 0x1a,
	0xb1, 0x05, 0x8b, 0xaa, 0x01, 0x61, 0x9a, 0x57, 0xdd, 0x0a, 0xaf, 0x18, 0x97, 0x36, 0x1a, 0x5b,
	0x93, 0x5e, 0x70, 0x58, 0x52, 0x2f, 0x2b, 0xa6, 0x99, 0xf4, 0xbe, 0xe6, 0x7d, 0x62, 0x5c, 0xe2,
	0x53, 0xf4, 0xa0, 0xae, 0x5c, 0x18, 0x25, 0xb9, 0x54, 0xe5, 0x1d, 0x6e, 0xeb, 0xcb, 0x77, 0x5c,
	0xd3, 0x6a, 0xb7, 0x6e, 0xb5, 0xeb, 0x03, 0x24, 0x7a, 0x97, 0xc1, 0xb0, 0x16, 0x9d, 0x54, 0x9a,
	0xa3, 0x19, 0x6a, 0x97, 0xf5, 0xc0, 0x1f, 0xd0, 0x6e, 0xe3, 0x31, 0xe9, 0x8e, 0x6e, 0xd4, 0x50,
	0x1f, 0x9b, 0xe9, 0xf4, 0xa5, 0xa6, 0xf9, 0xbd, 0x6b, 0xbf, 0xf3, 0xd5, 0xda, 0x1a, 0x5a, 0xc1,
	0x20, 0xba, 0xf9, 0xff, 0xe9, 0xef, 0x1f, 0x7f, 0xbe, 0x75, 0x46, 0x78, 0xdf, 0xe8, 0x0d, 0x56,
	0xd5, 0xb8, 0x38, 0x46, 0x0e, 0x03, 0x33, 0x39, 0x13, 0xf0, 0x65, 0xdd, 0x30, 0xf1, 0x91, 0x7e,
	0x39, 0x17, 0x65, 0xfa, 0x0b, 0x6b, 0xde, 0xd5, 0xdb, 0x38, 0xfe, 0x1b, 0x00, 0x00, 0xff, 0xff,
	0xe4, 0xa4, 0xd0, 0x95, 0x44, 0x04, 0x00, 0x00,
}
