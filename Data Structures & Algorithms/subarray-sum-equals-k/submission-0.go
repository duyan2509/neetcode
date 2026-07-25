func subarraySum(nums []int, k int) int {
	count := map[int]int{
		0: 1,
	}
	prefix := 0
	ans := 0
	for _, x := range nums {
		prefix += x
		ans += count[prefix-k]
		count[prefix]++
	}
	return ans
}
