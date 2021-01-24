package hackerrank

import "fmt"

/*
* problem url : https://www.hackerrank.com/challenges/new-year-chaos/problem?h_l=interview&playlist_slugs%5B%5D=interview-preparation-kit&playlist_slugs%5B%5D=arrays

 */

// rozwiązanie max pkt
// kluczem jest uzycie bubble sort i sprawdzenie czy odległość indeksu od wartości jest większe niż 2 (po dodaniu 1 do indeksu)
// zlozonosc czasowa : O(n2)
//stopień trudności : średni (65% rozwiązań)
func minimumBribes(q []int32) {
	bribes := 0
	l := len(q)

	for i := 0; i < l-1; i++ {
		wasSwapped := false

		for j := 0; j < l-i-1; j++ {

			if int(q[j])-(j+1) > 2 {
				fmt.Println("Too chaotic")
				return
			}
			if q[j] > q[j+1] {
				q[j], q[j+1] = q[j+1], q[j]

				bribes++
				wasSwapped = true
			}
		}

		if !wasSwapped {
			fmt.Println(bribes)

			return
		}

	}
	fmt.Println(bribes)
}
