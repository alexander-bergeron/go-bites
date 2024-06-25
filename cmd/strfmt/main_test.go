package main

import (
	"testing"
)

func TestFormatString(t *testing.T) {
	tests := []struct {
		template string
		values   map[string]string
		expected string
	}{
		{
			template: "this is a {foo} and {bar}",
			values: map[string]string{
				"foo": "also test",
				"bar": "test",
			},
			expected: "this is a also test and test",
		},
		{
			template: "no placeholders here",
			values:   map[string]string{},
			expected: "no placeholders here",
		},
		{
			template: "missing {foo} value",
			values:   map[string]string{},
			expected: "missing {foo} value",
		},
		{
			template: "extra {foo} placeholder",
			values: map[string]string{
				"foo": "value",
			},
			expected: "extra value placeholder",
		},
		{
			template: "escaped braces {{foo}} should not be replaced",
			values: map[string]string{
				"foo": "value",
			},
			expected: "escaped braces {foo} should not be replaced",
		},
		{
			template: "nested {foo{bar}} placeholders",
			values: map[string]string{
				"foo{bar}": "value",
			},
			expected: "nested value placeholders",
		},
	}

	for _, test := range tests {
		result := formatString(test.template, test.values)
		if result != test.expected {
			t.Errorf("formatString(%q, %v) = %q; want %q", test.template, test.values, result, test.expected)
		}
	}
}
