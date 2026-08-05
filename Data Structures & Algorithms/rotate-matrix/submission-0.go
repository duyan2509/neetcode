func rotate(matrix [][]int)  {
	m := len(matrix)
	n := len(matrix[0])
	l:=0
	r:=n-1
	for l<=r {
		for i:=0;i<m;i++{
			tmp:=matrix[i][l]
			matrix[i][l]=matrix[i][r]
			matrix[i][r]=tmp
		}
		l++
		r--
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ {
			matrix[i][j], matrix[n-1-j][n-1-i] =
				matrix[n-1-j][n-1-i], matrix[i][j]
		}
	}
}
