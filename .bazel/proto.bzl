"""Proto Workspace Configuration for Repository"""

load("@rules_proto//proto:repositories.bzl", "rules_proto_dependencies", "rules_proto_toolchains")

def workspace_proto():
    rules_proto_dependencies()
    rules_proto_toolchains()
