package dsears

import "testing"

func TestDaysToShortUnit(t *testing.T) {
	tables := []struct {
		input    int
		expected string
	}{
		{0, "0.0 days"},
		{1, "1.0 days"},
		{7, "1.0 weeks"},
		{14, "2.0 weeks"},
		{45, "1.5 months"},
		{90, "3.0 months"},
		{365, "1.0 years"},
		{500, "1.4 years"},
		{700, "1.9 years"},
	}
	for _, table := range tables {
		actual := DaysToShortUnit(table.input)
		if actual != table.expected {
			t.Errorf("For input %d, expected %s, got %s", table.input, table.expected, actual)
		}
	}
}
