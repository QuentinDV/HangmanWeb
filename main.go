package main

import (
	functions "HangmanWeb/assets/functions"
	"bufio"
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"os"
)

type Game struct {
	RevealdWord []string
	TabWord     []string
	Letter      []string
	Difficulty  string
	Word        string
	Pseudo      string
	Tryleft     int
	LenTabWord  int
}

var revealdWord []string
var tabWord []string
var letter []string

var pseudo string
var difficulty string
var tryleft = -2

func Hangman(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./assets/html/hangman.html")
}

func ChooseDifficulty(w http.ResponseWriter, r *http.Request) {
	Game := Game{Pseudo: pseudo}
	tmpl := template.Must(template.ParseFiles("assets/html/choose_difficulty.html"))
	tmpl.Execute(w, Game)
}

func Defeat(w http.ResponseWriter, r *http.Request) {
	Game := Game{Pseudo: pseudo, Word: functions.TabtoString(tabWord)}
	tmpl := template.Must(template.ParseFiles("assets/html/defeat.html"))
	tmpl.Execute(w, Game)
}

func Victory(w http.ResponseWriter, r *http.Request) {
	Game := Game{Pseudo: pseudo, Word: functions.TabtoString(tabWord)}
	tmpl := template.Must(template.ParseFiles("assets/html/victory.html"))
	tmpl.Execute(w, Game)
}

func Save(w http.ResponseWriter, r *http.Request) {
	Data := functions.DisplaySave(pseudo)
	Data.CurrentDiff = difficulty
	tmpl := template.Must(template.ParseFiles("assets/html/save.html"))
	tmpl.Execute(w, Data)
}

func Load(w http.ResponseWriter, r *http.Request) {
	Data := functions.DisplaySave(pseudo)
	tmpl := template.Must(template.ParseFiles("assets/html/loadsave.html"))
	tmpl.Execute(w, Data)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	Data := functions.DisplaySave(pseudo)
	tmpl := template.Must(template.ParseFiles("assets/html/delete.html"))
	tmpl.Execute(w, Data)
}

func PseudoForm(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Form data parsing error", http.StatusInternalServerError)
		return
	}
	pseudo = r.Form.Get("pseudo")
	load := r.Form.Get("load")
	delete := r.Form.Get("delete")

	if load == "true" {
		http.Redirect(w, r, "/hangman/load", http.StatusSeeOther)
	} else if delete == "true" {
		http.Redirect(w, r, "/hangman/delete", http.StatusSeeOther)
	}

	fmt.Println("Name:", pseudo)

	http.Redirect(w, r, "/hangman/choose_difficulty", http.StatusSeeOther)
}

func ChooseDifficultyForm(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Form data parsing error", http.StatusInternalServerError)
		return
	}
	if difficulty == "" {
		difficulty = r.Form.Get("difficulty")
		fmt.Println("New Difficulty:", difficulty)
		if difficulty == "easy" {
			http.Redirect(w, r, "/hangman/easy", http.StatusSeeOther)
		} else if difficulty == "normal" {
			http.Redirect(w, r, "/hangman/normal", http.StatusSeeOther)
		} else if difficulty == "hard" {
			http.Redirect(w, r, "/hangman/hard", http.StatusSeeOther)
		}
	} else {
		tabWord, revealdWord = nil, nil
		letter = nil
		tryleft = -2
		difficulty = r.Form.Get("difficulty")
		fmt.Println("New Difficulty:", difficulty)
		if difficulty == "easy" {
			http.Redirect(w, r, "/hangman/easy", http.StatusSeeOther)
		} else if difficulty == "normal" {
			http.Redirect(w, r, "/hangman/normal", http.StatusSeeOther)
		} else if difficulty == "hard" {
			http.Redirect(w, r, "/hangman/hard", http.StatusSeeOther)
		}
	}
}

func SaveForm(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Form data parsing error", http.StatusInternalServerError)
		return
	}
	number := r.Form.Get("number")

	gamestate := Game{
		RevealdWord: revealdWord,
		TabWord:     tabWord,
		Letter:      letter,
		Difficulty:  difficulty,
		Pseudo:      pseudo,
		Tryleft:     tryleft,
	}

	pseudofile := "assets/save/" + pseudo
	err = os.MkdirAll(pseudofile, os.ModePerm)
	if err != nil {
		panic(err)
	}
	file, err := os.Create("assets/save/" + pseudo + "/" + number + ".txt")

	if err != nil {
		panic(err)
	}
	defer file.Close()
	encodedData, err := json.Marshal(gamestate)
	if err != nil {
		panic(err)
	}
	_, err = file.Write(encodedData)
	if err != nil {
		panic(err)
	}

	http.Redirect(w, r, "/hangman", http.StatusSeeOther)
}

