FROM golang:1.17-alpine as build

RUN apk add --no-cache git

WORKDIR /go/src/app
ADD . /go/src/app

RUN go get -d -v ./...

ENV CGO_ENABLED=0

RUN go build -trimpath \
    -a -ldflags "-w -X main.gitCommit=$(git rev-parse HEAD) -X main.gitVersion=$(git describe --tags --always --dirty)" \
    -o /go/bin/app

FROM gcr.io/distroless/base
COPY --from=build /go/bin/app /
CMD ["/app"]

