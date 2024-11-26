package logic

import (
	"context"
	"errors"
	"github.com/jinzhu/copier"
	"goim/apps/user/rpc/internal/svc"
	"goim/apps/user/rpc/user"
	"goim/apps/user/usermodel"

	"github.com/zeromicro/go-zero/core/logx"
)

var ErrUserNotFound = errors.New("该用户不存在")

type GetUserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserInfoLogic) GetUserInfo(in *user.GetUserInfoReq) (*user.GetUserInfoResp, error) {
	userEntiy, err := l.svcCtx.UsersModel.FindOne(l.ctx, in.Id)
	if err != nil {
		if err == usermodel.ErrNotFound {
			return nil, ErrUserNotFound
		}
		return nil, err
	}
	var resp user.UserEntity
	copier.Copy(&resp, userEntiy)

	return &user.GetUserInfoResp{
		User: &resp,
	}, nil
}
