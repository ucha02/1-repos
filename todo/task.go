package todo

import "time"

type Task struct {
	Title      string
	Text       string
	IsDone     bool
	CreateTime time.Time
	DoneTime   *time.Time
}

func NewTask(title, text string) *Task {
	return &Task{
		Title:      title,
		Text:       text,
		CreateTime: time.Now(),

		IsDone:   false,
		DoneTime: nil,
	}
}

func (t *Task) StatusChange(isDone bool) {
	if isDone {
		t.IsDone = true

		now := time.Now()
		t.DoneTime = &now
	} else {
		t.IsDone = false
		t.DoneTime = nil
	}
}
