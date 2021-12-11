OS = $(shell go env GOHOSTOS)

.PHONY: format
format:
	bin/format.sh

.PHONY: lint
lint:
	golangci-lint run ./...

.PHONY: tidy
tidy:
	env GO111MODULE=on go mod tidy

.PHONY: pretty
pretty: tidy format lint

.PHONY: test
test:
	go test -v -race ./...

.PHONY: mockgen
mockgen:
	bin/generate-mock.sh

.PHONY: vendor
vendor:
	GO111MODULE=on go mod vendor

.PHONY: cover
cover:
	go test -v -race ./... -coverprofile=coverage.out
	go tool cover -html=coverage.out -o coverage.html
	go tool cover -func coverage.out 

.PHONY: coverhtml
coverhtml:
	go test -v -race ./... -coverprofile=coverage.out
	go tool cover -html=coverage.out

.PHONY: compile.cli
compile.cli:
	GO111MODULE=on CGO_ENABLED=0 GOOS=$(OS) go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o sulong cmd/cli/main.go

.PHONY: compile.cron
compile.cron:
	GO111MODULE=on CGO_ENABLED=0 GOOS=$(OS) go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o sulong cmd/cron/main.go

.PHONY: docker-build
docker-build:
	docker build --no-cache -t sulong:latest .

.PHONY: docker-run
docker-run:
	docker run -p 8080:8080 --env-file .env sulong:latest

.PHONY: docker-down
docker-down:
	docker compose --env-file .env down