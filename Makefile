BUILD_BIN=main
LDFLAGS=

.PHONY: build build-prod

build:
	GOOS=linux go build -tags dev -ldflags="-s -w $(LDFLAGS)" -o $(BUILD_BIN) main.go

build-local:
	go build -tags prod -ldflags="-s -w $(LDFLAGS)" -o $(BUILD_BIN) main.go

build-prod:
	GOOS=linux go build -tags prod -ldflags="-s -w $(LDFLAGS)" -o $(BUILD_BIN) main.go
