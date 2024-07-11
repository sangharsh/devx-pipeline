#!/bin/zsh
VAULT_ADDR='http://127.0.0.1:8200'

if [[ -z "${VAULT_TOKEN}" ]]; then
    echo "Error: VAULT_TOKEN is not set." >&2
    exit 1
fi

vault secrets enable -path=concourse -version=2 kv

vault policy write concourse ./concourse-policy.hcl
vault token create --policy concourse --period 1d

# SECURITY alert: Puts your SSH private key on vault
echo -n "Put your SSH private key on Vault? (y/N)"
if [[ $answer == "y" || $answer == "Y" ]]; then
    echo "Proceeding..."
    cat ~/.ssh/id_ed25519 | vault kv put -mount=concourse shared ssh_key=-
else
    echo "Skipping..."
fi