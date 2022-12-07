package logic

import (
	"context"
	"errors"

	"cloud-disk/core/define"
	"cloud-disk/core/helper"
	"cloud-disk/core/internal/svc"
	"cloud-disk/core/internal/types"
	"cloud-disk/core/models"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserLoginLogic {
	return &UserLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserLoginLogic) UserLogin(req *types.LoginRequest) (resp *types.LoginReply, err error) {
	//1.从数据库查询当前用户
	user := new(models.UserBasic)
	//向数据库查询name,password
	//将请求里的密码进行md5后查询
	has, err := l.svcCtx.Engine.Where("name=? AND password = ?",
		req.Name, helper.Md5(req.Password)).Get(user)
	if err != nil {
		return nil, err
	}
	// 用户不存在
	if !has {
		return nil, errors.New("用户名或密码错误")
	}

	//2.生成token
	//使用数据库里的user.Id, user.Identity, user.Name进行生成
	token, err := helper.GenerateToken(user.Id, user.Identity, user.Name, define.TokenExpire)
	if err != nil {
		return nil, err
	}

	// 3. 生成用于刷新 token 的 token
	refreshToken, err := helper.GenerateToken(user.Id, user.Identity, user.Name, define.RefreshTokenExpire)
	if err != nil {
		return
	}
	resp = new(types.LoginReply)
	resp.Token = token
	resp.RefreshToken = refreshToken

	return

}
