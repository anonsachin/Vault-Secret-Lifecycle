# Vault Renew

Looking at how to manage vault secrets and lifecycle them properly.

--------------------------------------------------------------------

## Dependencies

- [Docker](https://docs.docker.com/engine/install/)
- [Go](https://golang.org/doc/install)
- [Terraform](https://learn.hashicorp.com/tutorials/terraform/install-cli)

--------------------------------------------------------------------

## Steps

- Start the vault container. Command - `make vault-dev`
- Configure Vault.
- Genrate a renewable periodic token. Command - `make vault-token`
- Use the genrated token in the program. `export VAULT_TOKEN=... `
- Run the code. Commmand - `make run`

--------------------------------------------------------------------

## Tip

- the `Write` and `Read` functions require the input in the same form as the http request, so have a look at the **HTTP API** for what you want to manage.