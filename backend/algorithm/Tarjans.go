package algorithm

import (
	"backend/utils"
	"strconv"
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
	lowKey.Push([]string{fk, "0", "0", fk})

	cT := 1
	for !stackBefore.IsEmpty() {
		valueStack, _ := stackBefore.Pop()

		nextNode := t.Value(valueStack)

		for _, val := range nextNode {
			if !lowKey.IsExist(val) {
				lowKey.Push([]string{val, strconv.Itoa(cT), strconv.Itoa(cT), valueStack})
				stackBefore.Push(val)
				cT++
			} else {
				a, _ := strconv.Atoi(lowKey[lowKey.GetIdxFromKey(val)][1])
				b, _ := strconv.Atoi(lowKey[lowKey.GetIdxFromKey(valueStack)][2])
				if a < b {
					lowKey[lowKey.GetIdxFromKey(valueStack)][2] = lowKey[lowKey.GetIdxFromKey(val)][1]
					if lowKey[lowKey.GetIdxFromKey(val)][3] > valueStack {
						lowKey[lowKey.GetIdxFromKey(val)][3] = valueStack
					}
				}
			}
		}
	}

	for i := len(lowKey) - 1; i >= 0; i-- {
		n := lowKey[i]
		nParent := lowKey[lowKey.GetIdxFromKey(n[3])]

		if n[2] < nParent[1] {
			lowKey[lowKey.GetIdxFromKey(nParent[0])][2] = n[2]
		}
	}

	separate := make(map[string][]string)

	for key := range lowKey {
		n := lowKey[key]
		separate[n[0]] = []string{n[0]}
		for _, cek := range lowKey {
			if n[1] == cek[2] && n[0] != cek[0] {
				separate[n[0]] = append(separate[n[0]], cek[0])
			}
		}
	}

	var result [][]string
	for s := range separate {
		val, _ := separate[s]
		if len(val) == 1 {
			i := 0
			for k := range separate {
				cek, _ := separate[k]
				if !utils.CheckElmt(cek, val[0]) {
					i++
				}
			}
			if i == len(separate)-1 {
				result = append(result, val)
			}
		} else {
			result = append(result, val)
		}
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
	lowKey.Push([]string{fk, "0", "0", fk})

	cT := 1
	for !stackBefore.IsEmpty() {
		valueStack, _ := stackBefore.Pop()

		nextNode := t.Value(valueStack)

		for _, val := range nextNode {
			if !lowKey.IsExist(val) {
				lowKey.Push([]string{val, strconv.Itoa(cT), strconv.Itoa(cT), valueStack})
				stackBefore.Push(val)
				cT++
			} else {
				a, _ := strconv.Atoi(lowKey[lowKey.GetIdxFromKey(val)][1])
				b, _ := strconv.Atoi(lowKey[lowKey.GetIdxFromKey(valueStack)][2])
				if a < b {
					lowKey[lowKey.GetIdxFromKey(valueStack)][2] = lowKey[lowKey.GetIdxFromKey(val)][1]
					if lowKey[lowKey.GetIdxFromKey(val)][3] > valueStack {
						lowKey[lowKey.GetIdxFromKey(val)][3] = valueStack
					}
				}
			}
		}
	}

	for i := len(lowKey) - 1; i >= 0; i-- {
		n := lowKey[i]
		nParent := lowKey[lowKey.GetIdxFromKey(n[3])]

		if n[2] < nParent[1] {
			lowKey[lowKey.GetIdxFromKey(nParent[0])][2] = n[2]
		}
	}

	var result [][]string

	for i := len(lowKey) - 1; i >= 0; i-- {
		n := lowKey[i]
		nParent := lowKey[lowKey.GetIdxFromKey(n[3])]
		if n[2] > nParent[1] {
			result = append(result, []string{nParent[0], n[0]})
		}
	}

	separate := make(map[string][]string)

	for key := range lowKey {
		n := lowKey[key]
		_, isExist := separate[n[2]]
		if isExist {
			separate[n[2]] = append(separate[n[2]], n[0])
		} else {
			separate[n[2]] = []string{n[0]}
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
