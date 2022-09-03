package game

import (
	"github.com/gin-gonic/gin"
	"github.com/programzheng/base/config"
	"github.com/programzheng/base/internal/grpc/invokegrpc"
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

	agentCode := config.Cfg.GetString("GAMES_AGENT_CODE")

	invokegrpc.AssignRandomIssuedTicketToThirdPartyUser(agentCode, u.UUID)
}
