load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "backdoor.go",
        "gccgoinstallation.go",
        "importer.go",
        "parser.go",
    ],
    importpath = "golang.org/x/tools/go/internal/gccgoimporter",
    visibility = ["//vendor/golang.org/x/tools/go:__subpackages__"],
)

go_test(
    name = "go_default_test",
    srcs = [
        "gccgoinstallation_test.go",
        "importer19_test.go",
        "importer_test.go",
        "parser_test.go",
        "testenv_test.go",
    ],
    data = glob(["testdata/**"]),
    importpath = "golang.org/x/tools/go/internal/gccgoimporter",
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
