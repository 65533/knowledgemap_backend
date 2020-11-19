// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: api/knowledgemap.proto

/*
	Package api is a generated protocol buffer package.

	It is generated from these files:
		api/knowledgemap.proto

	It has these top-level messages:
		Empty
		UserReq
		CRqQueryMapBySubject
		KnowledegeMapInfo
		CRqQueryMyMapBySubject
		CreateKnowledegeReq
		KnowledegeInfoReply
		QueryKnowledegeInfoReq
*/
package api

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Empty struct {
}

func (m *Empty) Reset()                    { *m = Empty{} }
func (m *Empty) String() string            { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()               {}
func (*Empty) Descriptor() ([]byte, []int) { return fileDescriptorKnowledgemap, []int{0} }

type UserReq struct {
	Userid string `protobuf:"bytes,1,opt,name=userid,proto3" json:"uid"`
}

func (m *UserReq) Reset()                    { *m = UserReq{} }
func (m *UserReq) String() string            { return proto.CompactTextString(m) }
func (*UserReq) ProtoMessage()               {}
func (*UserReq) Descriptor() ([]byte, []int) { return fileDescriptorKnowledgemap, []int{1} }

func (m *UserReq) GetUserid() string {
	if m != nil {
		return m.Userid
	}
	return ""
}

type CRqQueryMapBySubject struct {
	Subject string `protobuf:"bytes,1,opt,name=subject,proto3" json:"subject"`
}

func (m *CRqQueryMapBySubject) Reset()                    { *m = CRqQueryMapBySubject{} }
func (m *CRqQueryMapBySubject) String() string            { return proto.CompactTextString(m) }
func (*CRqQueryMapBySubject) ProtoMessage()               {}
func (*CRqQueryMapBySubject) Descriptor() ([]byte, []int) { return fileDescriptorKnowledgemap, []int{2} }

func (m *CRqQueryMapBySubject) GetSubject() string {
	if m != nil {
		return m.Subject
	}
	return ""
}

type KnowledegeMapInfo struct {
	Knowledgemap string `protobuf:"bytes,1,opt,name=knowledgemap,proto3" json:"knowledgemap"`
}

func (m *KnowledegeMapInfo) Reset()                    { *m = KnowledegeMapInfo{} }
func (m *KnowledegeMapInfo) String() string            { return proto.CompactTextString(m) }
func (*KnowledegeMapInfo) ProtoMessage()               {}
func (*KnowledegeMapInfo) Descriptor() ([]byte, []int) { return fileDescriptorKnowledgemap, []int{3} }

func (m *KnowledegeMapInfo) GetKnowledgemap() string {
	if m != nil {
		return m.Knowledgemap
	}
	return ""
}

type CRqQueryMyMapBySubject struct {
	Uid     string `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Subject string `protobuf:"bytes,2,opt,name=subject,proto3" json:"subject"`
	Endtime int64  `protobuf:"varint,3,opt,name=endtime,proto3" json:"endtime"`
}

func (m *CRqQueryMyMapBySubject) Reset()         { *m = CRqQueryMyMapBySubject{} }
func (m *CRqQueryMyMapBySubject) String() string { return proto.CompactTextString(m) }
func (*CRqQueryMyMapBySubject) ProtoMessage()    {}
func (*CRqQueryMyMapBySubject) Descriptor() ([]byte, []int) {
	return fileDescriptorKnowledgemap, []int{4}
}

func (m *CRqQueryMyMapBySubject) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *CRqQueryMyMapBySubject) GetSubject() string {
	if m != nil {
		return m.Subject
	}
	return ""
}

func (m *CRqQueryMyMapBySubject) GetEndtime() int64 {
	if m != nil {
		return m.Endtime
	}
	return 0
}

type CreateKnowledegeReq struct {
	Name    string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty" form:"name"`
	Subject string `protobuf:"bytes,2,opt,name=subject,proto3" json:"subject,omitempty" form:"subject"`
	Course  string `protobuf:"bytes,3,opt,name=course,proto3" json:"course,omitempty" form:"course"`
}

func (m *CreateKnowledegeReq) Reset()                    { *m = CreateKnowledegeReq{} }
func (m *CreateKnowledegeReq) String() string            { return proto.CompactTextString(m) }
func (*CreateKnowledegeReq) ProtoMessage()               {}
func (*CreateKnowledegeReq) Descriptor() ([]byte, []int) { return fileDescriptorKnowledgemap, []int{5} }

func (m *CreateKnowledegeReq) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CreateKnowledegeReq) GetSubject() string {
	if m != nil {
		return m.Subject
	}
	return ""
}

