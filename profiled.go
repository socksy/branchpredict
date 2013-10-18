package main

import "math/rand"

//import "fmt"

type Profiled struct {
	BasePredictor
	Predictions map[int]TakenRecord
}

type TakenRecord struct {
	taken    int
	nottaken int
}

func (p *Profiled) Profile(pc int, taken bool) {
	//ugh (can't access map[key].property, need to make tmp)
	current := p.Predictions[pc]

	if taken {
		current.taken++
	} else {
		current.nottaken++
	}
	p.Predictions[pc] = current
}

func (p *Profiled) Predict(pc int, taken bool) {
	//fmt.Println(p.Predictions[pc]);
	var prediction bool
	current := p.Predictions[pc]
	if current.taken+current.nottaken == 0 {
		panic("Profile didn't work!")
	}
	certainty := float64(current.taken) / float64(current.taken+current.nottaken)
	//fmt.Println(certainty)
	if rand.Float64() < certainty {
		prediction = true
	} else {
		prediction = false
	}

	if prediction == taken {
		p.correct++
	} else {
		p.incorrect++
	}
}
