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
		return
	}

	u, err := user.Auth(&user.UserAuthRequest{
		Token: token,
	})
	if err != nil {
		resource.Unauthorized(ctx, err)
		return
	}

	agentCode := config.Cfg.GetString("GAMES_AGENT_CODE")

	userTicket, err := invokegrpc.AssignOnceRandomIssuedTicketToThirdPartyUser(agentCode, u.UUID)
	if err != nil {
		resource.Fail(ctx, err)
		return
	}

	resource.Success(ctx, userTicket, nil)
}
