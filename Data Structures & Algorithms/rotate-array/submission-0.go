func rotate(nums []int, k int) {
	l:=0
	r:=len(nums)-1
	for l<=r{
		tmp:=nums[l]
		nums[l]=nums[r]
		nums[r]=tmp
		l++
		r--
	}
	l=0
	r=k%len(nums)-1
	for l<=r{
		tmp:=nums[l]
		nums[l]=nums[r]
		nums[r]=tmp
		l++
		r--
	}
	l=k%len(nums)
	r=len(nums)-1
	for l<=r{
		tmp:=nums[l]
		nums[l]=nums[r]
		nums[r]=tmp
		l++
		r--
	}
}
