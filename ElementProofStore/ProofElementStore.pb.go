// Code generated by protoc-gen-go.
// source: ElementProofStore.proto
// DO NOT EDIT!

/*
Package ElementProofStore is a generated protocol buffer package.

It is generated from these files:
	ElementProofStore.proto

It has these top-level messages:
	SECPElementProofStore
	SECPSHA2ElementProofStore
*/
package ElementProofStore

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type SECPElementProofStore_ElementProofState int32

const (
	SECPElementProofStore_Initialized SECPElementProofStore_ElementProofState = 0
	SECPElementProofStore_Signed      SECPElementProofStore_ElementProofState = 1
	SECPElementProofStore_Revoked     SECPElementProofStore_ElementProofState = 2
	SECPElementProofStore_SuperCeded  SECPElementProofStore_ElementProofState = 3
)

var SECPElementProofStore_ElementProofState_name = map[int32]string{
	0: "Initialized",
	1: "Signed",
	2: "Revoked",
	3: "SuperCeded",
}
var SECPElementProofStore_ElementProofState_value = map[string]int32{
	"Initialized": 0,
	"Signed":      1,
	"Revoked":     2,
	"SuperCeded":  3,
}

func (x SECPElementProofStore_ElementProofState) String() string {
	return proto.EnumName(SECPElementProofStore_ElementProofState_name, int32(x))
}

type SECPSHA2ElementProofStore_ElementProofState int32

const (
	SECPSHA2ElementProofStore_Initialized SECPSHA2ElementProofStore_ElementProofState = 0
	SECPSHA2ElementProofStore_Signed      SECPSHA2ElementProofStore_ElementProofState = 1
	SECPSHA2ElementProofStore_Revoked     SECPSHA2ElementProofStore_ElementProofState = 2
	SECPSHA2ElementProofStore_SuperCeded  SECPSHA2ElementProofStore_ElementProofState = 3
)

var SECPSHA2ElementProofStore_ElementProofState_name = map[int32]string{
	0: "Initialized",
	1: "Signed",
	2: "Revoked",
	3: "SuperCeded",
}
var SECPSHA2ElementProofStore_ElementProofState_value = map[string]int32{
	"Initialized": 0,
	"Signed":      1,
	"Revoked":     2,
	"SuperCeded":  3,
}

func (x SECPSHA2ElementProofStore_ElementProofState) String() string {
	return proto.EnumName(SECPSHA2ElementProofStore_ElementProofState_name, int32(x))
}

type SECPElementProofStore struct {
	Name       string   `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Threshold  int32    `protobuf:"varint,3,opt,name=threshold" json:"threshold,omitempty"`
	Data       string   `protobuf:"bytes,4,opt,name=data" json:"data,omitempty"`
	PublicKeys [][]byte `protobuf:"bytes,5,rep,name=PublicKeys,proto3" json:"PublicKeys,omitempty"`
	Signatures [][]byte `protobuf:"bytes,6,rep,name=Signatures,proto3" json:"Signatures,omitempty"`
	Supercede  string   `protobuf:"bytes,7,opt,name=supercede" json:"supercede,omitempty"`
}

func (m *SECPElementProofStore) Reset()         { *m = SECPElementProofStore{} }
func (m *SECPElementProofStore) String() string { return proto.CompactTextString(m) }
func (*SECPElementProofStore) ProtoMessage()    {}

type SECPSHA2ElementProofStore struct {
	Name       string   `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Threshold  int32    `protobuf:"varint,3,opt,name=threshold" json:"threshold,omitempty"`
	Data       string   `protobuf:"bytes,4,opt,name=data" json:"data,omitempty"`
	PublicKeys [][]byte `protobuf:"bytes,5,rep,name=PublicKeys,proto3" json:"PublicKeys,omitempty"`
	Signatures [][]byte `protobuf:"bytes,6,rep,name=Signatures,proto3" json:"Signatures,omitempty"`
	Digests    [][]byte `protobuf:"bytes,7,rep,name=Digests,proto3" json:"Digests,omitempty"`
	Preimages  [][]byte `protobuf:"bytes,8,rep,name=Preimages,proto3" json:"Preimages,omitempty"`
	Supercede  string   `protobuf:"bytes,9,opt,name=supercede" json:"supercede,omitempty"`
}

func (m *SECPSHA2ElementProofStore) Reset()         { *m = SECPSHA2ElementProofStore{} }
func (m *SECPSHA2ElementProofStore) String() string { return proto.CompactTextString(m) }
func (*SECPSHA2ElementProofStore) ProtoMessage()    {}

func init() {
	proto.RegisterEnum("SECPElementProofStore_ElementProofState", SECPElementProofStore_ElementProofState_name, SECPElementProofStore_ElementProofState_value)
	proto.RegisterEnum("SECPSHA2ElementProofStore_ElementProofState", SECPSHA2ElementProofStore_ElementProofState_name, SECPSHA2ElementProofStore_ElementProofState_value)
}