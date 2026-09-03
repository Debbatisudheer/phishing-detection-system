.PHONY: all fmt vet lint build backend-test frontend-build ui-test ci

fmt:
	gofmt -w .

vet:
	go vet ./...

build:
	go build ./...

backend-test:
	go test ./... -v

frontend-build:
	cd frontend && npm install --include=dev && npm run build

ci: fmt vet build frontend-build backend-test

all: ci