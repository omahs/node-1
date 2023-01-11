// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: akash/escrow/v1beta2/types.proto

package v1beta2

import (
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
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

// State stores state for an escrow account
type Account_State int32

const (
	// AccountStateInvalid is an invalid state
	AccountStateInvalid Account_State = 0
	// AccountOpen is the state when an account is open
	AccountOpen Account_State = 1
	// AccountClosed is the state when an account is closed
	AccountClosed Account_State = 2
	// AccountOverdrawn is the state when an account is overdrawn
	AccountOverdrawn Account_State = 3
)

var Account_State_name = map[int32]string{
	0: "invalid",
	1: "open",
	2: "closed",
	3: "overdrawn",
}

var Account_State_value = map[string]int32{
	"invalid":   0,
	"open":      1,
	"closed":    2,
	"overdrawn": 3,
}

func (x Account_State) String() string {
	return proto.EnumName(Account_State_name, int32(x))
}

func (Account_State) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_5b25ee303c78038d, []int{1, 0}
}

// Payment State
type FractionalPayment_State int32

const (
	// PaymentStateInvalid is the state when the payment is invalid
	PaymentStateInvalid FractionalPayment_State = 0
	// PaymentStateOpen is the state when the payment is open
	PaymentOpen FractionalPayment_State = 1
	// PaymentStateClosed is the state when the payment is closed
	PaymentClosed FractionalPayment_State = 2
	// PaymentStateOverdrawn is the state when the payment is overdrawn
	PaymentOverdrawn FractionalPayment_State = 3
)

var FractionalPayment_State_name = map[int32]string{
	0: "invalid",
	1: "open",
	2: "closed",
	3: "overdrawn",
}

var FractionalPayment_State_value = map[string]int32{
	"invalid":   0,
	"open":      1,
	"closed":    2,
	"overdrawn": 3,
}

func (x FractionalPayment_State) String() string {
	return proto.EnumName(FractionalPayment_State_name, int32(x))
}

func (FractionalPayment_State) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_5b25ee303c78038d, []int{2, 0}
}

// AccountID is the account identifier
type AccountID struct {
	Scope string `protobuf:"bytes,1,opt,name=scope,proto3" json:"scope" yaml:"scope"`
	XID   string `protobuf:"bytes,2,opt,name=xid,proto3" json:"xid" yaml:"xid"`
}

func (m *AccountID) Reset()         { *m = AccountID{} }
func (m *AccountID) String() string { return proto.CompactTextString(m) }
func (*AccountID) ProtoMessage()    {}
func (*AccountID) Descriptor() ([]byte, []int) {
	return fileDescriptor_5b25ee303c78038d, []int{0}
}
func (m *AccountID) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AccountID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AccountID.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AccountID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccountID.Merge(m, src)
}
func (m *AccountID) XXX_Size() int {
	return m.Size()
}
func (m *AccountID) XXX_DiscardUnknown() {
	xxx_messageInfo_AccountID.DiscardUnknown(m)
}

var xxx_messageInfo_AccountID proto.InternalMessageInfo

func (m *AccountID) GetScope() string {
	if m != nil {
		return m.Scope
	}
	return ""
}

func (m *AccountID) GetXID() string {
	if m != nil {
		return m.XID
	}
	return ""
}

// Account stores state for an escrow account
type Account struct {
	// unique identifier for this escrow account
	ID AccountID `protobuf:"bytes,1,opt,name=id,proto3" json:"id" yaml:"id"`
	// bech32 encoded account address of the owner of this escrow account
	Owner string `protobuf:"bytes,2,opt,name=owner,proto3" json:"owner" yaml:"owner"`
	// current state of this escrow account
	State Account_State `protobuf:"varint,3,opt,name=state,proto3,enum=akash.escrow.v1beta2.Account_State" json:"state" yaml:"state"`
	// unspent coins received from the owner's wallet
	Balance types.DecCoin `protobuf:"bytes,4,opt,name=balance,proto3" json:"balance" yaml:"balance"`
	// total coins spent by this account
	Transferred types.DecCoin `protobuf:"bytes,5,opt,name=transferred,proto3" json:"transferred" yaml:"transferred"`
	// block height at which this account was last settled
	SettledAt int64 `protobuf:"varint,6,opt,name=settled_at,json=settledAt,proto3" json:"settledAt" yaml:"settledAt"`
	// bech32 encoded account address of the depositor.
	// If depositor is same as the owner, then any incoming coins are added to the Balance.
	// If depositor isn't same as the owner, then any incoming coins are added to the Funds.
	Depositor string `protobuf:"bytes,7,opt,name=depositor,proto3" json:"depositor" yaml:"depositor"`
	// Funds are unspent coins received from the (non-Owner) Depositor's wallet.
	// If there are any funds, they should be spent before spending the Balance.
	Funds types.DecCoin `protobuf:"bytes,8,opt,name=funds,proto3" json:"funds" yaml:"funds"`
}

