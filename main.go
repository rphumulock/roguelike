package main

import (
	_ "image/png"
	"log"

	"github.com/bytearena/ecs"
	"github.com/hajimehoshi/ebiten/v2"
)

// Game holds all data the entire game will need.
type Game struct {
	Map       GameMap
	World     *ecs.Manager
	WorldTags map[string]ecs.Tag
}

// NewGame creates a new Game Object and initializes the data
func NewGame() *Game {
	g := &Game{}
	world, tags := InitializeWorld()
	g.Map = NewGameMap()
	g.WorldTags = tags
	g.World = world
	return g
}

// Update is called each tic.
func (g *Game) Update() error {
	TryMovePlayer(g)
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	//Draw the Map
	level := g.Map.Dungeons[0].Levels[0]
	level.DrawLevel(screen)
	ProcessRenderables(g, level, screen)
}

// Layout will return the screen dimensions.
func (g *Game) Layout(w, h int) (int, int) {
	gd := NewGameData()
	return gd.TileWidth * gd.ScreenWidth, gd.TileHeight * gd.ScreenHeight
}

func main() {

	g := NewGame()
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)

	ebiten.SetWindowTitle("Tower")

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
