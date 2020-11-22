FROM golang:1.14-alpine as build_base
RUN apk add bash ca-certificates git gcc g++ libc-dev

ARG WORKSPACE
WORKDIR /${WORKSPACE}

# Force the go compiler to use modules
ENV GO111MODULE=on
# Copy go mod and sum files
COPY go.mod .
COPY go.sum .
RUN go mod download

# This image builds the weavaite server
FROM build_base AS server_builder
# Here we copy the rest of the source code
COPY . .

RUN go build -o main ./server/http

EXPOSE 80

CMD ["./main"]