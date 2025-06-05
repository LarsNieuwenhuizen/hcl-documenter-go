package variable_files

import (
	markdownTable "github.com/larsnieuwenhuizen/go-markdown-table/pkg/table"
	terraform "github.com/larsnieuwenhuizen/hcl-parser-go/pkg/parser/hcl"
	"path/filepath"
)

func parseVariablesFile(filePath string) ([]*terraform.Variable, error) {
	file := []string{filepath.Clean(filePath)}
	config, err := terraform.Parse([]string{}, file)
	if err != nil {
		return nil, err
	}

	return config.Variables, nil
}

func variablesFileToMarkdownTable(filePath string) (string, error) {
	variables, err := parseVariablesFile(filePath)
	if err != nil {
		return "", err
	}

	table := markdownTable.InitiateMarkdownTable()
	table.AddHeaderColumnsFromStringSlice([]string{
		"Name",
		"Type",
		"Description",
		"Default",
	})

	for _, variable := range variables {
		table.AddRowFromStringSlice([]string{
			variable.Name,
			variable.Type,
			variable.Description,
			variable.Default,
		})
	}

	return table.ToString(), nil
}
