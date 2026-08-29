func subarraysDivByK(nums []int, k int) int {
	extraTime:=make(map[int]int)
	var prev []int
	count:=0
	prev=append(prev,nums[0])
	if nums[0]%k==0 {
		count++
	}
	extraTime[prev[0]%k]++
	for i:=1;i<len(nums);i++{
		prev=append(prev, nums[i]+prev[i-1])
		if prev[i]%k==0 {
			count++
		}
		time, ok:=extraTime[prev[i]%k]
		if ok {
			count+=time
		}
		extraTime[prev[i]%k]++
	}
	return count
}

