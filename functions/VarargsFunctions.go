package functions

func TotalVariadic(numbers ...int) int {
	total := 0

	for i := 0; i < len(numbers); i++ {
		total += numbers[i]
	}
	return total
}
