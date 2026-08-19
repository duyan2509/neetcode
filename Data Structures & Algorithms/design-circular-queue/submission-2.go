type MyCircularQueue struct {
    Head *Node
	Tail *Node
	Capacity int
	Size int
}

type Node struct {
	Val int
	Next *Node 
}
func Constructor(k int) MyCircularQueue {
    return MyCircularQueue{
		Head:nil,
		Tail:nil,
		Capacity:k,
		Size:0,
	}
}


func (this *MyCircularQueue) EnQueue(value int) bool {
    if this.Size==this.Capacity {
		return false
	}

	if this.Head==nil {
		var newNode = &Node{
			Val:value,
		}
		this.Head=newNode
		this.Tail=newNode
		this.Tail.Next=this.Head
	} else {
		var newNode = &Node{
			Val:value,
			Next:this.Head,
		}
		this.Tail.Next=newNode
		this.Tail=newNode
	}
	this.Size++
	return true
}


func (this *MyCircularQueue) DeQueue() bool {
	if this.Size==0 || this.Head==nil {
		return false
	}
	if this.Size==1 {
		this.Head=nil
		this.Tail=nil 
	} else {
		this.Tail.Next=this.Head.Next
		this.Head=this.Head.Next
	}
	this.Size--
	return true
}


func (this *MyCircularQueue) Front() int {
    if this.Size==0 {
		return -1
	}

	return this.Head.Val
}


func (this *MyCircularQueue) Rear() int {
    if this.Size==0 {
		return -1
	}
	
	return this.Tail.Val
}


func (this *MyCircularQueue) IsEmpty() bool {
    return this.Size==0
}


func (this *MyCircularQueue) IsFull() bool {
    return this.Capacity==this.Size
}


/**
 * Your MyCircularQueue object will be instantiated and called as such:
 * obj := Constructor(k);
 * param1 := obj.EnQueue(value);
 * param2 := obj.DeQueue();
 * param3 := obj.Front();
 * param4 := obj.Rear();
 * param5 := obj.IsEmpty();
 * param6 := obj.IsFull();
 */
 