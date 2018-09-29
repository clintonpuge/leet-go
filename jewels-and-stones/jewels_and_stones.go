package jewelsandstones

import "strings"

func jewelsAndStones(jewels string, stones string) int {
	count := 0
	for _, stone := range stones {
		if strings.Contains(jewels, string(stone)) {
			count++
		}
	}
	return count
}
