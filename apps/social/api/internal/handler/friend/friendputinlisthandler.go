package friend

import (
	"goim/apps/social/api/internal/logic/friend"
	"goim/apps/social/api/internal/svc"
	"goim/apps/social/api/internal/types"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func FriendPutInListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.FriendPutInListReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := friend.NewFriendPutInListLogic(r.Context(), svcCtx)
		resp, err := l.FriendPutInList(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}