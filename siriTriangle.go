package main

import (
	"github.com/veandco/go-sdl2/sdl"
	"lsystem/turt"
)

type SiriTriangle struct {
	cur []Symbol
	pen *turt.Turtle
}

func NewSiriTriangle() *SiriTriangle {
	return &SiriTriangle{
		cur: []Symbol{F, Minus, G, Minus, G},
		pen: turt.NewTurtle(300, 200, 10),
	}
}

func (s *SiriTriangle) currentSymbolSet() []Symbol {
	return s.cur
}

func (st *SiriTriangle) setCurrentSymbolSet(s []Symbol) {
	st.cur = s
}

func (st *SiriTriangle) applyDrawRule(s Symbol, surf *sdl.Surface) {
	//	log.Printf("in draw loop %+v\n", d.pen.Location)
	switch s {
	case F, G:
		st.pen.Forward()
		//		log.Println("forward", d.pen.Location.X, d.pen.Location.Y)
		rect := sdl.Rect{int32(st.pen.Location.X), int32(st.pen.Location.Y), 1, 1}
		surf.FillRect(&rect, 0xffff0000)
	case Minus:
		//		log.Println("left", d.pen.Location.X, d.pen.Location.Y)
		st.pen.Turn(turt.RIGHT, 120)
	case Plus:
		//		log.Println("right", d.pen.Location.X, d.pen.Location.Y)
		st.pen.Turn(turt.LEFT, 120)
	default:
		return
	}
}

//rules  : (F → F−G+F+G−F), (G → GG)

func (st *SiriTriangle) applyRule(s Symbol) []Symbol {
	switch s {
	case F:
		return []Symbol{F, Minus, G, Plus, F, Plus, G, Minus, F}
	case G:
		return []Symbol{G, G}
	case Minus:
		return []Symbol{Minus}
	case Plus:
		return []Symbol{Plus}
	default:
		return []Symbol{s}
	}
}

// variables : F G
// constants : + −
// start  : F−G−G

func (st *SiriTriangle) String() string {
	var s string
	bod := st.currentSymbolSet()
	lbod := len(bod)
	for i := 0; i < lbod; i++ {
		sym := bod[i]
		switch sym {
		case F:
			s += "F "
		case G:
			s += "G "
		case Minus:
			s += "Minus "
		case Plus:
			s += "Plus "
		}
	}
	return s
}
