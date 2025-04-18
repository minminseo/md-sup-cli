package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var detailsCmd = &cobra.Command{
	Use:   "details <file>",
	Short: "折りたたみ詳細をファイル末尾に追記",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		path := args[0]
		snippet := `<details><summary>折りたたみタイトル</summary>

</details>
`
		f, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			return err
		}
		defer f.Close()
		fmt.Fprint(f, "\n\n"+snippet)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(detailsCmd)
}
