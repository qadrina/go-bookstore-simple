package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/qadrina/go-bookstore-simple/pkg/models"
	"github.com/qadrina/go-bookstore-simple/pkg/utils"
)

var NewBook models.Book

func GetBook(w http.ResponseWriter, r *http.Request) {
	//get all books from DB dan store it in var newBooks
	newBooks := models.GetAllBooks()
	//convert the newbooks to json
	res, _ := json.Marshal(newBooks)
	//add header
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	//print list json that has been retrieved from DB
	w.Write(res)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	//access the request n access the book id from our request
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	// in golang, use blank character -> _ if there's something that's not needed, bcs defining something without using it would cause error
	bookDetails, _ := models.GetBookById(ID)
	//send the json response to user
	res, _ := json.Marshal(bookDetails)
	//add header
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	//sending the response. the res is the json version of the book details found in the db
	w.Write(res)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	// receive json
	createBook := &models.Book{}
	//pass/send it into something the DB would understand
	utils.ParseBody(r, createBook)
	//fungsi dari models create book
	b := createBook.CreateBook()
	//response convert to json
	res, _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	//to be able to access the request that we've received from the users
	vars := mux.Vars(r)
	//access the book id
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	book := models.DeleteBook(ID)
	//send the response to the user
	res, _ := json.Marshal(book)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	var updateBook = &models.Book{}
	//whatever u sent in the request, the new body to be updated in the DB of the same ID, take from json and deserialize it into a diff format
	utils.ParseBody(r, updateBook)
	//get the id
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	bookDetails, db := models.GetBookById(ID)
	if updateBook.Name != "" {
		bookDetails.Name = updateBook.Name
	}
	if updateBook.Author != "" {
		bookDetails.Author = updateBook.Author
	}
	if updateBook.Publication != "" {
		bookDetails.Publication = updateBook.Publication
	}
	db.Save(&bookDetails)
	// tstamp 2:11:00
	// send response (the updated book details) to the user, it has to be in json , it's basically serializing to json
	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