func (m *Account) Reset()         { *m = Account{} }
func (m *Account) String() string { return proto.CompactTextString(m) }
func (*Account) ProtoMessage()    {}
func (*Account) Descriptor() ([]byte, []int) {
	return fileDescriptor_5b25ee303c78038d, []int{1}
}
func (m *Account) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Account) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Account.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Account) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Account.Merge(m, src)
}
func (m *Account) XXX_Size() int {
	return m.Size()
}
func (m *Account) XXX_DiscardUnknown() {
	xxx_messageInfo_Account.DiscardUnknown(m)
}

var xxx_messageInfo_Account proto.InternalMessageInfo

func (m *Account) GetID() AccountID {
	if m != nil {
		return m.ID
	}
	return AccountID{}
}

func (m *Account) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *Account) GetState() Account_State {
	if m != nil {
		return m.State
	}
	return AccountStateInvalid
}

func (m *Account) GetBalance() types.DecCoin {
	if m != nil {
		return m.Balance
	}
	return types.DecCoin{}
}

func (m *Account) GetTransferred() types.DecCoin {
	if m != nil {
		return m.Transferred
	}
	return types.DecCoin{}
}

func (m *Account) GetSettledAt() int64 {
	if m != nil {
		return m.SettledAt
	}
	return 0
}

func (m *Account) GetDepositor() string {
	if m != nil {
		return m.Depositor
	}
	return ""
}

func (m *Account) GetFunds() types.DecCoin {
	if m != nil {
		return m.Funds
	}
	return types.DecCoin{}
}

// Payment stores state for a payment
type FractionalPayment struct {
	AccountID AccountID               `protobuf:"bytes,1,opt,name=account_id,json=accountId,proto3" json:"accountID" yaml:"accountID"`
	PaymentID string                  `protobuf:"bytes,2,opt,name=payment_id,json=paymentId,proto3" json:"paymentID" yaml:"paymentID"`
	Owner     string                  `protobuf:"bytes,3,opt,name=owner,proto3" json:"owner" yaml:"owner"`
	State     FractionalPayment_State `protobuf:"varint,4,opt,name=state,proto3,enum=akash.escrow.v1beta2.FractionalPayment_State" json:"state" yaml:"state"`
	Rate      types.DecCoin           `protobuf:"bytes,5,opt,name=rate,proto3" json:"rate" yaml:"rate"`
	Balance   types.DecCoin           `protobuf:"bytes,6,opt,name=balance,proto3" json:"balance" yaml:"balance"`
	Withdrawn types.Coin              `protobuf:"bytes,7,opt,name=withdrawn,proto3" json:"withdrawn" yaml:"withdrawn"`
}

func (m *FractionalPayment) Reset()         { *m = FractionalPayment{} }
func (m *FractionalPayment) String() string { return proto.CompactTextString(m) }
func (*FractionalPayment) ProtoMessage()    {}
func (*FractionalPayment) Descriptor() ([]byte, []int) {
	return fileDescriptor_5b25ee303c78038d, []int{2}
}
func (m *FractionalPayment) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *FractionalPayment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_FractionalPayment.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *FractionalPayment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FractionalPayment.Merge(m, src)
}
func (m *FractionalPayment) XXX_Size() int {
	return m.Size()
}
func (m *FractionalPayment) XXX_DiscardUnknown() {
	xxx_messageInfo_FractionalPayment.DiscardUnknown(m)
}

var xxx_messageInfo_FractionalPayment proto.InternalMessageInfo

