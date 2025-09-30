#!/usr/bin/env bash

set -eux

MODE=$1

IN_ARCHIVE=/in/distributive/ArchLinuxARM-rpi-armv7-latest.tar.gz
OUT_IMAGE=/out/firmware.img

# Download tarball
if [ ! -f ${IN_ARCHIVE} ]; then
    wget http://os.archlinuxarm.org/os/ArchLinuxARM-rpi-armv7-latest.tar.gz -O ${IN_ARCHIVE}
fi

# Create image
rm -f ${OUT_IMAGE}
dd if=/dev/zero of=${OUT_IMAGE} bs=1M count=4096

# Create partitions
parted -s ${OUT_IMAGE} mklabel msdos
parted -s ${OUT_IMAGE} mkpart primary fat32 2048s 512M
parted -s ${OUT_IMAGE} mkpart primary ext4 512M 100%
parted -s ${OUT_IMAGE} set 1 boot on

# Format partitions
LOOPBACK=$(losetup --partscan --find --show ${OUT_IMAGE})
mkfs.vfat ${LOOPBACK}p1
mkfs.ext4 ${LOOPBACK}p2

# Mount partitions
rm -rf root
mkdir -p root
mount ${LOOPBACK}p2 root
mkdir -p root/boot
mount ${LOOPBACK}p1 root/boot

# Extract tarball
rm -rf arch
mkdir arch
bsdtar -xpf ${IN_ARCHIVE} -C arch

# Copy tarball contents
cp -Rp arch/boot/* root/boot/.
rm -rf arch/root
cp -Rp arch/* root/.

# Copy rook script
cp -Rp /in/firmware root/firmware

# Copy rook
cp /usr/bin/rook root/usr/bin/rook

# Copy qemu
cp /usr/sbin/qemu-arm-static root/usr/sbin/qemu-arm-static

# Enable binary translation
mount binfmt_misc -t binfmt_misc /proc/sys/fs/binfmt_misc
echo ':arm:M::\x7fELF\x01\x01\x01\x00\x00\x00\x00\x00\x00\x00\x00\x00\x02\x00\x28\x00:\xff\xff\xff\xff\xff\xff\xff\x00\xff\xff\xff\xff\xff\xff\xff\xff\xfe\xff\xff\xff:/usr/sbin/qemu-arm-static:' > /proc/sys/fs/binfmt_misc/register | true

# Configure chroot
mount -t proc /proc root/proc
mount -o bind /dev root/dev
mount -o bind /dev/pts root/dev/pts
mount -o bind /sys root/sys
mount binfmt_misc -t binfmt_misc root/proc/sys/fs/binfmt_misc

# Bootstrap
chroot root rook apply local --values /firmware/values_${MODE}.yaml --scripts /firmware/script.sh

# Cleanup
rm root/usr/bin/rook
rm root/usr/sbin/qemu-arm-static
rm -rf root/firmware

# Sync
sync
killall gpg-agent # TODO: Move to other script?
sleep 10 # TODO: wait for lock

# Unmount
umount /proc/sys/fs/binfmt_misc
umount root/proc/sys/fs/binfmt_misc
umount -f root/proc
umount -f root/dev/pts
umount -f root/dev
umount -f root/sys
umount -f root/boot root
losetup -d ${LOOPBACK}
