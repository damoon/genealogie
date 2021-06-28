#!/bin/bash
set -euo pipefail

MINIO_VERSION=RELEASE.2021-06-17T00-10-46Z

curl --fail -o docker-compose.yaml https://raw.githubusercontent.com/minio/minio/${MINIO_VERSION}/docs/orchestration/docker-compose/docker-compose.yaml

curl --fail -o nginx.conf https://raw.githubusercontent.com/minio/minio/${MINIO_VERSION}/docs/orchestration/docker-compose/nginx.conf
