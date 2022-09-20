package main

import (
	"dialog-interview/src/router"

	_ "github.com/swaggo/echo-swagger" // echo-swagger middleware
)

func main() {

	// services.ReadAuthorsCSV()
	router.Handler()
}
