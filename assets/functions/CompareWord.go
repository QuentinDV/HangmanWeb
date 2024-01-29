package functions

func CompareWord(RevealdWord []string, TabWord []string) bool {
	F := len(RevealdWord)
	T := len(TabWord)

	if F != T {
		return false
	}

	for i := 0; i < F; i++ {
		if RevealdWord[i] != TabWord[i] {
			return false
		}
	}

	return true
}
