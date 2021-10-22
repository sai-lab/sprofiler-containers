FROM docker.io/library/gcc:11.2.0 as builder

LABEL org.opencontainers.image.source https://github.com/sai-lab/sprofiler-containers

WORKDIR /src
COPY . .
RUN gcc main.c -o hello

FROM gcr.io/distroless/cc as ruuner

COPY --from=builder /src/hello /hello

ENTRYPOINT ["/hello"]
