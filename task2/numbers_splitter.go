package task2

/*
Напишите функцию разделения массива чисел на массивы простых и составных чисел. Для записи в массивы используйте два
разных канала и горутины.

Важно, чтобы были использованы владельцы каналов.
*/

func splitNumbers(numbers []int) ([]int, []int) {
	primeChan, compositeChan := numberSplitter(numbers)

	var primes []int
	var composites []int
	done := make(chan struct{}, 2)

	go func() {
		for n := range primeChan {
			primes = append(primes, n)
		}
		done <- struct{}{}
	}()

	go func() {
		for n := range compositeChan {
			composites = append(composites, n)
		}
		done <- struct{}{}
	}()

	<-done
	<-done

	return primes, composites
}

// функция-владелец, которая создает и возвращает каналы
func numberSplitter(numbers []int) (<-chan int, <-chan int) {
	primeChan := make(chan int)
	compositeChan := make(chan int)

	go func() {
		defer close(primeChan)
		defer close(compositeChan)

		for _, n := range numbers {
			if isPrime(n) {
				primeChan <- n
			} else {
				compositeChan <- n
			}
		}
	}()

	return primeChan, compositeChan
}

// see more: https://stackoverflow.com/questions/55010252/why-the-iteration-is-done-by-i6-every-time-and-why-the-condition-is-ii-n-for
func isPrime(n int) bool {
	if n <= 1 {
		return false
	}

	if n <= 3 {
		return true
	}

	if n%2 == 0 || n%3 == 0 {
		return false
	}

	for i := 5; i*i <= n; i += 6 {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
	}

	return true
}
