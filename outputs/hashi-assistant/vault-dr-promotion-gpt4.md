To automate the promotion of a secondary DR (Disaster Recovery) Vault cluster, you will need to follow these steps using the Vault API:

Check the replication status: Verify the secondary DR Vault cluster's replication status to ensure it's in the correct state for promotion.

API Endpoint: /sys/replication/dr/status

Curl example:

```
curl -s --header "X-Vault-Token: <your_secondary_dr_token>" \
  --request GET \
  https://<your_secondary_dr_vault_address>:8200/v1/sys/replication/dr/status
```

Generate a DR operation token: Use the primary cluster's DR operation token to generate a new DR operation token on the secondary cluster.

API Endpoint: /sys/replication/dr/secondary/token

Curl example:

```
curl -s --header "X-Vault-Token: <your_primary_dr_token>" \
  --request POST \
  https://<your_primary_vault_address>:8200/v1/sys/replication/dr/secondary/token
```

Save the response, which includes the newly generated DR operation token. You'll use this token in the next step.

Promote the secondary DR Vault cluster: Use the DR operation token from the previous step to promote the secondary DR cluster to a primary cluster.

API Endpoint: /sys/replication/dr/secondary/promote

Curl example:

```
curl -s --header "X-Vault-Token: <your_new_dr_operation_token>" \
  --request POST \
  https://<your_secondary_dr_vault_address>:8200/v1/sys/replication/dr/secondary/promote
```

### Gotchas:

1. Ensure that you use the correct Vault tokens for each API request. Using an incorrect token may lead to unauthorized access or unexpected behavior.

2. Make sure to check the secondary DR Vault cluster's replication status before attempting to promote it. Promoting a cluster with an incorrect status can lead to issues.

3. The primary DR Vault cluster must be unsealed to generate a DR operation token. If the primary cluster is sealed or unavailable, you may need to manually unseal it or follow a different disaster recovery procedure.

4. After promoting the secondary DR Vault cluster, your clients and applications should switch to using the newly promoted primary cluster. Update any configuration or environment variables as needed.

5. In the case of promoting a secondary DR Vault cluster to a primary cluster, it's essential to have a well-documented and tested disaster recovery plan in place. This plan should include regular backups and proper monitoring and alerting mechanisms to ensure a smooth recovery process.

6. Be aware of the differences between HashiCorp's Enterprise and Open Source products. Some features, such as DR replication, are available only in the Enterprise version of Vault. Ensure that you're using the correct version of Vault that supports the desired features.
