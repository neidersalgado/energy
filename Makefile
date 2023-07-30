.PHONY: build-db run-db migrate load-data

build-db:
	docker build -t energy-db -f ./Dockerfile .

run-db:
	docker-compose up -d

migrate:
	docker-compose exec db sh -c 'mysql -uroot -proot energy < /docker-entrypoint-initdb.d/migration.sql'

load-data:
	docker-compose exec db sh -c 'mysql --local-infile=1 -uroot -proot energy < /docker-entrypoint-initdb.d/load_data.sql'

make start:
	go run main.go