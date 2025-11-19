package util

import (
	"strconv"
	"strings"
	"unicode"
)

var (
	slRes []int
	text  []string
)

// Validinput валидирует строку в виде массива чисел
func Validinput(str string) ([]int, error) {
	text = strings.FieldsFunc(strings.TrimSpace(str), unicode.IsSpace)
	slRes = make([]int, 0, len(text))
	for i, v := range text {
		intVal, err := strconv.Atoi(v)
		if err != nil {
			return nil, err
		}
		slRes[i] = intVal
	}
	return slRes, nil
}
