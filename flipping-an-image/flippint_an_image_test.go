package flippinganimage

import (
	"reflect"
	"testing"
)

func TestFlipAndInvertImage(t *testing.T) {
	tests := []struct {
		input  [][]int
		result [][]int
	}{
		{
			input: [][]int{
				{1, 1, 0},
				{1, 0, 1},
				{0, 0, 0},
			},
			result: [][]int{
				{1, 0, 0},
				{0, 1, 0},
				{1, 1, 1},
			},
		},
		{
			input: [][]int{
				{1, 1, 0, 0},
				{1, 0, 0, 1},
				{0, 1, 1, 1},
				{1, 0, 1, 0},
			},
			result: [][]int{
				{1, 1, 0, 0},
				{0, 1, 1, 0},
				{0, 0, 0, 1},
				{1, 0, 1, 0},
			},
		},
	}
	for _, test := range tests {
		actual := flipAndInvertImage(test.input)
		if !reflect.DeepEqual(actual, test.result) {
			t.Errorf("It should return %v instead got %v", test.result, actual)
		}
	}
}
