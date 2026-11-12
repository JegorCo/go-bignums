package main

import (
	"fmt"
	"math/rand"
	"sync"
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

	if len(data) == 0 {
		return 0
	}

	max := data[0]

	for _, number := range data[1:] {
		if number > max {
			max = number
		}
	}

	return max

}

// maxChunks returns the maximum number of elements in a chunks.
func maxChunks(data []int) int {
	if len(data) == 0 {
		return 0
	}

	maxValues := make([]int, CHUNKS)
	chunkSize := len(data) / CHUNKS

	var wg sync.WaitGroup

	for i:=0; i < CHUNKS; i++ {
		end = len(data)
	}

	// Начальное значение для максимума чанка
	maxValues[i] = data[0]

	wg.Add(1)

	go func(index int, part []int) {
		defer wg.Done()

		if len(part) == 0 {
			return nil
		}
	}
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
