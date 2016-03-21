load("@io_bazel_rules_go//go:def.bzl", "go_binary")

package(
    default_visibility = ["//visibility:public"],
)

go_binary(
    name = "bare",
    srcs = ["bare.go"],
    deps = [
        "//submodule/src:bare",
    ],
)

go_binary(
    name = "remote",
    srcs = ["remote.go"],
    deps = [
        "//submodule/src/remote",
    ],
)
