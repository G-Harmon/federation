load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "block.go",
        "cipher.go",
        "const.go",
    ],
    importpath = "golang.org/x/crypto/blowfish",
    visibility = ["//visibility:public"],
)

go_test(
    name = "go_default_test",
    srcs = ["blowfish_test.go"],
    importpath = "golang.org/x/crypto/blowfish",
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
