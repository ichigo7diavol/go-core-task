// #### Задание 6
// Напишите генератор случайных чисел используя небуфферизированный канал.
// * Напишите unit тесты к созданным функциям

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func RandomGenerator() <-chan int {
	ch := make(chan int)

	go func() {
		for {
			ch <- rand.Intn(1000)
		}
	}()

	return ch
}

func main() {
	randomChan := RandomGenerator()

	fmt.Println("Генератор случайных чисел (небуферизованный канал):")
	for i := 0; i < 5; i++ {
		fmt.Println(<-randomChan)
		time.Sleep(200 * time.Millisecond)
	}
}
