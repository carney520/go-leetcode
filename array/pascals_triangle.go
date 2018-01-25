package array

// 这个文件实现了Pascals Triangle

// PascalsTriangle 给定一个行数, 打印一个Pacals三角形
func PascalsTriangle(numRows int) [][]int {
	var res [][]int
	var lastRow []int

	for r := 1; r <= numRows; r++ {
		var row []int
		if lastRow == nil {
			lastRow = []int{1}
			res = append(res, lastRow)
			continue
		}

		for i := 0; i < r; i++ {
			var pre, next int
			if i == 0 {
				pre = 0
			} else {
				pre = lastRow[i-1]
			}

			if i == r-1 {
				next = 0
			} else {
				next = lastRow[i]
			}
			row = append(row, pre+next)
		}

		lastRow = row
		res = append(res, row)
	}

	return res
}
