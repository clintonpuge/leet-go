package hammingdistance

import "testing"

func TestHammingDistance(t *testing.T) {
	tests := []struct {
		x      int
		y      int
		result int
	}{
		{x: 1, y: 4, result: 2},
	}

	for _, test := range tests {
		actual := HammingDistance(test.x, test.y)
		if actual != test.result {
			t.Errorf("It should return %d instead got %d", test.result, actual)
		}
	}
}
