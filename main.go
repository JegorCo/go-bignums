package main

import (
	"fmt"
	"math/rand"
)

const (
	SIZE   = 100_000_000
	CHUNKS = 8
)

// generateRandomElements generates random elements.
func generateRandomElements(size int) []int {
	
	if size <= 0 {
		return []int{}
	}

	numbers := make([]int, size)
	for i := range numbers {
		numbers[i] = rand.Intn(1000)
	}

	return numbers
}

// maximum returns the maximum number of elements.
func maximum(data []int) int {
	// ваш код здесь
}

// maxChunks returns the maximum number of elements in a chunks.
func maxChunks(data []int) int {
	// ваш код здесь
}

func main() {
	fmt.Printf("Генерируем %d целых чисел", SIZE)
	// ваш код здесь

	fmt.Println("Ищем максимальное значение в один поток")
	// ваш код здесь

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)

	fmt.Printf("Ищем максимальное значение в %d потоков", CHUNKS)
	// ваш код здесь

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)
}
