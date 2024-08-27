package thirdPayment

import (
	"net/http"

	"github.com/Sanagiig/zero-look/common/result"

	"github.com/zeromicro/go-zero/rest/httpx"
	"zero-look/app/payment/cmd/api/internal/logic/thirdPayment"
	"zero-look/app/payment/cmd/api/internal/svc"
	"zero-look/app/payment/cmd/api/internal/types"
)

func ThirdPaymentwxPayHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ThirdPaymentWxPayReq
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := thirdPayment.NewThirdPaymentwxPayLogic(r.Context(), svcCtx)
		resp, err := l.ThirdPaymentwxPay(&req)
		result.HttpResult(r, w, resp, err)
	}
}
