package task10

import (
	"fmt"
	"sort"

	"github.com/karolgil/AdventOfCode/2018/utils"
)

func Solution1(inputFile string) string {
	lines, _ := utils.ReadLinesFrom(inputFile)
	cartsMap, carts := CreateCartsMap(lines)
	for {
		//cartsMap.ShowWithCarts(carts)
		cartsMap.MoveCarts(carts)
		if crashed := carts.AnyCrashed(); crashed != nil {
			return fmt.Sprintf("%d,%d", crashed.x, crashed.y)
		}
	}
}

func CreateCartsMap(lines []string) (CartsMap, Carts) {
	var carts Carts
	cartsMap := make(CartsMap, len(lines))
	for y, line := range lines {
		cartsMap[y] = make(LineMap, len(line))
		for x, char := range line {
			switch string(char) {
			case "^":
				carts = append(carts, NewCart(x, y, 0, -1, "^"))
				cartsMap[y][x] = "|"
			case "v":
				carts = append(carts, NewCart(x, y, 0, 1, "v"))
				cartsMap[y][x] = "|"
			case ">":
				carts = append(carts, NewCart(x, y, 1, 0, ">"))
				cartsMap[y][x] = "-"
			case "<":
				carts = append(carts, NewCart(x, y, -1, 0, "<"))
				cartsMap[y][x] = "-"
			default:
				cartsMap[y][x] = string(char)
			}
		}
	}
	return cartsMap, carts
}

func Solution2(inputFile string) string {
	lines, _ := utils.ReadLinesFrom(inputFile)
	cartsMap, carts := CreateCartsMap(lines)
	for {
		// cartsMap.ShowWithCarts(carts) // Uncomment to see carts progressing
		cartsMap.MoveCarts(carts)
		if oneLeftAlive := carts.IsOneLeftAlive(); oneLeftAlive {
			stillAlive := carts.FindAnyAlive()
			return fmt.Sprintf("%d,%d", stillAlive.x, stillAlive.y)
		}
	}
}

type LineMap map[int]string

func (lm LineMap) SortedKeys() []int {
	keys := make([]int, len(lm))
	for key := range lm {
		keys = append(keys, key)
	}
	sort.Ints(keys)
	return keys
}

type CartsMap map[int]LineMap

func (cm CartsMap) SortedKeys() []int {
	var keys []int
	for key := range cm {
		keys = append(keys, key)
	}
	sort.Ints(keys)
	return keys
}

func (cm CartsMap) ShowWithCarts(carts Carts) {
	fmt.Println()
	keys := cm.SortedKeys()
	for y := range keys {
		for x := range cm[y].SortedKeys() {
			if cart := carts.FindNotCrashed(x, y); cart != nil {
				fmt.Print(cart.representation)
			} else {
				fmt.Print(string(cm[y][x]))
			}
		}
		fmt.Print("\n")
	}
}

func (cm CartsMap) MoveCarts(carts Carts) {
	carts.SortByYThenX()
	for _, cart := range carts {
		if cart.crashed {
			continue
		}
		newX := cart.x + cart.vx
		newY := cart.y + cart.vy
		cartAtSamePosition := carts.FindNotCrashed(newX, newY)
		cart.x = newX
		cart.y = newY
		if cartAtSamePosition != nil {
			cart.crashed = true
			cartAtSamePosition.crashed = true
			continue
		}

		switch cm[cart.y][cart.x] {
		case "/":
			cart.vx, cart.vy = -cart.vy, -cart.vx
			break
		case "\\":
			cart.vx, cart.vy = cart.vy, cart.vx
			break
		case "+":
			switch cart.state {
			case 0:
				cart.vx, cart.vy = cart.vy, -cart.vx
				break
			case 1:
				break
			case 2:
				cart.vx, cart.vy = -cart.vy, cart.vx
				break
			}
			cart.state = (cart.state + 1) % 3
		}
	}
}

type Carts []*Cart

func (cs Carts) SortByYThenX() {
	sort.Slice(cs, func(i, j int) bool {
		if cs[i].y < cs[j].y {
			return true
		} else if cs[i].y == cs[j].y {
			if cs[i].x < cs[j].x {
				return true
			}
		}
		return false
	})
}

func (cs Carts) FindNotCrashed(x, y int) *Cart {
	for _, cart := range cs {
		if cart.x == x && cart.y == y && !cart.crashed {
			return cart
		}
	}
	return nil
}

func (cs Carts) Show() {
	for _, cart := range cs {
		fmt.Printf("Cart{x=%d, y=%d, vx=%d, vy=%d, state=%d, crashed=%v}\n", cart.x, cart.y, cart.vx, cart.vy, cart.state, cart.crashed)
	}
}

func (cs Carts) AnyCrashed() *Cart {
	for _, cart := range cs {
		if cart.crashed {
			return cart
		}
	}
	return nil
}

func (cs Carts) FindAnyAlive() *Cart {
	for _, cart := range cs {
		if !cart.crashed {
			return cart
		}
	}
	return nil
}

func (cs Carts) IsOneLeftAlive() bool {
	alive := 0
	for _, cart := range cs {
		if !cart.crashed {
			alive += 1
		}
	}
	return alive == 1
}

type Cart struct {
	x              int
	y              int
	vx             int
	vy             int
	state          int
	crashed        bool
	representation string
}

func NewCart(x, y, vx, vy int, representation string) *Cart {
	return &Cart{
		x:              x,
		y:              y,
		vx:             vx,
		vy:             vy,
		state:          0,
		crashed:        false,
		representation: representation,
	}
}
