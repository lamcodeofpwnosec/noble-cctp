// Code generated by protoc-gen-go-pulsar. DO NOT EDIT.
package cctpv1

import (
	fmt "fmt"
	runtime "github.com/cosmos/cosmos-proto/runtime"
	_ "github.com/cosmos/gogoproto/gogoproto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoiface "google.golang.org/protobuf/runtime/protoiface"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	io "io"
	reflect "reflect"
	sync "sync"
)

var (
	md_PerMessageBurnLimit        protoreflect.MessageDescriptor
	fd_PerMessageBurnLimit_denom  protoreflect.FieldDescriptor
	fd_PerMessageBurnLimit_amount protoreflect.FieldDescriptor
)

func init() {
	file_circle_cctp_v1_per_message_burn_limit_proto_init()
	md_PerMessageBurnLimit = File_circle_cctp_v1_per_message_burn_limit_proto.Messages().ByName("PerMessageBurnLimit")
	fd_PerMessageBurnLimit_denom = md_PerMessageBurnLimit.Fields().ByName("denom")
	fd_PerMessageBurnLimit_amount = md_PerMessageBurnLimit.Fields().ByName("amount")
}

var _ protoreflect.Message = (*fastReflection_PerMessageBurnLimit)(nil)

type fastReflection_PerMessageBurnLimit PerMessageBurnLimit

func (x *PerMessageBurnLimit) ProtoReflect() protoreflect.Message {
	return (*fastReflection_PerMessageBurnLimit)(x)
}

func (x *PerMessageBurnLimit) slowProtoReflect() protoreflect.Message {
	mi := &file_circle_cctp_v1_per_message_burn_limit_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_PerMessageBurnLimit_messageType fastReflection_PerMessageBurnLimit_messageType
var _ protoreflect.MessageType = fastReflection_PerMessageBurnLimit_messageType{}

type fastReflection_PerMessageBurnLimit_messageType struct{}

func (x fastReflection_PerMessageBurnLimit_messageType) Zero() protoreflect.Message {
	return (*fastReflection_PerMessageBurnLimit)(nil)
}
func (x fastReflection_PerMessageBurnLimit_messageType) New() protoreflect.Message {
	return new(fastReflection_PerMessageBurnLimit)
}
func (x fastReflection_PerMessageBurnLimit_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_PerMessageBurnLimit
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_PerMessageBurnLimit) Descriptor() protoreflect.MessageDescriptor {
	return md_PerMessageBurnLimit
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_PerMessageBurnLimit) Type() protoreflect.MessageType {
	return _fastReflection_PerMessageBurnLimit_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_PerMessageBurnLimit) New() protoreflect.Message {
	return new(fastReflection_PerMessageBurnLimit)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_PerMessageBurnLimit) Interface() protoreflect.ProtoMessage {
	return (*PerMessageBurnLimit)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_PerMessageBurnLimit) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Denom != "" {
		value := protoreflect.ValueOfString(x.Denom)
		if !f(fd_PerMessageBurnLimit_denom, value) {
			return
		}
	}
	if x.Amount != "" {
		value := protoreflect.ValueOfString(x.Amount)
		if !f(fd_PerMessageBurnLimit_amount, value) {
			return
		}
	}
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_PerMessageBurnLimit) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "circle.cctp.v1.PerMessageBurnLimit.denom":
		return x.Denom != ""
	case "circle.cctp.v1.PerMessageBurnLimit.amount":
		return x.Amount != ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: circle.cctp.v1.PerMessageBurnLimit"))
		}
		panic(fmt.Errorf("message circle.cctp.v1.PerMessageBurnLimit does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_PerMessageBurnLimit) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "circle.cctp.v1.PerMessageBurnLimit.denom":
		x.Denom = ""
	case "circle.cctp.v1.PerMessageBurnLimit.amount":
		x.Amount = ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: circle.cctp.v1.PerMessageBurnLimit"))
		}
		panic(fmt.Errorf("message circle.cctp.v1.PerMessageBurnLimit does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_PerMessageBurnLimit) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "circle.cctp.v1.PerMessageBurnLimit.denom":
		value := x.Denom
		return protoreflect.ValueOfString(value)
	case "circle.cctp.v1.PerMessageBurnLimit.amount":
		value := x.Amount
		return protoreflect.ValueOfString(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: circle.cctp.v1.PerMessageBurnLimit"))
		}
		panic(fmt.Errorf("message circle.cctp.v1.PerMessageBurnLimit does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_PerMessageBurnLimit) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "circle.cctp.v1.PerMessageBurnLimit.denom":
		x.Denom = value.Interface().(string)
	case "circle.cctp.v1.PerMessageBurnLimit.amount":
		x.Amount = value.Interface().(string)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: circle.cctp.v1.PerMessageBurnLimit"))
		}
		panic(fmt.Errorf("message circle.cctp.v1.PerMessageBurnLimit does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_PerMessageBurnLimit) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "circle.cctp.v1.PerMessageBurnLimit.denom":
		panic(fmt.Errorf("field denom of message circle.cctp.v1.PerMessageBurnLimit is not mutable"))
	case "circle.cctp.v1.PerMessageBurnLimit.amount":
		panic(fmt.Errorf("field amount of message circle.cctp.v1.PerMessageBurnLimit is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: circle.cctp.v1.PerMessageBurnLimit"))
		}
		panic(fmt.Errorf("message circle.cctp.v1.PerMessageBurnLimit does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_PerMessageBurnLimit) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "circle.cctp.v1.PerMessageBurnLimit.denom":
		return protoreflect.ValueOfString("")
	case "circle.cctp.v1.PerMessageBurnLimit.amount":
		return protoreflect.ValueOfString("")
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: circle.cctp.v1.PerMessageBurnLimit"))
		}
		panic(fmt.Errorf("message circle.cctp.v1.PerMessageBurnLimit does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_PerMessageBurnLimit) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in circle.cctp.v1.PerMessageBurnLimit", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_PerMessageBurnLimit) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_PerMessageBurnLimit) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_PerMessageBurnLimit) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_PerMessageBurnLimit) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*PerMessageBurnLimit)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		l = len(x.Denom)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.Amount)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*PerMessageBurnLimit)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if len(x.Amount) > 0 {
			i -= len(x.Amount)
			copy(dAtA[i:], x.Amount)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Amount)))
			i--
			dAtA[i] = 0x12
		}
		if len(x.Denom) > 0 {
			i -= len(x.Denom)
			copy(dAtA[i:], x.Denom)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Denom)))
			i--
			dAtA[i] = 0xa
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*PerMessageBurnLimit)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: PerMessageBurnLimit: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: PerMessageBurnLimit: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Denom", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Denom = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Amount = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

