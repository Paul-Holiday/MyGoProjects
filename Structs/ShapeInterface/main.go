package main

import (
	"fmt"
	"math"
)

// Интерфейс Shape
type Shape interface {
	Area() float64
	Perimeter() float64
}

// Структуры фигур
type Rectangle struct {
	Height, Width float64
}

type Triangle struct {
	A, B, C float64
}

type Circle struct {
	Radius float64
}

// КОНСТРУКТОРЫ
func NewRectangle(height, width float64) *Rectangle {
	return &Rectangle{height, width}
}

func NewTriangle(a, b, c float64) *Triangle {
	return &Triangle{a, b, c}
}

func NewCircle(r float64) *Circle {
	return &Circle{r}
}

// ==Методы для RECTANGLE==
func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

func (r Rectangle) Perimeter() float64 {
	return (r.Height + r.Width) * 2
}

// ==Методы для TRIANGLE==
func (t Triangle) IsValid() bool {
	return (t.A > 0) && (t.B > 0) && (t.C > 0) &&
		(t.A+t.B > t.C) && (t.A+t.C > t.B) && (t.B+t.C > t.A)
}

func (t Triangle) Area() float64 {
	if !t.IsValid() {
		return 0
	}
	p := (t.A + t.B + t.C) / 2
	return math.Sqrt(p * (p - t.A) * (p - t.B) * (p - t.C))
}

func (t Triangle) Perimeter() float64 {
	return t.A + t.B + t.C
}

// ==Методы для Circle==
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

// ==Функции для интерфейса Shape==
func PrintShapeInfo(s Shape) {
	fmt.Printf("%s: Площадь: %0.2f. Периметр: %0.2f.\n", GetShapeType(s), s.Area(), s.Perimeter())
}

func GetShapeType(s Shape) string {
	switch s.(type) {
	case Rectangle, *Rectangle:
		return "Прямоугольник"
	case Triangle, *Triangle:
		return "Треугольник"
	case Circle, *Circle:
		return "Круг"
	default:
		return "Неизвестная фигура"
	}
}

func main() {
	rec := NewRectangle(5, 7)
	tri := NewTriangle(3, 4, 5)
	circ := NewCircle(3)

	// Можно создать целый слайс из структур имплементирующих интерфейс Shape
	shapes := []Shape{rec, tri, circ}
	for _, shape := range shapes {
		PrintShapeInfo(shape)
	}

	// Можно вызвать функцию PrintShapeInfo() с любой структурой имплеметирующей Shape
	PrintShapeInfo(rec)
	PrintShapeInfo(tri)
	PrintShapeInfo(circ)
}
