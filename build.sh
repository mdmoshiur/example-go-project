#! /bin/sh
set -e

if ! [ -x "$(command -v go)" ]; then
    echo "go is not installed"
    exit
fi
if ! [ -x "$(command -v git)" ]; then
    echo "git is not installed"
    exit
fi

COMMIT=`git rev-parse --short HEAD`
TAG=$(git describe --exact-match --abbrev=0 --tags ${COMMIT} 2> /dev/null || true)

if [ -z "$TAG" ]; then
    VERSION=$COMMIT
else
    VERSION=$TAG
fi

if [ -n "$(git diff --shortstat 2> /dev/null | tail -n1)" ]; then
    VERSION="${VERSION}-dirty"
fi

BUILD_DATE=$(date "+%Y-%m-%d")
BRANCH_NAME=$(git rev-parse --abbrev-ref HEAD)

PATH="$PATH:$GOPATH/bin"

export GO111MODULE=on
export GOARCH="amd64"
export GOOS="linux"
#export GOOS="darwin"
export CGO_ENABLED=1

export GOPROXY="https://proxy.golang.org"
export GOSUMDB=sum.golang.org
export GOPRIVATE="github.com"

# go mod tidy -v
go mod download -x
go build -v -ldflags="-s -w -X github.com/mdmoshiur/example-go/config.Version=${BUILD_DATE}:${BRANCH_NAME}:${VERSION}" -o example-go
