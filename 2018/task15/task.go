package task10

import (
	"fmt"
	"sort"

	"github.com/karolgil/AdventOfCode/2018/utils"
)

func Solution1(inputFile string) int {
	lines, _ := utils.ReadLinesFrom(inputFile)
	game := NewGame(lines, 3)
	game.Play()
	return game.GetResult()
}

func Solution2(inputFile string) int {
	for i := 4; true; i++ {
		lines, _ := utils.ReadLinesFrom(inputFile)
		game := NewGame(lines, i)
		game.Play()
		if game.AllElvesAlive() {
			return game.GetResult()
		}
	}
	return -1
}

type Race int

const (
	Elf    Race = iota
	Goblin Race = iota
)

type Point struct {
	x, y int
}

func (p Point) Neighbours() []Point {
	return []Point{
		{p.x, p.y - 1},
		{p.x - 1, p.y},
		{p.x + 1, p.y},
		{p.x, p.y + 1},
	}
}

type Warrior struct {
	race            Race
	currentPosition Point
	hitPoints       int
	power           int
}

func NewElf(position Point, elfPower int) *Warrior {
	return &Warrior{
		race:            Elf,
		currentPosition: position,
		hitPoints:       200,
		power:           elfPower,
	}
}

func NewGoblin(position Point) *Warrior {
	return &Warrior{
		race:            Goblin,
		currentPosition: position,
		hitPoints:       200,
		power:           3,
	}
}

func (w *Warrior) IsAlive() bool {
	return w.hitPoints > 0
}

func (w *Warrior) WeakSpots() []Point {
	return w.currentPosition.Neighbours()
}

func (w *Warrior) InRangeToAny(opponentsFreeWeakSpots map[Point]*Warrior) bool {
	if _, exists := opponentsFreeWeakSpots[w.currentPosition]; exists {
		return true
	}
	return false
}

func (g *Game) MoveWarriorTowards(warrior *Warrior, opponentsFreeWeakSpots map[Point]*Warrior) {
	visitedLocations := make(map[Point]struct{})
	possibleTargetLocations := make(map[Point]LocationStep)
	possibleTargetLocations[warrior.currentPosition] = LocationStep{
		warrior.currentPosition,
		0,
	}
	locationsToVisit := []LocationStep{
		{warrior.currentPosition, 0},
	}
	occupiedLocations := g.GetAllOccupiedLocationsAsideFrom(warrior)

	for len(locationsToVisit) > 0 {
		currentlyVisitedLocation := locationsToVisit[0]
		locationsToVisit = locationsToVisit[1:]
		for _, neighbour := range currentlyVisitedLocation.parent.Neighbours() {
			if _, isOccupied := occupiedLocations[neighbour]; isOccupied || g.WallsMap[neighbour] == true {
				continue
			}
			if _, isTargetAlready := possibleTargetLocations[neighbour]; !isTargetAlready || possibleTargetLocations[neighbour].distance > currentlyVisitedLocation.distance+1 {
				possibleTargetLocations[neighbour] = LocationStep{currentlyVisitedLocation.parent, currentlyVisitedLocation.distance + 1}
			}
			if _, alreadyVisited := visitedLocations[neighbour]; alreadyVisited {
				continue
			}
			alreadyPlannedVisit := false
			for _, plannedVisit := range locationsToVisit {
				if plannedVisit.parent == neighbour {
					alreadyPlannedVisit = true
				}
			}
			if !alreadyPlannedVisit {
				locationsToVisit = append(locationsToVisit, LocationStep{neighbour, currentlyVisitedLocation.distance + 1})
			}
		}
		visitedLocations[currentlyVisitedLocation.parent] = struct{}{}
	}

	if len(possibleTargetLocations) == 0 {
		return
	}

	var closestLocationSteps []LocationStep
	for point, locationStep := range possibleTargetLocations {
		if _, isWeakSpot := opponentsFreeWeakSpots[point]; isWeakSpot {
			closestLocationSteps = append(closestLocationSteps, LocationStep{
				parent:   point,
				distance: locationStep.distance,
			})
		}
	}
	sort.Slice(closestLocationSteps, func(i, j int) bool {
		if closestLocationSteps[i].distance != closestLocationSteps[j].distance {
			return closestLocationSteps[i].distance < closestLocationSteps[j].distance
		}
		if closestLocationSteps[i].parent.y != closestLocationSteps[j].parent.y {
			return closestLocationSteps[i].parent.y < closestLocationSteps[j].parent.y
		}
		return closestLocationSteps[i].parent.x < closestLocationSteps[j].parent.x
	})

	if len(closestLocationSteps) > 0 {
		move := closestLocationSteps[0].parent
		for possibleTargetLocations[move].distance > 1 {
			move = possibleTargetLocations[move].parent
		}
		if move.x != 0 && move.y != 0 {
			warrior.currentPosition = move
		}
	}
}

