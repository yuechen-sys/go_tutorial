package simplemath

import "fmt"

// Max returns the larger number among `a` and `b`
func Max(a, b int) int {
	fmt.Println("This is a public function")
	if a > b {
		return a
	}

	return b
}

func max(a, b int) int {
	fmt.Println("This is a private function")
	if a > b {
		return a
	}

	return b
}
