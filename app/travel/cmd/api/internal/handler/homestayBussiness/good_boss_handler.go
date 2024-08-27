package homestayBussiness

import (
	"net/http"

	"github.com/Sanagiig/zero-look/common/result"

	"github.com/zeromicro/go-zero/rest/httpx"
	"zero-look/app/travel/cmd/api/internal/logic/homestayBussiness"
	"zero-look/app/travel/cmd/api/internal/svc"
	"zero-look/app/travel/cmd/api/internal/types"
)

func GoodBossHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GoodBossReq
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := homestayBussiness.NewGoodBossLogic(r.Context(), svcCtx)
		resp, err := l.GoodBoss(&req)
		result.HttpResult(r, w, resp, err)
	}
}
