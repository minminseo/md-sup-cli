package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var allCmd = &cobra.Command{
	Use:   "all <file>",
	Short: "すべてのテンプレをファイル末尾に追記",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		path := args[0]
		snippets := []string{
			"[タイトル](URL \"\")\n\n---\n",
			`<details><summary>折りたたみタイトル</summary>

</details>

---` + "\n",
			"```\n\n```" + "\n\n---\n",
			`<table>
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
`,
		}
		f, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			return err
		}
		defer f.Close()
		for _, sn := range snippets {
			fmt.Fprint(f, "\n\n"+sn)
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(allCmd)
}
