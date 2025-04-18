package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var snipCmd = &cobra.Command{
	Use:   "snip [<言語>] <file>",
	Short: "コードスニペットをファイル末尾に追記（言語指定可）",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		var lang, path string

		// 引数が2つ以上ならargs[0]を言語（など）、args[1]をファイルパスとして扱う
		if len(args) == 1 {
			path = args[0]
		} else {
			lang = args[0]
			path = args[1]
		}

		// コードスニペットのテンプレ。langが未指定%sは出力されない
		snippet := fmt.Sprintf("```%s\n\n```\n", strings.TrimSpace(lang))
		f, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			return err
		}
		defer f.Close()
		fmt.Fprint(f, "\n\n"+snippet)
		return nil

		// 書き方2
		/*
			// テンプレートとすでにあるコンテンツを合体させて、ファイル全体を新たな内容で上書きする。
			// os.WriteFileは、ファイル全体を置き換えるのに向いていて、小規模ファイルやテンプレート生成に向いている。
			newContent := []byte(tmpl)
			newContent = append(newContent, content...)

			if err := os.WriteFile(path, newContent, 0o644); err != nil {
				return fmt.Errorf("ファイル書き込みエラー: %w", err)
			}
		*/
	},
}

func init() {
	rootCmd.AddCommand(snipCmd)
}
