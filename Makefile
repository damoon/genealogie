up:
	docker-compose up -d
	until curl --fail http://127.0.0.1:9000/minio/health/cluster; do sleep 1; done

logs:
	docker-compose logs -f

test:
	go test ./...

lint:
	golangci-lint run ./...
