load("@io_bazel_rules_go//go:def.bzl", "go_prefix", "go_binary")

package(
    default_visibility = ["//visibility:public"],
)

go_prefix("github.com/pjjw/bazel-example-golang/")

go_binary(
    name = "hello",
    srcs = ["hello.go"],
    deps = [
        "//local",
    ],
)

go_binary(
    name = "remote",
    srcs = ["remote.go"],
    deps = [
        "@ws_remote//:remote",
    ],
)

go_binary(
    name = "bare",
    srcs = ["bare.go"],
    deps = [
        "@ws_bare//:bare",
    ],
)