type LocationStep struct {
	parent   Point
	distance int
}

type Wall bool

type Game struct {
	WallsMap map[Point]Wall
	Warriors []*Warrior
	Rounds   int
}

func NewGame(inputLines []string, elfPower int) *Game {
	g := Game{
		WallsMap: make(map[Point]Wall),
		Rounds:   0,
	}
	for y, line := range inputLines {
		for x, char := range line {
			currentPoint := Point{x, y}
			switch string(char) {
			case "#":
				g.WallsMap[currentPoint] = true
				break
			case ".":
				g.WallsMap[currentPoint] = false
				break
			case "E":
				g.WallsMap[currentPoint] = false
				g.Warriors = append(g.Warriors, NewElf(currentPoint, elfPower))
				break
			case "G":
				g.WallsMap[currentPoint] = false
				g.Warriors = append(g.Warriors, NewGoblin(currentPoint))
				break
			}
		}
	}
	return &g
}

func (g *Game) GetResult() int {
	sumHitPoints := 0
	for _, warrior := range g.Warriors {
		if warrior.IsAlive() {
			sumHitPoints += warrior.hitPoints
		}
	}
	return g.Rounds * sumHitPoints
}

func (g *Game) Play() {
	for {
		//g.ShowMap()
		if !g.PlayRound() {
			//g.ShowMap()
			return
		}
		g.Rounds += 1
	}
}

func (g *Game) ShowMap() {
	for y := 0; y < 35; y++ {
		for x := 0; x < 35; x++ {
			if g.WallsMap[Point{x, y}] {
				print("#")
			} else {
				isWarrior := false
				for _, warrior := range g.Warriors {
					if warrior.currentPosition.x == x && warrior.currentPosition.y == y {
						isWarrior = true
						if warrior.race == Elf && warrior.IsAlive() {
							fmt.Print("E")
						} else if warrior.race == Goblin && warrior.IsAlive() {
							fmt.Print("G")
						} else {
							fmt.Print(".")
						}
					}
				}
				if !isWarrior {
					fmt.Print(".")
				}
			}
		}
		fmt.Print("\n")
	}
	fmt.Println()
}

func (g *Game) PlayRound() bool {
	g.SortWarriors()
	for _, warrior := range g.Warriors {
		if !warrior.IsAlive() {
			continue
		}
		if !g.HasAnyOpponentFor(warrior) {
			return false
		}
		g.PlayWarriorsRound(warrior)
	}
	return true
}

func (g *Game) SortWarriors() {
	sort.Slice(g.Warriors, func(i, j int) bool {
		if g.Warriors[i].currentPosition.y != g.Warriors[j].currentPosition.y {
			return g.Warriors[i].currentPosition.y < g.Warriors[j].currentPosition.y
		} else {
			return g.Warriors[i].currentPosition.x < g.Warriors[j].currentPosition.x
		}
	})
}

