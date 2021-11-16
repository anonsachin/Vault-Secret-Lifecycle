# Vault Secret Lifecycle

Example on how to manage vault secrets and lifecycle them properly.

--------------------------------------------------------------------

## Dependencies

- [Docker](https://docs.docker.com/engine/install/)
- [Go](https://golang.org/doc/install)
- [Terraform](https://learn.hashicorp.com/tutorials/terraform/install-cli)

--------------------------------------------------------------------

## Steps

- Start the vault container. Command - `make vault-dev`
- Configure Vault. Command - `make init` & `make apply`
- Genrate a renewable periodic token. Command - `make vault-token`
- Use the genrated token in the program. `export VAULT_TOKEN=... `
- Run the code. Commmand - `make run`
- `Ctrl + c` twice to stop the renewal.
- Clean up the vault env. Commad - `make destroy` & `make vault-dev-down`

--------------------------------------------------------------------

## Resources
- Vault token API Docs
    - [Token renew](https://www.vaultproject.io/api/auth/token#renew-a-token-self)
    - [Generate certs](https://www.vaultproject.io/api-docs/secret/pki#generate-certificate)
    - [PKI engine](https://learn.hashicorp.com/tutorials/vault/pki-engine)
    - [Tokens](https://www.vaultproject.io/docs/concepts/tokens)

--------------------------------------------------------------------

## Tip

- the `Write` and `Read` functions require the input in the same form as the http request, so have a look at the **HTTP API** for what you want to manage.

- certs cannot be renewed, they need to reissued.