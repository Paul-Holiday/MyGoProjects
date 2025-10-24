package main

import (
	"fmt"
	"log"
	"time"
)

// ==СТРУКТУРЫ ВИДОВ ТРАНСПОРТА==
type Car struct {
	model     string
	tankCap   float64 // литры
	fuelLevel float64 // литры
	fuelRate  float64 // литры/100км
	locationX float64 // км
	moving    bool
	avgSpeed  int // км/ч
	startTime time.Time
}

type Hybrid struct {
	model      string
	tankCap    float64 // литры
	batCap     float64 // А*ч
	fuelLevel  float64 // литры
	batLevel   float64 // А*ч
	fuelRate   float64 // литры/100км
	chargeRate float64 // А*ч/100км
	locationX  float64 // км
	moving     bool
	avgSpeed   int // км/ч
	startTime  time.Time
}

type Escooter struct {
	model      string
	batСap     float64 // А*ч
	batLevel   float64 // А*ч
	chargeRate float64 // А*ч/100км
	locationX  float64 // км
	moving     bool
	avgSpeed   int // км/ч
	startTime  time.Time
}

// ==МЕТОДЫ CAR==
func (c *Car) Move() {
	if c == nil {
		log.Println("Машина не объявлена, движение невозможно.")
		return
	}
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
	if c == nil {
		log.Println("Машина не объявлена, невозможно остановить.")
		return
	}
	if !c.moving {
		log.Println("Машина уже стоит.")
		return
	}
	now := time.Now()
	// расходованное топливо = время * ср.скорость * расход на 1 км пути
	usedFuel := now.Sub(c.startTime).Hours() * float64(c.avgSpeed) * (c.fuelRate / 100)

	// если расход топлива за время движения превысил количество топлива в баке,
	// то выяснить когда топливо кончилось и где соответственно остановилась машина.
	// логика:
	// если (расходованное топливо) > количество топлива, то {
	// путь = количество топлива / (расход на 100км / 100)
	// координата += путь
	// время остановки машины = время старта + путь / среднюю скорость машины
	// лог: "Машина %s остановилась раньше запроса, так закончилось топливо. Время остановки: %s, координата остановки: %0.f\n"
	// флаг = машина остановлена
	// количество топлива в баке = 0
	// }
	if usedFuel > c.fuelLevel {
		S := c.fuelLevel / (c.fuelRate / 100)
		c.locationX += S
		hoursPassed := S / float64(c.avgSpeed)
		timeStopped := c.startTime.Add(time.Hour * time.Duration(hoursPassed))
		log.Printf("Машина %s остановилась раньше запроса, так как закончилось топливо. Время остановки: %s, координата остановки: %0.f\n",
			c.model, timeStopped, c.locationX)
		c.fuelLevel = 0
	} else {
		// если топлива хватило, то смотрим где остановилась машина
		// координата += время поездки * среднюю скорость
		// лог: "Машина %s остановилась раньше запроса, так как закончилось топливо. Время остановки: %s, координата остановки: %0.f\n"
		// флаг = машина остановлена
		// количество топлива в баке -= расход
		S := now.Sub(c.startTime).Hours() * float64(c.avgSpeed)
		c.locationX += S
		log.Printf("Машина %s была остановлена. Время остановки: %s, координата остановки: %0.f\n", c.model, now, c.locationX)
		c.fuelLevel -= c.fuelRate / 100 * S
	}
	c.moving = false
}

// Метод должен заправлять машину на положительное количество топлива,
// если заправляется больше, чем вмещает в себя бак,
// то выводим лог, что бак заправлен полностью, а лишнее топливо не было заправлено
func (c *Car) Fuel(amount float64) {
	if c == nil {
		log.Println("Машина не объявлена, заправка невозможна.")
		return
	}
	if c.moving {
		log.Printf("Машина %s двигается, остановитесь, чтобы заправиться!\n", c.model)
		return
	}
	if amount <= 0 {
		log.Println("Количество заправляемого топлива должно быть положительным!")
		return
	}

	emptySpace := c.tankCap - c.fuelLevel

	// если места в баке больше, чем планируется заправить, то просто добавляем топлива
	if emptySpace > amount {
		c.fuelLevel += amount
		log.Printf("Бак машины %s заправлен на %0.1f/%0.1f литров.\n",
			c.model, c.fuelLevel, c.tankCap)
		return
	}
	c.fuelLevel = c.tankCap
	fuelLeft := emptySpace - amount
	log.Printf("Бак машины %s заправлен полностью - %0.1f литров. Не поместилось в бак: %0.1f л\n",
		c.model, c.fuelLevel, fuelLeft)
}

func (c *Car) GetFuelLevel() {

}

func main() {
	log.Println()
	fmt.Println()
}
