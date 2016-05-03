package systems

import (
	"engo.io/ecs"
	"engo.io/engo"
)

// The border
type Border struct {
	ecs.BasicEntity
	engo.CollisionComponent
	engo.RenderComponent
	engo.SpaceComponent
}

func NewBorder(world *ecs.World, position engo.Point) {
	texture := engo.Files.Image("rock.png")

	border := Border{BasicEntity: ecs.NewBasic()}
	border.RenderComponent = engo.NewRenderComponent(texture, engo.Point{1, 1})
	border.SpaceComponent = engo.SpaceComponent{
		Position: position,
		Width:    texture.Width() * border.RenderComponent.Scale().X,
		Height:   texture.Height() * border.RenderComponent.Scale().Y,
	}

	border.CollisionComponent = engo.CollisionComponent{Solid: true}

	for _, system := range world.Systems() {
		switch sys := system.(type) {
		case *engo.RenderSystem:
			sys.Add(&border.BasicEntity, &border.RenderComponent, &border.SpaceComponent)
		case *engo.CollisionSystem:
			sys.Add(&border.BasicEntity, &border.CollisionComponent, &border.SpaceComponent)
		case *RoadSpeedSystem:
			sys.Add(&border.BasicEntity, &border.SpaceComponent)
		}
	}
}
