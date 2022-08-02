package game

import (
	"fmt"

	"github.com/programzheng/base/config"
)

func VerifyAgentCode(agentCode string) error {
	originAgentCode := config.Cfg.GetString("GAMES_AGENT_CODE")
	if originAgentCode != agentCode {
		return fmt.Errorf("agent code %s is not valid", agentCode)
	}

	return nil
}
