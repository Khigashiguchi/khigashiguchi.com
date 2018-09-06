package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/Khigashiguchi/khigashiguchi.com/api/infrastructure/repository"
	"github.com/Khigashiguchi/khigashiguchi.com/api/infrastructure/router"
)

// Run exec to start http server.
func serve() {
	db, err := repository.NewSqlite3("./tmp/tmp.db")
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to start server: %s\n", err)
		os.Exit(1)
	}

	r := router.New(db)
	port := 8080
	fmt.Fprintf(os.Stdout, "start to run http server at port %d\n", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), r); err != nil {
		fmt.Fprintf(os.Stderr, "failed to start server: %s\n", err)
		os.Exit(1)
	}
}

func main() {
	serve()
}
