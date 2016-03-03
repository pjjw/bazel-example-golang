# bazel-example-golang

An example project demonstrating bazel build rules for go with
a local binary (go_binary) and a local library (go_library).

This requires bazel release > `bazel-0.20`.

```
$ git clone https://github.com/laramiel/bazel-example-golang.git
$ cd bazel-example-golang

$ bazel run :hello

$ bazel run :local
$ bazel run :remote
$ bazel run :bare

$ git submodule update --init --recursive
$ bazel run //submodule:bare
$ bazel run //submodule:remote
```

