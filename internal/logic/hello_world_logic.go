package logic

import (
	rpc_demo "api-demo/pb/rpc-demo"
	"context"

	"api-demo/internal/svc"
	"api-demo/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type HelloWorldLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHelloWorldLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HelloWorldLogic {
	return &HelloWorldLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HelloWorldLogic) HelloWorld(req *types.HelloworldReq) (resp *types.HelloworldResp, err error) {
	in := &rpc_demo.HelloRequest{
		Msg: req.Msg,
	}
	r, _ := l.svcCtx.RpcDemoCli.Hello(l.ctx, in)
	resp = &types.HelloworldResp{
		Msg:  r.Msg,
		Code: int(r.Code),
	}
	return
}
