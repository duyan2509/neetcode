func numSubarrayProductLessThanK(nums []int, k int) int {
	count:=0
	current:=1
	l:=0
	for r:=0;r<len(nums);r++ {
		if nums[r]<k {
			count++
		}
		current*=nums[r]
		if current>=k {
			for l<r && current>=k {
				current/=nums[l]
				l++
			}

		}
		count+=r-l
	}
	return count
}
