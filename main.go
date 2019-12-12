package main

import (
	"fmt"
	"os"
)

func main () {
	for _, k := range os.Args[1:] {
		fmt.Println(k)
	}
}
