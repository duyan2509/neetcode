/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var s1 []int
	var s2 []int
	for l1!=nil{
		s1=append(s1,l1.Val)
		l1=l1.Next
	}
	for l2!=nil {
		s2=append(s2,l2.Val)
		l2=l2.Next
	}
	var rs *ListNode
	promise:=0
	for len(s1)!=0 && len(s2)!=0{
		val:=s1[len(s1)-1]+s2[len(s2)-1]+promise
		if val>9 {
			val=val%10
			promise=1
		} else {
			promise=0
		}
		newNode := &ListNode{
			Val: val,
			Next: rs,
		}
		rs=newNode
		s1=s1[0:len(s1)-1]
		s2=s2[0:len(s2)-1]
	}

	for len(s1)!=0{
		val:=s1[len(s1)-1]+promise
		if val>9 {
			val=val%10
			promise=1
		} else {
			promise=0
		}
		newNode := &ListNode{
			Val: val,
			Next: rs,
		}
		rs=newNode
		s1=s1[0:len(s1)-1]
	}
	for len(s2)!=0{
		val:=s2[len(s2)-1]+promise
		if val>9 {
			val=val%10
			promise=1
		} else {
			promise=0
		}
		newNode := &ListNode{
			Val: val,
			Next: rs,
		}
		rs=newNode
		s2=s2[0:len(s2)-1]
	}
	if promise==1 {
		newNode := &ListNode{
			Val: 1,
			Next: rs,
		}
		rs=newNode
	}
	return rs
}
