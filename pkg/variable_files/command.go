package variable_files

import (
	"fmt"
	"github.com/spf13/cobra"
)

func VariablesFileToMarkdownTableCommand() *cobra.Command {
	command := &cobra.Command{
		Use:     "variables-file",
		Short:   "Convert variables file to markdown table",
		Args:    cobra.ExactArgs(1),
		Example: "test",
		RunE: func(cmd *cobra.Command, args []string) error {
			table, err := variablesFileToMarkdownTable(args[0])
			if err != nil {
				fmt.Println(err)
			}

			fmt.Println(table)

			return nil
		},
	}
	return command
}
