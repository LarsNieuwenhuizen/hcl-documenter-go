package variable_files

import (
	"errors"
	markdownTable "github.com/larsnieuwenhuizen/go-markdown-table/pkg/table"
	terraform "github.com/larsnieuwenhuizen/hcl-parser-go/pkg/parser/hcl"
	"os"
	"path/filepath"
	"strings"
)

var ErrFileNotExist = errors.New("file does not exist")

func parseVariablesFile(filePath string) ([]*terraform.Variable, error) {
	file := []string{filepath.Clean(filePath)}
	if err := doesFileExist(filePath); err != nil {
		return nil, err
	}

	config, err := terraform.Parse([]string{}, file)

	return config.Variables, err
}

func doesFileExist(filePath string) error {
	_, err := os.Stat(filePath)

	if os.IsNotExist(err) {
		return ErrFileNotExist
	}

	return nil
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

	return strings.TrimSpace(table.ToString()), nil
}
