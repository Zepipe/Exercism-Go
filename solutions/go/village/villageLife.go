package main

import (
	"fmt"
	"math/rand/v2"
	"strings"
)

// =============================Resident==============================
type Resident struct {
	Name    string
	Age     int
	Married bool
	Alive   bool
	Events  []string
}

func NewResident(name string, age int) *Resident {
	return &Resident{
		Name:    name,
		Age:     age,
		Married: false,
		Alive:   true,
		Events:  []string{},
	}
}

// IncAge increments resident's age by 1 year
func (r *Resident) IncAge() {
	if r.Alive {
		r.Age++
	}
}

// ChangeMarriageStatus toggles the current marriage status
func (r *Resident) ChangeMarriageStatus() {
	r.Married = !r.Married

	if r.Married {
		r.Events = append(r.Events, "Наконец-то в браке!")
	} else {
		r.Events = append(r.Events, "Я свобоооден(-на) словно птица в небесах. Брак закончен, я забыл, что значит страх")
	}
}

// ChangeAliveStatus toggles the current alive status
func (r *Resident) ChangeAliveStatus() {
	r.Alive = !r.Alive

	if !r.Alive {
		r.Events = append(r.Events, "Этот парень/дама был(-а) из тех, кто просто любит жизнь....")
	}
}

// FlushInfo returns resident information and clears the events list
func (r *Resident) FlushInfo() string {

	lifeStatus, marriageStatus := "", ""

	if r.Alive {
		lifeStatus = "жив(-а)"
	} else {
		lifeStatus = "умер(-ла)"
	}

	if r.Married {
		marriageStatus = "в браке/замужем"
	} else {
		marriageStatus = "холост/одна"
	}

	residentInfo := fmt.Sprintf("Имя: %s.\nВозраст: %d.\nСтатус брака: %s.\nЖив ли: %s.\nСписок событий: %s.\n",
		r.Name, r.Age, marriageStatus, lifeStatus, strings.Join(r.Events, ", "))

	r.Events = nil

	return residentInfo
}

func (r *Resident) Update() {
	if !r.Alive {
		return
	}

	if rand.IntN(100) < 20 {
		r.ChangeMarriageStatus()
	}

	if rand.IntN(100) < 15 {
		r.Events = append(r.Events, "Новая работа")
	}

	if rand.IntN(100) < 5 {
		r.ChangeAliveStatus()
	}

	r.IncAge()
}

// ============================Animal=================================

type Animal struct {
	Name   string
	Age    int
	Type   string
	Alive  bool
	Events []string
}

func NewAnimal(name string, age int, animalType string) *Animal {
	return &Animal{
		Name:   name,
		Age:    age,
		Type:   animalType,
		Alive:  true,
		Events: []string{},
	}
}

// IncAge increments animal's age by 1 year
func (a *Animal) IncAge() {
	if a.Alive {
		a.Age++
	}
}

// ChangeAliveStatus toggles the current alive status
func (a *Animal) ChangeAliveStatus() {
	a.Alive = !a.Alive

	if !a.Alive {
		a.Events = append(a.Events, "Этот мякиш был из тех, кто просто любит жизнь....")
	}
}

// FlushInfo returns animal information and clears the events list
func (a *Animal) FlushInfo() string {
	lifeStatus := ""

	if a.Alive {
		lifeStatus = "жив(-а)"
	} else {
		lifeStatus = "умер(-ла)"
	}

	animalInfo := fmt.Sprintf("Имя: %s.\nВозраст: %d.\nТип животного: %s.\nЖив ли: %s.\nСписок событий: %s.\n",
		a.Name, a.Age, a.Type, lifeStatus, strings.Join(a.Events, ", "))

	a.Events = nil

	return animalInfo
}

func (a *Animal) Update() {
	if !a.Alive {
		return
	}

	if rand.IntN(100) < 15 {
		a.Events = append(a.Events, "Купили новый корм")
	}

	if rand.IntN(100) < 20 {
		a.Events = append(a.Events, "Сводили к ветеринару")
	}

	if rand.IntN(100) < 5 {
		a.ChangeAliveStatus()
	}

	a.IncAge()
}

