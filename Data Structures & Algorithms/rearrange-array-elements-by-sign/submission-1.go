func rearrangeArray(nums []int) []int {
	sign:=1
	for i:=0;i<len(nums);i++ {
		if sign*nums[i]<0 {
			next:=i
			for next<len(nums) && nums[next]*sign<0{
				next++
			}
			nextVal:=nums[next]
			for next!=i {
				nums[next]=nums[next-1]
				next--
			}
			nums[i]=nextVal
		}
		sign=-sign
	}
	return nums
}
