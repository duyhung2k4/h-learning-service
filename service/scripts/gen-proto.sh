find proto -name "*.proto" | xargs protoc \
	--go_opt=paths=source_relative \
	--go-grpc_opt=paths=source_relative \
	--go_out=generated/grpc \
	--go-grpc_out=generated/grpc \
    --proto_path=proto
