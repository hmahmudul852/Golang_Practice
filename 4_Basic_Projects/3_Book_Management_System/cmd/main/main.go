package main

import (
	"log"
	"net/http"

	"3_Book_Management_System/pkg/routes"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":9010", r))
}

// go mod init 3_Book_Management_System
// go mod init 3_Book_Management_System
// export GO111MODULE="on"
// go get github.com/gin-gonic/gin