func (m *CreateKnowledegeReq) GetCourse() string {
	if m != nil {
		return m.Course
	}
	return ""
}

type KnowledegeInfoReply struct {
	Id      string `protobuf:"bytes,1,opt,name=id,proto3" json:"id"`
	Name    string `protobuf:"bytes,2,opt,name=name,proto3" json:"name"`
	Subject string `protobuf:"bytes,3,opt,name=subject,proto3" json:"subject"`
	Course  string `protobuf:"bytes,4,opt,name=course,proto3" json:"course"`
}

func (m *KnowledegeInfoReply) Reset()                    { *m = KnowledegeInfoReply{} }
func (m *KnowledegeInfoReply) String() string            { return proto.CompactTextString(m) }
func (*KnowledegeInfoReply) ProtoMessage()               {}
func (*KnowledegeInfoReply) Descriptor() ([]byte, []int) { return fileDescriptorKnowledgemap, []int{6} }

func (m *KnowledegeInfoReply) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *KnowledegeInfoReply) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *KnowledegeInfoReply) GetSubject() string {
	if m != nil {
		return m.Subject
	}
	return ""
}

func (m *KnowledegeInfoReply) GetCourse() string {
	if m != nil {
		return m.Course
	}
	return ""
}

type QueryKnowledegeInfoReq struct {
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id"`
}

func (m *QueryKnowledegeInfoReq) Reset()         { *m = QueryKnowledegeInfoReq{} }
func (m *QueryKnowledegeInfoReq) String() string { return proto.CompactTextString(m) }
func (*QueryKnowledegeInfoReq) ProtoMessage()    {}
func (*QueryKnowledegeInfoReq) Descriptor() ([]byte, []int) {
	return fileDescriptorKnowledgemap, []int{7}
}

func (m *QueryKnowledegeInfoReq) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func init() {
	proto.RegisterType((*Empty)(nil), "api.Empty")
	proto.RegisterType((*UserReq)(nil), "api.UserReq")
	proto.RegisterType((*CRqQueryMapBySubject)(nil), "api.CRqQueryMapBySubject")
	proto.RegisterType((*KnowledegeMapInfo)(nil), "api.KnowledegeMapInfo")
	proto.RegisterType((*CRqQueryMyMapBySubject)(nil), "api.CRqQueryMyMapBySubject")
	proto.RegisterType((*CreateKnowledegeReq)(nil), "api.CreateKnowledegeReq")
	proto.RegisterType((*KnowledegeInfoReply)(nil), "api.KnowledegeInfoReply")
	proto.RegisterType((*QueryKnowledegeInfoReq)(nil), "api.QueryKnowledegeInfoReq")
}
func (m *Empty) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Empty) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	return i, nil
}

func (m *UserReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UserReq) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Userid) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintKnowledgemap(dAtA, i, uint64(len(m.Userid)))
		i += copy(dAtA[i:], m.Userid)
	}
	return i, nil
}

func (m *CRqQueryMapBySubject) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CRqQueryMapBySubject) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Subject) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintKnowledgemap(dAtA, i, uint64(len(m.Subject)))
		i += copy(dAtA[i:], m.Subject)
	}
	return i, nil
}

func (m *KnowledegeMapInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *KnowledegeMapInfo) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Knowledgemap) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintKnowledgemap(dAtA, i, uint64(len(m.Knowledgemap)))
		i += copy(dAtA[i:], m.Knowledgemap)
	}
	return i, nil
}

func (m *CRqQueryMyMapBySubject) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CRqQueryMyMapBySubject) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Uid) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintKnowledgemap(dAtA, i, uint64(len(m.Uid)))
		i += copy(dAtA[i:], m.Uid)
	}
	if len(m.Subject) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintKnowledgemap(dAtA, i, uint64(len(m.Subject)))
		i += copy(dAtA[i:], m.Subject)
	}
	if m.Endtime != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintKnowledgemap(dAtA, i, uint64(m.Endtime))
	}
	return i, nil
}

func (m *CreateKnowledegeReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CreateKnowledegeReq) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Name) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintKnowledgemap(dAtA, i, uint64(len(m.Name)))
		i += copy(dAtA[i:], m.Name)
	}
	if len(m.Subject) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintKnowledgemap(dAtA, i, uint64(len(m.Subject)))
		i += copy(dAtA[i:], m.Subject)
	}
	if len(m.Course) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintKnowledgemap(dAtA, i, uint64(len(m.Course)))
		i += copy(dAtA[i:], m.Course)
	}
	return i, nil
}

