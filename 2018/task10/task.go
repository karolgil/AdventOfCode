package task10

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/otiai10/gosseract"

	"golang.org/x/tools/container/intsets"

	"github.com/karolgil/AdventOfCode/2018/utils"
)

func Solution(inputFile string) (string, int, error) {
	lines, err := utils.ReadLinesFrom(inputFile)
	if err != nil {
		return "", 0, err
	}
	points := NewPoints(lines)
	text := ""
	timePassed := 0
	for text == "" {
		timePassed++
		points.Rotate()
		if points.HeightLowerThan(20) {
			points.SavePNG(inputFile)
			text = points.Text(inputFile)
		}
	}
	fmt.Printf("Text found: %s\n", text)
	return text, timePassed, nil
}

type Points struct {
	points                 []*Point
	minX, minY, maxX, maxY int
}

func (ps *Points) HeightLowerThan(limit int) bool {
	return ps.maxY-ps.minY < limit
}

func (ps *Points) SavePNG(fileName string) {
	imgRect := image.Rect(ps.minX, ps.minY, ps.maxX, ps.maxY)
	img := image.NewNRGBA(imgRect)
	for _, p := range ps.points {
		img.Set(p.x, p.y, color.NRGBA{
			R: 0,
			G: 0,
			B: 0,
			A: 255,
		})
	}
	f, err := os.Create(fmt.Sprintf("%s.png", fileName))
	if err != nil {
		log.Fatal(err)
	}
	enc := &png.Encoder{
		CompressionLevel: png.NoCompression,
	}

	if err := enc.Encode(f, img); err != nil {
		f.Close()
		log.Fatal(err)
	}

	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}

func (ps *Points) Rotate() {
	ps.minX = intsets.MaxInt
	ps.minY = intsets.MaxInt
	ps.maxX = intsets.MinInt
	ps.maxY = intsets.MinInt
	for _, p := range ps.points {
		newX, newY := p.Rotate()
		ps.MoveBoundaries(newX, newY)
	}
}

func (ps *Points) MoveBoundaries(newX int, newY int) {
	if newX < ps.minX {
		ps.minX = newX - 4
	}
	if newX > ps.maxX {
		ps.maxX = newX + 4
	}
	if newY < ps.minY {
		ps.minY = newY - 4
	}
	if newY > ps.maxY {
		ps.maxY = newY + 4
	}
}

func (ps *Points) Text(fileName string) string {
	client := gosseract.NewClient()
	defer client.Close()
	client.SetImage(fmt.Sprintf("%s.png", fileName))
	text, _ := client.Text()
	// gosseract/tesseract are not handling blacklists properly, so we need to switch common number/alpha problems
	// we assume that we can have only alphabet uppercase characters in output
	text = strings.Replace(text, "0", "O", -1)
	text = strings.Replace(text, "6", "G", -1)
	text = strings.Replace(text, "2", "Z", -1)
	return text
}

func NewPoints(inputLines []string) *Points {
	ps := &Points{
		minX: intsets.MaxInt,
		minY: intsets.MaxInt,
		maxX: intsets.MinInt,
		maxY: intsets.MinInt,
	}
	pointRegex := regexp.MustCompile(`^position=<\s*(.+),\s*(.+)> velocity=<\s*(.+),\s*(.+)>$`)
	for _, line := range inputLines {
		pointDetails := pointRegex.FindAllStringSubmatch(line, -1)
		x, y, xVel, yVel := getPointDetails(pointDetails[0])
		ps.MoveBoundaries(x, y)
		ps.points = append(ps.points, NewPoint(x, y, xVel, yVel))
	}
	return ps
}

func getPointDetails(pointDetails []string) (x int, y int, xVel int, yVel int) {
	x, _ = strconv.Atoi(pointDetails[1])
	y, _ = strconv.Atoi(pointDetails[2])
	xVel, _ = strconv.Atoi(pointDetails[3])
	yVel, _ = strconv.Atoi(pointDetails[4])
	return
}

type Point struct {
	x, y, xVel, yVel int
}

func (p *Point) Rotate() (int, int) {
	p.x += p.xVel
	p.y += p.yVel
	return p.x, p.y
}

func NewPoint(x, y, xVel, yVel int) *Point {
	return &Point{
		x:    x,
		y:    y,
		xVel: xVel,
		yVel: yVel,
	}
}
