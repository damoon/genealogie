# Backup bucket to bucket

Create a backup of a bucket and restore at a later time.

## Purpose

Running Rook to operate ceph still needs backups to protect against data loss.  
- ceph can fail
- rook can fail
- kubernetes nodes can become inaccessible

## Features

- none

## Wishlist

- create backup
- list backups
- list files in backup
- list backups of file
- restore bucket
- restore file
- skip unchanged files
- trim backups
- add parity data
- scrub
- compression
- deduplication
- encryption
- helm chart

## Contribute

1. start minio via `docker-compose up`
2. verify tests via `go test ./...`
