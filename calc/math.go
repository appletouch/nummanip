package calc

// fixed loop calc
func Add(numbers ...int) int {
	sum := 0
	for num := range numbers {
		sum = +num
	}
	return sum
}
