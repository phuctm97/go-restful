FROM golang:1

# Configure to avoid redundant error "debconf: delaying package configuration, since apt-utils is not installed".
# See https://code.visualstudio.com/docs/remote/containers-advanced#_debconf-delaying-package-configuration-since-aptutils-is-not-installed.
ENV DEBIAN_FRONTEND=noninteractive
RUN apt-get update && apt-get -y install --no-install-recommends apt-utils

# Verify git, process tools, lsb-release (common in install instructions for CLIs) installed.
RUN apt-get -y install git iproute2 procps lsb-release

# Enable Go Modules.
ENV GO111MODULE=on

# Install essential tools for Go development.
RUN go get \
        golang.org/x/tools/gopls@latest \
        github.com/mdempsky/gocode \
        github.com/uudashr/gopkgs/cmd/gopkgs \
        github.com/ramya-rao-a/go-outline \
        github.com/acroca/go-symbols \
        golang.org/x/tools/cmd/guru \
        golang.org/x/tools/cmd/gorename \
        github.com/cweill/gotests/... \
        github.com/fatih/gomodifytags \
        github.com/josharian/impl \
        github.com/davidrjenni/reftools/cmd/fillstruct \
        github.com/haya14busa/goplay/cmd/goplay \
        github.com/godoctor/godoctor \
        github.com/go-delve/delve/cmd/dlv \
        github.com/rogpeppe/godef \
        golang.org/x/tools/cmd/goimports \
        golang.org/x/lint/golint \
    && go build -o $GOPATH/bin/gocode-gomod github.com/stamblerre/gocode

# Clean up.
RUN apt-get autoremove -y && apt-get clean -y \
    && rm -rf /var/lib/apt/lists/*

# Revert workaround of avoiding redundant error "debconf: delaying package configuration, since apt-utils is not installed".
# See https://code.visualstudio.com/docs/remote/containers-advanced#_debconf-delaying-package-configuration-since-aptutils-is-not-installed.
ENV DEBIAN_FRONTEND=dialog

# Expose service ports.
EXPOSE 8000
