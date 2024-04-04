package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	f, err := os.Open("name.txt")

	if err != nil {
		fmt.Println(err)
		return
	}

	defer f.Close()

	bs, err := io.ReadAll(f)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(bs))

}
