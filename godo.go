package godo

import (
	"errors"
	"time"
)

type Item struct {
	Task       string
	Done       bool
	CreatedAt  time.Time
	FinishedAt time.Time
}

type Godos []Item

func (t *Godos) Add(task string) {
	godo := Item{
		Task:       task,
		Done:       false,
		CreatedAt:  time.Now(),
		FinishedAt: time.Time{},
	}

	*t = append(*t, godo)
}

func (t *Godos) Complete(index int) error {
	ls := *t
	if index <= 0 || index > len(ls) {
		return errors.New("index is not valid")
	}

	ls[index-1].FinishedAt = time.Now()
	ls[index-1].Done = true

	return nil
}

func (t *Godos) Delete(index int) error {
	ls := *t
	if index <= 0 || index > len(ls) {
		return errors.New("index is not valid")
	}

	*t = append(ls[:index-1], ls[index:]...)

	return nil
}
