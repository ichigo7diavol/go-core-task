// Реализуйте структуру данных StringIntMap, которая будет использоваться для хранения пар "строка - целое число". Ваша
// структура должна поддерживать следующие операции:
// 1. Добавление элемента: Метод Add(key string, value int), который добавляет новую пару "ключ-значение" в карту.
// 2. Удаление элемента: Метод Remove(key string), который удаляет элемент по ключу из карты.
// 3. Копирование карты: Метод Copy() map[string]int, который возвращает новую карту, содержащую все элементы текущей карты.
// 4. Проверка наличия ключа: Метод Exists(key string) bool, который проверяет, существует ли ключ в карте.
// 5. Получение значения: Метод Get(key string) (int, bool), который возвращает значение по ключу и булевый флаг, указывающий на успешность операции.
// * Напишите unit тесты к созданным функциям

package main

import "fmt"

func main() {
	m := NewStringIntMap()

	m.Add("apple", 10)
	m.Add("banana", 20)
	m.Add("cherry", 30)

	fmt.Println("Исходная карта:", m.data)

	value, exists := m.Get("banana")
	fmt.Printf("banana = %d (существует: %v)\n", value, exists)

	fmt.Println("Существует 'orange':", m.Exists("orange"))

	copyMap := m.Copy()
	fmt.Println("Копия карты:", copyMap)

	m.Remove("apple")
	fmt.Println("После удаления 'apple':", m.data)
}
