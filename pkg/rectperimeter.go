package piscine

// 1st checkpoint test  level

func Rectperimeter(length, width int) int {
	if length < 0 || width < 0 {
			return -1
		} 
	return 2 * (length + width)
}