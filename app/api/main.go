package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	//w.Header().Set("Content-Type", "application/json")

	db, err := sql.Open("mysql", "root:root@tcp(docker.for.mac.localhost:3306)/goflight")
	if err != nil {
		log.Fatalln(err)
	}

	var id string
	var model string
	err2 := db.QueryRow("select id, model from planes where id = 1").Scan(&id, &model)

	if err2 != nil {
		log.Fatal(err2)
	}

	fmt.Println(id, model)

	/*
		repository := repository2.PlaneMySQLRepository{Db: db}
		service := service2.ListPlanes{Repository: repository}

		output, err := service.Exec()
		if err != nil {
			fmt.Println("error: ", err)
		} else {
			fmt.Println("else:", output)
		}*/
}
func main() {

	http.HandleFunc("/", indexHandler)
	http.ListenAndServe(":8080", nil)
}
