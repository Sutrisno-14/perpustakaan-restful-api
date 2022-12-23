package librarycontroller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/Sutrisno-14/perpustakaan-restful-api/helper"
	"github.com/Sutrisno-14/perpustakaan-restful-api/models"
	"github.com/gorilla/mux"
)

var Responsejson = helper.ResponJson
var ResponseError = helper.ResponseError

func CreateBook(w http.ResponseWriter, r *http.Request) {
	book := models.Book{}
	decoder := json.NewDecoder(r.Body)
	//mendecode to variabel book
	if err := decoder.Decode(&book); err != nil {
		ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	defer r.Body.Close()
	if err := models.DB.Create(&book).Error; err != nil {
		ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}
	Responsejson(w, http.StatusOK, book)
	log.Printf("Success create book id :%d", book.Id)
}

func GetBooks(w http.ResponseWriter, r *http.Request) {
	var books []models.Book

	if err := models.DB.Find(&books).Error; err != nil {
		ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}
	Responsejson(w, http.StatusOK, books)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}

	var book models.Book
	if err := models.DB.First(&book, id).Error; err != nil {
		ResponseError(w, http.StatusNotFound, "Data not found")
		return
	}
	Responsejson(w, http.StatusOK, book)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)

	if err != nil {
		ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}

	var book models.Book
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&book); err != nil {
		ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	book.Id = id
	defer r.Body.Close()
	if models.DB.Where("id = ?", id).Updates(&book).RowsAffected == 0 {
		message := fmt.Sprintf("Updating is failde, id : %d not found", book.Id)
		ResponseError(w, http.StatusBadRequest, message)
		return
	}

	Responsejson(w, http.StatusOK, book)
	log.Printf("Success update data with id : %d", book.Id)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	input := map[string]string{"id":""}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&input); err != nil {
		ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}

	defer r.Body.Close()
	var book models.Book
	if models.DB.Delete(&book, input["id"]).RowsAffected == 0 {
		message := fmt.Sprintf("Deleting is failed, id :%s not found", input["id"])
		ResponseError(w, http.StatusBadRequest, message)
		return
	}

	message := map[string]string{"message":"Success delete data book"}
	Responsejson(w, http.StatusOK, message)
	log.Printf("Success delete book id :%s", input["id"])
}
