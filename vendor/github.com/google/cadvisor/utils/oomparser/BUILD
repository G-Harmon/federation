load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = ["oomparser.go"],
    importpath = "github.com/google/cadvisor/utils/oomparser",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/euank/go-kmsg-parser/kmsgparser:go_default_library",
        "//vendor/github.com/golang/glog:go_default_library",
    ],
)

go_test(
    name = "go_default_test",
    srcs = ["oomparser_test.go"],
    importpath = "github.com/google/cadvisor/utils/oomparser",
    library = ":go_default_library",
    deps = [
        "//vendor/github.com/euank/go-kmsg-parser/kmsgparser:go_default_library",
        "//vendor/github.com/stretchr/testify/assert:go_default_library",
    ],
)

filegroup(
    name = "package-srcs",
    srcs = glob(["**"]),
    tags = ["automanaged"],
    visibility = ["//visibility:private"],
)

filegroup(
    name = "all-srcs",
    srcs = [
        ":package-srcs",
        "//vendor/github.com/google/cadvisor/utils/oomparser/oomexample:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
