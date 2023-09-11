package user

import (
	"context"
	"fmt"
	"zero/user-api/model"

	"zero/user-api/internal/svc"
	"zero/user-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserInfoLogic) UserInfo(req *types.UserInfoReq) (resp *types.UserInfoResp, err error) {
	user, err := l.svcCtx.UserModel.FindOne(l.ctx, req.ID)
	if err != nil && err != model.ErrNotFound {
		return nil, fmt.Errorf("查询数据失败")
	}
	if user == nil {
		return nil, fmt.Errorf("用户不存在")
	}

	return &types.UserInfoResp{
		ID:       user.Id,
		Nickname: user.Name,
	}, nil
}
