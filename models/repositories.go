package models

import "time"

type Main struct {
	Book     Book
	Employee Employee
	Visitor  Visitor
}
type Book struct {
	Id    		int64  `gorm:"primarykey" json:"id"`
	Name  		string `gorm:"varchar(50)" json:"name"`
	Stock 		int32  `gorm:"integer(100)" json:"stock"`
	CreatedAt   time.Time 
	UpdatedAt   time.Time 
}
type Employee struct {
	Id        int64  `gorm:"primarykey" json:"id"`
	Name      string `gorm:"varchar(50)" json:"name"`
	Position  string `gorm:"varchar(50)" json:"position"`
	Sex       string `gorm:"varchar(1)" json:"sex"`
	CreatedAt time.Time 
	UpdatedAt time.Time 
}
type Visitor struct {
	Id         int64     `gorm:"primarykey" json:"id"`
	Name       string    `gorm:"varchar(50)" json:"name"`
	LoanDate   time.Time `gorm:"type:date" json:"loan_date"`
	Status     string    `gorm:"varchar(20)" json:"status"`
	ReturnDate time.Time `gorm:"type:date"  json:"return_date"`
	CreatedAt  time.Time 
	UpdatedAt  time.Time 
}
