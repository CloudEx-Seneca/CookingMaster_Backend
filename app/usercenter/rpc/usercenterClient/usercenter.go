// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.5
// Source: usercenter.proto

package usercenterClient

import (
	"context"

	"CookingMaster_Backend/app/usercenter/rpc/usercenter"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	GenerateTokenReq  = usercenter.GenerateTokenReq
	GenerateTokenResp = usercenter.GenerateTokenResp
	VarifyTokenReq    = usercenter.VarifyTokenReq
	VarifyTokenResp   = usercenter.VarifyTokenResp

	Usercenter interface {
		GenerateToken(ctx context.Context, in *GenerateTokenReq, opts ...grpc.CallOption) (*GenerateTokenResp, error)
		VarifyToken(ctx context.Context, in *VarifyTokenReq, opts ...grpc.CallOption) (*VarifyTokenResp, error)
	}

	defaultUsercenter struct {
		cli zrpc.Client
	}
)

func NewUsercenter(cli zrpc.Client) Usercenter {
	return &defaultUsercenter{
		cli: cli,
	}
}

func (m *defaultUsercenter) GenerateToken(ctx context.Context, in *GenerateTokenReq, opts ...grpc.CallOption) (*GenerateTokenResp, error) {
	client := usercenter.NewUsercenterClient(m.cli.Conn())
	return client.GenerateToken(ctx, in, opts...)
}

func (m *defaultUsercenter) VarifyToken(ctx context.Context, in *VarifyTokenReq, opts ...grpc.CallOption) (*VarifyTokenResp, error) {
	client := usercenter.NewUsercenterClient(m.cli.Conn())
	return client.VarifyToken(ctx, in, opts...)
}
