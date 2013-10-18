package main

type CorrelatingTwoBit struct {
	BasePredictor
	predictions [4][4096]int
	history     int
	bitmask     int
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
		p.history = p.history | 2
	}

}
