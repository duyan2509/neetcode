func lastStoneWeight(stones []int) int {
	maxHeap := []int{}
	for _, stone := range stones {
		maxHeap=append(maxHeap, stone)
		index:=len(maxHeap)-1
		for index>0 {
			parent:=(index-1)/2
			if maxHeap[index]<=maxHeap[parent] {
				break
			}
			maxHeap[index],maxHeap[parent]=maxHeap[parent],maxHeap[index]
			index=parent
		}
	}
	for len(maxHeap)>1 {
		max1:=maxHeap[0]
		last:=maxHeap[len(maxHeap)-1]
		maxHeap=maxHeap[:len(maxHeap)-1]
		if len(maxHeap)>0 {
			maxHeap[0]=last
			index:=0
			for {
				left:=index*2+1
				right:=index*2+2
				largest:=index
				if left<len(maxHeap) && maxHeap[left]>maxHeap[largest] {
					largest=left
				}
				if right<len(maxHeap) && maxHeap[right]>maxHeap[largest] {
					largest=right
				}
				if largest==index {
					break
				}
				maxHeap[index],maxHeap[largest]=maxHeap[largest],maxHeap[index]
				index=largest
			}
		}
		max2:=maxHeap[0]
		last=maxHeap[len(maxHeap)-1]
		maxHeap=maxHeap[:len(maxHeap)-1]
		if len(maxHeap)>0 {
			maxHeap[0]=last
			index:=0
			for {
				left:=index*2+1
				right:=index*2+2
				largest:=index
				if left<len(maxHeap) && maxHeap[left]>maxHeap[largest] {
					largest=left
				}
				if right<len(maxHeap)&&maxHeap[right]>maxHeap[largest] {
					largest=right
				}
				if largest==index {
					break
				}
				maxHeap[index],maxHeap[largest]=maxHeap[largest],maxHeap[index]
				index=largest
			}
		}
		if max1!=max2 {
			maxHeap=append(maxHeap,max1-max2)
			index:=len(maxHeap) - 1
			for index>0 {
				parent:=(index-1)/2
				if maxHeap[index]<=maxHeap[parent]{
					break
				}
				maxHeap[index],maxHeap[parent]=maxHeap[parent],maxHeap[index]
				index=parent
			}
		}
	}
	if len(maxHeap)==0 {
		return 0
	}
	return maxHeap[0]
}