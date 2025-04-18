/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "msc <command> [options]",
	Short: "Markdown補助CLI（md-sup-cli）",
	Long: `md-sup-cliはMarkdownの少しだけ面倒な記法のテンプレをファイル末尾に追記するCLIです。
`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
