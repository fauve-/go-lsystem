package main

import (
	"container/list"
	"github.com/veandco/go-sdl2/sdl"
	"lsystem/turt"
)

type Plant struct {
	cur   []Symbol
	pen   *turt.Turtle
	stack *list.List
}

func NewPlant() *Plant {
	return &Plant{
		cur:   []Symbol{X},
		pen:   turt.NewTurtle(300, 200, 10),
		stack: list.New(),
	}
}

func (p *Plant) currentSymbolSet() []Symbol {
	return p.cur
}

func (p *Plant) setCurrentSymbolSet(s []Symbol) {
	p.cur = s
}

// variables : X F
// constants : + − [ ]
// rules  : (X → F-[[X]+X]+F[+FX]-X), (F → FF)
// angle  : 25°

func (p *Plant) applyDrawRule(s Symbol, surf *sdl.Surface) {
	switch s {
	case F:
		p.pen.Forward()
		rect := sdl.Rect{int32(p.pen.Location.X), int32(p.pen.Location.Y), 1, 1}
		surf.FillRect(&rect, 0xffff0000)
	case Minus:
		p.pen.Turn(turt.LEFT, 25)
	case Plus:
		p.pen.Turn(turt.RIGHT, 25)
	case RBracket:
		popped := p.stack.Front().Value.(turt.Turtle)
		p.pen = &popped
	case LBracket:
		p.stack.PushFront(*p.pen)
	}
}

func (p *Plant) applyRule(s Symbol) []Symbol {
	switch s {
	case F:
		return []Symbol{F, F}
	case X:
		return []Symbol{F, Minus, LBracket, LBracket, X, RBracket, Plus, X, RBracket, Plus, F, LBracket, Plus, F, X, RBracket, Minus, X}

		//, Plus, X, RBracket, Plus, F, LBracket, Plus, F, RBracket, Minus, X}
	case Minus:
		return []Symbol{Minus}
	case Plus:
		return []Symbol{Plus}
	case LBracket:
		return []Symbol{LBracket}
	case RBracket:
		return []Symbol{RBracket}
	default:
		return []Symbol{s}
	}
}

// variables : F G
// constants : + −
// start  : F−G−G

func (p *Plant) String() string {
	var s string
	bod := p.currentSymbolSet()
	lbod := len(bod)
	for i := 0; i < lbod; i++ {
		sym := bod[i]
		switch sym {
		case F:
			s += "F "
		case X:
			s += "X "
		case Minus:
			s += "Minus "
		case Plus:
			s += "Plus "
		case LBracket:
			s += "LBracket "
		case RBracket:
			s += "RBracket "
		}
	}
	return s
}
