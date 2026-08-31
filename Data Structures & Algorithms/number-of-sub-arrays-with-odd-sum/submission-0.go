func numOfSubarrays(arr []int) int {
    count:=0
	current:=0
	countOdds:=make(map[bool]int)
	for r:=0;r<len(arr);r++{

		current+=arr[r]
		if current%2==0 {
			count+=countOdds[false]
			countOdds[true]++
		} else {
			count+=1+countOdds[true]
			countOdds[false]++
		}
	}
	return count
}