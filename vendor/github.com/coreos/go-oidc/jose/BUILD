load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "claims.go",
        "doc.go",
        "jose.go",
        "jwk.go",
        "jws.go",
        "jwt.go",
        "sig.go",
        "sig_rsa.go",
    ],
    importpath = "github.com/coreos/go-oidc/jose",
    visibility = ["//visibility:public"],
)

go_test(
    name = "go_default_test",
    srcs = [
        "claims_test.go",
        "jwk_test.go",
        "jws_test.go",
        "jwt_test.go",
    ],
    importpath = "github.com/coreos/go-oidc/jose",
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
