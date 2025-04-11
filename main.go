package main
import (
	"fmt"
	"os"
	"strconv"
)

// возвращает n-й элемент последовательности по переданному правилу
func recursiveSequence(n int, rule func(int, int, int) int) int {
	seq := []int{1, 2, 3}
	if n < 3 {
		return seq[n]
	}

	for i := 3; i <= n; i++ {
		next := rule(seq[i-1], seq[i-2], seq[i-3])
		seq = append(seq, next)
	}

	return seq[n]
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Укажите номер элемента последовательности!")
		return
	}

	n, _ := strconv.Atoi(os.Args[1])

	rule := func(a, b, c int) int {
		return (a * b - c) % 10 + 1
	}

	result := recursiveSequence(n, rule)
	fmt.Printf("%d-й элемент последовательности: %d\n", n, result)
}
