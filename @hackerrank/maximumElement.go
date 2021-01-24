package hackerrank
// PASSED 100% TEST CASES
import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getInput() [][]int {
	scanner := bufio.NewScanner(os.Stdin)
	input := [][]int{}
	for scanner.Scan() {
		strs := strings.Fields(scanner.Text())
		parsed := []int{}
		for _, s := range strs {
			v, _ := strconv.Atoi(s)
			parsed = append(parsed, v)
		}
		input = append(input, parsed)
	}
	return input
}
func getMax(v1 int, v2 int) int {
	if v1 >= v2 {
		return v1
	}
	return v2
}
func main() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	input := getInput()
	stack := []int{}
	for i := 1; i < len(input); i++ {
		query := input[i][0]
		if query == 1 {
			if len(stack) == 0 {
				stack = append(stack, input[i][1])
			} else {
				stack = append(stack, getMax(input[i][1], stack[len(stack)-1]))
			}
		} else if query == 2 {
			stack = stack[:len(stack)-1]
		} else {
			fmt.Println(stack[len(stack)-1])
		}
	}

}
