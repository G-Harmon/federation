package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "go_default_library",
    srcs = ["doc.go"],
    importpath = "k8s.io/apiserver/pkg/registry",
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
        "//vendor/k8s.io/apiserver/pkg/registry/generic:all-srcs",
        "//vendor/k8s.io/apiserver/pkg/registry/rest:all-srcs",
    ],
    tags = ["automanaged"],
)
