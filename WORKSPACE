load("@bazel_tools//tools/build_rules/go:def.bzl", "go_repositories")

go_repositories()

git_repository(
  name = "remote",
  remote = "https://github.com/laramiel/bazel-example-golang-remote.git",
  commit = "8f2e405",
)


BARE_BUILD = """
load("@bazel_tools//tools/build_rules/go:def.bzl", "go_prefix", "go_library")

go_prefix("github.com/laramiel/bazel-example-golang-bare")

go_library(
    name = "bare",
    srcs = [ "bare.go" ],
    visibility = ["//visibility:public"],
)

"""

new_git_repository(
  name = "bare",
  remote = "https://github.com/laramiel/bazel-example-golang-bare.git",
  commit = "3bd848f",
  build_file_content = BARE_BUILD
)
