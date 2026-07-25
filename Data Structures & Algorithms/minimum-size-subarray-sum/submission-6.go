func minSubArrayLen(target int, nums []int) int {
	l:=0
	current:=nums[l]
	if current>=target {
		return 1
	}
	rs:=len(nums)+1
	for r:=1;r<len(nums);r++{
		sum:=current+nums[r]
		if sum>=target {
			for current+nums[r]>=target {
				if r-l+1<=rs {
					rs=r-l+1
				}
				current-=nums[l]
				l++
			}
		}
		current+=nums[r]
	}
	if rs==len(nums)+1 {
		return 0
	}
	return rs
}
