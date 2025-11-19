package util

import (
	"bufio"
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
