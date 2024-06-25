package main

import (
	"testing"
)

func TestDetermineTemplateType(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected TemplateType
	}{
		{
			name:     "Valid JSON",
			input:    `{"field1": "value1", "field2": 123}`,
			expected: JSONTemplate,
		},
		{
			name:     "Valid YAML Single Line",
			input:    "field1: value1\nfield2: 123",
			expected: YAMLTemplate,
		},
		{
			name:     "Valid YAML Multi Line",
			input:    "field1: value1\nfield2: 123\n",
			expected: YAMLTemplate,
		},
		{
			name:     "Valid TOML",
			input:    "field1 = \"value1\"\nfield2 = 123\n",
			expected: TOMLTemplate,
		},
		{
			name:     "Invalid JSON, YAML, TOML - Should be Text",
			input:    "This is just plain text.",
			expected: TextTemplate,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := determineTemplateType(tt.input)
			if actual != tt.expected {
				t.Errorf("determineTemplateType(%q) = %v, want %v", tt.input, actual, tt.expected)
			}
		})
	}
}
