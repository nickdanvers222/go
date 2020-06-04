package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"net/http"
	"github.com/joho/godotenv"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB
var err error

type Page struct {
	Title string
	Body []byte
}
type Character struct {
	gorm.Model
	Name string
	Class string
}


func (p *Page) save() error {
    filename := p.Title + ".txt"
    return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page,error) {
    filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
    return &Page{Title: title, Body: body}, nil
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view/"):]
	p, _ := loadPage(title)
	fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)
}

func initialMigration() {
	godotenv.Load()
	db, err = gorm.Open("postgres", os.Getenv("ELEPHANT_URL"))
	
	if err != nil {
		fmt.Println(err.Error())
		panic("Failed to connect to database")
	}
	defer db.Close()

	db.AutoMigrate(&Character{})
}

func main() {
	fmt.Println("Kickoff")
	initialMigration()

	http.HandleFunc("/view/", viewHandler)
    log.Fatal(http.ListenAndServe(":8081", nil))
}