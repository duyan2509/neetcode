func numSubarraysWithSum(nums []int, goal int) int {
	return numSubarrayLessThanK(nums, goal)-numSubarrayLessThanK(nums, goal-1)
}

func numSubarrayLessThanK(nums []int, k int) int {
	if k<0 {
		return 0
	}
	count:=0
	current:=0
	l:=0
	for r:=0;r<len(nums);r++ {
		current+=nums[r]
		for l<=r && current>k {
			current-=nums[l]
			l++
		}
		count+=r-l+1
	}
	return count
}

