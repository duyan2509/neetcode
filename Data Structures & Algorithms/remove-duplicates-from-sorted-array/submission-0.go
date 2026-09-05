func removeDuplicates(nums []int) int {
	if len(nums)<2 {
		return len(nums)
	}
	l:=0
	for r:=1;r<len(nums);r++ {
		if nums[l]==nums[r] {
			continue
		} else {
			nums[l+1]=nums[r]
			l++
		}
	}
	return l+1
}

