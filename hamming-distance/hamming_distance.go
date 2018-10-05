package hammingdistance

import (
	"strconv"
)

// HammingDistance The Hamming distance between two integers is the number of positions at which the corresponding bits are different.
func HammingDistance(x int, y int) int {
	num1 := toBinary(x)
	num2 := toBinary(y)

	if len(num1) != len(num2) {
		greater, lesser := getGreaterLesser(num1, num2)
		bits1, bits2 := matchNumsBitsLength(greater, lesser)
		return countDiff(bits1, bits2)
	}

	return countDiff(num1, num2)
}

func toBinary(num int) string {
	return strconv.FormatInt(int64(num), 2)
}

func getGreaterLesser(num1 string, num2 string) (string, string) {
	greater := ""
	lesser := ""

	if len(num1) > len(num2) {
		greater, lesser = num1, num2
	} else {
		greater, lesser = num2, num1
	}

	return greater, lesser
}

func matchNumsBitsLength(greater string, lesser string) (string, string) {
	bits1 := greater
	bits2 := lesser

	diff := len(greater) - len(lesser)

	for i := 0; i < diff; i++ {
		bits2 = string("0") + bits2
	}

	return bits1, bits2
}

func countDiff(num1 string, num2 string) int {
	diff := 0
	for i, num := range num2 {
		x, num1Err := strconv.Atoi(string(num))
		if num1Err != nil {
			return 0
		}
		y, num2Err := strconv.Atoi(string(num1[i]))
		if num2Err != nil {
			return 0
		}
		if x != y {
			diff++
		}
	}
	return diff
}
