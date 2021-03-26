package utils

// 计算总页数
func TotalPage(size, total int64) int64 {
	if size == 0 {
		return 0
	}
	t := total / size
	if total%size > 0 {
		t += 1
	}
	return t
}
