/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	var ln,rn *ListNode
	current:=head
	var prevL *ListNode
	index:=1
	for current!=nil {
		next:=current.Next
		if index>left && index<=right && ln!=nil{
			current.Next=ln
			ln=current
			if index==right{
				if prevL!=nil {
					prevL.Next=ln
				}
				rn.Next=next
			}
			if left==1 {
				head=ln
			}
		} 
		if index==left-1 && left>1 {
			rn=current.Next
			ln=current.Next
			prevL=current
			next=next.Next
			index++
		} else if index==left && left==1 {
			ln=current
			rn=current
		}
		current=next
		index++
	}
	return head
}
