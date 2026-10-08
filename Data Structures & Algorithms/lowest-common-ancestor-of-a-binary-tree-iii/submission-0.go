/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Parent *Node
 * }
 */

func lowestCommonAncestor(p *Node, q *Node) *Node {
	ancestors:=map[*Node]bool{}
	for p!=nil {
		ancestors[p]=true
		p=p.Parent
	}
	for q!=nil {
		if ancestors[q] {
			return q
		}
		q=q.Parent
	}
	return nil
}	

