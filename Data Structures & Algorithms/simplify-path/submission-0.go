func simplifyPath(path string) string {
	if path=="/" {
		return path
	}
	i:=0
	for i<len(path) && path[i]!='/' {
		i++
	}
	st := []string{}
	i++
	current:=""
	for i=i;i<len(path);i++{
		if path[i] == '/' {
			if current==".." {
				if len(st)>0 {
					st=st[0:len(st)-1]
				}
				current=""
			} else if current=="." {
				current=""
			} else if current!="" {
				st=append(st,current)
				current=""
			}
		} else {
			current+=string(path[i])
		}
	}
	if current==".." {
		if len(st)>0 {
			st=st[0:len(st)-1]
		}
	} else if current!="" && current!="."{
		st=append(st,current)
	}
	rs := ""
	for _, s := range st {
		rs += "/" + s 
	}
	if rs=="" {
		return "/"
	}
	return rs 
}

