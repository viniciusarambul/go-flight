package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	repository2 "github.com/viniciusarambul/go-flight/app/infrastructure/repository"
	service2 "github.com/viniciusarambul/go-flight/app/usecase/plane"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	//w.Header().Set("Content-Type", "application/json")

	db, err := sql.Open("mysql", "root:root@tcp(docker.for.mac.localhost:3306)/goflight")
	if err != nil {
		log.Fatalln(err)
	}

	repository := repository2.PlaneMySQLRepository{Db: db}
	service := service2.ListPlanes{Repository: repository}

	output, err := service.Exec()
	if err != nil {
		fmt.Println("error: ", err)
	} else {
		fmt.Println("else:", output)
	}
}
func main() {

	http.HandleFunc("/", indexHandler)
	http.ListenAndServe(":8080", nil)
}
