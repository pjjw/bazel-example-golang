load("@bazel_tools//tools/build_rules/go:def.bzl", "go_prefix", "go_binary")

package(
    default_visibility = ["//visibility:public"],
)

go_prefix("github.com/laramiel/bazel-example-golang")

go_binary(
	name ="hello",
	srcs = [ "hello.go" ],
	deps = [
	    "//local",
	],
)

go_binary(
	name ="remote",
	srcs = [ "remote.go" ],
	deps = [
	    "@remote//:remote",
	],
)

go_binary(
	name ="bare",
	srcs = [ "bare.go" ],
	deps = [
	    "@bare//:bare",
	],
)

go_binary(
	name ="subbare",
	srcs = [ "subbare.go" ],
	deps = [
	    "//submodule/src:bare",
	],
)

go_binary(
	name ="subremote",
	srcs = [ "subremote.go" ],
	deps = [
	    "//submodule/src/remote",
	],
)
