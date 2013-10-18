package main

//Credits go to "cespare" and "foobaz" for setting me on the straight and narrow
//with regards to go structs and interfaces in #go-nuts on irc.freenode.net

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
)

func GoThroughLines(r io.Reader, p Predictor) Predictor {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		pc, _ := strconv.Atoi(scanner.Text())
		scanner.Scan()
		taken, _ := strconv.ParseBool(scanner.Text())
		p.Predict(pc, taken)
	}
	return p
}

type Predictor interface {
	Predict(pc int, taken bool)
	Correct() int
	Incorrect() int
}

type BasePredictor struct {
	correct   int
	incorrect int
}

func (b *BasePredictor) Correct() int {
	return b.correct
}
func (b *BasePredictor) Incorrect() int {
	return b.incorrect
}

func PrintStats(name string, p Predictor) {
	total := float64(p.Correct()+p.Incorrect())
	fmt.Printf("%s — %d/%d correct %.2f%% of the time, incorrect %.2f%%\n", name, p.Correct(), p.Incorrect(), 100*float64(p.Correct())/total, 100*float64(p.Incorrect())/total)
}
