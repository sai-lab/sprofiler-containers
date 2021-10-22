FROM docker.io/golang:1.17-buster as builder

LABEL org.opencontainers.image.source https://github.com/sai-lab/sprofiler-containers

WORKDIR /src

RUN apt-get update && apt-get install -y make

COPY go.mod go.mod
COPY go.sum go.sum
RUN go mod download

COPY . /src

RUN make build

FROM scratch as runner

COPY --from=builder /src/bin/hello-go /hello-go

ENTRYPOINT ["/hello-go"]
