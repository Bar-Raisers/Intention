FROM gitpod/workspace-base:latest

USER gitpod

ENV GOROOT=$HOME/go
ENV GOPATH=$HOME/go-packages
ENV PATH=$GOROOT/bin:$GOPATH/bin:$PATH

ENV GO_VERSION=1.21.3
ENV USE_BAZEL_VERSION=6.4.0

SHELL ["/bin/bash", "-o", "pipefail", "-c"]

# Install Go
RUN curl -fsSL "https://dl.google.com/go/go$GO_VERSION.linux-amd64.tar.gz" | tar -xzs -C "$HOME"

# Install Bazelisk
RUN go install github.com/bazelbuild/bazelisk@latest;

# Install Trunk
RUN curl -fsSL https://get.trunk.io | bash -s -- -y;
