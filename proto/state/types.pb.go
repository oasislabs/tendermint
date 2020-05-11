// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: proto/state/types.proto

package state

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/timestamp"
	types "github.com/tendermint/tendermint/abci/types"
	types1 "github.com/tendermint/tendermint/proto/types"
	version "github.com/tendermint/tendermint/proto/version"
	math "math"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// ABCIResponses retains the responses
// of the various ABCI calls during block processing.
// It is persisted to disk for each height before calling Commit.
type ABCIResponses struct {
	DeliverTxs           []*types.ResponseDeliverTx `protobuf:"bytes,1,rep,name=deliver_txs,json=deliverTxs,proto3" json:"deliver_txs,omitempty"`
	EndBlock             *types.ResponseEndBlock    `protobuf:"bytes,2,opt,name=end_block,json=endBlock,proto3" json:"end_block,omitempty"`
	BeginBlock           *types.ResponseBeginBlock  `protobuf:"bytes,3,opt,name=begin_block,json=beginBlock,proto3" json:"begin_block,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *ABCIResponses) Reset()         { *m = ABCIResponses{} }
func (m *ABCIResponses) String() string { return proto.CompactTextString(m) }
func (*ABCIResponses) ProtoMessage()    {}
func (*ABCIResponses) Descriptor() ([]byte, []int) {
	return fileDescriptor_00e69fef8162ea9b, []int{0}
}
func (m *ABCIResponses) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ABCIResponses.Unmarshal(m, b)
}
func (m *ABCIResponses) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ABCIResponses.Marshal(b, m, deterministic)
}
func (m *ABCIResponses) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ABCIResponses.Merge(m, src)
}
func (m *ABCIResponses) XXX_Size() int {
	return xxx_messageInfo_ABCIResponses.Size(m)
}
func (m *ABCIResponses) XXX_DiscardUnknown() {
	xxx_messageInfo_ABCIResponses.DiscardUnknown(m)
}

var xxx_messageInfo_ABCIResponses proto.InternalMessageInfo

func (m *ABCIResponses) GetDeliverTxs() []*types.ResponseDeliverTx {
	if m != nil {
		return m.DeliverTxs
	}
	return nil
}

func (m *ABCIResponses) GetEndBlock() *types.ResponseEndBlock {
	if m != nil {
		return m.EndBlock
	}
	return nil
}

func (m *ABCIResponses) GetBeginBlock() *types.ResponseBeginBlock {
	if m != nil {
		return m.BeginBlock
	}
	return nil
}

// ValidatorsInfo represents the latest validator set, or the last height it changed
type ValidatorsInfo struct {
	ValidatorSet         *types1.ValidatorSet `protobuf:"bytes,1,opt,name=validator_set,json=validatorSet,proto3" json:"validator_set,omitempty"`
	LastHeightChanged    int64                `protobuf:"varint,2,opt,name=last_height_changed,json=lastHeightChanged,proto3" json:"last_height_changed,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *ValidatorsInfo) Reset()         { *m = ValidatorsInfo{} }
func (m *ValidatorsInfo) String() string { return proto.CompactTextString(m) }
func (*ValidatorsInfo) ProtoMessage()    {}
func (*ValidatorsInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_00e69fef8162ea9b, []int{1}
}
func (m *ValidatorsInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ValidatorsInfo.Unmarshal(m, b)
}
func (m *ValidatorsInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ValidatorsInfo.Marshal(b, m, deterministic)
}
func (m *ValidatorsInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ValidatorsInfo.Merge(m, src)
}
func (m *ValidatorsInfo) XXX_Size() int {
	return xxx_messageInfo_ValidatorsInfo.Size(m)
}
func (m *ValidatorsInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ValidatorsInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ValidatorsInfo proto.InternalMessageInfo

func (m *ValidatorsInfo) GetValidatorSet() *types1.ValidatorSet {
	if m != nil {
		return m.ValidatorSet
	}
	return nil
}

func (m *ValidatorsInfo) GetLastHeightChanged() int64 {
	if m != nil {
		return m.LastHeightChanged
	}
	return 0
}

// ConsensusParamsInfo represents the latest consensus params, or the last height it changed
type ConsensusParamsInfo struct {
	ConsensusParams      types1.ConsensusParams `protobuf:"bytes,1,opt,name=consensus_params,json=consensusParams,proto3" json:"consensus_params"`
	LastHeightChanged    int64                  `protobuf:"varint,2,opt,name=last_height_changed,json=lastHeightChanged,proto3" json:"last_height_changed,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *ConsensusParamsInfo) Reset()         { *m = ConsensusParamsInfo{} }
func (m *ConsensusParamsInfo) String() string { return proto.CompactTextString(m) }
func (*ConsensusParamsInfo) ProtoMessage()    {}
func (*ConsensusParamsInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_00e69fef8162ea9b, []int{2}
}
func (m *ConsensusParamsInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConsensusParamsInfo.Unmarshal(m, b)
}
func (m *ConsensusParamsInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConsensusParamsInfo.Marshal(b, m, deterministic)
}
func (m *ConsensusParamsInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConsensusParamsInfo.Merge(m, src)
}
func (m *ConsensusParamsInfo) XXX_Size() int {
	return xxx_messageInfo_ConsensusParamsInfo.Size(m)
}
func (m *ConsensusParamsInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ConsensusParamsInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ConsensusParamsInfo proto.InternalMessageInfo

func (m *ConsensusParamsInfo) GetConsensusParams() types1.ConsensusParams {
	if m != nil {
		return m.ConsensusParams
	}
	return types1.ConsensusParams{}
}

func (m *ConsensusParamsInfo) GetLastHeightChanged() int64 {
	if m != nil {
		return m.LastHeightChanged
	}
	return 0
}

type Version struct {
	Consensus            version.Consensus `protobuf:"bytes,1,opt,name=consensus,proto3" json:"consensus"`
	Software             string            `protobuf:"bytes,2,opt,name=software,proto3" json:"software,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Version) Reset()         { *m = Version{} }
func (m *Version) String() string { return proto.CompactTextString(m) }
func (*Version) ProtoMessage()    {}
func (*Version) Descriptor() ([]byte, []int) {
	return fileDescriptor_00e69fef8162ea9b, []int{3}
}
func (m *Version) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Version.Unmarshal(m, b)
}
func (m *Version) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Version.Marshal(b, m, deterministic)
}
func (m *Version) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Version.Merge(m, src)
}
func (m *Version) XXX_Size() int {
	return xxx_messageInfo_Version.Size(m)
}
func (m *Version) XXX_DiscardUnknown() {
	xxx_messageInfo_Version.DiscardUnknown(m)
}

