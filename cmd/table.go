package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var tableCmd = &cobra.Command{
	Use:   "table <file>",
	Short: "HTMLテーブルをファイル末尾に追記",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		path := args[0]
		snippet := `<table>
	<caption>HTMLの要素</caption>
	<thead>
	    <tr>
    		<th>名前</th> <th>説明</th>
    	</tr>
  	</thead>
	<tr>
		<td> table </td> <td>テーブル</td>
	</tr>
  	<tr>
    	<td> caption </td> <td>テーブルのキャプション</td>
  	</tr>
</table>
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
	rootCmd.AddCommand(tableCmd)
}
