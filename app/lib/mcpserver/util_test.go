package mcpserver

import (
	"testing"

	"github.com/kyleu/pftest/app/util"
)

func TestExtractURIPatternArgs(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name     string
		pattern  string
		uri      string
		expected util.ValueMap
	}{
		{name: "basic pattern matching", pattern: "resource://{section}/{id}", uri: "resource://x/123", expected: util.ValueMap{"section": "x", "id": "123"}},
		{name: "single variable", pattern: "resource://data/{name}", uri: "resource://data/test", expected: util.ValueMap{"name": "test"}},
		{name: "no variables", pattern: "resource://static/path", uri: "resource://static/path", expected: util.ValueMap{}},
		{name: "pattern doesn't match", pattern: "resource://example/{id}", uri: "resource://different/123", expected: util.ValueMap{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			result := extractURIPatternArgs(tt.pattern, tt.uri)
			if len(result) != len(tt.expected) {
				t.Errorf("Expected %d args, got %d", len(tt.expected), len(result))
				return
			}
			for key, expectedValue := range tt.expected {
				if actualValue, exists := result[key]; !exists {
					t.Errorf("Expected key %s not found in result", key)
				} else if actualValue != expectedValue {
					t.Errorf("Expected %s=%s, got %s=%s", key, expectedValue, key, actualValue)
				}
			}
		})
	}
}
