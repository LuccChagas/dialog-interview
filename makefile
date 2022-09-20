generate-docs: 
	swag init -g src/router/handler.go --output docs/app
	"Test"

start-server:
	go run src/cmd/main.go

insert-csv:
	go run src/cmd/csv/main.go
