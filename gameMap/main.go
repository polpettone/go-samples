package main

import "fmt"

type Map struct {
	MaxX int
	MaxY int

	Fields [][]Field
}

type Element struct {
	Symbol rune
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

	id := 0
	for x := 0; x < gameMap.MaxX; x++ {
		for y := 0; y < gameMap.MaxY; y++ {
			fields[x][y] = Field{Element{Symbol: '_', ID: id}}
			id++
		}
	}
	gameMap.Fields = fields
	return gameMap
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

}
