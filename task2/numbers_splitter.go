package task2

/*
Напишите функцию разделения массива чисел на массивы простых и составных чисел. Для записи в массивы используйте два
разных канала и горутины.

Важно, чтобы были использованы владельцы каналов.
*/
func splitNumbers(numbers []int) ([]int, []int) {
	primeChan := make(chan int)
	compositeChan := make(chan int)
	done := make(chan bool)

	// Владелец primeChan
	go func() {
		defer close(primeChan)
		for _, n := range numbers {
			if isPrime(n) {
				primeChan <- n
			}
		}
	}()

	// Владелец compositeChan
	go func() {
		defer close(compositeChan)
		for _, n := range numbers {
			if !isPrime(n) {
				compositeChan <- n
			}
		}
	}()

	var primes []int
	var composites []int

	go func() {
		for n := range primeChan {
			primes = append(primes, n)
		}
		done <- true
	}()

	go func() {
		for n := range compositeChan {
			composites = append(composites, n)
		}
		done <- true
	}()

	<-done
	<-done

	return primes, composites
}

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
