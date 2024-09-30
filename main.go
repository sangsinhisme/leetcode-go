package main

func main() {

	/**
	 * Your CustomStack object will be instantiated and called as such:
	 * obj := Constructor(maxSize);
	 * obj.Push(x);
	 * param_2 := obj.Pop();
	 * obj.Increment(k,val);
	 */

	obj := Constructor(3)
	obj.Push(1)
	obj.Push(2)
	obj.Pop()
	obj.Push(2)
	obj.Push(3)
	obj.Push(4)
	obj.Increment(5, 100)
	obj.Increment(2, 100)
	obj.Pop()
	obj.Pop()
	obj.Pop()
	obj.Pop()
	obj.Pop()

}
