package game

import (
	"github.com/gin-gonic/gin"
	"github.com/programzheng/base/config"
	"github.com/programzheng/base/internal/grpc/call_grpc"
	"github.com/programzheng/base/pkg/controller"
	"github.com/programzheng/base/pkg/resource"
	"github.com/programzheng/base/pkg/service/user"
)

func Play(ctx *gin.Context) {
	token, err := controller.GetTokenByGinContext(ctx)
	if err != nil {
		resource.Unauthorized(ctx, err)
	}

	u, err := user.Auth(&user.UserAuthRequest{
		Token: token,
	})
	if err != nil {
		resource.Unauthorized(ctx, err)
	}

	call_grpc.AssignRandomIssuedTicketToThirdPartyUser(config.Cfg.GetString("AGENT_CODE"), u.UUID)
}
