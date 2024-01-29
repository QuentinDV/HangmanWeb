package functions

import (
	"bufio"
	"encoding/json"
	"os"
)

type SaveData struct {
	CurrentDiff string
	Pseudo      string

	RevealdWord1 string
	Difficulty1  string
	Tryleft1     int
	Exist1       bool
	BorderStyle1 string

	RevealdWord2 string
	Difficulty2  string
	Tryleft2     int
	Exist2       bool
	BorderStyle2 string

	RevealdWord3 string
	Difficulty3  string
	Tryleft3     int
	Exist3       bool
	BorderStyle3 string

	RevealdWord4 string
	Difficulty4  string
	Tryleft4     int
	Exist4       bool
	BorderStyle4 string

	RevealdWord5 string
	Difficulty5  string
	Tryleft5     int
	Exist5       bool
	BorderStyle5 string
}

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

func SetBorderStyle(difficulty string) string {
	switch difficulty {
	case "easy":
		return "border-green"
	case "normal":
		return "border-orange"
	case "hard":
		return "border-red"
	default:
		return "border-default"
	}
}

func DisplaySave(pseudo string) SaveData {
	var SaveData SaveData

	SaveData.Pseudo = pseudo

	if FileExists("assets/save/" + pseudo + "/1.txt") {
		file, err := os.Open("assets/save/" + pseudo + "/1.txt")
		if err != nil {
			panic(err)
		}
		defer file.Close()
		scanner := bufio.NewScanner(file)
		err = nil
		var firstLine string
		if scanner.Scan() {
			firstLine = scanner.Text()
		}
		err = nil
		var Game Game
		err = json.Unmarshal([]byte(firstLine), &Game)
		if err != nil {
			panic(err)
		}
		SaveData.RevealdWord1 = TabtoString(Game.RevealdWord)
		SaveData.Difficulty1 = Game.Difficulty
		SaveData.Tryleft1 = Game.Tryleft
		SaveData.BorderStyle1 = SetBorderStyle(Game.Difficulty)
		SaveData.Exist1 = true
	} else {
		SaveData.Exist1 = false
	}

	if FileExists("assets/save/" + pseudo + "/2.txt") {
		file, err := os.Open("assets/save/" + pseudo + "/2.txt")
		if err != nil {
			panic(err)
		}
		defer file.Close()
		scanner := bufio.NewScanner(file)
		err = nil
		var firstLine string
		if scanner.Scan() {
			firstLine = scanner.Text()
		}
		err = nil
		var Game Game
		err = json.Unmarshal([]byte(firstLine), &Game)
		if err != nil {
			panic(err)
		}
		SaveData.RevealdWord2 = TabtoString(Game.RevealdWord)
		SaveData.Difficulty2 = Game.Difficulty
		SaveData.Tryleft2 = Game.Tryleft
		SaveData.BorderStyle2 = SetBorderStyle(Game.Difficulty)
		SaveData.Exist2 = true
	} else {
		SaveData.Exist2 = false
	}

	if FileExists("assets/save/" + pseudo + "/3.txt") {
		file, err := os.Open("assets/save/" + pseudo + "/3.txt")
		if err != nil {
			panic(err)
		}
		defer file.Close()
		scanner := bufio.NewScanner(file)
		err = nil
		var firstLine string
		if scanner.Scan() {
			firstLine = scanner.Text()
		}
		err = nil
		var Game Game
		err = json.Unmarshal([]byte(firstLine), &Game)
		if err != nil {
			panic(err)
		}
		SaveData.RevealdWord3 = TabtoString(Game.RevealdWord)
		SaveData.Difficulty3 = Game.Difficulty
		SaveData.Tryleft3 = Game.Tryleft
		SaveData.BorderStyle3 = SetBorderStyle(Game.Difficulty)
		SaveData.Exist3 = true
	} else {
		SaveData.Exist3 = false
	}

	if FileExists("assets/save/" + pseudo + "/4.txt") {
		file, err := os.Open("assets/save/" + pseudo + "/4.txt")
		if err != nil {
			panic(err)
		}
		defer file.Close()
		scanner := bufio.NewScanner(file)
		err = nil
		var firstLine string
		if scanner.Scan() {
			firstLine = scanner.Text()
		}
		err = nil
		var Game Game
		err = json.Unmarshal([]byte(firstLine), &Game)
		if err != nil {
			panic(err)
		}
		SaveData.RevealdWord4 = TabtoString(Game.RevealdWord)
		SaveData.Difficulty4 = Game.Difficulty
		SaveData.Tryleft4 = Game.Tryleft
		SaveData.BorderStyle4 = SetBorderStyle(Game.Difficulty)
		SaveData.Exist4 = true
	} else {
		SaveData.Exist4 = false
	}

	if FileExists("assets/save/" + pseudo + "/5.txt") {
		file, err := os.Open("assets/save/" + pseudo + "/5.txt")
		if err != nil {
			panic(err)
		}
		defer file.Close()
		scanner := bufio.NewScanner(file)
		err = nil
		var firstLine string
		if scanner.Scan() {
			firstLine = scanner.Text()
		}
		err = nil
		var Game Game
		err = json.Unmarshal([]byte(firstLine), &Game)
		if err != nil {
			panic(err)
		}
		SaveData.RevealdWord5 = TabtoString(Game.RevealdWord)
		SaveData.Difficulty5 = Game.Difficulty
		SaveData.Tryleft5 = Game.Tryleft
		SaveData.BorderStyle5 = SetBorderStyle(Game.Difficulty)
		SaveData.Exist5 = true
	} else {
		SaveData.Exist5 = false
	}

	return SaveData
}

func FileExists(filename string) bool {
	_, err := os.Stat(filename)
	return !os.IsNotExist(err)
}
