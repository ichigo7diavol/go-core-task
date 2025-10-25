// ### Задание 1
// Напишите программу на Go, которая:
// 1. Создает несколько переменных различных типов данных:
// ```
// int (три числа в десятичной, восьмеричной и шеснадцатиричной системах)
// float64
// string
// bool
// complex64
// ```
// 2. Определяет тип каждой переменной и выводит его на экран.
// 3. Преобразует все переменные в строковый тип и объединяет их в одну строку.
// 4. Преобразовать эту строку в срез рун.
// 5. Захэшировать этот срез рун SHA256, добавив в середину соль "go-2024" и вывести результат.
// * Напишите unit тесты к созданным функциям
// Напишите main функцию, в которой протестируете весь вышеописанный функционал. Выведите результаты на экран.
// Входные числа из пункта 1 могут быть:
// ```
// var numDecimal int = 42          	// Десятичная система
// var numOctal int = 052           	// Восьмеричная система
// var numHexadecimal int = 0x2A     	// Шестнадцатиричная система
// var pi float64 = 3.14             	// Тип float64
// var name string = "Golang"         	// Тип string
// var isActive bool = true           	// Тип bool
// var complexNum complex64 = 1 + 2i  	// Тип complex64
// ```

package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"reflect"
)

var (
	numDecimal     int       = 42
	numOctal       int       = 052
	numHexadecimal int       = 0x2A
	pi             float64   = 3.14
	name           string    = "Golang"
	isActive       bool      = true
	complexNum     complex64 = 1 + 2i
)

func GetType(v any) string {
	return reflect.TypeOf(v).String()
}

func CombineToString(values ...any) string {
	var result string
	for _, v := range values {
		result += fmt.Sprintf("%v", v)
	}
	return result
}

func ToRunes(s string) []rune {
	return []rune(s)
}

func HashWithSaltSHA256(runes []rune, salt string) string {
	mid := len(runes) / 2
	withSalt := string(runes[:mid]) + salt + string(runes[mid:])
	hash := sha256.Sum256([]byte(withSalt))
	return hex.EncodeToString(hash[:])
}

func main() {
	fmt.Println("=== Типы переменных ===")
	fmt.Println("numDecimal:", GetType(numDecimal))
	fmt.Println("numOctal:", GetType(numOctal))
	fmt.Println("numHexadecimal:", GetType(numHexadecimal))
	fmt.Println("pi:", GetType(pi))
	fmt.Println("name:", GetType(name))
	fmt.Println("isActive:", GetType(isActive))
	fmt.Println("complexNum:", GetType(complexNum))

	combined := CombineToString(numDecimal, numOctal, numHexadecimal, pi, name, isActive, complexNum)
	fmt.Println("\nОбъединённая строка:", combined)

	runes := ToRunes(combined)
	fmt.Println("Количество рун:", len(runes))

	hash := HashWithSaltSHA256(runes, "go-2024")
	fmt.Println("\nSHA256 hash с солью:", hash)
}
