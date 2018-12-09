package task1

import (
	"container/ring"
)

func Solution(numOfPlayers, lastRound int) int {
	game := NewGame(numOfPlayers, lastRound)
	game.Play()
	return game.GetHighestScore()
}

type Game struct {
	circle    *ring.Ring
	lastRound int
	results   map[int]int
}

func NewGame(numOfPlayers, lastRound int) *Game {
	r := ring.New(1)
	r.Value = 0
	r = r.Next()
	g := &Game{
		circle:    r,
		lastRound: lastRound,
		results:   make(map[int]int, numOfPlayers),
	}
	g.setupPlayers(numOfPlayers)
	return g
}

func (g *Game) Play() {
	for marble := 1; marble <= g.lastRound; marble++ {
		if marble%23 != 0 {
			g.playSpecialRound(marble)
		} else {
			g.playNormalRound(marble)
		}
	}
}

func (g *Game) playNormalRound(marble int) {
	g.circle = g.circle.Move(-6)
	taken := g.circle.Move(-2).Link(g.circle)
	g.results[marble%len(g.results)] += marble + taken.Value.(int)
}

func (g *Game) playSpecialRound(marble int) {
	currentMarble := ring.New(1)
	currentMarble.Value = marble
	g.circle = g.circle.Next()
	g.circle.Link(currentMarble)
	g.circle = g.circle.Next()
}

func (g *Game) GetHighestScore() int {
	best := 0
	for _, score := range g.results {
		if score > best {
			best = score
		}
	}
	return best
}

func (g *Game) setupPlayers(numOfPlayers int) {
	for i := range make([]struct{}, numOfPlayers) {
		g.results[i] = 0
	}
}
