pullmysql:
	docker pull mysql:8.0

runmysql:
	docker run --name notifier-mysql -p 33060:3306 -e MYSQL_ROOT_PASSWORD=password -d mysql:8.0

createdb:
	docker exec -i notifier-mysql mysql -uroot -ppassword  <<< "CREATE DATABASE notifier_db; USE notifier_db;" 2> /dev/null

dropdb:
	docker exec -i notifier-mysql mysql -uroot -ppassword  <<< "DROP DATABASE notifier_db;" 2> /dev/null

up:
	migrate -path db/migrations -database "mysql://root:password@tcp(localhost:33060)/notifier_db" -verbose up

down: 
	migrate -path db/migrations -database "mysql://root:password@tcp(localhost:33060)/notifier_db" -verbose down

.PHONY: pullmysql runmysql createdb dropdb up down