var xxx_messageInfo_Version proto.InternalMessageInfo

func (m *Version) GetConsensus() version.Consensus {
	if m != nil {
		return m.Consensus
	}
	return version.Consensus{}
}

func (m *Version) GetSoftware() string {
	if m != nil {
		return m.Software
	}
	return ""
}

type State struct {
	Version Version `protobuf:"bytes,1,opt,name=version,proto3" json:"version"`
	// immutable
	ChainID string `protobuf:"bytes,2,opt,name=chain_Id,json=chainId,proto3" json:"chain_Id,omitempty"`
	// LastBlockHeight=0 at genesis (ie. block(H=0) does not exist)
	LastBlockHeight int64          `protobuf:"varint,3,opt,name=last_block_height,json=lastBlockHeight,proto3" json:"last_block_height,omitempty"`
	LastBlockID     types1.BlockID `protobuf:"bytes,4,opt,name=last_block_id,json=lastBlockId,proto3" json:"last_block_id"`
	LastBlockTime   time.Time      `protobuf:"bytes,5,opt,name=last_block_time,json=lastBlockTime,proto3,stdtime" json:"last_block_time"`
	// LastValidators is used to validate block.LastCommit.
	// Validators are persisted to the database separately every time they change,
	// so we can query for historical validator sets.
	// Note that if s.LastBlockHeight causes a valset change,
	// we set s.LastHeightValidatorsChanged = s.LastBlockHeight + 1 + 1
	// Extra +1 due to nextValSet delay.
	NextValidators              *types1.ValidatorSet `protobuf:"bytes,6,opt,name=next_validators,json=nextValidators,proto3" json:"next_validators,omitempty"`
	Validators                  *types1.ValidatorSet `protobuf:"bytes,7,opt,name=validators,proto3" json:"validators,omitempty"`
	LastValidators              *types1.ValidatorSet `protobuf:"bytes,8,opt,name=last_validators,json=lastValidators,proto3" json:"last_validators,omitempty"`
	LastHeightValidatorsChanged int64                `protobuf:"varint,9,opt,name=last_height_validators_changed,json=lastHeightValidatorsChanged,proto3" json:"last_height_validators_changed,omitempty"`
	// Consensus parameters used for validating blocks.
	// Changes returned by EndBlock and updated after Commit.
	ConsensusParams                  types1.ConsensusParams `protobuf:"bytes,10,opt,name=consensus_params,json=consensusParams,proto3" json:"consensus_params"`
	LastHeightConsensusParamsChanged int64                  `protobuf:"varint,11,opt,name=last_height_consensus_params_changed,json=lastHeightConsensusParamsChanged,proto3" json:"last_height_consensus_params_changed,omitempty"`
	// Merkle root of the results from executing prev block
	LastResultsHash []byte `protobuf:"bytes,12,opt,name=LastResultsHash,proto3" json:"LastResultsHash,omitempty"`
	// the latest AppHash we've received from calling abci.Commit()
	AppHash              []byte   `protobuf:"bytes,13,opt,name=AppHash,proto3" json:"AppHash,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *State) Reset()         { *m = State{} }
func (m *State) String() string { return proto.CompactTextString(m) }
func (*State) ProtoMessage()    {}
func (*State) Descriptor() ([]byte, []int) {
	return fileDescriptor_00e69fef8162ea9b, []int{4}
}
func (m *State) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_State.Unmarshal(m, b)
}
func (m *State) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_State.Marshal(b, m, deterministic)
}
func (m *State) XXX_Merge(src proto.Message) {
	xxx_messageInfo_State.Merge(m, src)
}
func (m *State) XXX_Size() int {
	return xxx_messageInfo_State.Size(m)
}
func (m *State) XXX_DiscardUnknown() {
	xxx_messageInfo_State.DiscardUnknown(m)
}

var xxx_messageInfo_State proto.InternalMessageInfo

func (m *State) GetVersion() Version {
	if m != nil {
		return m.Version
	}
	return Version{}
}

func (m *State) GetChainID() string {
	if m != nil {
		return m.ChainID
	}
	return ""
}

func (m *State) GetLastBlockHeight() int64 {
	if m != nil {
		return m.LastBlockHeight
	}
	return 0
}

func (m *State) GetLastBlockID() types1.BlockID {
	if m != nil {
		return m.LastBlockID
	}
	return types1.BlockID{}
}

func (m *State) GetLastBlockTime() time.Time {
	if m != nil {
		return m.LastBlockTime
	}
	return time.Time{}
}

func (m *State) GetNextValidators() *types1.ValidatorSet {
	if m != nil {
		return m.NextValidators
	}
	return nil
}

func (m *State) GetValidators() *types1.ValidatorSet {
	if m != nil {
		return m.Validators
	}
	return nil
}

func (m *State) GetLastValidators() *types1.ValidatorSet {
	if m != nil {
		return m.LastValidators
	}
	return nil
}

func (m *State) GetLastHeightValidatorsChanged() int64 {
	if m != nil {
		return m.LastHeightValidatorsChanged
	}
	return 0
}

func (m *State) GetConsensusParams() types1.ConsensusParams {
	if m != nil {
		return m.ConsensusParams
	}
	return types1.ConsensusParams{}
}

func (m *State) GetLastHeightConsensusParamsChanged() int64 {
	if m != nil {
		return m.LastHeightConsensusParamsChanged
	}
	return 0
}

func (m *State) GetLastResultsHash() []byte {
	if m != nil {
		return m.LastResultsHash
	}
	return nil
}

func (m *State) GetAppHash() []byte {
	if m != nil {
		return m.AppHash
	}
	return nil
}

func init() {
	proto.RegisterType((*ABCIResponses)(nil), "tendermint.proto.state.ABCIResponses")
	proto.RegisterType((*ValidatorsInfo)(nil), "tendermint.proto.state.ValidatorsInfo")
	proto.RegisterType((*ConsensusParamsInfo)(nil), "tendermint.proto.state.ConsensusParamsInfo")
	proto.RegisterType((*Version)(nil), "tendermint.proto.state.Version")
	proto.RegisterType((*State)(nil), "tendermint.proto.state.State")
}

func init() { proto.RegisterFile("proto/state/types.proto", fileDescriptor_00e69fef8162ea9b) }

var fileDescriptor_00e69fef8162ea9b = []byte{
	// 722 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x55, 0x5d, 0x6a, 0xdb, 0x40,
	0x10, 0xae, 0xea, 0x24, 0xb6, 0x47, 0x71, 0xdc, 0x6e, 0x20, 0x15, 0x0e, 0xd4, 0xc6, 0x0d, 0x89,
	0x5b, 0x8a, 0x0c, 0xe9, 0x01, 0x4a, 0x64, 0x97, 0x46, 0x25, 0x2d, 0x45, 0x09, 0x21, 0xf4, 0x45,
	0xc8, 0xd6, 0x46, 0x12, 0xb5, 0x25, 0xa1, 0x5d, 0xbb, 0xc9, 0x19, 0xfa, 0xd2, 0x1b, 0xf4, 0x3a,
	0xbd, 0x43, 0x21, 0x85, 0x3c, 0xf7, 0x10, 0x65, 0x7f, 0x24, 0x6d, 0x9c, 0x84, 0x60, 0xe8, 0x93,
	0x57, 0x33, 0xf3, 0x7d, 0xf3, 0xcd, 0xee, 0x37, 0x18, 0x9e, 0xa5, 0x59, 0x42, 0x93, 0x3e, 0xa1,
	0x1e, 0xc5, 0x7d, 0x7a, 0x99, 0x62, 0x62, 0xf2, 0x08, 0xda, 0xa2, 0x38, 0xf6, 0x71, 0x36, 0x8d,
	0x62, 0x2a, 0x22, 0x26, 0xaf, 0x69, 0xed, 0xd2, 0x30, 0xca, 0x7c, 0x37, 0xf5, 0x32, 0x7a, 0xd9,
	0x17, 0xe0, 0x20, 0x09, 0x92, 0xf2, 0x24, 0xaa, 0x5b, 0x5b, 0xde, 0x68, 0x1c, 0x09, 0x46, 0x95,
	0xb7, 0x25, 0x1b, 0xde, 0x4e, 0x6c, 0xab, 0x89, 0xb9, 0x37, 0x89, 0x7c, 0x8f, 0x26, 0x99, 0x4c,
	0x1a, 0x6a, 0x32, 0xf5, 0x32, 0x6f, 0xba, 0x00, 0x9b, 0xe3, 0x8c, 0x44, 0x49, 0x9c, 0xff, 0xca,
	0x64, 0x3b, 0x48, 0x92, 0x60, 0x82, 0x85, 0xce, 0xd1, 0xec, 0xbc, 0x4f, 0xa3, 0x29, 0x26, 0xd4,
	0x9b, 0xa6, 0xa2, 0xa0, 0xfb, 0x57, 0x83, 0xc6, 0x81, 0x35, 0xb0, 0x1d, 0x4c, 0xd2, 0x24, 0x26,
	0x98, 0x20, 0x1b, 0x74, 0x1f, 0x4f, 0xa2, 0x39, 0xce, 0x5c, 0x7a, 0x41, 0x0c, 0xad, 0x53, 0xe9,
	0xe9, 0xfb, 0x3d, 0x53, 0xb9, 0x0d, 0x36, 0x98, 0x29, 0x94, 0xe7, 0xb0, 0xa1, 0x40, 0x9c, 0x5c,
	0x38, 0xe0, 0xe7, 0x47, 0x82, 0x86, 0x50, 0xc7, 0xb1, 0xef, 0x8e, 0x26, 0xc9, 0xf8, 0xab, 0xf1,
	0xb8, 0xa3, 0xf5, 0xf4, 0xfd, 0xbd, 0x07, 0x88, 0xde, 0xc5, 0xbe, 0xc5, 0xca, 0x9d, 0x1a, 0x96,
	0x27, 0xf4, 0x01, 0xf4, 0x11, 0x0e, 0xa2, 0x58, 0xf2, 0x54, 0x38, 0xcf, 0xcb, 0x07, 0x78, 0x2c,
	0x86, 0x10, 0x4c, 0x30, 0x2a, 0xce, 0xdd, 0xef, 0x1a, 0x6c, 0x9c, 0xe6, 0x57, 0x4b, 0xec, 0xf8,
	0x3c, 0x41, 0x36, 0x34, 0x8a, 0xcb, 0x76, 0x09, 0xa6, 0x86, 0xc6, 0x1b, 0xec, 0x98, 0xb7, 0xde,
	0x5f, 0x74, 0x28, 0xe0, 0xc7, 0x98, 0x3a, 0xeb, 0x73, 0xe5, 0x0b, 0x99, 0xb0, 0x39, 0xf1, 0x08,
	0x75, 0x43, 0x1c, 0x05, 0x21, 0x75, 0xc7, 0xa1, 0x17, 0x07, 0xd8, 0xe7, 0x93, 0x57, 0x9c, 0xa7,
	0x2c, 0x75, 0xc8, 0x33, 0x03, 0x91, 0xe8, 0xfe, 0xd4, 0x60, 0x73, 0xc0, 0xd4, 0xc6, 0x64, 0x46,
	0x3e, 0xf3, 0x47, 0xe5, 0x92, 0xce, 0xe0, 0xc9, 0x38, 0x0f, 0xbb, 0xe2, 0xb1, 0xa5, 0xaa, 0xbd,
	0xfb, 0x54, 0x2d, 0xd0, 0x58, 0x2b, 0xbf, 0xae, 0xda, 0x8f, 0x9c, 0xe6, 0xf8, 0x66, 0x78, 0x69,
	0x85, 0x31, 0x54, 0x4f, 0x85, 0xa1, 0xd0, 0x7b, 0xa8, 0x17, 0x6c, 0x52, 0xcd, 0x8b, 0xdb, 0x6a,
	0x72, 0xfb, 0x15, 0x7a, 0xa4, 0x92, 0x12, 0x8b, 0x5a, 0x50, 0x23, 0xc9, 0x39, 0xfd, 0xe6, 0x65,
	0x98, 0x37, 0xae, 0x3b, 0xc5, 0x77, 0xf7, 0xf7, 0x1a, 0xac, 0x1e, 0xb3, 0x35, 0x43, 0x6f, 0xa1,
	0x2a, 0xb9, 0x64, 0xb3, 0xb6, 0x79, 0xf7, 0x42, 0x9a, 0x52, 0xa0, 0x6c, 0x94, 0xa3, 0xd0, 0x2e,
	0xd4, 0xc6, 0xa1, 0x17, 0xc5, 0xae, 0x2d, 0xe6, 0xab, 0x5b, 0xfa, 0xf5, 0x55, 0xbb, 0x3a, 0x60,
	0x31, 0x7b, 0xe8, 0x54, 0x79, 0xd2, 0xf6, 0xd1, 0x2b, 0xe0, 0x73, 0x0b, 0x77, 0xc9, 0x8b, 0xe1,
	0x26, 0xab, 0x38, 0x4d, 0x96, 0xe0, 0xc6, 0x11, 0xb7, 0x82, 0xce, 0xa0, 0xa1, 0xd4, 0x46, 0xbe,
	0xb1, 0x72, 0x9f, 0x34, 0xf1, 0x2a, 0x1c, 0x6b, 0x0f, 0xad, 0x4d, 0x26, 0xed, 0xfa, 0xaa, 0xad,
	0x1f, 0xe5, 0x84, 0xf6, 0xd0, 0xd1, 0x0b, 0x76, 0xdb, 0x47, 0x47, 0xd0, 0x54, 0x98, 0xd9, 0x96,
	0x1a, 0xab, 0x9c, 0xbb, 0x65, 0x8a, 0x15, 0x36, 0xf3, 0x15, 0x36, 0x4f, 0xf2, 0x15, 0xb6, 0x6a,
	0x8c, 0xf6, 0xc7, 0x9f, 0xb6, 0xe6, 0x34, 0x0a, 0x2e, 0x96, 0x45, 0x1f, 0xa1, 0x19, 0xe3, 0x0b,
	0xea, 0x16, 0xee, 0x24, 0xc6, 0xda, 0x12, 0xae, 0xde, 0x60, 0xe0, 0x72, 0x4d, 0xd0, 0x10, 0x40,
	0x61, 0xaa, 0x2e, 0xc1, 0xa4, 0xe0, 0x98, 0x28, 0x3e, 0xa2, 0x42, 0x55, 0x5b, 0x46, 0x14, 0x03,
	0x2b, 0xa2, 0x06, 0xf0, 0x5c, 0xb5, 0x72, 0xc9, 0x5a, 0xb8, 0xba, 0xce, 0x1f, 0x71, 0xbb, 0x74,
	0x75, 0x89, 0x96, 0xfe, 0xbe, 0x73, 0xd3, 0xe0, 0xbf, 0x6c, 0xda, 0x27, 0xd8, 0xb9, 0xb1, 0x69,
	0x0b, 0x5d, 0x0a, 0x91, 0x3a, 0x17, 0xd9, 0x51, 0x56, 0xef, 0x26, 0x51, 0xae, 0xb4, 0x07, 0x4d,
	0x66, 0x1e, 0x07, 0x93, 0xd9, 0x84, 0x92, 0x43, 0x8f, 0x84, 0xc6, 0x7a, 0x47, 0xeb, 0xad, 0x3b,
	0x8b, 0x61, 0x64, 0x40, 0xf5, 0x20, 0x4d, 0x79, 0x45, 0x83, 0x57, 0xe4, 0x9f, 0x96, 0xf9, 0xe5,
	0x75, 0x10, 0xd1, 0x70, 0x36, 0x32, 0xc7, 0xc9, 0xb4, 0x5f, 0xce, 0xa7, 0x1e, 0x95, 0xbf, 0xc3,
	0xd1, 0x1a, 0xff, 0x78, 0xf3, 0x2f, 0x00, 0x00, 0xff, 0xff, 0x78, 0x53, 0xd2, 0x08, 0x24, 0x07,
	0x00, 0x00,
}