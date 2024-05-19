package main

import (
	"bufio"
	"os"
	"tz-yadro/internal"
)

func main() {
	var in *bufio.Reader
	in = bufio.NewReader(os.Stdin)
	var out *bufio.Writer
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	internal.SortBalls(in, out)
}
