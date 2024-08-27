package homestayComment

import (
	"net/http"

	"github.com/Sanagiig/zero-look/common/result"

	"github.com/zeromicro/go-zero/rest/httpx"
	"zero-look/app/travel/cmd/api/internal/logic/homestayComment"
	"zero-look/app/travel/cmd/api/internal/svc"
	"zero-look/app/travel/cmd/api/internal/types"
)

func CommentListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CommentListReq
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := homestayComment.NewCommentListLogic(r.Context(), svcCtx)
		resp, err := l.CommentList(&req)
		result.HttpResult(r, w, resp, err)
	}
}