//
// Copyright (c) 2023, © Circle Internet Financial, LTD.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.0
// 	protoc        (unknown)
// source: circle/cctp/v1/per_message_burn_limit.proto

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// *
// PerMessageBurnLimit is the maximum amount of a certain denom that can be
// burned in an single burn
// @param denom the denom
// @param amount the amount that can be burned (in microunits).  An amount of
// 1000000 uusdc is equivalent to 1USDC
type PerMessageBurnLimit struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Denom  string `protobuf:"bytes,1,opt,name=denom,proto3" json:"denom,omitempty"`
	Amount string `protobuf:"bytes,2,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (x *PerMessageBurnLimit) Reset() {
	*x = PerMessageBurnLimit{}
	if protoimpl.UnsafeEnabled {
		mi := &file_circle_cctp_v1_per_message_burn_limit_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PerMessageBurnLimit) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PerMessageBurnLimit) ProtoMessage() {}

// Deprecated: Use PerMessageBurnLimit.ProtoReflect.Descriptor instead.
func (*PerMessageBurnLimit) Descriptor() ([]byte, []int) {
	return file_circle_cctp_v1_per_message_burn_limit_proto_rawDescGZIP(), []int{0}
}

func (x *PerMessageBurnLimit) GetDenom() string {
	if x != nil {
		return x.Denom
	}
	return ""
}

func (x *PerMessageBurnLimit) GetAmount() string {
	if x != nil {
		return x.Amount
	}
	return ""
}

var File_circle_cctp_v1_per_message_burn_limit_proto protoreflect.FileDescriptor

var file_circle_cctp_v1_per_message_burn_limit_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x63, 0x69, 0x72, 0x63, 0x6c, 0x65, 0x2f, 0x63, 0x63, 0x74, 0x70, 0x2f, 0x76, 0x31,
	0x2f, 0x70, 0x65, 0x72, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x62, 0x75, 0x72,
	0x6e, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x63,
	0x69, 0x72, 0x63, 0x6c, 0x65, 0x2e, 0x63, 0x63, 0x74, 0x70, 0x2e, 0x76, 0x31, 0x1a, 0x14, 0x67,
	0x6f, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x67, 0x6f, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x62, 0x0a, 0x13, 0x50, 0x65, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x42, 0x75, 0x72, 0x6e, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x64, 0x65,
	0x6e, 0x6f, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x64, 0x65, 0x6e, 0x6f, 0x6d,
	0x12, 0x35, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x1d, 0xc8, 0xde, 0x1f, 0x00, 0xda, 0xde, 0x1f, 0x15, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73,
	0x73, 0x64, 0x6b, 0x2e, 0x69, 0x6f, 0x2f, 0x6d, 0x61, 0x74, 0x68, 0x2e, 0x49, 0x6e, 0x74, 0x52,
	0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0xc3, 0x01, 0x0a, 0x12, 0x63, 0x6f, 0x6d, 0x2e,
	0x63, 0x69, 0x72, 0x63, 0x6c, 0x65, 0x2e, 0x63, 0x63, 0x74, 0x70, 0x2e, 0x76, 0x31, 0x42, 0x18,
	0x50, 0x65, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x42, 0x75, 0x72, 0x6e, 0x4c, 0x69,
	0x6d, 0x69, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x39, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x69, 0x72, 0x63, 0x6c, 0x65, 0x66, 0x69, 0x6e,
	0x2f, 0x6e, 0x6f, 0x62, 0x6c, 0x65, 0x2d, 0x63, 0x63, 0x74, 0x70, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x63, 0x69, 0x72, 0x63, 0x6c, 0x65, 0x2f, 0x63, 0x63, 0x74, 0x70, 0x2f, 0x76, 0x31, 0x3b, 0x63,
	0x63, 0x74, 0x70, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x43, 0x43, 0x58, 0xaa, 0x02, 0x0e, 0x43, 0x69,
	0x72, 0x63, 0x6c, 0x65, 0x2e, 0x43, 0x63, 0x74, 0x70, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x0e, 0x43,
	0x69, 0x72, 0x63, 0x6c, 0x65, 0x5c, 0x43, 0x63, 0x74, 0x70, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x1a,
	0x43, 0x69, 0x72, 0x63, 0x6c, 0x65, 0x5c, 0x43, 0x63, 0x74, 0x70, 0x5c, 0x56, 0x31, 0x5c, 0x47,
	0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x10, 0x43, 0x69, 0x72,
	0x63, 0x6c, 0x65, 0x3a, 0x3a, 0x43, 0x63, 0x74, 0x70, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_circle_cctp_v1_per_message_burn_limit_proto_rawDescOnce sync.Once
	file_circle_cctp_v1_per_message_burn_limit_proto_rawDescData = file_circle_cctp_v1_per_message_burn_limit_proto_rawDesc
)

func file_circle_cctp_v1_per_message_burn_limit_proto_rawDescGZIP() []byte {
	file_circle_cctp_v1_per_message_burn_limit_proto_rawDescOnce.Do(func() {
		file_circle_cctp_v1_per_message_burn_limit_proto_rawDescData = protoimpl.X.CompressGZIP(file_circle_cctp_v1_per_message_burn_limit_proto_rawDescData)
	})
	return file_circle_cctp_v1_per_message_burn_limit_proto_rawDescData
}

var file_circle_cctp_v1_per_message_burn_limit_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_circle_cctp_v1_per_message_burn_limit_proto_goTypes = []interface{}{
	(*PerMessageBurnLimit)(nil), // 0: circle.cctp.v1.PerMessageBurnLimit
}
var file_circle_cctp_v1_per_message_burn_limit_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_circle_cctp_v1_per_message_burn_limit_proto_init() }
func file_circle_cctp_v1_per_message_burn_limit_proto_init() {
	if File_circle_cctp_v1_per_message_burn_limit_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_circle_cctp_v1_per_message_burn_limit_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PerMessageBurnLimit); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_circle_cctp_v1_per_message_burn_limit_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_circle_cctp_v1_per_message_burn_limit_proto_goTypes,
		DependencyIndexes: file_circle_cctp_v1_per_message_burn_limit_proto_depIdxs,
		MessageInfos:      file_circle_cctp_v1_per_message_burn_limit_proto_msgTypes,
	}.Build()
	File_circle_cctp_v1_per_message_burn_limit_proto = out.File
	file_circle_cctp_v1_per_message_burn_limit_proto_rawDesc = nil
	file_circle_cctp_v1_per_message_burn_limit_proto_goTypes = nil
	file_circle_cctp_v1_per_message_burn_limit_proto_depIdxs = nil
}