func (m *FractionalPayment) GetAccountID() AccountID {
	if m != nil {
		return m.AccountID
	}
	return AccountID{}
}

func (m *FractionalPayment) GetPaymentID() string {
	if m != nil {
		return m.PaymentID
	}
	return ""
}

func (m *FractionalPayment) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *FractionalPayment) GetState() FractionalPayment_State {
	if m != nil {
		return m.State
	}
	return PaymentStateInvalid
}

func (m *FractionalPayment) GetRate() types.DecCoin {
	if m != nil {
		return m.Rate
	}
	return types.DecCoin{}
}

func (m *FractionalPayment) GetBalance() types.DecCoin {
	if m != nil {
		return m.Balance
	}
	return types.DecCoin{}
}

func (m *FractionalPayment) GetWithdrawn() types.Coin {
	if m != nil {
		return m.Withdrawn
	}
	return types.Coin{}
}

func init() {
	proto.RegisterEnum("akash.escrow.v1beta2.Account_State", Account_State_name, Account_State_value)
	proto.RegisterEnum("akash.escrow.v1beta2.FractionalPayment_State", FractionalPayment_State_name, FractionalPayment_State_value)
	proto.RegisterType((*AccountID)(nil), "akash.escrow.v1beta2.AccountID")
	proto.RegisterType((*Account)(nil), "akash.escrow.v1beta2.Account")
	proto.RegisterType((*FractionalPayment)(nil), "akash.escrow.v1beta2.FractionalPayment")
}

func init() { proto.RegisterFile("akash/escrow/v1beta2/types.proto", fileDescriptor_5b25ee303c78038d) }

