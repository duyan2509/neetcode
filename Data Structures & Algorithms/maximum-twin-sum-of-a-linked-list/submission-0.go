/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func pairSum(head *ListNode) int {
    size:=0
	current:=head
	for current!=nil {
		size++
		current=current.Next
	}
	var st []int
	current=head
	i:=0
	for i<size/2{
		st=append(st,current.Val)
		current=current.Next
		i++
	}
	rs:=0
	for i<size  {
		rs=max(rs,st[len(st)-1]+current.Val)
		current=current.Next
		st=st[0:len(st)-1]
		i++
	}
	return rs
}

func max(i1 int, i2 int) int {
	if i1>i2 {
		return i1
	}
	return i2
}