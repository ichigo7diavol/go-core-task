// Сделать кастомную waitGroup на семафоре, не используя sync.WaitGroup.
// * Напишите unit тесты к созданным функциям

package main

import (
	"fmt"
	"time"
)

func main() {
	wg := NewSemaphoreWaitGroup()

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			time.Sleep(time.Duration(id) * 200 * time.Millisecond)
			fmt.Println("Горутина", id, "завершилась")
		}(i)
	}

	fmt.Println("Ожидание завершения всех горутин...")
	wg.Wait()
	fmt.Println("Все горутины завершились!")
}
