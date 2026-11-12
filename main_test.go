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
