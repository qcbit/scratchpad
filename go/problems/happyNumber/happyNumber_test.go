package happynumber

import "testing"

func TestHappyNumber(t *testing.T) {
	tests := []struct {
		n       int
		isHappy bool
	}{
		{19, true},
		{2, false},
		{86, true},
		{1, true},
	}

	for _, tt := range tests {
		got := happyNumber(tt.n)
		if got != tt.isHappy {
			t.Errorf("got %v want %v", got, tt.isHappy)
		}
	}
}
