package structsproject

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"
)

/*

output in json
input :
-title
-content
-created-at

*/

type Note struct {
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created"`
}

// constructor
func New(title, content string) (*Note, error) {

	if title == "" || content == "" {
		return nil, errors.New("eror")
	}

	return &Note{
		Title:     title,
		Content:   content,
		CreatedAt: time.Now(),
	}, nil

}

func Notes() {
	fmt.Println("Notes app")

	title := UserInput("Enter Title : ")
	fmt.Println(title)
	content := UserInput("Enter content: ")

	fmt.Println(content)

	userNote, err := New(title, content)

	noteOutput, _ := json.Marshal(userNote)

	if err != nil {
		return
	}

	byteOutput := []byte(noteOutput)

	writeOutputToFile := os.WriteFile("note", byteOutput, 0644)

	fmt.Println(writeOutputToFile)

	if err != nil {
		return
	}

	fmt.Println(userNote.Title)
	fmt.Println(userNote.Content)
	fmt.Println(userNote.CreatedAt)

}

func UserInput(text string) string {
	var userInput string
	fmt.Print(text)
	fmt.Scan(&userInput)

	return userInput

}
