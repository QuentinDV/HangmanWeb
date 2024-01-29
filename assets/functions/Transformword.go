package functions

func StringtoTab(str string) []string {
	var tab []string
	for i := 0; i < len(str); i++ {
		tab = append(tab, string(str[i]))
	}
	return tab
}

func TabtoString(tab []string) string {
	var str string
	for i := 0; i < len(tab); i++ {
		str += tab[i]
	}
	return str
}

func Displayable(tab []string) string {
	var str string
	for i := 0; i < len(tab); i++ {
		str += tab[i] + " "
	}

	return str
}
