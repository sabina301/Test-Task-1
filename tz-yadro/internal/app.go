package internal

import (
	"bufio"
	"fmt"
	"sort"
)

// Scan the entered data and count the number of balls in each container
// and count the balls of each color
func scanData(in *bufio.Reader) ([]int, []int, error) {
	var quantity int
	_, err := fmt.Fscan(in, &quantity)
	if err != nil {
		return nil, nil, err
	}
	colorBallsCount := make([]int, quantity)
	containerBallsCount := make([]int, quantity)

	for i := 0; i < quantity; i++ {
		for j := 0; j < quantity; j++ {
			var balls int
			_, err = fmt.Fscan(in, &balls)
			if err != nil {
				return nil, nil, err
			}
			containerBallsCount[i] += balls
			colorBallsCount[j] += balls
		}
	}
	return colorBallsCount, containerBallsCount, nil
}

// Comparison of the number of balls in each container and balls of each color
func isEqual(colorBallsCount []int, containerBallsCount []int) bool {
	sort.Ints(colorBallsCount)
	sort.Ints(containerBallsCount)
	for i := 0; i < len(colorBallsCount); i++ {
		if colorBallsCount[i] != containerBallsCount[i] {
			return false
		}
	}
	return true
}

// Print result ("yes" or "no")
func printResult(out *bufio.Writer, result bool) {
	if result {
		fmt.Fprint(out, "yes")
	} else {
		fmt.Fprint(out, "no")
	}
}

func SortBalls(in *bufio.Reader, out *bufio.Writer) {
	colorBallsCount, containerBallsCount, err := scanData(in)
	if err != nil {
		fmt.Fprintf(out, "Error while reading input: %v\n", err)
		return
	}
	result := isEqual(colorBallsCount, containerBallsCount)
	printResult(out, result)
}
