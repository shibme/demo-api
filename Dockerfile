FROM ghcr.io/shibme/cg/static
ARG TARGETARCH
COPY ./dist/*$TARGETARCH*/demo-api /demo-api
ENTRYPOINT ["/demo-api"]