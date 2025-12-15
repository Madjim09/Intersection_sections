package util

import (
	"bufio"
	"fmt"
	"strings"
)

// InputUser ввод чисел пользователем
func InputUser(scanner *bufio.Scanner) ([]int, error) {
	if !scanner.Scan() {
		return nil, scanner.Err()
	}

	slise, err := Validinput(scanner.Text())

	if err != nil {
		return nil, err
	}

	return slise, nil
}

// SortSlices сортировка слайсов
func SortSlices(sl1, sl2 []int) []int {
	slMap := make(map[int]int, len(sl2))

	for _, v := range sl2 {
		slMap[v]++
	}

	slRes := make([]int, 0, len(sl1))
	for _, v := range sl1 {
		if slMap[v] > 0 {
			slRes = append(slRes, v)
			slMap[v]--
		}
	}

	return slRes
}

// OutoutRes вывод результата и запрос на продолжение
func OutputRes(sl []int, scanner *bufio.Scanner) (bool, error) {
	if len(sl) == 0 {
		fmt.Print("Empty intersection")
	} else {
		for _, v := range sl {
			fmt.Printf("%d ", v)
		}
	}
	fmt.Println()

	fmt.Print("Хотите продолжить? [y/n]: ")
	for {
		if !scanner.Scan() {
			return false, scanner.Err()
		}
		answer := strings.TrimSpace(strings.ToLower(scanner.Text()))
		switch answer {
		case "y", "yes", "да", "д":
			fmt.Println()
			return true, nil
		case "n", "no", "нет", "н":
			return false, nil
		default:
			fmt.Println("Неверный ввод.")
			fmt.Print("Введите y (да) или n (нет): ")
		}
	}
}
