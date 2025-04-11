package main
import ("testing")

func testRule(a, b, c int) int {
	return (a * b - c) % 10 + 1
}

func TestRecursiveSequence(t *testing.T) {
	tests := []struct {
		n      int
		result int
	}{
		{0, 1},  // seq[0] = 1
		{1, 2},  // seq[1] = 2
		{2, 3},  // seq[2] = 3
		{3, 6},  // (3*2 - 1)%10 + 1 = 5%10 + 1 = 6
		{4, 7},  // (6*3 - 2)%10 + 1 = 16%10 + 1 = 6 + 1 = 7
	}

	for _, test := range tests {
		t.Run("TestRecursiveSequence", func(t *testing.T) {
			got := recursiveSequence(test.n, testRule)
			if got != test.result {
				t.Errorf("Для n=%d ожидали %d, но получили %d", test.n, test.result, got)
			}
		})
	}
}