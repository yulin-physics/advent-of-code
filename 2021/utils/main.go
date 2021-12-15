package utils

const(
	MIN = true
	MAX = false
)

func minMaxofInts(a int, b int, min bool) int {
	if min {
		if b < a {
			return b
		}
		return a
	}
	if b > a {
		return b
	}
	return a
}