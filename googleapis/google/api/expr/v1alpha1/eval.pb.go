// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: google/api/expr/v1alpha1/eval.proto

package google_api_expr_v1alpha1

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	rpc "istio.io/gogo-genproto/googleapis/google/rpc"
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

// The state of an evaluation.
//
// Can represent an inital, partial, or completed state of evaluation.
type EvalState struct {
	// The unique values referenced in this message.
	Values []*ExprValue `protobuf:"bytes,1,rep,name=values,proto3" json:"values,omitempty"`
	// An ordered list of results.
	//
	// Tracks the flow of evaluation through the expression.
	// May be sparse.
	Results []*EvalState_Result `protobuf:"bytes,3,rep,name=results,proto3" json:"results,omitempty"`
}

func (m *EvalState) Reset()         { *m = EvalState{} }
func (m *EvalState) String() string { return proto.CompactTextString(m) }
func (*EvalState) ProtoMessage()    {}
func (*EvalState) Descriptor() ([]byte, []int) {
	return fileDescriptor_1e95f32326d4b8b7, []int{0}
}
func (m *EvalState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EvalState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EvalState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EvalState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EvalState.Merge(m, src)
}
func (m *EvalState) XXX_Size() int {
	return m.Size()
}
func (m *EvalState) XXX_DiscardUnknown() {
	xxx_messageInfo_EvalState.DiscardUnknown(m)
}

var xxx_messageInfo_EvalState proto.InternalMessageInfo

func (m *EvalState) GetValues() []*ExprValue {
	if m != nil {
		return m.Values
	}
	return nil
}

func (m *EvalState) GetResults() []*EvalState_Result {
	if m != nil {
		return m.Results
	}
	return nil
}

