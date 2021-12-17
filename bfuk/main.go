package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/arunmurugan78/bfuk"
)

const (
	version = "1.0.0"
)

func main() {
	var reader *bufio.Reader

	flag.Usage = func() {
		fmt.Println("Usage:\n\tbfuk [OPTIONS] COMMAND\nExample:\n\tbfuk /filepath/filename.bf\nOptions:")
		flag.PrintDefaults()
	}

	bfString := flag.String("e", "", "Execute given string of brainf**k program")

	flag.Parse()

	if len(*bfString) != 0 {
		reader = bufio.NewReader(strings.NewReader(*bfString))
	} else {
		filepath := flag.CommandLine.Arg(0)

		if len(filepath) == 0 {
			// When filepath is not specified
			fmt.Println("bfuk", version)
			fmt.Println("Description: a Brainf**k programming language interpreter.")
			fmt.Println("Author: Arun Murugan")
			flag.Usage()
			os.Exit(0)
		}

		file, err := os.Open(filepath)

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		reader = bufio.NewReader(file)
	}

	bf := bfuk.NewBfuk(reader)

	input := bufio.NewReader(os.Stdin)
	output := bufio.NewWriter(os.Stdout)

	bf.Run(input, output)
	output.Flush()
}
