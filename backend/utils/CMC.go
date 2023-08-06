package utils

func ChangeToMultipleChild(input [][]string) [][]string {
	var ss SpecialStack

	if len(input) == 0 {
		return nil
	}

	for _, vI := range input {
		if !ss.IsExist(vI[0]) {
			ss.Push([]string{vI[0], vI[1]})
		} else {
			ss[ss.GetIdxFromKey(vI[0])] = append(ss[ss.GetIdxFromKey(vI[0])], vI[1])
		}
	}

	return ss
}