func (g *Game) HasAnyOpponentFor(warrior *Warrior) bool {
	for _, potentialOpponent := range g.Warriors {
		if potentialOpponent.IsAlive() && potentialOpponent.race != warrior.race {
			return true
		}
	}
	return false
}

func (g *Game) PlayWarriorsRound(warrior *Warrior) {
	opponents := g.GetOpponentsFor(warrior)
	notOccupiedOpponentsWeakSpots := g.GetNotOccupiedOpponentsWeakSpotsFor(warrior, opponents)
	if !warrior.InRangeToAny(notOccupiedOpponentsWeakSpots) {
		g.MoveWarriorTowards(warrior, notOccupiedOpponentsWeakSpots)
	}
	opponentsNearby := g.GetOpponentsNearby(warrior, opponents)
	if len(opponentsNearby) > 0 {
		opponentsNearby[0].hitPoints -= warrior.power
	}
}

func (g *Game) GetOpponentsFor(warrior *Warrior) []*Warrior {
	var opponents []*Warrior
	for _, potentialOpponent := range g.Warriors {
		if potentialOpponent.IsAlive() && potentialOpponent.race != warrior.race {
			opponents = append(opponents, potentialOpponent)
		}
	}
	return opponents
}

func (g *Game) GetOtherWarriorsPositions(warrior *Warrior) map[Point]*Warrior {
	warriorsPositions := make(map[Point]*Warrior)
	for _, otherWarrior := range g.Warriors {
		if otherWarrior.IsAlive() && otherWarrior != warrior {
			warriorsPositions[otherWarrior.currentPosition] = otherWarrior
		}
	}
	return warriorsPositions
}

func (g *Game) GetNotOccupiedOpponentsWeakSpotsFor(warrior *Warrior, opponents []*Warrior) map[Point]*Warrior {
	otherWarriorsPositions := g.GetOtherWarriorsPositions(warrior)
	opponentsWeakSpots := make(map[Point]*Warrior)
	for _, opponent := range opponents {
		for _, weakSpot := range opponent.WeakSpots() {
			if isWall := g.WallsMap[weakSpot]; !isWall {
				if _, isOccupiedPosition := otherWarriorsPositions[weakSpot]; !isOccupiedPosition {
					opponentsWeakSpots[weakSpot] = opponent
				}
			}
		}
	}
	return opponentsWeakSpots
}

func (g *Game) GetAllOccupiedLocationsAsideFrom(warrior *Warrior) map[Point]struct{} {
	occupied := make(map[Point]struct{})
	for point, isWall := range g.WallsMap {
		if isWall {
			occupied[point] = struct{}{}
		}
	}
	for _, otherWarrior := range g.Warriors {
		if otherWarrior.IsAlive() && otherWarrior != warrior {
			occupied[otherWarrior.currentPosition] = struct{}{}
		}
	}
	return occupied
}

func (g *Game) GetOpponentsNearby(warrior *Warrior, opponents []*Warrior) []*Warrior {
	var opponentsNearby []*Warrior
	for _, neighbour := range warrior.currentPosition.Neighbours() {
		for _, opponent := range opponents {
			if opponent.currentPosition == neighbour {
				opponentsNearby = append(opponentsNearby, opponent)
			}
		}
	}
	sort.Slice(opponentsNearby, func(i, j int) bool {
		if opponentsNearby[i].hitPoints < opponentsNearby[j].hitPoints {
			return true
		}
		if opponentsNearby[i].currentPosition.y != opponentsNearby[j].currentPosition.y {
			return opponentsNearby[i].currentPosition.y < opponentsNearby[j].currentPosition.y
		} else {
			return opponentsNearby[i].currentPosition.x < opponentsNearby[j].currentPosition.x
		}
	})
	return opponentsNearby
}

func (g *Game) AllElvesAlive() bool {
	for _, warrior := range g.Warriors {
		if warrior.race == Elf && !warrior.IsAlive() {
			return false
		}
	}
	return true
}
