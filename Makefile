.PHONY: generate clean

PROTO_FILES=$(shell find . -maxdepth 1 -name "*.proto")
GO_OUT_DIR=.

generate:
	protoc -I. \
		--go_out=$(GO_OUT_DIR) --go_opt=paths=source_relative \
		--go-grpc_out=$(GO_OUT_DIR) --go-grpc_opt=paths=source_relative \
		$(PROTO_FILES)

clean:
	find . -name "*.pb.go" -type f -delete 