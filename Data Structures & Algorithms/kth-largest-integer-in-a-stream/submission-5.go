type KthLargest struct {
	Q []int
	K int
}

func Constructor(k int, nums []int) KthLargest {
	q := []int{}
	for _, val := range nums {
		if len(q) < k {
			q = append(q, val)
			index := len(q) - 1
			for index > 0 && q[index-1] < val {
				q[index] = q[index-1]
				index--
			}
			q[index] = val
			continue
		}
		if val <= q[k-1] {
			continue
		}
		index := k - 1
		for index > 0 && q[index-1] < val {
			q[index] = q[index-1]
			index--
		}
		q[index] = val
	}

	return KthLargest{
		Q: q,
		K: k,
	}
}

func (this *KthLargest) Add(val int) int {
	if len(this.Q) < this.K {
		this.Q = append(this.Q, val)
		index := len(this.Q) - 1
		for index > 0 && this.Q[index-1] < val {
			this.Q[index] = this.Q[index-1]
			index--
		}
		this.Q[index] = val
		return this.Q[this.K-1]
	}
	if val <= this.Q[this.K-1] {
		return this.Q[this.K-1]
	}
	index := this.K - 1
	for index > 0 && this.Q[index-1] < val {
		this.Q[index] = this.Q[index-1]
		index--
	}
	this.Q[index] = val
	return this.Q[this.K-1]
}