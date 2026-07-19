func majorityElement(nums []int) []int {
	if len(nums)==1 {
		return nums
	}
	var rs []int
	sort.Ints(nums)
	left:=0
	count:=1
	for left<len(nums)-1{
		if nums[left]==nums[left+1] {
			count++
		} else {
			if count>len(nums)/3{
				rs = append(rs,nums[left])
			}
			count=1
		}
		left++
	}
	if nums[len(nums)-1]==nums[len(nums)-2] {
		if count>len(nums)/3{
			rs = append(rs,nums[left])
		}
	} else if len(nums)/3<1{
		rs = append(rs,nums[left])
	}
	return rs
}
