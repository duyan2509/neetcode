func checkSubarraySum(nums []int, k int) bool {
	current:=0
	first:=make(map[int]int)
	first[0]=-1
	for i:=0;i<len(nums);i++ {
		current+=nums[i]
		rem:=current % k
		if prev,ok:=first[rem]; ok {
			if i-prev >= 2 {
				return true
			}
		} else {
			first[rem] = i
		}
	}

	return false
}