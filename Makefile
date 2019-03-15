up:
	docker-compose up -d

down:
	docker-compose down

restart:
	make down
	make up

images:
	docker-compose build --no-cache

dep:
	docker-compose exec gobm bash -c 'dep ensure'

build:
	make clean
	docker-compose exec gobm bash -c './build.sh'

exec:
	docker-compose exec gobm bash

mysql:
	docker-compose exec gobm bash -c 'mysql -h mysql -uroot -ppassword gobm_d'

ps:
	docker-compose ps

logs:
	docker-compose logs -f

fmt:
	docker-compose exec gobm bash -c 'go fmt ./...'

.PHONY: mysql test build dep builds deploy