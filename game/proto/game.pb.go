// Code generated by protoc-gen-go.
// source: game/game.proto
// DO NOT EDIT!

/*
Package game is a generated protocol buffer package.

It is generated from these files:
	game/game.proto

It has these top-level messages:
	Row
	Col
*/
package game

import proto "code.google.com/p/goprotobuf/proto"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type Row struct {
	X                []int32 `protobuf:"varint,1,rep,name=x" json:"x,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Row) Reset()         { *m = Row{} }
func (m *Row) String() string { return proto.CompactTextString(m) }
func (*Row) ProtoMessage()    {}

func (m *Row) GetX() []int32 {
	if m != nil {
		return m.X
	}
	return nil
}

type Col struct {
	Rows             []*Row `protobuf:"bytes,2,rep,name=rows" json:"rows,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *Col) Reset()         { *m = Col{} }
func (m *Col) String() string { return proto.CompactTextString(m) }
func (*Col) ProtoMessage()    {}

func (m *Col) GetRows() []*Row {
	if m != nil {
		return m.Rows
	}
	return nil
}

func init() {
}
