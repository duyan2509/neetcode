func mergeAlternately(word1 string, word2 string) string {
	rs:=""
	i1:=0
	i2:=0
	i:=0
	for i1<len(word1) && i2<len(word2) {
		if i%2==0 {
			rs+=string(word1[i1])
			i1++
		} else {
			rs+=string(word2[i2])
			i2++
		}
		i++
	}
	for i1<len(word1) {
		rs+=string(word1[i1])
		i1++
	}
	for i2<len(word2) {
		rs+=string(word2[i2])
		i2++
	}
	return rs
}
