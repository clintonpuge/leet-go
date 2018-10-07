package flippinganimage

func flipAndInvertImage(A [][]int) [][]int {
	for _, rows := range A {
		reverse(rows)
		flip(rows)
	}
	return A
}

func flip(list []int) {
	for i, num := range list {
		if num == 1 {
			list[i] = 0
		} else {
			list[i] = 1
		}
	}
}

func reverse(list []int) {
	for i, j := 0, len(list)-1; i < j; i, j = i+1, j-1 {
		list[i], list[j] = list[j], list[i]
	}
}
