package main

import (
	"fmt"
	"log"
	"os"
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
	TransportInfoGetter
	Mover
	LocationGetter
	LocationSetter
	TripCalculator
	EnergyConsumer
	EnergyChecker
}

type TransportInfoGetter interface {
	GetModel() string
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

type LocationSetter interface {
	ChangeLocation(distance float64)
	IncreaseDistanceTravelled(distance float64)
}

type LocationGetter interface {
	GetLocation() float64
	GetDistanceTravelled() float64
}

type TripCalculator interface {
	CalculateTrip(distance float64) (fuel, energy float64, err error)
}

type EnergyChecker interface {
	CheckEnergy(fuel, energy float64) bool
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

func (c *Car) GetModel() string {
	return c.model
}

func (c *Car) Move() error {
	if c.isMoving {
		return fmt.Errorf("car %s is already moving", c.model)

	}
	if c.fuelLevel == 0 {
		return fmt.Errorf("no fuel in %s, refuel to move", c.model)
	}
	c.isMoving = true
	return nil
}

func (c *Car) Stop() error {
	if !c.isMoving {
		return fmt.Errorf("car %s is already stopped", c.model)
	}
	c.isMoving = false
	return nil
}

func (c *Car) IsMoving() bool {
	return c.isMoving
}

func (c *Car) Refuel(amount float64) error {
	if c.isMoving {
		return fmt.Errorf("car %s is moving, stop to refuel", c.model)
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
		return 0, 0, fmt.Errorf("car %s is moving", c.model)
	}
	if distance <= 0 {
		return 0, 0, fmt.Errorf("distance must be positive")
	}

	return distance * c.fuelRate / 100, 0, nil
}

func (c *Car) ConsumeEnergy(fuel, energy float64) error {
	if fuel < 0 {
		return fmt.Errorf("amount of fuel must be positive")
	}
	c.fuelLevel -= fuel
	return nil
}

func (c *Car) CheckEnergy(fuel, energy float64) bool {
	return fuel <= c.fuelLevel
}

// ==МЕТОДЫ HYBRID==

func (h *Hybrid) GetModel() string {
	return h.model
}

func (h *Hybrid) Move() error {
	if h.isMoving {
		return fmt.Errorf("hybrid %s is already moving", h.model)
	}
	if h.fuelLevel == 0 && h.batLevel == 0 {
		return fmt.Errorf("no fuel or charge in hybrid %s", h.model)
	}
	h.isMoving = true
	return nil
}

func (h *Hybrid) Stop() error {
	if !h.isMoving {
		return fmt.Errorf("hybrid %s is already stopped", h.model)
	}
	h.isMoving = false
	return nil
}

func (h *Hybrid) IsMoving() bool {
	return h.isMoving
}

func (h *Hybrid) Refuel(amount float64) error {
	if h.isMoving {
		return fmt.Errorf("hybrid %s is moving, stop to refuel", h.model)
	}
	var err error
	h.fuelLevel, err = addUntilFull(h.fuelLevel, h.tankCap, amount)

	return err
}

func (h *Hybrid) Charge(energy float64) error {
	if h.isMoving {
		return fmt.Errorf("hybrid %s is moving, stop to charge", h.model)
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
		return 0, 0, fmt.Errorf("hybrid %s is moving", h.model)
	}
	if distance <= 0 {
		return 0, 0, fmt.Errorf("distance must be positive")
	}

	if distance*h.chargeRate/100 <= h.batLevel {
		return 0, distance * h.chargeRate / 100, nil
	}
	fuel = (distance - h.batLevel/(h.chargeRate/100)) * h.fuelRate / 100 // (всё расстояние - расстояние которое проехали на батарее) * расход топлива на 1км
	return fuel, h.batLevel, nil
}

func (h *Hybrid) ConsumeEnergy(fuel, energy float64) error {
	if fuel < 0 || energy < 0 {
		return fmt.Errorf("amount of fuel and energy must be positive")
	}

	h.fuelLevel -= fuel
	h.batLevel -= energy
	return nil
}

func (h *Hybrid) CheckEnergy(fuel, energy float64) bool {
	if fuel > h.fuelLevel || energy > h.batLevel {
		return false
	}
	return true
}

// ==МЕТОДЫ ESCOOTER==

func (esc *Escooter) GetModel() string {
	return esc.model
}

func (esc *Escooter) Move() error {
	if esc.isMoving {
		return fmt.Errorf("electric scooter %s is already moving", esc.model)
	}
	if esc.batLevel == 0 {
		return fmt.Errorf("no charge in Escooter %s", esc.model)
	}
	esc.isMoving = true
	return nil
}

func (esc *Escooter) Stop() error {
	if !esc.isMoving {
		return fmt.Errorf("Escooter %s already stopped", esc.model)
	}
	esc.isMoving = false
	return nil
}

func (esc *Escooter) IsMoving() bool {
	return esc.isMoving
}

func (esc *Escooter) Charge(energy float64) error {
	if esc.isMoving {
		return fmt.Errorf("scooter %s is moving, stop to charge", esc.model)
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
	return fmt.Sprintf("Electric scooter %s info:\nBattery capacity: %0.1f\nBattery level: %0.1f\nCharge Rate: %0.1f\nCurrent location: %0.1f\nLast travelled distance: %0.1f\nIs this scooter moving? - %v\n",
		esc.model, esc.batCap, esc.batLevel, esc.chargeRate, esc.locationX, esc.distanceTravelled, esc.isMoving)
}

func (esc *Escooter) CalculateTrip(distance float64) (fuel, energy float64, err error) {
	if esc.isMoving {
		return 0, 0, fmt.Errorf("e.scooter %s is moving", esc.model)
	}
	if distance <= 0 {
		return 0, 0, fmt.Errorf("distance must be positive")
	}

	return 0, distance * esc.chargeRate / 100, nil
}

func (esc *Escooter) ConsumeEnergy(fuel, energy float64) error {
	if energy < 0 {
		return fmt.Errorf("amount of energy must be positive")
	}
	esc.batLevel -= energy
	return nil
}

func (esc *Escooter) CheckEnergy(fuel, energy float64) bool {
	return energy <= esc.batLevel
}

// ==Общая логика транспорта==
// - заправка транспорта
func addUntilFull(current, capacity, amount float64) (float64, error) {
	if amount <= 0 {
		return current, fmt.Errorf("amount must be positive")
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

func NewTransportService() *TransportService {
	return &TransportService{
		logger: log.New(os.Stdout, "TRANSPORT:", log.LstdFlags),
	}
}

func (ts *TransportService) TestDrive(vehicle TestDrivable, distance float64) error {
	model := vehicle.GetModel()
	// лог что тест драйв начался
	ts.logger.Printf("TEST DRIVE STARTED - vehicle: %s, distance: %.1f km", model, distance)

	// 1. Рассчитать необходимое топливо для поездки
	fuel, energy, err := vehicle.CalculateTrip(distance)
	if err != nil {
		// лог об ошибке
		ts.logger.Printf("ERROR: Trip calculation failed - vehicle: %s, error: %v", model, err)
		return fmt.Errorf("trip calculation failed, %w", err)
	}

	// лог о том сколько топлива надо для тест драйва
	ts.logger.Printf("CALCULATION - vehicle: %s, fuel_needed: %.1f, energy_needed: %.1f", model, fuel, energy)

	// 2. Определить хватает ли топлива
	if !vehicle.CheckEnergy(fuel, energy) {
		// лог об ошибке
		err = fmt.Errorf("insufficient energy: need fuel=%0.1f, energy=%0.1f", fuel, energy)
		ts.logger.Printf("ERROR: Energy check failed - vehicle: %s, error: %v", model, err)
		return err
	}

	// 3. Запустить поездку и потребить топливо, сделать лог об этом
	// лог что vehicle начинает движение
	ts.logger.Printf("MOVING - vehicle %s is starting", model)
	err = vehicle.Move()
	if err != nil {
		// лог об ошибке
		ts.logger.Printf("ERROR: Failed to start moving - vehicle: %s, error: %v", model, err)
		return fmt.Errorf("failed to start moving, %w", err)
	}

	err = vehicle.ConsumeEnergy(fuel, energy)
	if err != nil {
		ts.logger.Printf("ERROR: Failed to consume energy - vehicle: %s, error: %v", model, err)
		if stopErr := vehicle.Stop(); stopErr != nil {
			// лог об ошибке
			ts.logger.Printf("ERROR: Failed to stop - vehicle: %s, error: %v", model, stopErr)
			return fmt.Errorf("failed to consume energy and failed to stop: %w; %w", err, stopErr)
		}
		return fmt.Errorf("energy consumption failed, %w", err)
	}
	// лог что было потреблено опр количество топлива/энергии
	ts.logger.Printf("ENERGY_CONSUMED - vehicle: %s, fuel: %.1f, energy: %.1f", model, fuel, energy)

	// 4. Посчитать координату и добавить пробег
	vehicle.ChangeLocation(distance)
	vehicle.IncreaseDistanceTravelled(distance)
	// лог об обновлении локации
	ts.logger.Printf("LOCATION_UPDATED - vehicle: %s, new_location: %.1f, total_distance: %.1f", model, vehicle.GetLocation(), vehicle.GetDistanceTravelled())

	// 5. Закончить поездку сделать лог об этом
	//лог что vehicle останавливается
	ts.logger.Printf("STOPPING - vehicle %s is stopping", model)
	err = vehicle.Stop()
	if err != nil {
		ts.logger.Printf("ERROR: Failed to stop - vehicle: %s, error: %v", model, err)
		return fmt.Errorf("failed to stop: %w", err)
	}

	// 6. Лог что транспорт прошел тест драйв
	ts.logger.Printf("TEST DRIVE COMPLETED - vehicle: %s, distance: %.1f", model, distance)
	return nil
}

func main() {
	fmt.Println("=== СИСТЕМА ТЕСТ-ДРАЙВОВ ===")
	ts := NewTransportService()

	car := NewCar("Toyota Camry", 60, 30, 8.5, 0, 0, false)
	hybrid := NewHybrid("Toyota Prius", 45, 5, 20, 3, 5, 1.5, 0, 0, false)
	scooter := NewEscooter("Xiaomi Pro", 0.5, 0.3, 0.1, 0, 0, false)

	vehicles := []TestDrivable{car, hybrid, scooter}

	for i, vehicle := range vehicles {
		fmt.Printf("\n--- Тест-драйв %d ---\n", i+1)
		fmt.Println("Состояние до тест-драйва:")
		fmt.Println(vehicle)

		err := ts.TestDrive(vehicle, 50)
		if err != nil {
			fmt.Printf("Тест-драйв завершился ошибкой: %v\n", err)
		} else {
			fmt.Println("Тест-драйв успешно завершен!")
		}

		fmt.Println("Состояние после тест-драйва:")
		fmt.Println(vehicle)
	}
}
