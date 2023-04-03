package domain

import "testing"

func TestMoney_FloatToMoney(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		args     float64
		expected Money
	}{
		{
			name:     "should convert to float to money type",
			args:     100.0,
			expected: 10000,
		},
		{
			name:     "should convert to float to money type",
			args:     .50,
			expected: 50,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if FloatToMoney(tt.args) != tt.expected {
				t.Errorf("[TestCase '%s'] Result: '%v' | Expected: '%v'",
					tt.name,
					FloatToMoney(tt.args),
					tt.expected,
				)
			}
		})
	}
}
