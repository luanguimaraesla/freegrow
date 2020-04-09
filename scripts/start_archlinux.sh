#! /bin/bash

PROJECT_HOME=$(pwd)
RPI_KERNEL="$PROJECT_HOME/images/kernel-qemu-4.4.34-jessie"
RPI_FS="$PROJECT_HOME/images/2017-03-02-raspbian-jessie.img"
QEMU=$(which qemu-system-arm)

$QEMU \
    -kernel $RPI_KERNEL \
    -cpu arm1176 \
    -m 256 \
    -M versatilepb \
    -no-reboot \
    -serial stdio \
    -append "root=/dev/sda2 panic=1 rootfstype=ext4 rw" \
    -drive "file=$RPI_FS,index=0,media=disk,format=raw" \
    -device e1000,netdev=net0 \
    -netdev user,id=net0,hostfwd=tcp::5555-:22
