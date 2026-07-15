func reverseString(s []byte) {
	l,r:=0,len(s)-1;
	for l<r {
		tmp:=s[l]
		s[l]=s[r]
		s[r]=tmp
		l++;
		r--
	}
}
