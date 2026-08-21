# Uses pre-built binaries from CI - no build stage needed
# Binary is selected based on TARGETARCH (amd64 or arm64)
FROM gcr.io/distroless/static-debian12:nonroot@sha256:afa5c872c891853ca7fcf1f12c3edb23f7eeef36189728842dd51042ff57f7ab

ARG TARGETARCH
COPY dist/readability_linux_${TARGETARCH} /usr/local/bin/readability

ENTRYPOINT ["readability"]
