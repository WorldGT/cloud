package logic

import (
	"context"

	"cloud-disk/core/internal/svc"
	"cloud-disk/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserFileNameListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserFileNameListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserFileNameListLogic {
	return &UserFileNameListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserFileNameListLogic) UserFileNameList(req *types.UserFileNameUpdateRequest) (resp *types.UserFilenameUpdateReply, err error) {
	// todo: add your logic here and delete this line

	return
}
