func numOfSubarrays(arr []int, k int, threshold int) int {
	rs:=0
	l:=0
    currentSum := 0
	for r:=0;r<len(arr);r++ {
		currentSum+=arr[r]

		if r-l+1==k {
			if currentSum/k>=threshold {
				rs++
			}
			currentSum-=arr[l]
			l++
		}
	}
	return rs
}
