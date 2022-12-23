package librarycontroller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/Sutrisno-14/perpustakaan-restful-api/models"
	"github.com/gorilla/mux"
)

func CreateVisitor(w http.ResponseWriter, r *http.Request) {
	visitor := models.Visitor{}
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&visitor); err != nil {
		ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}
	defer r.Body.Close()
	if err := models.DB.Create(&visitor).Error; err != nil {
		ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	Responsejson(w, http.StatusOK, visitor)
	log.Printf("Success Create data visitor with id :%d", visitor.Id)
}

func GetVisitors(w http.ResponseWriter, r *http.Request) {
	visitors := []models.Visitor{}

	if err := models.DB.Find(&visitors).Error; err != nil {
		ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	Responsejson(w, http.StatusOK, visitors)
}

func GetVisitorById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)

	if err != nil {
		ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	visitor := models.Visitor{}
	visitor.Id = id
	if err := models.DB.First(&visitor, visitor.Id).Error; err != nil {
		message := fmt.Sprintf("Data visitor with id :%d not found", visitor.Id)
		ResponseError(w, http.StatusBadRequest, message)
		return
	}

	Responsejson(w, http.StatusOK, visitor)
}

func UpdateVisitor(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)

	if err != nil {
		ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	visitor := models.Visitor{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&visitor); err != nil {
		ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	visitor.Id = id
	defer r.Body.Close()
	if models.DB.Updates(&visitor).Where("id = ?", visitor.Id).RowsAffected == 0 {
		message := fmt.Sprintf("Data visitor with id :%d not foud", visitor.Id)
		ResponseError(w, http.StatusNotFound, message)
		return
	}

	Responsejson(w, http.StatusOK, visitor)
	log.Printf("Success update data visitor with id :%d", visitor.Id)

}

func DeleteVisitor(w http.ResponseWriter, r *http.Request) {
	input := map[string]int{"id":0}
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&input); err != nil {
		ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	defer r.Body.Close()
	visitor := models.Visitor{}
	visitor.Id = int64(input["id"])
	if models.DB.Delete(&visitor, visitor.Id).RowsAffected == 0 {
		message := fmt.Sprintf("Data visitor id : %d not found", visitor.Id)
		ResponseError(w, http.StatusNotFound, message)
		return
	}

	message := map[string]string{"message":"Success delete data visitor"}
	Responsejson(w, http.StatusOK, message)
	log.Printf("Success delete data visitor with id :%d", visitor.Id)

}