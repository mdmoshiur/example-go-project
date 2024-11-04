FROM golang:1-alpine AS builder

RUN apk add --no-cache --update git gcc g++ bash openssh tzdata
RUN apk add --update git

# Set go env
ENV GOPATH=/go
ENV GO111MODULE=on
ENV GOPRIVATE="github.com"
ENV CGO_ENABLED=1

# Take SSH private key from argument
# create ssh dir and known_hosts file
# copy ssh private key, add magic as known host and
# use ssh over https for cloning

ARG SSH_KEY

RUN mkdir /root/.ssh/
RUN touch /root/.ssh/known_hosts
RUN echo "$SSH_KEY" > /root/.ssh/id_rsa
RUN chmod 600 /root/.ssh/id_rsa

COPY . $GOPATH/src/github.com/mdmoshiur/example-go
WORKDIR $GOPATH/src/github.com/mdmoshiur/example-go

RUN chmod +x ./build.sh
RUN ./build.sh

RUN mv ./example-go /go/bin/

FROM alpine:latest
RUN apk add --no-cache --update ca-certificates openssl && apk add --no-cache tzdata
COPY --from=0 /go/bin/example-go /usr/local/bin/example-go

ENTRYPOINT ["example-go"]
