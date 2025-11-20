package types

type TaskItem struct {
	Title  string `json:"title"`
	Status string `json:"status"`
}

type TaskFile struct {
	Tasks []TaskItem `json:"tasks"`
}
