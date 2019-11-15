package dsears

import "testing"

func TestFuzzyStrings(t *testing.T) {
	tables := []struct {
		input    string
		expected string
	}{
		{"War & Peace", "warpeace"},
		{"Secret of N.I.M.H.", "secretofnimh"},
		{"Amélie", "amelie"},
		{"8½ (1963)", "8x1891963"},
		{"banana", "bernana"},
	}
	for _, table := range tables {
		actual := ToFuzzyString(table.input)
		if actual != table.expected {
			t.Errorf("Expected %s, got %s", table.expected, actual)
		}
	}
}