var fileDescriptor_5b25ee303c78038d = []byte{
	// 810 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x55, 0xcd, 0x6e, 0xeb, 0x44,
	0x14, 0x8e, 0xf3, 0x8b, 0x27, 0x70, 0xc9, 0x1d, 0x2a, 0xe1, 0x86, 0x5b, 0x8f, 0xef, 0x5c, 0x90,
	0xca, 0xe2, 0xda, 0x6a, 0x60, 0xd5, 0x0d, 0x6a, 0x5a, 0x21, 0x75, 0xd1, 0x02, 0x2e, 0x0b, 0x60,
	0xd1, 0x6a, 0x62, 0x4f, 0x5b, 0xab, 0x89, 0xc7, 0xb2, 0xa7, 0x4d, 0xfb, 0x06, 0x28, 0x62, 0x81,
	0x58, 0xb1, 0x09, 0x42, 0xe2, 0x65, 0xba, 0xec, 0x92, 0x95, 0x85, 0xd2, 0x5d, 0x96, 0x79, 0x02,
	0xe4, 0x99, 0xb1, 0x1d, 0x44, 0x14, 0x65, 0x71, 0x57, 0xf6, 0xf9, 0xce, 0x39, 0xdf, 0x9c, 0x9f,
	0xf9, 0x34, 0xc0, 0x22, 0x37, 0x24, 0xb9, 0x76, 0x68, 0xe2, 0xc5, 0x6c, 0xec, 0xdc, 0xed, 0x0d,
	0x28, 0x27, 0x3d, 0x87, 0x3f, 0x44, 0x34, 0xb1, 0xa3, 0x98, 0x71, 0x06, 0xb7, 0x44, 0x84, 0x2d,
	0x23, 0x6c, 0x15, 0xd1, 0xdd, 0xba, 0x62, 0x57, 0x4c, 0x04, 0x38, 0xd9, 0x9f, 0x8c, 0xed, 0x9a,
	0x1e, 0x4b, 0x46, 0x2c, 0x71, 0x06, 0x24, 0xa1, 0x8a, 0x6c, 0xcf, 0xf1, 0x58, 0x10, 0x4a, 0x3f,
	0x1e, 0x02, 0xfd, 0xc0, 0xf3, 0xd8, 0x6d, 0xc8, 0x8f, 0x8f, 0xa0, 0x03, 0x1a, 0x89, 0xc7, 0x22,
	0x6a, 0x68, 0x96, 0xb6, 0xab, 0xf7, 0xb7, 0xe7, 0x29, 0x92, 0xc0, 0x22, 0x45, 0xef, 0x3f, 0x90,
	0xd1, 0x70, 0x1f, 0x0b, 0x13, 0xbb, 0x12, 0x86, 0x36, 0xa8, 0xdd, 0x07, 0xbe, 0x51, 0x15, 0xe1,
	0xaf, 0x66, 0x29, 0xaa, 0xfd, 0x70, 0x7c, 0x34, 0x4f, 0x51, 0x86, 0x2e, 0x52, 0x04, 0x64, 0xce,
	0x7d, 0xe0, 0x63, 0x37, 0x83, 0xf0, 0x1f, 0x4d, 0xd0, 0x52, 0xc7, 0xc1, 0x53, 0x50, 0x0d, 0x7c,
	0x71, 0x52, 0xbb, 0x87, 0xec, 0x55, 0x2d, 0xd9, 0x45, 0x65, 0xfd, 0x9d, 0xc7, 0x14, 0x55, 0x66,
	0x29, 0xaa, 0x0a, 0xfa, 0xaa, 0x60, 0xd7, 0x25, 0x7b, 0x46, 0x5e, 0x0d, 0xfc, 0xac, 0x78, 0x36,
	0x0e, 0x69, 0xac, 0xaa, 0x11, 0xc5, 0x0b, 0xa0, 0x2c, 0x5e, 0x98, 0xd8, 0x95, 0x30, 0xfc, 0x1e,
	0x34, 0x12, 0x4e, 0x38, 0x35, 0x6a, 0x96, 0xb6, 0xfb, 0xa2, 0xf7, 0x66, 0x6d, 0x0d, 0xf6, 0x59,
	0x16, 0xaa, 0x46, 0x92, 0xfd, 0x2e, 0x8d, 0x24, 0x33, 0xb3, 0x91, 0x64, 0x5f, 0xf8, 0x23, 0x68,
	0x0d, 0xc8, 0x90, 0x84, 0x1e, 0x35, 0xea, 0xa2, 0xb7, 0x57, 0xb6, 0x5c, 0x81, 0x9d, 0xad, 0x40,
	0xd1, 0xee, 0xd9, 0x47, 0xd4, 0x3b, 0x64, 0x41, 0xd8, 0x7f, 0x9d, 0x35, 0x36, 0x4f, 0x51, 0x9e,
	0xb4, 0x48, 0xd1, 0x0b, 0x49, 0xab, 0x00, 0xec, 0xe6, 0x2e, 0x18, 0x80, 0x36, 0x8f, 0x49, 0x98,
	0x5c, 0xd2, 0x38, 0xa6, 0xbe, 0xd1, 0xd8, 0x80, 0xfe, 0x73, 0x45, 0xbf, 0x9c, 0xb8, 0x48, 0x11,
	0x94, 0x47, 0x2c, 0x81, 0xd8, 0x5d, 0x0e, 0x81, 0x27, 0x00, 0x24, 0x94, 0xf3, 0x21, 0xf5, 0x2f,
	0x08, 0x37, 0x9a, 0x96, 0xb6, 0x5b, 0xeb, 0xdb, 0xb3, 0x14, 0xe9, 0x67, 0x12, 0x3d, 0xe0, 0xf3,
	0x14, 0xe9, 0x49, 0x6e, 0x2c, 0x52, 0xd4, 0x51, 0xc3, 0xc8, 0x21, 0xec, 0x96, 0x6e, 0xf8, 0x15,
	0xd0, 0x7d, 0x1a, 0xb1, 0x24, 0xe0, 0x2c, 0x36, 0x5a, 0x62, 0x3f, 0xaf, 0x33, 0x82, 0x02, 0x2c,
	0x09, 0x0a, 0x08, 0xbb, 0xa5, 0x1b, 0x7e, 0x07, 0x1a, 0x97, 0xb7, 0xa1, 0x9f, 0x18, 0xef, 0x6d,
	0xd0, 0xf4, 0x8e, 0x6a, 0x5a, 0xa6, 0x94, 0x8b, 0x12, 0x26, 0x76, 0x25, 0x8c, 0x7f, 0xd3, 0x40,
	0x43, 0x2c, 0x15, 0x7e, 0x0a, 0x5a, 0x41, 0x78, 0x47, 0x86, 0x81, 0xdf, 0xa9, 0x74, 0x3f, 0x9e,
	0x4c, 0xad, 0x8f, 0xd4, 0xd2, 0x85, 0xfb, 0x58, 0xba, 0xe0, 0x36, 0xa8, 0xb3, 0x88, 0x86, 0x1d,
	0xad, 0xfb, 0xe1, 0x64, 0x6a, 0xb5, 0x55, 0xc8, 0x37, 0x11, 0x0d, 0xe1, 0x0e, 0x68, 0x7a, 0x43,
	0x96, 0x50, 0xbf, 0x53, 0xed, 0xbe, 0x9c, 0x4c, 0xad, 0x0f, 0x94, 0xf3, 0x50, 0x80, 0xf0, 0x0d,
	0xd0, 0xd9, 0x1d, 0x8d, 0xfd, 0x98, 0x8c, 0xc3, 0x4e, 0xad, 0xbb, 0x35, 0x99, 0x5a, 0x9d, 0x3c,
	0x3d, 0xc7, 0xbb, 0xf5, 0x9f, 0xff, 0x32, 0x2b, 0xf8, 0x97, 0x26, 0x78, 0xf9, 0x75, 0x4c, 0x3c,
	0x1e, 0xb0, 0x90, 0x0c, 0xbf, 0x25, 0x0f, 0x23, 0x1a, 0x72, 0x18, 0x03, 0x40, 0x64, 0xfc, 0xc5,
	0xe6, 0x92, 0xe9, 0x29, 0xc9, 0x94, 0xfa, 0xce, 0x26, 0x4e, 0x72, 0xa3, 0x9c, 0x78, 0x01, 0x61,
	0xb7, 0x70, 0x8b, 0x1b, 0x10, 0xc9, 0xe3, 0x2f, 0x0a, 0x85, 0x8b, 0x1b, 0xa0, 0x8a, 0x92, 0x74,
	0x51, 0x6e, 0x94, 0x74, 0x05, 0x84, 0xdd, 0xc2, 0xbd, 0xa4, 0xce, 0xda, 0x86, 0xea, 0x3c, 0xcf,
	0xd5, 0x59, 0x17, 0xea, 0x7c, 0xbb, 0xba, 0xdd, 0xff, 0xcd, 0x6a, 0x63, 0x9d, 0x9e, 0x80, 0x7a,
	0x9c, 0xd1, 0x6f, 0xa2, 0xa2, 0x4f, 0xd4, 0x85, 0x12, 0x19, 0x8b, 0x14, 0xb5, 0x25, 0x61, 0x2c,
	0xf8, 0x04, 0xb8, 0x2c, 0xfb, 0xe6, 0x3b, 0x96, 0xfd, 0x39, 0xd0, 0xc7, 0x01, 0xbf, 0x16, 0xd7,
	0x44, 0x88, 0xa7, 0xdd, 0xdb, 0x5e, 0x49, 0x2e, 0x98, 0x3f, 0x53, 0xcc, 0x65, 0x4e, 0xb9, 0x9a,
	0x02, 0xc2, 0x6e, 0xe9, 0x5e, 0x2b, 0x04, 0x35, 0xd5, 0x75, 0x42, 0x50, 0x21, 0xab, 0x85, 0xa0,
	0x9c, 0x6b, 0x84, 0x90, 0xa7, 0xff, 0x57, 0x08, 0xfb, 0xf5, 0xdf, 0xff, 0x44, 0x5a, 0xff, 0xf4,
	0x71, 0x66, 0x6a, 0x4f, 0x33, 0x53, 0xfb, 0x67, 0x66, 0x6a, 0xbf, 0x3e, 0x9b, 0x95, 0xa7, 0x67,
	0xb3, 0xf2, 0xf7, 0xb3, 0x59, 0xf9, 0xe9, 0xcb, 0xab, 0x80, 0x5f, 0xdf, 0x0e, 0x6c, 0x8f, 0x8d,
	0x1c, 0x71, 0x33, 0xde, 0x86, 0x94, 0x8f, 0x59, 0x7c, 0xe3, 0x84, 0xcc, 0xa7, 0xce, 0x7d, 0xfe,
	0x7e, 0x8a, 0x77, 0x33, 0x7f, 0x45, 0x07, 0x4d, 0xf1, 0xe8, 0x7d, 0xf1, 0x6f, 0x00, 0x00, 0x00,
	0xff, 0xff, 0xa4, 0x8f, 0x3b, 0x8a, 0x64, 0x07, 0x00, 0x00,
}

