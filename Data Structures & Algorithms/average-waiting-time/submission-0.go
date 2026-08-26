func averageWaitingTime(customers [][]int) float64 {
    sum:=0
	availableAt:=0
	for i:=0;i<len(customers);i++{
		if customers[i][0]>=availableAt {
			sum+=customers[i][1]
			availableAt=customers[i][1]+customers[i][0]
		} else {
			sum+=availableAt+customers[i][1]-customers[i][0]
			availableAt+=customers[i][1]
		}
	}
    return float64(sum)/float64(len(customers))
}