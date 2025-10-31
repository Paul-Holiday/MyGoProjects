package main

import (
	"fmt"
	"log"
)

// КОМПОЗИЦИЯ ИНТЕРФЕЙСОВ

type FuelVehicle interface {
	Mover
	Fueler
}

type ElectricVehicle interface {
	Mover
	Charger
}

type HybridVehicle interface {
	Mover
	Fueler
	Charger
}

type TestDrivable interface {
	Mover
	Locator
	TripCalculator
	EnergyConsumer
}

type Mover interface {
	Move() error
	Stop() error
	IsMoving() bool
}

type Fueler interface {
	Refuel(float64) error
	GetFuelLevel() float64
	GetTankCapacity() float64
}

type Charger interface {
	Charge(float64) error
	GetBatteryLevel() float64
	GetBatteryCapacity() float64
}

type Locator interface {
	GetLocation() float64
	GetDistanceTravelled() float64
	ChangeLocation(distance float64)
	IncreaseDistanceTravelled(distance float64)
}

type TripCalculator interface {
	CalculateTrip(distance float64) (fuel, energy float64, err error)
}

type EnergyConsumer interface {
	ConsumeEnergy(fuel, energy float64) error
}

// ==СТРУКТУРЫ ВИДОВ ТРАНСПОРТА==
type Car struct {
	model             string
	tankCap           float64 // литры
	fuelLevel         float64 // литры
	fuelRate          float64 // литры/100км
	locationX         float64 // км
	distanceTravelled float64
	isMoving          bool
}

type Hybrid struct {
	model             string
	tankCap           float64 // литры
	batCap            float64 // кВт*ч
	fuelLevel         float64 // литры
	batLevel          float64 // кВт*ч
	fuelRate          float64 // литры/100км
	chargeRate        float64 // кВт*ч/100км
	locationX         float64 // км
	distanceTravelled float64 // км
	isMoving          bool
}

type Escooter struct {
	model             string
	batCap            float64 // кВт*ч
	batLevel          float64 // кВт*ч
	chargeRate        float64 // кВт*ч/100км
	locationX         float64 // км
	distanceTravelled float64 // км
	isMoving          bool
}

// КОНСТРУКТОРЫ

func NewCar(model string, tankCap, fuelLevel, fuelRate, locationX, distanceTravelled float64, isMoving bool) *Car {
	return &Car{model, tankCap, fuelLevel, fuelRate, locationX, distanceTravelled, isMoving}
}

func NewHybrid(model string, tankCap, batCap, fuelLevel, batLevel, fuelRate, chargeRate, locationX, distanceTravelled float64, isMoving bool) *Hybrid {
	return &Hybrid{model, tankCap, batCap, fuelLevel, batLevel, fuelRate, chargeRate, locationX, distanceTravelled, isMoving}
}

func NewEscooter(model string, batCap, batLevel, chargeRate, locationX, distanceTravelled float64, isMoving bool) *Escooter {
	return &Escooter{model, batCap, batLevel, chargeRate, locationX, distanceTravelled, isMoving}
}

// ==МЕТОДЫ CAR==
func (c *Car) Move() error {
	if c.isMoving {
		return fmt.Errorf("error: car %s is already moving", c.model)

	}
	if c.fuelLevel == 0 {
		return fmt.Errorf("error: no fuel in %s, refuel to move", c.model)
	}
	c.isMoving = true
	log.Printf("Машина %s начала движение.", c.model)
	return nil
}

func (c *Car) Stop() error {
	if !c.isMoving {
		return fmt.Errorf("error: car %s is already stopped", c.model)
	}
	c.isMoving = false
	log.Printf("Машина %s остановилась.\n", c.model)
	return nil
}

func (c *Car) IsMoving() bool {
	return c.isMoving
}

func (c *Car) Refuel(amount float64) error {
	if c.isMoving {
		return fmt.Errorf("error: car %s is moving, stop to refuel", c.model)
	}
	var err error
	c.fuelLevel, err = addUntilFull(c.fuelLevel, c.tankCap, amount)

	return err
}

func (c *Car) GetFuelLevel() float64 {
	return c.fuelLevel
}

func (c *Car) GetTankCapacity() float64 {
	return c.tankCap
}

func (c *Car) GetLocation() float64 {
	return c.locationX
}

func (c *Car) GetDistanceTravelled() float64 {
	return c.distanceTravelled
}

func (c *Car) ChangeLocation(distance float64) {
	c.locationX += distance
}

func (c *Car) IncreaseDistanceTravelled(distance float64) {
	c.distanceTravelled += distance
}

func (c *Car) String() string {
	return fmt.Sprintf("Car %s info:\nTank Capacity: %0.1f\nFuel Level: %0.1f\nFuel Rate: %0.1f\nCurrent location: %0.1f\nLast travelled distance: %0.1f\nIs this car moving? - %v\n",
		c.model, c.tankCap, c.fuelLevel, c.fuelRate, c.locationX, c.distanceTravelled, c.isMoving)
}

