package main

import (
	"fmt"
	"github.com/eiannone/keyboard"
)

func main() {

	if err := keyboard.Open(); err != nil {
		panic(err)
	}
	defer func() {
		_ = keyboard.Close()
	}()

	fmt.Println("Нажмите стрелку влево, стрелку вправо или Esc для выхода...")
	fmt.Println("Добро пожаловать!")
	fmt.Println("\nМеню действий")
	fmt.Println("\nЧтобы выйти, нажми кнопку 5")

	for {
		char, key, err := keyboard.GetKey()
		if err != nil {
			panic(err)
		}

		if key == keyboard.KeyArrowLeft {
			fmt.Println("Стрелка влево нажата")
		} else if key == keyboard.KeyArrowRight {
			fmt.Println("Стрелка вправо нажата")
		} else if key == keyboard.KeyEsc {
			fmt.Println("Выход")
			break
		}

		if char == 'q' || char == 'Q' {
			fmt.Println("Выход")
			break
		}
	}
}

//package main
//
//import (
//	"fmt"
//)
//
//func main() {
//	s := "gopher"
//	fmt.Println("Hello and welcome, %s!", s)
//
//	for {
//		fmt.Println("Добро пожаловать!")
//		fmt.Println("\nМеню действий")
//		fmt.Println("\nЧтобы выйти, нажми кнопку 5")
//
//	}
//
//	for i := 1; i <= 5; i++ {
//		fmt.Println("i =", 100/i)
//	}
//}
