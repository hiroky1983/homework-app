include .env

.PHONY: test

build:
	docker-compose build

up:
	docker-compose up -d

tidy:
	docker-compose exec ${APP_NAME} go mod tidy

api-goget:
	docker-compose exec app go get ${MOD}

lint:
	docker-compose exec app go fmt ./...
	docker-compose exec app sh -c 'staticcheck -go 1.0 $$(go list ./... | grep -v "seeder_api_service/gen")'

api-test:
	docker-compose exec app go test -cover ./... -coverprofile=../cover.out

destroy:
	docker-compose down --rmi all --volumes --remove-orphans
