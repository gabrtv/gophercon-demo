IMAGE=gabrtv/logarchiver:canary

all: docker-build

build-linux:
	mkdir -p rootfs/opt/logarchiver/bin
	CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"' -o rootfs/opt/logarchiver/bin/logarchiver

docker-build: build-linux
	docker build -t ${IMAGE} rootfs

