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

build: build-monolith build-services

rebuild: clean-monolith clean-services build

clean-monolith:
	docker image rm mallbots-monolith

clean-services:
	docker image rm mallbots-baskets mallbots-cosec mallbots-customers mallbots-depot mallbots-notifications mallbots-ordering mallbots-payments mallbots-search mallbots-stores

build-monolith:
	docker build -t mallbots-monolith --file docker/Dockerfile .

build-services:
	docker build -t mallbots-baskets --file docker/Dockerfile.microservices --build-arg=service=baskets .
	docker build -t mallbots-cosec --file docker/Dockerfile.microservices --build-arg=service=cosec .
	docker build -t mallbots-customers --file docker/Dockerfile.microservices --build-arg=service=customers .
	docker build -t mallbots-depot --file docker/Dockerfile.microservices --build-arg=service=depot .
	docker build -t mallbots-notifications --file docker/Dockerfile.microservices --build-arg=service=notifications .
	docker build -t mallbots-ordering --file docker/Dockerfile.microservices --build-arg=service=ordering .
	docker build -t mallbots-payments --file docker/Dockerfile.microservices --build-arg=service=payments .
	docker build -t mallbots-search --file docker/Dockerfile.microservices --build-arg=service=search .
	docker build -t mallbots-stores --file docker/Dockerfile.microservices --build-arg=service=stores .

test-e2e:
	@go1.18 test -v -tags=e2e ./testing/e2e/...

up-monolith:
	docker compose --profile mallbots-monolith up -d
