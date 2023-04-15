To automate the promotion of a secondary DR Vault cluster using the Vault API, the following steps can be taken:

1. Update the configuration of the secondary DR Vault cluster to promote it to a primary cluster.

2. Migrate any data from the old primary cluster to the new one.

3. Reconfigure any applications that use the Vault API to point to the new primary cluster.

The relevant API endpoints for this process would be:

1. The /sys/storage/dr/secondary-promotion endpoint - this endpoint allows for the primary promotion of a secondary DR Vault cluster.

2. The /sys/replication/performance/primary/secondary/migrate endpoint - this endpoint allows for the migration of data from the old primary cluster to the new one.

3. The /sys/tools/discover endpoint - this endpoint returns the active Vault instance for a given path.

A successful API request to promote a secondary DR Vault cluster would look like:

```
curl \
    --header "X-Vault-Token: <root_token>" \
    --request PUT \
    --data @config.json \
    http://<new_primary_vault>:8200/v1/sys/storage/dr/secondary-promotion
```
where the `config.json` file contains the updated configuration for the Vault cluster being promoted to primary.

A successful API request to migrate data from the old primary cluster to the new one would look like:

```
curl \
    --header "X-Vault-Token: <root_token>" \
    --request POST \
    http://<old_primary_vault>:8200/v1/sys/replication/performance/primary/<new_primary_vault>/migrate
```
where `<old_primary_vault>` and `<new_primary_vault>` are the addresses of the old and new primary Vault clusters, respectively.

A successful API request to discover the active Vault instance for a given path would look like:

```
curl \
    --header "X-Vault-Token: <client_token>" \
    --request GET \
    http://<vault_address>:8200/v1/sys/tools/discover?format=json&path=<secret_path>
```
where `<client_token>` is the client token for the application using the Vault API, `<vault_address>` is the address of the Vault cluster, and `<secret_path>` is the path to the secret being accessed.

Potential gotchas to watch out for include ensuring that the configuration of the promoted Vault cluster is correct, ensuring that all required data is migrated from the old primary cluster to the new one, and ensuring that all applications using the Vault API are reconfigured to point to the new primary cluster. Additionally, if using Hashicorp Enterprise version of Vault, there may be additional steps or considerations for promotion and migration.