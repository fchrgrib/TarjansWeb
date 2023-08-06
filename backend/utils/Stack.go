package utils

type Stack []string

func (s *Stack) Push(letter string) {
	*s = append(*s, letter)
}

func (s *Stack) Pop() (string, bool) {
	if s.IsEmpty() {
		return "", false
	}

	lengthStack := s.Length() - 1
	result := (*s)[lengthStack]
	*s = (*s)[:lengthStack]
	return result, true
}

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) Length() int {
	return len(*s)
}

type SpecialStack [][]string

func (s *SpecialStack) Push(letter []string) {
	*s = append(*s, letter)
}

func (s *SpecialStack) Pop() ([]string, bool) {
	if s.IsEmpty() {
		return nil, false
	}

	lengthStack := s.Length() - 1
	result := (*s)[lengthStack]
	*s = (*s)[:lengthStack]
	return result, true
}

func (s *SpecialStack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *SpecialStack) Length() int {
	return len(*s)
}

func (s *SpecialStack) IsExist(letter string) bool {
	for _, value := range *s {
		if value[0] == letter {
			return true
		}
	}
	return false
}

func (s *SpecialStack) GetIdxFromKey(letter string) int {
	for i := 0; i < len(*s); i++ {
		if (*s)[i][0] == letter {
			return i
		}
	}
	for _, cek := range *s {
		for _, tes := range cek {
			println(tes)
		}
	}
	println(letter)
	return -1
}

func (s *SpecialStack) Print() {

	for _, value := range *s {
		print(value[0], ": [")
		for _, cek := range value {
			print(cek + " ")
		}
		println("]")
	}
}
