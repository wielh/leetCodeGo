package main

type MyQueue struct {
	data []int
}

func Constructor() MyQueue {
	return MyQueue{data: []int{}}
}

func (this *MyQueue) Push(x int) {
	this.data = append(this.data, x)
}

func (this *MyQueue) Pop() int {
	answer := this.Peek()
	this.data = this.data[1:]
	return answer
}

func (this *MyQueue) Peek() int {
	return this.data[0]
}

func (this *MyQueue) Empty() bool {
	return len(this.data) == 0
}
