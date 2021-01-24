package hackerrank

/*
* url : https://www.hackerrank.com/challenges/minimum-swaps-2/problem?h_l=interview&playlist_slugs%5B%5D=interview-preparation-kit&playlist_slugs%5B%5D=arrays
* kluczem do rozwiązania jest zamienianie miejscami elementow tablicy dopóki ich index nie bedzie odpowiadał wartości( -1 )
 */
func minimumSwaps(arr []int32) int32 {
	currentIndex := 0
	var swaps int32 = 0
	for currentIndex < len(arr)-1 {
		if int(arr[currentIndex]) != currentIndex+1 {
			arr[currentIndex], arr[arr[currentIndex]-1] = arr[arr[currentIndex]-1], arr[currentIndex]
			swaps++
		} else {
			currentIndex++
		}
	}

	return swaps
}
