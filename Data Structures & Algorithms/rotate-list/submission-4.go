/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 
func rotateRight(head *ListNode, k int) *ListNode {
	size := 0
	prevHead:=head
	current := head
	for current != nil {
		size++
		current = current.Next
	}
	    if size <= 1 || k == 0 {
        return head
    }

    k %= size
    if k == 0 {
        return head
    }

    k = size - k

	for k>0 {
		head=head.Next
		k--
	}
	current=head
	for current.Next!=nil{
		current=current.Next
	}
	current.Next=prevHead
	for current.Next!=head{
		current=current.Next
	}
	current.Next=nil
	return head
}
