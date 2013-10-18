package main

//including a few comments pointing out some quirks in golang

import (
	"os"
)

type Predictors struct {
	BasePredictor
	predictors []Predictor
}

func (p *Predictors) Predict(pc int, taken bool) {
	for _, predictor := range p.predictors {
		predictor.Predict(pc, taken)
	}
}

func main() {
	//(:= works out type for you and declares implicitly)
	p := new(Predictors)
	p.predictors = make([]Predictor, 6)
	p.predictors[0] = new(AlwaysTaken)
	p.predictors[1] = new(NeverTaken)
	p.predictors[2] = &TwoBit {bitmask: 0x1FF}
	p.predictors[3] = &TwoBit {bitmask: 0x3FF}
	p.predictors[4] = &TwoBit {bitmask: 0x7FF}
	p.predictors[5] = &TwoBit {bitmask: 0xFFF}

	f, err := os.Open("pathological.out")
	if err != nil {
		panic(err)
	}
	defer f.Close() //who needs finally?
	GoThroughLines(f, p)
	PrintStats("Always taken", p.predictors[0])
	PrintStats("Never taken", p.predictors[1])
	PrintStats("2-bit predictor, 512 table", p.predictors[2])
	PrintStats("2-bit predictor, 1024 table", p.predictors[3])
	PrintStats("2-bit predictor, 2048 table", p.predictors[4])
	PrintStats("2-bit predictor, 4096 table", p.predictors[5])
}
