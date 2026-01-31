package todo

import "time"

type Task struct {
	Title       string
	Description string
	IsDone      bool

	CreatedAt time.Time
	DoneAt    *time.Time
}

func NewTast(title string, description string) Task {
	return Task{
		Title:       title,
		Description: description,
		IsDone:      false,
		CreatedAt:   time.Now(),
		DoneAt:      nil,
	}
}

func (t *Task) Done() {
	doneTime := time.Now()

	t.IsDone = true
	t.DoneAt = &doneTime
}
