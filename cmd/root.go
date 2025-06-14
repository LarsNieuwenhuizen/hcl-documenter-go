package cmd

import (
	"github.com/larsnieuwenhuizen/hcl-documenter-go/pkg/variable_files"
	"github.com/spf13/cobra"
)

func RootCmd() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "hcldoc",
		Short: "Use this app to document your terraform files",
	}

	rootCmd.AddCommand(variable_files.VariablesFileToMarkdownTableCommand())

	return rootCmd
}