func LoadForm(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Form data parsing error", http.StatusInternalServerError)
		return
	}
	number := r.Form.Get("number")

	file, err := os.Open("assets/save/" + pseudo + "/" + number + ".txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	err = nil
	var firstLine string
	if scanner.Scan() {
		firstLine = scanner.Text()
	} else {
		if err := scanner.Err(); err != nil {
			panic(err)
		}
		fmt.Println("The file is empty.")
	}
	err = nil
	var Game Game
	err = json.Unmarshal([]byte(firstLine), &Game)
	if err != nil {
		panic(err)
	}

	revealdWord = Game.RevealdWord
	tabWord = Game.TabWord
	letter = Game.Letter
	difficulty = Game.Difficulty
	pseudo = Game.Pseudo
	tryleft = Game.Tryleft

	http.Redirect(w, r, "/hangman/"+difficulty, http.StatusSeeOther)
}

func DeleteForm(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Form data parsing error", http.StatusInternalServerError)
		return
	}
	number := r.Form.Get("number")
	err = os.Remove("assets/save/" + pseudo + "/" + number + ".txt")
	if err != nil {
		fmt.Println("Erreur lors de la suppression du fichier:", err)
		return
	}

	http.Redirect(w, r, "/hangman", http.StatusSeeOther)
}

func Redirect(w http.ResponseWriter, r *http.Request) {
	if functions.CompareWord(tabWord, revealdWord) && tryleft != -2 && tabWord != nil && revealdWord != nil { // condition victoire
		http.Redirect(w, r, "/hangman/victory", http.StatusSeeOther)
	} else if tabWord == nil && revealdWord == nil || functions.CompareWord(tabWord, revealdWord) {
		if difficulty == "easy" {
			tabWord, revealdWord = functions.ChooseWord(difficulty)
			tryleft = 14
			letter = nil
		} else if difficulty == "normal" {
			tabWord, revealdWord = functions.ChooseWord(difficulty)
			tryleft = 10
			letter = nil
		} else if difficulty == "hard" {
			tabWord, revealdWord = functions.ChooseWord(difficulty)
			tryleft = 8
			letter = nil
		} else {
			fmt.Println("Bug Redirect")
		}
	} else if tryleft <= 0 { // condition defaite
		http.Redirect(w, r, "/hangman/defeat", http.StatusSeeOther)
	}

	Game := Game{RevealdWord: revealdWord, TabWord: tabWord, Letter: letter, Tryleft: tryleft, LenTabWord: len(tabWord), Word: functions.Displayable(revealdWord)}
	tmpl := template.Must(template.ParseFiles("assets/html/" + difficulty + ".html"))
	tmpl.Execute(w, Game)
}

func Guessletter(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Form data parsing error", http.StatusInternalServerError)
		return
	}
	guess := r.Form.Get("guess")

	tabWord, revealdWord, letter, tryleft = functions.TestGuess(guess, tabWord, revealdWord, letter, tryleft)

	http.Redirect(w, r, "/hangman/"+difficulty, http.StatusSeeOther)
}

func main() {
	// Pages
	http.HandleFunc("/hangman", Hangman)
	http.HandleFunc("/hangman/choose_difficulty", ChooseDifficulty)
	http.HandleFunc("/hangman/easy", Redirect)
	http.HandleFunc("/hangman/normal", Redirect)
	http.HandleFunc("/hangman/hard", Redirect)
	http.HandleFunc("/hangman/load", Load)
	http.HandleFunc("/hangman/delete", Delete)

	// Result Of Game
	http.HandleFunc("/hangman/defeat", Defeat)
	http.HandleFunc("/hangman/victory", Victory)
	http.HandleFunc("/hangman/save", Save)

	// Form
	http.HandleFunc("/pseudoform", PseudoForm)
	http.HandleFunc("/guessletter", Guessletter)
	http.HandleFunc("/ChooseDifficultyForm", ChooseDifficultyForm)
	http.HandleFunc("/saveform", SaveForm)
	http.HandleFunc("/loadform", LoadForm)
	http.HandleFunc("/deleteform", DeleteForm)

	// Elements
	http.Handle("/assets/static/", http.StripPrefix("/assets/static/", http.FileServer(http.Dir("./assets/static"))))
	http.Handle("/assets/img/", http.StripPrefix("/assets/img/", http.FileServer(http.Dir("./assets/img"))))

	// Links
	fmt.Println("\nPlay : http://localhost:8080/hangman")
	http.ListenAndServe(":8080", nil)
}