func (c *Car) CalculateTrip(distance float64) (fuel, energy float64, err error) {
	if c.isMoving {
		return 0, 0, fmt.Errorf("error: car %s is moving", c.model)
	}
	if distance <= 0 {
		return 0, 0, fmt.Errorf("error: distance must be positive")
	}

	return distance * c.fuelRate, 0, nil
}

func (c *Car) ConsumeEnergy(fuel, energy float64) error {
	if fuel > c.tankCap {
		return fmt.Errorf("error: cannot consume that much fuel due to max tank capacity (%0.1f)", c.tankCap)
	}
	if fuel > c.fuelLevel {
		return fmt.Errorf("error: not enough fuel in car %s", c.model)
	}
	c.fuelLevel -= fuel
	return nil
}

// ==МЕТОДЫ HYBRID==

func (h *Hybrid) Move() error {
	if h.isMoving {
		return fmt.Errorf("error: hybrid %s is already moving", h.model)
	}
	if h.fuelLevel == 0 && h.batLevel == 0 {
		return fmt.Errorf("error: no fuel or charge in hybrid %s", h.model)
	}
	h.isMoving = true
	log.Printf("Гибрид %s начал движение.\n", h.model)
	return nil
}

func (h *Hybrid) Stop() error {
	if !h.isMoving {
		return fmt.Errorf("error: hybrid %s is already stopped", h.model)
	}
	log.Printf("Гибрид %s остановился.\n", h.model)
	h.isMoving = false
	return nil
}

func (h *Hybrid) IsMoving() bool {
	return h.isMoving
}

func (h *Hybrid) Refuel(amount float64) error {
	if h.isMoving {
		return fmt.Errorf("error: hybrid %s is moving, stop to refuel", h.model)
	}
	var err error
	h.fuelLevel, err = addUntilFull(h.fuelLevel, h.tankCap, amount)

	return err
}

func (h *Hybrid) Charge(energy float64) error {
	if h.isMoving {
		return fmt.Errorf("error: hybrid %s is moving, stop to charge", h.model)
	}
	var err error
	h.batLevel, err = addUntilFull(h.batLevel, h.batCap, energy)

	return err
}

func (h *Hybrid) GetFuelLevel() float64 {
	return h.fuelLevel
}

func (h *Hybrid) GetTankCapacity() float64 {
	return h.tankCap
}

func (h *Hybrid) GetBatteryLevel() float64 {
	return h.batLevel
}

func (h *Hybrid) GetBatteryCapacity() float64 {
	return h.batCap
}

func (h *Hybrid) GetLocation() float64 {
	return h.locationX
}

func (h *Hybrid) GetDistanceTravelled() float64 {
	return h.distanceTravelled
}

func (h *Hybrid) ChangeLocation(distance float64) {
	h.locationX += distance
}

func (h *Hybrid) IncreaseDistanceTravelled(distance float64) {
	h.distanceTravelled += distance
}

func (h *Hybrid) String() string {
	return fmt.Sprintf("Hybrid %s info:\nTank Capacity: %0.1f\nBattery capacity: %0.1f\nFuel Level: %0.1f\nBattery level: %0.1f\nFuel Rate: %0.1f\nCharge Rate: %0.1f\nCurrent location: %0.1f\nLast travelled distance: %0.1f\nIs this hybrid moving? - %v\n",
		h.model, h.tankCap, h.batCap, h.fuelLevel, h.batLevel, h.fuelRate, h.chargeRate, h.locationX, h.distanceTravelled, h.isMoving)
}

func (h *Hybrid) CalculateTrip(distance float64) (fuel, energy float64, err error) {
	if h.isMoving {
		return 0, 0, fmt.Errorf("error: hybrid %s is moving", h.model)
	}
	if distance <= 0 {
		return 0, 0, fmt.Errorf("error: distance must be positive")
	}

	if distance*h.chargeRate < h.batLevel {
		return 0, distance * h.chargeRate, nil
	}
	fuel = (distance - h.batLevel/h.chargeRate) * h.fuelRate // (всё расстояние - расстояние которое проехали на батарее) * расход топлива на 1км
	return fuel, h.batLevel, nil
}

func (h *Hybrid) ConsumeEnergy(fuel, energy float64) error {
	if energy > h.batCap {
		return fmt.Errorf("error: cannot consume that much energy due to max battery capacity (%0.1f)", h.batCap)
	}
	if fuel > h.tankCap {
		return fmt.Errorf("error: cannot consume that much fuel due to max tank capacity (%0.1f)", h.tankCap)
	}

	if energy > h.batLevel {
		return fmt.Errorf("error: not enough battery charge to consume %0.1f of energy, %0.1f more needed", energy, energy-h.batLevel)
	} else if fuel > h.fuelLevel {
		return fmt.Errorf("error: not enough fuel in tank to consume %0.1f of fuel, %0.1f more needed", fuel, fuel-h.fuelLevel)
	}

	h.fuelLevel -= fuel
	h.batLevel -= energy
	return nil
}

