package main

import (
	"bufio"
	"fmt"
	"github.com/go-notes/note"
	"github.com/go-notes/todo"
	"os"
	"strings"
)

type saver interface {
	Save() error
}

type outputtable interface {
	saver
	Display()
}

func main() {
	title, content := getNoteData()
	todoText := getUserInput("Todo text: ")

	todoData, err := todo.New(todoText)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = outPutData(todoData)

	fmt.Println("Successfully saved the note to the file")
	userNote, err := note.New(title, content)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = outPutData(userNote)
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

func saveData(data saver) error {
	err := data.Save()
	if err != nil {
		fmt.Println("Saving node failed")
		return err
	}

	fmt.Println("Successfully saved the note to the file")
	return nil
}

func outPutData(data outputtable) error {
	data.Display()
	return saveData(data)
}
