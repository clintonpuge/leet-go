package jewelsandstones

import "testing"

// TestjewelsAndStones test function
func TestJewelsAndStones(t *testing.T) {
	tests := []struct {
		jewel  string
		stones string
		result int
	}{
		{jewel: "aA", stones: "aAAbbbb", result: 3},
		{jewel: "z", stones: "ZZ", result: 0},
	}
	for _, test := range tests {
		actual := jewelsAndStones(test.jewel, test.stones)
		if actual != test.result {
			t.Errorf("It should return %d instead got %d", test.result, actual)
		}
	}
}
