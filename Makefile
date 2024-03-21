include .env

.PHONY: test

build:
	docker-compose build

up:
	docker-compose up -d

tidy:
	docker-compose exec ${APP_NAME} go mod tidy

init:
	docker-compose exec ${APP_NAME} go run ./migrate/main.go db init

migrate:
	docker-compose exec ${APP_NAME} go run ./migrate/main.go db migrate

migration:
	docker-compose exec ${APP_NAME} go run ./migrate/main.go db create_go ${name}

api-goget:
	docker-compose exec app go get ${MOD}

gen-moq:
	cd src/domain/repository && go generate

swagger:
	docker-compose exec homework swag init
lint:
	docker-compose exec ${APP_NAME}  go fmt ./...
	docker-compose exec ${APP_NAME}  sh -c 'staticcheck -go 1.0 $$(go list ./... | grep -v "moq/fakerepository")'

api-test:
	docker-compose exec ${APP_NAME} go test -cover ./... -coverprofile=../cover.out

destroy:
	docker-compose down --rmi all --volumes --remove-orphans