// ==МЕТОДЫ ESCOOTER==

func (esc *Escooter) Move() error {
	if esc.isMoving {
		return fmt.Errorf("error: electric scooter %s is already moving", esc.model)
	}
	if esc.batLevel == 0 {
		return fmt.Errorf("error: no charge in Escooter %s", esc.model)
	}
	esc.isMoving = true
	log.Printf("Электросамокат %s начал движение.\n", esc.model)
	return nil
}

func (esc *Escooter) Stop() error {
	if !esc.isMoving {
		return fmt.Errorf("error: Escooter %s already stopped", esc.model)
	}
	log.Printf("Электросамокат %s остановился.\n", esc.model)
	esc.isMoving = false
	return nil
}

func (esc *Escooter) IsMoving() bool {
	return esc.isMoving
}

func (esc *Escooter) Charge(energy float64) error {
	if esc.isMoving {
		return fmt.Errorf("error: scooter %s is moving, stop to charge", esc.model)
	}
	var err error
	esc.batLevel, err = addUntilFull(esc.batLevel, esc.batCap, energy)

	return err
}

func (esc *Escooter) GetBatteryLevel() float64 {
	return esc.batLevel
}

func (esc *Escooter) GetBatteryCapacity() float64 {
	return esc.batCap
}

func (esc *Escooter) GetLocation() float64 {
	return esc.locationX
}

func (esc *Escooter) GetDistanceTravelled() float64 {
	return esc.distanceTravelled
}

func (esc *Escooter) ChangeLocation(distance float64) {
	esc.locationX += distance
}

func (esc *Escooter) IncreaseDistanceTravelled(distance float64) {
	esc.distanceTravelled += distance
}

func (esc *Escooter) String() string {
	return fmt.Sprintf("Electric scooter %s info:\nBattery capacity: %0.1f\nBattery level: %0.1f\nCharge Rate: %0.1f\nCurrent location: %0.1f\nLast travelled distance: %0.1f\nIs this hybrid moving? - %v\n",
		esc.model, esc.batCap, esc.batLevel, esc.chargeRate, esc.locationX, esc.distanceTravelled, esc.isMoving)
}

func (esc *Escooter) CalculateTrip(distance float64) (fuel, energy float64, err error) {
	if esc.isMoving {
		return 0, 0, fmt.Errorf("error: E.scooter %s is moving", esc.model)
	}
	if distance <= 0 {
		return 0, 0, fmt.Errorf("error: distance must be positive")
	}

	return 0, distance * esc.chargeRate, nil
}

func (esc *Escooter) ConsumeEnergy(fuel, energy float64) error {
	if energy > esc.batCap {
		return fmt.Errorf("error: cannot consume that much energy due to max tank capacity (%0.1f)", esc.batCap)
	}

	if energy > esc.batLevel {
		return fmt.Errorf("error: not enough energy in E.Scooter %s", esc.model)
	}
	esc.batLevel -= energy
	return nil
}

// ==Общая логика транспорта==
// - заправка транспорта
func addUntilFull(current, capacity, amount float64) (float64, error) {
	if amount <= 0 {
		return current, fmt.Errorf("error: amount must be positive")
	}
	emptySpace := capacity - current

	if emptySpace > amount {
		current += amount
		return current, nil
	}
	return capacity, nil
}

// ==TRANSPORT SERVICE==
type TransportService struct {
	logger *log.Logger
}

func (ts *TransportService) TestDrive(vehicle TestDrivable, distance float64) error {
	// 1. Рассчитать необходимое топливо для поездки
	fuel, energy, err := vehicle.CalculateTrip(float64(distance))
	if err != nil {
		return err
	}
	// 2. Определить хватает ли топлива
	err = vehicle.ConsumeEnergy(fuel, energy)
	if err != nil {
		return err
	}
	// 3. Запустить поездку и сделать лог об этом
	err = vehicle.Move()
	if err != nil {
		return err
	}
	// сделать лог
	// 4. Посчитать координату и добавить пробег
	vehicle.ChangeLocation(distance)
	vehicle.IncreaseDistanceTravelled(distance)
	// 5. Закончить поездку сделать лог об этом
	err = vehicle.Stop()
	if err != nil {
		return err
	}
	// сделать лог что поездка окончена
	// 6. Лог что транспорт прошел тест драйв
	// сделать лог что транспорт прошел тестдрайв
	return nil
}

func (ts *TransportService) Refuel(vehicle interface{}, amount float64) error {
	// Универсальная заправка - работает с любым транспортом, поддерживающим Fueler
	// здесь походу typeSwitch
}

func (ts *TransportService) Charge(vehicle interface{}, energy float64) error {
	// Универсальная зарядка - работает с любым транспортом, поддерживающим Charger
}

func (ts *TransportService) GetVehicleInfo(vehicle interface{}) string {
	// Возвращает информацию о транспорте в зависимости от его возможностей
}

func main() {
	log.Println()
	fmt.Println()
}
