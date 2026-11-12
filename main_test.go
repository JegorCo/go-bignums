package main

import "testing"
// Пишите тесты в этом файле

// Тест генерации чисел
func TestGenerateRandomElements(t *testing.T) {
	for _, size := range []int{-1, 0, 1, 100} {
		numbers := generateRandomElements(size)

		wantLen := size
		if size < 0 {
			wantLen = 0
		}

		if len(numbers) != wantLen {
			t.Errorf("size=%d: длина %d, ожидалась %d",
				size, len(numbers), wantLen)
		}

		for i, number := range numbers {
			if number < 0 || number >= 1000 {
				t.Errorf("numbers[%d]=%d: число вне диапазона [0, 1000)",
					i, number)
			}
		}
	}
}

func TestMaximum(t *testing.T) {
	tests := []struct {
		name string
		data []int
		want int
	}{
		{name: "nil", data: nil, want: 0},
		{name: "пустой слайс", data: []int{}, want: 0},
		{name: "один элемент", data: []int{7}, want: 7},
		{name: "один отрицательный", data: []int{-7}, want: -7},
		{name: "максимум в начале", data: []int{9, 3, 1}, want: 9},
		{name: "максимум в середине", data: []int{1, 9, 3}, want: 9},
		{name: "максимум в конце", data: []int{1, 3, 9}, want: 9},
		{name: "отрицательные", data: []int{-5, -2, -8}, want: -2},
		{name: "смешанные знаки", data: []int{-5, 4, 0}, want: 4},
		{name: "одинаковые", data: []int{3, 3, 3}, want: 3},
		{name: "нули", data: []int{0, 0}, want: 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maximum(tt.data)

			if got != tt.want {
				t.Errorf("maximum(%v) = %d, ожидалось %d",
					tt.data, got, tt.want)
			}
		})
	}
}