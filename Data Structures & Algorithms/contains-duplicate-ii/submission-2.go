func containsNearbyDuplicate(nums []int, k int) bool {
	current:=make(map[int]int) //[value]index
	if k==0 {
		return false
	}
	l:=0
	current[nums[l]]=l
	for r:=1;r<len(nums);r++ {
		if r-l>k{
			delete(current,nums[l])
			l++
		}
		_ ,ok := current[nums[r]]
		if ok {
			return true
		} else {
			current[nums[r]]=r
		}

	}	
	return false
}
