package logic

import (
	"context"

	"cloud-disk/core/internal/svc"
	"cloud-disk/core/internal/types"
	"cloud-disk/core/models"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserFileDeletLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserFileDeletLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserFileDeletLogic {
	return &UserFileDeletLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserFileDeletLogic) UserFileDelet(req *types.UserFileDeleteRequest, userIdentity string) (resp *types.UserFileDeleteReply, err error) {
	// todo: add your logic here and delete this line
	_, err = l.svcCtx.Engine.Where("user_identity =? AND identity = ?", userIdentity, req.Identity).
		Delete(new(models.UserRepository))

	return
}
