development.setup:
	pacman -Sy qemu qemu-user-static-binfmt
	mount binfmt_misc -t binfmt_misc /proc/sys/fs/binfmt_misc
	docker run --privileged --rm tonistiigi/binfmt --install all

distributive.download:
	wget -O in/distributive/ArchLinuxARM-rpi-armv7-latest.tar.gz http://os.archlinuxarm.org/os/ArchLinuxARM-rpi-armv7-latest.tar.gz

image.build: version ?= latest
image.build:
	docker build --no-cache -t raspberry:${version} .

firmware.build: version ?= latest
firmware.build: mode ?= acs
firmware.build:
	docker run -it --privileged \
		-v /dev:/dev \
		-v $(shell pwd)/in/distributive:/in/distributive \
		-v $(shell pwd)/in/disk:/in/disk \
		-v $(shell pwd)/in/firmware:/in/firmware \
		-v $(shell pwd)/out:/out \
		raspberry:${version} \
			${mode}

firmware.adapter.write: disk ?= /dev/sda
firmware.adapter.write:
	dd if=out/firmware.img of=${disk} bs=16M status=progress iflag=direct oflag=direct
	parted ${disk} resizepart 2 100%
	e2fsck -f ${disk}2
	resize2fs ${disk}2

firmware.slot.write: disk ?= /dev/mmcblk0
firmware.slot.write:
	dd if=out/firmware.img of=${disk} bs=16M status=progress iflag=direct oflag=direct
	parted ${disk} resizepart 2 100%
	e2fsck -f ${disk}p2
	resize2fs ${disk}p2
