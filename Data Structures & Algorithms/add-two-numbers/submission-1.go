/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    rs := &ListNode{}
    current := rs
    tmp:=0
    for l1!=nil || l2!=nil{
        val:=tmp
        tmp=0
        if l1!=nil {
            val+=l1.Val
            l1=l1.Next
        }
        if l2!=nil {
            val+=l2.Val
            l2=l2.Next
        }
        if val+tmp>=10 {
            tmp=1
            val=val%10
        } else {
            val=val+tmp
            tmp=0
        }
        current.Val=val
        next:=&ListNode{}
        if l1!=nil || l2!=nil {
            current.Next=next
            current=current.Next
        } else if l1==nil && l2==nil && tmp==1{
            next.Val=1
            current.Next=next
            current=current.Next
        }
    }
    return rs
}
// 327
// 654
//   1
//  8