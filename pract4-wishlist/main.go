package main

import (
	"fmt"

	"example.com/note/note"
)

func main() {
	tittle, content := getNoteData()
	userNote, err := note.New(tittle, content)

	if err != nil {
		fmt.Println(err)
		return
	}

	userNote.Display()

}

func getNoteData() (string, string) {
	tittle := getUserInput("Note tittle:")
	content := getUserInput("Note content:")

	return tittle, content
}

func getUserInput(prompt string) string {
	fmt.Print(prompt)
	var value string
	fmt.Scan(&value)

	return value
}
