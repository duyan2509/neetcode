func longestCommonPrefix(strs []string) string {
    compareStr:=strs[0];
	right:=len(compareStr)-1
	for i:=1;i<len(strs);i++{
		limit:=len(strs[i])
		if limit==0 {
			return ""
		}
		if right<len(strs[i]) {
			limit=right+1
		}
		for k:=0;k<limit ;k++{
			if strs[i][k]==compareStr[k]{
				right=k
			} else {
				if k==0 {
					return ""
				}
				break
			}
		}
	}
	return compareStr[:right+1]
}
