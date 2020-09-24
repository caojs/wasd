package cmd

import (
	"fmt"

	"github.com/caojs/wasd/cmd/server"

	"github.com/spf13/cobra"
)

var (
	cmd = &cobra.Command{
		Use: "wasd",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("root cmd")
		},
	}
)

func init() {
	cmd.AddCommand(server.Cmd)
}

func Execute() error {
	return cmd.Execute()
}
