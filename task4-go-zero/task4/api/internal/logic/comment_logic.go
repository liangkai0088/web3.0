package logic

import (
	"context"
	errorx "task4-go-zero/internal/error"

	"task4-go-zero/internal/types"
	"task4-go-zero/task4/api/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type CommentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CommentLogic {
	return &CommentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CommentLogic) Create(req *types.CommentCreateReq, userID uint) error {
	// 检查文章是否存在
	var post types.Post
	if err := l.svcCtx.DB.Where("id = ?", req.PostID).First(&post).Error; err != nil {
		return errorx.ErrPostNotFound
	}

	// 创建评论
	comment := types.Comment{
		Content: req.Content,
		UserID:  userID,
		PostID:  req.PostID,
	}

	if err := l.svcCtx.DB.Create(&comment).Error; err != nil {
		return errorx.ErrSystem
	}

	return nil
}

func (l *CommentLogic) List(req *types.CommentListReq) (resp []types.Comment, err error) {
	var comments []types.Comment
	if err := l.svcCtx.DB.Where("post_id = ?", req.PostID).Find(&comments).Error; err != nil {
		return nil, errorx.ErrSystem
	}

	return comments, nil
}
