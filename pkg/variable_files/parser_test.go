package variable_files

import (
	"errors"
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

func TestVariablesFileToMarkdownTableErrors(t *testing.T) {
	testCases := []struct {
		name          string
		input         string
		expectedError error
	}{
		{
			"non-existing file",
			"test.tf",
			ErrFileNotExist,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := variablesFileToMarkdownTable(tc.input)
			if err == nil || !errors.Is(err, tc.expectedError) {
				t.Errorf("expected error: %v, got: %v", tc.expectedError, err)
			}
		})
	}
}
