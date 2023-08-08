package utils

func CheckElmt(arr []string, elmt string) bool {
	for _, val := range arr {
		if elmt == val {
			return true
		}
	}
	return false
}
