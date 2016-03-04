package main

import (
	"fmt"

	// HACK ALERT
	// The correct import is this:
	//   "github.com/laramiel/bazel-example-golang-bare/bare"
	// But the bazel build rules for subversion packages don't support alternate
	// go_prefix() commands.
	// We hack around it by importing where those external dependencies end up:
	"github.com/laramiel/bazel-example-golang/submodule/src/bare"
)

func main() {
	fmt.Println("Hello", bare.World())
}
