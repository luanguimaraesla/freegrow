#! /bin/bash

set -euxo pipefail

PROJECT_HOME=$(pwd)
EMULATOR_HOME="$PROJECT_HOME/images"

KERNEL_FILE="kernel-qemu-4.19.50-buster"
DTB_FILE="versatile-pb-buster.dtb"
RPI_IMAGE_RELEASE="2020-02-13"
RPI_IMAGE="$RPI_IMAGE_RELEASE-raspbian-buster"

RPI_KERNEL="$EMULATOR_HOME/$KERNEL_FILE"
RPI_VERSATILE_DTB="$EMULATOR_HOME/$DTB_FILE"
RPI_FS="$EMULATOR_HOME/$RPI_IMAGE.img"

TMP_DIR="$EMULATOR_HOME/tmp"
LOCK_FILE="$EMULATOR_HOME/.prepared"


# Download QEMU compatible kernel to boot our system
if [[ ! -f $RPI_KERNEL ]]; then
  echo "Downloading RPI Kernel"
  curl  \
    -o $RPI_KERNEL \
    -L https://github.com/dhruvvyas90/qemu-rpi-kernel/raw/master/$KERNEL_FILE
fi

if [[ ! -f $RPI_VERSATILE_DTB ]]; then
  echo "Downloading versatile dtb file"
  curl  \
    -o $RPI_VERSATILE_DTB \
    -L https://github.com/dhruvvyas90/qemu-rpi-kernel/raw/master/$DTB_FILE
fi

# Download Raspbian image
if [[ ! -f $RPI_FS ]]; then
  echo "Downloading RPI filesystem"
  curl \
    -o "$EMULATOR_HOME/$RPI_IMAGE.zip" \
    -L http://downloads.raspberrypi.org/raspbian/images/raspbian-2020-02-14/$RPI_IMAGE.zip

  cd $EMULATOR_HOME
  unzip $RPI_IMAGE.zip
fi

if [[ ! -f $LOCK_FILE ]]; then
  qemu-img resize $RPI_FS +10G
  echo "Don't forget to resize the partition inside the VM"
  echo
  echo "(0)  run make emulator"
  echo "(1)  sudo fdisk -l  # check the sector where /dev/sda2 begins"
  echo "(2)  sudo fdisk /dev/sda"
  echo "(3)  in fdisk press d to delete a partiton"
  echo "(4)  select the second one"
  echo "(5)  press n to create a new partition"
  echo "(6)  accept default primary"
  echo "(7)  enter the start sector you got in (1) for /dev/sda2"
  echo "(8)  if doing mistake quit without saving change with q"
  echo "(9)  otherwize w will write the new table to disk"
  echo "(10) you'll see an error saying that the volume is busy, just reboot"
  echo "(11) run make emulator again"
  echo "(12) sudo resize2fs /dev/sda2"
  echo

    touch $LOCK_FILE
fi

echo "finished"
