package robotreturntoorigin

import "testing"

func TestJudgeCircle(t *testing.T) {
	tests := []struct {
		move   string
		result bool
	}{
		{move: "UD", result: true},
		{move: "LL", result: false},
	}
	for _, test := range tests {
		actual := judgeCircle(test.move)
		if actual != test.result {
			t.Errorf("It should return %t instead got %t", test.result, actual)
		}
	}
}
