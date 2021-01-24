package hackerrank

// url : https://www.hackerrank.com/challenges/equal-stacks/problem
// PASSED 100% TEST CASES

type stack struct {
	height int32
	blocks []int32
}

func (stack *stack) isEmpty() bool {
	return len(stack.blocks) == 0
}

func (stack *stack) shift() {
	if !stack.isEmpty() {
		blockHeight := stack.blocks[0]
		stack.blocks = stack.blocks[1:]
		stack.height -= blockHeight
	}

}
func initStack(blocks []int32) *stack {
	var height int32 = 0
	for _, h := range blocks {
		height += h
	}
	return &stack{height: height, blocks: blocks}
}
func areStacksEqual(s1 *stack, s2 *stack, s3 *stack) bool {
	if s1.height != s2.height || s1.height != s3.height || s2.height != s3.height {
		return false
	}
	return true
}
func getHighestStack(s1 *stack, s2 *stack, s3 *stack) *stack {
	heighest := s1
	if s2.height > heighest.height {
		heighest = s2
	}
	if s3.height > heighest.height {
		heighest = s3
	}
	return heighest
}
func equalStacks(h1 []int32, h2 []int32, h3 []int32) int32 {
	s1 := initStack(h1)
	s2 := initStack(h2)
	s3 := initStack(h3)
	if areStacksEqual(s1, s2, s3) {
		return s1.height
	}
	for !areStacksEqual(s1, s2, s3) {
		heighest := getHighestStack(s1, s2, s3)
		heighest.shift()
	}
	return s1.height
}
