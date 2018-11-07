.PHONY: help

status:
		docker-compose ps

down:
		docker-compose down

migrate:
		docker-compose exec users_service go run db/dbmigrate.go