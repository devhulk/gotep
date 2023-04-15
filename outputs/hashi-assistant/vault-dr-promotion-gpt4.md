To automate the promotion of a secondary DR vault cluster using the Vault API, you can follow these steps:

1. Generate a batch DR operation token on the primary cluster (cluster A).
2. Promote the secondary DR cluster (cluster B) to become the new primary.
3. Demote the original primary cluster (cluster A) to become secondary.
4. Update the original primary cluster (cluster A) to point to the new primary cluster (cluster B).

Here are the relevant API endpoints and curl examples:

### Generate a batch DR operation token on the primary cluster (cluster A):

Endpoint:
POST /sys/replication/dr/primary/secondary-token
Curl example:

```
curl --request POST --header "X-Vault-Token: <PRIMARY_CLUSTER_VAULT_TOKEN>" --data '{"id": "cluster-b"}' http://primary-cluster:8200/v1/sys/replication/dr/primary/secondary-token
```
### Promote the secondary DR cluster (cluster B) to become the new primary:

Endpoint:
POST /sys/replication/dr/secondary/promote
Curl example:
```
curl --request POST --data '{"dr_operation_token": "<DR_OPERATION_TOKEN>"}' http://secondary-cluster:8200/v1/sys/replication/dr/secondary/promote
```
### Demote the original primary cluster (cluster A) to become secondary:
Endpoint:

POST /sys/replication/dr/primary/demote
Curl example:

```
curl --request POST --header "X-Vault-Token: <PRIMARY_CLUSTER_VAULT_TOKEN>" http://primary-cluster:8200/v1/sys/replication/dr/primary/demote
```
### Update the original primary cluster (cluster A) to point to the new primary cluster (cluster B):
Endpoint:

POST /sys/replication/dr/secondary/update-primary
Curl example:
```
curl --request POST --header "X-Vault-Token: <PRIMARY_CLUSTER_VAULT_TOKEN>" --data '{"primary_api_addr": "http://new-primary-cluster:8200", "dr_operation_token": "<DR_OPERATION_TOKEN>"}' http://primary-cluster:8200/v1/sys/replication/dr/secondary/update-primary
```
### Potential gotchas:

1. Ensure that you have a recent backup of the Vault data before performing failover or failback operations.
2. Make sure to use the DR operation token when promoting and demoting clusters.
3. Be cautious about the order of operations, as promoting the new primary before demoting the old primary may be necessary in some production scenarios due to unavailability of the old primary.
4. Always keep the DR operation token secure and use it judiciously.
