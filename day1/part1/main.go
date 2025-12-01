package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

const filename = "day1/input.txt"

func main() {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)

	d := newDial()
	numZeroes := 0

	for scanner.Scan() {
		line := scanner.Text()

		dir := line[0:1]
		num, err1 := strconv.Atoi(line[1:])
		if err1 != nil {
			log.Fatal(err1)
		}

		switch dir {
		case "R":
			d.right(num)
		case "L":
			d.left(num)
		}

		if d.isZero() {
			numZeroes++
		}
	}

	if err2 := scanner.Err(); err2 != nil {
		log.Fatal(err2)
	}

	fmt.Println(numZeroes)
}

// dial - Representative of the safe dial. Handles integer wrapping just like a safe dial.
// (Yes, I know this reeks of OOP and that is a no-no for Go, but it worked!)
type dial struct {
	value  int
	maxVal int
}

func newDial() *dial {
	return &dial{
		value:  50,
		maxVal: 99,
	}
}

func (d *dial) add() {
	d.value += 1
	if d.value > d.maxVal {
		d.value = 0
	}
}

func (d *dial) sub() {
	d.value -= 1
	if d.value < 0 {
		d.value = d.maxVal
	}
}

func (d *dial) right(v int) {
	for range v {
		d.add()
	}
}

func (d *dial) left(v int) {
	for range v {
		d.sub()
	}
}

func (d *dial) isZero() bool {
	return d.value == 0
}
