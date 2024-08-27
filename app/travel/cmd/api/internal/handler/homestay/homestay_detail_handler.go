package homestay

import (
	"net/http"

	"github.com/Sanagiig/zero-look/common/result"

	"github.com/zeromicro/go-zero/rest/httpx"
	"zero-look/app/travel/cmd/api/internal/logic/homestay"
	"zero-look/app/travel/cmd/api/internal/svc"
	"zero-look/app/travel/cmd/api/internal/types"
)

func HomestayDetailHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.HomestayDetailReq
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := homestay.NewHomestayDetailLogic(r.Context(), svcCtx)
		resp, err := l.HomestayDetail(&req)
		result.HttpResult(r, w, resp, err)
	}
}
