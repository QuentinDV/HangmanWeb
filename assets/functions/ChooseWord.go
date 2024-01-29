package functions

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
)

func ChooseWord(difficulty string) ([]string, []string) {
	var word string
	var TabWord []string
	var RevealdWord []string
	file, err := os.Open("assets/txt/" + difficulty + ".txt")
	if err != nil {
		fmt.Println("Erreur lors de l'ouverture du fichier :", err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var TabWords []string
	for scanner.Scan() {
		TabWords = append(TabWords, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Erreur lors de la lecture du fichier :", err)
	}

	// Mot aleatoire
	for len(word) == 0 {
		nombre := rand.Intn(len(TabWords))
		word = TabWords[nombre]
	}

	// nombre de lettre revélé
	TabWord = StringtoTab(word)
	n := len(word)/2 - 1

	for i := 0; i < len(TabWord); i++ {
		RevealdWord = append(RevealdWord, "_")
	}

	for n != 0 {
		i := rand.Intn(len(RevealdWord))
		if RevealdWord[i] == "_" {
			RevealdWord[i] = TabWord[i]
			n -= 1
		}
	}

	return TabWord, RevealdWord
}
