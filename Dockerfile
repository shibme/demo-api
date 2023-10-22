FROM alpine AS build-env
ARG TARGETARCH
WORKDIR /build
COPY ./*.zip .
RUN unzip $TARGETARCH.zip

FROM ghcr.io/shibme/cg/static
COPY --from=build-env /build/demo-api /app
ENTRYPOINT ["/app"]