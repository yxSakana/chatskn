package user

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"

	"user/api/internal/svc"
	"user/api/internal/types"
	"user/rpc/user"
)

type ResendActivateEmailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewResendActivateEmailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ResendActivateEmailLogic {
	return &ResendActivateEmailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ResendActivateEmailLogic) ResendActivateEmail(req *types.ResendActivateEmailReq) (resp *types.ResendActivateEmailResp, err error) {
	_, err = l.svcCtx.UserRpc.SendActivateEmail(l.ctx,
		&user.SendActivateEmailReq{Email: req.Email})

	return &types.ResendActivateEmailResp{}, err
}
