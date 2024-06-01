package task2

import "testing"

func TestSplitNumbers(t *testing.T) {
	tests := []struct {
		name       string
		numbers    []int
		primes     []int
		composites []int
	}{
		{
			name:       "Базовый тест",
			numbers:    []int{2, 3, 4, 5, 6, 7, 8, 9, 10},
			primes:     []int{2, 3, 5, 7},
			composites: []int{4, 6, 8, 9, 10},
		},
		{
			name:       "С нулём и единицей",
			numbers:    []int{0, 1, 2, 3, 4, 5, 6},
			primes:     []int{2, 3, 5},
			composites: []int{0, 1, 4, 6},
		},
		{
			name:       "Все простые",
			numbers:    []int{2, 3, 5, 7, 11, 13, 17, 19, 23},
			primes:     []int{2, 3, 5, 7, 11, 13, 17, 19, 23},
			composites: []int{},
		},
		{
			name:       "Все составные",
			numbers:    []int{4, 6, 8, 9, 10, 12, 14, 15, 16},
			primes:     []int{},
			composites: []int{4, 6, 8, 9, 10, 12, 14, 15, 16},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			primes, composites := splitNumbers(tt.numbers)
			if !equalSlices(primes, tt.primes) {
				t.Errorf("splitNumbers() простые = %v, ожидаем %v", primes, tt.primes)
			}
			if !equalSlices(composites, tt.composites) {
				t.Errorf("splitNumbers() составные = %v, ожидаем %v", composites, tt.composites)
			}
		})
	}
}

// Сравниваем слайсы
func equalSlices(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
