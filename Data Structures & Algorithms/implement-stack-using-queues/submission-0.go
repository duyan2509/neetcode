type MyStack struct {
	Data []int
}

func Constructor() MyStack {
    return MyStack{
        Data: []int{},
    }
}

func (this *MyStack) Push(x int) {
	this.Data=append(this.Data,x)
}

func (this *MyStack) Pop() int {
	val:=this.Data[len(this.Data)-1]
	this.Data=this.Data[0:len(this.Data)-1]
	return val	
}

func (this *MyStack) Top() int {
	val:=this.Data[len(this.Data)-1]
	return val		
}

func (this *MyStack) Empty() bool {
	return len(this.Data)==0
}


/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param2 := obj.Pop();
 * param3 := obj.Top();
 * param4 := obj.Empty();
 */
