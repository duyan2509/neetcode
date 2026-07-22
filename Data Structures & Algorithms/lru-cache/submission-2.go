type LRUCache struct {
    cache map[int]*Node
	capacity int
	head *Node
	tail *Node
}
type Node struct {
	Key int
	Val int
	Next *Node
	Prev *Node
}
func Constructor(capacity int) LRUCache {
    return LRUCache{
		cache:make(map[int]*Node),
		capacity:capacity,
		head:nil,
		tail:nil,
	}
}

func (this *LRUCache) Get(key int) int {
    n, ok := this.cache[key]
	if !ok {
		return -1
	}
	if n==this.tail {
		return n.Val
	}
	if n==this.head{
		this.tail.Next=n
		n.Next.Prev=nil
		this.head=n.Next
		n.Next=nil
		n.Prev=this.tail
		this.tail=n
	} else {
		n.Prev.Next=n.Next
		n.Next.Prev=n.Prev
		n.Prev=this.tail
		n.Next=nil
		this.tail.Next=n
		this.tail=n
	}
	return n.Val
}

func (this *LRUCache) Put(key int, value int) {
    n, ok := this.cache[key]
	if !ok {
		if this.capacity == len(this.cache) {
			if this.head==this.tail {
				delete(this.cache, this.head.Key)
				this.head = nil
				this.tail = nil
			} else {
				headKey:=this.head.Key
				this.head.Next.Prev=nil
				this.head=this.head.Next
				delete(this.cache,headKey)
			}
		}
		newNode:=&Node{
			Key:key,
			Val:value,
		}
		if this.tail == nil {
			this.head = newNode
			this.tail = newNode
		} else {
			newNode.Prev = this.tail
			this.tail.Next = newNode
			this.tail = newNode
		}

		this.cache[key]=newNode
	} else {
		this.cache[key].Val=value
		if n == this.tail {
		return
		}
		if n==this.head{
			this.tail.Next=n
			n.Next.Prev=nil
			this.head=n.Next
			n.Next=nil
			n.Prev=this.tail
			this.tail=n
		} else {
			n.Prev.Next=n.Next
			n.Next.Prev=n.Prev
			n.Prev=this.tail
			n.Next=nil
			this.tail.Next=n
			this.tail=n
		}
	}
}
