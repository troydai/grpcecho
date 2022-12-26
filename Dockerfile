FROM golang:alpine AS builder

RUN apk update && apk add --no-cache make protobuf-dev
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

WORKDIR /src
COPY go.mod /src
COPY go.sum /src
RUN go mod download

COPY . /src
RUN make bin

FROM alpine

WORKDIR /opt/bin

COPY --from=builder /src/artifacts/server /opt/bin/server

EXPOSE 8080

ENTRYPOINT [ "/opt/bin/server" ]
