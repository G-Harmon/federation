load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "override.go",
        "unicode.go",
    ],
    importpath = "golang.org/x/text/encoding/unicode",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/golang.org/x/text/encoding:go_default_library",
        "//vendor/golang.org/x/text/encoding/internal:go_default_library",
        "//vendor/golang.org/x/text/encoding/internal/identifier:go_default_library",
        "//vendor/golang.org/x/text/internal/utf8internal:go_default_library",
        "//vendor/golang.org/x/text/runes:go_default_library",
        "//vendor/golang.org/x/text/transform:go_default_library",
    ],
)

go_test(
    name = "go_default_test",
    srcs = ["unicode_test.go"],
    importpath = "golang.org/x/text/encoding/unicode",
    library = ":go_default_library",
    deps = [
        "//vendor/golang.org/x/text/encoding:go_default_library",
        "//vendor/golang.org/x/text/encoding/charmap:go_default_library",
        "//vendor/golang.org/x/text/encoding/internal/enctest:go_default_library",
        "//vendor/golang.org/x/text/transform:go_default_library",
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
        "//vendor/golang.org/x/text/encoding/unicode/utf32:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
