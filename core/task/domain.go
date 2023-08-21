package task

type Task struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Complete string `json:"complete"`
}
