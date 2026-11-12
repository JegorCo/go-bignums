package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
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
		numbers[i] = rand.Intn(1000) + 1
	}

	return numbers
}

// maximum returns the largest element, or zero for an empty slice.
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

// maxChunks finds the largest element using CHUNKS goroutines.
func maxChunks(data []int) int {
	if len(data) == 0 {
		return 0
	}

	maxValues := make([]int, CHUNKS)
	chunkSize := len(data) / CHUNKS

	var wg sync.WaitGroup

	for i := 0; i < CHUNKS; i++ {
		start := i * chunkSize
		end := start + chunkSize
		if i == CHUNKS-1 {
			end = len(data)
		}

		// Empty chunks keep an actual input value, including for negative data.
		maxValues[i] = data[0]
		wg.Add(1)

		go func(index int, part []int) {
			defer wg.Done()
			if len(part) == 0 {
				return
			}
			maxValues[index] = maximum(part)
		}(i, data[start:end])
	}

	wg.Wait()
	return maximum(maxValues)
}

func main() {
	fmt.Printf("Генерируем %d целых чисел\n", SIZE)
	numbers := generateRandomElements(SIZE)

	fmt.Println("Ищем максимальное значение последовательно")
	start := time.Now()
	maxValue := maximum(numbers)
	elapsedMicroseconds := time.Since(start).Microseconds()

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d мкс\n", maxValue, elapsedMicroseconds)

	fmt.Printf("Ищем максимальное значение в %d горутинах\n", CHUNKS)
	start = time.Now()
	maxValue = maxChunks(numbers)
	elapsedMicroseconds = time.Since(start).Microseconds()

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d мкс\n", maxValue, elapsedMicroseconds)
}
