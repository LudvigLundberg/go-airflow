package main

import (
	"fmt"
	"os"

	"github.com/LudvigLundberg/go-airflow/cli"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		panic("No argument passed")
	}
	if len(args) == 1 {
		switch args[0] {
		case "list_dags":
			listDags()
			break
		default:
			panic("Command not supported: " + args[0])
		}
	}
}

func listDags() {
	dags, err := cli.ListDags()

	if err != nil {
		panic(err)
	}

	for _, dag := range dags {
		fmt.Println(dag)
	}
}
