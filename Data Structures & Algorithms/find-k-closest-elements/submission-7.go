func findClosestElements(arr []int, k int, x int) []int {
	lAbs:=10001
	//rAbs:=10001	
	rsl:=0
	for l:=0;l<=len(arr)-k;l++ {
		labs:=arr[l]-x
		if labs<=0 {
			labs=-labs
		}
		rabs:=arr[l+k-1]-x
		if rabs<=0 {
			rabs=-rabs
		}
		if labs<=lAbs && rabs<lAbs {
			lAbs=labs
			//rAbs=rabs
			rsl=l
		}
	}
	return arr[rsl:rsl+k]
}
