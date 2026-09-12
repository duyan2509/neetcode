type MyHashMap struct {
	Data map[int]int
}

func Constructor() MyHashMap {
    return MyHashMap{
		Data:make(map[int]int),
	}
}

func (this *MyHashMap) Put(key int, value int) {
    this.Data[key]=value
}

func (this *MyHashMap) Get(key int) int {
    val,ok:=this.Data[key]
	if !ok {
		return -1
	}
	return val
}

func (this *MyHashMap) Remove(key int) {
    _,ok:=this.Data[key]
	if ok {
		delete(this.Data,key)
	}
}

/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */