package main

import (
	"github.com/veandco/go-sdl2/sdl"
	"lsystem/turt"
)

type DragonCurve struct {
	cur []Symbol
	pen *turt.Turtle
}

func NewDragonCurve() *DragonCurve {
	return &DragonCurve{
		cur: []Symbol{F, X},
		pen: turt.NewTurtle(300, 200, 10),
	}
}

func (d *DragonCurve) String() string {
	var out string
	for i := 0; i < len(d.cur); i++ {
		switch d.cur[i] {
		case F:
			out += "F "
		case X:
			out += "X "
		case Y:
			out += "Y "
		case Minus:
			out += "Minus "
		case Plus:
			out += "Plus "

		}
	}
	return out
}

func (d *DragonCurve) currentSymbolSet() []Symbol {
	return d.cur
}

func (d *DragonCurve) setCurrentSymbolSet(s []Symbol) {
	d.cur = s
}

func (d *DragonCurve) applyDrawRule(s Symbol, surf *sdl.Surface) {
	//	log.Printf("in draw loop %+v\n", d.pen.Location)
	switch s {
	case F:
		d.pen.Forward()
		//		log.Println("forward", d.pen.Location.X, d.pen.Location.Y)
		rect := sdl.Rect{int32(d.pen.Location.X), int32(d.pen.Location.Y), 1, 1}
		surf.FillRect(&rect, 0xffff0000)
	case Minus:
		//		log.Println("left", d.pen.Location.X, d.pen.Location.Y)
		d.pen.Turn(turt.LEFT, 90)
	case Plus:
		//		log.Println("right", d.pen.Location.X, d.pen.Location.Y)
		d.pen.Turn(turt.RIGHT, 90)
	default:
		return
	}
}

func (d *DragonCurve) applyRule(s Symbol) []Symbol {
	switch s {
	case X:
		return []Symbol{X, Plus, Y, F}
	case Y:
		return []Symbol{F, X, Minus, Y}
	case Minus:
		return []Symbol{Minus}
	case Plus:
		return []Symbol{Plus}
	case F:
		return []Symbol{F}
	default:
		return []Symbol{}
	}
}
