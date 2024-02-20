.PHONY: help
include .env

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
orm:
	sqlboiler mysql -c config/database.toml -o model -p model --no-tests --wipe
migrate:
	migrate create -ext sql -dir app/infrastructure/migrations/ -seq add_timestamps_to_users_and_todos
mup:
	migrate -path app/infrastructure/migrations/ -database 'mysql://${MYSQL_USER}:${MYSQL_PASSWORD}@tcp(127.0.0.1:3306)/${MYSQL_DATABASE}' -verbose up
mdown:
	migrate -path app/infrastructure/migrations/  -database 'mysql://${MYSQL_USER}:${MYSQL_PASSWORD}@tcp(127.0.0.1:3306)/${MYSQL_DATABASE}' -verbose down
add:
	docker compose exec -it app sh -c "go get ${name}"
app-con:
	docker compose exec app ash