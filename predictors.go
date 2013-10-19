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
		pc, err := strconv.Atoi(scanner.Text())
		if err != nil {
			continue
		}
		if (pc == 1) || (pc == 0) {
			continue
		}

		scanner.Scan()

		var taken bool
		taken, err = strconv.ParseBool(scanner.Text())
		if err != nil {
			continue
		}
		p.Predict(pc, taken)
	}
	return p
}

//who needs DRY?
func ProfileAll(r io.Reader, p *Profiled) Predictor {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		pc, err := strconv.Atoi(scanner.Text())
		if err != nil {
			continue
		}
		if (pc == 1) || (pc == 0) {
			continue
		}

		scanner.Scan()

		var taken bool
		taken, err = strconv.ParseBool(scanner.Text())
		if err != nil {
			continue
		}
		p.Profile(pc, taken)
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
	total := float64(p.Correct() + p.Incorrect())
	fmt.Printf("%s — \n%d/%d correct %.2f%% of the time, incorrect %.2f%%\n", name, p.Correct(), p.Incorrect(), 100*float64(p.Correct())/total, 100*float64(p.Incorrect())/total)
}

func Missed(p Predictor) float64 {
	total := float64(p.Correct() + p.Incorrect())
	return 100 * float64(p.Incorrect()) / total
}
