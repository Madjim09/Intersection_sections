package util

import (
	"strconv"
	"strings"
)

// Validinput валидирует строку в виде массива чисел
func Validinput(str string) ([]int, error) {
	text := strings.Fields(strings.TrimSpace(str))
	slRes := make([]int, 0, len(text))
	for _, v := range text {
		intVal, err := strconv.Atoi(v)
		if err != nil {
			return nil, err
		}
		slRes = append(slRes, intVal)
	}
	return slRes, nil
}
