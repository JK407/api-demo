package svc

import (
	"api-demo/internal/config"
	"api-demo/pb/rpcdemo"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config     config.Config
	RpcDemoCli rpcdemo.RpcDemo // 手动代码
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:     c,
		RpcDemoCli: rpcdemo.NewRpcDemo(zrpc.MustNewClient(c.RpcDemo)), // 手动代码
	}
}
