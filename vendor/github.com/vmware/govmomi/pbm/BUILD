load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "client.go",
        "pbm_util.go",
    ],
    importpath = "github.com/vmware/govmomi/pbm",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/vmware/govmomi/pbm/methods:go_default_library",
        "//vendor/github.com/vmware/govmomi/pbm/types:go_default_library",
        "//vendor/github.com/vmware/govmomi/vim25:go_default_library",
        "//vendor/github.com/vmware/govmomi/vim25/soap:go_default_library",
        "//vendor/github.com/vmware/govmomi/vim25/types:go_default_library",
    ],
)

go_test(
    name = "go_default_test",
    srcs = ["client_test.go"],
    importpath = "github.com/vmware/govmomi/pbm",
    library = ":go_default_library",
    deps = [
        "//vendor/github.com/vmware/govmomi:go_default_library",
        "//vendor/github.com/vmware/govmomi/pbm/types:go_default_library",
        "//vendor/github.com/vmware/govmomi/property:go_default_library",
        "//vendor/github.com/vmware/govmomi/view:go_default_library",
        "//vendor/github.com/vmware/govmomi/vim25/mo:go_default_library",
        "//vendor/github.com/vmware/govmomi/vim25/soap:go_default_library",
        "//vendor/github.com/vmware/govmomi/vim25/types:go_default_library",
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
        "//vendor/github.com/vmware/govmomi/pbm/methods:all-srcs",
        "//vendor/github.com/vmware/govmomi/pbm/types:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
