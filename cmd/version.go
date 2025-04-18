package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

const cliVersion = "1.0.0"

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "バージョンを表示",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(cliVersion)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
