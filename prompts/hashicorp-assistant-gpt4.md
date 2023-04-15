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
