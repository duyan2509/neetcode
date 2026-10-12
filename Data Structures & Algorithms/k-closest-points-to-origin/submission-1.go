func kClosest(points [][]int, k int) [][]int {
	var rs [][]int
	for i:=0;i<len(points);i++ {
		if len(rs)<k{
			index:=len(rs)
			rs=append(rs,points[i])
			for index>=1 && dis2(rs[index])>dis2(rs[(index-1)/2]) {
				rs[index],rs[(index-1)/2]=rs[(index-1)/2],rs[index]
				index=(index-1)/2
			}
		} else if dis2(points[i])<dis2(rs[0]){
			rs[0]=points[i]
			index:=0
			for{
				left:=2*index+1
				right:=2*index+2
				largest:=index
				if left<len(rs)&&dis2(rs[left])>dis2(rs[largest]){
					largest=left
				}
				if right<len(rs)&&dis2(rs[right])>dis2(rs[largest]){
					largest=right
				}
				if largest==index{
					break
				}
				rs[index],rs[largest]=rs[largest],rs[index]
				index=largest
			}
		}
	}
	return rs
}

func dis2(point []int) int{
	return point[1]*point[1]+point[0]*point[0]
}