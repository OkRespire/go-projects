package main

import (
	"fmt"
	"os"
)

func main() {
	thing, err := os.UserConfigDir()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(thing)
}
