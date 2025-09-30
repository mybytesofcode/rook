FROM rook-x86_64:latest AS build
FROM archlinux:base-20250420.0.338771
WORKDIR /build
RUN pacman -Syy && \
    pacman -Sy --noconfirm \
    qemu-user-static \
    parted \
    e2fsprogs \
    dosfstools \
    wget
COPY --from=build /usr/bin/rook /usr/bin/rook
ENTRYPOINT [ "/in/disk/script.sh" ]
