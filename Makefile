build:
	GOROOT=${PWD} go install ./cmd/renew

run:
	go run ./cmd/renew/main.go