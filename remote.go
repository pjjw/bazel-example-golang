package main

import (
	"fmt"

	// HACK ALERT
	// The correct import is this:
	//   "github.com/laramiel/bazel-example-golang-remote/remote"
	// But the bazel build rules for external golang packages like git_repository() are busted.
	// We hack around it by importing where those external dependencies end up:
	"github.com/pjjw/bazel-example-golang/external/ws_remote/remote"
)

func main() {
	fmt.Println("Hello ", remote.World())
}
