package task

type Task struct {
	Id       int    `json:"id,omitempty"`
	Question string `json:"question"`
	Answer   string `json:"answer"`
}
