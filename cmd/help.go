package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var helpText = `USAGE:
msc <COMMAND> [<OPTIONS>]

COMMAND:
	link [<ファイルパス>]             タイトル付きリンクのテンプレ生成
	details [<ファイルパス>]          詳細折りたたみのテンプレ生成
	snip [<言語>] [<ファイルパス>]    コードスニペットを追加
	table [<ファイルパス>]            テーブルのテンプレ生成
	all [<ファイルパス>]              全テンプレ生成
	version                           バージョン表示
	help                              このヘルプを表示
`

var helpCmd = &cobra.Command{
	Use:   "help",
	Short: "ヘルプを表示",
	Run: func(cmd *cobra.Command, args []string) { // 標準出力だけなのでRunEではなくRun
		fmt.Print(helpText)
	},
}

func init() {
	// SetHelpFuncで-hと--helpを上書き
	rootCmd.SetHelpFunc(func(cmd *cobra.Command, args []string) {
		fmt.Print(helpText)
	})

	rootCmd.AddCommand(helpCmd)
}
