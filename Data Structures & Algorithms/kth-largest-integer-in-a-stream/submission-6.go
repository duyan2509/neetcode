type KthLargest struct {
	Q []int
	K int
}

func Constructor(k int, nums []int) KthLargest {
	obj := KthLargest{
		Q: []int{},
		K: k,
	}

	for _, num := range nums {
		obj.Add(num)
	}

	return obj
}

func (this *KthLargest) Add(val int) int {
	if len(this.Q)<this.K {
		this.Q=append(this.Q, val)
		index:=len(this.Q) - 1
		for index>0 {
			parent:=(index - 1) / 2
			if this.Q[parent]<=this.Q[index] {
				break
			}
			this.Q[parent],this.Q[index]=this.Q[index],this.Q[parent]
			index = parent
		}
	} else if val>this.Q[0] {
		this.Q[0]=val
		this.down(0)
	}
	return this.Q[0]
}

func (this *KthLargest) down(index int) {
	for {
		left:=2*index+1
		right:=2*index+2
		smallest := index
		if left<len(this.Q) && this.Q[left]<this.Q[smallest] {
			smallest = left
		}
		if right<len(this.Q) && this.Q[right]<this.Q[smallest] {
			smallest = right
		}
		if smallest==index {
			break
		}
		this.Q[index],this.Q[smallest]=this.Q[smallest],this.Q[index]
		index = smallest
	}
}