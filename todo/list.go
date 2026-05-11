package todo

import "sync"

type List struct {
	// ListName string
	TaskList map[string]*Task
	mtx      sync.RWMutex
}

func NewList( /*listName string*/ ) *List {
	return &List{
		// ListName: listName,
		TaskList: make(map[string]*Task),
	}
}

func (l *List) AddTask(t *Task) error {
	l.mtx.Lock()
	defer l.mtx.Unlock()

	if _, exists := l.TaskList[t.Title]; exists {
		return ErrTaskAlreadyExists
	}
	l.TaskList[t.Title] = t
	return nil
}

func (l *List) GetTask(title string) (*Task, error) {
	l.mtx.RLock()
	defer l.mtx.RUnlock()

	task, exists := l.TaskList[title]
	if !exists {
		return nil, ErrTaskNotFound
	}
	return task, nil
}

func (l *List) DelTask(title string) error {
	l.mtx.Lock()
	defer l.mtx.Unlock()

	_, exists := l.TaskList[title]
	if !exists {
		return ErrTaskNotFound
	}
	delete(l.TaskList, title)
	return nil
}

func (l *List) GetAllTasks() map[string]*Task {
	l.mtx.RLock()
	defer l.mtx.RUnlock()

	copy := make(map[string]*Task)
	for k, v := range l.TaskList {
		copy[k] = v
	}
	return copy
}

func (l *List) GetAllUncompletedTasks() map[string]*Task {
	l.mtx.RLock()
	defer l.mtx.RUnlock()

	uncompletedTasks := make(map[string]*Task)
	for k, v := range l.TaskList {
		if !v.IsDone {
			uncompletedTasks[k] = v
		}
	}
	return uncompletedTasks
}

func (l *List) TaskStateChange(isDone bool, title string) (*Task, error) {
	l.mtx.Lock()
	defer l.mtx.Unlock()

	task, exists := l.TaskList[title]
	if !exists {
		return nil, ErrTaskNotFound
	}

	task.StatusChange(isDone)
	return l.TaskList[title], nil
}
