package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"zero/apps/app/api/internal/logic"
	"zero/apps/app/api/internal/svc"
	"zero/apps/app/api/internal/types"
)

// 分类商品列表
func CategoryListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CategoryListRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewCategoryListLogic(r.Context(), svcCtx)
		resp, err := l.CategoryList(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
