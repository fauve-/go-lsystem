package main

import (
	"github.com/veandco/go-sdl2/sdl"
	"lsystem/turt"
)

type OtherSiriTriangle struct {
	cur []Symbol
	pen *turt.Turtle
}

func NewOtherSiriTriangle() *OtherSiriTriangle {
	return &OtherSiriTriangle{
		cur: []Symbol{A},
		pen: turt.NewTurtle(300, 200, 10),
	}
}

func (s *OtherSiriTriangle) currentSymbolSet() []Symbol {
	return s.cur
}

func (st *OtherSiriTriangle) setCurrentSymbolSet(s []Symbol) {
	st.cur = s
}

func (st *OtherSiriTriangle) applyDrawRule(s Symbol, surf *sdl.Surface) {
	//	log.Printf("in draw loop %+v\n", d.pen.Location)
	switch s {
	case A, B:
		st.pen.Forward()
		//		log.Println("forward", d.pen.Location.X, d.pen.Location.Y)
		rect := sdl.Rect{int32(st.pen.Location.X), int32(st.pen.Location.Y), 1, 1}
		surf.FillRect(&rect, 0xffff0000)
	case Minus:
		//		log.Println("left", d.pen.Location.X, d.pen.Location.Y)
		st.pen.Turn(turt.RIGHT, 60)
	case Plus:
		//		log.Println("right", d.pen.Location.X, d.pen.Location.Y)
		st.pen.Turn(turt.LEFT, 60)
	default:
		return
	}
}

func (st *OtherSiriTriangle) applyRule(s Symbol) []Symbol {
	switch s {
	case A:
		return []Symbol{B, Minus, A, Minus, B}
	case B:
		return []Symbol{A, Plus, B, Plus, A}
	case Minus:
		return []Symbol{Minus}
	case Plus:
		return []Symbol{Plus}
	default:
		return []Symbol{s}
	}
}

func (st *OtherSiriTriangle) String() string {
	var s string
	bod := st.currentSymbolSet()
	lbod := len(bod)
	for i := 0; i < lbod; i++ {
		sym := bod[i]
		switch sym {
		case A:
			s += "A "
		case B:
			s += "B "
		case Minus:
			s += "Minus "
		case Plus:
			s += "Plus "
		}
	}
	return s
}
