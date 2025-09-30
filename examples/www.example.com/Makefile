connect: host ?= 127.0.0.1
connect: port ?= 22
connect: user ?= root
connect:
	TERM=xterm \
	ssh -v -p ${port} ${user}@${host}

configure: host ?= 127.0.0.1
configure: port ?= 22
configure: user ?= root
configure:
	rook apply remote --ssh.host ${host} --ssh.port ${port} --ssh.user ${user} --values values.yaml --scripts script.configure.sh

reboot: host ?= 127.0.0.1
reboot: port ?= 22
reboot: user ?= root
reboot:
	rook apply remote --ssh.host ${host} --ssh.port ${port} --ssh.user ${user} --values values.yaml --scripts script.reboot.sh
