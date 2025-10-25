// На вход подаются два неупорядоченных слайса строк.
// Например:
// ```
// slice1 := []string{"apple", "banana", "cherry", "date", "43", "lead", "gno1"}
// slice2 := []string{"banana", "date", "fig"}
// ```
// Напишите функцию, которая возвращает слайс строк, содержащий элементы, которые есть в первом слайсе, но отсутствуют во втором.
// * Напишите unit тесты к созданным функциям

package main

import (
	"fmt"
)

func Difference(slice1, slice2 []string) []string {
	exists := make(map[string]struct{}, len(slice2))
	for _, s := range slice2 {
		exists[s] = struct{}{}
	}

	var result = []string{}
	for _, s := range slice1 {
		if _, found := exists[s]; !found {
			result = append(result, s)
		}
	}
	return result
}

func main() {
	slice1 := []string{"apple", "banana", "cherry", "date", "43", "lead", "gno1"}
	slice2 := []string{"banana", "date", "fig"}

	diff := Difference(slice1, slice2)
	fmt.Println("Элементы, которых нет во втором слайсе:", diff)
}
