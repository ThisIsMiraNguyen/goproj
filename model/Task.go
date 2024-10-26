package model
type Task struct {
	Id int `json:"id"`
	Title string `json:"title"`
	Completed bool `json:"completed"`
}