package godo

import (
	"encoding/json"
	"errors"
	"os"
	"time"
)

type Item struct {
	Task       string
	Done       bool
	CreatedAt  time.Time
	FinishedAt time.Time
}

type Godos []Item

func (g *Godos) Add(task string) {
	godo := Item{
		Task:       task,
		Done:       false,
		CreatedAt:  time.Now(),
		FinishedAt: time.Time{},
	}

	*g = append(*g, godo)
}

func (g *Godos) Complete(index int) error {
	ls := *g
	if index <= 0 || index > len(ls) {
		return errors.New("index is not valid")
	}

	ls[index-1].FinishedAt = time.Now()
	ls[index-1].Done = true

	return nil
}

func (g *Godos) Delete(index int) error {
	ls := *g
	if index <= 0 || index > len(ls) {
		return errors.New("index is not valid")
	}

	*g = append(ls[:index-1], ls[index:]...)

	return nil
}

func (g *Godos) Load(filename string) error {
	file, err := os.ReadFile(filename)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}
		return err
	}

	if len(file) == 0 {
		return err
	}

	err = json.Unmarshal(file, g)
	if err != nil {
		return err
	}

	return nil
}

func (g *Godos) Write(filename string) error {
	data, err := json.Marshal(g)
	if err != nil {
		return err
	}

	return os.WriteFile(filename, data, 0644)
}
