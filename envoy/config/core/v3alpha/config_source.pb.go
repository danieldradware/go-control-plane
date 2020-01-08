// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/core/v3alpha/config_source.proto

package envoy_config_core_v3alpha

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
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

type ApiVersion int32

const (
	ApiVersion_AUTO    ApiVersion = 0
	ApiVersion_V2      ApiVersion = 1
	ApiVersion_V3ALPHA ApiVersion = 2
)

var ApiVersion_name = map[int32]string{
	0: "AUTO",
	1: "V2",
	2: "V3ALPHA",
}

var ApiVersion_value = map[string]int32{
	"AUTO":    0,
	"V2":      1,
	"V3ALPHA": 2,
}

func (x ApiVersion) String() string {
	return proto.EnumName(ApiVersion_name, int32(x))
}

func (ApiVersion) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_fb345dcf1abd8590, []int{0}
}

type ApiConfigSource_ApiType int32

const (
	ApiConfigSource_hidden_envoy_deprecated_UNSUPPORTED_REST_LEGACY ApiConfigSource_ApiType = 0 // Deprecated: Do not use.
	ApiConfigSource_REST                                            ApiConfigSource_ApiType = 1
	ApiConfigSource_GRPC                                            ApiConfigSource_ApiType = 2
	ApiConfigSource_DELTA_GRPC                                      ApiConfigSource_ApiType = 3
)

var ApiConfigSource_ApiType_name = map[int32]string{
	0: "hidden_envoy_deprecated_UNSUPPORTED_REST_LEGACY",
	1: "REST",
	2: "GRPC",
	3: "DELTA_GRPC",
}

var ApiConfigSource_ApiType_value = map[string]int32{
	"hidden_envoy_deprecated_UNSUPPORTED_REST_LEGACY": 0,
	"REST":       1,
	"GRPC":       2,
	"DELTA_GRPC": 3,
}

func (x ApiConfigSource_ApiType) String() string {
	return proto.EnumName(ApiConfigSource_ApiType_name, int32(x))
}

func (ApiConfigSource_ApiType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_fb345dcf1abd8590, []int{0, 0}
}

type ApiConfigSource struct {
	ApiType                   ApiConfigSource_ApiType `protobuf:"varint,1,opt,name=api_type,json=apiType,proto3,enum=envoy.config.core.v3alpha.ApiConfigSource_ApiType" json:"api_type,omitempty"`
	TransportApiVersion       ApiVersion              `protobuf:"varint,8,opt,name=transport_api_version,json=transportApiVersion,proto3,enum=envoy.config.core.v3alpha.ApiVersion" json:"transport_api_version,omitempty"`
	ClusterNames              []string                `protobuf:"bytes,2,rep,name=cluster_names,json=clusterNames,proto3" json:"cluster_names,omitempty"`
	GrpcServices              []*GrpcService          `protobuf:"bytes,4,rep,name=grpc_services,json=grpcServices,proto3" json:"grpc_services,omitempty"`
	RefreshDelay              *duration.Duration      `protobuf:"bytes,3,opt,name=refresh_delay,json=refreshDelay,proto3" json:"refresh_delay,omitempty"`
	RequestTimeout            *duration.Duration      `protobuf:"bytes,5,opt,name=request_timeout,json=requestTimeout,proto3" json:"request_timeout,omitempty"`
	RateLimitSettings         *RateLimitSettings      `protobuf:"bytes,6,opt,name=rate_limit_settings,json=rateLimitSettings,proto3" json:"rate_limit_settings,omitempty"`
	SetNodeOnFirstMessageOnly bool                    `protobuf:"varint,7,opt,name=set_node_on_first_message_only,json=setNodeOnFirstMessageOnly,proto3" json:"set_node_on_first_message_only,omitempty"`
	XXX_NoUnkeyedLiteral      struct{}                `json:"-"`
	XXX_unrecognized          []byte                  `json:"-"`
	XXX_sizecache             int32                   `json:"-"`
}

