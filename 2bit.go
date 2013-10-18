package main

const bitmask512 int = 0x1FF


type TwoBit struct {
	BasePredictor
	//size of array here doesn't matter, the bitmask for indexing does
	predictions [4096]int
	bitmask int
}

func (p *TwoBit) Predict(pc int, taken bool) {
	current := p.predictions[pc & p.bitmask]

	var t int
	if taken {
		t = 1
	} else {
		t = 0
	}

	if current >= 2 {
		p.correct += t
		p.incorrect += 1-t
	} else {
		p.incorrect += t
		p.correct += 1-t
	}

	p.predictions[pc & p.bitmask] = NextState(current, taken)
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
