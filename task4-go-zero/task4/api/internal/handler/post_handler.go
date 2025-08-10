package handler

import (
	"fmt"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"task4-go-zero/internal/types"
	"task4-go-zero/task4/api/internal/logic"
	"task4-go-zero/task4/api/internal/middleware"
	"task4-go-zero/task4/api/internal/svc"
)

func PostCreateHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.PostCreateReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		// 获取用户ID
		userID, ok := r.Context().Value(middleware.UserIDKey).(uint)
		if !ok {
			httpx.ErrorCtx(r.Context(), w, fmt.Errorf("无法获取用户ID"))
			return
		}

		l := logic.NewPostLogic(r.Context(), svcCtx)
		err := l.Create(&req, userID)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, map[string]interface{}{
				"code":    200,
				"message": "创建文章成功",
				"data":    nil,
			})
		}
	}
}

func PostPageHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.PostPageReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		// 获取用户ID
		userID, ok := r.Context().Value(middleware.UserIDKey).(uint)
		if !ok {
			httpx.ErrorCtx(r.Context(), w, fmt.Errorf("无法获取用户ID"))
			return
		}

		l := logic.NewPostLogic(r.Context(), svcCtx)
		resp, err := l.Page(&req, userID)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, map[string]interface{}{
				"code":    200,
				"message": "查询文章成功",
				"data":    resp,
			})
		}
	}
}

func PostDetailHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.PostDetailReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		// 获取用户ID
		userID, ok := r.Context().Value(middleware.UserIDKey).(uint)
		if !ok {
			httpx.ErrorCtx(r.Context(), w, fmt.Errorf("无法获取用户ID"))
			return
		}

		l := logic.NewPostLogic(r.Context(), svcCtx)
		resp, err := l.Detail(&req, userID)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, map[string]interface{}{
				"code":    200,
				"message": "查询文章成功",
				"data":    resp,
			})
		}
	}
}

func PostEditHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.PostEditReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		// 获取用户ID
		userID, ok := r.Context().Value(middleware.UserIDKey).(uint)
		if !ok {
			httpx.ErrorCtx(r.Context(), w, fmt.Errorf("无法获取用户ID"))
			return
		}

		l := logic.NewPostLogic(r.Context(), svcCtx)
		err := l.Edit(&req, userID)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, map[string]interface{}{
				"code":    200,
				"message": "更新文章成功",
				"data":    nil,
			})
		}
	}
}

func PostDeleteHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.PostDeleteReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		// 获取用户ID
		userID, ok := r.Context().Value(middleware.UserIDKey).(uint)
		if !ok {
			httpx.ErrorCtx(r.Context(), w, fmt.Errorf("无法获取用户ID"))
			return
		}

		l := logic.NewPostLogic(r.Context(), svcCtx)
		err := l.Delete(&req, userID)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, map[string]interface{}{
				"code":    200,
				"message": "删除文章成功",
				"data":    nil,
			})
		}
	}
}
