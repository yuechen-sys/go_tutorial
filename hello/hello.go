package hello

import "fmt"

import "github.com/yuechen-sys/go_tutorial/math/simplemath"

// import "rsc.io/quote"

func main() {
	var a int = 1
	var b int = 2
	var c int32

	c = int32(simplemath.Max(a, b))
	fmt.Printf("Hello World%d\n", c)
}
