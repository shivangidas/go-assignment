package handler

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/google/uuid"
	"github.com/shivangidas/go-to-do-app/taskInterface"
)

var inMemoryTasks = taskInterface.TaskList{}

func InjectData() {
	sampleTask := []taskInterface.Task{{Name: "Buy shares", Status: taskInterface.StatusEnum(2)},
		{Name: "Check news", Status: taskInterface.StatusEnum(0)},
		{Name: "Complete assignment", Status: taskInterface.StatusEnum(1)},
		{Name: "Send email", Status: taskInterface.StatusEnum(2)},
		{Name: "Call access to work", Status: taskInterface.StatusEnum(0)}}
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

func interactHandler(writer http.ResponseWriter, req *http.Request) {
	tmpl, err := template.ParseFiles("templates/view.html")
	checkErr(err)
	err = tmpl.Execute(writer, inMemoryTasks)
	checkErr(err)
}

func addTaskHandler(writer http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodPost:
		createHandler(writer, req)
	case http.MethodGet:
		navigateToAdd(writer)
	default:
		http.Error(writer, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
func navigateToAdd(writer http.ResponseWriter) {
	tmpl, err := template.ParseFiles("templates/addTask.html")
	checkErr(err)
	err = tmpl.Execute(writer, nil)
	checkErr(err)
}

func createHandler(writer http.ResponseWriter, req *http.Request) {
	todo := req.FormValue("todo")
	status, err := strconv.ParseInt(req.FormValue("status"), 10, 64)
	checkErrHTTP(writer, err, "Wrong status ", http.StatusBadRequest)
	_, err2 := inMemoryTasks.AddTask(taskInterface.Task{Name: todo, Status: taskInterface.StatusEnum(status)})
	checkErrHTTP(writer, err2, "Did not add task", http.StatusBadRequest)
	http.Redirect(writer, req, "/", http.StatusFound)
}

func updateTaskHandler(writer http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodPost:
		updateHandler(writer, req)
	case http.MethodGet:
		navigateToUpdate(writer, req)
	default:
		http.Error(writer, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
func navigateToUpdate(writer http.ResponseWriter, req *http.Request) {
	idStr := req.URL.Query().Get("id")
	id, err := uuid.Parse(idStr)
	checkErrHTTP(writer, err, "Invalid task ID", http.StatusBadRequest)
	task, err := inMemoryTasks.SearchTask(id)
	checkErrHTTP(writer, err, "Task not found", http.StatusNotFound)
	data := struct {
		ID   uuid.UUID
		Task taskInterface.Task
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
	checkErrHTTP(writer, err, "Invalid task ID", http.StatusBadRequest)
	todo := req.FormValue("todo")
	status, err := strconv.ParseInt(req.FormValue("status"), 10, 64)
	checkErrHTTP(writer, err, "Invalid status", http.StatusBadRequest)
	oldTask, err := inMemoryTasks.SearchTask(id)
	checkErrHTTP(writer, err, "Task not found", http.StatusNotFound)
	if oldTask.Name != todo {
		inMemoryTasks.UpdateTaskName(id, todo)
	}
	if oldTask.Status != taskInterface.StatusEnum(status) {
		inMemoryTasks.UpdateStatus(id, taskInterface.StatusEnum(status))
	}
	http.Redirect(writer, req, "/", http.StatusFound)
}

func deleteHandler(writer http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet { //should be delete but would need javascript
		http.Error(writer, "Method not allowed", http.StatusMethodNotAllowed)
	}
	idStr := req.URL.Query().Get("id")
	id, err := uuid.Parse(idStr)
	checkErrHTTP(writer, err, "Invalid task ID", http.StatusBadRequest)
	inMemoryTasks.DeleteTask(id)
	http.Redirect(writer, req, "/", http.StatusFound)
}

func Handlers() {
	http.HandleFunc("/", interactHandler)
	http.HandleFunc("/task", addTaskHandler)
	http.HandleFunc("/edit", updateTaskHandler)
	http.HandleFunc("/delete", deleteHandler)

	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))
}
