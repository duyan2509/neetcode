func compress(chars []byte) int {
    s:=0
	l:=0
	r:=0
	for r<len(chars){
		if r==l {
			chars[s]=chars[r]
			s++
			r++
		} else if chars[r]==chars[l] {
			for r<len(chars) && chars[r]==chars[l] {
				r++
			}
			size:=r-l
			fmt.Println(size)
			left:=s
			for size>0{
				chars[s]=byte(size%10)+'0'
				s++
				size=size/10
			}
			right:=s-1
			for left<right{
				tmp:=chars[right]
				chars[right]=chars[left]
				chars[left]=tmp
				left++
				right--
			}
			l=r
		} else {
			l++
		}
	}
	return s
}
