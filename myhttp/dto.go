package myhttp

import (
	"encoding/json"
	"errors"
	"time"
)

type TaskDTO struct {
	Title string
	Text  string
}

type CompleteTaskDTO struct {
	IsDone bool
}

type ErrorDTO struct {
	Message string
	Time    time.Time
}

func (t TaskDTO) ValidateForCreate() error {
	if t.Title == "" {
		return errors.New("title is empty")
	}

	if t.Text == "" {
		return errors.New("description is empty")
	}

	return nil
}

func NewErrDTO(err error) ErrorDTO {
	return ErrorDTO{
		Message: err.Error(),
		Time:    time.Now(),
	}
}

func (e ErrorDTO) ToString() string {
	b, err := json.MarshalIndent(e, "", "	")
	if err != nil {
		panic(err)
	}
	return string(b)
}
