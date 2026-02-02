package todo

import "sync"

type List struct {
	tasks map[string]Task
	mtx   sync.RWMutex
}

// Конструктор листа. Инициализируем мапу
func NewList() *List {
	return &List{
		tasks: make(map[string]Task),
	}
}

// Добавление задачи по ключу
func (l *List) AddTask(task Task) error {
	l.mtx.Lock()
	defer l.mtx.Unlock()

	_, ok := l.tasks[task.Title]
	if ok {
		return ErrTaskAlreadyExist
	}

	l.tasks[task.Title] = task

	return nil
}

func (l *List) GetTask(title string) (Task, error) {
	l.mtx.RLock()
	defer l.mtx.RUnlock()

	task, ok := l.tasks[title]
	if !ok {
		return Task{}, ErrTaskNotFound
	}

	return task, nil
}

// Метод отображения всех задач через временную переменную для ограничения доступа к оригинальному листу
func (l *List) ListTasks() map[string]Task {
	l.mtx.RLock()
	defer l.mtx.RUnlock()

	tmp := make(map[string]Task, len(l.tasks))

	for k, v := range l.tasks {
		tmp[k] = v
	}

	return tmp
}

// Метод завершения задачи с проверкой на существования этой задачи
func (l *List) CompleteTask(title string) (Task, error) {
	l.mtx.Lock()
	defer l.mtx.Unlock()

	task, ok := l.tasks[title]
	if !ok {
		return Task{}, ErrTaskNotFound
	}

	task.Done()

	l.tasks[title] = task

	return l.tasks[title], nil
}

func (l *List) UncomleteTask(title string) (Task, error) {
	l.mtx.Lock()
	defer l.mtx.Unlock()

	task, ok := l.tasks[title]
	if !ok {
		return Task{}, ErrTaskNotFound
	}

	task.Uncomlete()

	l.tasks[title] = task

	return l.tasks[title], nil
}

// Метод отображения не завершенных задач
func (l *List) ListUncomletedTasks() map[string]Task {
	l.mtx.RLock()
	defer l.mtx.RUnlock()

	uncompletedTask := make(map[string]Task)

	for title, task := range l.tasks {
		if !task.Completed {
			uncompletedTask[title] = task
		}
	}

	return uncompletedTask
}

// Метод удаления задачи
func (l *List) DeleteTask(title string) error {
	l.mtx.Lock()
	defer l.mtx.Unlock()

	_, ok := l.tasks[title]
	if !ok {
		return ErrTaskNotFound
	}

	delete(l.tasks, title)

	return nil
}
