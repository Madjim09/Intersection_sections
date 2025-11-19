package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/Madjim09/Intersection_sections/internal/util"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	flag := true

	for flag {
		fmt.Print("Введите первый слайс: ")
		sl1, err := util.InputUser(scanner)

		if err != nil {
			err = fmt.Errorf("ошибка ввода: %w", err)
			fmt.Println(err)
		}

		fmt.Print("Введите второй слайс: ")
		sl2, err := util.InputUser(scanner)

		if err != nil {
			err = fmt.Errorf("ошибка ввода: %w", err)
			fmt.Println(err)
		}

		result := util.SortSlices(sl1, sl2)

		fmt.Print("Результат: ")
		flag, err = util.OutoutRes(result, scanner)

		if err != nil {
			err = fmt.Errorf("ошибка ввода: %w", err)
			fmt.Println(err)
		}
	}
}
