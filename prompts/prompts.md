# hashi-assistant prompts

# Introduction

The prompts intent is for a Vault User to simply ask.
```
How do I automate the promotion of my DR cluster?
```
We may want it to be a little manual but I just wanted to see what ChatGPT would do since its a pretty niche
technical topic.

# Comparison

Massive differences in accuracy observed between the GPT-3.5 and GPT-4 model. 
The character limit may play a role here because with GPT4 I was able to submit
a prompt with 10K plus characters (the model supports 30K plus). That meant I was able to add most of the updated documentation to the prompt so GPT4 didn't have to guess anywhere near as much. 

ChatGPT3.5 completely missed the mark and completely fabricated a ton of nonsense. For example it completely made up API endpoints that don't exist. 

ChatGPT4 had much less obvious errors. 

Its worth noting that even with the ChatGPT3.5 prompt above the results were "more" accurate then the ChatGPT3.5 response. GPT4 just gave a lot more room for additional context.

I could still add around another 20K characters worth of content to my prompt with GPT4 which would could potentially provide a much higher degree of accuracy.

# Prompts and Responses

## GPT 3.5 Prompt

```
## Context
You are an assistant that is an expert in all Hashicorp tooling.
You specialize in understanding each potential api call for Terraform cloud, Vault, Consul, and Nomad.
Your background is in Cloud Architecture and infrastructure automation.
You can give expert level advice on SRE, Devops, and Platform Engineering.
You understand the difference between Hashicorps Enterprise and Open Source versions of their products and can spek to the differences.

## Question
How would I automate the promotion of a secendary DR vault cluster using the Vault API?
List the relevant API endpoints.
Give curl examples of what successful API requests would look like.
Highlight any potential gotchas.

## Answer

```

## GPT 3.5 Response

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



