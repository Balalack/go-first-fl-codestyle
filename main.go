package main

import (
	"fmt"
	"math/rand"
	"strings"
)

func randint(min, max int) int {
	return rand.Intn(max-min) + min
}

func attack(charName, charClass string) string {
	switch charClass {
	case "warrior":
		return fmt.Sprintf("%s нанес урон противнику равный %d.", charName, 5+randint(3, 5))
	case "mage":
		return fmt.Sprintf("%s нанес урон противнику равный %d.", charName, 5+randint(5, 10))
	case "healer":
		return fmt.Sprintf("%s нанес урон противнику равный %d.", charName, 5+randint(-3, -1))
	default:
		return "неизвестный класс персонажа"
	}
}

// обратите внимание на "if else" и на "else"
func defence(charName, charClass string) string {
	switch charClass {
	case "warrior":
		return fmt.Sprintf("%s блокировал %d урона.", charName, 10+randint(5, 10))
	case "mage":
		return fmt.Sprintf("%s блокировал %d урона.", charName, 10+randint(-2, 2))
	case "healer":
		return fmt.Sprintf("%s блокировал %d урона.", charName, 10+randint(2, 5))
	default:
		return "неизвестный класс персонажа"
	}
}

// обратите внимание на "if else" и на "else"
func special(charName, charClass string) string {
	switch charClass {
	case "warrior":
		return fmt.Sprintf("%s применил специальное умение `Выносливость %d`", charName, 80+25)
	case "mage":
		return fmt.Sprintf("%s применил специальное умение `Атака %d`", charName, 5+40)
	case "healer":
		return fmt.Sprintf("%s применил специальное умение `Защита %d`", charName, 10+30)
	default:
		return "неизвестный класс персонажа"
	}
}

// здесь обратите внимание на имена параметров
func start_training(charName, charClass string) string {

	switch charClass {
	case "warrior":
		fmt.Printf("%s, ты Воитель - отличный боец ближнего боя.\n", charName)
	case "mage":
		fmt.Printf("%s, ты Маг - превосходный укротитель стихий.\n", charName)
	case "healer":
		fmt.Printf("%s, ты Лекарь - чародей, способный исцелять раны.\n", charName)
	default:
		return "неизвестный класс персонажа"
	}

	fmt.Println("Потренируйся управлять своими навыками.")
	fmt.Println("Введи одну из команд: attack — чтобы атаковать противника,")
	fmt.Println("defence — чтобы блокировать атаку противника,")
	fmt.Println("special — чтобы использовать свою суперсилу.")
	fmt.Println("Если не хочешь тренироваться, введи команду skip.")

	var cmd string
	for cmd != "skip" {
		fmt.Print("Введи команду: ")
		fmt.Scanf("%s\n", &cmd)

		switch cmd {
		case "attack":
			fmt.Println(attack(charName, charClass))
		case "defence":
			fmt.Println(defence(charName, charClass))
		case "special":
			fmt.Println(special(charName, charClass))
		}
	}

	return "тренировка окончена"
}

// обратите внимание на имя функции и имена переменных
func choiseCharClass() string {
	var approveChoice string
	var charClass string

	fmt.Print("Введи название персонажа, за которого хочешь играть:")
	fmt.Println("Воитель — warrior, Маг — mage, Лекарь — healer: ")
	fmt.Scanf("%s\n", &charClass)

	switch charClass {
	case "warrior":
		fmt.Println("Воитель — дерзкий воин ближнего боя. Сильный, выносливый и отважный.")
	case "mage":
		fmt.Println("Маг — находчивый воин дальнего боя. Обладает высоким интеллектом.")
	case "healer":
		fmt.Println("Лекарь — могущественный заклинатель. Черпает силы из природы, веры и духов.")
	}

	fmt.Print("Нажми (Y), чтобы подтвердить выбор,")
	fmt.Println("или любую другую кнопку, чтобы выбрать другого персонажа: ")
	fmt.Scanf("%s\n", &approveChoice)
	approveChoice = strings.ToLower(approveChoice)

	return charClass
}

// обратите внимание на имена переменных
func main() {
	fmt.Println("Приветствую тебя, искатель приключений!")
	fmt.Println("Прежде чем начать игру...")

	var charName string
	fmt.Print("...назови себя: ")
	fmt.Scanf("%s\n", &charName)

	fmt.Printf("Здравствуй, %s\n", charName)
	fmt.Println("Сейчас твоя выносливость — 80, атака — 5 и защита — 10.")
	fmt.Println("Ты можешь выбрать один из трёх путей силы:")
	fmt.Println("Воитель, Маг, Лекарь")

	charClass := choiseCharClass()

	fmt.Println(start_training(charName, charClass))
}
