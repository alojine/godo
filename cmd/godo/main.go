package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/alojine/godo"
)

const (
	godosFile = ".godos.json"
)

func main() {
	add := flag.Bool("add", false, "Add a new godo task")
	complete := flag.Int("complete", 0, "Check godo task as completed")
	del := flag.Int("del", 0, "Delete a godo task")

	flag.Parse()

	godos := &godo.Godos{}

	if err := godos.Load(godosFile); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	switch {
	case *add:
		godos.Add("make godo")
		err := godos.Write(godosFile)
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

	default:
		fmt.Fprintln(os.Stdout, "invalid command")
		os.Exit(0)
	}
}
