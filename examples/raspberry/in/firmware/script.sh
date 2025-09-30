#!/usr/bin/env bash

set -eux

# Hostname
render /etc/hostname

# DNS
unlink /etc/resolv.conf
render /etc/resolv.conf

# Vconsole
render /etc/vconsole.conf

# Pacman mirror
render /etc/pacman.d/mirrorlist

# Pacman packages
pacman-key --init
pacman-key --populate archlinuxarm
pacman -Syyu --noconfirm

{{ range .packages.libraries }}
    pacman -Syu --noconfirm {{ . }}
{{ end }}

{{ range .packages.tools }}
    pacman -Syu --noconfirm {{ . }}
{{ end }}

{{ range .packages.network }}
    pacman -Syu --noconfirm {{ . }}
{{ end }}

{{ range .packages.services }}
    pacman -Syu --noconfirm {{ . }}
{{ end }}

{{ range .packages.console }}
    pacman -Syu --noconfirm {{ . }}
{{ end }}

{{ if .desktop.enable }}
    {{ range .packages.desktop.components }}
       pacman -Syu --noconfirm {{ . }}
    {{ end }}

    {{ range .packages.desktop.utils }}
        pacman -Syu --noconfirm {{ . }}
    {{ end }}
{{ end }}

# Configure user
usermod -l user alarm
groupmod -n user alarm
usermod -d /home/user -m user
usermod -a -G docker user
usermod -a -G seat user
render /etc/sudoers.d/user

# Locale
render /etc/locale.conf
render /etc/locale.gen
locale-gen

# SSH server
render /etc/ssh/sshd_config
systemctl enable sshd

# SSH keys
mkdir -p /home/user/.ssh
chown -R user:user /home/user/.ssh
chmod 0700 /home/user/.ssh
touch /home/user/.ssh/authorized_keys
chown -R user:user /home/user/.ssh/authorized_keys
chmod 0600 /home/user/.ssh/authorized_keys
render /home/user/.ssh/authorized_keys

# Enable services
systemctl enable ntpdate
systemctl enable docker

# Iptables
render /etc/iptables/iptables.rules
systemctl enable iptables

# Dhcpcd
# systemctl enable dhcpcd

# Sysctl
render /etc/sysctl.d/0-base.conf

{{ if .ethernet.enable }}
    # Ethernet
    render /etc/systemd/system/macspoof@.service
    systemctl enable macspoof@end0
{{ end }}

{{ if .wifi.client }}
    # WPA supplicant
    render /etc/wpa_supplicant/wpa_supplicant-wlan0.conf
    systemctl enable wpa_supplicant@wlan0

    # Hook
    ln -s /usr/share/dhcpcd/hooks/10-wpa_supplicant /usr/lib/dhcpcd/dhcpcd-hooks
{{ end }}

# Seatd
systemctl enable seatd

{{ if .desktop.enable }}
    # Ly
    render /etc/ly/config.ini
    systemctl enable ly.service

    # Sway
    render /etc/sway/config

    # Foot
    render /etc/xdg/foot/foot.ini
{{ end }}

# Boot
render /boot/config.txt
echo 'root=/dev/mmcblk0p2 rw rootwait console=serial0,115200 console=tty1 fsck.repair=yes quiet splash' > /boot/cmdline.txt
