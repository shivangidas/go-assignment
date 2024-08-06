package handler

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/google/uuid"
	cmdLineApp "github.com/shivangidas/go-to-do-app/taskInterface"
)

var inMemoryTasks = cmdLineApp.TaskList{}

func Setup() {
	sampleTask := []cmdLineApp.Task{{Name: "Email doc for letter", Status: cmdLineApp.StatusEnum(2)},
		{Name: "Format email", Status: cmdLineApp.StatusEnum(3)},
		{Name: "Attach letter", Status: cmdLineApp.StatusEnum(3)},
		{Name: "Send email", Status: cmdLineApp.StatusEnum(3)},
		{Name: "Call access to work", Status: cmdLineApp.StatusEnum(3)}}
	for _, task := range sampleTask {
		inMemoryTasks.AddTask(task)
	}
}

// TODO: improve error handling
func checkErrHTTP(writer http.ResponseWriter, err error, msg string, status int) {
	if err != nil {
		fmt.Println(err)
		http.Error(writer, msg, status)
	}
}
func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func write(writer http.ResponseWriter, msg string) {
	_, err := writer.Write([]byte(msg))
	checkErr(err)
}

func helloHandler(writer http.ResponseWriter, req *http.Request) {
	write(writer, "Hello from server")
}

func interactHandler(writer http.ResponseWriter, req *http.Request) {
	tmpl, err := template.ParseFiles("templates/view.html")
	checkErr(err)
	err = tmpl.Execute(writer, inMemoryTasks)
	checkErr(err)
}

func addTaskHandler(writer http.ResponseWriter, req *http.Request) {
	tmpl, err := template.ParseFiles("templates/addTask.html")
	checkErr(err)
	err = tmpl.Execute(writer, nil)
	checkErr(err)
}

func createHandler(writer http.ResponseWriter, req *http.Request) {
	todo := req.FormValue("todo")
	status, err := strconv.ParseInt(req.FormValue("status"), 10, 64)
	checkErr(err)
	_, err2 := inMemoryTasks.AddTask(cmdLineApp.Task{Name: todo, Status: cmdLineApp.StatusEnum(status)})
	checkErrHTTP(writer, err2, "Cannot add new task", http.StatusBadRequest)
	http.Redirect(writer, req, "/", http.StatusFound)
}

func editLinkHandler(writer http.ResponseWriter, req *http.Request) {
	idStr := req.URL.Query().Get("id")
	id, err := uuid.Parse(idStr)
	fmt.Println(id)
	checkErrHTTP(writer, err, "Invalid task ID", http.StatusBadGateway)
	task, err := inMemoryTasks.SearchTask(id)
	checkErr(err)
	data := struct {
		ID   uuid.UUID
		Task cmdLineApp.Task
	}{
		ID:   id,
		Task: task,
	}
	tmpl, err := template.ParseFiles("templates/updateTask.html")
	checkErr(err)
	err = tmpl.Execute(writer, data)
	checkErr(err)
}

func updateHandler(writer http.ResponseWriter, req *http.Request) {
	idStr := req.FormValue("id")
	id, err := uuid.Parse(idStr)
	checkErrHTTP(writer, err, "Invalid task ID", http.StatusBadGateway)
	todo := req.FormValue("todo")
	status, err := strconv.ParseInt(req.FormValue("status"), 10, 64)
	checkErr(err)
	oldTask, err := inMemoryTasks.SearchTask(id)
	checkErr(err)
	if oldTask.Name != todo {
		inMemoryTasks.UpdateTaskName(id, todo)
	}
	if oldTask.Status != cmdLineApp.StatusEnum(status) {
		inMemoryTasks.UpdateStatus(id, cmdLineApp.StatusEnum(status))
	}
	http.Redirect(writer, req, "/", http.StatusFound)
}

func deleteLinkHandler(writer http.ResponseWriter, req *http.Request) {
	idStr := req.URL.Query().Get("id")
	id, err := uuid.Parse(idStr)
	fmt.Println(id)
	checkErrHTTP(writer, err, "Invalid task ID", http.StatusBadGateway)
	inMemoryTasks.DeleteTask(id)
	http.Redirect(writer, req, "/", http.StatusFound)
}

func Handlers() {
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/", interactHandler)
	http.HandleFunc("/task", addTaskHandler)
	http.HandleFunc("/create", createHandler)
	http.HandleFunc("/edit", editLinkHandler)
	http.HandleFunc("/update", updateHandler)
	http.HandleFunc("/delete", deleteLinkHandler)
}
