func asteroidCollision(asteroids []int) []int {
	if len(asteroids)==0{
		return asteroids
	}
	var st []int
	st=append(st,asteroids[0])
	for i:=1;i<len(asteroids);i++ {
		if asteroids[i]<0 {
			absAs:=-asteroids[i]
			//prevL:=len(st)
			win:=true
			for len(st)-1>=0 && st[len(st)-1]>0 && st[len(st)-1]<absAs{
				st=st[0:len(st)-1]
			}
			if len(st)-1>=0 && st[len(st)-1]>0 && st[len(st)-1]>absAs {
				win=false
			}
			if len(st)-1>=0  && st[len(st)-1]==absAs {
				st=st[0:len(st)-1]
			} else if win {
				st=append(st,-absAs)
			}
		} else {
			st=append(st,asteroids[i])
		}
	}	
	return st
}
