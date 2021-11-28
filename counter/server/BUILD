load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library", "go_test")

go_binary(
    name = "server",
    embed = [":server_lib"],
    visibility = ["//visibility:public"],
)

go_library(
    name = "server_lib",
    srcs = [
        "counter_server.go",
        "main.go",
    ],
    importpath = "github.com/mikehelmick/grpc-test/counter/server",
    visibility = ["//visibility:private"],
    deps = [
        "//counter/proto",
        "@org_golang_google_grpc//:go_default_library",
    ],
)

go_test(
    name = "server_test",
    srcs = ["counter_server_test.go"],
    embed = [":server_lib"],
    deps = ["//counter/proto"],
)
