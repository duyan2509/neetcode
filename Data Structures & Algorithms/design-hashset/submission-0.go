type MyHashSet struct {
	Val map[int]int
}

func Constructor() MyHashSet {
    return MyHashSet{
		Val:make(map[int]int),
	}
}

func (this *MyHashSet) Add(key int) {
    this.Val[key]++
}

func (this *MyHashSet) Remove(key int) {
    _,ok:=this.Val[key]
	if ok {
		delete(this.Val,key)
	}
}

func (this *MyHashSet) Contains(key int) bool {
    _,ok:=this.Val[key]
	if ok {
		return true
	}
	return false
}

/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */
 