func (m *ApiConfigSource) Reset()         { *m = ApiConfigSource{} }
func (m *ApiConfigSource) String() string { return proto.CompactTextString(m) }
func (*ApiConfigSource) ProtoMessage()    {}
func (*ApiConfigSource) Descriptor() ([]byte, []int) {
	return fileDescriptor_fb345dcf1abd8590, []int{0}
}

func (m *ApiConfigSource) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ApiConfigSource.Unmarshal(m, b)
}
func (m *ApiConfigSource) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ApiConfigSource.Marshal(b, m, deterministic)
}
func (m *ApiConfigSource) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ApiConfigSource.Merge(m, src)
}
func (m *ApiConfigSource) XXX_Size() int {
	return xxx_messageInfo_ApiConfigSource.Size(m)
}
func (m *ApiConfigSource) XXX_DiscardUnknown() {
	xxx_messageInfo_ApiConfigSource.DiscardUnknown(m)
}

var xxx_messageInfo_ApiConfigSource proto.InternalMessageInfo

func (m *ApiConfigSource) GetApiType() ApiConfigSource_ApiType {
	if m != nil {
		return m.ApiType
	}
	return ApiConfigSource_hidden_envoy_deprecated_UNSUPPORTED_REST_LEGACY
}

func (m *ApiConfigSource) GetTransportApiVersion() ApiVersion {
	if m != nil {
		return m.TransportApiVersion
	}
	return ApiVersion_AUTO
}

func (m *ApiConfigSource) GetClusterNames() []string {
	if m != nil {
		return m.ClusterNames
	}
	return nil
}

func (m *ApiConfigSource) GetGrpcServices() []*GrpcService {
	if m != nil {
		return m.GrpcServices
	}
	return nil
}

func (m *ApiConfigSource) GetRefreshDelay() *duration.Duration {
	if m != nil {
		return m.RefreshDelay
	}
	return nil
}

func (m *ApiConfigSource) GetRequestTimeout() *duration.Duration {
	if m != nil {
		return m.RequestTimeout
	}
	return nil
}

func (m *ApiConfigSource) GetRateLimitSettings() *RateLimitSettings {
	if m != nil {
		return m.RateLimitSettings
	}
	return nil
}

func (m *ApiConfigSource) GetSetNodeOnFirstMessageOnly() bool {
	if m != nil {
		return m.SetNodeOnFirstMessageOnly
	}
	return false
}

type AggregatedConfigSource struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AggregatedConfigSource) Reset()         { *m = AggregatedConfigSource{} }
func (m *AggregatedConfigSource) String() string { return proto.CompactTextString(m) }
func (*AggregatedConfigSource) ProtoMessage()    {}
func (*AggregatedConfigSource) Descriptor() ([]byte, []int) {
	return fileDescriptor_fb345dcf1abd8590, []int{1}
}

func (m *AggregatedConfigSource) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AggregatedConfigSource.Unmarshal(m, b)
}
func (m *AggregatedConfigSource) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AggregatedConfigSource.Marshal(b, m, deterministic)
}
func (m *AggregatedConfigSource) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AggregatedConfigSource.Merge(m, src)
}
func (m *AggregatedConfigSource) XXX_Size() int {
	return xxx_messageInfo_AggregatedConfigSource.Size(m)
}
func (m *AggregatedConfigSource) XXX_DiscardUnknown() {
	xxx_messageInfo_AggregatedConfigSource.DiscardUnknown(m)
}

var xxx_messageInfo_AggregatedConfigSource proto.InternalMessageInfo

type SelfConfigSource struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SelfConfigSource) Reset()         { *m = SelfConfigSource{} }
func (m *SelfConfigSource) String() string { return proto.CompactTextString(m) }
func (*SelfConfigSource) ProtoMessage()    {}
func (*SelfConfigSource) Descriptor() ([]byte, []int) {
	return fileDescriptor_fb345dcf1abd8590, []int{2}
}

func (m *SelfConfigSource) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SelfConfigSource.Unmarshal(m, b)
}
func (m *SelfConfigSource) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SelfConfigSource.Marshal(b, m, deterministic)
}
func (m *SelfConfigSource) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SelfConfigSource.Merge(m, src)
}
func (m *SelfConfigSource) XXX_Size() int {
	return xxx_messageInfo_SelfConfigSource.Size(m)
}
func (m *SelfConfigSource) XXX_DiscardUnknown() {
	xxx_messageInfo_SelfConfigSource.DiscardUnknown(m)
}

