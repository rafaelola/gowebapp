package main

import (
	"GoWebApp/database"
	"GoWebApp/models"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net"
	"net/http"
	"text/template"
)

func main() {

	// connect to the database
	db := database.Connect()
	log.Println("DB Connected in main.")
	//jsonInventories := service.GetInventory(db)

	http.HandleFunc("/inventories", func(w http.ResponseWriter,
		r *http.Request) {
		inventories, err := models.All(db)
		if err != nil {
			log.Print(err)
			http.Error(w, http.StatusText(500), 500)
			return
		}
		//unformattedJson := string(jsonInventories)
		t, _ := template.New("foo").Parse(`{{define "T"}}Hello, {{.}}!{{end}}`)
		_ = t.ExecuteTemplate(w, "T", inventories)
		log.Println("Result html")

		log.Println("After logging inventories")
	})
	log.Println("Starting my GO server....")

	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	err = http.Serve(listener, nil)
	if err != nil {
		return // return from main
	}

}
