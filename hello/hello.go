package main

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/yuechen-sys/go_tutorial/math/simplemath"
)

func add_one(a uint32) uint32 {
	return a + 1
}

func add_one_pointer(a *int) {
	*a = *a + 1
}

func create_pointer() *int {
	a := 1
	return &a
}

// import "rsc.io/quote"

func main() {
	a := create_pointer()
	var b int = 2
	var c int32 = int32(simplemath.Max(*a, b))

	add_one_pointer(a)
	println("a: ", *a)

	var new_uuid = uuid.New()

	fmt.Printf("Hello World%d\n", c)
	fmt.Println("uuid: ", add_one(new_uuid.ID()))
}
