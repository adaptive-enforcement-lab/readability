# Uses pre-built binaries from CI - no build stage needed
# Binary is selected based on TARGETARCH (amd64 or arm64)
FROM gcr.io/distroless/static-debian12:nonroot@sha256:a9329520abc449e3b14d5bc3a6ffae065bdde0f02667fa10880c49b35c109fd1

ARG TARGETARCH
COPY dist/readability_linux_${TARGETARCH} /usr/local/bin/readability

ENTRYPOINT ["readability"]
