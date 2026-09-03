func totalFruit(fruits []int) int {
	rs:=0
	l:=0
	set:=make(map[int]int)
	for r:=0;r<len(fruits);r++{
		_ , ok:=set[fruits[r]]
		if l==r || ok || len(set)<2 {
			set[fruits[r]]++
			current:=0
			for _,v:= range set {
				current+=v
			}
			rs=getMax(rs,current)
		} else if len(set)==2 && !ok{
			for len(set)==2 && l<r{
				count, ok := set[fruits[l]]
				if ok && count==1 {
					delete(set,fruits[l])
				} else if ok  {
					set[fruits[l]]--
				}
				l++
			}
			r--
		}
	}
	return rs
}

func getMax (a int, b int) int {
	if a>b {
		return a
	}
	return b
}
