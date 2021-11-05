build:
	GOROOT=${PWD} go install ./cmd/renew

run:
	VAULT_ADDR=http://127.0.0.1:8200 VAULT_TOKEN=myroot go run ./cmd/renew/main.go

vault-dev:
	vault server -dev -dev-root-token-id="myroot"