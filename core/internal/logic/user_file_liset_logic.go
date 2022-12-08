package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"cloud-disk/core/define"
	"cloud-disk/core/internal/svc"
	"cloud-disk/core/internal/types"
	"cloud-disk/core/models"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserFileLisetLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserFileLisetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserFileLisetLogic {
	return &UserFileLisetLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserFileLisetLogic) UserFileLiset(req *types.UserFileListRequest, userIdentity string) (resp *types.UserFileLisetReply, err error) {
	// todo: add your logic here and delete this line
	uf := make([]*types.UserFile, 0)
	resp = new(types.UserFileLisetReply)

	// 分页参数
	size := req.Size
	if size == 0 {
		size = define.PageSize
	}
	page := req.Page
	if page == 0 {
		page = 1
	}
	offset := (page - 1) * size

	//查询表user_repository 里的数据 但不包括被删除的 和 有父级文件的

	err = l.svcCtx.Engine.Table("user_repository").
		Where("parent_id = ? AND user_identity = ? ", req.Id, userIdentity).
		Select("user_repository.id, user_repository.identity, user_repository.repository_identity, "+
			"user_repository.ext, user_repository.name, repository_pool.path, repository_pool.size").
		Join("LEFT", "repository_pool", "user_repository.repository_identity = repository_pool.identity").
		Where("user_repository.deleted_at = ? OR user_repository.deleted_at IS NULL", time.Time{}.Format(define.Datetime)).
		Limit(size, offset).
		Find(&uf)
	/*
		查询表user_repository 里的 parent_id =  req.Id AND user_identity =  userIdentity 的行
		选择表user_repository里的id，user_repository.identity, user_repository.repository_identity，
			user_repository.ext, user_repository.name, repository_pool.path, repository_pool.size
		user_repository左连接表repository_pool，"user_repository里的repository_identity = repository_pool里的identity"
		user_repository里的repository_identity = repository_pool里的identity"
			查找 user_repository.deleted_at 有时间或空的
		Limit(int, …int)限Limit(int, …int)限制获取的数目，第一个参数为条数，第二个参数表示开始位置，如果不传则为0
		返回结果到uf
	*/
	jsons, _ := json.Marshal(uf)
	fmt.Println("run:")
	fmt.Println("uf:", string(jsons))
	fmt.Println("end:")

	if err != nil {
		return
	}

	// 查询用户文件总数
	cnt, err := l.svcCtx.Engine.Where("parent_id = ? AND user_identity = ? ", req.Id, userIdentity).Count(new(models.UserRepository))
	if err != nil {
		return
	}
	resp.List = uf
	resp.Count = cnt

	return
}
