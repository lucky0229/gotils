package array

//InArray - check find value in array
func InArray[T comparable](arr []T, find T) bool {
	for _, v := range arr {
		if v == find {
			return true
		}
	}
	return false
}
