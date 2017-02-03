// Code generated by protoc-gen-go.
// source: common/configtx.proto
// DO NOT EDIT!

package common

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type ConfigItem_ConfigType int32

const (
	ConfigItem_POLICY  ConfigItem_ConfigType = 0
	ConfigItem_CHAIN   ConfigItem_ConfigType = 1
	ConfigItem_ORDERER ConfigItem_ConfigType = 2
	ConfigItem_PEER    ConfigItem_ConfigType = 3
	ConfigItem_MSP     ConfigItem_ConfigType = 4
)

var ConfigItem_ConfigType_name = map[int32]string{
	0: "POLICY",
	1: "CHAIN",
	2: "ORDERER",
	3: "PEER",
	4: "MSP",
}
var ConfigItem_ConfigType_value = map[string]int32{
	"POLICY":  0,
	"CHAIN":   1,
	"ORDERER": 2,
	"PEER":    3,
	"MSP":     4,
}

func (x ConfigItem_ConfigType) String() string {
	return proto.EnumName(ConfigItem_ConfigType_name, int32(x))
}
func (ConfigItem_ConfigType) EnumDescriptor() ([]byte, []int) { return fileDescriptor1, []int{7, 0} }

// ConfigEnvelope is designed to contain _all_ configuration for a chain with no dependency
// on previous configuration transactions.
//
// It is generated with the following scheme:
//   1. Retrieve the existing configuration
//   2. Note the highest configuration sequence number, store it and increment it by one
//   3. Modify desired ConfigItems, setting each LastModified to the stored and incremented sequence number
//     a) Note that the ConfigItem has a ChainHeader header attached to it, who's type is set to CONFIGURATION_ITEM
//   4. Create Config message containing the new configuration, marshal it into ConfigEnvelope.config and encode the required signatures
//     a) Each signature is of type ConfigSignature
//     b) The ConfigSignature signature is over the concatenation of signatureHeader and the Config bytes (which includes a ChainHeader)
//   5. Submit new Config for ordering in Envelope signed by submitter
//     a) The Envelope Payload has data set to the marshaled ConfigEnvelope
//     b) The Envelope Payload has a header of type Header.Type.CONFIGURATION_TRANSACTION
//
// The configuration manager will verify:
//   1. All configuration items and the envelope refer to the correct chain
//   2. Some configuration item has been added or modified
//   3. No existing configuration item has been ommitted
//   4. All configuration changes have a LastModification of one more than the last configuration's highest LastModification number
//   5. All configuration changes satisfy the corresponding modification policy
type ConfigEnvelope struct {
	Config     []byte             `protobuf:"bytes,1,opt,name=config,proto3" json:"config,omitempty"`
	Signatures []*ConfigSignature `protobuf:"bytes,2,rep,name=signatures" json:"signatures,omitempty"`
}

func (m *ConfigEnvelope) Reset()                    { *m = ConfigEnvelope{} }
func (m *ConfigEnvelope) String() string            { return proto.CompactTextString(m) }
func (*ConfigEnvelope) ProtoMessage()               {}
func (*ConfigEnvelope) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *ConfigEnvelope) GetSignatures() []*ConfigSignature {
	if m != nil {
		return m.Signatures
	}
	return nil
}

// ConfigTemplate is used as a serialization format to share configuration templates
// The orderer supplies a configuration template to the user to use when constructing a new
// chain creation transaction, so this is used to facilitate that.
type ConfigTemplate struct {
	Items []*ConfigItem `protobuf:"bytes,1,rep,name=items" json:"items,omitempty"`
}

func (m *ConfigTemplate) Reset()                    { *m = ConfigTemplate{} }
func (m *ConfigTemplate) String() string            { return proto.CompactTextString(m) }
func (*ConfigTemplate) ProtoMessage()               {}
func (*ConfigTemplate) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

func (m *ConfigTemplate) GetItems() []*ConfigItem {
	if m != nil {
		return m.Items
	}
	return nil
}

