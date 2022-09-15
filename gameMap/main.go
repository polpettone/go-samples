package main

import (
	"fmt"
	"strings"
)

const WHALE_RUNE = '🐋'
const WHALE_STRING = "🐋"
const LEFT_WALL_PIECE = "|"
const RIGHT_WALL_PIECE = "|"
const HORIZONTAL_WALL_PIECE = "_"
const EMPTY_FIELD = " "

const GOPHER = ""
const WATER = "~"
const FIELD = " "
const BOX = ""
const GITHUB = ""

const ELEPHANT = "🐘"

const PENGUIN = "🐧"

const HEART = "❤️'"

const BALLON = "🎈"

const FIELD_WIDTH = 2

type Map struct {
	MaxX int
	MaxY int

	Fields [][]Field
}

type Element struct {
	Symbol string
	ID     int
}

type Field struct {
	Element Element
}

func initMap() Map {
	gameMap := Map{
		MaxX: 10,
		MaxY: 10,
	}

	fields := make([][]Field, gameMap.MaxX+1)
	for n := range fields {
		fields[n] = make([]Field, gameMap.MaxY+1)
	}

	fieldSymbol := strings.Repeat(FIELD, FIELD_WIDTH)

	id := 0
	for x := 0; x < gameMap.MaxX; x++ {
		for y := 0; y < gameMap.MaxY; y++ {
			fields[x][y] = Field{Element{Symbol: fieldSymbol, ID: id}}
			id++
		}
	}

	fields[1][4].Element.Symbol = WHALE_STRING + strings.Repeat(FIELD, FIELD_WIDTH-len(WHALE_STRING)+2)
	fields[1][3].Element.Symbol = HEART + strings.Repeat(FIELD, FIELD_WIDTH-len(HEART)+5)
	fields[2][3].Element.Symbol = GOPHER + strings.Repeat(FIELD, FIELD_WIDTH-len(GOPHER)+2)

	fields[3][4].Element.Symbol = BALLON + strings.Repeat(FIELD, FIELD_WIDTH-len(BALLON)+2)
	fields[3][3].Element.Symbol = WHALE_STRING + strings.Repeat(FIELD, FIELD_WIDTH-len(WHALE_STRING)+2)

	gameMap.Fields = fields

	makeLeftWall(&gameMap)
	makeRightWall(&gameMap)

	makeTopWall(&gameMap)
	makeBottomWall(&gameMap)

	return gameMap
}

func makeLeftWall(m *Map) {
	for y := 0; y < m.MaxY; y++ {
		m.Fields[0][y].Element.Symbol = LEFT_WALL_PIECE + strings.Repeat(FIELD, FIELD_WIDTH-len(LEFT_WALL_PIECE)+3)
	}
}

func makeRightWall(m *Map) {
	for y := 0; y < m.MaxY; y++ {
		m.Fields[m.MaxX-1][y].Element.Symbol = RIGHT_WALL_PIECE + strings.Repeat(FIELD, FIELD_WIDTH-len(RIGHT_WALL_PIECE)+3)
	}
}

func makeTopWall(m *Map) {
	for x := 0; x < m.MaxX; x++ {
		m.Fields[x][m.MaxY-1].Element.Symbol = strings.Repeat(HORIZONTAL_WALL_PIECE, FIELD_WIDTH)
	}
}

func makeBottomWall(m *Map) {
	for x := 0; x < m.MaxX; x++ {
		m.Fields[x][0].Element.Symbol = strings.Repeat(HORIZONTAL_WALL_PIECE, FIELD_WIDTH)
	}
}

func (m Map) Print() string {
	p := ""
	for y := m.MaxY - 1; y >= 0; y-- {
		for x := 0; x < m.MaxX; x++ {
			p += fmt.Sprintf(string(m.Fields[x][y].Element.Symbol))
		}
		p += "\n"
	}
	p += "\n"
	return p
}

func main() {

	gameMap := initMap()
	fmt.Printf("%s", gameMap.Print())

	printLenght(WHALE_STRING)
	printLenght(HORIZONTAL_WALL_PIECE)
	printLenght(LEFT_WALL_PIECE)
	printLenght(HEART)
	printLenght(PENGUIN)
	printLenght(BALLON)
	printLenght(GOPHER)
	printLenght(GITHUB)
	printLenght(FIELD)

	fmt.Println(strings.Repeat("-", 20))
	printEachCharacter(WHALE_STRING)

	fmt.Println(strings.Repeat("-", 20))
	printEachCharacter(GOPHER)

	fmt.Println(strings.Repeat("-", 20))
	printEachCharacter("12")

}

func printEachCharacter(s string) {
	for _, c := range s {
		fmt.Println(c)
	}
}

func printLenght(s string) {
	fmt.Printf("%s lenght %d \n", s, len(s))
}
