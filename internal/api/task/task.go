package task

type Task struct {
	Id     int64  `json:"id"`
	Text   string `json:"text" validate:"omitempty,max=255"`
	Status string `json:"status" validate:"omitempty,max=255"`
}
