package main

import (
	"fmt"
	"time"
)

type Mover interface {
	Move()
	Stop()
}

type Car struct {
	Model string
}

type Bicycle struct {
	Model string
}

type Person struct {
	Name string
}

func (c *Car) Move() {
	fmt.Printf("Машина %s поехала!\n", c.Model)
}

func (c Car) Stop() {
	fmt.Printf("Машина %s остановилась!\n", c.Model)
}

func (b *Bicycle) Move() {
	fmt.Printf("Велосипед %s поехал!\n", b.Model)
}

func (b Bicycle) Stop() {
	fmt.Printf("Велосипед %s остановился!\n", b.Model)
}

func (p *Person) Move() {
	fmt.Printf("%s поехал(а)!\n", p.Name)
}

func (p Person) Stop() {
	fmt.Printf("%s остановился(ась)!\n", p.Name)
}

func MoveThenStop(m Mover) {
	m.Move()
	time.Sleep(1 * time.Second)
	m.Stop()
}

func main() {
	car := &Car{"Toyota"}
	bicycle := &Bicycle{"PIVOT"}
	person := &Person{"Алексей"}

	movers := []Mover{car, bicycle, person}

	for _, person := range movers {
		MoveThenStop(person)
	}
}
