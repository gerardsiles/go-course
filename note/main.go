package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/note/note"
	"example.com/note/todo"
)

type saver interface {
	Save() error
}

type outputtable interface {
	saver
	Display()
}

// * Another way of accepting data can be done through any or interface{}, are equivalent
// func doSomething(value interface{}) {
// switch value.(type) this access the type of the value allowing ass to treat different cases
// case int:
// do something with the int
// case string:
// case float64:
// default:
// Another way to access the typed value is:
// }
// intVal, ok := value.(int)
// floatVal, ok := value.(float64)
// func doSomething(value any) {}

func main() {
	title, content := getNoteData()
	todoText := getUserInput("Todo text: ")

	todo, err := todo.New(todoText)
	if err != nil {
		fmt.Println(err)
		return
	}

	outputData(todo)
	created, err := note.New(title, content)

	if err != nil {
		fmt.Println(err)
		return
	}

	outputData(created)
}

func outputData(data outputtable) error {
	data.Display()
	return saveData(data)
}

func saveData(data saver) error {
	err := data.Save()
	if err != nil {
		fmt.Println("Saving the note failed")
		return err
	}
	fmt.Println("Saving operation succeeded.")
	return nil
}

func getNoteData() (string, string) {
	title := getUserInput("Note title: ")

	content := getUserInput("Note content: ")

	return title, content
}

func getUserInput(prompt string) string {
	fmt.Print(prompt)

	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}
