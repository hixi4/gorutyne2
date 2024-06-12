package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Генеруємо випадкові числа та надсилаємо їх у канал
func generateNumbers(numCh chan<- int, minMaxCh <-chan [2]int) {
	for {
		select {
		case minMax := <-minMaxCh:
			fmt.Printf("Найменше число: %d, Найбільше число: %d\n", minMax[0], minMax[1])
		default:
			num := rand.Intn(10) // Генеруємо випадкове число від 0 до 9
			numCh <- num
			time.Sleep(time.Second) // Затримка для наочності
		}
	}
}

// Отримуємо випадкові числа та знаходимо найбільше й найменше числа
func findMinMax(numCh <-chan int, minMaxCh chan<- [2]int) {
	var min, max int
	min, max = int(^uint(0)>>1), -int(^uint(0)>>1)-1 // Ініціалізуємо min та max

	for num := range numCh {
		if num < min {
			min = num
		}
		if num > max {
			max = num
		}
		minMaxCh <- [2]int{min, max}
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	numCh := make(chan int)
	minMaxCh := make(chan [2]int)

	go generateNumbers(numCh, minMaxCh)
	go findMinMax(numCh, minMaxCh)

	// Запобігаємо завершенню програми
	select {}
}
