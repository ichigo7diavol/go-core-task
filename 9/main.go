// #### Задание 9
// Сделать конвейер чисел
// Даны два канала.
// В первый пишутся числа типа uint8. Нужно, чтобы числа читались из первого канала по мере поступления,
// затем эти числа должны преобразовываться в float64 и возводиться в куб и результат записывался во второй канал.
// Напишите main функцию, в которой протестируете весь вышеописанный функционал. Выведите результаты на экран.
// * Напишите unit тесты к созданным функциям

package main

import (
	"fmt"
)

func CubePipeline(in <-chan uint8, out chan<- float64) {
	for n := range in {
		out <- float64(n * n * n)
	}
	close(out)
}

func main() {
	in := make(chan uint8)
	out := make(chan float64)

	go CubePipeline(in, out)

	go func() {
		for _, v := range []uint8{1, 2, 3, 4, 5, 10} {
			in <- v
		}
		close(in)
	}()
	for result := range out {
		fmt.Println(result)
	}
}
