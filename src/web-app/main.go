package main

import (
	"fmt"
	"log"
	"os"
	"net/http"
	"github.com/rs/cors"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB
var err error

type Character struct {
	gorm.Model
	Name string
	Class string
}
type Item struct {
	gorm.Model
	Name string
	ItemType string

}

func AllCharacters(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	fmt.Fprintf(w, "all characters endpoint hit")
}

func AllItems(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	db, err = gorm.Open("postgres", os.Getenv("ELEPHANT_URL"))
	if err != nil {
		panic("Could not connect to the database")
	}
	defer db.Close()
	
	var items []Item
	db.Find(&items)
	json.NewEncoder(w).Encode(items)
}

func AddItems(w http.ResponseWriter, r *http.Request) {
	db, err = gorm.Open("postgres", os.Getenv("ELEPHANT_URL"))
	if err!= nil {
		panic("Could not connect to the DB!")
	}
	defer db.Close()
	vars := mux.Vars(r)
	name := vars["name"]
	itemtype := vars["itemtype"]
	
	db.Create(&Item{Name:name, ItemType:itemtype})
	
	fmt.Fprintf(w, "New item added successfully")
}

func RemoveItems(w http.ResponseWriter, r *http.Request) {
	db, err = gorm.Open("postgres", os.Getenv("ELEPHANT_URL"))
	if err!= nil {
		panic("Could not connect to the DB!")
	}
	defer db.Close()
	
	vars := mux.Vars(r)
	name := vars["name"]
	
	var item Item 
	db.Where("name = ?", name).Find(&item)
	db.Delete(&item)
	fmt.Fprintf(w,"item successfully deleted")
}

func UpdateItems(w http.ResponseWriter, r *http.Request) {
	db, err = gorm.Open("postgres", os.Getenv("ELEPHANT_URL"))
	if err!= nil {
		panic("Could not connect to the DB!")
	}
	defer db.Close()
	
	vars := mux.Vars(r)
	name := vars["name"]
	itemtype := vars["itemtype"]
	
	var item Item
	db.Where("name = ?", name).Find(&item)
	item.ItemType = itemtype
	db.Save(&item)
	fmt.Fprintf(w,"Successfully updated")
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow_origin", "*")
}
func InitialMigration() {
	db, err = gorm.Open("postgres", os.Getenv("ELEPHANT_URL"))
	
	if err != nil {
		fmt.Println(err.Error())
		panic("Failed to connect to database")
	}
	defer db.Close()

	db.AutoMigrate(&Item{})
}

func HandleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", AllCharacters).Methods("GET")
	myRouter.HandleFunc("/all/", AllItems).Methods("GET")
	myRouter.HandleFunc("/add/{itemtype}/{name}", AddItems).Methods("POST")
	myRouter.HandleFunc("/remove/{name}", RemoveItems).Methods("DELETE")
	myRouter.HandleFunc("/update/{itemtype}/{name}", UpdateItems).Methods("PUT")
	handler := cors.Default().Handler(myRouter)
	log.Fatal(http.ListenAndServe(":8081", handler))
}


func main() {
	err := godotenv.Load()
  	if err != nil {
    	log.Fatal("Error loading .env file")
	  }
	  
	fmt.Println("Kickoff")
	InitialMigration()

	HandleRequests()
}