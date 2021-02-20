FROM gitpod/workspace-full

ENV GO_VERSION=1.16

RUN curl -fsSL https://storage.googleapis.com/golang/go$GO_VERSION.linux-amd64.tar.gz | tar xzs