package main

import (
	"fmt"

	"../utils"
)

type Probe struct {
	Xmin              int
	Xmax              int
	Ymin              int
	Ymax              int
	InitialVelocities [][]int
}

func main() {
	p := Probe{Xmin: 139, Xmax: 187, Ymin: -148, Ymax: -89}
	fmt.Printf("part one: %d\npart two: %d", p.maxVerticalPos(), p.findDistinctVelocities())
}

func (p *Probe) maxVerticalPos() int {
	//to maximise height, the max initial vertical velocity to hit a target area below y=0 is -ymin-1 (i.e. first instance vertical trajectory becomes negative)
	yPrime := -p.Ymin - 1
	y := 0
	for {
		y += yPrime
		yPrime -= 1
		if yPrime == 0 {
			break
		}
	}
	return y
}

func (p *Probe) findDistinctVelocities() int {
	for dx := 1; dx <= p.Xmax; dx++ {
		for dy := p.Ymin; dy <= -p.Ymin-1; dy++ {
			start := []int{0, 0}
			vel := []int{dx, dy}
			for {
				start[0] += vel[0]
				start[1] += vel[1]
				vel = p.nextStep(vel)
				if start[0]+vel[0] > p.Xmax || start[1]+vel[1] < p.Ymin {
					break
				}
			}
			if p.hasHit(start) {
				p.InitialVelocities = append(p.InitialVelocities, vel)
			}
		}
	}
	return len(p.InitialVelocities)
}

func (p *Probe) hasHit(pos []int) bool {
	if pos[0] <= p.Xmax && pos[0] >= p.Xmin && pos[1] <= p.Ymax && pos[1] >= p.Ymin {
		return true
	}
	return false
}

func (p *Probe) nextStep(vel []int) []int {
	new := make([]int, 2)
	new[0] = utils.MinMaxofInts(0, vel[0]-1, utils.MAX)
	new[1] = vel[1] - 1
	return new
}
