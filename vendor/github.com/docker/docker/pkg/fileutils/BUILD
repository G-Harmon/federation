load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "fileutils.go",
    ] + select({
        "@io_bazel_rules_go//go/platform:darwin_amd64": [
            "fileutils_darwin.go",
        ],
        "@io_bazel_rules_go//go/platform:linux_amd64": [
            "fileutils_unix.go",
        ],
        "@io_bazel_rules_go//go/platform:windows_amd64": [
            "fileutils_windows.go",
        ],
        "//conditions:default": [],
    }),
    importpath = "github.com/docker/docker/pkg/fileutils",
    visibility = ["//visibility:public"],
    deps = ["//vendor/github.com/sirupsen/logrus:go_default_library"],
)

go_test(
    name = "go_default_test",
    srcs = ["fileutils_test.go"],
    importpath = "github.com/docker/docker/pkg/fileutils",
    library = ":go_default_library",
    deps = [
        "//vendor/github.com/stretchr/testify/assert:go_default_library",
        "//vendor/github.com/stretchr/testify/require:go_default_library",
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
    srcs = [":package-srcs"],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
