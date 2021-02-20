FROM golang:1.16-alpine

# Install Go tools.
RUN apk update \
    # Install other tools.
    && go get -u -v \
        github.com/mdempsky/gocode \
        github.com/uudashr/gopkgs/v2/cmd/gopkgs \
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
        github.com/stamblerre/gocode \
        github.com/rogpeppe/godef \
        github.com/sqs/goreturns \
        golang.org/x/lint/golint && \
    GO111MODULE=on go get -u -v \
        golang.org/x/tools/gopls@v0.6.5

ENV GOPATH=/workspace/go
ENV GO111MODULE=on
ENV GOPRIVATE=github.com/notnoobmaster

RUN apk add git

RUN git config --global url."https://${username}:${personal_access_token}@github.com/notnoobmaster".insteadOf "https://github.com/notnoobmaster"
