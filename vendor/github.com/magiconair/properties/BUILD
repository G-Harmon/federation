load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "decode.go",
        "doc.go",
        "integrate.go",
        "lex.go",
        "load.go",
        "parser.go",
        "properties.go",
        "rangecheck.go",
    ],
    importpath = "github.com/magiconair/properties",
    visibility = ["//visibility:public"],
)

go_test(
    name = "go_default_test",
    srcs = [
        "benchmark_test.go",
        "decode_test.go",
        "example_test.go",
        "integrate_test.go",
        "load_test.go",
        "properties_test.go",
    ],
    importpath = "github.com/magiconair/properties",
    library = ":go_default_library",
    deps = ["//vendor/github.com/magiconair/properties/_third_party/gopkg.in/check.v1:go_default_library"],
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
