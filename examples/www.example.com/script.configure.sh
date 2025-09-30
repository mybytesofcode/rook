#!/usr/bin/env bash

set -eux

export DEBIAN_FRONTEND=noninteractive

apt update -y
apt upgrade -y
apt install -y \
    netfilter-persistent \
    iptables-persistent \
    knockd \
    docker.io \
    docker-compose \
    apache2-utils

# User user
useradd user | true
chsh user -s /usr/bin/bash
mkdir -p /home/user/.ssh
touch /home/user/.ssh/authorized_keys
chown -R user:user /home/user
chmod 0700 /home/user/.ssh
chmod 0600 /home/user/.ssh/authorized_keys
usermod -a -G docker user
render /home/user/.ssh/authorized_keys
render /etc/sudoers.d/user

# Iptables
systemctl enable netfilter-persistent
systemctl start netfilter-persistent
render /etc/iptables/rules.v4

# Workaround
iptables -A INPUT -i eth0 -p tcp --dport {{ .openssh.port }} -j ACCEPT

# Docker
systemctl enable docker
systemctl start docker
docker login \
    --username {{ .registry.user }} \
    --password {{ .registry.password }} \
    registry.example.com

# Certbot
if [ ! -e "/etc/letsencrypt/live/www.example.com/privkey.pem" ]; then
    docker run \
        -p 80:80 \
        --rm \
        --name certbot \
        -v /etc/letsencrypt:/etc/letsencrypt \
        -v /var/lib/letsencrypt:/var/lib/letsencrypt \
        certbot/certbot \
        certonly \
            --noninteractive \
            --agree-tos \
            -m queue@example.com \
            -d www.example.com \
            --standalone

    docker run \
        -p 80:80 \
        --rm \
        --name certbot \
        -v /etc/letsencrypt:/etc/letsencrypt \
        -v /var/lib/letsencrypt:/var/lib/letsencrypt \
        certbot/certbot \
        certonly \
            --noninteractive \
            --agree-tos \
            -m queue@example.com \
            -d example.com \
            --standalone
fi

# Nginx
mkdir -p /etc/nginx
render /etc/nginx/nginx.conf
docker run \
    --name nginx \
    --restart always \
    --network host \
    -d \
    -v /etc/letsencrypt:/etc/letsencrypt \
    -v /etc/nginx/nginx.conf:/etc/nginx/nginx.conf \
        nginx:{{ .nginx.version }} | true

# OpenSSH
render /etc/ssh/sshd_config
systemctl enable ssh
systemctl start ssh

# Knockd
render /etc/knockd.conf
systemctl enable knockd
systemctl start knockd

# Landing
docker container rm -f landing | true
docker run \
    -p 127.0.0.1:8080:8080 \
    -d \
    --restart always \
    --name landing \
    -v /etc/letsencrypt:/etc/letsencrypt \
    registry.example.com/landing-x86_64:{{ .landing.version }}

# Workaround
iptables -D INPUT -i eth0 -p tcp --dport {{ .openssh.port }} -j ACCEPT
