package homestay

import (
	"net/http"

	"github.com/Sanagiig/zero-look/common/result"

	"github.com/zeromicro/go-zero/rest/httpx"
	"zero-look/app/travel/cmd/api/internal/logic/homestay"
	"zero-look/app/travel/cmd/api/internal/svc"
	"zero-look/app/travel/cmd/api/internal/types"
)

func GuessListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GuessListReq
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := homestay.NewGuessListLogic(r.Context(), svcCtx)
		resp, err := l.GuessList(&req)
		result.HttpResult(r, w, resp, err)
	}
}
