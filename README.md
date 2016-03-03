# bazel-example-golang

An example project demonstrating bazel build rules for go with
a local binary (go_binary) and a local library (go_library).

This requires bazel release > `bazel-0.20`.

```
$ bazel run :hello

$ bazel run :remote
$ bazel run :bare
$ bazel run :subbare
$ bazel run :subremote
```

