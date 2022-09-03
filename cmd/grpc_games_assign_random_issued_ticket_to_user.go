/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"

	"github.com/programzheng/base/internal/grpc/invokegrpc"
	"github.com/programzheng/base/pkg/service/game"
	"github.com/programzheng/base/pkg/service/user"
	"github.com/spf13/cobra"
)

// grpcGamesAssignRandomIssuedTicketToUserCmd represents the grpcGamesAssignRandomIssuedTicketToUser command
var grpcGamesAssignRandomIssuedTicketToUserCmd = &cobra.Command{
	Use:   "grpcGamesAssignRandomIssuedTicketToUser",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("agent code is required")
			return
		}
		agentCode := args[0]
		err := game.VerifyAgentCode(agentCode)
		if err != nil {
			log.Fatal(err)
		}
		if len(args) == 1 {
			fmt.Println("user uuid is required")
			return
		}
		userUUID := args[1]
		user, err := user.GetUserByUUID(userUUID)
		if err != nil {
			log.Fatal(err)
		}
		invokegrpc.AssignRandomIssuedTicketToThirdPartyUser(agentCode, user.UUID)
	},
}

func init() {
	rootCmd.AddCommand(grpcGamesAssignRandomIssuedTicketToUserCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// grpcGamesAssignRandomIssuedTicketToUserCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// grpcGamesAssignRandomIssuedTicketToUserCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
