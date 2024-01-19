.PHONY: run
run:
	docker compose -f docker-compose.yml up --build

.PHONY: destroy
destroy:
	docker compose -f docker-compose.dev.yml -f docker-compose.test.yml down --volumes --remove-orphans
	rm -rf mysql/mysql-data

.PHONY: tools
tools:
	go install github.com/cosmtrek/air@latest
	go install github.com/golang/mock/mockgen@latest
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.54.1 

.PHONY: test
test:
	docker compose -f docker-compose.test.yml run --rm test-app go test -v ./...