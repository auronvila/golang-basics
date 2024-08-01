package main

import (
	"bufio"
	"fmt"
	"github.com/go-notes/note"
	"os"
	"strings"
)

func main() {
	title, content := getNoteData()

	noteData, err := note.New(title, content)
	if err != nil {
		fmt.Println(err)
		return
	}

	noteData.Display()
	err = noteData.Save()
	if err != nil {
		fmt.Println("Saving node failed")
		return
	}

	fmt.Println("Successfully saved the note to the file")
}

func getUserInput(prompt string) string {
	fmt.Print(prompt, ": ")
	reader := bufio.NewReader(os.Stdin)

	text, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println(err)
		return ""
	}
	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}

func getNoteData() (string, string) {
	title := getUserInput("Enter the note title")
	content := getUserInput("Enter note content")
	return title, content
}
