package functions

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// book struct (Model)
type Book struct {
	ID     string  `json:"id"`
	Isbn   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

type Author struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

// init books var as a slice book struct
var books = []Book{
	Book{
		ID:     "1",
		Isbn:   "44874",
		Title:  "Book 1: Harry Potter",
		Author: &Author{Firstname: "JK", Lastname: "Rowling"},
	},
	Book{
		ID:     "2",
		Isbn:   "34872",
		Title:  "Book 2: Lord of the rings",
		Author: &Author{Firstname: "JRR", Lastname: "Tolkien"},
	},
}

func BaseRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

func GetBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

func GetBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // get params
	// loop through books and find id
	for _, item := range books {
		if item.ID == params["id"] {
			fmt.Println(item)
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(Book{})
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book Book
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		// panic(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"status": 500, "error": "something happened"}`))
		return
	}
	book.ID = strconv.Itoa(rand.Intn(10000000))
	fmt.Println(book)
	books = append(books, book)
	json.NewEncoder(w).Encode(books)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // get params

	for i, item := range books {
		if item.ID == params["id"] {
			books = append(books[:i], books[i+1:]...)
			break
		}
	}
	fmt.Println(books)
	json.NewEncoder(w).Encode(books)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // get params
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)

	for i, item := range books {
		if item.ID == params["id"] {
			book.ID = params["id"]
			books = append(books[:i], books[i+1:]...)
			books = append(books, book)
			break
		}
	}
	fmt.Println(books)
	json.NewEncoder(w).Encode(books)
}
