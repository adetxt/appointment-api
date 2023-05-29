package utils

func CalculateOffset(page, pageSize int) int {
	return (page - 1) * pageSize
}
