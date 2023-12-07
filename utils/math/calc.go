package math

// CalculateOffset 通过pageSize和pageNumber来计算偏移量
func CalculateOffset(pageSize, pageNumber int) int {
	offset := 0
	if pageNumber == 1 {
		offset = 0
	} else {
		num := pageNumber - 1
		offset = num * pageSize
	}
	return offset
}
