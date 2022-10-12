package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten"
)

const (
	screenWidth  = 1000
	screenHeight = 756

	keyboardWidth  = 800
	keyboardHeigth = 300
	capSize        = 45
)

type Cap struct {
	Image *ebiten.Image
	X     int
	Y     int
}

type Game struct {
	Caps            []Cap
	KeyboardImage   *ebiten.Image
	BackgroundImage *ebiten.Image
}

func (g *Game) Update(screen *ebiten.Image) error {

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {

	for _, c := range g.Caps {
		drawCap(c, g.KeyboardImage)
	}

	op := &ebiten.DrawImageOptions{}
	screen.DrawImage(g.BackgroundImage, op)

	op.GeoM.Translate(float64(100), float64(200))
	screen.DrawImage(g.KeyboardImage, op)

}

func drawCap(c Cap, keyboardImage *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(c.X), float64(c.Y))
	keyboardImage.DrawImage(c.Image, op)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Visual Keyboard")

	backgroundImage, err := ebiten.NewImage(screenWidth, screenHeight, ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
		return
	}
	backgroundImage.Fill(color.RGBA{240, 255, 240, 0xff})

	keyboardImage, err := ebiten.NewImage(keyboardWidth, keyboardHeigth, ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
		return
	}
	keyboardImage.Fill(color.RGBA{0, 0, 0, 0xff})

	capImage, err := ebiten.NewImage(capSize, capSize, ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
		return
	}
	capImage.Fill(color.RGBA{255, 255, 255, 0xff})

	c := Cap{
		Image: capImage,
		X:     0,
		Y:     0,
	}

	c1 := Cap{
		Image: capImage,
		X:     45,
		Y:     0,
	}

	g := &Game{
		Caps:            []Cap{c, c1},
		KeyboardImage:   keyboardImage,
		BackgroundImage: backgroundImage,
	}

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
