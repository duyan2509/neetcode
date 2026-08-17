type RandomizedSet struct {
    Data map[int]bool
}

func Constructor() RandomizedSet {
    return RandomizedSet{
		Data:make(map[int]bool),
	}
}

func (this *RandomizedSet) Insert(val int) bool {
    _ , ok := this.Data[val]
	if !ok {
		this.Data[val]=true
		return true
	}
	return false
}

func (this *RandomizedSet) Remove(val int) bool {
    _ , ok := this.Data[val]
	if ok {
		delete(this.Data,val)
		return true
	}
	return false
}

func (this *RandomizedSet) GetRandom() int {
    keys := make([]int, 0, len(this.Data))

    for val := range this.Data {
        keys = append(keys, val)
    }

    return keys[rand.Intn(len(keys))]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
 