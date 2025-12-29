install-tools:
	@echo installing tools
	@go1.18 install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@v2.11.3
	@go1.18 install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@v2.11.3
	@go1.18 install google.golang.org/protobuf/cmd/protoc-gen-go@v1.34.0
	@go1.18 install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2.0
	@go1.18 install github.com/bufbuild/buf/cmd/buf@v1.3.0
	@go1.18 install github.com/vektra/mockery/v2@v2.14.0
	@go1.18 install github.com/go-swagger/go-swagger/cmd/swagger@v0.29
	@go1.18 install github.com/cucumber/godog/cmd/godog@v0.15.1
	@echo done

generate:
	@echo running code generation
	@go1.18 generate ./...
	@echo done

test-e2e:
	@go1.18 test -v -tags=e2e ./testing/e2e/...
