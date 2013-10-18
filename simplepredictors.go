package main

type AlwaysTaken struct {
	BasePredictor
}

type NeverTaken struct {
	BasePredictor
}

//the type before indicates it's a method for AlwaysTaken struct
//implementing the Predictor
func (p *AlwaysTaken) Predict(pc int, taken bool) {
	if taken {
		p.correct++
	} else {
		p.incorrect++
	}
}

func (p *NeverTaken) Predict(pc int, taken bool) {
	if taken {
		p.incorrect++
	} else {
		p.correct++
	}
}
