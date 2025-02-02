package user

import (
	"context"
	"github.com/jinzhu/copier"
	"zero-look/app/usercenter/cmd/rpc/usercenter"
	"zero-look/app/usercenter/model"

	"zero-look/app/usercenter/cmd/api/internal/svc"
	"zero-look/app/usercenter/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (*types.LoginResp, error) {
	loginResp, err := l.svcCtx.UserCenterRpc.Login(l.ctx, &usercenter.LoginReq{
		AuthType: model.UserAuthTypeSystem,
		AuthKey:  req.Mobile,
		Password: req.Password,
	})
	if err != nil {
		return nil, err
	}

	resp := &types.LoginResp{}
	err = copier.Copy(resp, loginResp)
	return resp, err
}
