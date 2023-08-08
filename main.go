package main

import (
	"GoWebApp/database"
	"GoWebApp/models"
	"GoWebApp/utils"
	"encoding/json"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"log"
	"net"
	"net/http"
	"strconv"
)

func main() {
	r := mux.NewRouter()

	// connect to the database
	db := database.Connect()

	log.Println("DB Connected in main.")

	r.HandleFunc("/inventories/all/", func(w http.ResponseWriter,
		r *http.Request) {
		inventories, err := models.All(db)
		if err != nil {
			log.Print(err)
			utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}
		utils.RespondWithJson(w, http.StatusOK, inventories)
		log.Println("Result inventories/all/")

	})

	r.HandleFunc("/inventories/one/{id:[0-9]+}", func(w http.ResponseWriter,
		r *http.Request) {
		// get the ID from the request
		vars := mux.Vars(r)
		idStr := vars["id"]
		id, err := strconv.Atoi(idStr)

		if err != nil {
			log.Print(err)
			utils.RespondWithError(w, http.StatusBadRequest, "Invalid ID")
			return
		}

		inventories, err := models.One(db, id)
		if err != nil {
			log.Print(err)
			utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}

		utils.RespondWithJson(w, http.StatusOK, inventories)
		log.Println("Result inventories/one/")

	})

	r.HandleFunc("/inventories/add/", func(w http.ResponseWriter,
		r *http.Request) {
		// get request body and json decode it
		decoder := json.NewDecoder(r.Body)
		var inventory models.InventoryModel
		err := decoder.Decode(&inventory)
		if err != nil {
			log.Print(err)
			utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
			return
		}
		err = models.Create(db, (*models.Inventory)(&inventory))
		if err != nil {
			log.Print(err)
			utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}
		// return the new inventory ID with 201 Created
		utils.RespondWithJson(w, http.StatusCreated, inventory.Id)
		log.Println("Result inventories/add/")

	})

	log.Println("Starting my GO server....")

	http.Handle("/", r)

	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	err = http.Serve(listener, nil)
	if err != nil {
		return // return from main
	}
	defer db.Close()
}
