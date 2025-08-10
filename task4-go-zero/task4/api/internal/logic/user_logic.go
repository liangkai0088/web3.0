package logic

import (
	"context"
	"golang.org/x/crypto/bcrypt"
	errorx "task4-go-zero/internal/error"
	"task4-go-zero/internal/types"
	"task4-go-zero/pkg/auth"
	"task4-go-zero/task4/api/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserLogic {
	return &UserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserLogic) Register(req *types.UserRegisterReq) error {
	// 检查用户名是否已存在
	var existingUser types.User
	if err := l.svcCtx.DB.Where("username = ?", req.Username).First(&existingUser).Error; err == nil {
		return errorx.ErrInvalidParams
	}

	// 密码加密
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return errorx.ErrSystem
	}

	// 创建用户
	user := types.User{
		Username: req.Username,
		Password: string(hashedPassword),
		Email:    req.Email,
		Role:     "user",
	}

	if err := l.svcCtx.DB.Create(&user).Error; err != nil {
		return errorx.ErrSystem
	}

	return nil
}

func (l *UserLogic) Login(req *types.UserLoginReq) (resp *types.UserLoginResp, err error) {
	// 根据用户名查询用户信息
	var user types.User
	if err := l.svcCtx.DB.Where("username = ?", req.Username).First(&user).Error; err != nil {
		return nil, errorx.ErrUserNotFound
	}

	// 验证密码
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return nil, errorx.ErrInvalidCredentials
	}

	// 生成token
	token, err := auth.GenerateToken(user.ID, user.Role, l.svcCtx.Config.Auth.JwtSecret, l.svcCtx.Config.Auth.TokenExpiry)
	if err != nil {
		return nil, errorx.ErrSystem
	}

	return &types.UserLoginResp{
		Token: token,
	}, nil
}

func (l *UserLogic) Page(req *types.UserPageReq) (resp interface{}, err error) {
	var users []types.User
	var total int64

	// 查询总数
	if err := l.svcCtx.DB.Model(&types.User{}).Count(&total).Error; err != nil {
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
	if err := l.svcCtx.DB.Offset(offset).Limit(pageSize).Find(&users).Error; err != nil {
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
		"data":        users,
	}, nil
}
