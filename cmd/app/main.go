package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/Madjim09/Intersection_sections/internal/util"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	isActive := true

	for isActive {
		fmt.Print("Введите первый слайс: ")
		sl1, err := util.InputUser(scanner)

		if err != nil {
			fmt.Println("Ошибка ввода:", err)
		}

		fmt.Print("Введите второй слайс: ")
		sl2, err := util.InputUser(scanner)

		if err != nil {
			fmt.Println("Ошибка ввода:", err)
		}

		result := util.SortSlices(sl1, sl2)

		fmt.Print("Результат: ")
		isActive, err = util.OutputRes(result, scanner)

		if err != nil {
			err = fmt.Errorf("ошибка ввода: %w", err)
			fmt.Println(err)
		}
	}
}
