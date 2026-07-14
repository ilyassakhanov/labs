package main

import (
	"errors"
	"fmt"
)

type Note struct {
	noteTitle string
	noteText  string
}

func (n Note) NewNote(noteTitle, noteText string) (*Note, error) {
	return &Note{
		noteTitle: noteTitle,
		noteText:  noteText,
	}, nil
}

func getNoteData() (string, string, error) {
	title, err := getuserInput("Note title:")

	if err != nil {
		return "", "", err
	}

	content, err := getuserInput("Note content:")

	if err != nil {
		return "", "", err
	}

	return title, content, nil
}

func main() {
	title, content, err := getNoteData()

	if err != nil {
		fmt.Println(err)
	}

}

func getuserInput(prompt string) (string, error) {
	fmt.Print(prompt)
	var value string
	fmt.Scanln(&value)

	if value == "" {
		return "", errors.New("Invalid input")
	}
	return value, nil
}
