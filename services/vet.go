package services

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type VetClinic interface {
	CheckHealth() bool
}

type vetClinicImpl struct{}

func NewVetClinic() VetClinic {
	return &vetClinicImpl{}
}

func (v *vetClinicImpl) CheckHealth() bool {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Введите состояние здоровья животного (здоров/нездоров): ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Ошибка ввода, попробуйте снова.")
			continue
		}
		input = strings.TrimSpace(input)
		if input == "здоров" {
			return true
		} else if input == "нездоров" {
			return false
		} else {
			fmt.Println("Неверный ввод. Введите 'здоров' или 'нездоров'.")
		}
	}
}