## GPT 4 Prompt
```
## Context
You are an assistant that is an expert in all Hashicorp tooling.
You specialize in understanding each potential api call for Terraform cloud, Vault, Consul, and Nomad.
Your background is in Cloud Architecture and infrastructure automation.
You can give expert level advice on SRE, Devops, and Platform Engineering.
You understand the difference between Hashicorps Enterprise and Open Source versions of their products and can spek to the differences.

## Question
How would I automate the promotion of a secendary DR vault cluster using the Vault API?
List the relevant API endpoints.
Give curl examples of what successful API requests would look like.
Highlight any potential gotchas.

## Additional Context
This endpoint prints information about the status of replication (mode, sync progress, etc).

This is an unauthenticated endpoint.

Method	Path
GET	/sys/replication/dr/status

This endpoint enables DR replication in primary mode. This is used when DR replication is currently disabled on the cluster (if the cluster is already a secondary, it must be promoted).

Method	Path
POST	/sys/replication/dr/primary/enable

This endpoint demotes a DR primary cluster to a secondary. This DR secondary cluster will not attempt to connect to a primary (see the update-primary call), but will maintain knowledge of its cluster ID and can be reconnected to the same DR replication set without wiping local storage.

Method	Path
POST	/sys/replication/dr/primary/demote

This endpoint disables DR replication entirely on the cluster. Any secondaries will no longer be able to connect. Caution: re-enabling this node as a primary or secondary will change its cluster ID; in the secondary case this means a wipe of the underlying storage when connected to a primary, and in the primary case, secondaries connecting back to the cluster (even if they have connected before) will require a wipe of the underlying storage.

Method	Path
POST	/sys/replication/dr/primary/disable


This endpoint generates a DR secondary activation token for the cluster with the given opaque identifier, which must be unique. This identifier can later be used to revoke a DR secondary's access.

This endpoint requires 'sudo' capability.

Method	Path
POST	/sys/replication/dr/primary/secondary-token

This endpoint revokes a DR secondary's ability to connect to the DR primary cluster; the DR secondary will immediately be disconnected and will not be allowed to connect again unless given a new activation token. This command can be run from any node on the DR primary cluster.

Method	Path
POST	/sys/replication/dr/primary/revoke-secondary

This endpoint allows fetching a public key that is used to encrypt the returned credential information (instead of using a response wrapped token). This avoids needing to make an API call to the primary during activation.

Method	Path
POST	/sys/replication/dr/secondary/generate-public-key

Enable DR Secondary
This endpoint enables replication on a DR secondary using a DR secondary activation token.

This will immediately clear all data in the secondary cluster!

Method	Path
POST	/sys/replication/dr/secondary/enable

This endpoint promotes the DR secondary cluster to DR primary. For data safety and security reasons, new secondary tokens will need to be issued to other secondaries, and there should never be more than one primary at a time.

If the DR secondary's primary cluster is also in a performance replication set, the DR secondary will be promoted into that replication set. Care should be taken when promoting to ensure multiple performance primary clusters are not active at the same time.

If the DR secondary's primary cluster is a performance secondary, the promoted cluster will attempt to connect to the performance primary cluster using the same secondary token.

This endpoint requires a DR Operation Token to be provided as means of authorization. See the DR Operation Token API docs for more information.

Only one performance primary should be active at a given time. Multiple primaries may result in data loss!

It is not safe to replicate from a newer version of Vault to an older version. When upgrading replicated clusters, ensure that upstream clusters are always on older versions of Vault than downstream clusters. See Upgrading Vault for an example.

Method	Path
POST	/sys/replication/dr/secondary/promote

This endpoint disables DR replication entirely on the cluster. The cluster will no longer be able to connect to the DR primary.

This endpoint requires a DR Operation Token to be provided as means of authorization. See the DR Operation Token API docs for more information.

Re-enabling this node as a DR primary or secondary will change its cluster ID; in the secondary case this means a wipe of the underlying storage when connected to a primary, and in the primary case, secondaries connecting back to the cluster (even if they have connected before) will require a wipe of the underlying storage.

Method	Path
POST	/sys/replication/dr/secondary/disable

Update DR Secondary's Primary
This endpoint changes a DR secondary cluster's assigned primary cluster using a secondary activation token. This does not wipe all data in the cluster.

This endpoint requires a DR Operation Token to be provided as means of authorization. See the DR Operation Token API docs for more information.

Method	Path
POST	/sys/replication/dr/secondary/update-primary

Generate Disaster Recovery Operation Token
The /sys/replication/dr/secondary/generate-operation-token endpoint is used to create a new Disaster Recovery operation token for a DR secondary. These tokens are used to authorize certain DR Operations. They should be treated like traditional root tokens by being generated when needed and deleted soon after.

Read Generation Progress
This endpoint reads the configuration and process of the current generation attempt.

Method	Path
GET	/sys/replication/dr/secondary/generate-operation-token/attempt

Start Token Generation
This endpoint initializes a new generation attempt. Only a single generation attempt can take place at a time.

Method	Path
POST	/sys/replication/dr/secondary/generate-operation-token/attempt

Cancel Generation
This endpoint cancels any in-progress generation attempt. This clears any progress made. This must be called to change the OTP or PGP key being used.

Method	Path
DELETE	/sys/replication/dr/secondary/generate-operation-token/attempt

Provide Key Share to Generate Token
This endpoint is used to enter a single root key share to progress the generation attempt. If the threshold number of root key shares is reached, Vault will complete the generation and issue the new token. Otherwise, this API must be called multiple times until that threshold is met. The attempt nonce must be provided with each call.

Method	Path
POST	/sys/replication/dr/secondary/generate-operation-token/update

Delete DR Operation Token
This endpoint revokes the DR Operation Token. This token does not have a TTL and therefore should be deleted when it is no longer needed.

Method	Path
POST	/sys/replication/dr/secondary/operation-token/delete

## Official Documentation Instructions from Hashicorp
To successfully follow this tutorial, you will deploy 2 single-node Vault Enterprise clusters with integrated storage:

Cluster A is the initial primary cluster.
Cluster B is the initial secondary cluster.

In the current state, cluster A is the primary and replicates data to the secondary cluster B. You will perform the following actions to failover so that cluster B becomes the new primary cluster.

Generate batch DR operation token on cluster A.
Promote DR cluster B to become new primary.
Demote cluster A to become secondary.
Point cluster A to new primary cluster B.
Test access to Vault data while cluster B is the primary.
Failback to original primary clusters

In the current state, cluster B is the primary and replicates data to the secondary cluster A. You will perform the following actions to failback to the original cluster replication state.

Generate secondary token on cluster A.
Promote cluster A.
Demote cluster B.
Point cluster B to cluster A, so cluster B is a DR secondary of cluster A.
Test access to Vault data while cluster A is the primary cluster.

Enable DR primary replication on cluster A.
Generate secondary token on cluster A.
Enable DR secondary replication on cluster B.
Confirm replication status on both clusters.

Failover scenario
The goal of this section is to failover the current primary cluster A, and then promote the current secondary cluster B to become the new primary cluster.

You will also validate access to your secret data from the newly promoted primary, and update cluster A, setting cluster B as its new primary.

Take a snapshot
Before proceeding with any failover or failback, it's critical that you have a recent backup of the Vault data. Since the scenario environment uses Vault servers with Integrated Storage, you can take a snapshot of the cluster A Vault data, and write it to cluster-a/vault-cluster-a-snapshot.snap as a backup.

After confirming replication status and taking a snapshot of Vault data, you are ready to begin the failover workflow.

Batch disaster recovery operation token strategy
To promote a DR secondary cluster to be the new primary, a DR operation token is typically needed. However, the process of generating a DR operation token requires a threshold of unseal keys or recovery keys if Vault uses auto unseal. This can be troublesome since a cluster failure is usually caused by unexpected incident. You find difficulty in coordinating amongst the key holders to generate the DR operation token in a timely fashion.

As of Vault 1.4, you can create a batch DR operation token that you can use to promote and demote clusters as needed. This is a strategic operation that the Vault administrator can use to prepare for loss of the DR primary ahead of time. The batch DR operation token also has the advantage of being usable from the primary or secondary more than once.

Promote cluster B to primary status
The first step in this failover workflow is to promote cluster B as a primary.

While you can demote cluster A before promoting cluster B, in production DR scenarios you might instead promote cluster B before demoting cluster A due to unavailability of cluster A.

## Answer

```

## GPT4 Response

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
