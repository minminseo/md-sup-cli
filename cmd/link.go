package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var linkCmd = &cobra.Command{
	Use:   "link <file>",
	Short: "タイトル付きリンクをファイル末尾に追記",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		path := args[0]
		snippet := "[タイトル](URL \"\")\n"
		f, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			return err
		}
		defer f.Close()
		// 2行改行してから追記
		fmt.Fprint(f, "\n\n"+snippet)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(linkCmd)
}
