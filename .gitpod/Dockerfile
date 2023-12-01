FROM gitpod/workspace-base:latest

ENV GOROOT=$HOME/go
ENV GOPATH=$HOME/go-packages
ENV PATH=$GOROOT/bin:$GOPATH/bin:$PATH

ENV GO_VERSION=1.21.3
ENV KUBECTL_VERSION=1.28.4
ENV MINIKUBE_VERSION=1.32.0
ENV SKAFFOLD_VERSION=2.9.0
ENV USE_BAZEL_VERSION=6.4.0

USER gitpod
SHELL ["/bin/bash", "-o", "pipefail", "-c"]

# Install Go
RUN curl -fsSL "https://dl.google.com/go/go$GO_VERSION.linux-amd64.tar.gz" | tar -xzs -C "$HOME"

# Install Bazelisk
RUN go install github.com/bazelbuild/bazelisk@latest;

# Install Trunk
RUN curl -fsSL https://get.trunk.io | bash -s -- -y;

# Install Skaffold
RUN curl --create-dirs -Lo "$HOME/.local/bin/skaffold" "https://storage.googleapis.com/skaffold/releases/v$SKAFFOLD_VERSION/skaffold-linux-amd64" && \
    chmod +x "$HOME/.local/bin/skaffold"

# Install Kubectl
RUN curl --create-dirs -Lo "$HOME/.local/bin/kubectl" "https://dl.k8s.io/release/v$KUBECTL_VERSION/bin/linux/amd64/kubectl" && \
    chmod +x "$HOME/.local/bin/kubectl"

# Install Minikube
RUN curl --create-dirs -Lo "$HOME/.local/bin/minikube" "https://storage.googleapis.com/minikube/releases/v$MINIKUBE_VERSION/minikube-linux-amd64" && \
    chmod +x "$HOME/.local/bin/minikube"
