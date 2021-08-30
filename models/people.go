package models

import "database/sql"

type People struct {
	Id int `json:"id"`
	Name string `json:"name"`
	BirthDay string `json:"birth_day"`
	Address string	`json:"address"`
	Phone string	`json:"phone"`
	Sex string	`json:"sex"`
	Job string	`json:"job"`
}
type accessDatabase interface {
	GetBuyId(db *sql.DB,id int) (People,error)
	GetAll(db *sql.DB) ([]People,error)
	CreatePeople(db *sql.DB,people People) error
	UpdatePeople(db *sql.DB, id int, newPeople People) error
	DeletePeople(db *sql.DB, id int) error
}
/*
create table people (
	id int primary key auto_increment,
	name varchar(50),
	birth_day Date,
	address varchar(50),
	phone varchar(15),
	sex varchar(10),
	job varchar(20)
	);*/