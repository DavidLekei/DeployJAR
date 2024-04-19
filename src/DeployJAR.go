package main

import (
	"fmt"
	"os"

	ap "mpi.mb.ca/DeployJAR/argparser"
)

func main() {
	args := os.Args[1:]

	fmt.Println("args: ", args)

	ap.Parse()

	fmt.Println("Done")
}
