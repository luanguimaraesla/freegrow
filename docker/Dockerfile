FROM ubuntu:20.04 AS qemu-system-arm-builder

ARG QEMU_VERSION=4.2.0
ENV QEMU_TARBALL="qemu-${QEMU_VERSION}.tar.xz" \
    DEBIAN_FRONTEND=noninteractive

WORKDIR /qemu

RUN apt update && \
    apt install -y \
    wget \
    gpg \
    pkg-config \
    python \
    build-essential \
    libglib2.0-dev \
    libpixman-1-dev \
    libfdt-dev \
    zlib1g-dev \
    flex \
    bison

RUN wget "https://download.qemu.org/${QEMU_TARBALL}" && \
    wget "https://download.qemu.org/${QEMU_TARBALL}.sig" && \
    gpg --keyserver keyserver.ubuntu.com --recv-keys CEACC9E15534EBABB82D3FA03353C9CEF108B584 && \
    gpg --verify "${QEMU_TARBALL}.sig" "${QEMU_TARBALL}"

RUN tar xvf "${QEMU_TARBALL}" && \
    "qemu-${QEMU_VERSION}/configure" --static --target-list=arm-softmmu && \
    make -j$(nproc) && \
    strip "arm-softmmu/qemu-system-arm"


FROM ubuntu:20.04 as rpi

COPY --from=qemu-system-arm-builder /qemu/arm-softmmu/qemu-system-arm /usr/local/bin/qemu-system-arm

ENV DEBIAN_FRONTEND=noninteractive
RUN apt update && \
    apt install -y \
    iproute2 \
    bridge-utils \
    kmod \
    expect && \
    rm -rf /var/lib/apt/lists/*

VOLUME /sdcard

ADD enable_ssh.exp /enable_ssh.exp
ADD networking.sh /networking.sh
ADD entrypoint.sh /entrypoint.sh

ENTRYPOINT /entrypoint.sh
