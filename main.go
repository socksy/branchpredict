package main

//including a few comments pointing out some quirks in golang

import (
	"flag"
	"os"
	"fmt"
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
	var filename string
	flag.StringVar(&filename, "i", "tests/ls.out", "Input tracefile to use")
	isCsv := flag.Bool("csv", false, "Print mispredictions as a comma separated list")
	flag.Parse()
	//(:= works out type for you and declares implicitly)
	p := new(Predictors)
	p.predictors = make([]Predictor, 11)
	p.predictors[0] = new(AlwaysTaken)
	p.predictors[1] = new(NeverTaken)
	p.predictors[2] = &TwoBit{bitmask: 0x1FF}
	p.predictors[3] = &TwoBit{bitmask: 0x3FF}
	p.predictors[4] = &TwoBit{bitmask: 0x7FF}
	p.predictors[5] = &TwoBit{bitmask: 0xFFF}
	profiled := Profiled{Predictions: make(map[int]TakenRecord)}
	p.predictors[7] = &CorrelatingTwoBit{bitmask: 0x1FF}
	p.predictors[8] = &CorrelatingTwoBit{bitmask: 0x3FF}
	p.predictors[9] = &CorrelatingTwoBit{bitmask: 0x7FF}
	p.predictors[10] = &CorrelatingTwoBit{bitmask: 0xFFF}

	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close() //who needs finally?
	ProfileAll(f, &profiled)
	p.predictors[6] = &profiled
	f, err = os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close() //who needs finally?
	GoThroughLines(f, p)
	if !*isCsv {
		PrettyPrint(p)
	} else {
		for _, pred := range p.predictors {
			fmt.Printf("%f,", Missed(pred))
		}
	}
}

func PrettyPrint(p *Predictors) {
	PrintStats("Always taken", p.predictors[0])
	PrintStats("Never taken", p.predictors[1])
	PrintStats("2-bit predictor, 512 table", p.predictors[2])
	PrintStats("2-bit predictor, 1024 table", p.predictors[3])
	PrintStats("2-bit predictor, 2048 table", p.predictors[4])
	PrintStats("2-bit predictor, 4096 table", p.predictors[5])
	PrintStats("Profiled", p.predictors[6])
	PrintStats("2-bit correlating predictor, 512 table", p.predictors[7])
	PrintStats("2-bit correlating predictor, 1024 table", p.predictors[8])
	PrintStats("2-bit correlating predictor, 2048 table", p.predictors[9])
	PrintStats("2-bit correlating predictor, 4096 table", p.predictors[10])
}
