# Backup bucket to bucket

Create a backup of a bucket and restore at a later time.

This is designed to backup a bucket to another bucket.  
The reasoning is:
  - object versioning does not protect against bucket deletion
  - buckets provided via rook/ceph running in kubernetes are a complex system, software bugs can lead to data loss

## Contribute

The project uses `golang`.

For CI `github actions` are used.  

For local development `tilt` sets up and live reloads an environment.
