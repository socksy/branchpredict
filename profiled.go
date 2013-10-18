package main

//import "math/rand"

//import "fmt"

type Profiled struct {
	BasePredictor
	Predictions map[int]TakenRecord
}

type TakenRecord struct {
	taken    int
	nottaken int
}

func (p *Profiled) Predict(pc int, taken bool) {
	var prediction bool
	current := p.Predictions[pc]
	certainty := float64(current.taken) / float64(1+current.taken+current.nottaken)
	//fmt.Println(certainty)
	if certainty > 0.5 {
		prediction = true
	} else {
		prediction = false
	}

	if prediction && taken {
		p.correct++
	} else {
		p.incorrect++
	}

	if taken {
		//ugh
		current.taken++
	} else {
		current.nottaken++
	}
	p.Predictions[pc] = current
}
