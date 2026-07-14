func checkInclusion(s1 string, s2 string) bool {
	map1:= make(map[byte]int)
	for i:=0;i<len(s1);i++{
		map1[s1[i]]++
	}
	crtMap:=make(map[byte]int)
	start := 0
	for i:=0;i<len(s2);i++{
		count, ok := map1[s2[i]]
		if ok && count!=0 {
			crtMap[s2[i]]++
			map1[s2[i]]--
			if i-start+1 == len(s1) {
				return true
			}
		} else {
			if s2[start]==s2[i] {
				start++
			} else {
				start = i + 1
				for k, v := range crtMap {
					map1[k]+=v
					v=0
				} 
			}
		}
	}
	return false
}


