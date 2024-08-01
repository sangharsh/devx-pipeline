# Vault for secret management

## Install Vault CLI

https://developer.hashicorp.com/vault/tutorials/getting-started/getting-started-install

## Run

```
docker run --cap-add=IPC_LOCK -d --name=dev-vault -p 8200:8200 hashicorp/vault:1.17
```

Currently running in development mode. Not suitable for production use.

## Set environment variable

```
export VAULT_ADDR='http://127.0.0.1:8200'

# Replace with root token from vault logs (container)
export VAULT_TOKEN=ROOT_TOKEN
```

## Setup

```
# Enable a KV v2 secret store for concourse
vault secrets enable -path=concourse -version=2 kv

# Apply policy
vault policy write concourse ./concourse-policy.hcl

# Create a temporary token to use in concourse pipeline
vault token create --policy concourse --period 1d

```
Copy token to apps/inventory/build.yml at `var_sources.name['vault].config.client_token`

### Upload your SSH private key
Required for fetching git repo in concourse job

```
cat ~/.ssh/id_ed25519 | vault kv put -mount=concourse shared ssh_key=-
```
