load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["main.go"],
    importpath = "github.com/pkg/sftp/examples/streaming-write-benchmark",
    visibility = ["//visibility:private"],
    deps = [
        "//vendor/github.com/pkg/sftp:go_default_library",
        "//vendor/golang.org/x/crypto/ssh:go_default_library",
        "//vendor/golang.org/x/crypto/ssh/agent:go_default_library",
    ],
)

go_binary(
    name = "streaming-write-benchmark",
    importpath = "github.com/pkg/sftp/examples/streaming-write-benchmark",
    library = ":go_default_library",
    visibility = ["//visibility:public"],
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
