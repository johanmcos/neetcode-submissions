var bracketMap = map[string]string{
	")" : "(",
	"]" : "[",
	"}" : "{",
}

func isValid(s string) bool {
	sArr := strings.Split(s, "")
	var brackets []string
	for i := range sArr {
		if strings.Contains("({[", sArr[i]) {
			brackets = append(brackets, sArr[i])
		} else {
			if i == 0 || len(brackets) == 0 || bracketMap[sArr[i]] != brackets[len(brackets) -1] {
				return false
			}
			brackets = brackets[:len(brackets)-1]
		}
	}

	return len(brackets) == 0
}