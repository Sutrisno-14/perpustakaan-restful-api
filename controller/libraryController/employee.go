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

func CreateEmployee(w http.ResponseWriter, r *http.Request) {
	var employee models.Employee
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&employee); err != nil {
		ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	defer r.Body.Close()
	if err := models.DB.Create(&employee).Error; err != nil {
		ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	Responsejson(w, http.StatusOK, employee)
	log.Printf("Succes create employee id :%d", employee.Id)
}

func GetEmployees(w http.ResponseWriter, r *http.Request) {
	employee := []models.Employee{}

	if err := models.DB.Find(&employee).Error; err != nil {
		ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	Responsejson(w, http.StatusOK, employee)
}

func GetEmployeeById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)

	if err != nil {
		ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}

	employee := models.Employee{}
	employee.Id = id
	if err :=  models.DB.First(&employee, employee.Id).Error; err != nil {
		message := fmt.Sprintf("Data with id :%d not found", employee.Id)
		ResponseError(w, http.StatusNotFound, message)
		return
	}

	Responsejson(w, http.StatusOK, employee)
}

func UpdateEmployee(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}

	employee := models.Employee{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&employee); err != nil {
		ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}
	
	employee.Id = id
	defer r.Body.Close()
	if models.DB.Updates(&employee).Where("id = ?", employee.Id).RowsAffected == 0 {
		message := fmt.Sprintf("Data with id :%d not found", employee.Id )
		ResponseError(w, http.StatusNotFound, message)
		return
	}
	
	Responsejson(w, http.StatusOK, employee)
	log.Printf("Succes update data with id :%d", employee.Id)
}

func DeleteEmployee(w http.ResponseWriter, r *http.Request) {
	input := map[string]int{"id":0}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&input); err != nil {
		ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}
	employee := models.Employee{}
	employee.Id = int64(input["id"])
	if models.DB.Delete(&employee, employee.Id).RowsAffected == 0 {
		message := fmt.Sprintf("Data with id :%d not found", employee.Id)
		ResponseError(w, http.StatusNotFound, message)
		return
	}

	message := map[string]string{"message":"Success delete data employee"}
	Responsejson(w, http.StatusOK, message)
	log.Printf("Success delete data with id :%d", employee.Id)
}