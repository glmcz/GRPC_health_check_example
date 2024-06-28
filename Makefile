



run:
	go run ./cmd/*.go

generate-proto:
	docker run --rm --volume "$(shell pwd):/workspace" --workdir /workspace/ bufbuild/buf generate \
 	--path api/schema/probe/healt_check.proto
