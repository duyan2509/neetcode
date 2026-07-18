func numRescueBoats(people []int, limit int) int {
	sort.Ints(people)
	rs:=0
	left:=0;
	right:=len(people)-1
	for left<=right {
		currentW := people[left]+people[right]
		if currentW>limit {
			rs++
			right--
		} else {
			right--
			left++
			rs++
		}
	}
	return rs
}

