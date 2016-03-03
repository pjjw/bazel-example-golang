load("@bazel_tools//tools/build_rules/go:def.bzl", "go_prefix", "go_binary")

package(
    default_visibility = ["//visibility:public"],
)

go_binary(
	name ="bare",
	srcs = [ "bare.go" ],
	deps = [
	    "//submodule/src:bare",
	],
)

go_binary(
	name ="remote",
	srcs = [ "remote.go" ],
	deps = [
	    "//submodule/src/remote",
	],
)
