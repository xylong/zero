package user

import (
	"context"
	"fmt"
	"zero/user-api/internal/svc"
	"zero/user-api/internal/types"
	"zero/user-api/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserListLogic {
	return &UserListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserListLogic) UserList(req *types.UserListReq) (resp *types.UserListResp, err error) {
	resp = &types.UserListResp{List: make([]types.User, 0)}

	users, err := l.svcCtx.UserModel.Query(l.ctx, req)
	if err != nil && err != model.ErrNotFound {
		return resp, fmt.Errorf("查询数据失败:%s", err.Error())
	}

	for _, user := range users {
		fmt.Println(user.Name)
		resp.List = append(resp.List, types.User{
			Id:     user.Id,
			Name:   user.Name,
			Phone:  user.Phone,
			Email:  user.Email,
			Gender: user.Gender,
			Avatar: user.Avatar,
		})
	}

	return
}
