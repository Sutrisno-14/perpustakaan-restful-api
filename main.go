package main

import (
	"github.com/Sutrisno-14/perpustakaan-restful-api/models"
	"github.com/Sutrisno-14/perpustakaan-restful-api/router"
)

func main() {
	models.ConnectionDB()
	router.Rooter()
}
