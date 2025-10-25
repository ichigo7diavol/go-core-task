// #### Задание 7
// Напишите программу на Go, которая сливает N каналов в один.
// * Напишите unit тесты к созданным функциям

package main

import (
	"fmt"
)

func Merge[T any](chans ...<-chan T) <-chan T {
	out := make(chan T)

	go func() {
		defer close(out)

		done := make(chan struct{})
		active := len(chans)

		for _, ch := range chans {
			go func(c <-chan T) {
				for v := range c {
					out <- v
				}
				done <- struct{}{}
			}(ch)
		}
		for i := 0; i < active; i++ {
			<-done
		}
		close(done)
	}()
	return out
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)

	go func() {
		defer close(ch1)
		for i := 0; i < 3; i++ {
			ch1 <- i
		}
	}()

	go func() {
		defer close(ch2)
		for i := 10; i < 13; i++ {
			ch2 <- i
		}
	}()

	go func() {
		defer close(ch3)
		for i := 100; i < 103; i++ {
			ch3 <- i
		}
	}()

	merged := Merge(ch1, ch2, ch3)

	for v := range merged {
		fmt.Println(v)
	}
}
