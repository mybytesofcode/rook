FROM golang:1.24.2-alpine3.21 AS build
RUN apk add -U make binutils
WORKDIR /build
RUN --mount=type=bind,source=.,destination=/build,rw \
    make binary.build && \
    mv rook /usr/bin/rook

FROM alpine:3.21.3
COPY --from=build /usr/bin/rook /usr/bin/rook
ENTRYPOINT [ "cp", "/usr/bin/rook", "/out/rook" ]
