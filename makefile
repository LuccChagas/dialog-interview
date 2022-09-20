generate-docs: 
	swag init -g src/router/handler.go --output docs/app

start-server:
	go run src/cmd/main.go

insert-csv:
	go run src/cmd/csv/main.go

database-up:
	docker-compose up

database-down:
	docker-compose up
