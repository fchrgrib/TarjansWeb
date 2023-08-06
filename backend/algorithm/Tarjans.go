package algorithm

import (
	"backend/utils"
)

type Tarjans utils.SpecialStack

func (t *Tarjans) SCC() ([][]string, bool) {
	var lowKey utils.SpecialStack
	var stackBefore utils.Stack

	if t.IsEmpty() {
		return nil, false
	}

	fk := (*t)[0][0]
	stackBefore.Push(fk)
	lowKey.Push([]string{fk, fk, fk})

	for !stackBefore.IsEmpty() {
		valueStack, _ := stackBefore.Pop()

		nextNode := t.Value(valueStack)

		for _, val := range nextNode {
			if !lowKey.IsExist(val) {
				lowKey.Push([]string{val, val, valueStack})
				stackBefore.Push(val)
			} else {
				if lowKey[lowKey.GetIdxFromKey(val)][0] < lowKey[lowKey.GetIdxFromKey(valueStack)][1] {
					lowKey[lowKey.GetIdxFromKey(valueStack)][1] = lowKey[lowKey.GetIdxFromKey(val)][0]
					if lowKey[lowKey.GetIdxFromKey(val)][2] > valueStack {
						lowKey[lowKey.GetIdxFromKey(val)][2] = valueStack
					}
				}
			}
		}
	}

	for i := len(lowKey) - 1; i >= 0; i-- {
		n := lowKey[i]
		nParent := lowKey[lowKey.GetIdxFromKey(n[2])]

		if n[1] < nParent[1] {
			lowKey[lowKey.GetIdxFromKey(nParent[0])][1] = n[1]
		}
	}

	separate := make(map[string][]string)

	for key := range lowKey {
		n := lowKey[key]
		_, isExist := separate[n[1]]
		if isExist {
			separate[n[1]] = append(separate[n[1]], n[0])
		} else {
			separate[n[1]] = []string{n[0]}
		}
	}

	var result [][]string
	for s := range separate {
		val, _ := separate[s]
		result = append(result, val)
	}

	for idx, value := range result {
		if len(value) > 1 {
			result[idx] = append(result[idx], value[0])
		}
	}

	return result, true
}

func (t *Tarjans) Bridge() ([][]string, bool) {

	var lowKey utils.SpecialStack
	var stackBefore utils.Stack

	if t.IsEmpty() {
		return nil, false
	}

	fk := (*t)[0][0]
	stackBefore.Push(fk)
	lowKey.Push([]string{fk, fk, fk})

	for !stackBefore.IsEmpty() {
		valueStack, _ := stackBefore.Pop()

		nextNode := t.Value(valueStack)

		for _, val := range nextNode {
			if !lowKey.IsExist(val) {
				lowKey.Push([]string{val, val, valueStack})
				stackBefore.Push(val)
			} else {
				if lowKey[lowKey.GetIdxFromKey(val)][0] < lowKey[lowKey.GetIdxFromKey(valueStack)][1] {
					lowKey[lowKey.GetIdxFromKey(valueStack)][1] = lowKey[lowKey.GetIdxFromKey(val)][0]
					if lowKey[lowKey.GetIdxFromKey(val)][2] > valueStack {
						lowKey[lowKey.GetIdxFromKey(val)][2] = valueStack
					}
				}
			}
		}
	}

	for i := len(lowKey) - 1; i >= 0; i-- {
		n := lowKey[i]
		nParent := lowKey[lowKey.GetIdxFromKey(n[2])]

		if n[1] < nParent[1] {
			lowKey[lowKey.GetIdxFromKey(nParent[0])][1] = n[1]
		}
	}

	var result [][]string

	for i := len(lowKey) - 1; i >= 0; i-- {
		n := lowKey[i]
		nParent := lowKey[lowKey.GetIdxFromKey(n[2])]
		if n[1] > nParent[0] {
			result = append(result, []string{nParent[0], n[0]})
		}
	}

	separate := make(map[string][]string)

	for key := range lowKey {
		n := lowKey[key]
		_, isExist := separate[n[1]]
		if isExist {
			separate[n[1]] = append(separate[n[1]], n[0])
		} else {
			separate[n[1]] = []string{n[0]}
		}
	}

	var res [][]string
	for s := range separate {
		val, _ := separate[s]
		res = append(res, val)
	}

	for i := 0; i < len(res); i++ {
		if len(res[i]) == 2 {
			result = append(result, []string{res[i][0], res[i][1]})
		}
	}

	return result, true
}

func (t *Tarjans) IsEmpty() bool {
	return len(*t) == 0
}

func (t *Tarjans) GetIdxFromKey(letter string) int {
	for i := 0; i < len(*t); i++ {
		if (*t)[i][0] == letter {
			return i
		}
	}
	for _, cek := range *t {
		for _, tes := range cek {
			println(tes)
		}
	}
	println(letter)
	return -1
}

func (t *Tarjans) Push(body []string) {
	*t = append(*t, body)
}

func (t *Tarjans) Value(letter string) []string {
	var result []string
	idx := t.GetIdxFromKey(letter)

	for i := 1; i < len((*t)[idx]); i++ {
		result = append(result, (*t)[idx][i])
	}

	return result
}
