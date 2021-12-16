package bfuk_test

import (
	"bufio"
	"bytes"
	"log"
	"os"
	"strings"
	"testing"

	"github.com/arunmurugan78/bfuk"
)

func TestBfuk(t *testing.T) {

	r := bufio.NewReader(strings.NewReader(">+++++++++++++++++++++++++++++++++++++++++++++++++++<>>>><<<."))

	bf := bfuk.NewBfuk(r)

	s := bytes.NewBufferString("")

	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(s)

	bf.Run(reader, writer)
	writer.Flush()

	if s.String() != "3" {
		log.Fatalln("expected 3")
	}
}
