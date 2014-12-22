package main

import (
	"flag"
	"github.com/veandco/go-sdl2/sdl"
	"log"
)

var (
	iters        int
	chosenSystem string
)

func init() {
	flag.IntVar(&iters, "iters", 5, "iteratoins")
	flag.StringVar(&chosenSystem, "lsys", "siri", "which lsystem do ya want? [tri,tri2,dragon]")
	flag.Parse()
}

//just a bunch of random symbols
const (
	F = Symbol(iota)
	X
	Y
	G
	Plus
	Minus
	LBracket
	RBracket
	A
	B
)

type Symbol uint8

type WorkingSet []Symbol

type LSystem interface {
	currentSymbolSet() []Symbol
	applyRule(Symbol) []Symbol
	setCurrentSymbolSet([]Symbol)
	applyDrawRule(Symbol, *sdl.Surface)
}

func Mutate(l LSystem, iterations int) {
	for i := 0; i < iterations; i++ {
		cursymset := l.currentSymbolSet()
		lcss := len(cursymset)
		newsymset := make([]Symbol, lcss)
		newSymsIdx := 0
		//sorry for that last line being so dumb
		//kind of fooling go's lexer here
		for j := 0; j < lcss; j++ {
			curSym := cursymset[j]
			toAppend := l.applyRule(curSym)
			for x := 0; x < len(toAppend); x, newSymsIdx = x+1, newSymsIdx+1 {
				if len(newsymset) <= newSymsIdx {
					//basically aping the append builtin
					newNewSyms := make([]Symbol, newSymsIdx*2)
					copy(newNewSyms[:newSymsIdx], newsymset)
					newsymset = newNewSyms
				}
				newsymset[newSymsIdx] = toAppend[x]

			}

			//			log.Println(toAppend)
		}
		//		log.Println("at the end of iteration", newsymset[:newSymsIdx])
		l.setCurrentSymbolSet(newsymset[:newSymsIdx])
		//		log.Println("iteration", i)
	}
	//	log.Println(l.currentSymbolSet())
}

func Draw(l LSystem) error {
	window, err := sdl.CreateWindow("test", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, 1200, 800, sdl.WINDOW_SHOWN)

	if err != nil {
		return err
	}
	surface := window.GetSurface()
	sys := l.currentSymbolSet()
	lsys := len(sys)
	for i := 0; i < lsys; i++ {
		sym := sys[i]
		l.applyDrawRule(sym, surface)
	}
	window.UpdateSurface()
	sdl.Delay(10000)
	window.Destroy()
	return nil
}

func main() {
	var sys LSystem
	switch chosenSystem {
	case "dragon":
		sys = NewDragonCurve()
	case "tri":
		sys = NewSiriTriangle()
	case "plant":
		sys = NewPlant()
	case "tri2":
		sys = NewOtherSiriTriangle()
	}
	Mutate(sys, iters)
	err := Draw(sys)
	if err != nil {
		log.Fatalf(err.Error())
	}
	log.Println(sys)
}
