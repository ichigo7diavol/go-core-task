// На вход подаются два неупорядоченных слайса int любой длины.
// Например:
// ```
// a := []int{65, 3, 58, 678, 64}
// b := []int{64, 2, 3, 43}
// ```
// Напишите функцию, которая проверяет, есть ли пересечения значений между двумя слайсами и возвращает:
// - bool значение есть ли хотя бы одно пересечение в значениях входных срезов
// - срез []int с пересеченными значениями (если таких значений нет, то возвращать пустой срез).
// Т. е. если взать слайсы a и b из премере, то вернет ```true, []int{64, 3}```.
// * Напишите unit тесты к созданным функциям

package main

import (
	"fmt"
)

func Intersect(a, b []int) (bool, []int) {
	if len(a) == 0 || len(b) == 0 {
		return false, []int{}
	}
	setB := make(map[int]struct{}, len(b))
	for _, v := range b {
		setB[v] = struct{}{}
	}
	var result []int
	seen := make(map[int]struct{})

	for _, v := range a {
		if _, exists := setB[v]; exists {
			if _, alreadyAdded := seen[v]; !alreadyAdded {
				result = append(result, v)
				seen[v] = struct{}{}
			}
		}
	}
	return len(result) > 0, result
}

func main() {
	a := []int{65, 3, 58, 678, 64}
	b := []int{64, 2, 3, 43}

	hasIntersection, intersection := Intersect(a, b)

	fmt.Println("Есть пересечение:", hasIntersection)
	fmt.Println("Пересечённые значения:", intersection)
}
