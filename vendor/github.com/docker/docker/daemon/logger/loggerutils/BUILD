load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "log_tag.go",
        "rotatefilewriter.go",
    ],
    importpath = "github.com/docker/docker/daemon/logger/loggerutils",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/docker/docker/daemon/logger:go_default_library",
        "//vendor/github.com/docker/docker/pkg/pubsub:go_default_library",
        "//vendor/github.com/docker/docker/pkg/templates:go_default_library",
    ],
)

go_test(
    name = "go_default_test",
    srcs = ["log_tag_test.go"],
    importpath = "github.com/docker/docker/daemon/logger/loggerutils",
    library = ":go_default_library",
    deps = ["//vendor/github.com/docker/docker/daemon/logger:go_default_library"],
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
