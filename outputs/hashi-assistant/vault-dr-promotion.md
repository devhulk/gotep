Automating the promotion of a secondary Disaster Recovery (DR) Vault cluster using the Vault API involves several steps. The relevant API endpoints include:

1. Enable a DR replication performance standby node. This can be done using the API endpoint:
/v1/sys/replication/dr/secondary/enable

2. Initialize the DR replication process for the standby node. You can use the API endpoint:
/v1/sys/replication/dr/secondary/init

3. Promote the DR replication performance standby to the primary cluster. Here, you will use the API endpoint:
/v1/sys/replication/dr/promote

Below are examples of successful API requests using curl:

1. Enable DR replication performance standby node
curl \
    --header "X-Vault-Token: ${VAULT_TOKEN}" \
    --request PUT \
    ${VAULT_ADDR}/v1/sys/replication/dr/secondary/enable

2. Initialize the DR replication process for the standby node
curl \
    --header "X-Vault-Token: ${VAULT_TOKEN}" \
    --request POST \
    ${VAULT_ADDR}/v1/sys/replication/dr/secondary/init

3. Promote the DR replication performance standby to primary cluster. 
curl \
    --header "X-Vault-Token: ${VAULT_TOKEN}" \
    --request POST \
    ${VAULT_ADDR}/v1/sys/replication/dr/promote

Gotchas to watch out for when automating the promotion of a secondary DR Vault cluster using the Vault API include ensuring that the DR node is healthy and connected to both the primary and secondary Vault clusters. Additionally, you will need to provide the correct Vault token with permissions to perform the necessary actions, and to supply the correct endpoint URLs to the command line arguments. Lastly, be sure to review the API response messages and check the Vault logs to confirm that the operations were successful.