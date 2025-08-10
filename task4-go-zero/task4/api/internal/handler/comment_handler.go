package handler

import (
	"errors"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"task4-go-zero/internal/types"
	"task4-go-zero/task4/api/internal/logic"
	"task4-go-zero/task4/api/internal/middleware"
	"task4-go-zero/task4/api/internal/svc"
)

func CommentCreateHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CommentCreateReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		// 获取用户ID
		userID, ok := r.Context().Value(middleware.UserIDKey).(uint)
		if !ok {
			httpx.ErrorCtx(r.Context(), w, errors.New("无法获取用户ID"))
			return
		}

		l := logic.NewCommentLogic(r.Context(), svcCtx)
		err := l.Create(&req, userID)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, map[string]interface{}{
				"code":    200,
				"message": "创建评论成功",
				"data":    nil,
			})
		}
	}
}

func CommentListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CommentListReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewCommentLogic(r.Context(), svcCtx)
		resp, err := l.List(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, map[string]interface{}{
				"code":    200,
				"message": "查询文章评论成功",
				"data":    resp,
			})
		}
	}
}
