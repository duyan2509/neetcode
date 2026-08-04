type BrowserHistory struct {
    Head *Node
	Current *Node	
}

type Node struct {
	Val string
	Next *Node
	Prev *Node
}

func Constructor(homepage string) BrowserHistory {
	hp := &Node{
        Val: homepage,
    }
	return BrowserHistory{
		Head:hp,
		Current:hp,
	}
}


func (this *BrowserHistory) Visit(url string)  {
	newPage:=&Node{
		Val:url,
	}
	this.Current.Next=newPage
	newPage.Prev=this.Current
	this.Current=newPage
}


func (this *BrowserHistory) Back(steps int) string {
    for steps>0{
		if this.Current.Prev!=nil{
			this.Current=this.Current.Prev
		}
		steps--
	}
	return this.Current.Val
}


func (this *BrowserHistory) Forward(steps int) string {
    for steps>0{
		if this.Current.Next!=nil{
			this.Current=this.Current.Next
		}
		steps--
	}
	return this.Current.Val
}


/**
 * Your BrowserHistory object will be instantiated and called as such:
 * obj := Constructor(homepage);
 * obj.Visit(url);
 * param_2 := obj.Back(steps);
 * param_3 := obj.Forward(steps);
 */