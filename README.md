# bazel-example-golang

An example project demonstrating bazel build rules for go with
a local binary (go_binary) and a local library (go_library).

This requires bazel release > `bazel-0.20`.

```
$ git clone https://github.com/laramiel/bazel-example-golang.git
$ cd bazel-example-golang
$ git submodule init
$ git submodule foreach git pull origin master

$ bazel run :hello

$ bazel run :local
$ bazel run :remote
$ bazel run :bare

$ git submodule update --init --recursive
$ bazel run //submodule:bare
$ bazel run //submodule:remote

$ bazel run //with_vendor:remote
$ bazel run //with_vendor:bare

```

# Explanation

This repository works in conjunction with other git repositories to
provide a demonstration of Bazel BUILD rules for golang.

* [bazel-example-golang](https://github.com/laramiel/bazel-example-golang)  
  The main repository, demonstrating binary, local, and remote repository usage.

* [bazel-example-golang-remote](https://github.com/laramiel/bazel-example-golang-remote)  
  An example remote repository which includes Bazel rules.

* [bazel-example-golang-bare](https://github.com/laramiel/bazel-example-golang-bare)
  An example remote repository without any Bazel rules.


Each of the remote repositories is linked into `go_binary()` targets in several
ways:

1. Using the WORKSPACE file to define `git_repository()` rules.
   This accounts for the `:remote` and `:bare` targets.

2. Using git submodules to import the repository source, and
   using those paths. This accounts for the `//submodule:bare`
   and `//submodule:remote` targets.

3. Using the VENDOR extension introduced in go1.6.

