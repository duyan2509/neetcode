func removeDuplicates(nums []int) int {
	canTwice:=true
	currentIndex:=1
	for i:=1;i<len(nums);i++{
		if nums[i]==nums[i-1] && canTwice {
			canTwice=false
			nums[currentIndex]=nums[i]
			currentIndex++
		} else if nums[i]!=nums[i-1] {
			canTwice=true
			nums[currentIndex]=nums[i]
			currentIndex++
		}
	}
	return currentIndex
}