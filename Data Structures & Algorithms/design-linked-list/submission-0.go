type MyLinkedList struct {
    Head *Node
	Tail *Node
	Size int 
}

type Node struct {
	Val int
	Next *Node
}

func Constructor() MyLinkedList {
    return MyLinkedList{
		Head:nil,
		Tail:nil,
		Size:0,
	}
}


func (this *MyLinkedList) Get(index int) int {
    if index >= this.Size {
		return -1
	}
   	current := this.Head
    for i := 0; i < index; i++ {
        current = current.Next
    }
	return current.Val
}


func (this *MyLinkedList) AddAtHead(val int)  {
	newNode := &Node{Val:val}
    if this.Head!=nil {
		newNode.Next=this.Head
		this.Head=newNode
	} else {
		this.Tail=newNode
		this.Head=newNode
	}
	this.Size++
}


func (this *MyLinkedList) AddAtTail(val int)  {
	newNode := &Node{Val:val}
    if this.Tail!=nil {
		this.Tail.Next=newNode
		this.Tail=newNode
	} else {
		this.Tail=newNode
		this.Head=newNode
	}
	this.Size++
}


func (this *MyLinkedList) AddAtIndex(index int, val int) {
    if index>this.Size{
        return
    }
    if index==0 {
        this.AddAtHead(val)
        return
    }
    if index==this.Size{
        this.AddAtTail(val)
        return
    }
    current:=this.Head
    for i:=0;i<index-1;i++{
        current = current.Next
    }
    newNode:=&Node{
        Val:val,
        Next:current.Next,
    }
    current.Next=newNode
    this.Size++
}


func (this *MyLinkedList) DeleteAtIndex(index int) {
    if index>=this.Size {
        return
    }
    if index==0 {
        this.Head=this.Head.Next
        this.Size--
        if this.Size==0 {
            this.Tail=nil
        }
        return
    }
    current:=this.Head
    for i:= 0;i<index-1;i++ {
        current=current.Next
    }
    if current.Next==this.Tail {
        this.Tail=current
    }
    current.Next=current.Next.Next
    this.Size--
}


/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */