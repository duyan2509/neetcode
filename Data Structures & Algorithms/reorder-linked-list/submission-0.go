/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reorderList(head *ListNode) {
    slow := head
	fast := head
	for fast!=nil && fast.Next!=nil {
		slow=slow.Next
		fast=fast.Next.Next
	}
	mid:=slow.Next
	slow.Next=nil

	var prev *ListNode 
	current:=mid
	for current!=nil{
		next:=current.Next
		current.Next=prev
		prev=current
		current=next
	}

	h1:=head
	h2:=prev


	for h1!=nil && h2!=nil{
		tmp1:=h1.Next
		tmp2:=h2.Next
		h1.Next=h2
		h2.Next=tmp1
		h2=tmp2
		h1=tmp1
	}
}

