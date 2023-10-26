"""Gazelle Workspace Configuration for Repository"""

load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies")

def workspace_gazelle():
    gazelle_dependencies()
