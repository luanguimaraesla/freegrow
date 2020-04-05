#! /bin/bash

PROJECT_HOME="$GOPATH/src/github.com/luanguimaraesla/freegrow"
RPI_KERNEL="$PROJECT_HOME/images/kernel-qemu-4.4.34-jessie"
RPI_FS="$PROJECT_HOME/images/2017-03-02-raspbian-jessie.img"
TMP_DIR="$PROJECT_HOME/images/tmp"

# Download QEMU
yay -S qemu qemu-arch-extra bridge-utils unzip

# Add image path
mkdir -p $PROJECT_HOME/images

# Download QEMU compatible kernel to boot our system
if [[ ! -f $RPI_KERNEL ]]; then
  curl  \
    -o $RPI_KERNEL \
    -L https://github.com/dhruvvyas90/qemu-rpi-kernel/raw/master/kernel-qemu-4.4.34-jessie
fi

# Download Raspbian image
if [[ ! -f $RPI_FS ]]; then
  curl \
    -o "$PROJECT_HOME/images/2017-03-02-raspbian-jessie.zip" \
    -L http://downloads.raspberrypi.org/raspbian/images/raspbian-2017-03-03/2017-03-02-raspbian-jessie.zip

  cd "$PROJECT_HOME/images/"
  unzip 2017-03-02-raspbian-jessie.zip
fi

if [[ ! -f $PROJECT_HOME/images/.prepared ]]; then
  # prepare the image
  SECTOR1=$( sudo fdisk -l $RPI_FS | grep FAT32 | awk '{ print $2 }' )
  SECTOR2=$( sudo fdisk -l $RPI_FS | grep Linux | awk '{ print $2 }' )
  OFFSET1=$(( SECTOR1 * 512 ))
  OFFSET2=$(( SECTOR2 * 512 ))

  mkdir -p $TMP_DIR
  sudo mount $RPI_FS -o offset=$OFFSET1 $TMP_DIR
  sudo touch $TMP_DIR/ssh   # this enables ssh
  sudo umount $TMP_DIR

  sudo mount $RPI_FS -o offset=$OFFSET2 $TMP_DIR
cat << EOT | sudo tee -a $TMP_DIR/etc/udev/rules.d/90-qemu.rules
KERNEL=="sda", SYMLINK+="mmcblk0"
KERNEL=="sda?", SYMLINK+="mmcblk0p%n"
KERNEL=="sda2", SYMLINK+="root"
EOT

  sudo umount -l $TMP_DIR
  rm -rf $TMP_DIR

  qemu-img resize $RPI_FS +10G
  echo "Don't forget to resize the partition inside VM with resize2fs"
  touch $PROJECT_HOME/images/.prepared
fi
