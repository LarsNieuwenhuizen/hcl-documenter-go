package variable_files

import (
	"strings"
	"testing"
)

func TestVariablesFileToMarkdownTableCommand(t *testing.T) {
	input := "testdata/variables.tf"
	expected := strings.TrimSpace(`
| Name | Type | Description         | Default |
|------|------|---------------------|---------|
| test | bool | This is just a test | false   |`,
	)

	result, _ := variablesFileToMarkdownTable(input)
	if result != expected {
		t.Errorf("expected: %s, got: \n %s", expected, result)
	}
}
