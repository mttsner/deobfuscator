FROM golang:1.16-alpine

RUN apk update

RUN apk -y install git

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
        golang.org/x/lint/golint \
        golang.org/x/tools/gopls \
    # Clean up.
    && apt-get autoremove -y \
    && apt-get clean -y \
    && rm -rf /var/lib/apt/lists/*