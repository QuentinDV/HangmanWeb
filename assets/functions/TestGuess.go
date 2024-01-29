package functions

import "fmt"

func TestGuess(guess string, TabWord []string, RevealdWord []string, Letter []string, tryleft int) ([]string, []string, []string, int) {
	if len(guess) == 1 {
		char := rune(guess[0]) // recuperation de la seul lettre du guess
		if char > 64 && char < 91 {
			char += 32
		}
		guess = string(char)
		Pos := AllPos(guess, TabWord)

		if len(AllPos(guess, Letter)) != 0 {
			fmt.Println("You have already entered this letter !")
		} else {
			if char > 96 && char < 123 {
				if len(Pos) != 0 {
					for i := 0; i < len(Pos); i++ {
						for j := 0; j < len(RevealdWord); j++ {
							if j == Pos[i] {
								RevealdWord[j] = guess
							}
						}
					}
				} else {
					tryleft--
					fmt.Println("Not present in the word,", tryleft, "attempts remaining")
				}
			} else {
				fmt.Println("Invalid character")
			}
			Letter = append(Letter, guess)
		}
	} else {
		Tabchoice := StringtoTab(guess)
		if CompareWord(Tabchoice, TabWord) {
			RevealdWord = TabWord
		} else {
			tryleft -= 2
			fmt.Println("Not the word")
		}
	}
	return TabWord, RevealdWord, Letter, tryleft
}

func AllPos(char string, Tab []string) []int {
	var Pos []int
	r := rune(char[0])
	if r > 64 && r < 91 {
		char = string(r + 32)
	}
	for i := 0; i < len(Tab); i++ {
		if char == Tab[i] {
			Pos = append(Pos, i)

		}
	}
	return Pos
}
