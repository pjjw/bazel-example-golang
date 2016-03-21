package main

import (
	"fmt"

	"github.com/laramiel/bazel-example-golang/local/local"
)

func main() {
	fmt.Println("Hello World")

	fmt.Println("Hello", local.World())
}
