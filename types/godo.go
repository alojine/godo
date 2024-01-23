package types

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/alexeyco/simpletable"
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

func (g *Godos) PrintTable() {

	table := simpletable.New()

	table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: "#"},
			{Align: simpletable.AlignCenter, Text: "Task"},
			{Align: simpletable.AlignCenter, Text: "Done"},
			{Align: simpletable.AlignRight, Text: "CreatedAt"},
			{Align: simpletable.AlignRight, Text: "FinishedAt"},
		},
	}

	var cells [][]*simpletable.Cell

	for i, item := range *g {
		i++
		cells = append(cells, []*simpletable.Cell{
			{Text: fmt.Sprintf("%d", i)},
			{Text: item.Task},
			{Text: fmt.Sprintf("%t", item.Done)},
			{Text: item.CreatedAt.Format(time.RFC822)},
			{Text: item.FinishedAt.Format(time.RFC822)},
		})
	}

	table.Body = &simpletable.Body{Cells: cells}

	table.Footer = &simpletable.Footer{Cells: []*simpletable.Cell{
		{Align: simpletable.AlignCenter, Span: 5, Text: "Godos!"},
	}}

	table.SetStyle(simpletable.StyleUnicode)
	table.Println()
}
