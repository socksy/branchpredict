package main

type TwoBit struct {
	BasePredictor
	//size of array here doesn't matter, the bitmask for indexing does
	predictions [4096]int
	bitmask     int
}

type CorrelatingTwoBit struct {
	BasePredictor
	predictions [4][4096]int
	history int
	bitmask int
}

func (p *CorrelatingTwoBit) Predict(pc int, taken bool) {
	current := p.predictions[p.history][pc&p.bitmask]
	var prediction bool
	if current >= 2 {
		prediction = true
	} else {
		prediction = false
	}

	if taken == prediction {
		p.correct++
	} else {
		p.incorrect++
	}

	p.predictions[p.history][pc&p.bitmask] = NextState(current, taken)

	p.history = p.history >> 1
	if taken {
		p.history =  p.history | 2
	}

}


func (p *TwoBit) Predict(pc int, taken bool) {
	current := p.predictions[pc&p.bitmask]

	var t int
	if taken {
		t = 1
	} else {
		t = 0
	}

	if current >= 2 {
		p.correct += t
		p.incorrect += 1 - t
	} else {
		p.incorrect += t
		p.correct += 1 - t
	}

	p.predictions[pc&p.bitmask] = NextState(current, taken)
}

func NextState(state int, taken bool) int {
	//not very elegant, but does the job
	switch {
	case state == 0 && taken:
		return 1
	case state == 0 && !taken:
		return 0
	case state == 1 && taken:
		return 3
	case state == 1 && !taken:
		return 0
	case state == 2 && taken:
		return 3
	case state == 2 && !taken:
		return 0
	case state == 3 && taken:
		return 3
	case state == 3 && !taken:
		return 2
	default:
		return -1
	}
}
