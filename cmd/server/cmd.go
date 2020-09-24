package server

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	Cmd = &cobra.Command{
		Use: "server",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("server cmd")
		},
	}
)
