/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/programzheng/base/internal/grpc/invokegrpc"
	"github.com/programzheng/base/pkg/helper"
	"github.com/spf13/cobra"
)

// grpcGamesRandomTicketCmd represents the grpcGamesRandomTicket command
var grpcGamesRandomTicketCmd = &cobra.Command{
	Use:   "grpcGamesRandomTicket",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("count is required")
			return
		}
		count := helper.ConvertToInt(args[0])
		if count <= 0 {
			fmt.Println("count is require greater than zero")
			return
		}
		invokegrpc.RandomTicket(count)
	},
}

func init() {
	rootCmd.AddCommand(grpcGamesRandomTicketCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// grpcGamesRandomTicketCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// grpcGamesRandomTicketCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