func (m *KnowledegeInfoReply) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *KnowledegeInfoReply) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Id) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintKnowledgemap(dAtA, i, uint64(len(m.Id)))
		i += copy(dAtA[i:], m.Id)
	}
	if len(m.Name) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintKnowledgemap(dAtA, i, uint64(len(m.Name)))
		i += copy(dAtA[i:], m.Name)
	}
	if len(m.Subject) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintKnowledgemap(dAtA, i, uint64(len(m.Subject)))
		i += copy(dAtA[i:], m.Subject)
	}
	if len(m.Course) > 0 {
		dAtA[i] = 0x22
		i++
		i = encodeVarintKnowledgemap(dAtA, i, uint64(len(m.Course)))
		i += copy(dAtA[i:], m.Course)
	}
	return i, nil
}

func (m *QueryKnowledegeInfoReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryKnowledegeInfoReq) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Id) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintKnowledgemap(dAtA, i, uint64(len(m.Id)))
		i += copy(dAtA[i:], m.Id)
	}
	return i, nil
}

func encodeVarintKnowledgemap(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Empty) Size() (n int) {
	var l int
	_ = l
	return n
}

func (m *UserReq) Size() (n int) {
	var l int
	_ = l
	l = len(m.Userid)
	if l > 0 {
		n += 1 + l + sovKnowledgemap(uint64(l))
	}
	return n
}

func (m *CRqQueryMapBySubject) Size() (n int) {
	var l int
	_ = l
	l = len(m.Subject)
	if l > 0 {
		n += 1 + l + sovKnowledgemap(uint64(l))
	}
	return n
}

func (m *KnowledegeMapInfo) Size() (n int) {
	var l int
	_ = l
	l = len(m.Knowledgemap)
	if l > 0 {
		n += 1 + l + sovKnowledgemap(uint64(l))
	}
	return n
}

func (m *CRqQueryMyMapBySubject) Size() (n int) {
	var l int
	_ = l
	l = len(m.Uid)
	if l > 0 {
		n += 1 + l + sovKnowledgemap(uint64(l))
	}
	l = len(m.Subject)
	if l > 0 {
		n += 1 + l + sovKnowledgemap(uint64(l))
	}
	if m.Endtime != 0 {
		n += 1 + sovKnowledgemap(uint64(m.Endtime))
	}
	return n
}

func (m *CreateKnowledegeReq) Size() (n int) {
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovKnowledgemap(uint64(l))
	}
	l = len(m.Subject)
	if l > 0 {
		n += 1 + l + sovKnowledgemap(uint64(l))
	}
	l = len(m.Course)
	if l > 0 {
		n += 1 + l + sovKnowledgemap(uint64(l))
	}
	return n
}

func (m *KnowledegeInfoReply) Size() (n int) {
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovKnowledgemap(uint64(l))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovKnowledgemap(uint64(l))
	}
	l = len(m.Subject)
	if l > 0 {
		n += 1 + l + sovKnowledgemap(uint64(l))
	}
	l = len(m.Course)
	if l > 0 {
		n += 1 + l + sovKnowledgemap(uint64(l))
	}
	return n
}

func (m *QueryKnowledegeInfoReq) Size() (n int) {
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovKnowledgemap(uint64(l))
	}
	return n
}

