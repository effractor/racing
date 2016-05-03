package systems

import (
	"engo.io/ecs"
	"engo.io/engo"

	"fmt"
)

type Car struct {
	ecs.BasicEntity
	engo.CollisionComponent
	engo.RenderComponent
	engo.SpaceComponent
}

func NewCar(world *ecs.World, position engo.Point, textureFile string, isPlayer bool) {
	fmt.Printf("Position: %v, texture: %s, isPlayer %v\n", position, textureFile, isPlayer)
	texture := engo.Files.Image(textureFile)

	if texture == nil {
		return
	}

	car := Car{BasicEntity: ecs.NewBasic()}

	car.SpaceComponent = engo.SpaceComponent{
		Position: position,
		Width:    73,
		Height:   157,
	}

	car.RenderComponent = engo.NewRenderComponent(texture, engo.Point{1, 1})
	car.CollisionComponent = engo.CollisionComponent{Solid: true, Main: isPlayer}

	for _, system := range world.Systems() {
		switch sys := system.(type) {
		case *engo.RenderSystem:
			sys.Add(&car.BasicEntity, &car.RenderComponent, &car.SpaceComponent)
		case *engo.CollisionSystem:
			sys.Add(&car.BasicEntity, &car.CollisionComponent, &car.SpaceComponent)
		case *RoadSpeedSystem:
			if !isPlayer {
				sys.Add(&car.BasicEntity, &car.SpaceComponent)
			}
		case *ControlSystem:
			if isPlayer {
				sys.Add(&car.BasicEntity, &car.SpaceComponent)
			}
		}
	}
}
