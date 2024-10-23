package main

type CustomStack struct {
	s    []int
	size int
	idx  int
}

func ConstructorStack(maxSize int) CustomStack {
	s := make([]int, maxSize)
	size := maxSize
	return CustomStack{s, size, -1}
}

func (s *CustomStack) Push(x int) {
	if s.idx < s.size-1 {
		s.idx++
		s.s[s.idx] = x
	}
}

func (s *CustomStack) Pop() int {
	if s.idx > -1 {
		pop := s.s[s.idx]
		s.s[s.idx] = 0
		s.idx--
		return pop
	}
	return -1
}

func (s *CustomStack) Increment(k int, val int) {
	for i := 0; i < k && i < len(s.s); i++ {
		s.s[i] += val
	}
}

/**
 * Your CustomStack object will be instantiated and called as such:
 * obj := Constructor(maxSize);
 * obj.Push(x);
 * param_2 := obj.Pop();
 * obj.Increment(k,val);
 */
