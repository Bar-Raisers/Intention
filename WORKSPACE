workspace(name = "intention")

load("@//:.bazel/base.bzl", "workspace_base")

workspace_base()

load("@//:.bazel/gazelle.bzl", "workspace_gazelle")
load("@//:.bazel/go.bzl", "go_dependencies", "workspace_go")
load("@//:.bazel/proto.bzl", "workspace_proto")

workspace_go()

workspace_gazelle()

workspace_proto()

# gazelle:repository_macro .bazel/go.bzl%go_dependencies
go_dependencies()
