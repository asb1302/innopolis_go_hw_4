package task3

import (
	"sort"
	"testing"
)

func TestMergeChannels(t *testing.T) {
	tests := []struct {
		name     string
		ch1Input []int
		ch2Input []int
		expected []int
	}{
		{
			name:     "Элементы в каждом канале не пересекаются",
			ch1Input: []int{1, 2, 3},
			ch2Input: []int{4, 5, 6},
			expected: []int{1, 2, 3, 4, 5, 6},
		},
		{
			name:     "Элементы в каждом канале пересекаются",
			ch1Input: []int{1, 3, 5},
			ch2Input: []int{2, 4, 6},
			expected: []int{1, 2, 3, 4, 5, 6},
		},
		{
			name:     "Канал 1 пустой",
			ch1Input: []int{},
			ch2Input: []int{1, 2, 3},
			expected: []int{1, 2, 3},
		},
		{
			name:     "Канал 2 пустой",
			ch1Input: []int{1, 2, 3},
			ch2Input: []int{},
			expected: []int{1, 2, 3},
		},
		{
			name:     "Оба канала пустые",
			ch1Input: []int{},
			ch2Input: []int{},
			expected: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ch1 := make(chan int)
			ch2 := make(chan int)

			go func() {
				for _, v := range tt.ch1Input {
					ch1 <- v
				}
				close(ch1)
			}()

			go func() {
				for _, v := range tt.ch2Input {
					ch2 <- v
				}
				close(ch2)
			}()

			var result []int
			for v := range mergeChannels(ch1, ch2) {
				result = append(result, v)
			}

			sort.Ints(result)
			sort.Ints(tt.expected)

			if !equalSlices(result, tt.expected) {
				t.Errorf("ожидаем %v, получено %v", tt.expected, result)
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
