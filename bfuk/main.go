package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/arunmurugan78/bfuk"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage:\n\tbfuk filename.bf")
		return
	}

	filepath := os.Args[1]

	f, _ := os.Open(filepath)

	reader := bufio.NewReader(f)

	bf := bfuk.NewBfuk(reader)

	input := bufio.NewReader(os.Stdin)
	output := bufio.NewWriter(os.Stdout)

	bf.Run(input, output)

	output.Flush()
}
