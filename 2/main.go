// #### Задание 2
// 1. Создайте слайс целых чисел originalSlice, содержащий 10 произвольных значений, которые генерируются случайным
// образом (при каждом запуске должны получаться новые значения)
// 2. Напишите функцию sliceExample, которая принимает слайс и возвращает новый слайс, содержащий только четные числа из исходного слайса.
// 3. Напишите функцию addElements, которая принимает слайс и число. Функция должна добавлять это число в конец слайса и возвращать новый слайс.
// 4. Напишите функцию copySlice, которая принимает слайс и возвращает его копию. Убедитесь, что изменения в оригинальном слайсе не влияют на его копию.
// 5. Напишите функцию removeElement, которая принимает слайс и индекс элемента, который нужно удалить. Функция должна возвращать новый слайс без элемента по указанному индексу.
// 6. Напишите main функцию, в которой протестируете все вышеописанные функции. Выведите результаты на экран.
// * Напишите unit тесты к созданным функциям
// Примечание.
// В качестве originalSlice можно использовать ```originalSlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}```

package main

import (
	"fmt"
	"math/rand"
)

func sliceExample(slice []int) []int {
	var result []int
	for _, v := range slice {
		if v%2 == 0 {
			result = append(result, v)
		}
	}
	return result
}

func addElements(slice []int, num int) []int {
	return append(slice, num)
}

func copySlice(slice []int) []int {
	copySlice := make([]int, len(slice))
	copy(copySlice, slice)
	return copySlice
}

func removeElement(slice []int, index int) []int {
	if index < 0 || index >= len(slice) {
		return slice
	}
	return append(slice[:index], slice[index+1:]...)
}

func main() {
	originalSlice := make([]int, 10)
	for i := range originalSlice {
		originalSlice[i] = rand.Intn(100)
	}

	fmt.Println("Оригинальный слайс:", originalSlice)

	evenSlice := sliceExample(originalSlice)
	fmt.Println("Четные числа:", evenSlice)

	addedSlice := addElements(originalSlice, 777)
	fmt.Println("После добавления числа 777:", addedSlice)

	copiedSlice := copySlice(originalSlice)
	fmt.Println("Копия слайса:", copiedSlice)

	originalSlice[0] = 999
	fmt.Println("Измененный оригинал:", originalSlice)
	fmt.Println("Копия после изменения оригинала:", copiedSlice)

	indexToRemove := 3
	removedSlice := removeElement(originalSlice, indexToRemove)
	fmt.Printf("После удаления элемента с индексом %d: %v\n", indexToRemove, removedSlice)
}
