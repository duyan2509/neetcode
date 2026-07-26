func sortArray(nums []int) []int {
    mergeSort(nums,0,len(nums)-1)
	return nums
}

func mergeSort(nums[]int, l int, r int) {
	if l>=r {
		return
	}
	m:=(l+r)/2
	mergeSort(nums, l, m)
	mergeSort(nums, m+1, r)
	merge(nums,l,m,r)
}

func merge(nums []int, l int, m int, r int) {
	var lefts []int
	for i:=l;i<m+1;i++{
		lefts=append(lefts,nums[i])
	}
	var rights []int
	for i:=m+1;i<r+1;i++{
		rights=append(rights,nums[i])
	}
	indexL:=0
	indexR:=0
	i:=0
	for i=l;i<r+1;i++{
		if indexL==len(lefts) ||indexR==len(rights){
			break
		}

		if lefts[indexL]<rights[indexR] {
			nums[i]=lefts[indexL]
			indexL++
		} else {
			nums[i]=rights[indexR]
			indexR++
		}
	}
	for indexL<len(lefts) {
		nums[i]=lefts[indexL]
		i++
		indexL++
	}
	for indexR<len(rights) {
		nums[i]=rights[indexR]
		i++
		indexR++
	}
}
