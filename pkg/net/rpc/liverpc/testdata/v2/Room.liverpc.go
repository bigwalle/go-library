// Code generated by protoc-gen-liverpc v0.1, DO NOT EDIT.
// source: v2/Room.proto

/*
Package v2 is a generated liverpc stub package.
This code was generated with go-common/app/tool/liverpc/protoc-gen-liverpc v0.1.

It is generated from these files:
	v2/Room.proto
*/
package v2

import context "context"

import proto "github.com/golang/protobuf/proto"
import "github.com/welcome112s/go-library/net/rpc/liverpc"

var _ proto.Message // generate to suppress unused imports
// Imports only used by utility functions:

// ==============
// Room Interface
// ==============

type RoomRPCClient interface {
	// * 根据房间id获取房间信息v2
	// 修正：原来的get_info_by_id 在传了fields字段但是不包含roomid的情况下 依然会返回所有字段， 新版修正这个问题， 只会返回指定的字段.
	GetByIds(ctx context.Context, req *RoomGetByIdsReq, opts ...liverpc.CallOption) (resp *RoomGetByIdsResp, err error)
}

// ====================
// Room Live Rpc Client
// ====================

type roomRPCClient struct {
	client *liverpc.Client
}

// NewRoomRPCClient creates a client that implements the RoomRPCClient interface.
func NewRoomRPCClient(client *liverpc.Client) RoomRPCClient {
	return &roomRPCClient{
		client: client,
	}
}

func (c *roomRPCClient) GetByIds(ctx context.Context, in *RoomGetByIdsReq, opts ...liverpc.CallOption) (*RoomGetByIdsResp, error) {
	out := new(RoomGetByIdsResp)
	err := doRPCRequest(ctx, c.client, 2, "Room.get_by_ids", in, out, opts)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// =====
// Utils
// =====

func doRPCRequest(ctx context.Context, client *liverpc.Client, version int, method string, in, out proto.Message, opts []liverpc.CallOption) (err error) {
	err = client.Call(ctx, version, method, in, out, opts...)
	return
}
