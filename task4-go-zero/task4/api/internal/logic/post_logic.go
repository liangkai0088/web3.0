package logic

import (
	"context"
	errorx "task4-go-zero/internal/error"
	"task4-go-zero/internal/types"
	"task4-go-zero/task4/api/internal/svc"
)

type PostLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPostLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PostLogic {
	return &PostLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PostLogic) Create(req *types.PostCreateReq, userID uint) error {
	post := types.Post{
		Title:   req.Title,
		Content: req.Content,
		UserID:  userID,
	}

	if err := l.svcCtx.DB.Create(&post).Error; err != nil {
		return errorx.ErrSystem
	}

	return nil
}

func (l *PostLogic) Page(req *types.PostPageReq, userID uint) (resp interface{}, err error) {
	var posts []types.Post
	var total int64

	// 查询总数
	if err := l.svcCtx.DB.Model(&types.Post{}).Where("user_id = ?", userID).Count(&total).Error; err != nil {
		return nil, errorx.ErrSystem
	}

	// 分页查询
	page := req.Page
	pageSize := req.PageSize
	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	offset := (page - 1) * pageSize
	if err := l.svcCtx.DB.Where("user_id = ?", userID).Offset(offset).Limit(pageSize).Find(&posts).Error; err != nil {
		return nil, errorx.ErrSystem
	}

	// 计算总页数
	totalPages := int((total + int64(pageSize) - 1) / int64(pageSize))

	return map[string]interface{}{
		"total":       total,
		"page":        page,
		"pageSize":    pageSize,
		"totalPages":  totalPages,
		"hasNextPage": page < totalPages,
		"data":        posts,
	}, nil
}

func (l *PostLogic) Detail(req *types.PostDetailReq, userID uint) (resp *types.Post, err error) {
	var post types.Post
	if err := l.svcCtx.DB.Where("id = ? AND user_id = ?", req.PostID, userID).First(&post).Error; err != nil {
		return nil, errorx.ErrPostNotFound
	}

	return &post, nil
}

func (l *PostLogic) Edit(req *types.PostEditReq, userID uint) error {
	var post types.Post
	if err := l.svcCtx.DB.Where("id = ? AND user_id = ?", req.ID, userID).First(&post).Error; err != nil {
		return errorx.ErrUnauthorized
	}

	// 更新文章
	if err := l.svcCtx.DB.Model(&post).Updates(types.Post{
		Title:   req.Title,
		Content: req.Content,
	}).Error; err != nil {
		return errorx.ErrSystem
	}

	return nil
}

func (l *PostLogic) Delete(req *types.PostDeleteReq, userID uint) error {
	var post types.Post
	if err := l.svcCtx.DB.Where("id = ? AND user_id = ?", req.PostID, userID).First(&post).Error; err != nil {
		return errorx.ErrUnauthorized
	}

	// 删除文章
	if err := l.svcCtx.DB.Delete(&post).Error; err != nil {
		return errorx.ErrSystem
	}

	return nil
}
