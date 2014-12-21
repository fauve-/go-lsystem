package turt

import (
	//	"log"
	"math"
)

type Heading float64

type Direction uint8

const (
	LEFT = Direction(iota)
	RIGHT
)

type Point struct {
	X int
	Y int
}

type Turtle struct {
	curHeading Heading
	Location   Point
	velocity   float64
	angle      float64
}

func NewTurtle(x, y int, velo float64) *Turtle {
	return &Turtle{
		curHeading: Heading(1),
		Location:   Point{X: x, Y: y},
		velocity:   velo,
	}
}

func (t *Turtle) Turn(d Direction, heading Heading) {
	//Math.PI / 180 * parseInt(args[0])
	t.curHeading = t.curHeading.turn(d, heading)
	t.angle = float64(t.curHeading) * math.Pi / 180.0
}

func (t *Turtle) Forward() {
	//ugh too many runtime checks
	tx := math.Cos(float64(t.angle))
	ty := math.Sin(float64(t.angle))
	//	log.Println(tx, ty)
	//	t.Location.X += int(tx * 180 / math.Pi)
	//	t.Location.Y += int(ty * 180 / math.Pi)
	t.Location.X += int(tx * t.velocity)
	t.Location.Y += int(ty * t.velocity)
	//	log.Printf("%v %v %+v\n", tx*t.velocity, ty*t.velocity, t.Location)
}

//The sin() and cos() functions take only one parameter: the angle. They return a number between -1 and 1. If you multiply this number by the length of the vector, you will get the exact Cartesian coordinates of the vector.

func (h Heading) turn(d Direction, adjust Heading) Heading {
	switch d {
	case RIGHT:
		return h + adjust
	case LEFT:
		return h - adjust
	}
	//	log.Println("deadbeef")
	return 0xdeadbeef
}
