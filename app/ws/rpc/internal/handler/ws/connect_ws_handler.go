package ws

import (
	"chatskn/app/ws/rpc/internal/webs"
	"github.com/zeromicro/go-zero/core/logx"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"chatskn/app/ws/rpc/internal/svc"
	"chatskn/app/ws/rpc/internal/types"
)

// connect
func ConnectWsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.WsConnectReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		webs.ServeWs(webs.ServeInput{
			HubCore: svcCtx.Hub,
			Writer:  w,
			Req:     r,
			UserId:  req.Id,
			Handler: func(msg []byte) error {
				// 是否有必要将发送ID强制改为当前连接的用户ID
				msgString := string(msg)
				logx.WithContext(r.Context()).Infof("[Broadcasts-Pusher]: %s", msgString)
				if err := svcCtx.KqPublisher.Push(r.Context(), msgString); err != nil {
					logx.WithContext(r.Context()).Errorf("[Broadcasts-Pusher]: %s", err.Error())
					return err
				}
				return nil
			},
		})
	}
}
