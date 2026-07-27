func findDuplicate(nums []int) int {
    slow := nums[0]
    fast := nums[nums[0]]
	for fast!=slow {
		fast=nums[nums[fast]]
		slow=nums[slow]
	}
	slow2:=0
	for slow2!=slow {
		slow2=nums[slow2]
		slow=nums[slow]
	}
	return slow
}