// This message may change slightly depending on the finalization of signature schemes for transactions
type Config struct {
	Header *ChainHeader  `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	Items  []*ConfigItem `protobuf:"bytes,2,rep,name=items" json:"items,omitempty"`
}

func (m *Config) Reset()                    { *m = Config{} }
func (m *Config) String() string            { return proto.CompactTextString(m) }
func (*Config) ProtoMessage()               {}
func (*Config) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

func (m *Config) GetHeader() *ChainHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *Config) GetItems() []*ConfigItem {
	if m != nil {
		return m.Items
	}
	return nil
}

// XXX this structure is to allow us to minimize the diffs in this change series
// it will be renamed Config once the original is ready to be removed
type ConfigNext struct {
	Header  *ChainHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	Channel *ConfigGroup `protobuf:"bytes,2,opt,name=channel" json:"channel,omitempty"`
}

func (m *ConfigNext) Reset()                    { *m = ConfigNext{} }
func (m *ConfigNext) String() string            { return proto.CompactTextString(m) }
func (*ConfigNext) ProtoMessage()               {}
func (*ConfigNext) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{3} }

func (m *ConfigNext) GetHeader() *ChainHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *ConfigNext) GetChannel() *ConfigGroup {
	if m != nil {
		return m.Channel
	}
	return nil
}

// ConfigGroup is the hierarchical data structure for holding config
type ConfigGroup struct {
	Version   uint64                   `protobuf:"varint,1,opt,name=version" json:"version,omitempty"`
	Groups    map[string]*ConfigGroup  `protobuf:"bytes,2,rep,name=groups" json:"groups,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Values    map[string]*ConfigValue  `protobuf:"bytes,3,rep,name=values" json:"values,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Policies  map[string]*ConfigPolicy `protobuf:"bytes,4,rep,name=policies" json:"policies,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	ModPolicy string                   `protobuf:"bytes,5,opt,name=mod_policy,json=modPolicy" json:"mod_policy,omitempty"`
}

func (m *ConfigGroup) Reset()                    { *m = ConfigGroup{} }
func (m *ConfigGroup) String() string            { return proto.CompactTextString(m) }
func (*ConfigGroup) ProtoMessage()               {}
func (*ConfigGroup) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{4} }

func (m *ConfigGroup) GetGroups() map[string]*ConfigGroup {
	if m != nil {
		return m.Groups
	}
	return nil
}

func (m *ConfigGroup) GetValues() map[string]*ConfigValue {
	if m != nil {
		return m.Values
	}
	return nil
}

func (m *ConfigGroup) GetPolicies() map[string]*ConfigPolicy {
	if m != nil {
		return m.Policies
	}
	return nil
}

// ConfigValue represents an individual piece of config data
type ConfigValue struct {
	Version   uint64 `protobuf:"varint,1,opt,name=version" json:"version,omitempty"`
	Value     []byte `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	ModPolicy string `protobuf:"bytes,3,opt,name=mod_policy,json=modPolicy" json:"mod_policy,omitempty"`
}

func (m *ConfigValue) Reset()                    { *m = ConfigValue{} }
func (m *ConfigValue) String() string            { return proto.CompactTextString(m) }
func (*ConfigValue) ProtoMessage()               {}
func (*ConfigValue) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{5} }

type ConfigPolicy struct {
	Version   uint64  `protobuf:"varint,1,opt,name=version" json:"version,omitempty"`
	Policy    *Policy `protobuf:"bytes,2,opt,name=policy" json:"policy,omitempty"`
	ModPolicy string  `protobuf:"bytes,3,opt,name=mod_policy,json=modPolicy" json:"mod_policy,omitempty"`
}

func (m *ConfigPolicy) Reset()                    { *m = ConfigPolicy{} }
func (m *ConfigPolicy) String() string            { return proto.CompactTextString(m) }
func (*ConfigPolicy) ProtoMessage()               {}
func (*ConfigPolicy) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{6} }

func (m *ConfigPolicy) GetPolicy() *Policy {
	if m != nil {
		return m.Policy
	}
	return nil
}

type ConfigItem struct {
	Type               ConfigItem_ConfigType `protobuf:"varint,1,opt,name=type,enum=common.ConfigItem_ConfigType" json:"type,omitempty"`
	LastModified       uint64                `protobuf:"varint,2,opt,name=last_modified,json=lastModified" json:"last_modified,omitempty"`
	ModificationPolicy string                `protobuf:"bytes,3,opt,name=modification_policy,json=modificationPolicy" json:"modification_policy,omitempty"`
	Key                string                `protobuf:"bytes,4,opt,name=key" json:"key,omitempty"`
	Value              []byte                `protobuf:"bytes,5,opt,name=value,proto3" json:"value,omitempty"`
}

func (m *ConfigItem) Reset()                    { *m = ConfigItem{} }
func (m *ConfigItem) String() string            { return proto.CompactTextString(m) }
func (*ConfigItem) ProtoMessage()               {}
func (*ConfigItem) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{7} }