func (m *AccountID) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AccountID) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AccountID) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.XID) > 0 {
		i -= len(m.XID)
		copy(dAtA[i:], m.XID)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.XID)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Scope) > 0 {
		i -= len(m.Scope)
		copy(dAtA[i:], m.Scope)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Scope)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Account) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Account) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Account) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Funds.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTypes(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x42
	if len(m.Depositor) > 0 {
		i -= len(m.Depositor)
		copy(dAtA[i:], m.Depositor)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Depositor)))
		i--
		dAtA[i] = 0x3a
	}
	if m.SettledAt != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.SettledAt))
		i--
		dAtA[i] = 0x30
	}
	{
		size, err := m.Transferred.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTypes(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	{
		size, err := m.Balance.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTypes(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	if m.State != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.State))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Owner) > 0 {
		i -= len(m.Owner)
		copy(dAtA[i:], m.Owner)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Owner)))
		i--
		dAtA[i] = 0x12
	}
	{
		size, err := m.ID.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTypes(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *FractionalPayment) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *FractionalPayment) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *FractionalPayment) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Withdrawn.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTypes(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x3a
	{
		size, err := m.Balance.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTypes(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x32
	{
		size, err := m.Rate.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTypes(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	if m.State != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.State))
		i--
		dAtA[i] = 0x20
	}
	if len(m.Owner) > 0 {
		i -= len(m.Owner)
		copy(dAtA[i:], m.Owner)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Owner)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.PaymentID) > 0 {
		i -= len(m.PaymentID)
		copy(dAtA[i:], m.PaymentID)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.PaymentID)))
		i--
		dAtA[i] = 0x12
	}
	{
		size, err := m.AccountID.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTypes(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintTypes(dAtA []byte, offset int, v uint64) int {
	offset -= sovTypes(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *AccountID) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Scope)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = len(m.XID)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}

func (m *Account) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.ID.Size()
	n += 1 + l + sovTypes(uint64(l))
	l = len(m.Owner)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	if m.State != 0 {
		n += 1 + sovTypes(uint64(m.State))
	}
	l = m.Balance.Size()
	n += 1 + l + sovTypes(uint64(l))
	l = m.Transferred.Size()
	n += 1 + l + sovTypes(uint64(l))
	if m.SettledAt != 0 {
		n += 1 + sovTypes(uint64(m.SettledAt))
	}
	l = len(m.Depositor)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = m.Funds.Size()
	n += 1 + l + sovTypes(uint64(l))
	return n
}

