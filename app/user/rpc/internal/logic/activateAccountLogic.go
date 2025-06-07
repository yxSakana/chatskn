package logic

import (
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/logx"

	"user/model"
	"user/rpc/internal/svc"
	"user/rpc/pb"
)

type ActivateAccountLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewActivateAccountLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ActivateAccountLogic {
	return &ActivateAccountLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ActivateAccountLogic) ActivateAccount(in *pb.ActivateAccountReq) (*pb.ActivateAccountResp, error) {
	info, err := ParseVerifyToken(in.VerifyToken, l.svcCtx.Config.JwtAuth.AccessSecret) // TODO: set secret in config
	if err != nil {
		return nil, err
	}

	user := &model.User{
		Username: "", // TODO: 生成/req
		Email:    info.Email,
		Password: info.Password,
	}
	_, err = l.svcCtx.UserModel.Insert(l.ctx, user)
	if err != nil {
		return nil, fmt.Errorf("user(%s) register failed: %v", info.Email, err)
	}

	_ = l.svcCtx.RegisterCache.Delete(l.ctx, info.Email)
	return &pb.ActivateAccountResp{}, nil
}
