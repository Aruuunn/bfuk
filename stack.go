package bfuk

type stack []int

func (s stack) Push(v int) stack {
	return append(s, v)
}

func (s stack) Pop() (stack, int) {
	l := len(s)
	return s[:l-1], s[l-1]
}

func (s stack) Top() int {
	if len(s) == 0 {
		return -1
	}

	return s[len(s)-1]
}
