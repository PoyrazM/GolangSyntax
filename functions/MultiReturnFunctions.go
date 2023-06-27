package functions

func FourTransactions(number1 int, number2 int) (int, int, int, float32) {
	total := number1 + number2
	inference := number1 - number2
	product := number1 * number2
	section := float32(number1 / number2)
	return total, inference, product, section
}
