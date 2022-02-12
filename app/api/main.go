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

/*
import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	db := database{"shoes": 50, "socks": 5}

	http.HandleFunc("/list", db.list)
	http.HandleFunc("/price", db.price)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

//!-main

type dollars float32

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

type database map[string]dollars

func (db database) list(w http.ResponseWriter, req *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}

func (db database) price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	if price, ok := db[item]; ok {
		fmt.Fprintf(w, "%s\n", price)
	} else {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
	}
}
*/
