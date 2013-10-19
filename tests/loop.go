package main

type Flip struct {
	x int
	y int
}

func main() {
	flip := new(Flip)
	for i := 0; i < 100; i++ {
		if flip.x < i {
			flip.x = i + 1
		} else {
			flip.y = i + 1
		}
	}
}
