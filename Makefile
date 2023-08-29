.PHONY: update trace_dumper geth clean docker start_docker trace_data

VERSION=scroll-v4.3.55

update: ## Let's keep it and docker version in consistent.
	go get -u github.com/scroll-tech/go-ethereum@${VERSION}
	go mod tidy

trace_dumper: ## Builds the trace_dumper instance.
	mkdir -p ${PWD}/bin
	go build -o ${PWD}/bin/trace_dumper

clean: ## Delete generated artifacts
	rm -rf ${PWD}/bin/*  ${PWD}/tracedata/*

start_docker:
	docker run --rm -p 8545:8545 -p 8546:8546 trace-dumper/l2geth:${VERSION}

docker: ## Build integration-test image
	docker build --platform linux/x86_64 --no-cache -t trace-dumper/l2geth:latest ./docker/l2geth/.

trace_data:
	./bin/trace_dumper -dump erc20
	./bin/trace_dumper -dump native
	./bin/trace_dumper -dump nft
	./bin/trace_dumper -dump greeter
	./bin/trace_dumper -dump sushi
	./bin/trace_dumper -dump dao
	./bin/trace_dumper -dump uniswapv2
	./bin/trace_dumper -dump multi_uniswapv2

abigen_precompiles:
	abigen --sol contracts/ecc/ecc.sol --pkg ecc --out contracts/ecc/ecc.go
	abigen --sol contracts/sha256/sha256.sol --pkg sha256 --out contracts/sha256/sha256.go
