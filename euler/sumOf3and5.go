package euler

func isMultiOf5(n int) bool {
	return n%5 == 0
}
func isMultiOf3(n int) bool {
	return n%3 == 0
}

// SumOfMultiesOf3or5 liczy sumę wszystkich liczb ze zbioru <3,1000> dzielących się bez reszty przez 3 bądź 5
func SumOfMultiesOf3or5() int {
	limit := 1000
	sum := 0
	for i := 3; i < limit; i++ {
		if isMultiOf3(i) || isMultiOf5(i) {
			sum += i
		}
	}
	return sum
}
