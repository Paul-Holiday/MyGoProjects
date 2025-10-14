package main

import (
	"fmt"
	"strings"
)

// Преподаватель
type Teacher struct {
	Name    string
	Subject string
	Salary  float64
	Years   int // лет работы
}

// Кафедра
type Department struct {
	Name     string
	Budget   float64
	Teachers []*Teacher
}

// Факультет
type Faculty struct {
	Name        string
	Dean        string
	Departments []*Department
}

// Университет
type University struct {
	Name      string
	City      string
	Founded   int // год основания
	Faculties []*Faculty
}

// --КОНСТРУКТОРЫ--
func NewTeacher(name, subj string, salary float64, years int) *Teacher {
	return &Teacher{name, subj, salary, years}
}
func NewDepartment(name string, budget float64, teachers []*Teacher) *Department {
	return &Department{name, budget, teachers}
}
func NewFaculty(name, dean string, departments []*Department) *Faculty {
	return &Faculty{name, dean, departments}
}
func NewUniversity(name, city string, founded int, faculties []*Faculty) *University {
	return &University{name, city, founded, faculties}
}

// --МЕТОДЫ TEACHER--
func (t Teacher) Introduce() {
	fmt.Printf("Преподаватель: %s. Предмет: %s. Месячная зарплата: %0.f. Стаж: %d лет.",
		t.Name, t.Subject, t.Salary, t.Years)
}
func (t Teacher) AnnualSalary() float64 {
	return t.Salary * 12
}

// --МЕТОДЫ DEPARTMENT--
func (d *Department) AddTeacher(t *Teacher) {
	if d == nil {
		return
	}
	d.Teachers = append(d.Teachers, t)
}

func (d *Department) TotalSalaryBudget() float64 {
	if d == nil {
		return 0
	}
	sum := 0.0
	for _, teacher := range d.Teachers {
		sum += teacher.AnnualSalary()
	}
	return sum
}

func (d *Department) FindTeacher(name string) *Teacher {
	if d == nil {
		return nil
	}

	for _, teacher := range d.Teachers {
		if strings.EqualFold(teacher.Name, name) {
			return teacher
		}
	}
	return nil
}

// --МЕТОДЫ FACULTY--
func (f *Faculty) AddDepartment(d *Department) {
	f.Departments = append(f.Departments, d)
}

func (f *Faculty) TotalTeachers() int {
	if f == nil {
		return 0
	}

	sum := 0
	for _, dep := range f.Departments {
		sum += len(dep.Teachers)
	}
	return sum
}

func (f *Faculty) DisplayStructure() {
	if f == nil {
		fmt.Println("Ошибка, параметры факультета не заданы.")
		return
	}

	fmt.Printf("\n==ФАКУЛЬТЕТ: %s (Декан: %s)==\n", f.Name, f.Dean)
	for _, dep := range f.Departments {
		fmt.Printf("\n=КАФЕДРА: %s=\n\nБюджет: %0.f\n\nПреподаватели:\n", dep.Name, dep.Budget)
		for i, teacher := range dep.Teachers {
			fmt.Printf("%d. %s - %s, Зарплата: %.fруб. (Стаж: %d лет).\n",
				i+1, teacher.Name, teacher.Subject, teacher.Salary, teacher.Years)
		}
	}
}

// --МЕТОДЫ UNIVERSITY--
func (u *University) AddFaculty(f *Faculty) {
	if u == nil {
		return
	}
	u.Faculties = append(u.Faculties, f)
}

func (u *University) TotalBudget() float64 {
	sum := 0.0
	for _, fac := range u.Faculties {
		for _, dep := range fac.Departments {
			sum += dep.Budget
		}
	}
	return sum
}

func (u *University) DisplayUniversity() {
	if u == nil {
		fmt.Println("Ошибка, параметры университета не заданы!")
		return
	}

	fmt.Printf("\n\n=== УНИВЕРСИТЕТ: %s (%s, основан в %d) ===\n\n",
		u.Name, u.City, u.Founded)
	for i, fac := range u.Faculties {
		fmt.Printf("\n%d. ", i+1)
		fac.DisplayStructure()
	}
}

func (u *University) FindTeacherEveryWhere(name string) *Teacher {
	return nil // пока не сделал
}

func main() {
	// Создаём преподавателей
	mathTeacher := NewTeacher("Иван Петров", "Математика", 50000, 10)
	physicsTeacher := NewTeacher("Анна Сидорова", "Физика", 55000, 15)

	// Создаём кафедру
	mathDept := NewDepartment("Кафедра высшей математики", 1000000, []*Teacher{})
	mathDept.AddTeacher(mathTeacher)

	physicsDept := NewDepartment("Кафедра физики", 1200000, []*Teacher{})
	physicsDept.AddTeacher(physicsTeacher)

	// Создаём факультет
	scienceFaculty := NewFaculty("Факультет естественных наук", "Проф. Смирнов", []*Department{})
	scienceFaculty.AddDepartment(mathDept)
	scienceFaculty.AddDepartment(physicsDept)

	// Создаём университет
	university := NewUniversity("МГУ", "Москва", 1755, []*Faculty{})
	university.AddFaculty(scienceFaculty)

	// Демонстрация работы методов
	university.DisplayUniversity()
	fmt.Printf("Общий бюджет: %.2f руб.\n", university.TotalBudget())
}