func sovKnowledgemap(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozKnowledgemap(x uint64) (n int) {
	return sovKnowledgemap(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Empty) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowKnowledgemap
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Empty: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Empty: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipKnowledgemap(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthKnowledgemap
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
func (m *UserReq) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowKnowledgemap
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: UserReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UserReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Userid", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowKnowledgemap
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthKnowledgemap
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Userid = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipKnowledgemap(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthKnowledgemap
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
func (m *CRqQueryMapBySubject) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowKnowledgemap
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: CRqQueryMapBySubject: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CRqQueryMapBySubject: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Subject", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowKnowledgemap
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthKnowledgemap
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Subject = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipKnowledgemap(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthKnowledgemap
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
func (m *KnowledegeMapInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowKnowledgemap
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: KnowledegeMapInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: KnowledegeMapInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Knowledgemap", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowKnowledgemap
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthKnowledgemap
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Knowledgemap = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipKnowledgemap(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthKnowledgemap
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
func (m *CRqQueryMyMapBySubject) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowKnowledgemap
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: CRqQueryMyMapBySubject: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CRqQueryMyMapBySubject: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Uid", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowKnowledgemap
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthKnowledgemap
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Uid = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Subject", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowKnowledgemap
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthKnowledgemap
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Subject = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Endtime", wireType)
			}
			m.Endtime = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowKnowledgemap
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Endtime |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipKnowledgemap(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthKnowledgemap
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
func (m *CreateKnowledegeReq) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowKnowledgemap
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: CreateKnowledegeReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CreateKnowledegeReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowKnowledgemap
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthKnowledgemap
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Subject", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowKnowledgemap
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthKnowledgemap
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Subject = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Course", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowKnowledgemap
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthKnowledgemap
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Course = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipKnowledgemap(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthKnowledgemap
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
func (m *KnowledegeInfoReply) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowKnowledgemap
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: KnowledegeInfoReply: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: KnowledegeInfoReply: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowKnowledgemap
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthKnowledgemap
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowKnowledgemap
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthKnowledgemap
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Subject", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowKnowledgemap
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthKnowledgemap
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Subject = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Course", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowKnowledgemap
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthKnowledgemap
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Course = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipKnowledgemap(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthKnowledgemap
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
func (m *QueryKnowledegeInfoReq) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowKnowledgemap
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryKnowledegeInfoReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryKnowledegeInfoReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowKnowledgemap
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthKnowledgemap
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipKnowledgemap(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthKnowledgemap
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
func skipKnowledgemap(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowKnowledgemap
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
					return 0, ErrIntOverflowKnowledgemap
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
					return 0, ErrIntOverflowKnowledgemap
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
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthKnowledgemap
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowKnowledgemap
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
				next, err := skipKnowledgemap(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
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
	ErrInvalidLengthKnowledgemap = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowKnowledgemap   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("api/knowledgemap.proto", fileDescriptorKnowledgemap) }

var fileDescriptorKnowledgemap = []byte{
	// 509 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x94, 0xd1, 0x6e, 0xd3, 0x3e,
	0x14, 0xc6, 0x9b, 0x64, 0xff, 0xf6, 0xbf, 0x33, 0x06, 0x9d, 0x8b, 0xa2, 0xae, 0x43, 0xf5, 0x64,
	0x84, 0x34, 0x10, 0xb4, 0x08, 0xb8, 0x42, 0xe2, 0xa6, 0x13, 0x1a, 0x13, 0x2a, 0x02, 0x03, 0x0f,
	0x90, 0x36, 0x6e, 0x08, 0x2c, 0x8d, 0x9b, 0x26, 0x82, 0x3c, 0x07, 0x12, 0xe2, 0x91, 0xb8, 0x84,
	0x17, 0xb0, 0x50, 0xb9, 0xcb, 0x65, 0x9f, 0x00, 0xc5, 0x76, 0xd3, 0xa6, 0x6b, 0x27, 0xee, 0x8e,
	0x3f, 0x9f, 0x7c, 0xfa, 0xf9, 0x3b, 0x47, 0x01, 0xdb, 0xe1, 0x7e, 0xf7, 0xd3, 0x38, 0xfc, 0x7c,
	0xc1, 0x5c, 0x8f, 0x05, 0x0e, 0xef, 0xf0, 0x28, 0x8c, 0x43, 0x64, 0x39, 0xdc, 0x6f, 0x3d, 0xf0,
	0xfc, 0xf8, 0x43, 0x32, 0xe8, 0x0c, 0xc3, 0xa0, 0xeb, 0x85, 0x5e, 0xd8, 0x95, 0x77, 0x83, 0x64,
	0x24, 0x4f, 0xf2, 0x20, 0x2b, 0xf5, 0x0d, 0xa9, 0xc1, 0x7f, 0xcf, 0x03, 0x1e, 0xa7, 0xe4, 0x1e,
	0xd4, 0xde, 0x4f, 0x59, 0x44, 0xd9, 0x04, 0x61, 0xa8, 0x26, 0x53, 0x16, 0xf9, 0x6e, 0xd3, 0x38,
	0x36, 0x4e, 0x76, 0x7b, 0xb5, 0x4c, 0x60, 0x2b, 0xf1, 0x5d, 0xaa, 0x65, 0xf2, 0x0c, 0x6e, 0x9e,
	0xd2, 0xc9, 0x9b, 0x84, 0x45, 0x69, 0xdf, 0xe1, 0xbd, 0xf4, 0x6d, 0x32, 0xf8, 0xc8, 0x86, 0x31,
	0xba, 0x03, 0xb5, 0xa9, 0x2a, 0xf5, 0x97, 0x7b, 0x99, 0xc0, 0x0b, 0x89, 0x2e, 0x0a, 0x72, 0x0e,
	0x07, 0x2f, 0x15, 0x3d, 0xf3, 0x58, 0xdf, 0xe1, 0xe7, 0xe3, 0x51, 0x88, 0x9e, 0xc0, 0xb5, 0xd5,
	0x27, 0x69, 0x83, 0x7a, 0x26, 0x70, 0x49, 0xa7, 0xa5, 0x13, 0xf9, 0x02, 0x76, 0x41, 0x52, 0x66,
	0xa9, 0x43, 0x8e, 0xac, 0x6c, 0x68, 0x5e, 0xae, 0xd2, 0x99, 0xdb, 0xe9, 0xf2, 0x36, 0x36, 0x76,
	0x63, 0x3f, 0x60, 0x4d, 0xeb, 0xd8, 0x38, 0xb1, 0x54, 0x9b, 0x96, 0xe8, 0xa2, 0x20, 0x5f, 0x0d,
	0x68, 0x9c, 0x46, 0xcc, 0x89, 0xd9, 0xf2, 0x2d, 0x79, 0x78, 0xb7, 0x61, 0x67, 0xec, 0x04, 0x4c,
	0xf3, 0xdf, 0x98, 0x0b, 0xbc, 0x37, 0x0a, 0xa3, 0xe0, 0x29, 0xc9, 0x55, 0x42, 0xe5, 0x25, 0xba,
	0xbf, 0x8e, 0x82, 0xe6, 0x02, 0x5f, 0x57, 0x7d, 0x8b, 0x94, 0x96, 0x44, 0x77, 0xa1, 0x3a, 0x0c,
	0x93, 0x68, 0xaa, 0x80, 0x76, 0x7b, 0x07, 0x73, 0x81, 0xf7, 0x55, 0xb3, 0xd2, 0x09, 0xd5, 0x0d,
	0xe4, 0x9b, 0x01, 0x8d, 0x25, 0x4f, 0x1e, 0x2c, 0x65, 0xfc, 0x22, 0x45, 0x36, 0x98, 0xc5, 0x38,
	0xab, 0x99, 0xc0, 0xa6, 0xef, 0x52, 0xd3, 0x77, 0xd1, 0x2d, 0x4d, 0xab, 0x28, 0xfe, 0xcf, 0x04,
	0x96, 0x67, 0x8d, 0xb9, 0x92, 0x98, 0x75, 0x45, 0x62, 0xa4, 0xe0, 0xdb, 0x91, 0x5d, 0x90, 0x09,
	0xac, 0x95, 0x02, 0xec, 0x21, 0xd8, 0x72, 0x4a, 0xeb, 0x70, 0x93, 0x6d, 0x68, 0x8f, 0x7e, 0x99,
	0xb0, 0x5f, 0x5a, 0x13, 0xf4, 0x1a, 0x0e, 0xcf, 0x58, 0x5c, 0xd2, 0x96, 0xf3, 0x3e, 0xec, 0x38,
	0xdc, 0xef, 0x6c, 0x5a, 0xcb, 0x96, 0x2d, 0xaf, 0x2e, 0xad, 0x1c, 0xa9, 0xa0, 0x77, 0x70, 0x74,
	0xc6, 0xe2, 0x7e, 0xba, 0xc5, 0xf3, 0xa8, 0xec, 0xf9, 0xaf, 0xae, 0x2f, 0xa0, 0xbe, 0xbe, 0x19,
	0xa8, 0xa9, 0xac, 0x2e, 0x2f, 0x4c, 0xab, 0xb9, 0xe6, 0x53, 0x0c, 0x8d, 0x54, 0xd0, 0x2b, 0x68,
	0x6c, 0x48, 0x4d, 0x73, 0x6d, 0xce, 0xf3, 0x2a, 0xbf, 0x5e, 0xfd, 0xc7, 0xac, 0x6d, 0xfc, 0x9c,
	0xb5, 0x8d, 0xdf, 0xb3, 0xb6, 0xf1, 0xfd, 0x4f, 0xbb, 0x32, 0xa8, 0xca, 0xdf, 0xc0, 0xe3, 0xbf,
	0x01, 0x00, 0x00, 0xff, 0xff, 0x02, 0xba, 0x08, 0xca, 0x54, 0x04, 0x00, 0x00,
}