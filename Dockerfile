FROM golang AS build-env
WORKDIR /build
COPY . /build
RUN go build -a -tags 'osusergo netgo static_build' -ldflags '-w -extldflags "-static"' -o app

FROM ubuntu:22.04
COPY --from=build-env /build/app /ws/
WORKDIR /ws
ENTRYPOINT ["/ws/app"]