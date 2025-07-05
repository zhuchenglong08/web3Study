package task1

func findLongestString(s []string) string {
	if len(s) == 0 {
		return ""
	}
	//首先找到最短字符串
	var short string = ""
	for _, v := range s {
		if len(v) < len(short) || len(short) == 0 {
			short = v
		}
	}
	for i := 0; i < len(short); i++ {
		for j := 0; j < len(s); j++ {
			if s[j][i] != short[i] {
				return s[j][:i]
			}
		}
	}
	return short
}
