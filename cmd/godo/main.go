package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/alojine/godo/types"
)

const (
	godosFile = ".godos.json"
)

func main() {
	list := flag.Bool("list", false, "List godo tasks")
	add := flag.Bool("add", false, "Add a new godo task")
	complete := flag.Int("complete", 0, "Check godo task as completed")
	del := flag.Int("del", 0, "Delete a godo task")
	clear := flag.Bool("clear", false, "Delete a godo task")

	flag.Parse()

	godos := &types.Godos{}

	if err := godos.Load(godosFile); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	switch {
	case *add:
		task, err := getInput(os.Stdin, flag.Args()...)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(0)
		}

		godos.Add(task)

		err = godos.Write(godosFile)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(0)
		}

	case *clear:
		err := godos.Clear()
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(0)
		}
		err = godos.Write(godosFile)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(0)
		}

	case *complete > 0:
		err := godos.Complete(*complete)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(0)
		}
		err = godos.Write(godosFile)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(0)
		}

	case *del > 0:
		err := godos.Delete(*del)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(0)
		}
		err = godos.Write(godosFile)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(0)
		}

	case *list:
		godos.PrintTable()

	default:
		fmt.Fprintln(os.Stdout, "invalid command")
		os.Exit(0)
	}
}

func getInput(reader io.Reader, args ...string) (string, error) {
	if len(args) > 0 {
		return strings.Join(args, " "), nil
	}

	scanner := bufio.NewScanner(reader)
	scanner.Scan()
	if err := scanner.Err(); err != nil {
		return "", err
	}

	text := scanner.Text()

	if len(text) == 0 {
		return "", errors.New("empty godo item is not allowed")
	}

	return text, nil
}
