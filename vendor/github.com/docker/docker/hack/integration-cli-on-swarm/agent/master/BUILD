load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "call.go",
        "master.go",
        "set.go",
    ],
    importpath = "github.com/docker/docker/hack/integration-cli-on-swarm/agent/master",
    visibility = ["//visibility:private"],
    deps = [
        "//vendor/github.com/bfirsh/funker-go:go_default_library",
        "//vendor/github.com/docker/docker/hack/integration-cli-on-swarm/agent/types:go_default_library",
    ],
)

go_binary(
    name = "master",
    importpath = "github.com/docker/docker/hack/integration-cli-on-swarm/agent/master",
    library = ":go_default_library",
    visibility = ["//visibility:public"],
)

go_test(
    name = "go_default_test",
    srcs = ["set_test.go"],
    importpath = "github.com/docker/docker/hack/integration-cli-on-swarm/agent/master",
    library = ":go_default_library",
)

filegroup(
    name = "package-srcs",
    srcs = glob(["**"]),
    tags = ["automanaged"],
    visibility = ["//visibility:private"],
)

filegroup(
    name = "all-srcs",
    srcs = [":package-srcs"],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
