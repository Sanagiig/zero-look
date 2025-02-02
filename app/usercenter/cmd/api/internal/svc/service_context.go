package svc

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
	"zero-look/app/usercenter/cmd/api/internal/config"
	"zero-look/app/usercenter/cmd/rpc/usercenter"
)

type ServiceContext struct {
	Config                config.Config
	UserCenterRpc         usercenter.Usercenter
	SetUidToCtxMiddleware rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:        c,
		UserCenterRpc: usercenter.NewUsercenter(zrpc.MustNewClient(c.UserCenterRpcConf)),
	}
}
