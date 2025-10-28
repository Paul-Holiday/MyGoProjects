package main

import (
	"fmt"
	"log"
	"time"
)

// КОМПОЗИЦИЯ ИНТЕРФЕЙСОВ

type Vehicle interface {
	Mover
	Fueler
	Charger
}

type Mover interface {
	Move()
	Stop()
}

type Fueler interface {
	Refuel(float64) error
	GetFuelLevel() float64
}

type Charger interface {
	Charge(float64) error
	GetBatteryLevel() float64
}

// ==СТРУКТУРЫ ВИДОВ ТРАНСПОРТА==
type Car struct {
	model     string
	tankCap   float64 // литры
	fuelLevel float64 // литры
	fuelRate  float64 // литры/100км
	locationX float64 // км
	moving    bool
	startTime time.Time
}

type Hybrid struct {
	model      string
	tankCap    float64 // литры
	batCap     float64 // кВт*ч
	fuelLevel  float64 // литры
	batLevel   float64 // кВт*ч
	fuelRate   float64 // литры/100км
	chargeRate float64 // кВт*ч/100км
	locationX  float64 // км
	moving     bool
	startTime  time.Time
}

type Escooter struct {
	model      string
	batСap     float64 // кВт*ч
	batLevel   float64 // кВт*ч
	chargeRate float64 // кВт*ч/100км
	locationX  float64 // км
	moving     bool
	startTime  time.Time
}

// ==МЕТОДЫ CAR==
func (c *Car) Move() {
	if c.moving {
		log.Println("Машина уже двигается.")
		return
	}
	if c.fuelLevel == 0 {
		log.Printf("У машины %s нет топлива, невозможно начать движение.", c.model)
		return
	}
	c.moving = true
	c.startTime = time.Now()
	log.Printf("Машина %s начала движение.", c.model)
}

func (c *Car) Stop() {
	if !c.moving {
		log.Println("Машина уже стоит.")
		return
	}
	c.moving = false
	log.Printf("Машина %s остановилась.\n", c.model)
}

func (c *Car) Refuel(amount float64) error {
	if c.moving {
		return fmt.Errorf("error: car %s is moving, stop to refuel", c.model)
	}
	if amount <= 0 {
		return fmt.Errorf("error: amount of fuel must be positive")
	}

	emptySpace := c.tankCap - c.fuelLevel
	// если места в баке больше, чем планируется заправить, то просто добавляем топлива
	if emptySpace > amount {
		c.fuelLevel += amount
		return nil
	}
	c.fuelLevel = c.tankCap
	return nil
}

func (c *Car) GetFuelLevel() float64 {
	return c.fuelLevel
}

// ==МЕТОДЫ HYBRID==

func (h *Hybrid) Move() {
	if h.moving {
		log.Println("Гибрид уже двигается.")
		return
	}
	if h.fuelLevel == 0 && h.batLevel == 0 {
		log.Printf("У гибрида %s нет топлива и заряда, невозможно начать движение.\n", h.model)
		return
	}
	h.moving = true
	h.startTime = time.Now()
	log.Printf("Гибрид %s начал движение.\n", h.model)
}

func (h *Hybrid) Stop() {
	if !h.moving {
		log.Printf("Гибрид %s уже стоит.\n", h.model)
		return
	}
	log.Printf("Гибрид %s остановился.\n", h.model)
	h.moving = false
}

func (h *Hybrid) Refuel(amount float64) error {
	if h.moving {
		return fmt.Errorf("error: hybrid %s is moving, stop to refuel", h.model)
	}
	if amount <= 0 {
		return fmt.Errorf("error: amount of fuel must be positive")
	}

	emptySpace := h.tankCap - h.fuelLevel
	// если места в баке больше, чем планируется заправить, то просто добавляем топлива
	if emptySpace > amount {
		h.fuelLevel += amount
		return nil
	}
	h.fuelLevel = h.tankCap
	return nil
}

func (h *Hybrid) Charge(energy float64) error {
	if h.moving {
		return fmt.Errorf("error: hybrid %s is moving, stop to charge", h.model)
	}
	if energy <= 0 {
		return fmt.Errorf("error: charging energy must be positive")
	}

	emptySpace := h.batCap - h.batLevel

	if emptySpace > energy {
		h.batLevel += energy
		return nil
	}
	h.batLevel = h.batCap
	return nil
}

func (h *Hybrid) GetFuelLevel() float64 {
	return h.fuelLevel
}

func (h *Hybrid) GetBatteryLevel() float64 {
	return h.batLevel
}

// ==МЕТОДЫ ESCOOTER==

func (esc *Escooter) Move() {
	if esc.moving {
		log.Println("Электросамокат уже двигается.")
		return
	}
	if esc.batLevel == 0 {
		log.Printf("У электросамоката %s нет заряда, невозможно начать движение.\n", esc.model)
		return
	}
	esc.moving = true
	esc.startTime = time.Now()
	log.Printf("Электросамокат %s начал движение.\n", esc.model)
}

func (esc *Escooter) Stop() {
	if !esc.moving {
		log.Printf("Электросамокат %s уже стоит.\n", esc.model)
		return
	}
	log.Printf("Электросамокат %s остановился.\n", esc.model)
	esc.moving = false
}

func (esc *Escooter) Charge(energy float64) error {
	if esc.moving {
		return fmt.Errorf("error: scooter %s is moving, stop to charge", esc.model)
	}
	if energy <= 0 {
		return fmt.Errorf("error: charging energy must be positive")
	}

	emptySpace := esc.batСap - esc.batLevel

	if emptySpace > energy {
		esc.batLevel += energy
		return nil
	}
	esc.batLevel = esc.batСap
	return nil
}

func (esc *Escooter) GetBatteryLevel() float64 {
	return esc.batLevel
}

func main() {
	log.Println()
	fmt.Println()
}