func (m *FractionalPayment) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.AccountID.Size()
	n += 1 + l + sovTypes(uint64(l))
	l = len(m.PaymentID)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = len(m.Owner)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	if m.State != 0 {
		n += 1 + sovTypes(uint64(m.State))
	}
	l = m.Rate.Size()
	n += 1 + l + sovTypes(uint64(l))
	l = m.Balance.Size()
	n += 1 + l + sovTypes(uint64(l))
	l = m.Withdrawn.Size()
	n += 1 + l + sovTypes(uint64(l))
	return n
}

func sovTypes(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTypes(x uint64) (n int) {
	return sovTypes(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *AccountID) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: AccountID: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AccountID: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Scope", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Scope = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field XID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.XID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Account) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Account: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Account: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Owner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field State", wireType)
			}
			m.State = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.State |= Account_State(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Balance", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Balance.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Transferred", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Transferred.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SettledAt", wireType)
			}
			m.SettledAt = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SettledAt |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Depositor", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Depositor = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Funds", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Funds.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *FractionalPayment) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: FractionalPayment: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: FractionalPayment: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AccountID", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.AccountID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PaymentID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PaymentID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Owner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field State", wireType)
			}
			m.State = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.State |= FractionalPayment_State(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Rate", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Rate.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Balance", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Balance.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Withdrawn", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Withdrawn.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipTypes(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTypes
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthTypes
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTypes
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTypes
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTypes        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTypes          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTypes = fmt.Errorf("proto: unexpected end of group")
)