// ===========================Village=============================

type Village struct {
	Elements []VillageElement
}

type VillageElement interface {
	Update()
	FlushInfo() string
}

// AddElement add new element in the village (animals or residents)
func (v *Village) AddElement(elem VillageElement) {
	v.Elements = append(v.Elements, elem)
}

// UpdateAll updates information about all elements in the Village
func (v *Village) UpdateAll() {
	for _, e := range v.Elements {
		e.Update()
	}
}

// ShowAllInfo shows the whole information about life in the Village
func (v Village) ShowAllInfo() string {
	result := ""
	for _, e := range v.Elements {
		result += e.FlushInfo() + "\n"
	}
	return result
}

// ============================Тестовый кейс=============================

func main() {
	fmt.Println("=== СИМУЛЯЦИЯ ЖИЗНИ В ДЕРЕВНЕ ===")

	// Создаем деревню
	village := &Village{}

	// Создаем жителей
	alice := NewResident("Алиса", 25)
	bob := NewResident("Боб", 30)

	// Создаем животных
	rex := NewAnimal("Рекс", 3, "Собака")
	murka := NewAnimal("Мурка", 2, "Кошка")

	// Добавляем всех в деревню
	village.AddElement(alice)
	village.AddElement(bob)
	village.AddElement(rex)
	village.AddElement(murka)

	fmt.Println("--- НАЧАЛЬНОЕ СОСТОЯНИЕ ---")
	fmt.Println(village.ShowAllInfo())

	// Симулируем 5 лет жизни в деревне
	fmt.Println("\n--- ПРОШЛО 5 ЛЕТ ---")
	for year := 1; year <= 5; year++ {
		fmt.Printf("\n🔄 Год %d:\n", year)
		village.UpdateAll()
		fmt.Printf("✅ Все обновления применены\n")
	}

	// Показываем финальное состояние
	fmt.Println("\n--- ФИНАЛЬНОЕ СОСТОЯНИЕ ---")
	fmt.Println(village.ShowAllInfo())

}

// Тестовый вывод
/* 
=== СИМУЛЯЦИЯ ЖИЗНИ В ДЕРЕВНЕ ===
--- НАЧАЛЬНОЕ СОСТОЯНИЕ ---
Имя: Алиса.
Возраст: 25.
Статус брака: холост/одна.
Жив ли: жив(-а).
Список событий: .

Имя: Боб.
Возраст: 30.
Статус брака: холост/одна.
Жив ли: жив(-а).
Список событий: .

Имя: Рекс.
Возраст: 3.
Тип животного: Собака.
Жив ли: жив(-а).
Список событий: .

Имя: Мурка.
Возраст: 2.
Тип животного: Кошка.
Жив ли: жив(-а).
Список событий: .

--- ПРОШЛО 5 ЛЕТ ---

🔄 Год 1:
✅ Все обновления применены

🔄 Год 2:
✅ Все обновления применены

🔄 Год 3:
✅ Все обновления применены

🔄 Год 4:
✅ Все обновления применены

🔄 Год 5:
✅ Все обновления применены

--- ФИНАЛЬНОЕ СОСТОЯНИЕ ---
Имя: Алиса.
Возраст: 30.
Статус брака: в браке/замужем.
Жив ли: жив(-а).
Список событий: Наконец-то в браке!, Новая работа.

Имя: Боб.
Возраст: 35.
Статус брака: холост/одна.
Жив ли: жив(-а).
Список событий: Новая работа.

Имя: Рекс.
Возраст: 8.
Тип животного: Собака.
Жив ли: жив(-а).
Список событий: Купили новый корм.

Имя: Мурка.
Возраст: 5.
Тип животного: Кошка.
Жив ли: умер(-ла).
Список событий: Купили новый корм, Сводили к ветеринару, Этот мякиш был из тех, кто просто любит жизнь.....
*/
