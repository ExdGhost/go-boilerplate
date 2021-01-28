FROM golang:alpine  AS builder

RUN apk update
RUN apk add git

# Specifies the working dir in the new container
WORKDIR /go/src/go-boilerplate

ADD ./ .

# Runs the command in WORKDIR in the image
COPY /cmd/api/main.go .
RUN  commitId="$(git rev-parse --short HEAD)" && go build -mod=vendor -ldflags "-X main.Version=$commitId" -o "go-boilerplate"

FROM golang:alpine
WORKDIR /bin/go-boilerplate

COPY --from=builder /go/src/go-boilerplate/go-boilerplate .
RUN mkdir config
COPY --from=builder /go/src/go-boilerplate/config/tier ./config/tier

CMD ["./go-boilerplate"]
EXPOSE 80 5001
