"""Workspace Configuration for Repository Deployment"""

load("@aspect_bazel_lib//lib:repositories.bzl", "aspect_bazel_lib_dependencies")
load("@bazel_skylib//:workspace.bzl", "bazel_skylib_workspace")
load("@rules_oci//oci:dependencies.bzl", "rules_oci_dependencies")
load("@rules_oci//oci:pull.bzl", "oci_pull")
load("@rules_oci//oci:repositories.bzl", "LATEST_CRANE_VERSION", "oci_register_toolchains")
load("@rules_pkg//:deps.bzl", "rules_pkg_dependencies")

def workspace_deployment():
    aspect_bazel_lib_dependencies()
    bazel_skylib_workspace()
    rules_pkg_dependencies()
    rules_oci_dependencies()

    oci_register_toolchains(
        name = "oci",
        crane_version = LATEST_CRANE_VERSION,
    )

    pull_images()

def pull_images():
    oci_pull(
        name = "distroless_base",
        digest = "sha256:ccaef5ee2f1850270d453fdf700a5392534f8d1a8ca2acda391fbb6a06b81c86",
        image = "gcr.io/distroless/base",
        platforms = [
            "linux/amd64",
        ],
    )
