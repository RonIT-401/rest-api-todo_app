package todo

type List struct {
	tasks map[string]Task
}

// Конструктор листа. Инициализируем мапу
func NewList() *List {
	return &List{
		tasks: make(map[string]Task),
	}
}

// Добавление задачи по ключу
func (l *List) AddTask(task Task) error {
	_, ok := l.tasks[task.Title]
	if !ok {
		return ErrTaskAlreadyExist
	}

	l.tasks[task.Title] = task

	return nil
}

// Метод отображения всех задач через временную переменную для ограничения доступа к оригинальному листу
func (l *List) ListTasks() map[string]Task {
	tmp := make(map[string]Task, len(l.tasks))

	for k, v := range l.tasks {
		tmp[k] = v
	}

	return tmp
}

// Метод завершения задачи с проверкой на существования этой задачи
func (l *List) CompleteTask(title string) error {
	task, ok := l.tasks[title]
	if !ok {
		return ErrTaskNotFound
	}

	task.Done()

	l.tasks[title] = task

	return nil
}

// Метод отображения не завершенных задач
func (l *List) ListNotComletedTasks() map[string]Task{
	notCompletedTask := make(map[string]Task)

	for title, task := range l.tasks {
		if !task.Completed {
			notCompletedTask[title] = task
		}
	}

	return notCompletedTask
}

// Метод удаления задачи
func (l *List) DeleteTask(title string) error {
	_, ok := l.tasks[title]
	if !ok {
		return ErrTaskNotFound
	}

	delete(l.tasks, title)

	return nil
}
