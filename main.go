package main

import (
	"image/color"
	"log"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Body struct {
	x, y       float64
	xspd, yspd float64
	mass       float64
}

var (
	bodies  []Body
	G       = 0.1
	screenW = 800
	screenH = 600
	img     *ebiten.Image
)

func init() {
	var err error
	img, _, err = ebitenutil.NewImageFromFile("assets/body.png")
	if err != nil {
		log.Fatal(err)
	}
}

type Game struct{}

func (g *Game) Update() error {
	for i := range bodies {
		ax, ay := 0.0, 0.0
		for j := range bodies {
			if i == j {
				continue
			}
			dx := bodies[j].x - bodies[i].x
			dy := bodies[j].y - bodies[i].y
			distSq := dx*dx + dy*dy
			dist := math.Sqrt(distSq)
			if dist < 1 {
				continue
			}
			force := G * bodies[j].mass / distSq
			ax += force * dx / dist
			ay += force * dy / dist
		}
		bodies[i].xspd += ax
		bodies[i].yspd += ay
	}

	for i := range bodies {
		bodies[i].x += bodies[i].xspd
		bodies[i].y += bodies[i].yspd
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{16, 16, 20, 255})

	imgW, imgH := img.Bounds().Dx(), img.Bounds().Dy()
	for _, b := range bodies {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(b.x-float64(imgW)/2, b.y-float64(imgH)/2)
		op.Filter = ebiten.FilterNearest
		screen.DrawImage(img, op)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenW, screenH
}

func main() {
	bodies = []Body{
		{x: 400, y: 300, mass: 10000},
		{x: 500, y: 300, yspd: -2.5, mass: 10},
		{x: 300, y: 300, yspd: 2.5, mass: 10},
	}

	ebiten.SetWindowSize(screenW, screenH)
	ebiten.SetWindowTitle("N-body")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
