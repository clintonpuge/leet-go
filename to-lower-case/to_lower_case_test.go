package tolowercase

import "testing"

func TestToLowerCase(t *testing.T) {
	tests := []struct {
		input  string
		result string
	}{
		{input: "Hello", result: "hello"},
		{input: "here", result: "here"},
		{input: "LOVELY", result: "lovely"},
	}

	for _, test := range tests {
		actual := ToLowerCase(test.input)
		if actual != test.result {
			t.Errorf("It should return %s instead got %s", test.result, actual)
		}
	}
}
