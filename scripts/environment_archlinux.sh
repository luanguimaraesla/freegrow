#! /bin/bash

PROJECT_HOME=$(pwd)
RPI_KERNEL="$PROJECT_HOME/images/kernel-qemu-4.4.34-jessie"
RPI_FS="$PROJECT_HOME/images/2017-03-02-raspbian-jessie.img"
TMP_DIR="$PROJECT_HOME/images/tmp"

echo $PROJECT_HOME
exit 0
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

  touch $PROJECT_HOME/images/.prepared
fi

echo "finished"
