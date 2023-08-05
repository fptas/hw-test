package hw05parallelexecution

import (
	"errors"
	"sync"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
func Run(tasks []Task, n, m int) error {
	wg := sync.WaitGroup{}
	mu := sync.Mutex{} /* один мьютекс на два блока кода,
	т.к. в первом блоке читается также переменная, которая изменяется во втором */

	// управляющие переменные:
	var errCount int            // текущее кол-во ошибок, сделано для наглядности, вместо уменьшения m: m--
	var curTask int             // индекс задачи из слайса задач, которая будет запущена при следующем запуске
	if l := len(tasks); l < n { // если массив задач меньше, то не будем запускать лишние горутины
		n = l
	}

	wg.Add(n)
	for i := 0; i < n; i++ {
		go func() {
			for {
				mu.Lock() /* залочим код перед чтением управляющих переменных,
				т.к. думаю читать нужно тоже внутри лока, иначе можем получить "грязное чтение" */
				if errCount >= m || curTask >= len(tasks) { // если число ошибок достигло лимита или все задачи обработаны
					mu.Unlock() // разблокировать и прекратить цикл обработки
					break
				}
				myCurTask := curTask // сохранить индекс своей задачи
				curTask++            // передвинуть индекс следующей запускаемой задачи
				mu.Unlock()

				err := tasks[myCurTask]() // запустить выполнние своей задачи
				if err != nil {           // если была ошибка, безопасно увеличть счетчик ошибок
					mu.Lock()
					errCount++
					mu.Unlock()
				}
			}
			wg.Done() // сообщить об остановке данной горутины
		}()
	}
	wg.Wait() // ждать окончание всех горутин

	if errCount >= m {
		return ErrErrorsLimitExceeded
	}
	return nil
}
