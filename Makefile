.PHONY: help
#include .env

help:
	@grep "^[a-zA-Z\-]*:" Makefile | grep -v "grep" | sed -e 's/^/make /' | sed -e 's/://'

# go
run:
	docker compose exec app sh -c "go run ./cmd/main.go"
test: lint
	cd app && go test ./...
lint:
	cd app && go vet ./...
gen:
	docker compose exec app sh -c "go generate ./..."
tidy:
	docker compose exec app sh -c "go mod tidy"

# 環境
up:
	docker compose up -d
ps:
	docker compose ps
stop:
	docker compose stop
down:
	docker compose down
logs:
	docker compose logs -f

cp-env:
	cp .env.example .env
db:
	docker compose exec -it db mysql -u ${MYSQL_USER} -p${MYSQL_PASSWORD} -D ${MYSQL_DATABASE}
add:
	docker compose exec -it app sh -c "go get ${name}"
app-con:
	docker compose exec app ash


##################
##### DB関連 #####
##################
# マイグレーション
build-cli: # cliのビルド
	cd app && go build -o ./cli/main ./cli/main.go

migrate-dry-run: up build-cli # migration dry-run
	shema_path=$$(find . -name "schema.sql"); \
	./app/cli/main migration $$shema_path
	cd app && rm ./cli/main

migrate-apply: up build-cli # migration apply
	shema_path=$$(find . -name "schema.sql"); \
	DB_HOST=$(DB_HOST) ./app/cli/main migration $$shema_path apply
	cd app && rm ./cli/main

migrate-local-dry-run:
	@make migrate-dry-run DB_HOST=127.0.0.1

migrate-local-apply:
	@make migrate-apply DB_HOST=127.0.0.1

sqlc-gen:
	docker compose exec app sh -c "sqlc generate"