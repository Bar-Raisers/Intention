workspace(name = "intention")

load("@//:.bazel/base.bzl", "workspace_base")

workspace_base()

load("@//:.bazel/deployment.bzl", "workspace_deployment")
load("@//:.bazel/gazelle.bzl", "workspace_gazelle")
load("@//:.bazel/go.bzl", "workspace_go")
load("@//:.bazel/proto.bzl", "workspace_proto")

workspace_deployment()

workspace_go()

workspace_gazelle()

workspace_proto()
