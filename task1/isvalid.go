package task1

func isValid(menber string) bool {
	stacks := []rune{}
	for _, v := range menber {
		switch v {
		case '(', '[', '{':
			stacks = append(stacks, v)
		case ')', ']', '}':
			stacks = stacks[:len(stacks)-1]
		}
	}
	if len(stacks) > 0 {
		return false
	}
	return true
}
