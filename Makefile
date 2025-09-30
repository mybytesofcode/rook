development-setup:
	docker run --privileged --rm tonistiigi/binfmt --install all

source-test:
	go test -v ./...

binary-build:
	go build -o rook ./cmd/rook
	strip rook

image-build-x86_64: version ?= latest
image-build-x86_64:
	docker buildx build \
		--platform linux/amd64 \
		--file Dockerfile \
		--tag rook-x86_64:${version} \
		.

image-build-armv7l: version ?= latest
image-build-armv7l:
	docker buildx build \
		--platform linux/arm/v7 \
		--file Dockerfile \
		--tag rook-armv7l:${version} \
		.

package-install: version ?= latest
package-install:
	docker run -it -v /usr/bin:/out rook-x86_64:${version}
