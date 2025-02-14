package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"KPO-miniHW-1/container"
	"KPO-miniHW-1/models"
	"KPO-miniHW-1/services"
)

func main() {
	c := container.BuildContainer()
	zoo := c.Zoo
	zoo.AddThing(models.NewTable())
	zoo.AddThing(models.NewComputer())
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("\n--- Московский зоопарк ERP ---")
		fmt.Println("1. Добавить новое животное")
		fmt.Println("2. Вывести отчет по количеству животных и общему потреблению еды")
		fmt.Println("3. Вывести список животных, пригодных для контактного зоопарка")
		fmt.Println("4. Вывести инвентаризацию животных и вещей")
		fmt.Println("5. Выход")
		fmt.Print("Выберите опцию: ")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		switch input {
		case "1":
			addAnimal(zoo, reader)
		case "2":
			totalAnimals := len(zoo.GetAnimals())
			totalFood := zoo.GetTotalFoodConsumption()
			fmt.Printf("Всего животных: %d, Общее потребление еды: %d кг в день\n", totalAnimals, totalFood)
		case "3":
			eligible := zoo.GetContactZooEligibleAnimals()
			if len(eligible) == 0 {
				fmt.Println("Нет животных, пригодных для контактного зоопарка.")
			} else {
				fmt.Println("Животные, пригодные для контактного зоопарка:")
				for _, a := range eligible {
					fmt.Printf("Наименование: %s, Инв. номер: %d\n", a.GetName(), a.GetNumber())
				}
			}
		case "4":
			zoo.PrintInventory()
		case "5":
			fmt.Println("Выход из программы.")
			return
		default:
			fmt.Println("Неверный выбор, попробуйте снова.")
		}
	}
}

func addAnimal(zoo *services.Zoo, reader *bufio.Reader) {
	fmt.Println("\nВыберите тип животного для добавления:")
	fmt.Println("1. Monkey")
	fmt.Println("2. Rabbit")
	fmt.Println("3. Tiger")
	fmt.Println("4. Wolf")
	fmt.Print("Ваш выбор: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	var animal models.IAnimal

	switch input {
	case "1":
		fmt.Print("Введите потребление еды (кг в день): ")
		dailyFood := readInt(reader)
		m := models.NewMonkey(dailyFood)
		animal = m
	case "2":
		fmt.Print("Введите потребление еды (кг в день): ")
		dailyFood := readInt(reader)
		fmt.Print("Введите уровень доброты (0-10): ")
		kindness := readInt(reader)
		r := models.NewRabbit(dailyFood, kindness)
		animal = r
	case "3":
		fmt.Print("Введите потребление еды (кг в день): ")
		dailyFood := readInt(reader)
		t := models.NewTiger(dailyFood)
		animal = t
	case "4":
		fmt.Print("Введите потребление еды (кг в день): ")
		dailyFood := readInt(reader)
		w := models.NewWolf(dailyFood)
		animal = w
	default:
		fmt.Println("Неверный выбор.")
		return
	}

	zoo.AddAnimal(animal)
}

func readInt(reader *bufio.Reader) int {
	for {
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		val, err := strconv.Atoi(input)
		if err != nil {
			fmt.Print("Неверный ввод, введите число: ")
			continue
		}
		return val
	}
}
