git_repository(
    name = "io_bazel_rules_go",
    remote = "https://github.com/bazelbuild/rules_go.git",
    commit = "373feb6",
)

load("@io_bazel_rules_go//go:def.bzl", "go_repositories")

go_repositories()

git_repository(
  name = "ws_remote",
  remote = "https://github.com/pjjw/bazel-example-golang-remote.git",
  commit = "6422e74",
)


BARE_BUILD = """
load("@io_bazel_rules_go//go:def.bzl", "go_prefix", "go_library")

go_prefix("github.com/pjjw/bazel-example-golang-bare")

go_library(
    name = "bare",
    srcs = [ "bare.go" ],
    visibility = ["//visibility:public"],
)

"""

new_git_repository(
  name = "ws_bare",
  remote = "https://github.com/pjjw/bazel-example-golang-bare.git",
  commit = "3bd848f",
  build_file_content = BARE_BUILD
)
