load("@bazel_tools//tools/build_rules/go:def.bzl", "go_prefix", "go_binary")

package(
    default_visibility = ["//visibility:public"],
)

go_prefix("github.com/lar/bazel-example-golang")

go_binary(
	name ="hello",
	srcs = [ "hello.go" ],
	deps = [ "//local" ],
)