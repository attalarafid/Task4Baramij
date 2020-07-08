package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

var db *sql.DB
var err error

//Colours struct
type Colours struct {
	Colors struct {
		Category string `json:"category"`
		Code     struct {
			Hex  string `json:"hex"`
			Rgba string `json:"rgba"`
		} `json:"code"`
		Color string `json:"color"`
		Type  string `json:"type"`
	} `json:"colors"`
	Thumbnail struct {
		Height int64  `json:"height"`
		URL    string `json:"url"`
		Width  int64  `json:"width"`
	} `json:"thumbnail"`
}

func inputColours(w http.ResponseWriter, r *http.Request) {

	body, _ := ioutil.ReadAll(r.Body)

	var request Colours

	if err = json.Unmarshal(body, &request); err != nil {
		fmt.Println("Failed decoding json message")
	}

	Kategori := request.Colors.Category
	Heks := request.Colors.Code.Hex
	RedGBA := request.Colors.Code.Rgba
	Warna := request.Colors.Color
	Tipe := request.Colors.Type
	Tinggi := request.Thumbnail.Height
	ULink := request.Thumbnail.URL
	Lebar := request.Thumbnail.Width

	//Insert to Database
	stmt, err := db.Prepare("INSERT INTO colours (Color,Category,Hex,Rgba,Type,URL,Width,Height) VALUES(?,?,?,?,?,?,?,?)")
	_, err = stmt.Exec(Warna, Kategori, Heks, RedGBA, Tipe, ULink, Tinggi, Lebar)
	if err != nil {
		fmt.Fprintf(w, "Data Duplicate")
	} else {
		fmt.Fprintf(w, "Data Created")
	}
}

type Dessert struct {
	ID    string `json:"id"`
	Image struct {
		Height int64  `json:"height"`
		URL    string `json:"url"`
		Width  int64  `json:"width"`
	} `json:"image"`
	Name      string `json:"name"`
	Thumbnail struct {
		Height int64  `json:"height"`
		URL    string `json:"url"`
		Width  int64  `json:"width"`
	} `json:"thumbnail"`
	Type string `json:"type"`
}

func inputDessert(w http.ResponseWriter, r *http.Request) {

	body, _ := ioutil.ReadAll(r.Body)

	var request Dessert

	if err = json.Unmarshal(body, &request); err != nil {
		fmt.Println("Failed decoding json message")
	}

	dessertID := request.ID
	Tipe := request.Type
	Nama := request.Name
	FotoURL := request.Image.URL
	FotoWidth := request.Image.Width
	FotoHeight := request.Image.Height
	SketsaURL := request.Thumbnail.URL
	SketsaWidth := request.Thumbnail.Width
	SketsaHeight := request.Thumbnail.Height

	//Tugas insert kan ke table Customer
	stmt, err := db.Prepare("INSERT INTO dessert (DessertID,ImageHeight,ImageURL,ImageWidth,Name,ThumbnailHeight,ThumbnailURL,ThumbnailWidth,Type) VALUES(?,?,?,?,?,?,?,?,?)")
	_, err = stmt.Exec(dessertID, Tipe, Nama, FotoURL, FotoWidth, FotoHeight, SketsaURL, SketsaWidth, SketsaHeight)
	if err != nil {
		fmt.Fprintf(w, "Data Duplicate")
	} else {
		fmt.Fprintf(w, "Data Created")
	}
}

func main() {

	db, err = sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/task4")
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	// Init router
	r := mux.NewRouter()

	fmt.Println("Server on :8181")

	// Route handles & endpoints
	r.HandleFunc("/inputColours", inputColours).Methods("POST")
	r.HandleFunc("/inputDessert", inputDessert).Methods("POST")

	// Start server
	log.Fatal(http.ListenAndServe(":8181", r))

}
