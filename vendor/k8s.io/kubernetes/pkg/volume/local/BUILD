package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_library(
    name = "go_default_library",
    srcs = [
        "doc.go",
        "local.go",
    ],
    importpath = "k8s.io/kubernetes/pkg/volume/local",
    deps = [
        "//vendor/github.com/golang/glog:go_default_library",
        "//vendor/k8s.io/api/core/v1:go_default_library",
        "//vendor/k8s.io/apimachinery/pkg/apis/meta/v1:go_default_library",
        "//vendor/k8s.io/apimachinery/pkg/types:go_default_library",
        "//vendor/k8s.io/client-go/kubernetes/scheme:go_default_library",
        "//vendor/k8s.io/client-go/tools/record:go_default_library",
        "//vendor/k8s.io/kubernetes/pkg/kubelet/events:go_default_library",
        "//vendor/k8s.io/kubernetes/pkg/util/keymutex:go_default_library",
        "//vendor/k8s.io/kubernetes/pkg/util/mount:go_default_library",
        "//vendor/k8s.io/kubernetes/pkg/util/strings:go_default_library",
        "//vendor/k8s.io/kubernetes/pkg/volume:go_default_library",
        "//vendor/k8s.io/kubernetes/pkg/volume/util:go_default_library",
        "//vendor/k8s.io/kubernetes/pkg/volume/validation:go_default_library",
    ],
)

go_test(
    name = "go_default_test",
    srcs = ["local_test.go"],
    importpath = "k8s.io/kubernetes/pkg/volume/local",
    library = ":go_default_library",
    deps = [
        "//vendor/k8s.io/api/core/v1:go_default_library",
        "//vendor/k8s.io/apimachinery/pkg/apis/meta/v1:go_default_library",
        "//vendor/k8s.io/apimachinery/pkg/types:go_default_library",
        "//vendor/k8s.io/client-go/util/testing:go_default_library",
        "//vendor/k8s.io/kubernetes/pkg/volume:go_default_library",
        "//vendor/k8s.io/kubernetes/pkg/volume/testing:go_default_library",
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
)