type ConfigSignature struct {
	SignatureHeader []byte `protobuf:"bytes,1,opt,name=signature_header,json=signatureHeader,proto3" json:"signature_header,omitempty"`
	Signature       []byte `protobuf:"bytes,2,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (m *ConfigSignature) Reset()                    { *m = ConfigSignature{} }
func (m *ConfigSignature) String() string            { return proto.CompactTextString(m) }
func (*ConfigSignature) ProtoMessage()               {}
func (*ConfigSignature) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{8} }

func init() {
	proto.RegisterType((*ConfigEnvelope)(nil), "common.ConfigEnvelope")
	proto.RegisterType((*ConfigTemplate)(nil), "common.ConfigTemplate")
	proto.RegisterType((*Config)(nil), "common.Config")
	proto.RegisterType((*ConfigNext)(nil), "common.ConfigNext")
	proto.RegisterType((*ConfigGroup)(nil), "common.ConfigGroup")
	proto.RegisterType((*ConfigValue)(nil), "common.ConfigValue")
	proto.RegisterType((*ConfigPolicy)(nil), "common.ConfigPolicy")
	proto.RegisterType((*ConfigItem)(nil), "common.ConfigItem")
	proto.RegisterType((*ConfigSignature)(nil), "common.ConfigSignature")
	proto.RegisterEnum("common.ConfigItem_ConfigType", ConfigItem_ConfigType_name, ConfigItem_ConfigType_value)
}

func init() { proto.RegisterFile("common/configtx.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 643 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x54, 0x51, 0x4f, 0xdb, 0x3c,
	0x14, 0xfd, 0xda, 0xa4, 0x29, 0xbd, 0x2d, 0x10, 0x19, 0xbe, 0x2d, 0x42, 0x43, 0x63, 0x99, 0x34,
	0x95, 0x21, 0xa8, 0xc6, 0x1e, 0x98, 0x90, 0xf6, 0xb0, 0x75, 0xd1, 0x40, 0x1a, 0xd0, 0x19, 0x34,
	0x69, 0x68, 0x52, 0x15, 0x12, 0xd3, 0x58, 0x4b, 0xe2, 0x28, 0x71, 0x11, 0x79, 0xdd, 0xcf, 0xdd,
	0xaf, 0x98, 0x62, 0x3b, 0x21, 0xed, 0x58, 0x11, 0x2f, 0x10, 0xdf, 0x7b, 0xee, 0x39, 0xc7, 0xf7,
	0xd6, 0x17, 0xfe, 0xf7, 0x58, 0x14, 0xb1, 0x78, 0xe0, 0xb1, 0xf8, 0x9a, 0x4e, 0xf8, 0xed, 0x5e,
	0x92, 0x32, 0xce, 0x90, 0x21, 0xc3, 0x1b, 0x6b, 0x55, 0xba, 0xf8, 0x27, 0x93, 0x1b, 0x65, 0x4d,
	0xc2, 0x42, 0xea, 0x51, 0x92, 0xc9, 0xb0, 0xed, 0xc2, 0xca, 0x50, 0xb0, 0x38, 0xf1, 0x0d, 0x09,
	0x59, 0x42, 0xd0, 0x13, 0x30, 0x24, 0xaf, 0xd5, 0xd8, 0x6a, 0xf4, 0x7b, 0x58, 0x9d, 0xd0, 0x01,
	0x40, 0x46, 0x27, 0xb1, 0xcb, 0xa7, 0x29, 0xc9, 0xac, 0xe6, 0x96, 0xd6, 0xef, 0xee, 0x3f, 0xdd,
	0x53, 0x1a, 0x92, 0xe3, 0xbc, 0xcc, 0xe3, 0x1a, 0xd4, 0x3e, 0x2c, 0x25, 0x2e, 0x48, 0x94, 0x84,
	0x2e, 0x27, 0xa8, 0x0f, 0x2d, 0xca, 0x49, 0x94, 0x59, 0x0d, 0xc1, 0x82, 0x66, 0x59, 0x8e, 0x39,
	0x89, 0xb0, 0x04, 0xd8, 0x63, 0x30, 0x64, 0x10, 0xed, 0x80, 0x11, 0x10, 0xd7, 0x27, 0xa9, 0xb0,
	0xd5, 0xdd, 0x5f, 0xab, 0x8a, 0x02, 0x97, 0xc6, 0x47, 0x22, 0x85, 0x15, 0xe4, 0x4e, 0xa0, 0xf9,
	0x90, 0x40, 0x00, 0x20, 0x83, 0xa7, 0xe4, 0x96, 0x3f, 0x4e, 0x64, 0x17, 0xda, 0x5e, 0xe0, 0xc6,
	0x31, 0x09, 0xad, 0xe6, 0x1c, 0x5a, 0x30, 0x7e, 0x4e, 0xd9, 0x34, 0xc1, 0x25, 0xc6, 0xfe, 0xad,
	0x41, 0xb7, 0x96, 0x40, 0x16, 0xb4, 0x6f, 0x48, 0x9a, 0x51, 0x16, 0x0b, 0x31, 0x1d, 0x97, 0x47,
	0x74, 0x00, 0xc6, 0xa4, 0x80, 0x94, 0xf6, 0x9f, 0xdf, 0xc3, 0xbb, 0x27, 0xfe, 0x66, 0x4e, 0xcc,
	0xd3, 0x1c, 0x2b, 0x78, 0x51, 0x78, 0xe3, 0x86, 0x53, 0x92, 0x59, 0xda, 0xbf, 0x0b, 0xbf, 0x09,
	0x84, 0x2a, 0x94, 0x70, 0xf4, 0x1e, 0x96, 0xca, 0xdf, 0x85, 0xa5, 0x8b, 0xd2, 0x17, 0xf7, 0x95,
	0x8e, 0x14, 0x46, 0x16, 0x57, 0x25, 0x68, 0x13, 0x20, 0x62, 0xfe, 0x58, 0x9c, 0x73, 0xab, 0xb5,
	0xd5, 0xe8, 0x77, 0x70, 0x27, 0x62, 0xbe, 0xc0, 0xe7, 0x1b, 0xa7, 0xd0, 0xad, 0xb9, 0x45, 0x26,
	0x68, 0x3f, 0x49, 0x2e, 0x2e, 0xdd, 0xc1, 0xc5, 0x27, 0xda, 0x86, 0x96, 0x30, 0xb2, 0xa8, 0x8f,
	0x12, 0x71, 0xd8, 0x7c, 0xd7, 0x28, 0xf8, 0x6a, 0x97, 0x78, 0x34, 0x9f, 0xa8, 0xad, 0xf3, 0x7d,
	0x85, 0xe5, 0x99, 0x9b, 0xdd, 0xc3, 0xf8, 0x7a, 0x96, 0x71, 0x7d, 0x96, 0x51, 0xde, 0xb3, 0x46,
	0x69, 0xff, 0x28, 0x67, 0x2d, 0xc4, 0x16, 0xcc, 0x7a, 0xbd, 0x4e, 0xdc, 0x53, 0x14, 0x73, 0x0d,
	0xd5, 0xe6, 0x1a, 0x6a, 0x33, 0xe8, 0xd5, 0x85, 0x17, 0xd0, 0xbf, 0x02, 0x43, 0x91, 0x48, 0xe3,
	0x2b, 0xa5, 0x71, 0x65, 0x59, 0x65, 0x1f, 0x12, 0xfc, 0xd5, 0x2c, 0x9f, 0x49, 0xf1, 0x76, 0xd0,
	0x1b, 0xd0, 0x79, 0x9e, 0x10, 0x21, 0xb6, 0xb2, 0xbf, 0xf9, 0xf7, 0xeb, 0x52, 0x9f, 0x17, 0x79,
	0x42, 0xb0, 0x80, 0xa2, 0x97, 0xb0, 0x1c, 0xba, 0x19, 0x1f, 0x47, 0xcc, 0xa7, 0xd7, 0x94, 0xf8,
	0xc2, 0x8f, 0x8e, 0x7b, 0x45, 0xf0, 0x44, 0xc5, 0xd0, 0x00, 0xd6, 0x64, 0xde, 0x73, 0x39, 0x65,
	0xf1, 0xac, 0x1d, 0x54, 0x4f, 0xa9, 0x8b, 0xab, 0x41, 0xe9, 0x77, 0x83, 0xaa, 0xfa, 0xd9, 0xaa,
	0xf5, 0xd3, 0x1e, 0x96, 0xf6, 0x0b, 0x47, 0x08, 0xc0, 0x18, 0x9d, 0x7d, 0x39, 0x1e, 0x7e, 0x37,
	0xff, 0x43, 0x1d, 0x68, 0x0d, 0x8f, 0x3e, 0x1c, 0x9f, 0x9a, 0x0d, 0xd4, 0x85, 0xf6, 0x19, 0xfe,
	0xe4, 0x60, 0x07, 0x9b, 0x4d, 0xb4, 0x04, 0xfa, 0xc8, 0x71, 0xb0, 0xa9, 0xa1, 0x36, 0x68, 0x27,
	0xe7, 0x23, 0x53, 0xb7, 0x2f, 0x61, 0x75, 0x6e, 0xcd, 0xa1, 0x6d, 0x30, 0xab, 0x45, 0x37, 0xae,
	0x6d, 0x8e, 0x1e, 0x5e, 0xad, 0xe2, 0x72, 0x6b, 0xa0, 0x67, 0xd0, 0xa9, 0x42, 0x6a, 0xd8, 0x77,
	0x81, 0x8f, 0xbb, 0x97, 0x3b, 0x13, 0xca, 0x83, 0xe9, 0x55, 0xd1, 0xcb, 0x41, 0x90, 0x27, 0x24,
	0x0d, 0x89, 0x3f, 0x21, 0xe9, 0xe0, 0xda, 0xbd, 0x4a, 0xa9, 0x37, 0x10, 0xdb, 0x3a, 0x53, 0x2b,
	0xfd, 0xca, 0x10, 0xc7, 0xb7, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0xec, 0xdf, 0xfc, 0x61, 0x09,
	0x06, 0x00, 0x00,
}
