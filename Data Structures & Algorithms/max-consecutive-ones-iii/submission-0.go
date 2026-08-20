func longestOnes(nums []int, k int) int {
	rs:=0
	l:=0
	flip:=k
	r:=0
	for r=0;r<len(nums);r++{
		if nums[r]==0 && flip!=0 {
			flip--
		} else if nums[r]==0 && flip == 0 {
			if r-l>rs {
				rs=r-l
			}
			if nums[l]==0 {
				l++
			} else {
				next:=l+1
				for next<len(nums) && nums[next]!=0 {
					next++
				}
				l=next+1
			}
		}
	}	
	if r-l>rs {
		rs=r-l
	}
	return rs
}
