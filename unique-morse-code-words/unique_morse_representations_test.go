package uniquemorsecodewords

import "testing"

func TestUniqueMorseRepresentations(t *testing.T) {
	tests := []struct {
		words  []string
		result int
	}{
		{words: []string{"gin", "zen", "gig", "msg"}, result: 2},
	}
	for _, test := range tests {
		actual := UniqueMorseRepresentations(test.words)
		if actual != test.result {
			t.Errorf("It should return %d instead got %d", test.result, actual)
		}
	}
}
