package main

import (
	"GoWebApp/database"
	"GoWebApp/models"
	"GoWebApp/utils"
	"encoding/json"
	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"log"
	"net"
	"net/http"
	"strconv"
)

var validate *validator.Validate

func main() {
	r := mux.NewRouter()

	// connect to the database
	db := database.Connect()

	// Initialize the validator
	validate = validator.New()

	log.Println("DB Connected in main.")
	//path prefix
	apirouter := r.PathPrefix("/api/v1").Subrouter()

	apirouter.HandleFunc("/inventories/all/", func(w http.ResponseWriter,
		r *http.Request) {
		inventories, err := models.All(db)
		if err != nil {
			log.Print(err)
			utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}
		utils.RespondWithJson(w, http.StatusOK, inventories)
		log.Println("Result inventories/all/")

	}).Methods("GET")

	apirouter.HandleFunc("/inventories/one/{id:[0-9]+}", func(w http.ResponseWriter,
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

	}).Methods("GET")

	apirouter.HandleFunc("/inventories/create/", func(w http.ResponseWriter,
		r *http.Request) {
		// create a new instance of inventory
		var inventory models.InventoryModel

		// Decode the JSON request body into the inventory instance
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&inventory)
		if err != nil {
			log.Print(err)
			utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
			return
		}
		// Validate the inventory instance
		err = inventory.Validate()

		if err != nil {
			log.Print(err)
			utils.RespondWithError(w, http.StatusUnprocessableEntity, err.Error())
			return
		}
		// Create the inventory record
		err = models.Create(db, (*models.Inventory)(&inventory))
		if err != nil {
			log.Print(err)
			utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}
		// return the new inventory ID with 201 Created
		utils.RespondWithJson(w, http.StatusCreated, inventory.Id)
		log.Println("Result inventories/add/")
	}).Methods("POST")

	apirouter.HandleFunc("/inventories/update/{id:[0-9]+}", func(w http.ResponseWriter,
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
		// get request body and json decode it
		decoder := json.NewDecoder(r.Body)
		var inventory models.InventoryModel
		err = decoder.Decode(&inventory)
		if err != nil {
			log.Print(err)
			utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
			return
		}
		// update the inventory in the database
		err = models.Update(db, id, (*models.Inventory)(&inventory))
		if err != nil {
			log.Print(err)
			utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}
		// return the new inventory ID with 201 Created
		utils.RespondWithJson(w, http.StatusOK, inventory.Id)
		log.Println("Result inventories/update/")
	}).Methods("PUT")

	log.Println("Starting my GO server....")

	http.Handle("/", apirouter)

	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	err = http.Serve(listener, nil)
	if err != nil {
		return // return from main
	}
	//defer db.Close()
}