// A single evalution result.
type EvalState_Result struct {
	// The id of the expression this result if for.
	Expr int64 `protobuf:"varint,1,opt,name=expr,proto3" json:"expr,omitempty"`
	// The index in `values` of the resulting value.
	Value int64 `protobuf:"varint,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (m *EvalState_Result) Reset()         { *m = EvalState_Result{} }
func (m *EvalState_Result) String() string { return proto.CompactTextString(m) }
func (*EvalState_Result) ProtoMessage()    {}
func (*EvalState_Result) Descriptor() ([]byte, []int) {
	return fileDescriptor_1e95f32326d4b8b7, []int{0, 0}
}
func (m *EvalState_Result) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EvalState_Result) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EvalState_Result.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EvalState_Result) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EvalState_Result.Merge(m, src)
}
func (m *EvalState_Result) XXX_Size() int {
	return m.Size()
}
func (m *EvalState_Result) XXX_DiscardUnknown() {
	xxx_messageInfo_EvalState_Result.DiscardUnknown(m)
}

var xxx_messageInfo_EvalState_Result proto.InternalMessageInfo

func (m *EvalState_Result) GetExpr() int64 {
	if m != nil {
		return m.Expr
	}
	return 0
}

func (m *EvalState_Result) GetValue() int64 {
	if m != nil {
		return m.Value
	}
	return 0
}

// The value of an evaluated expression.
type ExprValue struct {
	// An expression can resolve to a value, error or unknown.
	//
	// Types that are valid to be assigned to Kind:
	//	*ExprValue_Value
	//	*ExprValue_Error
	//	*ExprValue_Unknown
	Kind isExprValue_Kind `protobuf_oneof:"kind"`
}

func (m *ExprValue) Reset()         { *m = ExprValue{} }
func (m *ExprValue) String() string { return proto.CompactTextString(m) }
func (*ExprValue) ProtoMessage()    {}
func (*ExprValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_1e95f32326d4b8b7, []int{1}
}
func (m *ExprValue) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ExprValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ExprValue.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ExprValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExprValue.Merge(m, src)
}
func (m *ExprValue) XXX_Size() int {
	return m.Size()
}
func (m *ExprValue) XXX_DiscardUnknown() {
	xxx_messageInfo_ExprValue.DiscardUnknown(m)
}

var xxx_messageInfo_ExprValue proto.InternalMessageInfo

type isExprValue_Kind interface {
	isExprValue_Kind()
	MarshalTo([]byte) (int, error)
	Size() int
}

type ExprValue_Value struct {
	Value *Value `protobuf:"bytes,1,opt,name=value,proto3,oneof"`
}
type ExprValue_Error struct {
	Error *ErrorSet `protobuf:"bytes,2,opt,name=error,proto3,oneof"`
}
type ExprValue_Unknown struct {
	Unknown *UnknownSet `protobuf:"bytes,3,opt,name=unknown,proto3,oneof"`
}

func (*ExprValue_Value) isExprValue_Kind()   {}
func (*ExprValue_Error) isExprValue_Kind()   {}
func (*ExprValue_Unknown) isExprValue_Kind() {}

func (m *ExprValue) GetKind() isExprValue_Kind {
	if m != nil {
		return m.Kind
	}
	return nil
}

func (m *ExprValue) GetValue() *Value {
	if x, ok := m.GetKind().(*ExprValue_Value); ok {
		return x.Value
	}
	return nil
}

func (m *ExprValue) GetError() *ErrorSet {
	if x, ok := m.GetKind().(*ExprValue_Error); ok {
		return x.Error
	}
	return nil
}

func (m *ExprValue) GetUnknown() *UnknownSet {
	if x, ok := m.GetKind().(*ExprValue_Unknown); ok {
		return x.Unknown
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*ExprValue) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*ExprValue_Value)(nil),
		(*ExprValue_Error)(nil),
		(*ExprValue_Unknown)(nil),
	}
}

// A set of errors.
//
// The errors included depend on the context. See `ExprValue.error`.
type ErrorSet struct {
	// The errors in the set.
	Errors []*rpc.Status `protobuf:"bytes,1,rep,name=errors,proto3" json:"errors,omitempty"`
}

func (m *ErrorSet) Reset()         { *m = ErrorSet{} }
func (m *ErrorSet) String() string { return proto.CompactTextString(m) }
func (*ErrorSet) ProtoMessage()    {}
func (*ErrorSet) Descriptor() ([]byte, []int) {
	return fileDescriptor_1e95f32326d4b8b7, []int{2}
}
func (m *ErrorSet) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ErrorSet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ErrorSet.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ErrorSet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ErrorSet.Merge(m, src)
}
func (m *ErrorSet) XXX_Size() int {
	return m.Size()
}
func (m *ErrorSet) XXX_DiscardUnknown() {
	xxx_messageInfo_ErrorSet.DiscardUnknown(m)
}

var xxx_messageInfo_ErrorSet proto.InternalMessageInfo

func (m *ErrorSet) GetErrors() []*rpc.Status {
	if m != nil {
		return m.Errors
	}
	return nil
}

// A set of expressions for which the value is unknown.
//
// The unknowns included depend on the context. See `ExprValue.unknown`.
type UnknownSet struct {
	// The ids of the expressions with unknown values.
	Exprs []int64 `protobuf:"varint,1,rep,packed,name=exprs,proto3" json:"exprs,omitempty"`
}

func (m *UnknownSet) Reset()         { *m = UnknownSet{} }
func (m *UnknownSet) String() string { return proto.CompactTextString(m) }
func (*UnknownSet) ProtoMessage()    {}
func (*UnknownSet) Descriptor() ([]byte, []int) {
	return fileDescriptor_1e95f32326d4b8b7, []int{3}
}
func (m *UnknownSet) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *UnknownSet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_UnknownSet.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *UnknownSet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UnknownSet.Merge(m, src)
}
func (m *UnknownSet) XXX_Size() int {
	return m.Size()
}
func (m *UnknownSet) XXX_DiscardUnknown() {
	xxx_messageInfo_UnknownSet.DiscardUnknown(m)
}

var xxx_messageInfo_UnknownSet proto.InternalMessageInfo

func (m *UnknownSet) GetExprs() []int64 {
	if m != nil {
		return m.Exprs
	}
	return nil
}

func init() {
	proto.RegisterType((*EvalState)(nil), "google.api.expr.v1alpha1.EvalState")
	proto.RegisterType((*EvalState_Result)(nil), "google.api.expr.v1alpha1.EvalState.Result")
	proto.RegisterType((*ExprValue)(nil), "google.api.expr.v1alpha1.ExprValue")
	proto.RegisterType((*ErrorSet)(nil), "google.api.expr.v1alpha1.ErrorSet")
	proto.RegisterType((*UnknownSet)(nil), "google.api.expr.v1alpha1.UnknownSet")
}

func init() {
	proto.RegisterFile("google/api/expr/v1alpha1/eval.proto", fileDescriptor_1e95f32326d4b8b7)
}

var fileDescriptor_1e95f32326d4b8b7 = []byte{
	// 367 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0xbf, 0x4e, 0xc3, 0x30,
	0x10, 0x87, 0x63, 0xd2, 0xa6, 0x70, 0xdd, 0x2c, 0x24, 0xa2, 0x0a, 0x85, 0x2a, 0xed, 0x50, 0x75,
	0x70, 0xd4, 0x20, 0x81, 0x04, 0x0b, 0xaa, 0x28, 0x62, 0xac, 0x52, 0xc1, 0x6e, 0x8a, 0x55, 0xaa,
	0x86, 0xc4, 0x72, 0xfe, 0xd0, 0xc7, 0xe0, 0x99, 0x98, 0x10, 0x53, 0x47, 0x46, 0xd4, 0xbe, 0x04,
	0x23, 0xb2, 0x1d, 0x8b, 0x29, 0xdd, 0x62, 0xe7, 0xfb, 0x7e, 0x77, 0xc9, 0x1d, 0xf4, 0x16, 0x69,
	0xba, 0x88, 0x59, 0x40, 0xf9, 0x32, 0x60, 0x6b, 0x2e, 0x82, 0x72, 0x44, 0x63, 0xfe, 0x42, 0x47,
	0x01, 0x2b, 0x69, 0x4c, 0xb8, 0x48, 0xf3, 0x14, 0xbb, 0x1a, 0x22, 0x94, 0x2f, 0x89, 0x84, 0x88,
	0x81, 0x3a, 0xfd, 0x5a, 0xbd, 0xa4, 0x71, 0xc1, 0xb4, 0xdf, 0x39, 0xa9, 0x28, 0xc1, 0xe7, 0x41,
	0x96, 0xd3, 0xbc, 0xc8, 0xf4, 0x0b, 0xff, 0x03, 0xc1, 0xd1, 0xa4, 0xa4, 0xf1, 0x2c, 0xa7, 0x39,
	0xc3, 0xd7, 0xe0, 0x28, 0x2b, 0x73, 0x51, 0xd7, 0x1e, 0xb4, 0xc3, 0x1e, 0xa9, 0xab, 0x4b, 0x26,
	0x6b, 0x2e, 0x1e, 0x25, 0x1b, 0x55, 0x0a, 0xbe, 0x85, 0x96, 0x60, 0x59, 0x11, 0xe7, 0x99, 0x6b,
	0x2b, 0x7b, 0xb8, 0xc7, 0x36, 0x25, 0x49, 0xa4, 0x94, 0xc8, 0xa8, 0x9d, 0x10, 0x1c, 0x7d, 0x85,
	0x31, 0x34, 0xa4, 0xe4, 0xa2, 0x2e, 0x1a, 0xd8, 0x91, 0x7a, 0xc6, 0xc7, 0xd0, 0x54, 0xd5, 0xdc,
	0x03, 0x75, 0xa9, 0x0f, 0xfe, 0x97, 0xfc, 0x08, 0xd3, 0x0f, 0xbe, 0x34, 0x8c, 0x14, 0xdb, 0xe1,
	0x59, 0x7d, 0x17, 0x8a, 0xbf, 0xb7, 0xaa, 0x18, 0x7c, 0x05, 0x4d, 0x26, 0x44, 0x2a, 0x54, 0x78,
	0x3b, 0xf4, 0xf7, 0xb4, 0x2f, 0xb1, 0x19, 0xcb, 0xa5, 0xab, 0x14, 0x7c, 0x03, 0xad, 0x22, 0x59,
	0x25, 0xe9, 0x5b, 0xe2, 0xda, 0xca, 0xee, 0xd7, 0xdb, 0x0f, 0x1a, 0xd4, 0xbe, 0xd1, 0xc6, 0x0e,
	0x34, 0x56, 0xcb, 0xe4, 0xd9, 0xbf, 0x80, 0x43, 0x13, 0x8f, 0x87, 0xe0, 0xa8, 0x78, 0x33, 0x0f,
	0x6c, 0x42, 0x05, 0x9f, 0x93, 0x99, 0x9a, 0x63, 0x54, 0x11, 0xbe, 0x0f, 0xf0, 0x1f, 0x2c, 0x7f,
	0x94, 0x2c, 0xaa, 0x45, 0x3b, 0xd2, 0x87, 0xf1, 0xdd, 0xe7, 0xd6, 0x43, 0x9b, 0xad, 0x87, 0x7e,
	0xb6, 0x1e, 0x7a, 0xdf, 0x79, 0xd6, 0x66, 0xe7, 0x59, 0xdf, 0x3b, 0xcf, 0x82, 0xd3, 0x79, 0xfa,
	0x5a, 0xdb, 0xf1, 0x58, 0xad, 0xc8, 0x54, 0x2e, 0xcc, 0x14, 0xfd, 0x22, 0xf4, 0xe4, 0xa8, 0xe5,
	0x39, 0xff, 0x0b, 0x00, 0x00, 0xff, 0xff, 0xc7, 0x91, 0xca, 0x40, 0xbc, 0x02, 0x00, 0x00,
}

func (m *EvalState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EvalState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EvalState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Results) > 0 {
		for iNdEx := len(m.Results) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Results[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintEval(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.Values) > 0 {
		for iNdEx := len(m.Values) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Values[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintEval(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *EvalState_Result) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EvalState_Result) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EvalState_Result) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Value != 0 {
		i = encodeVarintEval(dAtA, i, uint64(m.Value))
		i--
		dAtA[i] = 0x10
	}
	if m.Expr != 0 {
		i = encodeVarintEval(dAtA, i, uint64(m.Expr))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *ExprValue) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ExprValue) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ExprValue) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Kind != nil {
		{
			size := m.Kind.Size()
			i -= size
			if _, err := m.Kind.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
		}
	}
	return len(dAtA) - i, nil
}

func (m *ExprValue_Value) MarshalTo(dAtA []byte) (int, error) {
	return m.MarshalToSizedBuffer(dAtA[:m.Size()])
}

func (m *ExprValue_Value) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.Value != nil {
		{
			size, err := m.Value.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintEval(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}
func (m *ExprValue_Error) MarshalTo(dAtA []byte) (int, error) {
	return m.MarshalToSizedBuffer(dAtA[:m.Size()])
}

func (m *ExprValue_Error) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.Error != nil {
		{
			size, err := m.Error.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintEval(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	return len(dAtA) - i, nil
}
func (m *ExprValue_Unknown) MarshalTo(dAtA []byte) (int, error) {
	return m.MarshalToSizedBuffer(dAtA[:m.Size()])
}

func (m *ExprValue_Unknown) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.Unknown != nil {
		{
			size, err := m.Unknown.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintEval(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	return len(dAtA) - i, nil
}
func (m *ErrorSet) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ErrorSet) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ErrorSet) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Errors) > 0 {
		for iNdEx := len(m.Errors) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Errors[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintEval(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *UnknownSet) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UnknownSet) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *UnknownSet) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Exprs) > 0 {
		dAtA5 := make([]byte, len(m.Exprs)*10)
		var j4 int
		for _, num1 := range m.Exprs {
			num := uint64(num1)
			for num >= 1<<7 {
				dAtA5[j4] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j4++
			}
			dAtA5[j4] = uint8(num)
			j4++
		}
		i -= j4
		copy(dAtA[i:], dAtA5[:j4])
		i = encodeVarintEval(dAtA, i, uint64(j4))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintEval(dAtA []byte, offset int, v uint64) int {
	offset -= sovEval(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *EvalState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Values) > 0 {
		for _, e := range m.Values {
			l = e.Size()
			n += 1 + l + sovEval(uint64(l))
		}
	}
	if len(m.Results) > 0 {
		for _, e := range m.Results {
			l = e.Size()
			n += 1 + l + sovEval(uint64(l))
		}
	}
	return n
}

func (m *EvalState_Result) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Expr != 0 {
		n += 1 + sovEval(uint64(m.Expr))
	}
	if m.Value != 0 {
		n += 1 + sovEval(uint64(m.Value))
	}
	return n
}

func (m *ExprValue) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Kind != nil {
		n += m.Kind.Size()
	}
	return n
}

func (m *ExprValue_Value) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Value != nil {
		l = m.Value.Size()
		n += 1 + l + sovEval(uint64(l))
	}
	return n
}
func (m *ExprValue_Error) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Error != nil {
		l = m.Error.Size()
		n += 1 + l + sovEval(uint64(l))
	}
	return n
}
func (m *ExprValue_Unknown) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Unknown != nil {
		l = m.Unknown.Size()
		n += 1 + l + sovEval(uint64(l))
	}
	return n
}
func (m *ErrorSet) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Errors) > 0 {
		for _, e := range m.Errors {
			l = e.Size()
			n += 1 + l + sovEval(uint64(l))
		}
	}
	return n
}

func (m *UnknownSet) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Exprs) > 0 {
		l = 0
		for _, e := range m.Exprs {
			l += sovEval(uint64(e))
		}
		n += 1 + sovEval(uint64(l)) + l
	}
	return n
}

func sovEval(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozEval(x uint64) (n int) {
	return sovEval(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *EvalState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEval
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
			return fmt.Errorf("proto: EvalState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EvalState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Values", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEval
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
				return ErrInvalidLengthEval
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthEval
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Values = append(m.Values, &ExprValue{})
			if err := m.Values[len(m.Values)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Results", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEval
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
				return ErrInvalidLengthEval
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthEval
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Results = append(m.Results, &EvalState_Result{})
			if err := m.Results[len(m.Results)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEval(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthEval
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthEval
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
func (m *EvalState_Result) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEval
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
			return fmt.Errorf("proto: Result: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Result: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Expr", wireType)
			}
			m.Expr = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEval
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Expr |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			m.Value = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEval
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Value |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipEval(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthEval
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthEval
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
func (m *ExprValue) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEval
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
			return fmt.Errorf("proto: ExprValue: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ExprValue: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEval
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
				return ErrInvalidLengthEval
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthEval
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &Value{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Kind = &ExprValue_Value{v}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Error", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEval
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
				return ErrInvalidLengthEval
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthEval
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &ErrorSet{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Kind = &ExprValue_Error{v}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Unknown", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEval
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
				return ErrInvalidLengthEval
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthEval
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &UnknownSet{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Kind = &ExprValue_Unknown{v}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEval(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthEval
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthEval
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
func (m *ErrorSet) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEval
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
			return fmt.Errorf("proto: ErrorSet: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ErrorSet: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Errors", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEval
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
				return ErrInvalidLengthEval
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthEval
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Errors = append(m.Errors, &rpc.Status{})
			if err := m.Errors[len(m.Errors)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEval(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthEval
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthEval
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
func (m *UnknownSet) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEval
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
			return fmt.Errorf("proto: UnknownSet: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UnknownSet: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType == 0 {
				var v int64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowEval
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= int64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.Exprs = append(m.Exprs, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowEval
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthEval
				}
				postIndex := iNdEx + packedLen
				if postIndex < 0 {
					return ErrInvalidLengthEval
				}
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				var elementCount int
				var count int
				for _, integer := range dAtA[iNdEx:postIndex] {
					if integer < 128 {
						count++
					}
				}
				elementCount = count
				if elementCount != 0 && len(m.Exprs) == 0 {
					m.Exprs = make([]int64, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v int64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowEval
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= int64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.Exprs = append(m.Exprs, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field Exprs", wireType)
			}
		default:
			iNdEx = preIndex
			skippy, err := skipEval(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthEval
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthEval
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
func skipEval(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowEval
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
					return 0, ErrIntOverflowEval
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowEval
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
				return 0, ErrInvalidLengthEval
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthEval
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowEval
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipEval(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthEval
				}
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthEval = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowEval   = fmt.Errorf("proto: integer overflow")
)
