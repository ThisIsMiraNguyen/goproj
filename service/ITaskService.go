package service
import(
	"go-todo-app/model"
)
type ITaskService interface {
	CreateTask(task model.Task) error
	GetTasks()([]model.Task, error)
	GetTaskById(id int)(model.Task, error)
	UpdateTask(task model.Task) error
	DeleteTask(id int) error
}