"""Go Workspace Configuration for Repository"""

load("@bazel_gazelle//:deps.bzl", "go_repository")
load("@io_bazel_rules_go//go:deps.bzl", "go_register_toolchains", "go_rules_dependencies")

def workspace_go():
    go_rules_dependencies()
    go_register_toolchains(version = "1.21.3")

    load_dependencies()

def load_dependencies():
    go_repository(
        name = "org_golang_google_grpc",
        build_file_proto_mode = "disable",
        importpath = "google.golang.org/grpc",
        sum = "h1:f+PlOh7QV4iIJkPrx5NQ7qaNGFQ3OTse67yaDHfju4E=",
        version = "v1.41.0",
    )
