package hackerrank

/*
* problem url : https://www.hackerrank.com/challenges/ctci-array-left-rotation/problem?h_l=interview&playlist_slugs%5B%5D=interview-preparation-kit&playlist_slugs%5B%5D=arrays
* kluczem do rozwiazania jest sprawdzenie czy rotacje są potrzebne(czy modulo nie jest równe długości tablicy)
* jeśli rotacje są potrzebne to łączymy ze sobą dwie tablice :  jedną od indeksu rownego liczbie rotacji do końca tablicy a drugą  od indeksu 0 do liczby rotacji
 */
func rotLeft(a []int32, d int32) []int32 {
	shifts := int(d) % len(a)
	if shifts == 0 {
		return a
	}

	return append(a[shifts:], a[0:shifts]...)

}
