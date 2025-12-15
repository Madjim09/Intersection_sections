package util

import (
	"reflect"
	"testing"
)

func TestSortSlices(t *testing.T) {
	tests := []struct {
		name string
		sl1  []int
		sl2  []int
		want []int
	}{
		{
			name: "Стандартный ввод",
			sl1:  []int{1, 2, 3, 4, 5, 6},
			sl2:  []int{4, 5, 6, 7, 8, 9},
			want: []int{4, 5, 6},
		},
		{
			name: "Ввод пустых значений",
			sl1:  []int{},
			sl2:  []int{},
			want: []int{},
		},
		{
			name: "Второй слайс пуст",
			sl1:  []int{1, 2, 3},
			sl2:  []int{},
			want: []int{},
		},
		{
			name: "Первый слайс пуст",
			sl1:  []int{},
			sl2:  []int{1, 2, 3},
			want: []int{},
		},
		{
			name: "Возврат значенй по порядку первого слайса",
			sl1:  []int{5, 3, 7, 9, 3},
			sl2:  []int{1, 7, 2, 3, 0, 5, 3},
			want: []int{5, 3, 7, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SortSlices(tt.sl1, tt.sl2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SortSlices() = %v, want %v", got, tt.want)
			}
		})
	}
}
