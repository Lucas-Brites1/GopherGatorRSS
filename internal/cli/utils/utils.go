package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Reader struct {
	Text string
}

func (r *Reader) Input() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("\nGophergator ")
	scanner.Scan()
	r.Text = scanner.Text()
}

func (r *Reader) Tokenize() []string {
	inputLower := strings.ToLower(r.Text)
	return strings.Fields(inputLower)
}
