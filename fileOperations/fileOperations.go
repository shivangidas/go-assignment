package save

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/shivangidas/go-to-do-app/model"
)

const fileName = "data/todoList.json"

func checkError(e error) {
	if e != nil {
		panic(e)
	}
}
func createFile(p string) *os.File {
	f, err := os.Create(p)
	checkError(err)
	return f
}

func writeFile(f *os.File, items ...model.Todo) {
	jsonList, _ := json.Marshal(items)
	fmt.Fprintln(f, string(jsonList))
}

func closeFile(f *os.File) {
	err := f.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v \n", err)
		os.Exit(1)
	}
}
func CreateFileAndWrite(items ...model.Todo) {
	f := createFile(fileName)
	defer closeFile(f)
	writeFile(f, items...)
	data := readFile(fileName)
	fmt.Println(data)
}

func readFile(fileName string) string {
	data, err := os.ReadFile(fileName)
	checkError(err)
	return string(data)
}
