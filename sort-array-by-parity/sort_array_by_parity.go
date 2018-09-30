package sortarraybyparity

// SortArrayByParity sort list of numbers from even to odd
func SortArrayByParity(nums []int) []int {
	odd := []int{}
	even := []int{}
	for _, num := range nums {
		if num%2 == 0 {
			even = append(even, num)
		} else {
			odd = append(odd, num)
		}
	}
	return append(even, odd...)
}
