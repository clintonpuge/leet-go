package sortarraybyparity

import (
	"reflect"
	"testing"
)

// TestSortArrayByParity test sort number from even to odd
func TestSortArrayByParity(t *testing.T) {
	tests := []struct {
		list   []int
		result []int
	}{
		{list: []int{3, 1, 2, 4}, result: []int{2, 4, 3, 1}},
	}

	for _, test := range tests {
		actual := SortArrayByParity(test.list)
		if !reflect.DeepEqual(actual, test.result) {
			t.Errorf("It should return %v instead got %v", test.result, actual)
		}
	}
}
