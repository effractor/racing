package systems

import (
	"engo.io/ecs"
	"engo.io/engo"

	"math/rand"
)

// Road generation system
type BorderSpawnSystem struct {
	world *ecs.World
	count int
	cars  int
}

func (road *BorderSpawnSystem) New(w *ecs.World) {
	road.world = w
}

func (*BorderSpawnSystem) Remove(ecs.BasicEntity) {}

func (road *BorderSpawnSystem) Update(dt float32) {
	road.count += 1
	if road.count < 10 {
		return
	}

	road.count = 0
	center := engo.Width() / 2
	width := float32(300)

	delta := float32(50)

	left := center - delta
	right := left + width

	NewBorder(road.world, engo.Point{left, -32})
	NewBorder(road.world, engo.Point{right, -32})

	road.cars += 1
	if road.cars < 20 {
		return
	}

	road.cars = 0

	txId := rand.Float32()
	tx := "police.png"
	if txId < 0.3 {
		tx = "purple.png"
	} else if 0.3 <= txId && txId < 0.6 {
		tx = "red.png"
	}

	delta = width * rand.Float32()

	x := left + delta
	if x > (left + width - 80) {
		x = left + width - 80
	}

	NewCar(road.world, engo.Point{x, -100}, tx, false)
}
