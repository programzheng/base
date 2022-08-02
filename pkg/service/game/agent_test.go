package game

import (
	"log"
	"os"
	"testing"
)

func TestVerifyAgentCode(t *testing.T) {
	os.Setenv("GAMES_AGENT_CODE", "test")

	err := VerifyAgentCode("test")
	if err != nil {
		log.Fatal(err)
	}

	log.Println("successfully completed")
}
