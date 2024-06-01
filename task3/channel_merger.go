package task3

/*
Реализуйте функцию слияния двух каналов в один.
*/
func mergeChannels(ch1, ch2 <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)
		for ch1 != nil || ch2 != nil {
			select {
			case v, ok := <-ch1:
				if !ok {
					ch1 = nil
					continue
				}
				out <- v
			case v, ok := <-ch2:
				if !ok {
					ch2 = nil
					continue
				}
				out <- v
			}
		}
	}()
	return out
}
