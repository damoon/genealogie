# Backup bucket to bucket

Create a backup of a bucket and restore at a later time.

## Problem description

Rook, Ceph and Kubernetes operate in a high available manner and replicate data.

*Raid (Data replication) is not a backup* and does not protect against data loss.

Possible failure scenarious are:
- human error
- ceph can fail
- rook can fail
- kubernetes can fail
- hardware can fail

`bucket-backup` creates backups from buckets and stores the snapshots in external buckets.

## Features

- none

## Wishlist

- create snapshots
- list snapshots
- list files in snapshot
- list snapshots of file
- restore bucket
- restore file
- skip unchanged files
- trim snapshots
- parity data
- scrub
- compression
- deduplication
- encryption
- helm chart

## Contribute

1. start minio via `docker-compose up`
2. verify tests via `go test ./...`
