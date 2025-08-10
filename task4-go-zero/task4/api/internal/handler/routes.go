package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest"
	"task4-go-zero/task4/api/internal/middleware"
	"task4-go-zero/task4/api/internal/svc"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	// 创建认证中间件
	authMiddleware := middleware.NewAuthMiddleware(serverCtx.Config.Auth.JwtSecret)

	// 认证相关路由
	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/api/v1/auth/register",
					Handler: UserRegisterHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/api/v1/auth/login",
					Handler: UserLoginHandler(serverCtx),
				},
			}...,
		),
	)

	// 用户相关路由
	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{authMiddleware.Handle},
			[]rest.Route{
				{
					Method:  http.MethodGet,
					Path:    "/api/v1/user/page",
					Handler: UserPageHandler(serverCtx),
				},
			}...,
		),
	)

	// 文章相关路由
	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{authMiddleware.Handle},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/api/v1/post/create",
					Handler: PostCreateHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/api/v1/post/page",
					Handler: PostPageHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/api/v1/post/byId",
					Handler: PostDetailHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/api/v1/post/edit",
					Handler: PostEditHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/api/v1/post/delete",
					Handler: PostDeleteHandler(serverCtx),
				},
			}...,
		),
	)

	// 评论相关路由
	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{authMiddleware.Handle},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/api/v1/comment/create",
					Handler: CommentCreateHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/api/v1/comment/byPostId",
					Handler: CommentListHandler(serverCtx),
				},
			}...,
		),
	)
}
