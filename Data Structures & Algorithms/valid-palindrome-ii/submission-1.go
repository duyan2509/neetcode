func validPalindrome(s string) bool {
	l:=0
	r:=len(s)-1
	usedOne:=false
	for l<r{
		if s[r]!=s[l] {
			if usedOne || (s[r-1]!=s[l] && s[l+1]!=s[r]) {
				return false
			} else if r-1>=0 && s[r-1]==s[l] && l+1<len(s) && s[l+1]==s[r] {
				return validPalindrome(s[l:r])||validPalindrome(s[l+1:r+1])
			} else if r-1>=0 && s[r-1]==s[l] {
				r--
				usedOne=true
			} else if l+1<len(s) && s[l+1]==s[r] {
				l++
				usedOne=true
			}
		} else {
			l++
			r--
		}
	}
	return true
}