var xxx_messageInfo_SelfConfigSource proto.InternalMessageInfo

type RateLimitSettings struct {
	MaxTokens            *wrappers.UInt32Value `protobuf:"bytes,1,opt,name=max_tokens,json=maxTokens,proto3" json:"max_tokens,omitempty"`
	FillRate             *wrappers.DoubleValue `protobuf:"bytes,2,opt,name=fill_rate,json=fillRate,proto3" json:"fill_rate,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *RateLimitSettings) Reset()         { *m = RateLimitSettings{} }
func (m *RateLimitSettings) String() string { return proto.CompactTextString(m) }
func (*RateLimitSettings) ProtoMessage()    {}
func (*RateLimitSettings) Descriptor() ([]byte, []int) {
	return fileDescriptor_fb345dcf1abd8590, []int{3}
}

func (m *RateLimitSettings) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RateLimitSettings.Unmarshal(m, b)
}
func (m *RateLimitSettings) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RateLimitSettings.Marshal(b, m, deterministic)
}
func (m *RateLimitSettings) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RateLimitSettings.Merge(m, src)
}
func (m *RateLimitSettings) XXX_Size() int {
	return xxx_messageInfo_RateLimitSettings.Size(m)
}
func (m *RateLimitSettings) XXX_DiscardUnknown() {
	xxx_messageInfo_RateLimitSettings.DiscardUnknown(m)
}

var xxx_messageInfo_RateLimitSettings proto.InternalMessageInfo

func (m *RateLimitSettings) GetMaxTokens() *wrappers.UInt32Value {
	if m != nil {
		return m.MaxTokens
	}
	return nil
}

func (m *RateLimitSettings) GetFillRate() *wrappers.DoubleValue {
	if m != nil {
		return m.FillRate
	}
	return nil
}

type ConfigSource struct {
	// Types that are valid to be assigned to ConfigSourceSpecifier:
	//	*ConfigSource_Path
	//	*ConfigSource_ApiConfigSource
	//	*ConfigSource_Ads
	//	*ConfigSource_Self
	ConfigSourceSpecifier isConfigSource_ConfigSourceSpecifier `protobuf_oneof:"config_source_specifier"`
	InitialFetchTimeout   *duration.Duration                   `protobuf:"bytes,4,opt,name=initial_fetch_timeout,json=initialFetchTimeout,proto3" json:"initial_fetch_timeout,omitempty"`
	ResourceApiVersion    ApiVersion                           `protobuf:"varint,6,opt,name=resource_api_version,json=resourceApiVersion,proto3,enum=envoy.config.core.v3alpha.ApiVersion" json:"resource_api_version,omitempty"`
	XXX_NoUnkeyedLiteral  struct{}                             `json:"-"`
	XXX_unrecognized      []byte                               `json:"-"`
	XXX_sizecache         int32                                `json:"-"`
}

func (m *ConfigSource) Reset()         { *m = ConfigSource{} }
func (m *ConfigSource) String() string { return proto.CompactTextString(m) }
func (*ConfigSource) ProtoMessage()    {}
func (*ConfigSource) Descriptor() ([]byte, []int) {
	return fileDescriptor_fb345dcf1abd8590, []int{4}
}

func (m *ConfigSource) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfigSource.Unmarshal(m, b)
}
func (m *ConfigSource) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfigSource.Marshal(b, m, deterministic)
}
func (m *ConfigSource) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfigSource.Merge(m, src)
}
func (m *ConfigSource) XXX_Size() int {
	return xxx_messageInfo_ConfigSource.Size(m)
}
func (m *ConfigSource) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfigSource.DiscardUnknown(m)
}

var xxx_messageInfo_ConfigSource proto.InternalMessageInfo

type isConfigSource_ConfigSourceSpecifier interface {
	isConfigSource_ConfigSourceSpecifier()
}

type ConfigSource_Path struct {
	Path string `protobuf:"bytes,1,opt,name=path,proto3,oneof"`
}

type ConfigSource_ApiConfigSource struct {
	ApiConfigSource *ApiConfigSource `protobuf:"bytes,2,opt,name=api_config_source,json=apiConfigSource,proto3,oneof"`
}

type ConfigSource_Ads struct {
	Ads *AggregatedConfigSource `protobuf:"bytes,3,opt,name=ads,proto3,oneof"`
}

type ConfigSource_Self struct {
	Self *SelfConfigSource `protobuf:"bytes,5,opt,name=self,proto3,oneof"`
}

func (*ConfigSource_Path) isConfigSource_ConfigSourceSpecifier() {}

func (*ConfigSource_ApiConfigSource) isConfigSource_ConfigSourceSpecifier() {}

func (*ConfigSource_Ads) isConfigSource_ConfigSourceSpecifier() {}

func (*ConfigSource_Self) isConfigSource_ConfigSourceSpecifier() {}

func (m *ConfigSource) GetConfigSourceSpecifier() isConfigSource_ConfigSourceSpecifier {
	if m != nil {
		return m.ConfigSourceSpecifier
	}
	return nil
}

func (m *ConfigSource) GetPath() string {
	if x, ok := m.GetConfigSourceSpecifier().(*ConfigSource_Path); ok {
		return x.Path
	}
	return ""
}

func (m *ConfigSource) GetApiConfigSource() *ApiConfigSource {
	if x, ok := m.GetConfigSourceSpecifier().(*ConfigSource_ApiConfigSource); ok {
		return x.ApiConfigSource
	}
	return nil
}

func (m *ConfigSource) GetAds() *AggregatedConfigSource {
	if x, ok := m.GetConfigSourceSpecifier().(*ConfigSource_Ads); ok {
		return x.Ads
	}
	return nil
}

func (m *ConfigSource) GetSelf() *SelfConfigSource {
	if x, ok := m.GetConfigSourceSpecifier().(*ConfigSource_Self); ok {
		return x.Self
	}
	return nil
}

func (m *ConfigSource) GetInitialFetchTimeout() *duration.Duration {
	if m != nil {
		return m.InitialFetchTimeout
	}
	return nil
}

func (m *ConfigSource) GetResourceApiVersion() ApiVersion {
	if m != nil {
		return m.ResourceApiVersion
	}
	return ApiVersion_AUTO
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*ConfigSource) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*ConfigSource_Path)(nil),
		(*ConfigSource_ApiConfigSource)(nil),
		(*ConfigSource_Ads)(nil),
		(*ConfigSource_Self)(nil),
	}
}

func init() {
	proto.RegisterEnum("envoy.config.core.v3alpha.ApiVersion", ApiVersion_name, ApiVersion_value)
	proto.RegisterEnum("envoy.config.core.v3alpha.ApiConfigSource_ApiType", ApiConfigSource_ApiType_name, ApiConfigSource_ApiType_value)
	proto.RegisterType((*ApiConfigSource)(nil), "envoy.config.core.v3alpha.ApiConfigSource")
	proto.RegisterType((*AggregatedConfigSource)(nil), "envoy.config.core.v3alpha.AggregatedConfigSource")
	proto.RegisterType((*SelfConfigSource)(nil), "envoy.config.core.v3alpha.SelfConfigSource")
	proto.RegisterType((*RateLimitSettings)(nil), "envoy.config.core.v3alpha.RateLimitSettings")
	proto.RegisterType((*ConfigSource)(nil), "envoy.config.core.v3alpha.ConfigSource")
}

func init() {
	proto.RegisterFile("envoy/config/core/v3alpha/config_source.proto", fileDescriptor_fb345dcf1abd8590)
}

var fileDescriptor_fb345dcf1abd8590 = []byte{
	// 918 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0xdd, 0x4e, 0x1b, 0x47,
	0x14, 0xf6, 0xda, 0x0e, 0x98, 0xc3, 0x9f, 0x19, 0x92, 0xc6, 0xa0, 0x0a, 0x19, 0x53, 0x5a, 0x97,
	0xa4, 0xb6, 0x6a, 0x2e, 0x2a, 0x51, 0x35, 0x92, 0x0d, 0x04, 0x50, 0x09, 0x58, 0x6b, 0x43, 0x5b,
	0xa9, 0xd2, 0x68, 0xd8, 0x3d, 0xbb, 0x8c, 0xba, 0xde, 0xd9, 0xce, 0x8c, 0x5d, 0x7c, 0xd9, 0x5c,
	0xf5, 0x19, 0xfa, 0x08, 0x7d, 0x8f, 0xaa, 0xaf, 0x54, 0xe5, 0xaa, 0x9a, 0xdd, 0x25, 0x18, 0x1b,
	0x9c, 0x2a, 0x7b, 0xb5, 0x73, 0xce, 0xf7, 0x7d, 0x73, 0xce, 0x7c, 0x33, 0x07, 0xbe, 0xc2, 0x70,
	0x20, 0x86, 0x75, 0x47, 0x84, 0x1e, 0xf7, 0xeb, 0x8e, 0x90, 0x58, 0x1f, 0xec, 0xb2, 0x20, 0xba,
	0x66, 0x69, 0x8c, 0x2a, 0xd1, 0x97, 0x0e, 0xd6, 0x22, 0x29, 0xb4, 0x20, 0x6b, 0x31, 0xbc, 0x96,
	0xa4, 0x6a, 0x06, 0x5e, 0x4b, 0xe1, 0xeb, 0x2f, 0x1f, 0x57, 0xf2, 0x65, 0xe4, 0x50, 0x85, 0x72,
	0xc0, 0x6f, 0x85, 0xd6, 0x37, 0x7c, 0x21, 0xfc, 0x00, 0xeb, 0xf1, 0xea, 0xaa, 0xef, 0xd5, 0xdd,
	0xbe, 0x64, 0x9a, 0x8b, 0xf0, 0xb1, 0xfc, 0x6f, 0x92, 0x45, 0x11, 0x4a, 0x95, 0xe6, 0x37, 0xfb,
	0x6e, 0xc4, 0xea, 0x2c, 0x0c, 0x85, 0x8e, 0x69, 0xaa, 0x3e, 0x40, 0xa9, 0xb8, 0x08, 0x79, 0xe8,
	0xa7, 0x90, 0xe7, 0x03, 0x16, 0x70, 0x97, 0x69, 0xac, 0xdf, 0xfe, 0x24, 0x89, 0xca, 0xdb, 0x19,
	0x58, 0x6e, 0x46, 0x7c, 0x3f, 0xae, 0xb4, 0x13, 0xb7, 0x47, 0x7e, 0x80, 0x02, 0x8b, 0x38, 0xd5,
	0xc3, 0x08, 0x4b, 0x56, 0xd9, 0xaa, 0x2e, 0x35, 0x1a, 0xb5, 0x47, 0x7b, 0xad, 0x8d, 0xb1, 0xcd,
	0xba, 0x3b, 0x8c, 0xb0, 0x55, 0x78, 0xd7, 0x7a, 0xf2, 0xd6, 0xca, 0x16, 0x2d, 0x7b, 0x96, 0x25,
	0x21, 0xe2, 0xc0, 0x33, 0x2d, 0x59, 0xa8, 0x22, 0x21, 0x35, 0x35, 0x5b, 0xa4, 0x75, 0x96, 0x0a,
	0xf1, 0x2e, 0xdb, 0xd3, 0x77, 0xb9, 0x4c, 0xc0, 0x23, 0xc2, 0xab, 0xef, 0xd5, 0xee, 0xd2, 0x64,
	0x0b, 0x16, 0x9d, 0xa0, 0xaf, 0x34, 0x4a, 0x1a, 0xb2, 0x1e, 0xaa, 0x52, 0xb6, 0x9c, 0xab, 0xce,
	0xd9, 0x0b, 0x69, 0xf0, 0xcc, 0xc4, 0xc8, 0xf7, 0xb0, 0x38, 0x6a, 0x84, 0x2a, 0xe5, 0xcb, 0xb9,
	0xea, 0x7c, 0xe3, 0xf3, 0x29, 0x15, 0x1c, 0xc9, 0xc8, 0xe9, 0x24, 0x70, 0x7b, 0xc1, 0xbf, 0x5b,
	0x28, 0xf2, 0x0a, 0x16, 0x25, 0x7a, 0x12, 0xd5, 0x35, 0x75, 0x31, 0x60, 0xc3, 0x52, 0xae, 0x6c,
	0x55, 0xe7, 0x1b, 0x6b, 0xb5, 0xc4, 0xb7, 0xda, 0xad, 0x6f, 0xb5, 0x83, 0xd4, 0x57, 0x7b, 0x21,
	0xc5, 0x1f, 0x18, 0x38, 0x39, 0x85, 0x65, 0x89, 0xbf, 0xf6, 0x51, 0x69, 0xaa, 0x79, 0x0f, 0x45,
	0x5f, 0x97, 0x9e, 0x7c, 0x40, 0x21, 0x3e, 0x84, 0xbf, 0xac, 0xec, 0x4e, 0xc6, 0x5e, 0x4a, 0xb9,
	0xdd, 0x84, 0x4a, 0x7e, 0x86, 0x55, 0xc9, 0x34, 0xd2, 0x80, 0xf7, 0xb8, 0xa6, 0x0a, 0xb5, 0xe6,
	0xa1, 0xaf, 0x4a, 0x33, 0xb1, 0xe2, 0xcb, 0x29, 0x0d, 0xda, 0x4c, 0xe3, 0xa9, 0x21, 0x75, 0x52,
	0x8e, 0xbd, 0x22, 0xc7, 0x43, 0xa4, 0x09, 0x1b, 0x0a, 0x35, 0x0d, 0x85, 0x8b, 0x54, 0x84, 0xd4,
	0xe3, 0x52, 0x69, 0xda, 0x43, 0xa5, 0x98, 0x6f, 0x02, 0xc1, 0xb0, 0x34, 0x5b, 0xb6, 0xaa, 0x05,
	0x7b, 0x4d, 0xa1, 0x3e, 0x13, 0x2e, 0x9e, 0x87, 0xaf, 0x0d, 0xe4, 0x4d, 0x82, 0x38, 0x0f, 0x83,
	0x61, 0xc5, 0x83, 0xd9, 0xf4, 0x8e, 0x90, 0x6f, 0xa0, 0x7e, 0xcd, 0x5d, 0x17, 0x43, 0x1a, 0x97,
	0x45, 0x5d, 0x8c, 0x24, 0x3a, 0x4c, 0xa3, 0x4b, 0x2f, 0xce, 0x3a, 0x17, 0xed, 0xf6, 0xb9, 0xdd,
	0x3d, 0x3c, 0xa0, 0xf6, 0x61, 0xa7, 0x4b, 0x4f, 0x0f, 0x8f, 0x9a, 0xfb, 0x3f, 0x15, 0x33, 0xeb,
	0xd9, 0x82, 0x45, 0x0a, 0x90, 0x37, 0xc1, 0x62, 0xfc, 0x77, 0x64, 0xb7, 0xf7, 0x8b, 0x59, 0xb2,
	0x04, 0x70, 0x70, 0x78, 0xda, 0x6d, 0xd2, 0x78, 0x9d, 0xdb, 0xab, 0xfe, 0xf9, 0xf7, 0x1f, 0x1b,
	0x5b, 0xb0, 0x99, 0x74, 0xcc, 0x22, 0x5e, 0x1b, 0x34, 0x92, 0x8e, 0xc7, 0xae, 0x6c, 0xe5, 0x04,
	0x3e, 0x69, 0xfa, 0xbe, 0x44, 0xdf, 0xec, 0x3c, 0x9a, 0xd9, 0xab, 0x1b, 0x8d, 0x1d, 0xa8, 0x3e,
	0xa0, 0xf1, 0x20, 0xa1, 0xf2, 0x1d, 0x14, 0x3b, 0x18, 0x78, 0xf7, 0x44, 0xbe, 0x34, 0x22, 0x9f,
	0x41, 0x65, 0x52, 0x64, 0x1c, 0x5a, 0xf9, 0xc7, 0x82, 0x95, 0x09, 0x1f, 0xc8, 0xb7, 0x00, 0x3d,
	0x76, 0x43, 0xb5, 0xf8, 0x05, 0x43, 0x15, 0x3f, 0xc9, 0xf9, 0xc6, 0xa7, 0x13, 0x77, 0xe3, 0xe2,
	0x24, 0xd4, 0xbb, 0x8d, 0x4b, 0x16, 0xf4, 0xd1, 0x9e, 0xeb, 0xb1, 0x9b, 0x6e, 0x0c, 0x27, 0x27,
	0x30, 0xe7, 0xf1, 0x20, 0xa0, 0xc6, 0xcb, 0x52, 0xf6, 0x11, 0xee, 0x81, 0xe8, 0x5f, 0x05, 0x18,
	0x73, 0x5b, 0x4b, 0xef, 0x5a, 0xf3, 0x64, 0x6e, 0x33, 0x93, 0x7e, 0x76, 0xc1, 0xd0, 0x4d, 0x51,
	0x7b, 0x3b, 0xa6, 0x91, 0x6d, 0xd8, 0x9a, 0x6c, 0x64, 0xa2, 0xe6, 0xca, 0xef, 0x79, 0x58, 0xb8,
	0x37, 0x55, 0x9e, 0x42, 0x3e, 0x62, 0xfa, 0x3a, 0x2e, 0x7f, 0xee, 0x38, 0x63, 0xc7, 0x2b, 0xf2,
	0x23, 0xac, 0x98, 0x41, 0x70, 0x6f, 0xbe, 0xa6, 0x55, 0xee, 0xfc, 0xff, 0xa1, 0x73, 0x9c, 0xb1,
	0x97, 0xd9, 0xd8, 0x14, 0x3b, 0x84, 0x1c, 0x73, 0x55, 0xfa, 0x16, 0xbf, 0x9e, 0xa6, 0xf5, 0xa0,
	0x93, 0xc7, 0x19, 0xdb, 0xf0, 0x49, 0x13, 0xf2, 0x0a, 0x03, 0x2f, 0x7d, 0x91, 0x2f, 0xa6, 0xe8,
	0x8c, 0x9b, 0x69, 0x7a, 0x34, 0x54, 0xf2, 0x06, 0x9e, 0xf1, 0x90, 0x6b, 0xce, 0x02, 0xea, 0xa1,
	0x76, 0xae, 0xdf, 0xbf, 0xf2, 0xfc, 0x87, 0xe6, 0xc4, 0x6a, 0xca, 0x7b, 0x6d, 0x68, 0xb7, 0x0f,
	0x9c, 0xc1, 0x53, 0x89, 0xc9, 0x49, 0xdd, 0x1b, 0xa2, 0x33, 0x1f, 0x37, 0x44, 0xc9, 0xad, 0xd8,
	0x5d, 0x76, 0x6f, 0xdb, 0x18, 0x5d, 0x86, 0x8d, 0x49, 0xa3, 0x47, 0x1b, 0x6c, 0x6d, 0xc0, 0xf3,
	0x7b, 0xc6, 0x51, 0x15, 0xa1, 0xc3, 0x3d, 0x8e, 0x92, 0xe4, 0xfe, 0x6d, 0x59, 0x3b, 0x2f, 0x00,
	0x46, 0x06, 0x73, 0x01, 0xf2, 0xcd, 0x8b, 0xee, 0x79, 0x31, 0x43, 0x66, 0x20, 0x7b, 0xd9, 0x28,
	0x5a, 0x64, 0x1e, 0x66, 0x2f, 0x77, 0x9b, 0xa7, 0xed, 0xe3, 0x66, 0x31, 0xdb, 0x7a, 0x05, 0x5f,
	0x70, 0x91, 0x14, 0x1f, 0x49, 0x71, 0x33, 0x7c, 0xbc, 0x8f, 0xd6, 0xca, 0x68, 0x15, 0x6d, 0x73,
	0x6a, 0x6d, 0xeb, 0x6a, 0x26, 0x3e, 0xbe, 0xdd, 0xff, 0x02, 0x00, 0x00, 0xff, 0xff, 0x8a, 0x8c,
	0x3c, 0xcd, 0xc6, 0x07, 0x00, 0x00,
}