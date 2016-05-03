package main

import (
	"engo.io/ecs"
	"engo.io/engo"

	"fmt"

	"github.com/effractor/racing/systems"
)

// Scene with the road and the car
type roadScene struct{}

func (*roadScene) Type() string { return "RacingGame" }

func (*roadScene) Preload() {
	fmt.Println("roadScene->Preload()")
	engo.Files.AddFromDir("data", false)
}

func (*roadScene) Setup(world *ecs.World) {
	fmt.Println("roadScene->Setup()")

	// Adding the RenderSystem
	world.AddSystem(&engo.RenderSystem{})
	world.AddSystem(&engo.CollisionSystem{})
	world.AddSystem(&systems.ControlSystem{})
	world.AddSystem(&systems.BorderSpawnSystem{})
	world.AddSystem(&systems.RoadSpeedSystem{})

	/*engo.Mailbox.Listen("CollisionMessage", func(message engo.Message) {
		entity, to := message.(engo.CollisionMessage)
		fmt.Printf("Collision: %v and %v\n", entity, to)
	})*/

	// Adding a car
	systems.NewCar(world, engo.Point{engo.Width() / 2, 300}, "car.png", true)
}

// Main function
func main() {
	fmt.Println("Racing game")

	opts := engo.RunOptions{
		Title:  "Racing game",
		Width:  800,
		Height: 600,
	}
	engo.Run(opts, &roadScene{})
}
