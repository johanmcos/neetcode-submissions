type MinStack struct {
	stack []int
	lowest []int
}

func Constructor() MinStack {
	return MinStack{
		stack: []int{},
		lowest: []int{},
	}
}

func (this *MinStack) Push(val int) {
	this.stack = append(this.stack, val)
	if len(this.lowest) == 0 {
		this.lowest = append(this.lowest, val)
	} else {
		this.lowest = append(this.lowest, min(this.lowest[len(this.lowest)-1], val))
	}
}

func (this *MinStack) Pop() {
	this.stack = this.stack[:len(this.stack)-1]
	this.lowest = this.lowest[:len(this.lowest)-1]
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	return this.lowest[len(this.lowest)-1]
}
