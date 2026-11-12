package main

import (
	"fmt"
	"testing"
)

// TestGenerateRandomElements проверяет длину слайса и диапазон случайных чисел.
func TestGenerateRandomElements(t *testing.T) {
	// Проверяем отрицательный и нулевой размер, один элемент и несколько элементов.
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

		// Проверяем диапазон, а не конкретные значения: числа могут повторяться.
		for i, number := range numbers {
			if number < 1 || number > 1000 {
				t.Errorf("numbers[%d]=%d: число вне диапазона [1, 1000]",
					i, number)
			}
		}
	}
}

// TestMaximum проверяет последовательный поиск максимума и граничные случаи.
func TestMaximum(t *testing.T) {
	// Границы int зависят от разрядности платформы, поэтому вычисляем их здесь.
	maxInt := int(^uint(0) >> 1)
	minInt := -maxInt - 1
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
		{name: "минимальный int", data: []int{minInt}, want: minInt},
		{name: "границы int", data: []int{minInt, 0, maxInt}, want: maxInt},
	}

	// Каждый случай запускается отдельно, чтобы по имени было видно, какой не прошёл.
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

// TestMaxChunks проверяет параллельный поиск на пустых, коротких и обычных слайсах.
func TestMaxChunks(t *testing.T) {
	maxInt := int(^uint(0) >> 1)
	minInt := -maxInt - 1
	tests := []struct {
		name string
		data []int
		want int
	}{
		{name: "nil", data: nil, want: 0},
		{name: "пустой слайс", data: []int{}, want: 0},
		{name: "один элемент", data: []int{7}, want: 7},
		{name: "один отрицательный", data: []int{-7}, want: -7},
		// Пустые части не должны подменять отрицательный максимум нулём.
		{name: "короткий отрицательный слайс", data: []int{-5, -2, -8}, want: -2},
		{name: "смешанные знаки", data: []int{-5, 4, 0}, want: 4},
		{name: "одинаковые отрицательные", data: []int{-3, -3, -3, -3, -3, -3, -3, -3}, want: -3},
		{name: "нули", data: []int{0, 0}, want: 0},
		{name: "минимальный int", data: []int{minInt}, want: minInt},
		{name: "границы int", data: []int{minInt, 0, maxInt}, want: maxInt},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxChunks(tt.data); got != tt.want {
				t.Errorf("maxChunks(%v) = %d, ожидалось %d", tt.data, got, tt.want)
			}
		})
	}
}

// TestMaxChunksBoundaries проверяет, что при делении на части не теряются элементы.
func TestMaxChunksBoundaries(t *testing.T) {
	// Проверяем длины меньше CHUNKS, кратные CHUNKS и с остатком от деления.
	for _, size := range []int{CHUNKS - 1, CHUNKS, CHUNKS + 1, 2 * CHUNKS, 2*CHUNKS + 3} {
		// Перемещаем максимум по всем позициям, включая границы частей и остаток.
		for maxIndex := 0; maxIndex < size; maxIndex++ {
			t.Run(fmt.Sprintf("size=%d/maxIndex=%d", size, maxIndex), func(t *testing.T) {
				// Отрицательные значения также выявляют ошибочный максимум, равный нулю.
				data := make([]int, size)
				for i := range data {
					data[i] = -10
				}
				data[maxIndex] = -1

				if got := maxChunks(data); got != -1 {
					t.Errorf("maxChunks(%v) = %d, ожидалось -1", data, got)
				}
			})
		}
	}
}
