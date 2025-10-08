package main

import (
	"fmt"
	"math"
)

// структура треугольника
type Triangle struct {
	A float64
	B float64
	C float64
}

// структура прямоугольника
type Rectangle struct{ Height, Width float64 }

// конструктор triangle
func NewTriangle(a, b, c float64) *Triangle {
	return &Triangle{
		A: a,
		B: b,
		C: c,
	}
}

// конструктор rectangle
func NewRectangle(a, b float64) *Rectangle {
	return &Rectangle{a, b}
}

// метод для валидации треугольника
func (tri *Triangle) IsValid() bool {
	// все стороны должны быть положительными, сумма катетов больше гипотенузы
	// и сумма гипотенузы с катетом из катетов больше оставшегося катета
	if tri.A > 0 && tri.B > 0 && tri.C > 0 &&
		tri.A+tri.B > tri.C &&
		tri.C+tri.A > tri.B &&
		tri.C+tri.B > tri.A {
		return true
	} else {
		return false
	}
}

// метод для валидации прямоугольника
func (rec *Rectangle) IsValid() bool {
	if rec.Height > 0 && rec.Width > 0 {
		return true
	} else {
		return false
	}
}

// вычисляет и возвращает периметр треугольника rec и ошибку
func (tri *Triangle) Perimeter() (float64, error) {
	if tri.IsValid() {
		return tri.A + tri.B + tri.C, nil
	} else {
		return 0, fmt.Errorf("\nТреугольник недопустимых параметров, периметр рассчитать невозможно")
	}
}

// вычисляет и возвращает площадь треугольника и ошибку
func (tri *Triangle) Area() (float64, error) {
	// полупериметр и катеты для формулы площади Герона
	if tri.IsValid() {
		p, _ := tri.Perimeter()
		p /= 2
		a := tri.A
		b := tri.B
		c := tri.C
		return math.Sqrt(p * (p - a) * (p - b) * (p - c)), nil
	} else {
		return 0, fmt.Errorf("\nТреугольник недопустимых параметров, площадь рассчитать невозможно")
	}
}

// вычисляет периметр прямоугольника rec и ошибку
func (rec *Rectangle) Perimeter() (float64, error) {
	if rec.IsValid() {
		return 2 * (rec.Height + rec.Width), nil
	} else {
		return 0, fmt.Errorf("\nПрямоугольник недопустимых параметров, рассчитать периметр невозможно")
	}
}

// вычисляет площадь прямоугольника rec и ошибку
func (rec *Rectangle) Area() (float64, error) {
	if rec.IsValid() {
		return rec.Height * rec.Width, nil
	} else {
		return 0, fmt.Errorf("Прямоугольник недопустимых параметров, рассчитать площадь невозможно")
	}
}

func main() {
	tri := NewTriangle(3, 4, 5)

	triPer, err := tri.Perimeter()
	triArea, err := tri.Area()

	fmt.Println(triPer, triArea, err)

	rec := NewRectangle(3, 2)
	recPer, err2 := rec.Perimeter()
	recArea, err2 := rec.Area()

	fmt.Println(recPer, recArea, err2)
}
