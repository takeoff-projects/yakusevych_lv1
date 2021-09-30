package main

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"drehnstrom.com/go-pets/petsdb"
	"encoding/json"
	"io/ioutil"
    "github.com/gorilla/mux"
)

var projectID string 

func main() {
	projectID = os.Getenv("GOOGLE_CLOUD_PROJECT")
	if projectID == "" {
		log.Fatal(`You need to set the environment variable "GOOGLE_CLOUD_PROJECT"`)
	}
	log.Printf("GOOGLE_CLOUD_PROJECT is set to %s", projectID)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"

	}

	log.Printf("Port set to: %s", port)

	fs := http.FileServer(http.Dir("assets"))
//	mux := http.NewServeMux()
	router := mux.NewRouter().StrictSlash(true)

	// This serves the static files in the assets folder
//	mux.Handle("/assets/", http.StripPrefix("/assets/", fs))
	router.Handle("/assets/", http.StripPrefix("/assets/", fs))

	// The rest of the routes
 	router.HandleFunc("/", indexHandler)
 	router.HandleFunc("/about", aboutHandler)
 	router.HandleFunc("/pets", getPets).Methods("GET")
	router.HandleFunc("/pet/{id}",getPetbyID).Methods("GET")
	router.HandleFunc("/pets", createPet).Methods("POST")
	//router.HandleFunc("/pets/{id}", deletePet).Methods("DELETE")


	log.Printf("Webserver listening on Port: %s", port)
//	http.ListenAndServe(":"+port, mux)
	http.ListenAndServe(":"+port, router)
}
	
func indexHandler(w http.ResponseWriter, r *http.Request) {
	var pets []petsdb.Pet
	pets, error := petsdb.GetPets()
	if error != nil {
		fmt.Print(error)
	}

	data := HomePageData{
		PageTitle: "Pets Home Page",
		Pets: pets,
	}

	var tpl = template.Must(template.ParseFiles("templates/index.html", "templates/layout.html"))

	buf := &bytes.Buffer{}
	err := tpl.Execute(buf, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println(err.Error())
		return
	}

	buf.WriteTo(w)
	log.Println("Home Page Served")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	data := AboutPageData{
		PageTitle: "About Go Pets",
	}

	var tpl = template.Must(template.ParseFiles("templates/about.html", "templates/layout.html"))

	buf := &bytes.Buffer{}
	err := tpl.Execute(buf, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println(err.Error())
		return
	}

	buf.WriteTo(w)
	log.Println("About Page Served")
}

// HomePageData for Index template
type HomePageData struct {
	PageTitle string
	Pets []petsdb.Pet
}

// AboutPageData for About template
type AboutPageData struct {
	PageTitle string
}

func getPets(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: getPets")
	pets, error := petsdb.GetPets()
	if error != nil {
		fmt.Print(error)
	}
	json.NewEncoder(w).Encode(pets)
}

func getPetbyID(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: getPetbyID")
	vars := mux.Vars(r)
	key := vars["id"]

	fmt.Printf("Key: %s\n", key)
	pets, error := petsdb.GetPetbyId(key)
	if error != nil {
		fmt.Print(error)
	}
	json.NewEncoder(w).Encode(pets)
}

func createPet(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: createPet")
	//newID := uuid.New().String()
	//fmt.Println(newID)

	reqBody, _ := ioutil.ReadAll(r.Body)
	var pet petsdb.Pet
	json.Unmarshal(reqBody, &pet)
	//pet.id = newID

	petsdb.CretaePet(pet)

	json.NewEncoder(w).Encode(pet)
}