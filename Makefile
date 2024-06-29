.PHONY: build lint test

run:
	go run ./cmd/*.go

generate-proto:
	docker run --rm --volume "$(shell pwd):/workspace" --workdir /workspace/ bufbuild/buf generate \
 	--path api/schema/probe/healt_check.proto

build:
	CGO_ENABLED=0 \
		go build \
			-installsuffix cgo \
			-o ./build/app \
			-ldflags "-X main.Version=$(APP_VERSION)" \
			./cmd/*.go

build-linux:
	CGO_ENABLED=0 GOOS=linux \
		go build \
		-installsuffix cgo \
		-o ./build/app \
		-ldflags "-X main.Version=$(APP_VERSION)" \
		./cmd/*.go