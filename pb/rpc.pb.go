// Code generated by protoc-gen-gogo.
// source: rpc.proto
// DO NOT EDIT!

/*
Package floodsub_pb is a generated protocol buffer package.

It is generated from these files:
	rpc.proto

It has these top-level messages:
	RPC
	Message
*/
package floodsub_pb

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type RPC struct {
	Type             *string  `protobuf:"bytes,1,opt,name=type" json:"type,omitempty"`
	Topics           []string `protobuf:"bytes,2,rep,name=topics" json:"topics,omitempty"`
	Msg              *Message `protobuf:"bytes,3,opt,name=msg" json:"msg,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *RPC) Reset()         { *m = RPC{} }
func (m *RPC) String() string { return proto.CompactTextString(m) }
func (*RPC) ProtoMessage()    {}

func (m *RPC) GetType() string {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return ""
}

func (m *RPC) GetTopics() []string {
	if m != nil {
		return m.Topics
	}
	return nil
}

func (m *RPC) GetMsg() *Message {
	if m != nil {
		return m.Msg
	}
	return nil
}

type Message struct {
	From             *string `protobuf:"bytes,1,opt,name=from" json:"from,omitempty"`
	Data             []byte  `protobuf:"bytes,2,opt,name=data" json:"data,omitempty"`
	Seqno            *uint64 `protobuf:"varint,3,opt,name=seqno" json:"seqno,omitempty"`
	Topic            *string `protobuf:"bytes,4,opt,name=topic" json:"topic,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}

func (m *Message) GetFrom() string {
	if m != nil && m.From != nil {
		return *m.From
	}
	return ""
}

func (m *Message) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *Message) GetSeqno() uint64 {
	if m != nil && m.Seqno != nil {
		return *m.Seqno
	}
	return 0
}

func (m *Message) GetTopic() string {
	if m != nil && m.Topic != nil {
		return *m.Topic
	}
	return ""
}

func init() {
	proto.RegisterType((*RPC)(nil), "floodsub.pb.RPC")
	proto.RegisterType((*Message)(nil), "floodsub.pb.Message")
}
