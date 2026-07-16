/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func removeNthFromEnd(head *ListNode, n int) *ListNode {
    
	current :=head
	size:=1
	for current.Next!=nil {
		size++
		current=current.Next
	}
	if size == 1 && n==1 {
		return nil
	}
	if size==n {
		return head.Next
	}
	current = head
	for i:=0;i<size-n-1;i++ {
		current=current.Next
	}
	current.Next=current.Next.Next
	return head
}
