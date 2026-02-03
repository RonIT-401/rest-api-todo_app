package todo

import "time"

// Структура задачи
type Task struct {
	Title       string
	Description string
	Completed     bool

	CreatedAt time.Time
	CompletedAt    *time.Time
}

// Конструктор задачи
func NewTask(title string, description string) Task {
	return Task{
		Title:       title,
		Description: description,
		Completed:      false,
		CreatedAt:   time.Now(),
		CompletedAt:      nil,
	}
}

// Отметка задачи выполненой с фиксацией времени
func (t *Task) Done() {
	completeTime := time.Now()

	t.Completed = true
	t.CompletedAt = &completeTime
}

// Отметка задачи невыполненой
func (t *Task) Uncomlete() {
	t.Completed = false
	t.CompletedAt = nil
}
