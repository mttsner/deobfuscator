FROM golang:1.16-alpine

ENV GOPATH=$HOME/go-packages
ENV GOROOT=$HOME/go

RUN go get -u -v \
        github.com/mdempsky/gocode \
        github.com/uudashr/gopkgs/cmd/gopkgs \
        github.com/ramya-rao-a/go-outline \
        github.com/acroca/go-symbols \
        golang.org/x/tools/cmd/guru \
        golang.org/x/tools/cmd/gorename \
        github.com/fatih/gomodifytags \
        github.com/haya14busa/goplay/cmd/goplay \
        github.com/josharian/impl \
        github.com/tylerb/gotype-live \
        github.com/rogpeppe/godef \
        github.com/zmb3/gogetdoc \
        golang.org/x/tools/cmd/goimports \
        github.com/sqs/goreturns \
        winterdrache.de/goformat/goformat \
        golang.org/x/lint/golint \
        github.com/cweill/gotests/... \
        github.com/alecthomas/gometalinter \
        honnef.co/go/tools/... \
        github.com/mgechev/revive \
        github.com/sourcegraph/go-langserver \
        github.com/go-delve/delve/cmd/dlv \
        github.com/davidrjenni/reftools/cmd/fillstruct \
        github.com/godoctor/godoctor && \
    GO111MODULE=on go get -u -v \
        golang.org/x/tools/gopls@v0.6.5 && \
    go get -u -v -d github.com/stamblerre/gocode && \
    go build -o $GOPATH/bin/gocode-gomod github.com/stamblerre/gocode && \
    sudo rm -rf $GOPATH/src $GOPATH/pkg /home/gitpod/.cache/go /home/gitpod/.cache/go-build
# user Go packages
ENV GOPATH=/workspace/go \
    PATH=/workspace/go/bin:$PATH