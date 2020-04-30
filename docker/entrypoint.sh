#! /bin/bash

SDCARD="/sdcard"

IPADDR=$(hostname -I)
IMAGE=${RPI_IMAGE:-"$SDCARD/2020-02-13-raspbian-buster.img"}
KERNEL=${RPI_KERNEL:-"$SDCARD/kernel-qemu-4.19.50-buster"}
DTB_FILE=${RPI_VERSATILE_DTB:-"$SDCARD/versatile-pb-buster.dtb"}
INIT_SSH=${RPI_INIT_SSH:-"no"}

QEMU="/usr/local/bin/qemu-system-arm"

# setup networking
/networking.sh

if [[ $RPI_INIT_SSH != "yes" ]]; then
  echo "running directly for IP $IPADDR"
  $QEMU \
    -M versatilepb \
    -cpu arm1176 \
    -m 256M \
    --net nic \
    --net user,hostfwd=::2222-:22 \
    --net tap,ifname=tap0 \
    --dtb $DTB_FILE \
    --kernel $KERNEL \
    --append "root=/dev/sda2 panic=1 rootfstype=ext4 rw" \
    --drive "file=$IMAGE,index=0,media=disk,format=raw" \
    --no-reboot \
    --display none \
    --serial mon:stdio
else
  echo "running for enabling ssh"
  expect /enable_ssh.exp $IMAGE $KERNEL $DTB_FILE
fi
