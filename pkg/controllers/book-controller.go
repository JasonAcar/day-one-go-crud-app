package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/JasonAcar/test-crud-app/pkg/models"
	"github.com/JasonAcar/test-crud-app/pkg/utils"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

var NewBook models.Book

func CreateBook(w http.ResponseWriter, r *http.Request) {
	createBook := &models.Book{}
	utils.ParseBody(r, createBook)
	b := createBook.CreateBook()
	res, _ := json.Marshal(b)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(res)
}

func GetBooks(w http.ResponseWriter, r *http.Request) {
	newBooks := models.GetAllBooks()
	res, _ := json.Marshal(newBooks)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBookByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing book ID to fetch")
	}
	bookDetails, _ := models.GetBookByID(ID)
	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	var updateBook = &models.Book{}
	utils.ParseBody(r, updateBook)
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing book ID to update")
	}
	existingBook, db := models.GetBookByID(ID)
	if updateBook.Name != "" {
		existingBook.Name = updateBook.Name
	}
	if updateBook.Author != "" {
		existingBook.Author = updateBook.Author
	}
	if updateBook.Publication != "" {
		existingBook.Publication = updateBook.Publication
	}
	db.Save(&existingBook)
	res, _ := json.Marshal(existingBook)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid book ID")
		return
	}

	book, err := models.DeleteBook(ID)
	if err != nil {
		utils.RespondWithError(w, http.StatusNotFound, err.Error())
		return
	}

	res, err := json.Marshal(*book)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Error marshalling JSON")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
