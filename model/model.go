package model

import "gorm.io/gorm"

type Person struct {
	gorm.Model

	Name  string `json:"name"`
	Email string `gorm:"typevarchar(100); unique_index" json:"Email"`
	Books []Book `json:"book"`
}

type Book struct {
	gorm.Model

	Title      string `json:"title"`
	Author     string `json:"author"`
	CallNumber int    `gorm:"unique_index" json:"call_number"`
	PersonID   int    `json:"person_id"`
}

type People []*Person
type Books []*Book
