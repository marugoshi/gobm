up:
	docker-compose up -d

down:
	docker-compose down

restart:
	make down
	make up

.PHONY: images
images:
	docker-compose build --no-cache

dep:
	docker-compose exec gobm bash -c 'dep ensure'

.PHONY: build
build:
	make clean
	docker-compose exec gobm bash -c './build.sh'

exec:
	docker-compose exec gobm bash

run:
	docker-compose exec gobm bash -c 'go run main.go'

reset_db:
	docker-compose exec gobm bash -c 'mysql -h mysql -uroot -ppassword < ./sql/init.sql'

migrate_up:
	docker-compose exec gobm bash -c 'migrate -path sql/migrations -database mysql://root:password@tcp\(mysql:3306\)/gobm_d up'

migrate_create:
	docker-compose exec gobm bash -c 'migrate create -dir sql/migrations -ext sql ${NAME}'

.PHONY: mysql
mysql:
	docker-compose exec gobm bash -c 'mysql -h mysql -uroot -ppassword gobm_d'

ps:
	docker-compose ps

logs:
	docker-compose logs -f

fmt:
	docker-compose exec gobm bash -c 'go fmt ./...'

.PHONY: test
test:
	docker-compose exec gobm bash -c 'go test ./...'
