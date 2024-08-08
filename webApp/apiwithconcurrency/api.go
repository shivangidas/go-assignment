package apiwithconcurrency

// trial 2 using concurrency
import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/google/uuid"
	tasks "github.com/shivangidas/go-to-do-app/taskWithMutex"
)

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

type TaskServer struct {
	service tasks.TaskService
}

func NewTaskServer(service tasks.TaskService) *TaskServer {
	return &TaskServer{service: service}
}

func (ts *TaskServer) interactHandler(writer http.ResponseWriter, req *http.Request) {
	tmpl, err := template.ParseFiles("templates/view.html")
	checkErr(err)
	err = tmpl.Execute(writer, ts.service.GetAllTask())
	checkErr(err)
}

func (ts *TaskServer) AddTaskHandler(writer http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodPost:
		ts.CreateHandler(writer, req)
	case http.MethodGet:
		ts.navigateToAdd(writer)
	default:
		http.Error(writer, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (ts *TaskServer) navigateToAdd(writer http.ResponseWriter) {
	tmpl, err := template.ParseFiles("templates/addTask.html")
	checkErr(err)
	err = tmpl.Execute(writer, nil)
	checkErr(err)
}

func (ts *TaskServer) CreateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	todo := r.FormValue("todo")
	status, err := strconv.ParseInt(r.FormValue("status"), 10, 64)
	checkErrHTTP(w, err, "Wrong status ", http.StatusBadRequest)
	_, err2 := ts.service.AddTask(tasks.Task{Name: todo, Status: tasks.StatusEnum(status)})
	checkErrHTTP(w, err2, "Did not add task", http.StatusBadRequest)
	http.Redirect(w, r, "/", http.StatusFound)
}
func (ts *TaskServer) updateTaskHandler(writer http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodPost:
		ts.updateHandler(writer, req)
	case http.MethodGet:
		ts.navigateToUpdate(writer, req)
	default:
		http.Error(writer, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (ts *TaskServer) navigateToUpdate(writer http.ResponseWriter, req *http.Request) {
	idStr := req.URL.Query().Get("id")
	id, err := uuid.Parse(idStr)
	checkErrHTTP(writer, err, "Invalid task ID", http.StatusBadRequest)
	task, err := ts.service.SearchTask(id)
	checkErrHTTP(writer, err, "Task not found", http.StatusNotFound)
	data := struct {
		ID   uuid.UUID
		Task tasks.Task
	}{
		ID:   id,
		Task: task,
	}
	tmpl, err := template.ParseFiles("templates/updateTask.html")
	checkErr(err)
	err = tmpl.Execute(writer, data)
	checkErr(err)
}

func (ts *TaskServer) updateHandler(writer http.ResponseWriter, req *http.Request) {
	idStr := req.FormValue("id")
	id, err := uuid.Parse(idStr)
	checkErrHTTP(writer, err, "Invalid task ID", http.StatusBadRequest)
	todo := req.FormValue("todo")
	status, err := strconv.ParseInt(req.FormValue("status"), 10, 64)
	checkErrHTTP(writer, err, "Invalid status", http.StatusBadRequest)
	oldTask, err := ts.service.SearchTask(id)
	checkErrHTTP(writer, err, "Task not found", http.StatusNotFound)
	if oldTask.Name != todo {
		ts.service.UpdateTaskName(id, todo)
	}
	if oldTask.Status != tasks.StatusEnum(status) {
		ts.service.UpdateStatus(id, tasks.StatusEnum(status))
	}
	http.Redirect(writer, req, "/", http.StatusFound)
}

func (ts *TaskServer) deleteHandler(writer http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet { //should be delete but would need javascript
		http.Error(writer, "Method not allowed", http.StatusMethodNotAllowed)
	}
	idStr := req.URL.Query().Get("id")
	id, err := uuid.Parse(idStr)
	checkErrHTTP(writer, err, "Invalid task ID", http.StatusBadRequest)
	ts.service.DeleteTask(id)
	http.Redirect(writer, req, "/", http.StatusFound)
}
func (ts *TaskServer) StartServer() {
	http.HandleFunc("/", ts.interactHandler)
	http.HandleFunc("/task", ts.AddTaskHandler)
	http.HandleFunc("/edit", ts.updateTaskHandler)
	http.HandleFunc("/delete", ts.deleteHandler)

	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))

	log.Println("Server is running at http://localhost:8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func StartServer() {
	taskList := tasks.NewTaskList()
	taskServer := NewTaskServer(taskList)
	taskServer.StartServer()
}
