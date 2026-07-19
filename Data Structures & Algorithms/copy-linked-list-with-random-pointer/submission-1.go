/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

func copyRandomList(head *Node) *Node {
    var rs *Node 
	if head!=nil {
		rs = &Node{}
	}
	current:=rs
	addressMap := make(map[*Node]*Node)
	for head!=nil {
		current.Val=head.Val
		addressMap[head]=current
		var next *Node
		n, ok := addressMap[head.Next]
		if !ok {
			next = &Node{}
			addressMap[head.Next]=next
		} else{
			next  = n
		}	
		if head.Random!=nil {
			var random *Node
			r, ok := addressMap[head.Random]
			if !ok {
				random = &Node{}
				addressMap[head.Random]=random
			} else{
				random  = r
			}	
			current.Random = random
		}

		if head.Next!=nil {
			current.Next = next
			current = current.Next
		}
		head=head.Next
	}
	return rs
}
