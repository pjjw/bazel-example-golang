load("@io_bazel_rules_go//go:def.bzl", "go_binary")

package(
    default_visibility = ["//visibility:public"],
)

# BAZEL does not work correctly with BUILD files rooted elsewhere than
# the actual code directory.
#
#go_binary(
#	name ="bare",
#	srcs = [ "bare.go" ],
#	deps = [
#	    "//vendor/github.com/laramiel:bare",
#	],
#)

go_binary(
    name = "remote",
    srcs = ["remote.go"],
    deps = [
        "//vendor/github.com/laramiel/bazel-example-golang-remote:remote",
    ],
)
