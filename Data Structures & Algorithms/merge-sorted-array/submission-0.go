func merge(nums1 []int, m int, nums2 []int, n int) {
	var n1[]int
	for i:=0;i<m;i++{
		n1=append(n1,nums1[i])
	}
	i1:=0
	i2:=0
	i := 0
	for i1 < m && i2 < n {
		if n1[i1] < nums2[i2] {
			nums1[i] = n1[i1]
			i1++
		} else {
			nums1[i] = nums2[i2]
			i2++
		}
		i++
	}
	for i1<m{
		nums1[i]=n1[i1]
		i++
		i1++
	}
	for i2<n{
		nums1[i]=nums2[i2]
		i++
		i2++
	}
}
