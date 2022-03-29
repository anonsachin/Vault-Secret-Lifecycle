# BASE
#-----
build:
	GOROOT=${PWD} go install ./cmd/renew

run:
	VAULT_ADDR=http://127.0.0.1:8200 go run ./cmd/renew/main.go

fmt:
	terraform fmt -recursive
	go fmt ./...

docker-lint:
	docker run --rm -v $(shell pwd):/app -w /app golangci/golangci-lint:v1.33.0 golangci-lint run -v --enable gosec

# VAULT
#------
vault-dev:
	#vault server -dev -dev-root-token-id="myroot"
	(cd vault; docker-compose up -d)

vault-dev-down:
	(cd vault; docker-compose down -v)

vault-token:
	VAULT_ADDR=http://127.0.0.1:8200 VAULT_TOKEN=myroot vault token create -policy=root -orphan -period=1m -renewable


# Terraform
#----------
init:
	(cd vault-setup; terraform init)

plan:
	(cd vault-setup; terraform plan)

apply:
	(cd vault-setup; terraform apply)

destroy:
	(cd vault-setup; terraform destroy)

output:
	(cd vault-setup; terraform output)