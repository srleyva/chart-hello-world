# Build Binary
FROM golang:alpine as build
RUN apk update && apk add make git tree
RUN mkdir -p $GOPATH/src/github.com/srleyva/chart-hello-world
WORKDIR $GOPATH/src/github.com/srleyva/chart-hello-world
ADD ./ ./
RUN CGO_ENABLED=0 GOOS=linux go build -o hello main.go && cp hello /hello

FROM scratch
COPY --from=build /hello /bin/hello
ENTRYPOINT ["hello"]
